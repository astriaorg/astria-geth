package config

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/precompile"
)

var NullPrecompiles = precompile.PrecompileMap{}

// return precompiles that are enabled at height
func PrecompileConfig(chainConfig *params.ChainConfig, height uint64, timestamp uint64) precompile.PrecompileMap {
	// Example, enable the base64 precompile at address 0x0000000000000000000000000000000000001000:
	// (add `import pcbase64 "github.com/ethereum/go-ethereum/precompile/contracts/base64"`)
	// return precompile.PrecompileMap{
	// 	common.HexToAddress("0x01000"): pcbase64.NewBase64(),
	// }

	return NullPrecompiles
}
