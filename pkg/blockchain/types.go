// SPDX-License-Identifier: MIT

package blockchain

import (
	"encoding/json"
	"time"
)

// Op describes how an EntityMutation should be applied.
type Op uint8

const (
	// OpUpsert inserts or updates an entity row.
	OpUpsert Op = iota + 1
	// OpDelete removes an entity row.
	OpDelete
)

// DecodedEvent is the canonical, database-ready representation of a single EVM log.
// All address/hash fields are stored as lower-case hex strings.
type DecodedEvent struct {
	// ContractID is the UUID of the contract row in blockchain.contracts.
	ContractID string
	// FlavourID is the UUID of the flavour row in blockchain.glint_flavours.
	FlavourID string
	// EventType is the name matching a row in blockchain.flavour_event_types.
	EventType string
	// SchemaVersion mirrors flavour_event_types.schema_version.
	SchemaVersion int

	// Chain provenance fields.
	TxHash         string    // 0x-prefixed, lowercase
	LogIndex       int       // position within the transaction
	BlockNumber    int64
	BlockHash      string    // 0x-prefixed, lowercase; may be empty
	BlockTimestamp time.Time // zero if unknown
	Emitter        string    // 0x-prefixed, lowercase contract address

	// Topics contains the raw topic hex strings, index 0 is topic0 (event signature hash).
	Topics []string
	// Data holds the decoded event fields as a JSON-serialisable map.
	Data json.RawMessage
}

// EntityMutation describes a single write to blockchain.contract_entities.
type EntityMutation struct {
	Op           Op
	ContractID   string
	FlavourID    string
	EntityType   string
	EntityKey    string
	State        json.RawMessage
	SchemaVersion int
	// LastEventID is the BIGSERIAL id of the contract_events row that caused this mutation.
	LastEventID int64
	LastBlock   int64
}

// EntitySnapshot is a read-only view of a single blockchain.contract_entities row.
type EntitySnapshot struct {
	ContractID    string
	FlavourID     string
	EntityType    string
	EntityKey     string
	State         json.RawMessage
	SchemaVersion int
	LastEventID   int64
	LastBlock     int64
}

// EntityRef uniquely identifies an entity row.
type EntityRef struct {
	ContractID string
	EntityType string
	EntityKey  string
}

// EventType describes a single event that a Plugin can emit.
type EventType struct {
	// Name matches flavour_event_types.event_type.
	Name string
	// Category is free-form; e.g. "collateral_token", "mint", "transfer".
	Category string
	// Topic0 is keccak256(event signature) in 0x-prefixed, lowercase hex.
	Topic0 string
	// ABIFragment is the single-event ABI fragment JSON (optional; stored for tooling).
	ABIFragment json.RawMessage
	// DataSchema is the JSON Schema (draft 2020-12) for the decoded Data field.
	DataSchema json.RawMessage
	// SchemaVersion matches flavour_event_types.schema_version.
	SchemaVersion int
}

// EntityType describes a derived-state object that a Plugin maintains.
type EntityType struct {
	// Name matches flavour_entity_types.entity_type.
	Name string
	// Description is human-readable.
	Description string
	// StateSchema is the JSON Schema (draft 2020-12) for the entity state object.
	StateSchema json.RawMessage
	// SchemaVersion matches flavour_entity_types.schema_version.
	SchemaVersion int
}


