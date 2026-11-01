// SPDX-License-Identifier: MIT

// Package glintdough implements the blockchain.Plugin for the GlintDough
// ERC-1155 collateral-backed gift card contract.
//
// It is the reference implementation of the plugin model.
// The package registers itself into blockchain.GlobalRegistry via its init function.
package glintdough

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	bindings "github.com/pat3icki/glint/contract/out/bindings-go/glintdough"
)

// parsedABI is the lazily-parsed contract ABI, initialised once at package init.
var parsedABI abi.ABI

func init() {
	a, err := bindings.GlintdoughMetaData.GetAbi()
	if err != nil {
		// Panic is acceptable here: the ABI is baked into the binary and
		// must always be valid. Any error here is a programming mistake.
		panic(fmt.Sprintf("glintdough: parse ABI: %v", err))
	}
	parsedABI = *a
}

// abiJSON returns the full ABI JSON string from the generated metadata.
func abiJSON() string {
	return bindings.GlintdoughMetaData.ABI
}

// topic0 returns the keccak256 event signature hash (topic0) for an event by name,
// as a lower-case 0x-prefixed hex string.
func topic0(eventName string) string {
	ev, ok := parsedABI.Events[eventName]
	if !ok {
		panic(fmt.Sprintf("glintdough: ABI has no event %q", eventName))
	}
	return ev.ID.Hex() // already 0x-prefixed lower-case hex
}
