package testutil

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/testutil/testcontract"
	"golang.org/x/sync/errgroup"
)

// TestChainHandler is a handler for interacting with test contracts on a chain to aid in building extensive tests.
// It is returned when emitting events with the test contracts in the PopulateWithLogs function.
type TestChainHandler struct {
	Addresses           []common.Address
	ContractStartBlocks map[common.Address]uint64
	ContractRefs        map[common.Address]*testcontract.TestContractRef
	EventsEmitted       map[common.Address]uint64
}

type chainBackendPair struct {
	chainID uint32
	backend backend.ScribeBackend
}

type chainContractPair struct {
	chainID      uint32
	chainHandler *TestChainHandler
}

// PopulateChainsWithLogs creates scribe backends for each chain backend and emits events from various contracts on each chain.
func PopulateChainsWithLogs(ctx context.Context, t *testing.T, chainBackends map[uint32]geth.Backend, desiredBlockHeight uint64, managers []*DeployManager, handler metrics.Handler) (map[uint32]*TestChainHandler, map[uint32][]backend.ScribeBackend, error) {
	t.Helper()
	addressChan := make(chan chainContractPair, len(chainBackends))
	scribeBackendChan := make(chan chainBackendPair, len(chainBackends))
	g, groupCtx := errgroup.WithContext(ctx)
	for k, v := range chainBackends {
		chain := k
		chainBackend := v

		g.Go(func() error {
			contractHandler, err := PopulateWithLogs(groupCtx, t, &chainBackend, desiredBlockHeight, managers)

			if err != nil {
				return err
			}

			addressChan <- chainContractPair{chain, contractHandler}

			return nil
		})
		g.Go(func() error {
			host := StartOmnirpcServer(groupCtx, t, &chainBackend)
			scribeBackend, err := backend.DialBackend(ctx, host, handler)

			if err != nil {
				return err
			}

			scribeBackendChan <- chainBackendPair{chain, scribeBackend}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, nil, fmt.Errorf("error populating chains with logs: %w", err)
	}
	close(addressChan) // Close the channels after writing to them
	close(scribeBackendChan)
	// Unpack channels
	chainMap := make(map[uint32]*TestChainHandler)
	scribeBackendMap := make(map[uint32][]backend.ScribeBackend)
	for pair := range addressChan {
		chainMap[pair.chainID] = pair.chainHandler
	}

	for pair := range scribeBackendChan {
		scribeBackendMap[pair.chainID] = []backend.ScribeBackend{pair.backend}
	}

	return chainMap, scribeBackendMap, nil
}

// PopulateWithLogs populates a backend with logs until it reaches a desired block height.
//
// nolint:cyclop
func PopulateWithLogs(ctx context.Context, t *testing.T, backend backends.SimulatedTestBackend, desiredBlockHeight uint64, managers []*DeployManager) (*TestChainHandler, error) {
	t.Helper()

	startBlocks := map[common.Address]uint64{}
	contracts := map[common.Address]contracts.DeployedContract{}
	contractRefs := map[common.Address]*testcontract.TestContractRef{}
	eventsEmitted := map[common.Address]uint64{}
	// Get all the test contracts
	for j := range managers {
		manager := managers[j]
		testContract, testRef := manager.GetTestContract(ctx, backend)
		contracts[testContract.Address()] = testContract
		contractRefs[testContract.Address()] = testRef
		eventsEmitted[testContract.Address()] = 0
	}

	// Get start blocks for the deployed contracts
	for address := range contracts {
		deployTxHash := contracts[address].DeployTx().Hash()
		receipt, err := backend.TransactionReceipt(ctx, deployTxHash)
		if err != nil {
			return nil, fmt.Errorf("error getting receipt for tx: %w", err)
		}
		startBlocks[address] = receipt.BlockNumber.Uint64()
	}
	// Iterate and emit events until we reach the desired block height

	testChainHandler := &TestChainHandler{
		Addresses:           dumpAddresses(contracts),
		ContractStartBlocks: startBlocks,
		ContractRefs:        contractRefs,
		EventsEmitted:       eventsEmitted,
	}
	err := EmitEvents(ctx, t, backend, desiredBlockHeight, testChainHandler)
	if err != nil {
		return nil, fmt.Errorf("error emitting events: %w", err)
	}
	return testChainHandler, nil
}

// EmitEvents emits events from the test contracts until the desired block height is reached.
func EmitEvents(ctx context.Context, t *testing.T, backend backends.SimulatedTestBackend, desiredBlockHeight uint64, testChainHandler *TestChainHandler) error {
	t.Helper()
	i := 0
	for {
		select {
		case <-ctx.Done():
			t.Log(ctx.Err())
			return nil
		default:
			i++
			randomAddress := common.BigToAddress(big.NewInt(int64(i)))
			backend.FundAccount(ctx, randomAddress, *big.NewInt(params.Wei))
			latestBlock, err := backend.BlockNumber(ctx)
			if err != nil {
				return err
			}

			if latestBlock >= desiredBlockHeight {
				return nil
			}
			// Emit EventA for each contract
			g, groupCtx := errgroup.WithContext(ctx)
			transactOpts := backend.GetTxContext(groupCtx, nil)
			for k, v := range testChainHandler.ContractRefs {
				address := k
				ref := v
				// Pass if the contract's specified start block is greater than the current block height.
				// Used for testing livefill passing.
				if latestBlock <= testChainHandler.ContractStartBlocks[address] {
					continue
				}
				// Update number of events emitted
				testChainHandler.EventsEmitted[address]++

				g.Go(func() error {
					tx, err := ref.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
					if err != nil {
						return fmt.Errorf("error emitting event a for contract %s: %w", address.String(), err)
					}
					backend.WaitForConfirmation(groupCtx, tx)

					return nil
				})
			}
			err = g.Wait()
			if err != nil {
				return fmt.Errorf("error emitting events: %w", err)
			}
		}
	}
}

// GetTxBlockNumber gets the block number of a transaction.
func GetTxBlockNumber(ctx context.Context, chain backends.SimulatedTestBackend, tx *types.Transaction) (uint64, error) {
	receipt, err := chain.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return 0, fmt.Errorf("error getting receipt for tx: %w", err)
	}
	return receipt.BlockNumber.Uint64(), nil
}

// StartOmnirpcServer starts an omnirpc server and returns the url to it.
func StartOmnirpcServer(ctx context.Context, t *testing.T, backend backends.SimulatedTestBackend) string {
	t.Helper()
	baseHost := testhelper.NewOmnirpcServer(ctx, t, backend)
	return testhelper.GetURL(baseHost, backend)
}

// ReachBlockHeight reaches a block height on a backend.
func ReachBlockHeight(ctx context.Context, t *testing.T, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) error {
	t.Helper()
	i := 0
	for {
		select {
		case <-ctx.Done():
			t.Log(ctx.Err())
			return nil
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))

		latestBlock, err := backend.BlockNumber(ctx)
		if err != nil {
			return fmt.Errorf("error getting latest block number: %w", err)
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

// GetLogsUntilNoneLeft gets all receipts from the database until there are none left (iterates page num).
func GetLogsUntilNoneLeft(ctx context.Context, testDB db.EventDB, filter db.LogFilter) ([]*types.Log, error) {
	var logs []*types.Log
	page := 0
	for {
		page++
		newLogs, err := testDB.RetrieveLogsWithFilter(ctx, filter, page)
		if err != nil {
			return nil, fmt.Errorf("error getting logs: %w", err)
		}
		if len(newLogs) == 0 {
			return logs, nil
		}
		logs = append(logs, newLogs...)
	}
}

// GetReceiptsUntilNoneLeft gets all receipts from the database until there are none left (iterates page num).
func GetReceiptsUntilNoneLeft(ctx context.Context, testDB db.EventDB, filter db.ReceiptFilter) ([]types.Receipt, error) {
	var receipts []types.Receipt
	page := 0
	for {
		page++
		newReceipts, err := testDB.RetrieveReceiptsWithFilter(ctx, filter, page)
		if err != nil {
			return nil, fmt.Errorf("error getting receipts: %w", err)
		}
		if len(newReceipts) == 0 {
			return receipts, nil
		}
		receipts = append(receipts, newReceipts...)
	}
}

// MakeRandomLog makes a random log.
func MakeRandomLog(txHash common.Hash) types.Log {
	return types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		Topics:      []common.Hash{common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64()))},
		Data:        []byte(gofakeit.Sentence(10)),
		BlockNumber: gofakeit.Uint64(),
		TxHash:      txHash,
		TxIndex:     uint(gofakeit.Uint64()),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Index:       uint(gofakeit.Uint64()),
		Removed:     gofakeit.Bool(),
	}
}
