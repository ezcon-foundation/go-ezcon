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
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/crypto/blake256"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"golang.org/x/crypto/sha3"
)

// signMessage signs a message using the private key and returns the signature in hex
func signMessage(message []byte, privKey *secp256k1.PrivateKey) ([]byte, error) {
	// Hash the message using Keccak-256 (Ethereum standard)
	messageHash := blake256.Sum256(message)

	// Sign the message hash
	signature := ecdsa.Sign(privKey, messageHash[:])

	// Convert to hex
	return signature.Serialize(), nil
}

// publicKeyToAddress generates an Ethereum address from a secp256k1 public key
func publicKeyToAddress(pubKey *secp256k1.PublicKey) string {
	// Get uncompressed public key bytes (excluding the 0x04 prefix)
	pubBytes := pubKey.SerializeUncompressed()[1:] // Remove 0x04 prefix

	// Compute Keccak-256 hash of the public key
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubBytes)
	hashSum := hash.Sum(nil)

	// Take the last 20 bytes of the hash as the address
	addressBytes := hashSum[len(hashSum)-20:]

	// Convert to hexadecimal with 0x prefix
	return "0x" + hex.EncodeToString(addressBytes)
}

// compressedPubKeyToAddress generates an Ethereum address from a compressed public key (hex string)
func compressedPubKeyToAddress(pubKeyHex string) (string, error) {
	// Decode hex string to bytes
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key hex: %v", err)
	}

	// Parse compressed public key
	pubKey, err := secp256k1.ParsePubKey(pubKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse compressed public key: %v", err)
	}

	// Get uncompressed public key bytes (excluding the 0x04 prefix)
	pubBytes := pubKey.SerializeUncompressed()[1:] // Remove 0x04 prefix

	// Compute Keccak-256 hash of the public key
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubBytes)
	hashSum := hash.Sum(nil)

	// Take the last 20 bytes of the hash as the address
	addressBytes := hashSum[len(hashSum)-20:]

	// Convert to hexadecimal with 0x prefix
	return "0x" + hex.EncodeToString(addressBytes), nil
}
