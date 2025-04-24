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
	"fmt"
)

// TxType represents a transaction type
type TxType int

// Transaction transaction
const (
	TxTypeKYCSet TxType = iota
	TxTypePayment
	TxTypeTrustSet
	TxTypeTrustConfirm
	TxTypeOfferCreate
	TxTypeOfferCancel
	TxTypeAccountSet
	TxTypeSignerListSet
	TxTypeEscrowCreate
	TxTypeEscrowFinish
	TxTypeEscrowCancel
	TxTypePaymentChannelCreate
	TxTypePaymentChannelFund
	TxTypePaymentChannelClaim
)

// txTypeNames maps TxType to string for JSON
var txTypeNames = map[TxType]string{
	TxTypeKYCSet:               "KYCSet",
	TxTypePayment:              "Payment",
	TxTypeTrustSet:             "TrustSet",
	TxTypeTrustConfirm:         "TrustConfirm",
	TxTypeOfferCreate:          "OfferCreate",
	TxTypeOfferCancel:          "OfferCancel",
	TxTypeAccountSet:           "AccountSet",
	TxTypeSignerListSet:        "SignerListSet",
	TxTypeEscrowCreate:         "EscrowCreate",
	TxTypeEscrowFinish:         "EscrowFinish",
	TxTypeEscrowCancel:         "EscrowCancel",
	TxTypePaymentChannelCreate: "PaymentChannelCreate",
	TxTypePaymentChannelFund:   "PaymentChannelFund",
	TxTypePaymentChannelClaim:  "PaymentChannelClaim",
}

// txTypeValues maps string to TxType for unmarshaling
var txTypeValues = map[string]TxType{
	"KYCSet":               TxTypeKYCSet,
	"Payment":              TxTypePayment,
	"TrustSet":             TxTypeTrustSet,
	"TrustConfirm":         TxTypeTrustConfirm,
	"OfferCreate":          TxTypeOfferCreate,
	"OfferCancel":          TxTypeOfferCancel,
	"AccountSet":           TxTypeAccountSet,
	"SignerListSet":        TxTypeSignerListSet,
	"EscrowCreate":         TxTypeEscrowCreate,
	"EscrowFinish":         TxTypeEscrowFinish,
	"EscrowCancel":         TxTypeEscrowCancel,
	"PaymentChannelCreate": TxTypePaymentChannelCreate,
	"PaymentChannelFund":   TxTypePaymentChannelFund,
	"PaymentChannelClaim":  TxTypePaymentChannelClaim,
}

// String returns the string representation of TxType
func (t TxType) String() string {
	if name, exists := txTypeNames[t]; exists {
		return name
	}
	return fmt.Sprintf("Unknown(%d)", t)
}

// MarshalJSON serializes TxType to JSON string
func (t TxType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON deserializes TxType from JSON string
func (t *TxType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if val, exists := txTypeValues[s]; exists {
		*t = val
		return nil
	}
	return fmt.Errorf("invalid tx_type: %s", s)
}
