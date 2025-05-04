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
	"github.com/ezcon-foundation/go-ezcon/crypto/blake256"
	"github.com/ezcon-foundation/go-ezcon/crypto/secp256k1"
	"github.com/ezcon-foundation/go-ezcon/crypto/secp256k1/ecdsa"
	"log"
	"testing"
)

//	func TestPublicKeyConvert(t *testing.T) {
//		publicKeyHex := "02c6c04b1fac83851374aafd2735ebce3ca23cc2c4eb091373e25ef2fe4820c532" // Mẫu compressed
//
//		// Chuyển public key hex sang bytes
//		publicKeyBytes, err := hex.DecodeString(publicKeyHex)
//		if err != nil {
//			t.Errorf("Failed to decode public key hex: %v", err)
//		}
//
//		var pubKey *ecdsa.PublicKey
//		// Compressed public key (33 bytes)
//		x, y := decompressPublicKey(publicKeyBytes, true)
//		pubKey = &ecdsa.PublicKey{
//			Curve: elliptic.P256(),
//			X:     x,
//			Y:     y,
//		}
//
//		address := publicKeyToAddress(pubKey)
//		comparePubkey := "0x649AFc454F4CBAaC9a869F77971Ea02bbA947bf9"
//
//		// Kiểm tra địa chỉ
//		if address != comparePubkey {
//			t.Errorf("Address mismatch: got %s, want %s", address, publicKeyHex)
//		}
//	}

func TestCreateKeyPair(t *testing.T) {

	privKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	privateKeyBytes := privKey.Serialize()

	// Ensure private key is 32 bytes
	if len(privateKeyBytes) != 32 {
		log.Fatalf("Private key must be 32 bytes, got %d bytes", len(privateKeyBytes))
	}

	// Get the public key
	publicKey := privKey.PubKey()

	// Convert private key to hex
	privateKeyHexOutput := hex.EncodeToString(privateKeyBytes)

	// Convert public key to hex (uncompressed, including 0x04 prefix)
	publicKeyBytesUncompressed := publicKey.SerializeUncompressed()
	publicKeyHexUncompressed := hex.EncodeToString(publicKeyBytesUncompressed)

	// Convert public key to hex (compressed)
	publicKeyBytesCompressed := publicKey.SerializeCompressed()
	publicKeyHexCompressed := hex.EncodeToString(publicKeyBytesCompressed)

	// Generate address from public key
	address := publicKeyToAddress(publicKey)

	addressFromCompressed, err := compressedPubKeyToAddress(publicKeyHexCompressed)
	if err != nil {
		log.Fatalf("Failed to generate address from compressed public key: %v", err)
	}

	fmt.Printf("Private Key (Hex): %s\n", privateKeyHexOutput)
	fmt.Printf("Public Key Uncompressed (Hex): %s\n", publicKeyHexUncompressed)
	fmt.Printf("Public Key Compressed (Hex): %s\n", publicKeyHexCompressed)
	fmt.Printf("Address: %s\n", address)
	fmt.Printf("Address (from compressed): %s\n", addressFromCompressed)

	// Private Key (Hex): 4bc30aa5172c5133a04a5873f87fdf9c0043844d611e6951e2553933703eb874
	// Public Key Uncompressed (Hex): 04bf267139cfd2638fb1e8a54865aaa855329167190718a481ff374a44597c7f6cf9ffd7945f84fa617c7b0471515f23000613a3c1ad2117a8ba86db2571110c46
	// Public Key Compressed (Hex): 02bf267139cfd2638fb1e8a54865aaa855329167190718a481ff374a44597c7f6c
	// Address: 0xab79975d250d45f685f65e1b3d41234ced77f040
	// Address (from compressed): 0xab79975d250d45f685f65e1b3d41234ced77f040
}

func TestSignMessage(t *testing.T) {

	privateKeyHex := "4bc30aa5172c5133a04a5873f87fdf9c0043844d611e6951e2553933703eb874"

	// Decode private key from hex
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	// Ensure private key is 32 bytes
	if len(privateKeyBytes) != 32 {
		log.Fatalf("Private key must be 32 bytes, got %d bytes", len(privateKeyBytes))
	}

	// Create secp256k1 private key
	var scalar secp256k1.ModNScalar
	if overflow := scalar.SetByteSlice(privateKeyBytes); overflow {
		log.Fatalf("Invalid private key: overflow")
	}
	privKey := secp256k1.NewPrivateKey(&scalar)

	text := "test message"

	signature, err := signMessage([]byte(text), privKey)
	if err != nil {
		return
	}

	fmt.Println(hex.EncodeToString(signature))

}

func TestVerifyMessage(t *testing.T) {
	// Decode hex-encoded serialized public key.
	pubKeyBytes, err := hex.DecodeString("02bf267139cfd2638fb1e8a54865aaa855329167190718a481ff374a44597c7f6c")
	if err != nil {
		fmt.Println(err)
		return
	}
	pubKey, err := secp256k1.ParsePubKey(pubKeyBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode hex-encoded serialized signature.
	sigBytes, err := hex.DecodeString("304402205898c006fbcefdbdea07c95532ca0c3015694419153a99dd9d02112d3a99a772022071dd1c65de9d9fed8dcf71b8e4955fdb38445f2a86ee652e99e7876e44ace73b")
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, err := ecdsa.ParseDERSignature(sigBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Verify the signature for the message using the public key.
	message := "test message"
	messageHash := blake256.Sum256([]byte(message))
	verified := signature.Verify(messageHash[:], pubKey)
	fmt.Println("Signature Verified?", verified)
}
