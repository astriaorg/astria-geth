package base64

import (
	b64 "encoding/base64"

	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/abi"
)

type Base64 struct {
	precompile.StatefulPrecompiledContract
}

func NewBase64() *Base64 {
	return &Base64{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			abi.Base64ABI,
		),
	}
}

func (c *Base64) Encode(ctx precompile.StatefulContext, data []byte) (string, error) {
	return b64.StdEncoding.EncodeToString(data), nil
}

func (c *Base64) EncodeURL(ctx precompile.StatefulContext, data []byte) (string, error) {
	return b64.URLEncoding.EncodeToString(data), nil
}

func (c *Base64) Decode(ctx precompile.StatefulContext, data string) ([]byte, error) {
	decoded, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func (c *Base64) DecodeURL(ctx precompile.StatefulContext, data string) ([]byte, error) {
	decoded, err := b64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

// gas calcs
func (c *Base64) EncodeRequiredGas(ctx precompile.StatefulContext, data []byte) uint64 {
	return precompile.WordLength(data, 256) * precompile.GasQuickStep
}

func (c *Base64) EncodeURLRequiredGas(ctx precompile.StatefulContext, data []byte) uint64 {
	return precompile.WordLength(data, 256) * precompile.GasQuickStep
}

func (c *Base64) DecodeRequiredGas(ctx precompile.StatefulContext, data string) uint64 {
	return precompile.WordLength([]byte(data), 256) * precompile.GasQuickStep
}

func (c *Base64) DecodeURLRequiredGas(ctx precompile.StatefulContext, data string) uint64 {
	return precompile.WordLength([]byte(data), 256) * precompile.GasQuickStep
}
