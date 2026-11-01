// SPDX-License-Identifier: MIT

// Package store wraps pgx/v5 and provides the only authorised path for
// reading and writing the blockchain schema tables.
//
// All SQL lives exclusively in this package; no other package may construct
// or execute raw SQL.
package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pat3icki/glint/pkg/blockchain"
)

const uniqueViolation = "23505"

// Store wraps a pgx connection pool and exposes typed methods for the
// blockchain schema.  A Store obtained via WithTx wraps a single transaction.
type Store struct {
	pool *pgxpool.Pool // nil when running in a transaction
	tx   pgx.Tx       // non-nil when running in a transaction
}

// querier is the common interface satisfied by both *pgxpool.Pool and pgx.Tx.
type querier interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func (s *Store) q() querier {
	if s.tx != nil {
		return s.tx
	}
	return s.pool
}

// New creates a new Store backed by a connection pool.
func New(ctx context.Context, dsn string) (*Store, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("store: parse DSN: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("store: connect: %w", err)
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("store: ping: %w", err)
	}
	return &Store{pool: pool}, nil
}

// Close releases the connection pool.
func (s *Store) Close() {
	if s.pool != nil {
		s.pool.Close()
	}
}

// SetActor sets the per-transaction actor used by the audit trigger.
// Must be called inside a WithTx callback.
func (s *Store) SetActor(ctx context.Context, actor string) error {
	_, err := s.q().Exec(ctx, `SET LOCAL app.current_user = $1`, actor)
	if err != nil {
		return fmt.Errorf("store: SetActor: %w", err)
	}
	return nil
}

// WithTx runs fn inside a serialisable transaction.
// If fn returns an error the transaction is rolled back; otherwise it is committed.
func (s *Store) WithTx(ctx context.Context, fn func(tx *Store) error) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("store: acquire conn: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return fmt.Errorf("store: begin tx: %w", err)
	}
	txStore := &Store{tx: tx}
	if err = fn(txStore); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}
	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("store: commit: %w", err)
	}
	return nil
}

// SeedFlavour idempotently upserts all metadata for a plugin into the database.
// Safe to call at every startup; uses ON CONFLICT DO UPDATE so schema changes propagate.
func (s *Store) SeedFlavour(ctx context.Context, p blockchain.Plugin) error {
	return s.WithTx(ctx, func(tx *Store) error {
		abiJSON, _ := json.Marshal(p.ABI())
		cfgJSON, _ := json.Marshal(p.ConfigSchema())

		const upsertFlavour = `
INSERT INTO blockchain.glint_flavours
    (flavour_key, display_name, description, abi, config_schema, schema_version)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (flavour_key) DO UPDATE SET
    display_name   = EXCLUDED.display_name,
    description    = EXCLUDED.description,
    abi            = EXCLUDED.abi,
    config_schema  = EXCLUDED.config_schema,
    schema_version = EXCLUDED.schema_version
RETURNING id`
		var flavourID string
		err := tx.q().QueryRow(ctx, upsertFlavour,
			p.FlavourKey(),
			p.DisplayName(),
			p.Description(),
			abiJSON,
			cfgJSON,
			p.SchemaVersion(),
		).Scan(&flavourID)
		if err != nil {
			return fmt.Errorf("store: SeedFlavour upsert flavour: %w", err)
		}

		for _, et := range p.EventTypes() {
			dataSchemaJSON := json.RawMessage(et.DataSchema)
			abiFrag := json.RawMessage(et.ABIFragment)
			const upsertEvt = `
INSERT INTO blockchain.flavour_event_types
    (flavour_id, event_type, category, topic0, abi_fragment, data_schema, schema_version)
VALUES ($1,$2,$3,$4,$5,$6,$7)
ON CONFLICT (flavour_id, event_type) DO UPDATE SET
    category       = EXCLUDED.category,
    topic0         = EXCLUDED.topic0,
    abi_fragment   = EXCLUDED.abi_fragment,
    data_schema    = EXCLUDED.data_schema,
    schema_version = EXCLUDED.schema_version`
			if _, err = tx.q().Exec(ctx, upsertEvt,
				flavourID, et.Name, et.Category, et.Topic0, []byte(abiFrag), []byte(dataSchemaJSON), et.SchemaVersion,
			); err != nil {
				return fmt.Errorf("store: SeedFlavour upsert event_type %s: %w", et.Name, err)
			}
		}

		for _, ent := range p.EntityTypes() {
			stateSchemaJSON := json.RawMessage(ent.StateSchema)
			const upsertEnt = `
INSERT INTO blockchain.flavour_entity_types
    (flavour_id, entity_type, description, state_schema, schema_version)
VALUES ($1,$2,$3,$4,$5)
ON CONFLICT (flavour_id, entity_type) DO UPDATE SET
    description    = EXCLUDED.description,
    state_schema   = EXCLUDED.state_schema,
    schema_version = EXCLUDED.schema_version`
			if _, err = tx.q().Exec(ctx, upsertEnt,
				flavourID, ent.Name, ent.Description, []byte(stateSchemaJSON), ent.SchemaVersion,
			); err != nil {
				return fmt.Errorf("store: SeedFlavour upsert entity_type %s: %w", ent.Name, err)
			}
		}

		slog.Info("store: seeded flavour", "flavour_key", p.FlavourKey(), "flavour_id", flavourID)
		return nil
	})
}

