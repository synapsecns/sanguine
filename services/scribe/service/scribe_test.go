package service_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/service"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"

	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	"github.com/synapsecns/sanguine/services/scribe/logger"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Spins up three chains with three contracts on each. Each contract emits events across a span of 20 blocks.
// The generated chains and contracts are fed into a new scribe instance, which is then queried for logs.
func (s *ScribeSuite) TestSimulatedScribe() {
	if os.Getenv("CI") != "" {
		s.T().Skip("Test flake: 20 sec of livefilling may fail on CI")
	}
	const numberOfContracts = 3
	const desiredBlockHeight = 20
	chainIDs := []uint32{gofakeit.Uint32(), gofakeit.Uint32(), gofakeit.Uint32()}
	chainBackends := make(map[uint32]geth.Backend)
	for i := range chainIDs {
		newBackend := geth.NewEmbeddedBackendForChainID(s.GetTestContext(), s.T(), big.NewInt(int64(chainIDs[i])))
		chainBackends[chainIDs[i]] = *newBackend
	}

	managers := []*testutil.DeployManager{s.manager}
	if numberOfContracts > 1 {
		for i := 1; i < numberOfContracts; i++ {
			managers = append(managers, testutil.NewDeployManager(s.T()))
		}
	}

	testChainHandlerMap, chainBackendMap, err := testutil.PopulateChainsWithLogs(s.GetTestContext(), s.T(), chainBackends, desiredBlockHeight, managers, s.nullMetrics)
	Nil(s.T(), err)

	// Build scribe config
	var chainConfigs []config.ChainConfig
	for chainID, testChainHandler := range testChainHandlerMap {
		contractConfigs := config.ContractConfigs{}
		for i := range testChainHandler.Addresses {
			contractConfig := config.ContractConfig{
				Address: testChainHandler.Addresses[i].String(),
			}
			contractConfigs = append(contractConfigs, contractConfig)
		}

		chainConfig := config.ChainConfig{
			ChainID:            chainID,
			Confirmations:      0,
			GetLogsBatchAmount: 1,
			StoreConcurrency:   1,
			GetLogsRange:       1,
			Contracts:          contractConfigs,
		}
		chainConfigs = append(chainConfigs, chainConfig)
	}

	scribeConfig := config.Config{
		Chains: chainConfigs,
	}

	scribe, err := service.NewScribe(s.testDB, chainBackendMap, scribeConfig, s.nullMetrics)
	Nil(s.T(), err)
	killableContext, cancel := context.WithTimeout(s.GetTestContext(), 20*time.Second)
	defer cancel()
	_ = scribe.Start(killableContext)

	// Check that the events were recorded.
	for _, chainConfig := range scribeConfig.Chains {
		for _, contractConfig := range chainConfig.Contracts {
			// Check the storage of logs.
			logFilter := db.LogFilter{
				ChainID:         chainConfig.ChainID,
				ContractAddress: contractConfig.Address,
			}
			logs, err := s.testDB.RetrieveLogsWithFilter(s.GetTestContext(), logFilter, 1)
			Nil(s.T(), err)
			Equal(s.T(), 4, len(logs))
			lastIndexed, err := s.testDB.RetrieveLastIndexed(s.GetTestContext(), common.HexToAddress(contractConfig.Address), chainConfig.ChainID, scribeTypes.IndexingConfirmed)
			Nil(s.T(), err)
			LessOrEqual(s.T(), desiredBlockHeight, int(lastIndexed))
		}
		// Check the storage of receipts.
		receiptFilter := db.ReceiptFilter{
			ChainID: chainConfig.ChainID,
		}
		receipts, err := s.testDB.RetrieveReceiptsWithFilter(s.GetTestContext(), receiptFilter, 1)
		Nil(s.T(), err)
		Equal(s.T(), 12, len(receipts))
	}
}

