// SPDX-License-Identifier: MIT

// Package blockchain provides the plugin framework for indexing EVM smart contracts.
package blockchain

import "errors"

// Sentinel errors returned by the framework packages.
var (
	// ErrPluginNotFound is returned when a requested flavour key is not in the registry.
	ErrPluginNotFound = errors.New("blockchain: plugin not found")

	// ErrPluginAlreadyRegistered is returned when a second plugin tries to use the same flavour key.
	ErrPluginAlreadyRegistered = errors.New("blockchain: plugin already registered")

	// ErrUnknownEvent is returned when DecodeEvent cannot match a log's topic0.
	ErrUnknownEvent = errors.New("blockchain: unknown event topic0")

	// ErrInvalidData is returned when decoded data fails JSON-Schema validation.
	ErrInvalidData = errors.New("blockchain: invalid data schema")

	// ErrEventAlreadyIngested is returned by InsertEvent when the unique constraint fires.
	ErrEventAlreadyIngested = errors.New("blockchain: event already ingested")

	// ErrCursorNotFound is returned when there is no cursor row for the given contract+indexer.
	ErrCursorNotFound = errors.New("blockchain: cursor not found")
)
