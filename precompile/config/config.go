package config

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/precompile"
)

var NullPrecompiles = precompile.PrecompileMap{}

// return precompiles that are enabled at height
func PrecompileConfig(chainConfig *params.ChainConfig, height uint64, timestamp uint64) precompile.PrecompileMap {
	return NullPrecompiles
}
