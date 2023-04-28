package util

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/big"
)

// TxToCall converts a transaction to a call.
func TxToCall(transaction Transaction) (*ethereum.CallMsg, error) {
	rawMsg, err := transaction.AsMessage(types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not convert to call")
	}

	return &ethereum.CallMsg{
		From:      rawMsg.From(),
		To:        transaction.To(),
		Gas:       transaction.Gas(),
		GasPrice:  transaction.GasPrice(),
		GasFeeCap: transaction.GasTipCap(),
		GasTipCap: transaction.GasFeeCap(),
		Value:     transaction.Value(),
		Data:      transaction.Data(),
	}, nil
}

// Transaction is an interface for everything needed to convert a transaction to a call.
type Transaction interface {
	// ChainId returns the chain id.
	ChainId() *big.Int
	// To returns the to address.
	To() *common.Address
	// GasPrice returns the gas price.
	GasPrice() *big.Int
	// GasFeeCap returns the gas fee cap.
	GasFeeCap() *big.Int
	// GasTipCap returns the gas tip cap.
	GasTipCap() *big.Int
	// Gas returns the gas limit.
	Gas() uint64
	// Value returns the value of the tx.
	Value() *big.Int
	// Data returns the data of the tx.
	Data() []byte
	// AsMessage converts the transaction to a message.
	AsMessage(s types.Signer, baseFee *big.Int) (types.Message, error)
}

var _ Transaction = &types.Transaction{}
