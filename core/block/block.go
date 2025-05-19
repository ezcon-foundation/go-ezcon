/*
 * Copyright (c) 2025 EZCON Foundation.
 *
 * The go-ezcon library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The go-ezcon library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with the go-ezcon library. If not, see <http://www.gnu.org/licenses/>.
 */

package block

import (
	"time"
)

// Block thể hiện trạng thái của một block trọng mạng blockchain ezcon
type Block struct {
	Header       BlockHeader `json:"header"`
	Accounts     SHAMap      `json:"accounts"`
	Transactions SHAMap      `json:"transactions"`
}

// BlockHeader Thể hiện thông tin data header của 1 block
type BlockHeader struct {

	// Index là index của block
	Index uint64 `json:"index"`

	// Hash của block
	Hash []byte `json:"hash"`

	// ParentHash là hash của block trước đó
	ParentHash []byte `json:"parent_hash"`

	//
	StateHash  []byte    `json:"state_hash"`
	TotalCoins uint64    `json:"total_coins"`
	CloseTime  time.Time `json:"close_time"`
}

type SHAMap struct {
	RootHash []byte          `json:"root_hash"` // Root hash of Merkle Tree
	Nodes    map[string]Node `json:"nodes"`
}

type Node struct {
	Hash   []byte `json:"hash"` // Hash của node
	Data   []byte `json:"data"`
	Left   string `json:"left"`
	Right  string `json:"right"`
	IsLeaf bool   `json:"is_leaf"`
}

func NewBlock(index uint64, parentHash []byte, totalCoins uint64) *Block {
	return &Block{
		Header: BlockHeader{
			Index:      index,
			ParentHash: parentHash,
			TotalCoins: totalCoins,
		},
	}
}
