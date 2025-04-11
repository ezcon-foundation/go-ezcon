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

package kyc

import (
	"encoding/json"
	"time"
)

// KYCData chứa thông tin KYC của người dùng
type KYCData struct {
	FullName      string `json:"full_name"`      // Họ và tên
	IDNumber      string `json:"id_number"`      // Số CCCD/hộ chiếu
	DateOfBirth   string `json:"date_of_birth"`  // Ngày sinh (YYYY-MM-DD)
	Nationality   string `json:"nationality"`    // Quốc tịch
	Address       string `json:"address"`        // Địa chỉ thường trú
	BiometricHash []byte `json:"biometric_hash"` // Hash của dữ liệu sinh trắc học
	IsEncrypted   bool   `json:"is_encrypted"`   // Dữ liệu có mã hóa không
}

// KYCSet là giao dịch để thiết lập/cập nhật KYC
type KYCSet struct {
	Account      string    `json:"account"`       // Địa chỉ tài khoản
	KYCData      KYCData   `json:"kyc_data"`      // Thông tin KYC (mã hóa hoặc rỗng)
	KYCHash      []byte    `json:"kyc_hash"`      // Hash của KYCData (SHA-256)
	KYCSignature []byte    `json:"kyc_signature"` // Chữ ký của KYC provider
	Sequence     uint32    `json:"sequence"`      // Số thứ tự giao dịch
	Fee          uint64    `json:"fee"`           // Phí giao dịch (drops)
	Timestamp    time.Time `json:"timestamp"`     // Thời gian gửi
	Signature    []byte    `json:"signature"`     // Chữ ký tài khoản
}

// Serialize KYCSet để ký hoặc lưu trữ
func (k *KYCSet) Serialize() ([]byte, error) {
	return json.Marshal(k)
}
