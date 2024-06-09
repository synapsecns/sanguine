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

	callMessage := ethereum.CallMsg{
		From:      rawMsg.From,
		To:        transaction.To(),
		Gas:       transaction.Gas(),
		GasPrice:  transaction.GasPrice(),
		GasFeeCap: transaction.GasFeeCap(),
		GasTipCap: transaction.GasTipCap(),
		Value:     transaction.Value(),
		Data:      transaction.Data(),
	}

	// gas price will getset to gastipcap + gas fee cap to account for legacy behavior so if
	// tip/fee cap are set we need to make sure we nil gas price
	if transaction.Type() == types.LegacyTxType {
		callMessage.GasTipCap = nil
		callMessage.GasFeeCap = nil
	} else {
		callMessage.GasPrice = nil
	}
	return &callMessage, nil
}
