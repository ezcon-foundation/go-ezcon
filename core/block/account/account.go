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

package account

import (
	"github.com/ezcon-foundation/go-ezcon/core/block/account/asset"
	"github.com/ezcon-foundation/go-ezcon/core/block/account/kyc"
	"github.com/ezcon-foundation/go-ezcon/core/block/account/trustline"
	"time"
)

// Account represents a user account
type Account struct {
	AccountID    string                `json:"account_id"`
	Balance      uint64                `json:"balance"`
	Sequence     uint32                `json:"sequence"`
	Reserve      uint64                `json:"reserve"`
	KYCData      kyc.KYCData           `json:"kyc_data"`
	KYCHash      []byte                `json:"kyc_hash"`
	KYCVerified  bool                  `json:"kyc_verified"`
	KYCTimestamp time.Time             `json:"kyc_timestamp"`
	TrustLines   []trustline.TrustLine `json:"trust_lines"`
	Assets       []asset.Asset         `json:"assets"`
}
