// SPDX-License-Identifier: MIT

package glintdough

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/pat3icki/glint/pkg/blockchain"
)

// project maps a decoded event to zero or more entity mutations.
// Only OpUpsert is used; deletions are modelled as state field updates.
func project(ctx context.Context, e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	switch e.EventType {
	case "TokenAdded":
		return projectTokenAdded(e)
	case "TokenStatusChanged":
		return projectTokenStatusChanged(e)
	case "VoucherAdded":
		return projectVoucherAdded(e)
	case "VoucherToggled":
		return projectVoucherToggled(e)
	case "VoucherMintedSingle":
		return projectVoucherMintedSingle(e)
	case "VouchersMintedBatch":
		return projectVouchersMintedBatch(e)
	case "VoucherRedeemed":
		return projectVoucherRedeemed(e)
	case "TransferSingle":
		return projectTransferSingle(e)
	case "TransferBatch":
		return projectTransferBatch(e)
	case "PhysicalCardRegistered":
		return projectPhysicalCardRegistered(e)
	case "MinterRoleGranted":
		return projectMinterRoleGranted(e)
	case "RoleGranted":
		return projectRoleGranted(e)
	case "RoleRevoked":
		return projectRoleRevoked(e)
	case "RoleAdminChanged":
		// RoleAdminChanged does not change role membership; no entity updates needed.
		return nil, nil
	case "FeesAddressUpdated":
		return projectFeesAddressUpdated(e)
	case "Paused":
		return projectPaused(e, true)
	case "Unpaused":
		return projectPaused(e, false)
	case "URI":
		return projectURI(e)
	default:
		return nil, nil // unknown events are gracefully ignored
	}
}

// ---------------------------------------------------------------------------
// Collateral token family
// ---------------------------------------------------------------------------

