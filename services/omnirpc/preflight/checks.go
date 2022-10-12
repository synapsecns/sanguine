package preflight

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/omnirpc/types"
	"math/big"
)

// RunPreflightChecks runs the preflight checks based on the rpc method
func RunPreflightChecks(ctx context.Context, rpcRequest types.IRPCRequest, chainID uint) (ok bool, _ error) {
	// TODO: add way to disable the check if config not present, etc
	switch rpcRequest.GetMethod() {
	case types.SendRawTransactionMethod:
		return checkTransaction(rpcRequest, uint64(chainID))
	}
	return true, nil
}

func checkTransaction(rpcRequest types.IRPCRequest, chainID uint64) (ok bool, _ error) {
	validatorAddress := common.HexToAddress("0x0000000000000000000000000000000000000005")

	if len(rpcRequest.GetParams()) == 0 {
		return false, fmt.Errorf("rpc request %s requires at least one arg", types.SendRawTransactionMethod)
	}

	tx := ethTypes.Transaction{}
	err := tx.UnmarshalBinary(rpcRequest.GetParams()[0])
	if err != nil {
		return false, fmt.Errorf("could not unmarshal transaction: %w", err)
	}

	senderAddress, err := ethTypes.LatestSignerForChainID(new(big.Int).SetUint64(chainID)).Sender(&tx)
	if err != nil {
		return false, fmt.Errorf("could not recover sender address: %w", err)
	}

	if senderAddress != validatorAddress {
		return true, nil
	}

	// continue
	return false, nil
}
