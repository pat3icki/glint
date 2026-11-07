// SPDX-License-Identifier: MIT

// Package ingest implements the ordered, per-event pipeline that drives the
// blockchain indexer loop.
//
// Pipeline per log (in order):
//  1. Decode  — plugin.DecodeEvent; poisoned logs are logged and skipped.
//  2. Insert  — store.InsertEvent (idempotent); already-ingested logs are skipped.
//  3. Project — plugin.Project → []EntityMutation.
//  4. Apply   — store.UpsertEntity for each mutation, in the same transaction.
//  5. Advance — store.AdvanceCursor after the batch commits.
package ingest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"sort"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/pat3icki/glint/pkg/blockchain"
	"github.com/pat3icki/glint/pkg/blockchain/chain"
	"github.com/pat3icki/glint/pkg/blockchain/store"
)

// Config holds per-run configuration for the indexer.
type Config struct {
	// ContractID is the UUID of the blockchain.contracts row being indexed.
	ContractID string
	// FlavourID is the UUID of the blockchain.glint_flavours row.
	FlavourID string
	// ContractAddress is the EVM address of the contract being watched.
	ContractAddress string
	// IndexerName identifies this indexer instance (used as cursor key).
	IndexerName string
	// StartBlock is the block to start from if no cursor exists.
	StartBlock int64
	// BlockBatchSize is the number of blocks fetched per FilterLogs call.
	// Defaults to 2000 if zero.
	BlockBatchSize int64
}

// Indexer runs the ingest pipeline for a single contract.
type Indexer struct {
	cfg    Config
	plugin blockchain.Plugin
	chain  *chain.Client
	store  *store.Store
}

// New creates a new Indexer.
func New(cfg Config, p blockchain.Plugin, c *chain.Client, s *store.Store) *Indexer {
	if cfg.BlockBatchSize == 0 {
		cfg.BlockBatchSize = 2000
	}
	return &Indexer{cfg: cfg, plugin: p, chain: c, store: s}
}

// RunOnce fetches all new logs up to the current chain head and processes them.
// It is safe to call repeatedly (poll-based indexing).
func (idx *Indexer) RunOnce(ctx context.Context) error {
	fromBlock, _, err := idx.store.GetCursor(ctx, idx.cfg.ContractID, idx.cfg.IndexerName)
	if err != nil {
		if errors.Is(err, blockchain.ErrCursorNotFound) {
			fromBlock = idx.cfg.StartBlock
		} else {
			return fmt.Errorf("ingest: GetCursor: %w", err)
		} 
	} else {
		fromBlock++ // resume from the block after the last processed one
	}

	latestBlock, err := idx.chain.LatestBlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("ingest: LatestBlockNumber: %w", err)
	}
	if fromBlock > latestBlock {
		slog.Debug("ingest: nothing to do", "contract", idx.cfg.ContractID, "from", fromBlock, "latest", latestBlock)
		return nil
	}

	addr := common.HexToAddress(idx.cfg.ContractAddress)
	var toBlock int64

	for fromBlock <= latestBlock {
		toBlock = fromBlock + idx.cfg.BlockBatchSize - 1
		if toBlock > latestBlock {
			toBlock = latestBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock),
			ToBlock:   big.NewInt(toBlock),
			Addresses: []common.Address{addr},
		}
		logs, err := idx.chain.FilterLogs(ctx, q)
		if err != nil {
			return fmt.Errorf("ingest: FilterLogs [%d,%d]: %w", fromBlock, toBlock, err)
		}

		// Sort by (BlockNumber, LogIndex) for deterministic ordering.
		sort.Slice(logs, func(i, j int) bool {
			if logs[i].BlockNumber != logs[j].BlockNumber {
				return logs[i].BlockNumber < logs[j].BlockNumber
			}
			return logs[i].Index < logs[j].Index
		})

		if err := idx.processLogs(ctx, logs); err != nil {
			return err
		}

		// Advance cursor to end of this batch.
		batchBlockHash := ""
		if len(logs) > 0 {
			last := logs[len(logs)-1]
			batchBlockHash = last.BlockHash.Hex()
		}
		if err := idx.store.AdvanceCursor(ctx, idx.cfg.ContractID, idx.cfg.IndexerName, toBlock, batchBlockHash); err != nil {
			return fmt.Errorf("ingest: AdvanceCursor: %w", err)
		}

		fromBlock = toBlock + 1
	}
	return nil
}

func (idx *Indexer) processLogs(ctx context.Context, logs []ethtypes.Log) error {
	for i := range logs {
		log := logs[i]
		lg := slog.With(
			"tx", log.TxHash.Hex(),
			"log_index", log.Index,
			"block", log.BlockNumber,
		)

		// Step 1: Decode.
		event, err := idx.plugin.DecodeEvent(ctx, log)
		if err != nil {
			if errors.Is(err, blockchain.ErrUnknownEvent) {
				lg.Debug("ingest: skipping unknown event topic")
				continue
			}
			// Poisoned event — log and skip.
			lg.Error("ingest: decode failed (poisoned event, skipping)", "err", err)
			continue
		}
		event.ContractID = idx.cfg.ContractID
		event.FlavourID = idx.cfg.FlavourID

		// Steps 2–4 run inside a single transaction per event.
		if err := idx.store.WithTx(ctx, func(tx *store.Store) error {
			// Step 2: Insert event (idempotent).
			eventID, err := tx.InsertEvent(ctx, event)
			if err != nil {
				if errors.Is(err, blockchain.ErrEventAlreadyIngested) {
					lg.Debug("ingest: event already ingested, skipping")
					return nil
				}
				return fmt.Errorf("InsertEvent: %w", err)
			}

			// Step 3: Project.
			mutations, err := idx.plugin.Project(ctx, event)
			if err != nil {
				return fmt.Errorf("Project: %w", err)
			}

			// Step 4: Apply mutations.
			for j := range mutations {
				m := &mutations[j]
				m.LastEventID = eventID
				m.LastBlock = event.BlockNumber
				if err := tx.UpsertEntity(ctx, m); err != nil {
					return fmt.Errorf("UpsertEntity %s/%s: %w", m.EntityType, m.EntityKey, err)
				}
			}
			return nil
		}); err != nil {
			return fmt.Errorf("ingest: process log tx_hash=%s log_index=%d: %w",
				log.TxHash.Hex(), log.Index, err)
		}
	}
	return nil
}
