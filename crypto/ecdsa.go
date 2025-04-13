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

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// GenerateKeyPair tạo cặp khóa ECDSA
func GenerateKeyPair() (PublicKey, PrivateKey, error) {
	curve := secp256k1.S256()
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate key pair: %v", err)
	}

	pubKey := privKey.PublicKey
	pubKeyBytes := elliptic.Marshal(curve, pubKey.X, pubKey.Y)
	return pubKeyBytes, privKey.D.Bytes(), nil
}

// Sign ký dữ liệu bằng khóa riêng
func Sign(data []byte, privKey PrivateKey) (Signature, error) {
	curve := secp256k1.S256()
	ecdsaPrivKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{Curve: curve},
		D:         new(big.Int).SetBytes(privKey),
	}

	// Hash dữ liệu trước khi ký
	hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, ecdsaPrivKey, hash[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %v", err)
	}

	// Chuyển r, s thành byte slice
	sig := make([]byte, 64)
	rBytes := r.Bytes()
	sBytes := s.Bytes()
	copy(sig[32-len(rBytes):32], rBytes)
	copy(sig[64-len(sBytes):], sBytes)
	return sig, nil
}

// Verify xác minh chữ ký
func Verify(data []byte, sig Signature, pubKey PublicKey) bool {
	if len(sig) != 64 {
		return false
	}

	curve := secp256k1.S256()
	x, y := elliptic.Unmarshal(curve, pubKey)
	if x == nil || y == nil {
		return false
	}

	ecdsaPubKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// Hash dữ liệu
	hash := sha256.Sum256(data)

	// Tách r, s từ chữ ký
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])
	return ecdsa.Verify(ecdsaPubKey, hash[:], r, s)
}

// PubKeyFromPrivKey lấy khóa công khai từ khóa riêng
func PubKeyFromPrivKey(privKey PrivateKey) (PublicKey, error) {
	curve := secp256k1.S256()
	ecdsaPrivKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{Curve: curve},
		D:         new(big.Int).SetBytes(privKey),
	}

	pubKey := ecdsaPrivKey.PublicKey
	pubKeyBytes := elliptic.Marshal(curve, pubKey.X, pubKey.Y)
	return pubKeyBytes, nil
}
