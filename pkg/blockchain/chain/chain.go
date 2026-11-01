// SPDX-License-Identifier: MIT

// Package chain wraps go-ethereum's ethclient to provide the single authorised
// path to EVM node access for the blockchain framework.
//
// No other package in this module may import ethclient directly.
package chain

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client wraps an ethclient.Client, exposing only the operations the
// framework needs.
type Client struct {
	inner *ethclient.Client
}

// Dial connects to the Ethereum node at rpcURL and returns a Client.
func Dial(ctx context.Context, rpcURL string) (*Client, error) {
	c, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("chain: dial %s: %w", rpcURL, err)
	}
	return &Client{inner: c}, nil
}

// Close closes the underlying connection.
func (c *Client) Close() { c.inner.Close() }

// ChainID returns the chain ID of the connected network.
func (c *Client) ChainID(ctx context.Context) (*big.Int, error) {
	id, err := c.inner.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("chain: ChainID: %w", err)
	}
	return id, nil
}

// HeaderByNumber returns the block header for number n.
// Pass nil to get the latest header.
func (c *Client) HeaderByNumber(ctx context.Context, n *big.Int) (*types.Header, error) {
	h, err := c.inner.HeaderByNumber(ctx, n)
	if err != nil {
		return nil, fmt.Errorf("chain: HeaderByNumber(%v): %w", n, err)
	}
	return h, nil
}

// FilterLogs returns all logs matching q.
func (c *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	logs, err := c.inner.FilterLogs(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("chain: FilterLogs: %w", err)
	}
	return logs, nil
}

// LatestBlockNumber returns the number of the most recently finalised block.
func (c *Client) LatestBlockNumber(ctx context.Context) (int64, error) {
	h, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return h.Number.Int64(), nil
}
