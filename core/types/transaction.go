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
	"github.com/ezcon-foundation/go-ezcon/kyc"
	"time"
)

// Transaction là giao diện chung cho mọi loại giao dịch trong go-ezcon
type Transaction interface {
	Serialize() ([]byte, error) // Chuyển giao dịch thành JSON
	GetAccount() string         // Lấy địa chỉ tài khoản gửi
	GetTxType() string          // Lấy loại giao dịch
}

// BaseTransaction chứa các trường chung cho mọi giao dịch
type BaseTransaction struct {
	TxType    string    `json:"tx_type"`   // Loại giao dịch (KYCSet, Payment, ...)
	Account   string    `json:"account"`   // Tài khoản gửi giao dịch
	Sequence  uint32    `json:"sequence"`  // Số thứ tự giao dịch
	Fee       uint64    `json:"fee"`       // Phí giao dịch (drops)
	Timestamp time.Time `json:"timestamp"` // Thời gian gửi
	Signature []byte    `json:"signature"` // Chữ ký tài khoản (ECDSA)
}

// GetAccount trả về tài khoản gửi
func (b *BaseTransaction) GetAccount() string {
	return b.Account
}

// GetTxType trả về loại giao dịch
func (b *BaseTransaction) GetTxType() string {
	return b.TxType
}

// Amount biểu thị số tiền và loại tiền
type Amount struct {
	Value    uint64 `json:"value"`    // Số tiền (drops cho EZC)
	Currency string `json:"currency"` // Loại tiền (EZC, USD, ...)
	Issuer   string `json:"issuer"`   // Nhà phát hành (nếu không phải EZC)
}

// Path biểu thị đường dẫn thanh toán (cho cross-currency)
type Path struct {
	// Chi tiết đường dẫn sẽ được mở rộng sau
	Accounts []string `json:"accounts"` // Tài khoản trung gian
}

// SignerEntry biểu thị một người ký trong multi-signature
type SignerEntry struct {
	Account string `json:"account"` // Tài khoản ký
	Weight  uint32 `json:"weight"`  // Trọng số
}

// KYCSet cập nhật thông tin KYC cho tài khoản
type KYCSet struct {
	BaseTransaction
	KYCData      kyc.KYCData `json:"kyc_data"`      // Thông tin KYC
	KYCHash      []byte      `json:"kyc_hash"`      // Hash của KYCData
	KYCSignature []byte      `json:"kyc_signature"` // Chữ ký KYC provider
}

// Serialize chuyển KYCSet thành JSON
func (k *KYCSet) Serialize() ([]byte, error) {
	return json.Marshal(k)
}

// Payment chuyển tiền giữa các tài khoản
type Payment struct {
	BaseTransaction
	Destination string `json:"destination"` // Tài khoản nhận
	Amount      Amount `json:"amount"`      // Số tiền và loại tiền
	Paths       []Path `json:"paths"`       // Đường dẫn thanh toán
}

// Serialize chuyển Payment thành JSON
func (p *Payment) Serialize() ([]byte, error) {
	return json.Marshal(p)
}

// TrustSet thiết lập trust line với nhà phát hành
type TrustSet struct {
	BaseTransaction
	LimitAmount Amount `json:"limit_amount"` // Giới hạn trust
	QualityIn   uint32 `json:"quality_in"`   // Tỷ giá vào
	QualityOut  uint32 `json:"quality_out"`  // Tỷ giá ra
	Flags       uint32 `json:"flags"`        // Cờ (Rippling, NoRipple, ...)
}

// Serialize chuyển TrustSet thành JSON
func (t *TrustSet) Serialize() ([]byte, error) {
	return json.Marshal(t)
}

// OfferCreate tạo lệnh mua/bán trên DEX
type OfferCreate struct {
	BaseTransaction
	TakerPays  Amount `json:"taker_pays"` // Tài sản muốn nhận
	TakerGets  Amount `json:"taker_gets"` // Tài sản muốn bán
	Expiration uint32 `json:"expiration"` // Thời gian hết hạn
	OfferSeq   uint32 `json:"offer_seq"`  // Số thứ tự offer
}

// Serialize chuyển OfferCreate thành JSON
func (o *OfferCreate) Serialize() ([]byte, error) {
	return json.Marshal(o)
}

// OfferCancel hủy lệnh trên DEX
type OfferCancel struct {
	BaseTransaction
	OfferSeq uint32 `json:"offer_seq"` // Số thứ tự offer cần hủy
}

