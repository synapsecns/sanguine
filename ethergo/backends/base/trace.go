package base

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/parser/tracely"
)

// getTransactionLabels gets the labels for all contracts used in this transaction.
// if a list of labels cannot be obtained using a trace all registered contracts will be returned.
//
// TODO: this should be able to work with a simulated backend but can't right now.
// returns address->label
func (b *Backend) getTransactionLabelMap(tx *types.Transaction, from common.Address) (res map[common.Address]string) {
	b.verifiedMux.RLock()
	defer b.verifiedMux.RUnlock()

	res = map[common.Address]string{}
	txResult, err := tracely.GetTxResult(b.RPCAddress(), tx.Hash().String())
	if err != nil {
		logger.Warnf("could not get tx result: %v", err)
	}

	res[from] = "sender"

	// if we can't get the tx result, we'll just return all the verified contracts.
	if err != nil || b.RPCAddress() == "" {
		for _, contract := range b.verifiedContracts {
			res[contract.Address()] = contract.ContractName()
		}
		return res
	}

	if tx.To() != nil {
		if contract, ok := b.verifiedContracts[*tx.To()]; ok {
			res[*tx.To()] = contract.ContractName()
		}
	}

	// walk through each frame of the transaction and get the contract address
	for _, frame := range txResult {
		stackRes := map[string]string{}
		switch frame.Op {
		case "CALL":
			stackRes = tracely.ParseCall(frame)
		case "DELEGATECALL", "STATICCALL":
			stackRes = tracely.ParseDelegateCall(frame)
		case "REVERT":
			stackRes = tracely.ParseRevert(frame)

		}

		addr := common.HexToAddress(stackRes["addr"])

		if contract, ok := b.verifiedContracts[addr]; ok {
			res[addr] = contract.ContractName()
		}

		for _, addie := range frame.Stack {
			stackAddie, ok := addie.(string)
			if !ok {
				continue
			}

			address := common.HexToAddress(stackAddie)

			if contract, ok := b.verifiedContracts[address]; ok {
				res[address] = contract.ContractName()
			}
		}

	}

	return res
}

// addCastLabels adds the labels for all contracts used in this transaction to the cast.
func (b *Backend) addCastLabels(tx *types.Transaction, from common.Address) (res string) {
	labels := b.getTransactionLabelMap(tx, from)
	for address, name := range labels {
		res += fmt.Sprintf("--label %s:%s ", address.String(), name)
	}

	// remove trailing comma
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return res
}
