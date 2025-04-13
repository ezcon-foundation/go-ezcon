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

package types

import (
	"encoding/json"
	"time"
)

// Transaction defines the interface for all transactions
type Transaction interface {
	Serialize() ([]byte, error)
	GetAccount() string
	GetTxType() string
}

// BaseTransaction contains common fields
type BaseTransaction struct {
	TxType    TxType    `json:"tx_type"`
	Account   string    `json:"account"`
	Sequence  uint32    `json:"sequence"`
	Fee       uint64    `json:"fee"`
	Timestamp time.Time `json:"timestamp"`
	Signature []byte    `json:"signature"`
}

func (b *BaseTransaction) GetAccount() string { return b.Account }
func (b *BaseTransaction) GetTxType() TxType  { return b.TxType }

// Amount represents a currency amount
type Amount struct {
	Value    uint64 `json:"value"`
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"` // Optional for EZC
}

// TrustSet initiates a trust line
type TrustSet struct {
	BaseTransaction
	TrustAccount string    `json:"trust_account"`
	LimitAmount  Amount    `json:"limit_amount"`
	QualityIn    uint32    `json:"quality_in"`
	QualityOut   uint32    `json:"quality_out"`
	Flags        uint32    `json:"flags"`
	ExpiresAt    time.Time `json:"expires_at"`
	Conditions   []string  `json:"conditions"`
}

func (t *TrustSet) Serialize() ([]byte, error) {
	return json.Marshal(t)
}

// TrustConfirm confirms a trust line
type TrustConfirm struct {
	BaseTransaction
	TrustAccount string `json:"trust_account"`
	Currency     string `json:"currency"`
}

func (t *TrustConfirm) Serialize() ([]byte, error) {
	return json.Marshal(t)
}
