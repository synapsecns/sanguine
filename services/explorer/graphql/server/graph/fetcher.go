package graph

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/explorer/backfill"
	fetcherpkg "github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"github.com/synapsecns/sanguine/services/explorer/static"
	"time"
)

var logger = log.Logger("explorer-server-fetcher")

const maxTimeToWaitForTx = 15 * time.Second
const batchAmount = 3
const chunkSize = 1000
const defaultRange = 10000

const bridgeConfigAddress = "0x5217c83ca75559B1f8a8803824E5b7ac233A12a1" // TODO create a server config and have this there.
func (r Resolver) bwOriginFallback(ctx context.Context, chainID uint32, txHash string) (*model.BridgeWatcherTx, error) {
	txFetchContext, cancelTxFetch := context.WithTimeout(ctx, maxTimeToWaitForTx)
	defer cancelTxFetch()
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}

	timeout := time.Duration(0)
	//var backendClient backend.ScribeBackend
	backendClient, err := client.DialBackend(ctx, fmt.Sprintf("%s/%d", r.OmniRPCURL, chainID), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create backend client: %w", err)
	}

	for {
		select {
		case <-ctx.Done():

			return nil, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):

			reciept, err := backendClient.TransactionReceipt(txFetchContext, common.HexToHash(txHash))
			if err != nil {
				timeout = b.Duration()
				logger.Errorf("Could not get recipet %s/%d. Error: %v", r.OmniRPCURL, chainID, err)
				continue
			}

			var logs []ethTypes.Log
			for _, log := range reciept.Logs {
				logs = append(logs, *log)
			}
			return r.parseAndStoreLog(txFetchContext, backendClient, common.HexToAddress(bridgeConfigAddress), chainID, logs)
		}
	}
}

//
//func (r Resolver) bwDestinationFallback(ctx context.Context, chainID uint32, address string, kappa string, timestamp int) (*model.BridgeWatcherTx, error) {
//	b := &backoff.Backoff{
//		Factor: 2,
//		Jitter: true,
//		Min:    30 * time.Millisecond,
//		Max:    5 * time.Second,
//	}
//
//	timeout := time.Duration(0)
//	//var backendClient backend.ScribeBackend
//	backendClient, err := client.DialBackend(ctx, fmt.Sprintf("%s/%d", r.OmniRPCURL, chainID), nil)
//	if err != nil {
//		return nil, fmt.Errorf("could not create backend client: %w", err)
//	}
//
//	for {
//		select {
//		case <-ctx.Done():
//
//			return nil, fmt.Errorf("context canceled: %w", ctx.Err())
//		case <-time.After(timeout):
//
//			currentBlock, err := backendClient.BlockNumber(ctx)
//			if err != nil {
//				timeout = b.Duration()
//				logger.Errorf("Could not get current block %s/%d. Error: %v", r.OmniRPCURL, chainID, err)
//				continue
//			}
//
//			startBlock := currentBlock - 10000
//			config := scribeType.IndexerConfig{
//				ChainID: chainID,
//				GetLogsBatchAmount: batchAmount,
//				GetLogsRange: chunkSize,
//				Addresses:
//
//
//			}
//			scribeFetcher.NewLogFetcher(backendClient,startBlock, currentBlock )
//
//
//			iterator := util.NewChunkIterator(big.NewInt(int64(currentBlock-10000)), big.NewInt(int64(currentBlock)), chunkSize, false)
//
//			for {
//
//			}
//			getChunkArr(iterator)
//			reciept, err := scribeBackend.GetLogsInRange(ctx, backendClient, address, chainID)
//			if err != nil {
//				timeout = b.Duration()
//				logger.Errorf("Could not get recipet %s/%d. Error: %v", r.OmniRPCURL, chainID, err)
//				continue
//			}
//
//			var logs []ethTypes.Log
//			for _, log := range reciept.Logs {
//				logs = append(logs, *log)
//			}
//			return r.parseAndStoreLog(ctx, backendClient, common.HexToAddress(bridgeConfigAddress), chainID, logs)
//		}
//	}
//}

func (r Resolver) parseAndStoreLog(ctx context.Context, client bind.ContractBackend, address common.Address, chainID uint32, logs []ethTypes.Log) (*model.BridgeWatcherTx, error) {
	bridgeConfigRef, err := bridgeconfig.NewBridgeConfigRef(common.HexToAddress(bridgeConfigAddress), client)
	if err != nil || bridgeConfigRef == nil {
		return nil, fmt.Errorf("could not create bridge config ScribeFetcher: %w", err)
	}
	priceDataService, err := tokenprice.NewPriceDataService()
	if err != nil {
		return nil, fmt.Errorf("could not create price data service: %w", err)
	}
	newConfigFetcher, err := fetcherpkg.NewBridgeConfigFetcher(common.HexToAddress(bridgeConfigAddress), bridgeConfigRef)
	if err != nil || newConfigFetcher == nil {
		return nil, fmt.Errorf("could not get bridge abi: %w", err)
	}
	tokenSymbolToIDs, err := parser.ParseYaml(static.GetTokenSymbolToTokenIDConfig())
	if err != nil {
		return nil, fmt.Errorf("could not open yaml file: %w", err)
	}
	tokenDataService, err := tokendata.NewTokenDataService(newConfigFetcher, tokenSymbolToIDs)
	bridgeParser, err := parser.NewBridgeParser(r.DB, address, tokenDataService, r.Fetcher, priceDataService)

	parsedLogs, err := backfill.ProcessLogs(ctx, logs, chainID, bridgeParser)
	if err != nil {
		return nil, fmt.Errorf("could not parse logs: %w", err)
	}
	go func() {
		r.DB.StoreEvents(ctx, parsedLogs)
	}()

	bridgeEvent := parsedLogs[0].(sql.BridgeEvent)

	return bwOriginBridgeToBWTx(&bridgeEvent, model.BridgeTxTypeOrigin)
}

// func getChunkArr() (chunkArr []*util.Chunk) { gets the appropriate amount of block chunks (getLogs ranges).
func getChunkArr(iterator util.ChunkIterator) (chunkArr []*util.Chunk) {
	for i := uint64(0); i < batchAmount; i++ {
		chunk := iterator.NextChunk()
		if chunk == nil {
			return chunkArr
		}
		chunkArr = append(chunkArr, chunk)
	}
	return chunkArr
}
