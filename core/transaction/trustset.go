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
	"time"
)

// TrustSet transaction
type TrustSet struct {
	BaseTransaction
	Destination string    `json:"destination"`
	Currency    string    `json:"currency"`
	Limit       uint64    `json:"limit"`
	Conditions  []string  `json:"conditions"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func (t *TrustSet) GetTxType() TxType {
	return txTypeValues[t.TxType]
}

func (t *TrustSet) GetAccount() string {
	return t.Account
}

func (t *TrustSet) GetSequence() uint64 {
	return t.Sequence
}

func (t *TrustSet) GetFee() uint64 {
	return t.Fee
}

func (t *TrustSet) Serialize() ([]byte, error) {
	return json.Marshal(t)
}
