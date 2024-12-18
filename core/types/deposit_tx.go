package types

import (
	"bytes"
	"math/big"

	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

var _ TxData = &InjectedTx{}

type InjectedTx struct {
	// the caller address set in the genesis file (either bridge or oracle)
	// ie. the minter or the caller of the ERC20/oracle contract
	From common.Address
	// value to be minted to the recipient, if this is a native asset mint
	Value *big.Int
	// gas limit
	Gas uint64
	// if this is a native asset mint, this is set to the mint recipient
	// if this is an ERC20 mint, this is set to the ERC20 contract address
	// if this is an oracle update, this is set to the oracle contract address
	To *common.Address
	// if this is an ERC20 mint, the following field is set
	// to the `mint` function calldata.
	// if this is an oracle update, the following field is set to the
	// `initializeCurrencyPair` or `updatePriceData` function calldata.
	Data []byte
	// the transaction ID of the source action on the sequencer, consisting
	// of the transaction hash.
	SourceTransactionId primitivev1.TransactionId
	// index of the source action within its sequencer transaction
	SourceTransactionIndex uint64
}

func (tx *InjectedTx) copy() TxData {
	to := new(common.Address)
	if tx.To != nil {
		*to = *tx.To
	}

	cpy := &InjectedTx{
		From:                   tx.From,
		Value:                  new(big.Int),
		Gas:                    tx.Gas,
		To:                     to,
		Data:                   make([]byte, len(tx.Data)),
		SourceTransactionId:    tx.SourceTransactionId,
		SourceTransactionIndex: tx.SourceTransactionIndex,
	}

	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	copy(cpy.Data, tx.Data)
	return cpy
}

func (tx *InjectedTx) txType() byte           { return InjectedTxType }
func (tx *InjectedTx) chainID() *big.Int      { return common.Big0 }
func (tx *InjectedTx) accessList() AccessList { return nil }
func (tx *InjectedTx) data() []byte           { return tx.Data }
func (tx *InjectedTx) gas() uint64            { return tx.Gas }
func (tx *InjectedTx) gasFeeCap() *big.Int    { return new(big.Int) }
func (tx *InjectedTx) gasTipCap() *big.Int    { return new(big.Int) }
func (tx *InjectedTx) gasPrice() *big.Int     { return new(big.Int) }
func (tx *InjectedTx) value() *big.Int        { return tx.Value }
func (tx *InjectedTx) nonce() uint64          { return 0 }
func (tx *InjectedTx) to() *common.Address    { return tx.To }

func (tx *InjectedTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	return dst.Set(new(big.Int))
}

func (tx *InjectedTx) rawSignatureValues() (v, r, s *big.Int) {
	return common.Big0, common.Big0, common.Big0
}

func (tx *InjectedTx) setSignatureValues(chainID, v, r, s *big.Int) {
	// noop
}

func (tx *InjectedTx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, tx)
}

func (tx *InjectedTx) decode(input []byte) error {
	return rlp.DecodeBytes(input, tx)
}