// Serialize chuyển OfferCancel thành JSON
func (o *OfferCancel) Serialize() ([]byte, error) {
	return json.Marshal(o)
}

// AccountSet cập nhật thuộc tính tài khoản
type AccountSet struct {
	BaseTransaction
	Flags      uint32 `json:"flags"`       // Cờ (RequireDest, RequireAuth, ...)
	ClearFlags uint32 `json:"clear_flags"` // Xóa cờ
	Domain     string `json:"domain"`      // Domain liên kết
	EmailHash  []byte `json:"email_hash"`  // Hash email
}

// Serialize chuyển AccountSet thành JSON
func (a *AccountSet) Serialize() ([]byte, error) {
	return json.Marshal(a)
}

// SignerListSet thiết lập danh sách người ký
type SignerListSet struct {
	BaseTransaction
	SignerQuorum  uint32        `json:"signer_quorum"`  // Số chữ ký tối thiểu
	SignerEntries []SignerEntry `json:"signer_entries"` // Danh sách người ký
}

// Serialize chuyển SignerListSet thành JSON
func (s *SignerListSet) Serialize() ([]byte, error) {
	return json.Marshal(s)
}

// EscrowCreate khóa tiền có điều kiện
type EscrowCreate struct {
	BaseTransaction
	Amount      Amount `json:"amount"`       // Số tiền khóa
	Destination string `json:"destination"`  // Tài khoản nhận
	CancelAfter uint32 `json:"cancel_after"` // Thời gian hủy
	FinishAfter uint32 `json:"finish_after"` // Thời gian hoàn thành
	Condition   []byte `json:"condition"`    // Điều kiện (hash)
}

// Serialize chuyển EscrowCreate thành JSON
func (e *EscrowCreate) Serialize() ([]byte, error) {
	return json.Marshal(e)
}

// EscrowFinish hoàn thành Escrow
type EscrowFinish struct {
	BaseTransaction
	Owner       string `json:"owner"`       // Tài khoản tạo Escrow
	OfferSeq    uint32 `json:"offer_seq"`   // Số thứ tự Escrow
	Condition   []byte `json:"condition"`   // Điều kiện
	Fulfillment []byte `json:"fulfillment"` // Chứng minh điều kiện
}

// Serialize chuyển EscrowFinish thành JSON
func (e *EscrowFinish) Serialize() ([]byte, error) {
	return json.Marshal(e)
}

// EscrowCancel hủy Escrow
type EscrowCancel struct {
	BaseTransaction
	Owner    string `json:"owner"`     // Tài khoản tạo Escrow
	OfferSeq uint32 `json:"offer_seq"` // Số thứ tự Escrow
}

// Serialize chuyển EscrowCancel thành JSON
func (e *EscrowCancel) Serialize() ([]byte, error) {
	return json.Marshal(e)
}

// PaymentChannelCreate tạo kênh thanh toán
type PaymentChannelCreate struct {
	BaseTransaction
	Amount      Amount `json:"amount"`       // Số tiền khóa
	Destination string `json:"destination"`  // Tài khoản nhận
	SettleDelay uint32 `json:"settle_delay"` // Thời gian thanh toán
	PublicKey   []byte `json:"public_key"`   // Khóa công khai kênh
}

// Serialize chuyển PaymentChannelCreate thành JSON
func (p *PaymentChannelCreate) Serialize() ([]byte, error) {
	return json.Marshal(p)
}

// PaymentChannelFund nạp tiền vào kênh
type PaymentChannelFund struct {
	BaseTransaction
	Channel []byte `json:"channel"` // ID kênh
	Amount  Amount `json:"amount"`  // Số tiền nạp
}

// Serialize chuyển PaymentChannelFund thành JSON
func (p *PaymentChannelFund) Serialize() ([]byte, error) {
	return json.Marshal(p)
}

// PaymentChannelClaim yêu cầu/đóng kênh thanh toán
type PaymentChannelClaim struct {
	BaseTransaction
	Channel   []byte `json:"channel"`   // ID kênh
	Balance   Amount `json:"balance"`   // Số tiền yêu cầu
	Signature []byte `json:"signature"` // Chữ ký kênh
	Close     bool   `json:"close"`     // Đóng kênh
}

// Serialize chuyển PaymentChannelClaim thành JSON
func (p *PaymentChannelClaim) Serialize() ([]byte, error) {
	return json.Marshal(p)
}
