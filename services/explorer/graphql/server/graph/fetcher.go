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
	if r.Refs.BridgeRefs[chainID] == nil {
		return nil, fmt.Errorf("bridge contract not set for chain %d", chainID)
	}
	contractAddress := r.Refs.BridgeRefs[chainID].Address().String()

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
			var tokenData *types.SwapReplacementData
			for _, log := range receipt.Logs {
				if log.Topics[0].String() == r.Config.SwapTopicHash {
					tokenData, err = r.parseSwapLog(ctx, *log, chainID)
					if err != nil {
						logger.Errorf("Could not parse swap log on chain %d Error: %v", chainID, err)
					}
				}
				if log.Address.String() == contractAddress {
					logs = append(logs, *log)
				}
			}
			return r.parseAndStoreLog(txFetchContext, chainID, logs, tokenData)
		}
	}
}

func (r Resolver) bwDestinationFallback(ctx context.Context, chainID uint32, address string, kappa string, timestamp int, historical bool) (*model.BridgeWatcherTx, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}
	timeout := time.Duration(0)
	//var backendClient backend.ScribeBackend
	backendClient := r.Clients[chainID]
	if r.Refs.BridgeRefs[chainID] == nil {
		return nil, fmt.Errorf("bridge contract not set for chain %d", chainID)
	}
	contractAddress := r.Refs.BridgeRefs[chainID].Address()

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
			} else {
				startBlock, endBlock, err = r.getIteratorForDestinationLogs(ctx, chainID, backendClient)
				ascending = false
			}
			if err != nil {
				b.Duration()
				logger.Errorf("Could not get iterator for historical logs on chain %d Error: %v", chainID, err)
				continue
			}
			toAddressTopic := common.HexToHash(address)
			toKappaTopic := common.HexToHash(fmt.Sprintf("0x%s", kappa))
			indexerConfig := &scribeTypes.IndexerConfig{
				Addresses:            []common.Address{contractAddress},
				GetLogsRange:         r.Config.Chains[chainID].GetLogsRange,
				GetLogsBatchAmount:   r.Config.Chains[chainID].GetLogsBatchAmount,
				StoreConcurrency:     1,
				ChainID:              chainID,
				StartHeight:          *startBlock,
				EndHeight:            *endBlock,
				ConcurrencyThreshold: 0,
				Topics:               [][]common.Hash{nil, {toAddressTopic}, {toKappaTopic}},
			}

			logFetcher := indexer.NewLogFetcher(backendClient, big.NewInt(int64(*startBlock)), big.NewInt(int64(*endBlock)), indexerConfig, ascending)
			maturedBridgeEvent, err := r.getAndParseLogs(ctx, logFetcher, chainID, kappa)
			if err != nil {
				return nil, fmt.Errorf("could not get and parse logs: %v", err)
			}
			go func() {
				r.DB.StoreEvent(ctx, maturedBridgeEvent)
			}()
			bridgeEvent := maturedBridgeEvent.(*sql.BridgeEvent)
			return bwBridgeToBWTx(bridgeEvent, model.BridgeTxTypeDestination)

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

func (r Resolver) parseAndStoreLog(ctx context.Context, chainID uint32, logs []ethTypes.Log, tokenData *types.SwapReplacementData) (*model.BridgeWatcherTx, error) {
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

	bridgeEvent := parsedLog.(*sql.BridgeEvent)
	if tokenData != nil {
		bridgeEvent.Amount = tokenData.Amount
		bridgeEvent.Token = tokenData.Address.String()
	}
	return bwBridgeToBWTx(bridgeEvent, model.BridgeTxTypeOrigin)
}

func (r Resolver) getAndParseLogs(ctx context.Context, logFetcher *indexer.LogFetcher, chainID uint32, kappa string) (interface{}, error) {
	streamLogsCtx, cancelStreamLogs := context.WithCancel(ctx)
	defer cancelStreamLogs()

	logsChan := *logFetcher.GetFetchedLogsChan()
	destinationData := make(chan *types.IFaceBridgeEvent, 1)
	errorChan := make(chan error)

	// Start fetcher
	go func() {
		err := logFetcher.Start(streamLogsCtx)
		if err != nil {
			errorChan <- err
		}
		close(errorChan) // Close error channel after using to signal other routines.
	}()

	// Consume all the logs and check if there is one that is the same as the kappa
	go func() {
		defer close(destinationData) // Always close channel to signal receiver.

		for {
			select {
			case <-streamLogsCtx.Done():
				return

			case log, ok := <-logsChan:
				if !ok {
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
					destinationData <- ifaceBridgeEvent
				}

			case streamErr, ok := <-errorChan:
				if ok {
					logger.Errorf("error while streaming logs: %v", streamErr)
					cancelStreamLogs()
					close(errorChan)
				}
				return
			}
		}
	}()

	ifaceBridgeEvent, ok := <-destinationData
	if !ok {
		// Handle the case where destinationData was closed without sending data.
		return nil, fmt.Errorf("no log found with kappa %s", kappa)
	}
	var maturedBridgeEvent interface{}
	var err error

	maturedBridgeEvent, err = r.Parsers.BridgeParsers[chainID].MatureLogs(ctx, ifaceBridgeEvent.BridgeEvent, ifaceBridgeEvent.IFace, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not mature logs: %w", err)
	}
	if len(errorChan) > 0 {
		return nil, <-errorChan
	}
	return maturedBridgeEvent, nil

}

// parseSwapLog this is a swap event, we need to get the address from it
func (r Resolver) parseSwapLog(ctx context.Context, swapLog ethTypes.Log, chainID uint32) (*types.SwapReplacementData, error) {
	// parse swap with swap filter
	var swapReplacementData types.SwapReplacementData
	for _, filter := range r.SwapFilters[chainID] {
		swapEvent, err := filter.ParseTokenSwap(swapLog)
		if err != nil {
			continue
		}
		if swapEvent != nil {
			iFace, err := filter.ParseTokenSwap(swapLog)
			if err != nil {
				return nil, fmt.Errorf("could not parse swap event: %v", err)
			}
			soldId := iFace.SoldId
			address, err := r.DB.GetString(ctx, fmt.Sprintf("SELECT token_address FROM token_indices WHERE contract_address='%s' AND chain_id=%d AND token_index=%d", swapLog.Address.String(), chainID, soldId.Uint64()))
			if err != nil {
				return nil, fmt.Errorf("could not parse swap event: %v", err)
			}
			swapReplacementData = types.SwapReplacementData{
				Amount:  iFace.TokensSold,
				Address: common.HexToAddress(address),
			}
			break
		}
	}
	return &swapReplacementData, nil
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
