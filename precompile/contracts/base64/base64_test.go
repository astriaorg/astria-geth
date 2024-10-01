package base64

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/mocks"
	"github.com/holiman/uint256"
)

func NewMockStatefulContext() precompile.StatefulContext {
	return precompile.NewStatefulContext(
		mocks.NewMockStateDB(),
		common.BytesToAddress([]byte("0xSelf")),
		common.BytesToAddress([]byte("0xMsgSender")),
		uint256.NewInt(0),
	)
}
func TestEncode(t *testing.T) {
	c := NewBase64()
	ctx := NewMockStatefulContext()

	input := []byte("Hello, Astria!")

	// Test for Non-Error
	output, err := c.Encode(ctx, input)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Test for Correct Encoding
	expectedOutput := "SGVsbG8sIEFzdHJpYSE="
	if output != expectedOutput {
		t.Errorf("Encode output mismatch. Got %v, expected %v", output, expectedOutput)
	}
}

func TestDecode(t *testing.T) {
	c := NewBase64()
	ctx := NewMockStatefulContext()

	input := "SGVsbG8sIFJvbGx1cCE="
	expectedOutput := []byte("Hello, Rollup!")

	output, err := c.Decode(ctx, input)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("Decode output mismatch. Got %v, expected %v", output, expectedOutput)
	}
}

func TestEncodeURL(t *testing.T) {
	c := NewBase64()
	ctx := NewMockStatefulContext()

	input := []byte{255, 0, 23, 40, 33, 32, 1, 56, 89, 23, 156, 21}

	// Test for Non-Error
	output, err := c.EncodeURL(ctx, input)
	if err != nil {
		t.Fatalf("EncodeURL failed: %v", err)
	}

	// Test for Correct Encoding
	expectedOutput := "_wAXKCEgAThZF5wV"
	if output != expectedOutput {
		t.Errorf("EncodeURL output mismatch. Got %v, expected %v", output, expectedOutput)
	}
}

func TestDecodeURL(t *testing.T) {
	c := NewBase64()
	ctx := NewMockStatefulContext()

	input := "SGVsbG8sIEFzdHJpYSE="
	expectedOutput := []byte("Hello, Astria!")

	output, err := c.DecodeURL(ctx, input)
	if err != nil {
		t.Fatalf("DecodeURL failed: %v", err)
	}

	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("DecodeURL output mismatch. Got %v, expected %v", output, expectedOutput)
	}
}

func TestEncodeRequiredGas(t *testing.T) {
	c := NewBase64()
	ctx := NewMockStatefulContext()

	input := []byte("Hello, Astria!")
	expectedGas := uint64(2)

	gas := c.EncodeRequiredGas(ctx, input)
	if gas != expectedGas {
		t.Errorf("EncodeRequiredGas mismatch. Got %v, expected %v", gas, expectedGas)
	}
}
