# Writing a Precompile Contract

1. Create a Solidity interface in `contracts/interfaces`, e.g, IBase64.sol

2. Generate bindings with `./gen.sh`

3. Copy generate `ABI` definition from `./bindings/i_<precompile>.abigen.go` to `./abi/abi.go`, assigning it to a unique const name, i.e, `Base64ABI`.

4. Implement the precompile in Go at `./contracts/<precompile>/<precompile.go>`.
  - The struct should implement the `StatefulPrecompiledContract` interface
  - You must methods defined in the Solidity interface
  - Implement custom gas handlers as needed
  - You can use the `StatefulContext` to access and modify the evm state db

5. Enable the precompile by returning it from `PrecompileConfig()` in `./config/config.go`. Existing chains should only enable new precompiles at hard forks.

   For example, to enable the example base64 precompile at genesis, the config could look like this:

   ```go
   package config

   import (
   	"github.com/ethereum/go-ethereum/common"
   	"github.com/ethereum/go-ethereum/params"
   	"github.com/ethereum/go-ethereum/precompile"

   	pcbase64 "github.com/ethereum/go-ethereum/precompile/contracts/base64"
   )

   var NullPrecompiles = precompile.PrecompileMap{}

   var MyRollupGenesisPrecompiles = precompile.PrecompileMap{
   	common.HexToAddress("0x01000"): pcbase64.NewBase64(),
   }

   // return precompiles that are enabled at height
   func PrecompileConfig(chainConfig *params.ChainConfig, height uint64,    timestamp uint64) precompile.PrecompileMap {
   	return MyRollupGenesisPrecompiles
   }
   ```
