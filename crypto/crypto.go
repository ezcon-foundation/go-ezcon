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
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"math/big"
)

// SHA256 tạo hash SHA-256
func SHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// SHA512_256 tạo hash SHA-512/256 (như XRPL)
func SHA512_256(data []byte) []byte {
	hash := sha512.Sum512_256(data)
	return hash[:]
}

// Sign ký dữ liệu bằng private key (hex string)
func Sign(data, privKey []byte) ([]byte, error) {
	// Decode hex string
	privKeyBytes, err := hex.DecodeString(string(privKey))
	if err != nil {
		return nil, errors.New("invalid hex private key")
	}

	// Parse private key
	priv, err := parsePrivateKey(privKeyBytes)
	if err != nil {
		return nil, err
	}

	// Hash dữ liệu
	hash := SHA512_256(data)

	// Ký bằng ECDSA
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		return nil, err
	}

	// Gộp r, s thành chữ ký
	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

// Verify xác minh chữ ký
func Verify(data, sig, pubKey []byte) bool {
	// Parse public key
	pub, err := parsePublicKey(pubKey)
	if err != nil {
		return false
	}

	// Parse chữ ký
	if len(sig) != 64 {
		return false
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	// Hash dữ liệu
	hash := SHA512_256(data)

	// Xác minh ECDSA
	return ecdsa.Verify(pub, hash, r, s)
}

// PubKeyFromPrivKey lấy public key từ private key (hex string)
func PubKeyFromPrivKey(privKey []byte) ([]byte, error) {
	// Decode hex string
	privKeyBytes, err := hex.DecodeString(string(privKey))
	if err != nil {
		return nil, errors.New("invalid hex private key")
	}

	// Parse private key
	priv, err := parsePrivateKey(privKeyBytes)
	if err != nil {
		return nil, err
	}
	pub := &priv.PublicKey
	return elliptic.Marshal(elliptic.P256(), pub.X, pub.Y), nil
}

// PubKeyFromNode lấy public key từ node ID
func PubKeyFromNode(node string) ([]byte, error) {
	// Giả lập: node ID là hex của public key
	pubKey, err := hex.DecodeString(node)
	if err != nil {
		return nil, err
	}
	if _, err := parsePublicKey(pubKey); err != nil {
		return nil, err
	}
	return pubKey, nil
}

// parsePrivateKey parse private key từ bytes
func parsePrivateKey(privKey []byte) (*ecdsa.PrivateKey, error) {
	if len(privKey) != 32 {
		return nil, errors.New("invalid private key length")
	}
	priv := &ecdsa.PrivateKey{
		D: new(big.Int).SetBytes(privKey),
	}
	priv.PublicKey.Curve = elliptic.P256()
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(privKey)
	return priv, nil
}

// parsePublicKey parse public key từ bytes
func parsePublicKey(pubKey []byte) (*ecdsa.PublicKey, error) {
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKey)
	if x == nil || y == nil {
		return nil, errors.New("invalid public key")
	}
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}, nil
}
