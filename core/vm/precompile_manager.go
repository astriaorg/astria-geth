package vm

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/holiman/uint256"
)

type methodID [4]byte

type statefulMethod struct {
	abiMethod     abi.Method
	reflectMethod reflect.Method
}

type precompileMethods map[methodID]*statefulMethod
type gasMethods map[methodID]reflect.Method

type precompileManager struct {
	evm         *EVM
	precompiles map[common.Address]precompile.StatefulPrecompiledContract
	pMethods    map[common.Address]precompileMethods
	gMethods    map[common.Address]gasMethods
}

func NewPrecompileManager(evm *EVM) *precompileManager {
	precompiles := make(map[common.Address]precompile.StatefulPrecompiledContract)
	pMethods := make(map[common.Address]precompileMethods)
	gMethods := make(map[common.Address]gasMethods)
	return &precompileManager{
		evm:         evm,
		precompiles: precompiles,
		pMethods:    pMethods,
		gMethods:    gMethods,
	}
}

func (pm *precompileManager) IsPrecompile(addr common.Address) bool {
	_, isEvmPrecompile := pm.evm.precompile(addr)
	if isEvmPrecompile {
		return true
	}

	_, isStatefulPrecompile := pm.precompiles[addr]
	return isStatefulPrecompile
}

func (pm *precompileManager) Run(
	addr common.Address,
	input []byte,
	caller common.Address,
	value *uint256.Int,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {

	// run core evm precompile
	p, isEvmPrecompile := pm.evm.precompile(addr)
	if isEvmPrecompile {
		return RunPrecompiledContract(p, input, suppliedGas, pm.evm.Config.Tracer)
	}

	contract, ok := pm.precompiles[addr]
	if !ok {
		return nil, 0, fmt.Errorf("no precompiled contract at address %v", addr.Hex())
	}

	// Extract the method ID from the input
	methodId := methodID(input)
	// Try to get the method from the precompiled contracts using the method ID
	method, exists := pm.pMethods[addr][methodId]
	if !exists {
		return nil, 0, fmt.Errorf("no method with id %x in precompiled contract at address %v", methodId, addr.Hex())
	}

	// refund gas for act of calling custom precompile
	if suppliedGas > 0 {
		if pm.evm.chainRules.IsEIP150 && suppliedGas > params.CallGasEIP150 {
			suppliedGas += params.CallGasEIP150
		} else if suppliedGas > params.CallGasFrontier {
			suppliedGas += params.CallGasFrontier
		}
	}

	// Unpack the input arguments using the ABI method's inputs
	unpackedArgs, err := method.abiMethod.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, err
	}

	// Convert the unpacked args to reflect values.
	reflectedUnpackedArgs := make([]reflect.Value, 0, len(unpackedArgs))
	for _, unpacked := range unpackedArgs {
		reflectedUnpackedArgs = append(reflectedUnpackedArgs, reflect.ValueOf(unpacked))
	}

	// set precompile nonce to 1 to avoid state deletion for being considered an empty account
	// this conforms precompile contracts to EIP-161
	if !readOnly && pm.evm.StateDB.GetNonce(addr) == 0 {
		pm.evm.StateDB.SetNonce(addr, 1)
	}

	ctx := precompile.NewStatefulContext(pm.evm.StateDB, addr, caller, value)

	// Make sure the readOnly is only set if we aren't in readOnly yet.
	// This also makes sure that the readOnly flag isn't removed for child calls.
	if readOnly && !ctx.IsReadOnly() {
		ctx.SetReadOnly(true)
		defer func() { ctx.SetReadOnly(false) }()
	}

	// check if enough gas is supplied
	var gasCost uint64 = contract.DefaultGas(input)
	gasMethod, exists := pm.gMethods[addr][methodId]
	if exists {
		gasResult := gasMethod.Func.Call(append(
			[]reflect.Value{
				reflect.ValueOf(contract),
				reflect.ValueOf(ctx),
			},
			reflectedUnpackedArgs...,
		))
		if len(gasResult) > 0 {
			gasCost, ok = gasResult[0].Interface().(uint64)
			if !ok {
				gasCost = contract.DefaultGas(input)
			}
		}
	}

	if gasCost > suppliedGas {
		return nil, 0, ErrOutOfGas
	}

	// call the precompile method
	results := method.reflectMethod.Func.Call(append(
		[]reflect.Value{
			reflect.ValueOf(contract),
			reflect.ValueOf(ctx),
		},
		reflectedUnpackedArgs...,
	))

	// check if precompile returned an error
	if len(results) > 0 {
		if err, ok := results[len(results)-1].Interface().(error); ok && err != nil {
			return nil, 0, err
		}
	}

	// Pack the result
	var output []byte
	if len(results) > 1 {
		interfaceArgs := make([]interface{}, len(results)-1)
		for i, v := range results[:len(results)-1] {
			interfaceArgs[i] = v.Interface()
		}
		output, err = method.abiMethod.Outputs.Pack(interfaceArgs...)
		if err != nil {
			return nil, 0, err
		}
	}

	suppliedGas -= gasCost
	return output, suppliedGas, nil
}

func (pm *precompileManager) RegisterMap(m precompile.PrecompileMap) error {
	for addr, p := range m {
		err := pm.Register(addr, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pm *precompileManager) Register(addr common.Address, p precompile.StatefulPrecompiledContract) error {
	if _, isEvmPrecompile := pm.evm.precompile(addr); isEvmPrecompile {
		return fmt.Errorf("precompiled contract already exists at address %v", addr.Hex())
	}

	if _, exists := pm.precompiles[addr]; exists {
		return fmt.Errorf("precompiled contract already exists at address %v", addr.Hex())
	}

	// niaeve implementation; parsed abi method names must match precompile method names 1:1
	//
	// Note on method naming:
	// Method name is the abi method name used for internal representation. It's derived from
	// the abi raw name and a suffix will be added in the case of a function overload.
	//
	// e.g.
	// These are two functions that have the same name:
	// * foo(int,int)
	// * foo(uint,uint)
	// The method name of the first one will be resolved as Foo while the second one
	// will be resolved as Foo0.
	//
	// Alternatively could require each precompile to define the func mapping instead of doing this magic
	abiMethods := p.GetABI().Methods
	contractType := reflect.ValueOf(p).Type()
	precompileMethods := make(precompileMethods)
	gasMethods := make(gasMethods)
	for _, abiMethod := range abiMethods {
		mName := strings.ToUpper(string(abiMethod.Name[0])) + abiMethod.Name[1:]
		reflectMethod, exists := contractType.MethodByName(mName)
		if !exists {
			return fmt.Errorf("precompiled contract does not implement abi method %s with signature %s", abiMethod.Name, abiMethod.RawName)
		}
		mID := methodID(abiMethod.ID)
		precompileMethods[mID] = &statefulMethod{
			abiMethod:     abiMethod,
			reflectMethod: reflectMethod,
		}

		// precompile method has custom gas calc
		gName := mName + "RequiredGas"
		gasMethod, exists := contractType.MethodByName(gName)
		if exists {
			if gasMethod.Type.NumOut() != 1 || gasMethod.Type.Out(0).Kind() != reflect.Uint64 {
				return fmt.Errorf("gas method %s does not return uint64", gName)
			}
			gasMethods[mID] = gasMethod
		}
	}

	pm.precompiles[addr] = p
	pm.pMethods[addr] = precompileMethods
	pm.gMethods[addr] = gasMethods
	return nil
}
