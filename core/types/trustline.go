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

import "time"

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
