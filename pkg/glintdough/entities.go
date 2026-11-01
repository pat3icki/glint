// SPDX-License-Identifier: MIT

package glintdough

import (
	_ "embed"

	"github.com/pat3icki/glint/pkg/blockchain"
)

// Embedded entity state-schema files.
//
//go:embed schemas/entity_collateral_token.json
var schemaEntityCollateralToken []byte

//go:embed schemas/entity_voucher.json
var schemaEntityVoucher []byte

//go:embed schemas/entity_voucher_balance.json
var schemaEntityVoucherBalance []byte

//go:embed schemas/entity_role_assignment.json
var schemaEntityRoleAssignment []byte

//go:embed schemas/entity_contract_status.json
var schemaEntityContractStatus []byte

//go:embed schemas/entity_contract_config.json
var schemaEntityContractConfig []byte

//go:embed schemas/entity_physical_card.json
var schemaEntityPhysicalCard []byte

//go:embed schemas/entity_uri.json
var schemaEntityURI []byte

// entityTypes returns the 8 derived-state entity types for GlintDough.
func entityTypes() []blockchain.EntityType {
	return []blockchain.EntityType{
		{
			Name:          "collateral_token",
			Description:   "ERC-20 collateral token registered via addCollateralToken.",
			StateSchema:   schemaEntityCollateralToken,
			SchemaVersion: 1,
		},
		{
			Name:          "voucher",
			Description:   "ERC-1155 voucher type (gift card denomination) added via addVoucher.",
			StateSchema:   schemaEntityVoucher,
			SchemaVersion: 1,
		},
		{
			Name:          "voucher_balance",
			Description:   "Per-holder ERC-1155 balance derived from TransferSingle/TransferBatch — the sole source of truth for balances.",
			StateSchema:   schemaEntityVoucherBalance,
			SchemaVersion: 1,
		},
		{
			Name:          "role_assignment",
			Description:   "AccessControl role membership (DEFAULT_ADMIN_ROLE / MINTER_ROLE) derived from RoleGranted/RoleRevoked.",
			StateSchema:   schemaEntityRoleAssignment,
			SchemaVersion: 1,
		},
		{
			Name:          "contract_status",
			Description:   "Singleton entity tracking whether the contract is currently paused.",
			StateSchema:   schemaEntityContractStatus,
			SchemaVersion: 1,
		},
		{
			Name:          "contract_config",
			Description:   "Singleton entity tracking nextTokenId and feesAddress.",
			StateSchema:   schemaEntityContractConfig,
			SchemaVersion: 1,
		},
		{
			Name:          "physical_card",
			Description:   "Physical gift card registered via PhysicalCardRegistered event.",
			StateSchema:   schemaEntityPhysicalCard,
			SchemaVersion: 1,
		},
		{
			Name:          "uri",
			Description:   "ERC-1155 metadata URI per token ID from the URI event.",
			StateSchema:   schemaEntityURI,
			SchemaVersion: 1,
		},
	}
}
