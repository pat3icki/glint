// SPDX-License-Identifier: MIT

package glintdough

import (
	"context"
	_ "embed"
	"encoding/json"

	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/pat3icki/glint/pkg/blockchain"
)

//go:embed config_schema.json
var rawConfigSchema []byte

// GlintDoughPlugin is the Plugin implementation for the GlintDough contract family.
type GlintDoughPlugin struct{}

// New returns a ready-to-use GlintDoughPlugin.
func New() *GlintDoughPlugin { return &GlintDoughPlugin{} }

// FlavourKey implements blockchain.Plugin.
func (p *GlintDoughPlugin) FlavourKey() string { return "glintdough" }

// DisplayName implements blockchain.Plugin.
func (p *GlintDoughPlugin) DisplayName() string { return "GlintDough" }

// Description implements blockchain.Plugin.
func (p *GlintDoughPlugin) Description() string {
	return "ERC-1155 collateral-backed gift card contract. " +
		"Each voucher type is backed 1:1 by a registered ERC-20 collateral token."
}

// SchemaVersion implements blockchain.Plugin.
func (p *GlintDoughPlugin) SchemaVersion() int { return 1 }

// ConfigSchema implements blockchain.Plugin.
func (p *GlintDoughPlugin) ConfigSchema() json.RawMessage {
	return json.RawMessage(rawConfigSchema)
}

// ABI implements blockchain.Plugin.
// Returns the full ABI JSON from the generated abigen metadata.
func (p *GlintDoughPlugin) ABI() json.RawMessage {
	return json.RawMessage(abiJSON())
}

// EventTypes implements blockchain.Plugin.
func (p *GlintDoughPlugin) EventTypes() []blockchain.EventType { return eventTypes() }

// EntityTypes implements blockchain.Plugin.
func (p *GlintDoughPlugin) EntityTypes() []blockchain.EntityType { return entityTypes() }

// DecodeEvent implements blockchain.Plugin.
func (p *GlintDoughPlugin) DecodeEvent(ctx context.Context, log ethtypes.Log) (*blockchain.DecodedEvent, error) {
	return decodeEvent(ctx, log)
}

// Project implements blockchain.Plugin.
func (p *GlintDoughPlugin) Project(ctx context.Context, event *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	return project(ctx, event)
}
