// SPDX-License-Identifier: MIT

// Package schema wraps santhosh-tekuri/jsonschema/v6 for JSON Schema draft 2020-12 validation.
//
// No other package in this module may import the jsonschema library directly.
package schema

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Schema is a compiled JSON Schema ready for validation.
type Schema struct {
	compiled *jsonschema.Schema
}

// Compile parses and compiles a JSON Schema document.
// schemaJSON must be a valid JSON Schema draft 2020-12 object.
func Compile(schemaJSON []byte) (*Schema, error) {
	c := jsonschema.NewCompiler()
	// Unmarshal to generic map first so we can feed it to the compiler.
	var raw any
	if err := json.Unmarshal(schemaJSON, &raw); err != nil {
		return nil, fmt.Errorf("schema: unmarshal: %w", err)
	}
	const uri = "schema:///inline"
	if err := c.AddResource(uri, raw); err != nil {
		return nil, fmt.Errorf("schema: add resource: %w", err)
	}
	s, err := c.Compile(uri)
	if err != nil {
		return nil, fmt.Errorf("schema: compile: %w", err)
	}
	return &Schema{compiled: s}, nil
}

// MustCompile is like Compile but panics on error.
// Safe to use at package init for static schemas.
func MustCompile(schemaJSON []byte) *Schema {
	s, err := Compile(schemaJSON)
	if err != nil {
		panic(err)
	}
	return s
}

// Validate validates data against the schema.
// data may be a json.RawMessage, a map, or any JSON-serialisable value.
// Returns a descriptive error if validation fails.
func (s *Schema) Validate(data any) error {
	// Normalise: always work via a round-trip through encoding/json so we get
	// the same representation the jsonschema library expects.
	var v any
	switch d := data.(type) {
	case json.RawMessage:
		if err := json.Unmarshal(d, &v); err != nil {
			return fmt.Errorf("schema: unmarshal data: %w", err)
		}
	case []byte:
		if err := json.Unmarshal(d, &v); err != nil {
			return fmt.Errorf("schema: unmarshal data: %w", err)
		}
	default:
		b, err := json.Marshal(d)
		if err != nil {
			return fmt.Errorf("schema: marshal data: %w", err)
		}
		if err = json.Unmarshal(b, &v); err != nil {
			return fmt.Errorf("schema: unmarshal data: %w", err)
		}
	}
	if err := s.compiled.Validate(v); err != nil {
		// Give a compact single-line error for logging.
		return fmt.Errorf("schema: validation failed: %s", strings.ReplaceAll(err.Error(), "\n", "; "))
	}
	return nil
}
