// SPDX-License-Identifier: MIT

package glintdough

import (
	_ "embed"
	"encoding/json"

	"github.com/pat3icki/glint/pkg/blockchain"
)

// Embedded event data-schema files.
//
//go:embed schemas/event_token_added.json
var schemaEventTokenAdded []byte

//go:embed schemas/event_token_status_changed.json
var schemaEventTokenStatusChanged []byte

//go:embed schemas/event_voucher_added.json
var schemaEventVoucherAdded []byte

//go:embed schemas/event_voucher_toggled.json
var schemaEventVoucherToggled []byte

//go:embed schemas/event_voucher_minted_single.json
var schemaEventVoucherMintedSingle []byte

//go:embed schemas/event_vouchers_minted_batch.json
var schemaEventVouchersMintedBatch []byte

//go:embed schemas/event_voucher_redeemed.json
var schemaEventVoucherRedeemed []byte

//go:embed schemas/event_physical_card_registered.json
var schemaEventPhysicalCardRegistered []byte

//go:embed schemas/event_minter_role_granted.json
var schemaEventMinterRoleGranted []byte

//go:embed schemas/event_fees_address_updated.json
var schemaEventFeesAddressUpdated []byte

//go:embed schemas/event_transfer_single.json
var schemaEventTransferSingle []byte

//go:embed schemas/event_transfer_batch.json
var schemaEventTransferBatch []byte

//go:embed schemas/event_role_granted.json
var schemaEventRoleGranted []byte

//go:embed schemas/event_role_revoked.json
var schemaEventRoleRevoked []byte

//go:embed schemas/event_role_admin_changed.json
var schemaEventRoleAdminChanged []byte

//go:embed schemas/event_paused.json
var schemaEventPaused []byte

//go:embed schemas/event_unpaused.json
var schemaEventUnpaused []byte

//go:embed schemas/event_uri.json
var schemaEventURI []byte

