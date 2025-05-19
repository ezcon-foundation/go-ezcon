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

package asset

import (
	"encoding/json"
	"fmt"
)

type AssetType int

// Asset transaction
const (
	AssetTypeRealEstate AssetType = iota
	AssetTypeVehicle
	AssetTypeArt
	AssetTypeCollectible
	AssetTypeIntellectualProperty
	AssetTypeFinancialInstrument
	AssetTypeDigitalAsset
	AssetTypeCommodity
	AssetTypeCurrency
	AssetTypeBond
	AssetTypeStock
	AssetTypeFuture
	AssetTypeOption
	AssetTypeSwap
	AssetTypeInsurance
	AssetTypeLoan
	AssetTypeMortgage
	AssetTypeLease
	AssetTypeRoyalty
	AssetTypeFranchise
	AssetTypeMembership
	AssetTypeLicense
	AssetTypeContract
	AssetTypeWarranty
	AssetTypeGuarantee
	AssetTypePledge
	AssetTypeCollateral
	AssetTypeDeposit
	AssetTypeEscrow
	AssetTypeTrust
	AssetTypeSecuritization
	AssetTypeTokenization
	AssetTypeFractionalOwnership
	AssetTypeCrowdfunding
	AssetTypePeerToPeerLending
	AssetTypeMicrofinance
	AssetTypeRemittance
	AssetTypePayment
	AssetTypeTransfer
	AssetTypeExchange
	AssetTypeSettlement
	AssetTypeClearing
	AssetTypeCustody
)

// String returns the string representation of the AssetType
var assetTypeNames = map[AssetType]string{
	AssetTypeRealEstate:           "Real Estate",
	AssetTypeVehicle:              "Vehicle",
	AssetTypeArt:                  "Art",
	AssetTypeCollectible:          "Collectible",
	AssetTypeIntellectualProperty: "Intellectual Property",
	AssetTypeFinancialInstrument:  "Financial Instrument",
	AssetTypeDigitalAsset:         "Digital Asset",
	AssetTypeCommodity:            "Commodity",
	AssetTypeCurrency:             "Currency",
	AssetTypeBond:                 "Bond",
	AssetTypeStock:                "Stock",
	AssetTypeFuture:               "Future",
	AssetTypeOption:               "Option",
	AssetTypeSwap:                 "Swap",
	AssetTypeInsurance:            "Insurance",
	AssetTypeLoan:                 "Loan",
	AssetTypeMortgage:             "Mortgage",
	AssetTypeLease:                "Lease",
	AssetTypeRoyalty:              "Royalty",
	AssetTypeFranchise:            "Franchise",
	AssetTypeMembership:           "Membership",
	AssetTypeLicense:              "License",
	AssetTypeContract:             "Contract",
	AssetTypeWarranty:             "Warranty",
	AssetTypeGuarantee:            "Guarantee",
	AssetTypePledge:               "Pledge",
	AssetTypeCollateral:           "Collateral",
	AssetTypeDeposit:              "Deposit",
	AssetTypeEscrow:               "Escrow",
	AssetTypeTrust:                "Trust",
	AssetTypeSecuritization:       "Securitization",
	AssetTypeTokenization:         "Tokenization",
	AssetTypeFractionalOwnership:  "Fractional Ownership",
	AssetTypeCrowdfunding:         "Crowdfunding",
	AssetTypePeerToPeerLending:    "Peer to Peer Lending",
	AssetTypeMicrofinance:         "Microfinance",
	AssetTypeRemittance:           "Remittance",
	AssetTypePayment:              "Payment",
	AssetTypeTransfer:             "Transfer",
	AssetTypeExchange:             "Exchange",
	AssetTypeSettlement:           "Settlement",
	AssetTypeClearing:             "Clearing",
	AssetTypeCustody:              "Custody",
}

// String returns the string representation of the AssetType
var assetTypeValues = map[string]AssetType{
	"Real Estate":           AssetTypeRealEstate,
	"Vehicle":               AssetTypeVehicle,
	"Art":                   AssetTypeArt,
	"Collectible":           AssetTypeCollectible,
	"Intellectual Property": AssetTypeIntellectualProperty,
	"Financial Instrument":  AssetTypeFinancialInstrument,
	"Digital Asset":         AssetTypeDigitalAsset,
	"Commodity":             AssetTypeCommodity,
	"Currency":              AssetTypeCurrency,
	"Bond":                  AssetTypeBond,
	"Stock":                 AssetTypeStock,
	"Future":                AssetTypeFuture,
	"Option":                AssetTypeOption,
	"Swap":                  AssetTypeSwap,
	"Insurance":             AssetTypeInsurance,
	"Loan":                  AssetTypeLoan,
	"Mortgage":              AssetTypeMortgage,
	"Lease":                 AssetTypeLease,
	"Royalty":               AssetTypeRoyalty,
	"Franchise":             AssetTypeFranchise,
	"Membership":            AssetTypeMembership,
	"License":               AssetTypeLicense,
	"Contract":              AssetTypeContract,
	"Warranty":              AssetTypeWarranty,
	"Guarantee":             AssetTypeGuarantee,
	"Pledge":                AssetTypePledge,
	"Collateral":            AssetTypeCollateral,
	"Deposit":               AssetTypeDeposit,
	"Escrow":                AssetTypeEscrow,
	"Trust":                 AssetTypeTrust,
	"Securitization":        AssetTypeSecuritization,
	"Tokenization":          AssetTypeTokenization,
	"Fractional Ownership":  AssetTypeFractionalOwnership,
	"Crowdfunding":          AssetTypeCrowdfunding,
	"Peer to Peer Lending":  AssetTypePeerToPeerLending,
	"Microfinance":          AssetTypeMicrofinance,
	"Remittance":            AssetTypeRemittance,
	"Payment":               AssetTypePayment,
	"Transfer":              AssetTypeTransfer,
	"Exchange":              AssetTypeExchange,
	"Settlement":            AssetTypeSettlement,
	"Clearing":              AssetTypeClearing,
	"Custody":               AssetTypeCustody,
}

// String returns the string representation of the AssetType
func (t AssetType) String() string {
	if name, exists := assetTypeNames[t]; exists {
		return name
	}
	return "Unknown"
}

// MarshalJSON marshals the AssetType to JSON
func (t AssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON unmarshals the JSON to AssetType
func (t *AssetType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	if value, exists := assetTypeValues[name]; exists {
		*t = value
		return nil
	}
	return fmt.Errorf("unknown asset type: %s", name)
}
