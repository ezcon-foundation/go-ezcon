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
	"errors"
	"github.com/ezcon-foundation/go-ezcon/kyc"
)

func (l *Ledger) ProcessKYCSet(kycSet kyc.KYCSet) error {
	account, exists := l.Accounts[kycSet.Account]
	if !exists {
		return errors.New("account not found")
	}

	if kycSet.Sequence != account.Sequence {
		return errors.New("invalid sequence number")
	}

	const MIN_FEE = 10
	if kycSet.Fee < MIN_FEE || account.Balance < kycSet.Fee {
		return errors.New("fee too low or insufficient balance")
	}

	// Giả lập kiểm tra chữ ký (sẽ thêm crypto sau)
	account.KYCData = kycSet.KYCData
	account.KYCHash = kycSet.KYCHash
	account.KYCVerified = true
	account.KYCTimestamp = kycSet.Timestamp
	account.Sequence++
	account.Balance -= kycSet.Fee
	account.Reserve += 2

	l.Accounts[kycSet.Account] = account
	return nil
}
