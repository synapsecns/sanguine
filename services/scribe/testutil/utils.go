package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
	"golang.org/x/sync/errgroup"
	"math/big"
	"testing"
)

// PopulateWithLogs populates a backend with logs until it reaches a desired block height.
func PopulateWithLogs(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64, testingSuite *testing.T, managers []*DeployManager) ([]common.Address, map[common.Address]uint64, error) {
	i := 0
	startBlocks := map[common.Address]uint64{}
	contracts := map[common.Address]contracts.DeployedContract{}
	contractRefs := map[common.Address]*testcontract.TestContractRef{}

	// Get all the test contracts
	for j := range managers {
		manager := managers[j]
		testContract, testRef := manager.GetTestContract(ctx, backend)
		contracts[testContract.Address()] = testContract
		contractRefs[testContract.Address()] = testRef
	}

	// Get start blocks for the deployed contracts
	for address := range contracts {
		deployTxHash := contracts[address].DeployTx().Hash()
		receipt, err := backend.TransactionReceipt(ctx, deployTxHash)
		if err != nil {
			return nil, nil, fmt.Errorf("error getting receipt for tx: %w", err)
		}
		startBlocks[address] = receipt.BlockNumber.Uint64()
	}

	// Iterate and emit events until we reach the desired block height
	for {
		select {
		case <-ctx.Done():
			testingSuite.Log(ctx.Err())
			return dumpAddresses(contracts), startBlocks, nil
		default:
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))

		// Emit EventA for each contract
		g, groupCtx := errgroup.WithContext(ctx)
		transactOpts := backend.GetTxContext(groupCtx, nil)
		for k, v := range contractRefs {
			address := k
			ref := v
			g.Go(func() error {
				tx, err := ref.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
				if err != nil {
					return fmt.Errorf("error emitting event a for contract %s: %v", address.String(), err)
				}
				backend.WaitForConfirmation(ctx, tx)
				return nil
			})

		}
		err := g.Wait()
		if err != nil {
			return nil, nil, fmt.Errorf("error emitting events: %v", err)
		}

		latestBlock, err := backend.BlockNumber(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("error getting latest block number: %v", err)
		}

		if latestBlock >= desiredBlockHeight {
			return dumpAddresses(contracts), startBlocks, nil
		}

	}
}

func GetTxBlockNumber(ctx context.Context, chain backends.SimulatedTestBackend, tx *types.Transaction) (uint64, error) {
	receipt, err := chain.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, fmt.Errorf("error getting receipt for tx: %w", err)
	}
	return receipt.BlockNumber.Uint64(), nil
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

// dumpAddresses is a helper function to return all the addresses from a deployed contract.
func dumpAddresses(contracts map[common.Address]contracts.DeployedContract) []common.Address {
	var addresses []common.Address
	for address := range contracts {
		addresses = append(addresses, address)
	}
	return addresses
}
