package config

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/precompile"
	pcbase64 "github.com/ethereum/go-ethereum/precompile/contracts/base64"
)

var NullPrecompiles = precompile.PrecompileMap{}

// Cache of initialized precompiles for each fork
var precompileCache = make(map[string]precompile.PrecompileMap)

// return precompiles that are enabled at height
func GetPrecompiles(chainConfig *params.ChainConfig, height uint64, timestamp uint64) precompile.PrecompileMap {
	fork := chainConfig.GetAstriaForks().GetForkAtHeight(height)

	// Return early if no precompiles configured
	if len(fork.Precompiles) == 0 {
		return NullPrecompiles
	}

	// Check if we've already initialized precompiles for this fork
	if cached, exists := precompileCache[fork.Name]; exists {
		return cached
	}

	// Initialize precompiles for this fork
	precompiles := make(precompile.PrecompileMap)
	for addr, precompileType := range fork.Precompiles {
		switch *precompileType {
		case params.PrecompileBase64:
			precompiles[addr] = pcbase64.NewBase64()
		default:
			log.Error("Unknown precompile type", "type", precompileType, "address", addr)
		}
	}

	// Cache the initialized precompiles
	precompileCache[fork.Name] = precompiles
	return precompiles
}
