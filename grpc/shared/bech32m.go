package shared

// Copied from astria-cli-go bech32m module (https://github.com/astriaorg/astria-cli-go/blob/d5ef82f718325b2907634c108d42b503211c20e6/modules/bech32m/bech32m.go#L1)
// TODO: organize the bech32m usage throughout the codebase

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/bech32"
)

type Address struct {
	address string
	prefix  string
	bytes   [20]byte
}

// String returns the bech32m address as a string
func (a *Address) String() string {
	return a.address
}

// Prefix returns the prefix of the bech32m address
func (a *Address) Prefix() string {
	return a.prefix
}

// Bytes returns the underlying bytes for the bech32m address as a [20]byte array
func (a *Address) Bytes() [20]byte {
	return a.bytes
}

// ValidateBech32mAddress verifies that a string in a valid bech32m address. It
// will return nil if the address is valid, otherwise it will return an error.
func ValidateBech32mAddress(address string, intendedPrefix string) error {
	prefix, byteAddress, version, err := bech32.DecodeGeneric(address)
	if err != nil {
		return fmt.Errorf("address must be a bech32 encoded string")
	}
	if version != bech32.VersionM {
		return fmt.Errorf("address must be a bech32m address")
	}
	byteAddress, err = bech32.ConvertBits(byteAddress, 5, 8, false)
	if err != nil {
		return fmt.Errorf("failed to convert address to 8 bit")
	}
	if prefix == "" {
		return fmt.Errorf("address must have prefix")
	}
	if prefix != intendedPrefix {
		return fmt.Errorf("address must have prefix %s", intendedPrefix)
	}

	if len(byteAddress) != 20 {
		return fmt.Errorf("address must decode to a 20 length byte array: got len %d", len(byteAddress))
	}

	return nil
}

// EncodeFromBytes creates a *Address from a [20]byte array and string
// prefix.
func EncodeFromBytes(prefix string, data [20]byte) (string, error) {
	// Convert the data from 8-bit groups to 5-bit
	convertedBytes, err := bech32.ConvertBits(data[:], 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("failed to convert bits from 8-bit groups to 5-bit groups: %v", err)
	}

	// Encode the data as bech32m
	address, err := bech32.EncodeM(prefix, convertedBytes)
	if err != nil {
		return "", fmt.Errorf("failed to encode address as bech32m: %v", err)
	}

	return address, nil
}

// EncodeFromPublicKey takes an ed25519 public key and string prefix and encodes
// them into a *Address.
func EncodeFromPublicKey(prefix string, pubkey ed25519.PublicKey) (string, error) {
	hash := sha256.Sum256(pubkey)
	var addr [20]byte
	copy(addr[:], hash[:20])
	address, err := EncodeFromBytes(prefix, addr)
	if err != nil {
		return "", err
	}
	return address, nil
}