// InsertEvent inserts a decoded event row idempotently.
// Returns the BIGSERIAL id of the new (or existing) row.
// Returns ErrEventAlreadyIngested if the row already exists.
func (s *Store) InsertEvent(ctx context.Context, e *blockchain.DecodedEvent) (int64, error) {
	topicsJSON, err := json.Marshal(e.Topics)
	if err != nil {
		return 0, fmt.Errorf("store: marshal topics: %w", err)
	}

	const q = `
INSERT INTO blockchain.contract_events
    (contract_id, flavour_id, event_type, schema_version,
     tx_hash, log_index, block_number, block_hash, block_timestamp, emitter, topics, data)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
ON CONFLICT (contract_id, tx_hash, log_index) DO NOTHING
RETURNING id`

	var blockHash *string
	if e.BlockHash != "" {
		blockHash = &e.BlockHash
	}
	var blockTime *interface{}
	if !e.BlockTimestamp.IsZero() {
		v := interface{}(e.BlockTimestamp)
		blockTime = &v
	}

	var id int64
	err = s.q().QueryRow(ctx, q,
		e.ContractID, e.FlavourID, e.EventType, e.SchemaVersion,
		e.TxHash, e.LogIndex, e.BlockNumber, blockHash,
		func() any {
			if e.BlockTimestamp.IsZero() {
				return nil
			}
			return e.BlockTimestamp
		}(),
		e.Emitter, topicsJSON, []byte(e.Data),
	).Scan(&id)
	_ = blockTime
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// ON CONFLICT DO NOTHING fired — already exists.
			return 0, blockchain.ErrEventAlreadyIngested
		}
		return 0, fmt.Errorf("store: InsertEvent: %w", err)
	}
	return id, nil
}

// UpsertEntity applies a single EntityMutation to blockchain.contract_entities.
func (s *Store) UpsertEntity(ctx context.Context, m *blockchain.EntityMutation) error {
	switch m.Op {
	case blockchain.OpUpsert:
		const q = `
INSERT INTO blockchain.contract_entities
    (contract_id, flavour_id, entity_type, entity_key, state, schema_version, last_event_id, last_block)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
ON CONFLICT (contract_id, entity_type, entity_key) DO UPDATE SET
    state          = EXCLUDED.state,
    schema_version = EXCLUDED.schema_version,
    last_event_id  = EXCLUDED.last_event_id,
    last_block     = EXCLUDED.last_block`
		if _, err := s.q().Exec(ctx, q,
			m.ContractID, m.FlavourID, m.EntityType, m.EntityKey,
			[]byte(m.State), m.SchemaVersion, m.LastEventID, m.LastBlock,
		); err != nil {
			return fmt.Errorf("store: UpsertEntity %s/%s: %w", m.EntityType, m.EntityKey, err)
		}
	case blockchain.OpDelete:
		const q = `
DELETE FROM blockchain.contract_entities
WHERE contract_id=$1 AND entity_type=$2 AND entity_key=$3`
		if _, err := s.q().Exec(ctx, q, m.ContractID, m.EntityType, m.EntityKey); err != nil {
			return fmt.Errorf("store: DeleteEntity %s/%s: %w", m.EntityType, m.EntityKey, err)
		}
	default:
		return fmt.Errorf("store: unknown Op %d", m.Op)
	}
	return nil
}

// GetEntity fetches a single entity snapshot.
func (s *Store) GetEntity(ctx context.Context, ref blockchain.EntityRef) (*blockchain.EntitySnapshot, error) {
	const q = `
SELECT contract_id, flavour_id, entity_type, entity_key, state, schema_version, last_event_id, last_block
FROM blockchain.contract_entities
WHERE contract_id=$1 AND entity_type=$2 AND entity_key=$3`
	row := s.q().QueryRow(ctx, q, ref.ContractID, ref.EntityType, ref.EntityKey)
	var snap blockchain.EntitySnapshot
	var state []byte
	err := row.Scan(
		&snap.ContractID, &snap.FlavourID, &snap.EntityType, &snap.EntityKey,
		&state, &snap.SchemaVersion, &snap.LastEventID, &snap.LastBlock,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("store: GetEntity: %w", err)
		}
		return nil, fmt.Errorf("store: GetEntity: %w", err)
	}
	snap.State = state
	return &snap, nil
}

// AdvanceCursor upserts the ingestion cursor for a contract+indexer pair.
func (s *Store) AdvanceCursor(ctx context.Context, contractID, indexerName string, blockNumber int64, blockHash string) error {
	const q = `
INSERT INTO blockchain.event_ingestion_cursors
    (contract_id, indexer_name, last_block_number, last_block_hash)
VALUES ($1,$2,$3,$4)
ON CONFLICT (contract_id, indexer_name) DO UPDATE SET
    last_block_number = EXCLUDED.last_block_number,
    last_block_hash   = EXCLUDED.last_block_hash`
	var bh *string
	if blockHash != "" {
		bh = &blockHash
	}
	if _, err := s.q().Exec(ctx, q, contractID, indexerName, blockNumber, bh); err != nil {
		return fmt.Errorf("store: AdvanceCursor: %w", err)
	}
	return nil
}

// GetCursor returns the last ingested block number and hash for a contract+indexer.
// Returns ErrCursorNotFound if no cursor row exists yet.
func (s *Store) GetCursor(ctx context.Context, contractID, indexerName string) (int64, string, error) {
	const q = `
SELECT last_block_number, COALESCE(last_block_hash,'')
FROM blockchain.event_ingestion_cursors
WHERE contract_id=$1 AND indexer_name=$2`
	var blockNum int64
	var blockHash string
	err := s.q().QueryRow(ctx, q, contractID, indexerName).Scan(&blockNum, &blockHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, "", blockchain.ErrCursorNotFound
		}
		return 0, "", fmt.Errorf("store: GetCursor: %w", err)
	}
	return blockNum, blockHash, nil
}

// isUniqueViolation returns true if err is a Postgres unique-constraint violation.
func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == uniqueViolation
}

// suppress unused warning
var _ = isUniqueViolation
