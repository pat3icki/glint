// SPDX-License-Identifier: MIT

package glintdough

import "github.com/pat3icki/glint/pkg/blockchain"

func init() {
	blockchain.Register(New())
}