// eventTypes returns the full catalogue of events this plugin handles.
// Each entry maps 1-to-1 with an ABI event; topic0 is computed from the
// parsed ABI so it can never drift from the real contract.
func eventTypes() []blockchain.EventType {
	abiFrag := json.RawMessage(abiJSON())
	return []blockchain.EventType{
		{
			Name:          "TokenAdded",
			Category:      "collateral_token",
			Topic0:        topic0("TokenAdded"),
			ABIFragment:   abiFragment("TokenAdded", abiFrag),
			DataSchema:    schemaEventTokenAdded,
			SchemaVersion: 1,
		},
		{
			Name:          "TokenStatusChanged",
			Category:      "collateral_token",
			Topic0:        topic0("TokenStatusChanged"),
			ABIFragment:   abiFragment("TokenStatusChanged", abiFrag),
			DataSchema:    schemaEventTokenStatusChanged,
			SchemaVersion: 1,
		},
		{
			Name:          "VoucherAdded",
			Category:      "voucher",
			Topic0:        topic0("VoucherAdded"),
			ABIFragment:   abiFragment("VoucherAdded", abiFrag),
			DataSchema:    schemaEventVoucherAdded,
			SchemaVersion: 1,
		},
		{
			Name:          "VoucherToggled",
			Category:      "voucher",
			Topic0:        topic0("VoucherToggled"),
			ABIFragment:   abiFragment("VoucherToggled", abiFrag),
			DataSchema:    schemaEventVoucherToggled,
			SchemaVersion: 1,
		},
		{
			Name:          "VoucherMintedSingle",
			Category:      "mint",
			Topic0:        topic0("VoucherMintedSingle"),
			ABIFragment:   abiFragment("VoucherMintedSingle", abiFrag),
			DataSchema:    schemaEventVoucherMintedSingle,
			SchemaVersion: 1,
		},
		{
			Name:          "VouchersMintedBatch",
			Category:      "mint",
			Topic0:        topic0("VouchersMintedBatch"),
			ABIFragment:   abiFragment("VouchersMintedBatch", abiFrag),
			DataSchema:    schemaEventVouchersMintedBatch,
			SchemaVersion: 1,
		},
		{
			Name:          "VoucherRedeemed",
			Category:      "redeem",
			Topic0:        topic0("VoucherRedeemed"),
			ABIFragment:   abiFragment("VoucherRedeemed", abiFrag),
			DataSchema:    schemaEventVoucherRedeemed,
			SchemaVersion: 1,
		},
		{
			Name:          "PhysicalCardRegistered",
			Category:      "physical_card",
			Topic0:        topic0("PhysicalCardRegistered"),
			ABIFragment:   abiFragment("PhysicalCardRegistered", abiFrag),
			DataSchema:    schemaEventPhysicalCardRegistered,
			SchemaVersion: 1,
		},
		{
			Name:          "MinterRoleGranted",
			Category:      "role",
			Topic0:        topic0("MinterRoleGranted"),
			ABIFragment:   abiFragment("MinterRoleGranted", abiFrag),
			DataSchema:    schemaEventMinterRoleGranted,
			SchemaVersion: 1,
		},
		{
			Name:          "FeesAddressUpdated",
			Category:      "config",
			Topic0:        topic0("FeesAddressUpdated"),
			ABIFragment:   abiFragment("FeesAddressUpdated", abiFrag),
			DataSchema:    schemaEventFeesAddressUpdated,
			SchemaVersion: 1,
		},
		{
			Name:          "TransferSingle",
			Category:      "transfer",
			Topic0:        topic0("TransferSingle"),
			ABIFragment:   abiFragment("TransferSingle", abiFrag),
			DataSchema:    schemaEventTransferSingle,
			SchemaVersion: 1,
		},
		{
			Name:          "TransferBatch",
			Category:      "transfer",
			Topic0:        topic0("TransferBatch"),
			ABIFragment:   abiFragment("TransferBatch", abiFrag),
			DataSchema:    schemaEventTransferBatch,
			SchemaVersion: 1,
		},
		{
			Name:          "RoleGranted",
			Category:      "role",
			Topic0:        topic0("RoleGranted"),
			ABIFragment:   abiFragment("RoleGranted", abiFrag),
			DataSchema:    schemaEventRoleGranted,
			SchemaVersion: 1,
		},
		{
			Name:          "RoleRevoked",
			Category:      "role",
			Topic0:        topic0("RoleRevoked"),
			ABIFragment:   abiFragment("RoleRevoked", abiFrag),
			DataSchema:    schemaEventRoleRevoked,
			SchemaVersion: 1,
		},
		{
			Name:          "RoleAdminChanged",
			Category:      "role",
			Topic0:        topic0("RoleAdminChanged"),
			ABIFragment:   abiFragment("RoleAdminChanged", abiFrag),
			DataSchema:    schemaEventRoleAdminChanged,
			SchemaVersion: 1,
		},
		{
			Name:          "Paused",
			Category:      "status",
			Topic0:        topic0("Paused"),
			ABIFragment:   abiFragment("Paused", abiFrag),
			DataSchema:    schemaEventPaused,
			SchemaVersion: 1,
		},
		{
			Name:          "Unpaused",
			Category:      "status",
			Topic0:        topic0("Unpaused"),
			ABIFragment:   abiFragment("Unpaused", abiFrag),
			DataSchema:    schemaEventUnpaused,
			SchemaVersion: 1,
		},
		{
			Name:          "URI",
			Category:      "uri",
			Topic0:        topic0("URI"),
			ABIFragment:   abiFragment("URI", abiFrag),
			DataSchema:    schemaEventURI,
			SchemaVersion: 1,
		},
	}
}



// abiFragment extracts the single-event ABI fragment for the named event.
func abiFragment(eventName string, _ json.RawMessage) json.RawMessage {
	ev, ok := parsedABI.Events[eventName]
	if !ok {
		return json.RawMessage("{}")
	}
	b, err := json.Marshal(ev)
	if err != nil {
		return json.RawMessage("{}")
	}
	return b
}
