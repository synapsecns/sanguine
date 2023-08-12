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
	"github.com/synapsecns/sanguine/services/explorer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"github.com/synapsecns/sanguine/services/explorer/types"
	"github.com/synapsecns/sanguine/services/scribe/service/indexer"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
	"math/big"
	"time"
)

var logger = log.Logger("explorer-server-fetcher")

const maxTimeToWaitForTx = 15 * time.Second
const batchAmount = 3

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
	backendClient := r.Clients[chainID]
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			receipt, err := backendClient.TransactionReceipt(txFetchContext, common.HexToHash(txHash))
			if err != nil {
				timeout = b.Duration()
				logger.Errorf("Could not get receipt on chain %d Error: %v", chainID, err)
				continue
			}
			var logs []ethTypes.Log
			for _, log := range receipt.Logs {
				logs = append(logs, *log)
			}
			return r.parseAndStoreLog(txFetchContext, chainID, logs)
		}
	}
}

func (r Resolver) bwDestinationFallback(ctx context.Context, chainID uint32, contractAddress common.Address, address string, kappa string, timestamp int, historical bool) (*model.BridgeWatcherTx, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}
	timeout := time.Duration(0)
	//var backendClient backend.ScribeBackend
	backendClient := r.Clients[chainID]

	for {
		select {
		case <-ctx.Done():

			return nil, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			var err error
			var startBlock *uint64
			var endBlock *uint64
			ascending := true
			if historical {
				startBlock, endBlock, err = r.getIteratorForHistoricalDestinationLogs(ctx, chainID, uint64(timestamp), backendClient)
				ascending = false
			} else {
				startBlock, endBlock, err = r.getIteratorForDestinationLogs(ctx, chainID, backendClient)
			}
			if err != nil {
				b.Duration()
				logger.Errorf("Could not get iterator for historical logs on chain %d Error: %v", chainID, err)
				continue
			}
			toAddressTopic := common.HexToHash(address)
			indexerConfig := &scribeTypes.IndexerConfig{
				Addresses:            []common.Address{contractAddress},
				GetLogsRange:         r.Config.Chains[chainID].GetLogsRange,
				GetLogsBatchAmount:   r.Config.Chains[chainID].GetLogsBatchAmount,
				StoreConcurrency:     1,
				ChainID:              chainID,
				StartHeight:          *startBlock,
				EndHeight:            *endBlock,
				ConcurrencyThreshold: 0,
				Topics:               [][]common.Hash{{toAddressTopic}},
			}
			logFetcher := indexer.NewLogFetcher(backendClient, big.NewInt(int64(*startBlock)), big.NewInt(int64(*endBlock)), indexerConfig, ascending)

			if err != nil {
				return nil, nil
			}
			go func() {
				r.DB.StoreEvent(ctx, maturedBridgeEvent)
			}()
			bridgeEvent := maturedBridgeEvent.(sql.BridgeEvent)
			return bwBridgeToBWTx(&bridgeEvent, model.BridgeTxTypeDestination)

		}
	}

}

func (r Resolver) getIteratorForDestinationLogs(ctx context.Context, chainID uint32, backendClient client.EVM) (*uint64, *uint64, error) {
	currentBlock, err := backendClient.BlockNumber(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get current block%s/%d. Error: %v", r.Config.RPCURL, chainID, err)
	}
	zero := uint64(0)
	return &zero, &currentBlock, nil
}

func (r Resolver) getIteratorForHistoricalDestinationLogs(ctx context.Context, chainID uint32, timestamp uint64, backendClient client.EVM) (*uint64, *uint64, error) {
	currentBlock, err := backendClient.BlockNumber(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get current block%s/%d. Error: %v", r.Config.RPCURL, chainID, err)
	}
	currentTime := uint64(time.Now().Unix())
	postulatedBlock := currentBlock - (currentTime-timestamp)*r.Config.Chains[chainID].BlockTime
	blockHeader, err := backendClient.BlockByNumber(ctx, big.NewInt(int64(postulatedBlock)))
	if err != nil {
		return nil, nil, fmt.Errorf("could not get block %d on chain %d. Error: %v", postulatedBlock, chainID, err)
	}
	difference := blockHeader.Time() - timestamp
	if difference > 0 {
		postulatedBlock = postulatedBlock - difference*(r.Config.Chains[chainID].BlockTime+5)
	}
	return &postulatedBlock, &currentBlock, nil
}