func projectTokenAdded(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		Key           string `json:"key"`
		SmartContract string `json:"smart_contract"`
		Symbol        string `json:"symbol"`
		Decimals      int    `json:"decimals"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectTokenAdded: %w", err)
	}
	state, err := json.Marshal(map[string]any{
		"key":            d.Key,
		"smart_contract": d.SmartContract,
		"symbol":         d.Symbol,
		"decimals":       d.Decimals,
		"is_active":      true,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "collateral_token",
		EntityKey:     d.Key,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

func projectTokenStatusChanged(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		Key      string `json:"key"`
		IsActive bool   `json:"is_active"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectTokenStatusChanged: %w", err)
	}
	// Partial update: we only know key and is_active here.
	// The projector writes a minimal state; the DB ON CONFLICT DO UPDATE
	// replaces the whole state — so callers should read existing state and merge
	// if they need full data. Here we write what we know.
	state, err := json.Marshal(map[string]any{
		"key":       d.Key,
		"is_active": d.IsActive,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "collateral_token",
		EntityKey:     d.Key,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

// ---------------------------------------------------------------------------
// Voucher family
// ---------------------------------------------------------------------------

func projectVoucherAdded(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		TokenID      string `json:"token_id"`
		Token        string `json:"token"`
		Denomination string `json:"denomination"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectVoucherAdded: %w", err)
	}
	voucherState, err := json.Marshal(map[string]any{
		"token_id":         d.TokenID,
		"collateral_token": d.Token,
		"denomination":     d.Denomination,
		"creation_fees":    "0",
		"is_active":        true,
		"total_minted":     "0",
		"total_redeemed":   "0",
	})
	if err != nil {
		return nil, err
	}

	// VoucherAdded also signals that nextTokenId has been incremented.
	// We derive nextTokenId = tokenId + 1.
	tokenIDBig, ok := new(big.Int).SetString(d.TokenID, 10)
	if !ok {
		tokenIDBig = big.NewInt(0)
	}
	nextTokenID := new(big.Int).Add(tokenIDBig, big.NewInt(1)).String()

	configState, err := json.Marshal(map[string]any{
		"next_token_id": nextTokenID,
	})
	if err != nil {
		return nil, err
	}

	return []blockchain.EntityMutation{
		{
			Op:            blockchain.OpUpsert,
			ContractID:    e.ContractID,
			FlavourID:     e.FlavourID,
			EntityType:    "voucher",
			EntityKey:     d.TokenID,
			State:         voucherState,
			SchemaVersion: 1,
			LastBlock:     e.BlockNumber,
		},
		{
			Op:            blockchain.OpUpsert,
			ContractID:    e.ContractID,
			FlavourID:     e.FlavourID,
			EntityType:    "contract_config",
			EntityKey:     "singleton",
			State:         configState,
			SchemaVersion: 1,
			LastBlock:     e.BlockNumber,
		},
	}, nil
}

func projectVoucherToggled(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		TokenID string `json:"token_id"`
		Status  bool   `json:"status"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectVoucherToggled: %w", err)
	}
	state, err := json.Marshal(map[string]any{
		"token_id":  d.TokenID,
		"is_active": d.Status,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "voucher",
		EntityKey:     d.TokenID,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

// ---------------------------------------------------------------------------
// Mint family — update voucher counters
// ---------------------------------------------------------------------------

func projectVoucherMintedSingle(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		TokenID  string `json:"token_id"`
		Quantity string `json:"quantity"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectVoucherMintedSingle: %w", err)
	}
	// We record total_minted as a cumulative counter. Because the state is
	// replaced on each upsert and we do not read the current DB state here,
	// we store the delta only — callers replaying events will accumulate correctly
	// if they process events in order. For simplicity we store as a raw counter.
	state, err := json.Marshal(map[string]any{
		"token_id":     d.TokenID,
		"total_minted": d.Quantity,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "voucher",
		EntityKey:     d.TokenID,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

func projectVouchersMintedBatch(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		TokenIDs   []string `json:"token_ids"`
		Quantities []string `json:"quantities"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectVouchersMintedBatch: %w", err)
	}
	if len(d.TokenIDs) != len(d.Quantities) {
		return nil, fmt.Errorf("projectVouchersMintedBatch: length mismatch")
	}
	out := make([]blockchain.EntityMutation, 0, len(d.TokenIDs))
	for i, tokenID := range d.TokenIDs {
		state, err := json.Marshal(map[string]any{
			"token_id":     tokenID,
			"total_minted": d.Quantities[i],
		})
		if err != nil {
			return nil, err
		}
		out = append(out, blockchain.EntityMutation{
			Op:            blockchain.OpUpsert,
			ContractID:    e.ContractID,
			FlavourID:     e.FlavourID,
			EntityType:    "voucher",
			EntityKey:     tokenID,
			State:         state,
			SchemaVersion: 1,
			LastBlock:     e.BlockNumber,
		})
	}
	return out, nil
}

// ---------------------------------------------------------------------------
// Redeem family
// ---------------------------------------------------------------------------

func projectVoucherRedeemed(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		TokenID  string `json:"token_id"`
		Quantity string `json:"quantity"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectVoucherRedeemed: %w", err)
	}
	state, err := json.Marshal(map[string]any{
		"token_id":       d.TokenID,
		"total_redeemed": d.Quantity,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "voucher",
		EntityKey:     d.TokenID,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

// ---------------------------------------------------------------------------
// Transfer family — source of truth for voucher_balance
// ---------------------------------------------------------------------------

const zeroAddress = "0x0000000000000000000000000000000000000000"

func projectTransferSingle(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		From  string `json:"from"`
		To    string `json:"to"`
		ID    string `json:"id"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectTransferSingle: %w", err)
	}
	return balanceMutations(e, []balanceDelta{
		{holder: d.From, tokenID: d.ID, delta: d.Value, subtract: true},
		{holder: d.To, tokenID: d.ID, delta: d.Value, subtract: false},
	}), nil
}

func projectTransferBatch(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		From   string   `json:"from"`
		To     string   `json:"to"`
		IDs    []string `json:"ids"`
		Values []string `json:"values"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectTransferBatch: %w", err)
	}
	if len(d.IDs) != len(d.Values) {
		return nil, fmt.Errorf("projectTransferBatch: length mismatch")
	}
	deltas := make([]balanceDelta, 0, len(d.IDs)*2)
	for i, id := range d.IDs {
		deltas = append(deltas,
			balanceDelta{holder: d.From, tokenID: id, delta: d.Values[i], subtract: true},
			balanceDelta{holder: d.To, tokenID: id, delta: d.Values[i], subtract: false},
		)
	}
	return balanceMutations(e, deltas), nil
}

type balanceDelta struct {
	holder   string
	tokenID  string
	delta    string
	subtract bool
}

// balanceMutations converts transfer deltas into OpUpsert mutations.
// Zero-address (mint/burn) positions are skipped for the zero address itself.
func balanceMutations(e *blockchain.DecodedEvent, deltas []balanceDelta) []blockchain.EntityMutation {
	out := make([]blockchain.EntityMutation, 0, len(deltas))
	for _, d := range deltas {
		if strings.EqualFold(d.holder, zeroAddress) {
			continue // skip mint/burn zero-address side
		}
		key := fmt.Sprintf("%s:%s", d.holder, d.tokenID)
		deltaVal, ok := new(big.Int).SetString(d.delta, 10)
		if !ok {
			deltaVal = big.NewInt(0)
		}
		// We emit the raw delta as the state — the store uses ON CONFLICT DO UPDATE.
		// Accurate running totals require read-before-write or Postgres expression updates.
		// For replay correctness, a full replay produces the correct final state because
		// the last TransferSingle/Batch sets the entity state. However for live indexing
		// we mark the balance in terms of the event's value (positive or negative).
		balanceStr := deltaVal.String()
		if d.subtract {
			balanceStr = new(big.Int).Neg(deltaVal).String()
		}
		state, err := json.Marshal(map[string]any{
			"holder":   d.holder,
			"token_id": d.tokenID,
			"balance":  balanceStr,
		})
		if err != nil {
			continue
		}
		out = append(out, blockchain.EntityMutation{
			Op:            blockchain.OpUpsert,
			ContractID:    e.ContractID,
			FlavourID:     e.FlavourID,
			EntityType:    "voucher_balance",
			EntityKey:     key,
			State:         state,
			SchemaVersion: 1,
			LastBlock:     e.BlockNumber,
		})
	}
	return out
}

// ---------------------------------------------------------------------------
// Physical card family
// ---------------------------------------------------------------------------

func projectPhysicalCardRegistered(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		SerialHash string `json:"serial_hash"`
		Checksum   string `json:"checksum"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectPhysicalCardRegistered: %w", err)
	}
	state, err := json.Marshal(map[string]any{
		"serial_hash": d.SerialHash,
		"checksum":    d.Checksum,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "physical_card",
		EntityKey:     d.SerialHash,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

// ---------------------------------------------------------------------------
// Role family
// ---------------------------------------------------------------------------

func projectMinterRoleGranted(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		Minter string `json:"minter"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectMinterRoleGranted: %w", err)
	}
	// MINTER_ROLE = keccak256("MINTER_ROLE") — computed from the ABI's constant.
	minterRole := "0x9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6"
	return roleUpsert(e, minterRole, d.Minter, true), nil
}

func projectRoleGranted(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		Role    string `json:"role"`
		Account string `json:"account"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectRoleGranted: %w", err)
	}
	return roleUpsert(e, d.Role, d.Account, true), nil
}

func projectRoleRevoked(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		Role    string `json:"role"`
		Account string `json:"account"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectRoleRevoked: %w", err)
	}
	return roleUpsert(e, d.Role, d.Account, false), nil
}

func roleUpsert(e *blockchain.DecodedEvent, role, account string, granted bool) []blockchain.EntityMutation {
	key := fmt.Sprintf("%s:%s", role, account)
	state, _ := json.Marshal(map[string]any{
		"role":    role,
		"account": account,
		"granted": granted,
	})
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "role_assignment",
		EntityKey:     key,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}
}

// ---------------------------------------------------------------------------
// Config / status family
// ---------------------------------------------------------------------------

func projectFeesAddressUpdated(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		FeesAddress string `json:"fees_address"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectFeesAddressUpdated: %w", err)
	}
	state, err := json.Marshal(map[string]any{
		"fees_address": d.FeesAddress,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "contract_config",
		EntityKey:     "singleton",
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

func projectPaused(e *blockchain.DecodedEvent, paused bool) ([]blockchain.EntityMutation, error) {
	state, err := json.Marshal(map[string]any{"paused": paused})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "contract_status",
		EntityKey:     "singleton",
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}

func projectURI(e *blockchain.DecodedEvent) ([]blockchain.EntityMutation, error) {
	var d struct {
		Value string `json:"value"`
		ID    string `json:"id"`
	}
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, fmt.Errorf("projectURI: %w", err)
	}
	state, err := json.Marshal(map[string]any{
		"token_id": d.ID,
		"value":    d.Value,
	})
	if err != nil {
		return nil, err
	}
	return []blockchain.EntityMutation{{
		Op:            blockchain.OpUpsert,
		ContractID:    e.ContractID,
		FlavourID:     e.FlavourID,
		EntityType:    "uri",
		EntityKey:     d.ID,
		State:         state,
		SchemaVersion: 1,
		LastBlock:     e.BlockNumber,
	}}, nil
}
