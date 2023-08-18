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
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/cctp"
	"github.com/synapsecns/sanguine/services/scribe/service/indexer"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
	"math/big"
	"time"
)

var logger = log.Logger("explorer-server-fetcher")

type ifaceBridgeEvent struct {
	IFace       bridge.EventLog
	BridgeEvent *sql.BridgeEvent
}

type ifaceCCTPEvent struct {
	IFace     cctp.EventLog
	CCTPEvent *sql.CCTPEvent
}

type swapReplacementData struct {
	Address common.Address
	Amount  *big.Int
}

const maxTimeToWaitForTx = 15 * time.Second
const kappaDoesNotExist = "kappa does not exist on destination chain"

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
	// var backendClient backend.ScribeBackend
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
			var tokenData *swapReplacementData
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

func (r Resolver) bwOriginFallbackCCTP(ctx context.Context, chainID uint32, txHash string) (*model.BridgeWatcherTx, error) {
	txFetchContext, cancelTxFetch := context.WithTimeout(ctx, maxTimeToWaitForTx)
	defer cancelTxFetch()
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}
	timeout := time.Duration(0)
	// var backendClient backend.ScribeBackend
	backendClient := r.Clients[chainID]
	if r.Refs.CCTPRefs[chainID] == nil {
		return nil, fmt.Errorf("cctp contract not set for chain %d", chainID)
	}
	contractAddress := r.Refs.CCTPRefs[chainID].Address().String()

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
				if log.Address.String() == contractAddress {
					logs = append(logs, *log)
				}
			}
			return r.parseAndStoreLogCCTP(txFetchContext, chainID, logs)
		}
	}
}

func (r Resolver) bwDestinationFallback(ctx context.Context, chainID uint32, address string, kappa string, timestamp int, historical bool) (*model.BridgeWatcherTx, error) {
	txFetchContext, cancelTxFetch := context.WithTimeout(ctx, maxTimeToWaitForTx)
	defer cancelTxFetch()
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}
	timeout := time.Duration(0)
	// var backendClient backend.ScribeBackend
	backendClient := r.Clients[chainID]
	contractAddress := r.Refs.BridgeRefs[chainID].Address()
	if !r.checkKappaExists(txFetchContext, kappa, chainID) {
		return nil, fmt.Errorf(kappaDoesNotExist)
	}
	for {
		select {
		case <-txFetchContext.Done():

			return nil, fmt.Errorf("context canceled: %w", txFetchContext.Err())
		case <-time.After(timeout):
			var err error
			var startBlock *uint64
			var endBlock *uint64
			ascending := true
			if historical {
				startBlock, endBlock, err = r.getRangeForHistoricalDestinationLogs(txFetchContext, chainID, uint64(timestamp), backendClient)
			} else {
				startBlock, endBlock, err = r.getRangeForDestinationLogs(txFetchContext, chainID, backendClient)
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
			maturedBridgeEvent, err := r.getAndParseLogs(txFetchContext, logFetcher, chainID, kappa)
			if err != nil {
				logger.Errorf("could not get and parse logs: %v", err)
				continue
			}
			go func() {
				storeErr := r.DB.StoreEvent(txFetchContext, maturedBridgeEvent)
				if storeErr != nil {
					logger.Errorf("could not store log while storing origin bridge watcher tx %v", err)
				}
			}()
			bridgeEvent, ok := maturedBridgeEvent.(*sql.BridgeEvent)
			if !ok {
				logger.Errorf("type assertion failed when converting bridge event")
				continue
			}
			return bwBridgeToBWTx(bridgeEvent, model.BridgeTxTypeDestination)
		}
	}
}

func (r Resolver) bwDestinationFallbackCCTP(ctx context.Context, chainID uint32, address string, requestID string, timestamp int, historical bool) (*model.BridgeWatcherTx, error) {
	txFetchContext, cancelTxFetch := context.WithTimeout(ctx, maxTimeToWaitForTx)
	defer cancelTxFetch()
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}
	timeout := time.Duration(0)
	// var backendClient backend.ScribeBackend
	backendClient := r.Clients[chainID]
	contractAddress := r.Refs.CCTPRefs[chainID].Address()
	if !r.checkRequestIDExists(txFetchContext, requestID, chainID) {
		return nil, fmt.Errorf(kappaDoesNotExist)
	}
	for {
		select {
		case <-txFetchContext.Done():

			return nil, fmt.Errorf("context canceled: %w", txFetchContext.Err())
		case <-time.After(timeout):
			var err error
			var startBlock *uint64
			var endBlock *uint64
			ascending := true
			if historical {
				startBlock, endBlock, err = r.getRangeForHistoricalDestinationLogs(txFetchContext, chainID, uint64(timestamp), backendClient)
			} else {
				startBlock, endBlock, err = r.getRangeForDestinationLogs(txFetchContext, chainID, backendClient)
				ascending = false
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
				Topics:               [][]common.Hash{nil, {toAddressTopic}},
			}

			logFetcher := indexer.NewLogFetcher(backendClient, big.NewInt(int64(*startBlock)), big.NewInt(int64(*endBlock)), indexerConfig, ascending)
			maturedBridgeEvent, err := r.getAndParseLogsCCTP(txFetchContext, logFetcher, chainID, requestID)
			if err != nil {
				logger.Errorf("could not get and parse logs: %v", err)
				continue
			}
			go func() {
				storeErr := r.DB.StoreEvent(txFetchContext, maturedBridgeEvent)
				if storeErr != nil {
					logger.Errorf("could not store log while storing origin bridge watcher tx %w", err)
				}
			}()
			bridgeEvent, ok := maturedBridgeEvent.(sql.BridgeEvent)
			if !ok {
				logger.Errorf("type assertion failed when converting bridge event")
				continue
			}
			return bwBridgeToBWTx(&bridgeEvent, model.BridgeTxTypeDestination)
		}
	}
}

