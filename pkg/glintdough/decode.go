// SPDX-License-Identifier: MIT

package glintdough

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/pat3icki/glint/pkg/blockchain"
)

// topic0Map is built at init from the parsed ABI; maps topic0 hex → event name.
var topic0Map map[string]string

func init() {
	topic0Map = make(map[string]string, len(parsedABI.Events))
	for name, ev := range parsedABI.Events {
		topic0Map[ev.ID.Hex()] = name
	}
}

// decodeEvent decodes a raw EVM log into a DecodedEvent.
// It returns blockchain.ErrUnknownEvent if the topic0 is not in the ABI.
func decodeEvent(ctx context.Context, log ethtypes.Log) (*blockchain.DecodedEvent, error) {
	if len(log.Topics) == 0 {
		return nil, blockchain.ErrUnknownEvent
	}
	t0 := strings.ToLower(log.Topics[0].Hex())

	eventName, ok := topic0Map[t0]
	if !ok {
		return nil, fmt.Errorf("%w: %s", blockchain.ErrUnknownEvent, t0)
	}

	abiEvent := parsedABI.Events[eventName]

	// Decode non-indexed (data) fields.
	dataMap := make(map[string]any)
	if len(log.Data) > 0 {
		if err := abiEvent.Inputs.UnpackIntoMap(dataMap, log.Data); err != nil {
			return nil, fmt.Errorf("glintdough: unpack %s data: %w", eventName, err)
		}
	}

	// Decode indexed fields (topics[1:]).
	indexedArgs := make(abi.Arguments, 0)
	for _, arg := range abiEvent.Inputs {
		if arg.Indexed {
			indexedArgs = append(indexedArgs, arg)
		}
	}
	if len(indexedArgs) > 0 && len(log.Topics) > 1 {
		if err := abi.Arguments(indexedArgs).UnpackIntoMap(dataMap, encodeTopics(log.Topics[1:])); err != nil {
			return nil, fmt.Errorf("glintdough: unpack %s topics: %w", eventName, err)
		}
	}

	// Normalise all values to JSON-friendly types.
	normalised := normaliseMap(dataMap)

	dataJSON, err := json.Marshal(normalised)
	if err != nil {
		return nil, fmt.Errorf("glintdough: marshal %s data: %w", eventName, err)
	}

	// Build topics list as lower-case hex strings.
	topics := make([]string, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = strings.ToLower(t.Hex())
	}

	event := &blockchain.DecodedEvent{
		EventType:     eventName,
		SchemaVersion: 1,
		TxHash:        strings.ToLower(log.TxHash.Hex()),
		LogIndex:      int(log.Index),
		BlockNumber:   int64(log.BlockNumber),
		BlockHash:     strings.ToLower(log.BlockHash.Hex()),
		Emitter:       strings.ToLower(log.Address.Hex()),
		Topics:        topics,
		Data:          dataJSON,
	}
	return event, nil
}

// encodeTopics packs indexed topic values as raw 32-byte slices so that
// abi.Arguments.UnpackIntoMap can decode them correctly.
func encodeTopics(topics []common.Hash) []byte {
	buf := make([]byte, len(topics)*32)
	for i, t := range topics {
		copy(buf[i*32:], t.Bytes())
	}
	return buf
}

// normaliseMap converts all values in a decoded ABI map to JSON-serialisable
// types. *big.Int → decimal string; common.Address → lower-case hex;
// [32]byte → lower-case hex; []byte → lower-case hex.
func normaliseMap(m map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	for k, v := range m {
		out[k] = normaliseValue(v)
	}
	return out
}

func normaliseValue(v any) any {
	switch val := v.(type) {
	case *big.Int:
		if val == nil {
			return "0"
		}
		return val.String()
	case common.Address:
		return strings.ToLower(val.Hex())
	case common.Hash:
		return strings.ToLower(val.Hex())
	case [32]byte:
		return strings.ToLower(common.Bytes2Hex(val[:]))
	case []byte:
		return strings.ToLower(common.Bytes2Hex(val))
	case bool:
		return val
	case string:
		return val
	case uint8:
		return int(val)
	case []*big.Int:
		strs := make([]string, len(val))
		for i, n := range val {
			if n == nil {
				strs[i] = "0"
			} else {
				strs[i] = n.String()
			}
		}
		return strs
	case []common.Address:
		addrs := make([]string, len(val))
		for i, a := range val {
			addrs[i] = strings.ToLower(a.Hex())
		}
		return addrs
	default:
		return val
	}
}
