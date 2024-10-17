package precompile

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Gas costs
const (
	GasFree        uint64 = 0
	GasQuickStep   uint64 = 2
	GasFastestStep uint64 = 3
	GasFastStep    uint64 = 5
	GasMidStep     uint64 = 8
	GasSlowStep    uint64 = 10
	GasExtStep     uint64 = 20
)

func WordLength(input []byte, wordSize uint64) uint64 {
	return (uint64(len(input)) + wordSize - 1) / wordSize
}

type statefulPrecompiledContract struct {
	abi abi.ABI
}

func NewStatefulPrecompiledContract(abiStr string) StatefulPrecompiledContract {
	abi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		panic(err)
	}
	return &statefulPrecompiledContract{
		abi: abi,
	}
}

func (spc *statefulPrecompiledContract) GetABI() abi.ABI {
	return spc.abi
}

// This is a placeholder implementation. The actual gas required would depend on the specific contract.
// You should replace this with the actual implementation.
func (spc *statefulPrecompiledContract) DefaultGas(input []byte) uint64 {
	return GasFree
}
