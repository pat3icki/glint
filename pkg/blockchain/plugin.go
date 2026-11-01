// SPDX-License-Identifier: MIT

// Package blockchain provides the plugin framework for indexing EVM smart contracts.
//
// Plugins are registered at program start via [Register] (typically from a sub-package's
// init function) and looked up at runtime via [GlobalRegistry].
//
// Rule: this package never imports any sub-package; sub-packages never import siblings.
package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Plugin is the interface every smart-contract family must implement.
// All methods must be safe for concurrent use and must not panic.
type Plugin interface {
	// FlavourKey returns the short, stable identifier stored in glint_flavours.flavour_key.
	// Example: "glintdough"
	FlavourKey() string

	// DisplayName returns the human-readable name stored in glint_flavours.display_name.
	DisplayName() string

	// Description returns the optional long description.
	Description() string

	// SchemaVersion returns the current schema version for this plugin.
	// Increment when DataSchema or StateSchema changes in a breaking way.
	SchemaVersion() int

	// ConfigSchema returns a JSON Schema (draft 2020-12) that validates
	// blockchain.contracts.config for this flavour.
	ConfigSchema() json.RawMessage

	// ABI returns the full ABI JSON as a raw byte slice.
	// Used by the store to populate glint_flavours.abi.
	ABI() json.RawMessage

	// EventTypes returns metadata for every event this flavour can emit.
	EventTypes() []EventType

	// EntityTypes returns metadata for every derived-state object this flavour maintains.
	EntityTypes() []EntityType

	// DecodeEvent translates a raw EVM log into a DecodedEvent.
	// It returns [ErrUnknownEvent] if the log's topic0 does not match any event type.
	// Poisoned / malformed logs should return an error; the ingest pipeline will skip them.
	DecodeEvent(ctx context.Context, log gethtypes.Log) (*DecodedEvent, error)

	// Project maps a decoded event to zero or more entity mutations.
	// Mutations are applied atomically within the same database transaction as InsertEvent.
	Project(ctx context.Context, event *DecodedEvent) ([]EntityMutation, error)
}

// Registry is a thread-safe map from FlavourKey to Plugin.
type Registry struct {
	mu      sync.RWMutex
	plugins map[string]Plugin
}

// NewRegistry returns an empty, ready-to-use Registry.
func NewRegistry() *Registry {
	return &Registry{plugins: make(map[string]Plugin)}
}

// Register adds p to the registry.
// It panics if a plugin with the same FlavourKey is already registered —
// consistent with how Go's database/sql drivers and log handlers register themselves.
func (r *Registry) Register(p Plugin) {
	key := p.FlavourKey()
	if key == "" {
		panic("blockchain: plugin FlavourKey must not be empty")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.plugins[key]; exists {
		panic(fmt.Sprintf("blockchain: plugin %q already registered", key))
	}
	r.plugins[key] = p
}

// Lookup returns the plugin for the given flavour key, or ErrPluginNotFound.
func (r *Registry) Lookup(flavourKey string) (Plugin, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.plugins[flavourKey]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrPluginNotFound, flavourKey)
	}
	return p, nil
}

// All returns a snapshot of all registered plugins.
func (r *Registry) All() []Plugin {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Plugin, 0, len(r.plugins))
	for _, p := range r.plugins {
		out = append(out, p)
	}
	return out
}

// GlobalRegistry is the process-wide plugin registry.
// Sub-packages register into this via their init() functions.
var GlobalRegistry = NewRegistry()

// Register is a package-level convenience wrapper around GlobalRegistry.Register.
func Register(p Plugin) { GlobalRegistry.Register(p) }

// Lookup is a package-level convenience wrapper around GlobalRegistry.Lookup.
func Lookup(flavourKey string) (Plugin, error) { return GlobalRegistry.Lookup(flavourKey) }