func (r Resolver) parseAndStoreLog(ctx context.Context, chainID uint32, logs []ethTypes.Log) (*model.BridgeWatcherTx, error) {
	parsedLogs, err := backfill.ProcessLogs(ctx, logs, chainID, r.Parsers.BridgeParsers[chainID])
	if err != nil {
		return nil, fmt.Errorf("could not parse logs: %w", err)
	}
	go func() {
		r.DB.StoreEvents(ctx, parsedLogs)
	}()

	parsedLog := interface{}(nil)
	for _, log := range parsedLogs {
		if log == nil {
			continue
		}
		parsedLog = log
	}
	if parsedLog == nil {
		return nil, fmt.Errorf("could not parse logs: %w", err)
	}
	bridgeEvent := parsedLog.(sql.BridgeEvent)
	return bwBridgeToBWTx(&bridgeEvent, model.BridgeTxTypeOrigin)
}

func (r Resolver) checkKappaExists(ctx context.Context, kappa string, chainID uint32) bool {
	var kappaBytes [32]byte
	copy(kappaBytes[:], kappa)
	exists, err := r.Refs.BridgeRefs[chainID].KappaExists(&bind.CallOpts{
		Context: ctx,
	}, kappaBytes)
	if err != nil {
		logger.Errorf("Could not check if kappa exists on chain %d. Error: %v", chainID, err)
		return false
	}
	return exists
}
func (r Resolver) getAndParseLogs(ctx context.Context, logFetcher *indexer.LogFetcher, chainID uint32, kappa string) (interface{}, error) {
	streamLogsCtx, cancelStreamLogs := context.WithCancel(ctx)
	defer cancelStreamLogs()

	logsChan := *logFetcher.GetFetchedLogsChan()
	destinationData := make(chan *types.IFaceBridgeEvent)
	errorChan := make(chan error) // Capacity of 3 because we have 3 goroutines that might send errors

	// Start fetcher
	go func() {
		err := logFetcher.Start(streamLogsCtx)
		if err != nil {
			errorChan <- err
		}
	}()

	// Consume all the logs and check if there is one that is the same as the kappa
	go func() {
		for {
			select {
			case <-streamLogsCtx.Done():
				errorChan <- fmt.Errorf("context canceled while storing and retrieving logs: %w", streamLogsCtx.Err())
				return
			case log, ok := <-logsChan: // empty log passed when ok is false.
				if !ok {
					close(destinationData)
					return
				}
				bridgeEvent, iFace, err := r.Parsers.BridgeParsers[chainID].ParseLog(log, chainID)
				if err != nil {
					logger.Errorf("could not parse log: %v", err)
					continue
				}
				if bridgeEvent.Kappa.Valid && bridgeEvent.Kappa.String == kappa {
					ifaceBridgeEvent := &types.IFaceBridgeEvent{
						IFace:       iFace,
						BridgeEvent: bridgeEvent,
					}
					select {
					case destinationData <- ifaceBridgeEvent:
					case <-streamLogsCtx.Done():
						errorChan <- fmt.Errorf("context canceled while sending bridge event: %w", streamLogsCtx.Err())
						return
					}
				}
			}
		}
	}()

	var maturedBridgeEvent interface{}

	<-streamLogsCtx.Done()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		maturedBridgeEvent, err := r.Parsers.BridgeParsers[chainID].MatureLogs(ctx, ifaceBridgeEvent.BridgeEvent, ifaceBridgeEvent.IFace, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not mature logs: %w", err)
		}
		if len(errorChan) > 0 {
			return nil, <-errorChan
		}
		return maturedBridgeEvent, nil
	}

}