// TestLivefillParity runs livefill on certain prod chains. Then it checks parity with that chain's block explorer API.
//
// nolint:gocognit,cyclop,maintidx
func (s *ScribeSuite) TestLivefillParity() {
	if os.Getenv("CI") != "" {
		s.T().Skip("Network test flake")
	}
	const blockRange = uint64(100)
	const globalConfirmations = uint64(200)
	// ethRPCURL := "https://1rpc.io/eth"
	// arbRPCURL := "https://endpoints.omniatech.io/v1/arbitrum/one/public"
	// avaxRPCURL := "https://avalanche.public-rpc.com"

	ethRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/1"
	arbRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/42161"
	maticRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/137"
	avaxRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/43114"
	bscRPCURL := "https://rpc.interoperability.institute/confirmations/1/rpc/56"

	ethClient, err := backend.DialBackend(s.GetTestContext(), ethRPCURL, s.nullMetrics)
	Nil(s.T(), err)
	arbClient, err := backend.DialBackend(s.GetTestContext(), arbRPCURL, s.nullMetrics)
	Nil(s.T(), err)
	maticClient, err := backend.DialBackend(s.GetTestContext(), maticRPCURL, s.nullMetrics)
	Nil(s.T(), err)
	avaxClient, err := backend.DialBackend(s.GetTestContext(), avaxRPCURL, s.nullMetrics)
	Nil(s.T(), err)
	bscClient, err := backend.DialBackend(s.GetTestContext(), bscRPCURL, s.nullMetrics)
	Nil(s.T(), err)

	ethID := uint32(1)
	bscID := uint32(56)
	arbID := uint32(42161)
	maticID := uint32(137)
	avaxID := uint32(43114)
	chains := []uint32{ethID, bscID, arbID, maticID, avaxID}

	// Get the current block for each chain.
	ethCurrentBlock, err := ethClient.BlockNumber(s.GetTestContext())
	Nil(s.T(), err)
	ethCurrentBlock -= globalConfirmations
	arbCurrentBlock, err := arbClient.BlockNumber(s.GetTestContext())
	Nil(s.T(), err)
	arbCurrentBlock -= globalConfirmations
	maticCurrentBlock, err := maticClient.BlockNumber(s.GetTestContext())
	Nil(s.T(), err)
	maticCurrentBlock -= globalConfirmations
	avaxCurrentBlock, err := avaxClient.BlockNumber(s.GetTestContext())
	Nil(s.T(), err)
	avaxCurrentBlock -= globalConfirmations
	bscCurrentBlock, err := bscClient.BlockNumber(s.GetTestContext())
	Nil(s.T(), err)
	bscCurrentBlock -= globalConfirmations

	latestBlocks := map[uint32]uint64{
		ethID:   ethCurrentBlock,
		arbID:   arbCurrentBlock,
		maticID: maticCurrentBlock,
		avaxID:  avaxCurrentBlock,
		bscID:   bscCurrentBlock,
	}
	clients := map[uint32][]backend.ScribeBackend{
		ethID:   {ethClient, ethClient},
		bscID:   {bscClient, bscClient},
		arbID:   {arbClient, arbClient},
		maticID: {maticClient, maticClient},
		avaxID:  {avaxClient, avaxClient},
	}

	apiURLs := map[uint32]string{
		ethID:   "https://api.etherscan.io/api",
		arbID:   "https://api.arbiscan.io/api",
		avaxID:  "https://api.snowtrace.io/api",
		bscID:   "https://api.bscscan.com/api",
		maticID: "https://api.polygonscan.com/api",
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{
			{
				ChainID:              ethID,
				Confirmations:        0,
				GetLogsRange:         50,
				GetLogsBatchAmount:   3,
				GetBlockBatchAmount:  10,
				ConcurrencyThreshold: 20000,
				LivefillThreshold:    100,
				Contracts: []config.ContractConfig{
					{
						Address:    "0x2796317b0fF8538F253012862c06787Adfb8cEb6",
						StartBlock: ethCurrentBlock - blockRange,
					},
					{
						Address:    "0x1116898DdA4015eD8dDefb84b6e8Bc24528Af2d8",
						StartBlock: ethCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID:              bscID,
				Confirmations:        0,
				GetLogsRange:         50,
				GetLogsBatchAmount:   3,
				GetBlockBatchAmount:  10,
				ConcurrencyThreshold: 20000,
				LivefillThreshold:    100,
				Contracts: []config.ContractConfig{
					{
						Address:    "0x28ec0B36F0819ecB5005cAB836F4ED5a2eCa4D13",
						StartBlock: bscCurrentBlock - blockRange,
					},
					{
						Address:    "0x930d001b7efb225613aC7F35911c52Ac9E111Fa9",
						StartBlock: bscCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID:              arbID,
				Confirmations:        0,
				GetLogsRange:         50,
				GetLogsBatchAmount:   3,
				GetBlockBatchAmount:  10,
				ConcurrencyThreshold: 20000,
				LivefillThreshold:    100,
				Contracts: []config.ContractConfig{
					{
						Address:    "0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9",
						StartBlock: arbCurrentBlock - blockRange,
					},
					{
						Address:    "0x9Dd329F5411466d9e0C488fF72519CA9fEf0cb40",
						StartBlock: arbCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID:              maticID,
				Confirmations:        0,
				GetLogsRange:         50,
				GetLogsBatchAmount:   3,
				GetBlockBatchAmount:  10,
				ConcurrencyThreshold: 20000,
				LivefillThreshold:    100,
				Contracts: []config.ContractConfig{
					{
						Address:    "0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280",
						StartBlock: maticCurrentBlock - blockRange,
					},
					{
						Address:    "0x85fCD7Dd0a1e1A9FCD5FD886ED522dE8221C3EE5",
						StartBlock: maticCurrentBlock - blockRange,
					},
				},
			},
			{
				ChainID:              avaxID,
				Confirmations:        0,
				GetLogsRange:         50,
				GetLogsBatchAmount:   3,
				GetBlockBatchAmount:  10,
				ConcurrencyThreshold: 20000,
				LivefillThreshold:    100,
				Contracts: []config.ContractConfig{
					{
						Address:    "0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE",
						StartBlock: avaxCurrentBlock - blockRange,
					},
					{
						Address:    "0x77a7e60555bC18B4Be44C181b2575eee46212d44",
						StartBlock: avaxCurrentBlock - blockRange,
					},
				},
			},
		},
	}

	scribe, err := service.NewScribe(s.testDB, clients, scribeConfig, s.nullMetrics)
	Nil(s.T(), err)

	killableContext, cancel := context.WithCancel(s.GetTestContext())

	go func() {
		_ = scribe.Start(killableContext)
	}()

	doneChan := make(chan bool, len(chains))

	for i := range chains {
		go func(index int) {
			for {
				allContractsBackfilled := true
				chain := scribeConfig.Chains[index]
				for _, contract := range chain.Contracts {
					currentBlock, err := s.testDB.RetrieveLastIndexed(s.GetTestContext(), common.HexToAddress(contract.Address), chain.ChainID, scribeTypes.IndexingConfirmed)

					Nil(s.T(), err)
					if currentBlock <= latestBlocks[chain.ChainID] {
						allContractsBackfilled = false
					}
				}
				if allContractsBackfilled {
					doneChan <- true
					fmt.Println("Done with chain", chain.ChainID, "index", index, "of", len(chains), "chains")

					return
				}
				time.Sleep(time.Second)
			}
		}(i)
	}

	for range chains {
		<-doneChan
	}
	cancel()
	for i := range chains {
		chain := scribeConfig.Chains[i]
		for _, contract := range chain.Contracts {
			logFilter := db.LogFilter{
				ChainID:         chains[i],
				ContractAddress: contract.Address,
			}
			fromBlock := latestBlocks[chains[i]] - blockRange
			toBlock := latestBlocks[chains[i]]
			var dbLogCount int
			var dbLogs []*types.Log
			dbLogCount, dbLogs, err = getLogAmount(s.GetTestContext(), s.testDB, logFilter, fromBlock, toBlock)
			Nil(s.T(), err)

			txs := make(map[int64]string)
			var explorerLogCount int
			explorerLogCount, err = getLogs(s.GetTestContext(), contract.Address, fromBlock, toBlock, apiURLs[chain.ChainID], &txs)
			Nil(s.T(), err)

			for k := range dbLogs {
				logBlockNumber := int64(dbLogs[k].BlockNumber)

				txLog := txs[logBlockNumber]
				if dbLogs[k].TxHash.String() != txLog {
					Error(s.T(), fmt.Errorf("mismatched TX\nchainid %d\nstart %d end %d\ndb txhash %s\nexplorer tx %s", chain.ChainID, contract.StartBlock, dbLogs[k].BlockNumber, dbLogs[k].TxHash.String(), txLog))
				}
			}
			// fmt.Println("chain", chain.ChainID, "contract", contract.Address, "dbLogCount", dbLogCount, "explorerLogCount", explorerLogCount)
			if dbLogCount != explorerLogCount {
				fmt.Println("chain", chain.ChainID, "contract", contract.Address, "dbLogCount", dbLogCount, "explorerLogCount", explorerLogCount)
			}
			Equal(s.T(), dbLogCount, explorerLogCount)
		}
	}
}

func createHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
}

func processBatch(ctx context.Context, client *http.Client, url string, txs *map[int64]string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("error getting data: %w", err)
	}
	resRaw, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("could not get data from explorer %w", err)
	}

	var decodedRes map[string]json.RawMessage
	if err := json.NewDecoder(resRaw.Body).Decode(&decodedRes); err != nil {
		return 0, fmt.Errorf("error decoding response: %w", err)
	}

	var resultSlice []map[string]interface{}
	if err := json.Unmarshal(decodedRes["result"], &resultSlice); err != nil {
		return 0, fmt.Errorf("error unmarshaling result: %w", err)
	}

	if err = resRaw.Body.Close(); err != nil {
		logger.ReportScribeError(err, 0, logger.TestError)
	}

	for _, result := range resultSlice {
		hexBlock, ok := result["blockNumber"].(string)
		if !ok {
			return 0, fmt.Errorf("error parsing block number: %w", err)
		}

		txHashStr, ok := result["transactionHash"].(string)
		if !ok {
			return 0, fmt.Errorf("error parsing transaction hash: %w", err)
		}

		key, err := strconv.ParseInt(strings.TrimPrefix(hexBlock, "0x"), 16, 64)
		if err != nil {
			return 0, fmt.Errorf("error parsing block number: %w", err)
		}
		(*txs)[key] = txHashStr
	}
	return len(resultSlice), nil
}

func getLogs(ctx context.Context, contractAddress string, fromBlock uint64, toBlock uint64, apiURL string, txs *map[int64]string) (int, error) {
	blockRange := toBlock - fromBlock
	batchSize := uint64(400)
	numBatches := blockRange/batchSize + 1
	client := createHTTPClient()
	totalResults := 0

	for i := uint64(0); i < numBatches; i++ {
		startBlock := fromBlock + i*batchSize
		endBlock := startBlock + batchSize - 1
		if endBlock > toBlock {
			endBlock = toBlock
		}
		url := fmt.Sprintf("%s?module=logs&action=getLogs&address=%s&fromBlock=%d&toBlock=%d&page=1",
			apiURL, contractAddress, startBlock, endBlock)
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    5 * time.Second,
			Max:    10 * time.Second,
		}
		timeout := 3 * time.Second

	RETRY:
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			resultCount, err := processBatch(ctx, client, url, txs)
			if err != nil {
				fmt.Println("error getting explorer logs", err)
				timeout = b.Duration()
				goto RETRY
			}
			totalResults += resultCount
		}

		if i < numBatches-1 {
			time.Sleep(3 * time.Second)
		}
	}

	return totalResults, nil
}

func getLogAmount(ctx context.Context, db db.EventDB, filter db.LogFilter, startBlock uint64, endBlock uint64) (int, []*types.Log, error) {
	page := 1
	var retrievedLogs []*types.Log
	for {
		logs, err := db.RetrieveLogsInRangeAsc(ctx, filter, startBlock, endBlock, page)
		if err != nil {
			return 0, nil, fmt.Errorf("failure while retreiving logs from database %w", err)
		}
		retrievedLogs = append(retrievedLogs, logs...)
		if len(logs) < base.PageSize {
			break
		}
		page++
	}
	return len(retrievedLogs), retrievedLogs, nil
}
