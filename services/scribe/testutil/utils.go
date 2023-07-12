package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"math/big"
	"testing"
)

// PopulateWithLogs populates a backend with logs until it reaches a desired block height.
func PopulateWithLogs(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64, testingSuite *testing.T, manager *DeployManager) (*common.Address, error) {
	i := 0
	var address common.Address
	for {
		select {
		case <-ctx.Done():
			testingSuite.Log(ctx.Err())
			return &address, nil
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
		testContract, testRef := manager.GetTestContract(ctx, backend)
		address = testContract.Address()
		transactOpts := backend.GetTxContext(ctx, nil)
		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		if err != nil {
			return nil, fmt.Errorf("error getting latest block number: %v", err)
		}
		backend.WaitForConfirmation(ctx, tx)

		latestBlock, err := backend.BlockNumber(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting latest block number: %v", err)
		}

		if latestBlock >= desiredBlockHeight {
			return &address, nil
		}
	}
}

// MultiContractPopulateWithLogs populates a backend with logs from multiple contracts until it reaches a desired block height.
func MultiContractPopulateWithLogs(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64, testingSuite *testing.T, manager *DeployManager) (*common.Address, *common.Address, error) {
	i := 0
	var address1 common.Address
	var address2 common.Address
	for {
		select {
		case <-ctx.Done():
			testingSuite.Log(ctx.Err())
			return &address1, &address2, nil
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
		testContract1, testRef1 := manager.GetTestContract(ctx, backend)
		testContract2, testRef2 := manager.GetTestContract2(ctx, backend)

		address1 = testContract1.Address()
		address2 = testContract2.Address()

		transactOpts := backend.GetTxContext(ctx, nil)
		tx1, err := testRef1.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		if err != nil {
			return nil, nil, fmt.Errorf("error getting latest block number: %v", err)
		}
		tx2, err := testRef2.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		if err != nil {
			return nil, nil, fmt.Errorf("error getting latest block number: %v", err)
		}

		backend.WaitForConfirmation(ctx, tx1)
		backend.WaitForConfirmation(ctx, tx2)

		latestBlock, err := backend.BlockNumber(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("error getting latest block number: %v", err)
		}

		if latestBlock >= desiredBlockHeight {
			return &address1, &address2, nil
		}
	}
}

// StartOmnirpcServer starts an omnirpc server and returns the url to it.
func StartOmnirpcServer(ctx context.Context, backend backends.SimulatedTestBackend, testingSuite *testing.T) string {
	baseHost := testhelper.NewOmnirpcServer(ctx, testingSuite, backend)
	return testhelper.GetURL(baseHost, backend)
}

// ReachBlockHeight reaches a block height on a backend.
func ReachBlockHeight(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64, testingSuite *testing.T) error {
	i := 0
	for {
		select {
		case <-ctx.Done():
			testingSuite.Log(ctx.Err())
			return nil
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))

		latestBlock, err := backend.BlockNumber(ctx)
		if err != nil {
			fmt.Errorf("error getting latest block number: %v", err)
		}

		if latestBlock >= desiredBlockHeight {
			return nil
		}
	}
}