func (r Resolver) getRangeForDestinationLogs(ctx context.Context, chainID uint32, backendClient client.EVM) (*uint64, *uint64, error) {
	currentBlock, err := backendClient.BlockNumber(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get current block%s/%d. Error: %w", r.Config.RPCURL, chainID, err)
	}
	zero := uint64(0)
	return &zero, &currentBlock, nil
}

func (r Resolver) getRangeForHistoricalDestinationLogs(ctx context.Context, chainID uint32, timestamp uint64, backendClient client.EVM) (*uint64, *uint64, error) {
	currentBlock, err := backendClient.BlockNumber(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get current block%s/%d. Error: %w", r.Config.RPCURL, chainID, err)
	}
	currentTime := uint64(time.Now().Unix())
	blockTime := r.Config.Chains[chainID].BlockTime
	postulatedBlock := (currentBlock - (currentTime-timestamp)/blockTime) - (r.Config.Chains[chainID].GetLogsRange * r.Config.Chains[chainID].GetLogsBatchAmount)
	blockHeader, err := backendClient.BlockByNumber(ctx, big.NewInt(int64(postulatedBlock)))
	if err != nil {
		return nil, nil, fmt.Errorf("could not get block %d on chain %d. Error: %w", postulatedBlock, chainID, err)
	}

	difference := int64(blockHeader.Time()) - int64(timestamp)
	fmt.Println(currentTime, timestamp, blockHeader.Time(), difference, postulatedBlock, currentBlock, blockTime)

	if difference > 0 {
		postulatedBlock -= uint64(difference)
	}
	fmt.Println(currentTime, timestamp, difference, blockHeader.Time(), postulatedBlock, currentBlock, blockTime)
	return &postulatedBlock, &currentBlock, nil
}

