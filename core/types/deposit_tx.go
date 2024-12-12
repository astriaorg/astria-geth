package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

var _ TxData = &DepositTx{}

type DepositTx struct {
	// the bridge sender address set in the genesis file
	// ie. the minter or the caller of the ERC20 contract
	From common.Address
	// value to be minted to the recipient, if this is a native asset mint
	Value *big.Int
	// gas limit
	Gas uint64
	// if this is a native asset mint, this is set to the mint recipient
	// if this is an ERC20 mint, this is set to the ERC20 contract address
	To *common.Address
	// if this is an ERC20 mint, the following field is set
	// to the `mint` function calldata.
	Data []byte
	// the transaction ID of the source action for the deposit, consisting
	// of the transaction hash.
	SourceTransactionId string
	// index of the deposit's source action within its transaction
	SourceTransactionIndex uint64
}

func (tx *DepositTx) copy() TxData {
	to := new(common.Address)
	if tx.To != nil {
		*to = *tx.To
	}

	cpy := &DepositTx{
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

func (tx *DepositTx) txType() byte           { return DepositTxType }
func (tx *DepositTx) chainID() *big.Int      { return common.Big0 }
func (tx *DepositTx) accessList() AccessList { return nil }
func (tx *DepositTx) data() []byte           { return tx.Data }
func (tx *DepositTx) gas() uint64            { return tx.Gas }
func (tx *DepositTx) gasFeeCap() *big.Int    { return new(big.Int) }
func (tx *DepositTx) gasTipCap() *big.Int    { return new(big.Int) }
func (tx *DepositTx) gasPrice() *big.Int     { return new(big.Int) }
func (tx *DepositTx) value() *big.Int        { return tx.Value }
func (tx *DepositTx) nonce() uint64          { return 0 }
func (tx *DepositTx) to() *common.Address    { return tx.To }

func (tx *DepositTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	return dst.Set(new(big.Int))
}

func (tx *DepositTx) rawSignatureValues() (v, r, s *big.Int) {
	return common.Big0, common.Big0, common.Big0
}

func (tx *DepositTx) setSignatureValues(chainID, v, r, s *big.Int) {
	// noop
}

func (tx *DepositTx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, tx)
}

func (tx *DepositTx) decode(input []byte) error {
	return rlp.DecodeBytes(input, tx)
}
