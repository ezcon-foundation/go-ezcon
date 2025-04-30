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

package transaction

import (
	"encoding/json"
	"errors"
	"time"
)

// Transaction defines the interface for all transactions
type Transaction interface {
	GetTxType() TxType
	GetAccount() string
	GetSequence() uint64
	GetFee() uint64
	Serialize() ([]byte, error)
}

// BaseTransaction contains common fields
type BaseTransaction struct {
	TxType    string    `json:"tx_type"`
	Account   string    `json:"account"`
	Sequence  uint64    `json:"sequence"`
	Fee       uint64    `json:"fee"`
	Timestamp time.Time `json:"timestamp"`
	Signature string    `json:"signature"`
}

// Amount represents a currency amount
type Amount struct {
	Value    uint64 `json:"value"`
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"`
}

// ParseTransaction parses JSON to Transaction
func ParseTransaction(rawTx map[string]interface{}) (Transaction, error) {
	txType, ok := rawTx["tx_type"].(string)
	if !ok {
		return nil, errors.New("missing tx_type")
	}
	data, err := json.Marshal(rawTx)
	if err != nil {
		return nil, err
	}
	switch txTypeValues[txType] {
	case TxTypeTrustSet:
		var tx TrustSet
		if err := json.Unmarshal(data, &tx); err != nil {
			return nil, err
		}
		return &tx, nil
	default:
		return nil, errors.New("unsupported tx_type")
	}
}