func (r Resolver) parseAndStoreLog(ctx context.Context, chainID uint32, logs []ethTypes.Log, tokenData *swapReplacementData) (*model.BridgeWatcherTx, error) {
	parsedLogs, err := backfill.ProcessLogs(ctx, logs, chainID, r.Parsers.BridgeParsers[chainID])
	if err != nil {
		return nil, fmt.Errorf("could not parse logs: %w", err)
	}
	go func() {
		storeErr := r.DB.StoreEvents(ctx, parsedLogs)
		if storeErr != nil {
			logger.Errorf("could not store log while storing origin bridge watcher tx %v", err)
		}
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

	bridgeEvent, ok := parsedLog.(*sql.BridgeEvent)
	if !ok {
		return nil, fmt.Errorf("type assertion failed when converting bridge event")
	}

	if tokenData != nil {
		bridgeEvent.Amount = tokenData.Amount
		bridgeEvent.Token = tokenData.Address.String()
	}
	return bwBridgeToBWTx(bridgeEvent, model.BridgeTxTypeOrigin)
}

func (r Resolver) parseAndStoreLogCCTP(ctx context.Context, chainID uint32, logs []ethTypes.Log) (*model.BridgeWatcherTx, error) {
	parsedLogs, err := backfill.ProcessLogs(ctx, logs, chainID, r.Parsers.CCTParsers[chainID])
	if err != nil {
		return nil, fmt.Errorf("could not parse logs: %w", err)
	}
	go func() {
		storeErr := r.DB.StoreEvents(ctx, parsedLogs)
		if storeErr != nil {
			logger.Errorf("could not store cctp log while storing origin bridge watcher tx %v", err)
		}
	}()
	parsedLog := interface{}(nil)
	for i, log := range parsedLogs {
		if log == nil {
			continue
		}
		fmt.Println("j", i, log)

		parsedLog = log
	}
	if parsedLog == nil {
		return nil, fmt.Errorf("could not parse logs: %w", err)
	}

	bridgeEvent, ok := parsedLog.(sql.BridgeEvent)
	if !ok {
		return nil, fmt.Errorf("type assertion failed when converting bridge event")
	}
	return bwBridgeToBWTx(&bridgeEvent, model.BridgeTxTypeOrigin)
}

// nolint:cyclop
func (r Resolver) getAndParseLogs(ctx context.Context, logFetcher *indexer.LogFetcher, chainID uint32, kappa string) (interface{}, error) {
	streamLogsCtx, cancelStreamLogs := context.WithCancel(ctx)
	defer cancelStreamLogs()

	logsChan := *logFetcher.GetFetchedLogsChan()
	destinationData := make(chan *ifaceBridgeEvent, 1)
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
					bridgeEventIFace := &ifaceBridgeEvent{
						IFace:       iFace,
						BridgeEvent: bridgeEvent,
					}
					destinationData <- bridgeEventIFace
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

	bridgeEventIFace, ok := <-destinationData
	if !ok {
		// Handle the case where destinationData was closed without sending data.
		return nil, fmt.Errorf("no log found with kappa %s", kappa)
	}
	var maturedBridgeEvent interface{}
	var err error

	maturedBridgeEvent, err = r.Parsers.BridgeParsers[chainID].MatureLogs(ctx, bridgeEventIFace.BridgeEvent, bridgeEventIFace.IFace, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not mature logs: %w", err)
	}
	if len(errorChan) > 0 {
		return nil, <-errorChan
	}
	return maturedBridgeEvent, nil
}

// nolint:cyclop
func (r Resolver) getAndParseLogsCCTP(ctx context.Context, logFetcher *indexer.LogFetcher, chainID uint32, requestID string) (interface{}, error) {
	streamLogsCtx, cancelStreamLogs := context.WithCancel(ctx)
	defer cancelStreamLogs()

	logsChan := *logFetcher.GetFetchedLogsChan()
	destinationData := make(chan *ifaceCCTPEvent, 1)
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
				fmt.Println("from scribe log", log)
				cctpEvent, iFace, err := r.Parsers.CCTParsers[chainID].ParseLog(log, chainID)
				if err != nil {
					logger.Errorf("could not parse log: %v", err)
					continue
				}
				fmt.Println("from scribe log cctpEvent", cctpEvent.RequestID, requestID)

				if cctpEvent.RequestID == requestID {
					ifaceCctpEvent := &ifaceCCTPEvent{
						IFace:     iFace,
						CCTPEvent: cctpEvent,
					}
					destinationData <- ifaceCctpEvent
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

	ifaceCctpEvent, ok := <-destinationData
	if !ok {
		// Handle the case where destinationData was closed without sending data.
		return nil, fmt.Errorf("no log found with kappa %s", requestID)
	}
	var maturedBridgeEvent interface{}
	var err error

	maturedBridgeEvent, err = r.Parsers.CCTParsers[chainID].MatureLogs(ctx, ifaceCctpEvent.CCTPEvent, ifaceCctpEvent.IFace, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not mature logs: %w", err)
	}
	if len(errorChan) > 0 {
		return nil, <-errorChan
	}
	return maturedBridgeEvent, nil
}

// parseSwapLog this is a swap event, we need to get the address from it.
func (r Resolver) parseSwapLog(ctx context.Context, swapLog ethTypes.Log, chainID uint32) (*swapReplacementData, error) {
	// parse swap with swap filter
	var swapReplacement swapReplacementData
	for _, filter := range r.SwapFilters[chainID] {
		swapEvent, err := filter.ParseTokenSwap(swapLog)
		if err != nil {
			continue
		}
		if swapEvent != nil {
			iFace, err := filter.ParseTokenSwap(swapLog)
			if err != nil {
				return nil, fmt.Errorf("could not parse swap event: %w", err)
			}
			soldID := iFace.SoldId
			address, err := r.DB.GetString(ctx, fmt.Sprintf("SELECT token_address FROM token_indices WHERE contract_address='%s' AND chain_id=%d AND token_index=%d", swapLog.Address.String(), chainID, soldID.Uint64()))
			if err != nil {
				return nil, fmt.Errorf("could not parse swap event: %w", err)
			}
			swapReplacement = swapReplacementData{
				Amount:  iFace.TokensSold,
				Address: common.HexToAddress(address),
			}
			break
		}
	}
	return &swapReplacement, nil
}

func (r Resolver) checkKappaExists(ctx context.Context, kappa string, chainID uint32) bool {
	var kappaBytes32 [32]byte

	kappaBytes := common.Hex2Bytes(kappa)
	copy(kappaBytes32[:], kappaBytes)

	exists, err := r.Refs.BridgeRefs[chainID].KappaExists(&bind.CallOpts{
		Context: ctx,
	}, kappaBytes32)
	if err != nil {
		logger.Errorf("Could not check if kappa exists on chain %d. Error: %v", chainID, err)
		return false
	}
	return exists
}

func (r Resolver) checkRequestIDExists(ctx context.Context, requestID string, chainID uint32) bool {
	var kappaBytes32 [32]byte
	kappaBytes := common.Hex2Bytes(requestID)
	copy(kappaBytes32[:], kappaBytes)
	fmt.Println("kappaBytes32", kappaBytes32, "kappaBytes", kappaBytes, "requestID", requestID)
	exists, err := r.Refs.CCTPRefs[chainID].IsRequestFulfilled(&bind.CallOpts{
		Context: ctx,
	}, kappaBytes32)
	if err != nil {
		logger.Errorf("Could not check if request exists on chain %d. Error: %v", chainID, err)
		return false
	}
	return exists
}
