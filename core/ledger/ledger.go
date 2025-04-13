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

// Ledger represents a snapshot of the blockchain state
type Ledger struct {
	Header       LedgerHeader `json:"header"`
	Accounts     SHAMap       `json:"accounts"`
	Transactions SHAMap       `json:"transactions"`
}

// LedgerHeader contains metadata for the ledger
type LedgerHeader struct {
	Index      uint64    `json:"index"`
	Hash       []byte    `json:"hash"`
	ParentHash []byte    `json:"parent_hash"`
	StateHash  []byte    `json:"state_hash"`
	TxHash     []byte    `json:"tx_hash"`
	TotalCoins uint64    `json:"total_coins"`
	CloseTime  time.Time `json:"close_time"`
	CloseRes   uint32    `json:"close_res"`
}

// Account represents a user account
type Account struct {
	AccountID    string      `json:"account_id"`
	Balance      uint64      `json:"balance"`
	Sequence     uint32      `json:"sequence"`
	Reserve      uint64      `json:"reserve"`
	KYCData      kyc.KYCData `json:"kyc_data"`
	KYCHash      []byte      `json:"kyc_hash"`
	KYCVerified  bool        `json:"kyc_verified"`
	KYCTimestamp time.Time   `json:"kyc_timestamp"`
	TrustLines   []TrustLine `json:"trust_lines"`
	Assets       []Asset     `json:"assets"`
}

type Asset struct {
	Type        string `json:"type"`         // Loại tài sản (RealEstate, IntellectualProperty, Financial, ...)
	ID          string `json:"id"`           // Định danh pháp lý (Sổ đỏ, bằng sáng chế, mã chứng khoán)
	Description string `json:"description"`  // Mô tả (địa chỉ nhà, tên nhãn hiệu, ...)
	Value       uint64 `json:"value"`        // Giá trị ước tính (VND)
	LegalHash   []byte `json:"legal_hash"`   // Hash hồ sơ pháp lý
	IsTokenized bool   `json:"is_tokenized"` // Đã token hóa chưa
}

// TrustLine defines a trust relationship between two accounts
type TrustLine struct {
	Account    string    `json:"account"`     // Counterparty account
	Currency   string    `json:"currency"`    // EZC, USD, REALESTATE, IP, etc.
	Limit      uint64    `json:"limit"`       // Trust limit
	Balance    int64     `json:"balance"`     // Current balance
	QualityIn  uint32    `json:"quality_in"`  // Inbound rate
	QualityOut uint32    `json:"quality_out"` // Outbound rate
	Flags      uint32    `json:"flags"`       // NoRipple, Authorized, etc.
	IsVerified bool      `json:"is_verified"` // Mutually confirmed
	ExpiresAt  time.Time `json:"expires_at"`  // Expiration time
	Conditions []string  `json:"conditions"`  // e.g., ["only_token:REALESTATE:NFT123"]
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
