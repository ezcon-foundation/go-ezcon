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

package ledger

import (
	"crypto/sha512"
	"encoding/json"
	"github.com/ezcon-foundation/go-ezcon/kyc"
	"time"
)

type Ledger struct {
	Header       LedgerHeader `json:"header"`
	Accounts     SHAMap       `json:"accounts"`     // Account State Tree
	Transactions SHAMap       `json:"transactions"` // Transaction Tree
}

type LedgerHeader struct {
	Index      uint64    `json:"index"`       // Số thứ tự ledger
	Hash       []byte    `json:"hash"`        // Hash của ledger (SHA-512/256)
	ParentHash []byte    `json:"parent_hash"` // Hash của ledger trước
	StateHash  []byte    `json:"state_hash"`  // Hash của Account State Tree
	TxHash     []byte    `json:"tx_hash"`     // Hash của Transaction Tree
	TotalCoins uint64    `json:"total_coins"` // Tổng cung coin (Native Coin)
	CloseTime  time.Time `json:"close_time"`  // Thời gian đóng ledger
	CloseRes   uint32    `json:"close_res"`   // Độ phân giải thời gian (ms)
}

type Account struct {
	AccountID    string      `json:"account_id"`
	Balance      uint64      `json:"balance"`       // Số dư (drops)
	Sequence     uint32      `json:"sequence"`      // Số thứ tự giao dịch
	Reserve      uint64      `json:"reserve"`       // Dự trữ tối thiểu
	KYCData      kyc.KYCData `json:"kyc_data"`      // Thông tin KYC
	KYCHash      []byte      `json:"kyc_hash"`      // Hash của KYCData
	KYCVerified  bool        `json:"kyc_verified"`  // Trạng thái KYC
	KYCTimestamp time.Time   `json:"kyc_timestamp"` // Thời gian KYC
	TrustLines   []TrustLine `json:"trust_lines"`   // Trust lines (nếu có)
	Offers       []Offer     `json:"offers"`        // Offers trên DEX
}

type TrustLine struct {
	Issuer   string `json:"issuer"`
	Currency string `json:"currency"`
	Limit    uint64 `json:"limit"`
}

type Offer struct {
}

type Transaction struct {
	TxType   string      `json:"tx_type"`  // Loại giao dịch (KYCSet, Payment, ...)
	KYCSet   *kyc.KYCSet `json:"kyc_set"`  // Giao dịch KYCSet (nếu có)
	Hash     []byte      `json:"hash"`     // Hash giao dịch
	Result   string      `json:"result"`   // Kết quả (tesSUCCESS, tecFAILED, ...)
	Metadata []byte      `json:"metadata"` // Thay đổi trạng thái
}

type SHAMap struct {
	RootHash []byte          `json:"root_hash"` // Hash gốc của Merkle Tree
	Nodes    map[string]Node `json:"nodes"`     // Các node trong SHAMap
}

type Node struct {
	Hash   []byte `json:"hash"`    // Hash của node
	Data   []byte `json:"data"`    // Dữ liệu (Account hoặc Transaction)
	Left   string `json:"left"`    // ID node con trái
	Right  string `json:"right"`   // ID node con phải
	IsLeaf bool   `json:"is_leaf"` // Là node lá
}

// NewLedger tạo ledger mới
func NewLedger(index uint64, parentHash []byte, totalCoins uint64) *Ledger {
	return &Ledger{
		Header: LedgerHeader{
			Index:      index,
			ParentHash: parentHash,
			TotalCoins: totalCoins,
			CloseRes:   1000, // 1 giây
		},
		Accounts:     SHAMap{Nodes: make(map[string]Node)},
		Transactions: SHAMap{Nodes: make(map[string]Node)},
	}
}

// ComputeHash tính hash ledger (SHA-512/256)
func (l *Ledger) ComputeHash() []byte {
	data, _ := json.Marshal(l.Header)
	hash := sha512.Sum512_256(data)
	return hash[:]
}

// Close đóng ledger
func (l *Ledger) Close(closeTime time.Time) {
	l.Header.CloseTime = closeTime
	l.Header.StateHash = l.Accounts.RootHash
	l.Header.TxHash = l.Transactions.RootHash
	l.Header.Hash = l.ComputeHash()
}
