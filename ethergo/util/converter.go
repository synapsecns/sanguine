package util

import (
	"github.com/ethereum/go-ethereum"
	ethCore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// TxToCall converts a transaction to a call.
func TxToCall(transaction *types.Transaction) (*ethereum.CallMsg, error) {
	rawMsg, err := ethCore.TransactionToMessage(transaction, types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not convert to call")
	}

	return &ethereum.CallMsg{
		From:      rawMsg.From,
		To:        transaction.To(),
		Gas:       transaction.Gas(),
		GasPrice:  transaction.GasPrice(),
		GasFeeCap: transaction.GasTipCap(),
		GasTipCap: transaction.GasFeeCap(),
		Value:     transaction.Value(),
		Data:      transaction.Data(),
	}, nil
}
