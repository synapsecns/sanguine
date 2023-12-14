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

const maxTimeToWaitForTx = 25 * time.Second
const kappaDoesNotExist = "kappa does not exist on destination chain"

// nolint:cyclop
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

			for i := range receipt.Logs {
				// iterating in reverse order to get the latest swap log
				log := receipt.Logs[len(receipt.Logs)-i-1]
				if tokenData == nil && log.Topics[0].String() == r.Config.SwapTopicHash {
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
		case <-txFetchContext.Done():
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

// nolint:gocognit,cyclop
func (r Resolver) bwDestinationFallback(ctx context.Context, chainID uint32, address string, identifier string, timestamp int, historical bool, bridgeType model.BridgeType) (*model.BridgeWatcherTx, error) {
	txFetchContext, cancelTxFetch := context.WithTimeout(ctx, maxTimeToWaitForTx)
	defer cancelTxFetch()

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}

	timeout := time.Duration(0)
	backendClient := r.Clients[chainID]
	var contractAddress common.Address

	// Check if the kappa/request id exists on the destination chain
	switch bridgeType {
	case model.BridgeTypeBridge:
		contractAddress = r.Refs.BridgeRefs[chainID].Address()
		if !r.checkKappaExists(txFetchContext, identifier, chainID) {
			return nil, fmt.Errorf(kappaDoesNotExist)
		}
	case model.BridgeTypeCctp:
		contractAddress = r.Refs.CCTPRefs[chainID].Address()
		if !r.checkRequestIDExists(txFetchContext, identifier, chainID) {
			return nil, fmt.Errorf(kappaDoesNotExist)
		}
	}

	// Start trying to fetch logs
	for {
		select {
		case <-txFetchContext.Done():
			return nil, fmt.Errorf("context canceled: %w", txFetchContext.Err())
		case <-time.After(timeout):
			var err error

			// Get the range of blocks to fetch logs from
			var startBlock, endBlock *uint64
			ascending := true
			if historical {
				startBlock, endBlock, err = r.getRangeForHistoricalDestinationLogs(txFetchContext, chainID, uint64(timestamp), backendClient)
			} else {
				startBlock, endBlock, err = r.getRangeForDestinationLogs(txFetchContext, chainID, backendClient)
				ascending = false
			}
			if err != nil {
				b.Duration()
				logger.Errorf("Could not get iterator for logs on chain %d Error: %v", chainID, err)
				continue
			}
			toAddressTopic := common.HexToHash(address)
			topics := [][]common.Hash{nil, {toAddressTopic}}
			if bridgeType == model.BridgeTypeBridge { // can filter by kappa as well if bridge
				toKappaTopic := common.HexToHash(fmt.Sprintf("0x%s", identifier))
				topics = append(topics, []common.Hash{toKappaTopic})
			}

			indexerConfig := &scribeTypes.IndexerConfig{
				Addresses:            []common.Address{contractAddress},
				GetLogsRange:         r.Config.Chains[chainID].GetLogsRange,
				GetLogsBatchAmount:   r.Config.Chains[chainID].GetLogsBatchAmount,
				StoreConcurrency:     1,
				ChainID:              chainID,
				StartHeight:          *startBlock,
				EndHeight:            *endBlock,
				ConcurrencyThreshold: 0,
				Topics:               topics,
			}

			logFetcher := indexer.NewLogFetcher(backendClient, big.NewInt(int64(*startBlock)), big.NewInt(int64(*endBlock)), indexerConfig, ascending)
			maturedBridgeEvent, err := r.getAndParseLogs(txFetchContext, logFetcher, chainID, identifier, bridgeType)
			if err != nil {
				logger.Errorf("could not get and parse logs: %v", err)
				continue
			}
			go r.storeBridgeEvent(maturedBridgeEvent) // store events
			switch bridgeType {
			case model.BridgeTypeBridge:
				bridgeEvent, ok := maturedBridgeEvent.(*sql.BridgeEvent)
				if !ok {
					logger.Errorf("type assertion failed when converting bridge event")
					continue
				}
				return bwBridgeToBWTx(bridgeEvent, model.BridgeTxTypeDestination)

			case model.BridgeTypeCctp:
				bridgeEvent, ok := maturedBridgeEvent.(sql.BridgeEvent)
				if !ok {
					logger.Errorf("type assertion failed when converting bridge event")
					continue
				}
				return bwBridgeToBWTx(&bridgeEvent, model.BridgeTxTypeDestination)
			}
			return nil, fmt.Errorf("could not convert bridge event to bridge watcher tx")
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
	// Get the current block number.
	currentBlock, err := backendClient.BlockNumber(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get current block%s/%d. Error: %w", r.Config.RPCURL, chainID, err)
	}

	const maxIterations = 25 // max tries
	iteration := 0
	var mid uint64

	upper := currentBlock
	lower := uint64(0)
	for lower <= upper && iteration < maxIterations {
		mid = (lower + upper) / 2
		blockHeader, err := backendClient.HeaderByNumber(ctx, big.NewInt(int64(mid)))
		if err != nil {
			return nil, nil, fmt.Errorf("could not get block %d on chain %d. Error: %w", mid, chainID, err)
		}
		timeDifference := int64(blockHeader.Time) - int64(timestamp)

		switch {
		case -6000 < timeDifference && timeDifference <= 0:
			return &mid, &currentBlock, nil
		case timeDifference >= 0:
			upper = mid
		default:
			lower = mid
		}

		iteration++
	}

	return &mid, &currentBlock, nil
}

func (r Resolver) parseAndStoreLog(ctx context.Context, chainID uint32, logs []ethTypes.Log, tokenData *swapReplacementData) (*model.BridgeWatcherTx, error) {
	parsedLogs, err := backfill.ProcessLogs(ctx, logs, chainID, r.Parsers.BridgeParsers[chainID])
	if err != nil {
		return nil, fmt.Errorf("could not parse logs with explorer parser: %w", err)
	}
	go func() {
		for _, parsedLog := range parsedLogs {
			r.storeBridgeEvent(parsedLog)
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
		return nil, fmt.Errorf("parsed log is nil %w", err)
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
	if err != nil || len(parsedLogs) == 0 {
		return nil, fmt.Errorf("could not parse logs: %w", err)
	}
	go func() {
		for _, parsedLog := range parsedLogs {
			r.storeBridgeEvent(parsedLog)
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

	bridgeEvent, ok := parsedLog.(sql.BridgeEvent)
	if !ok {
		return nil, fmt.Errorf("type assertion failed when converting bridge event")
	}
	return bwBridgeToBWTx(&bridgeEvent, model.BridgeTxTypeOrigin)
}

// nolint:cyclop,gocognit
func (r Resolver) getAndParseLogs(ctx context.Context, logFetcher *indexer.LogFetcher, chainID uint32, kappa string, bridgeType model.BridgeType) (interface{}, error) {
	streamLogsCtx, cancelStreamLogs := context.WithCancel(ctx)
	logsChan := *logFetcher.GetFetchedLogsChan()
	destinationData := make(chan *ifaceBridgeEvent, 1)
	destinationDataCCTP := make(chan *ifaceCCTPEvent, 1)

	closeDestinationChannels := func() {
		close(destinationData)
		close(destinationDataCCTP)
	}
	errorChan := make(chan error)

	// Start log fetching
	go func() {
		if err := logFetcher.Start(streamLogsCtx); err != nil {
			errorChan <- err
		}
	}()

	// Process logs and identify the relevant one matching the provided kappa. Will cancel if there's an error on the fetcher.
	go func() {
		defer cancelStreamLogs() // cancel stream logs if we exit this goroutine
		defer closeDestinationChannels()
		for {
			select {
			case <-ctx.Done():
				return
			case log, ok := <-logsChan:
				if !ok {
					return
				}

				switch bridgeType {
				case model.BridgeTypeBridge:
					bridgeEvent, iFace, err := r.Parsers.BridgeParsers[chainID].ParseLog(log, chainID)
					if err != nil {
						logger.Errorf("could not parse log: %v", err)
						continue
					}
					if bridgeEvent.Kappa.Valid && bridgeEvent.Kappa.String == kappa {
						destinationData <- &ifaceBridgeEvent{
							IFace:       iFace,
							BridgeEvent: bridgeEvent,
						}
						return
					}

				case model.BridgeTypeCctp:
					cctpEvent, iFace, err := r.Parsers.CCTParsers[chainID].ParseLog(log, chainID)
					if err != nil {
						logger.Errorf("could not parse log: %v", err)
						continue
					}
					if cctpEvent.RequestID == kappa {
						destinationDataCCTP <- &ifaceCCTPEvent{
							IFace:     iFace,
							CCTPEvent: cctpEvent,
						}
						return
					}
				}

			case streamErr, ok := <-errorChan:
				if ok {
					logger.Errorf("error while streaming logs: %v", streamErr)
				}
				return
			}
		}
	}()

	<-streamLogsCtx.Done()
	var bridgeEvent interface{}
	var err error
	switch bridgeType {
	case model.BridgeTypeBridge:
		bridgeEventIFace, ok := <-destinationData
		if !ok {
			return nil, fmt.Errorf("no log found with kappa %s", kappa)
		}
		bridgeEvent, err = r.Parsers.BridgeParsers[chainID].MatureLogs(ctx, bridgeEventIFace.BridgeEvent, bridgeEventIFace.IFace, chainID)

	case model.BridgeTypeCctp:
		ifaceCctpEvent, ok := <-destinationDataCCTP
		if !ok {
			// Handle the case where destinationData was closed without sending data.
			return nil, fmt.Errorf("no cctp log found with request id %s", kappa)
		}
		bridgeEvent, err = r.Parsers.CCTParsers[chainID].MatureLogs(ctx, ifaceCctpEvent.CCTPEvent, ifaceCctpEvent.IFace, chainID)
	}
	if err != nil {
		return nil, fmt.Errorf("could not mature logs: %w", err)
	}
	return bridgeEvent, nil
}

// parseSwapLog this is a swap event, we need to get the address from it.
func (r Resolver) parseSwapLog(ctx context.Context, swapLog ethTypes.Log, chainID uint32) (*swapReplacementData, error) {
	// parse swap with swap filter
	var swapReplacement swapReplacementData
	filterKey := fmt.Sprintf("%d_%s", chainID, swapLog.Address.String())
	filter := r.SwapFilters[filterKey]
	if filter == nil {
		return nil, fmt.Errorf("this swap address is not in the server config, chainid: %d, server: %s", chainID, swapLog.Address.String())
	}
	iFace, err := filter.ParseTokenSwap(swapLog)
	if err != nil || iFace == nil {
		return nil, fmt.Errorf("error parsing log, chainid: %d, server: %s", chainID, swapLog.Address.String())
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
	exists, err := r.Refs.CCTPRefs[chainID].IsRequestFulfilled(&bind.CallOpts{
		Context: ctx,
	}, kappaBytes32)
	if err != nil {
		logger.Errorf("Could not check if request exists on chain %d. Error: %v", chainID, err)
		return false
	}
	return exists
}

// will ignore non-bridge events
func (r Resolver) storeBridgeEvent(bridgeEvent interface{}) {
	storeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	storeErr := r.DB.StoreEvent(storeCtx, bridgeEvent)
	if storeErr != nil {
		// Log the error and continue. This function is only called by the get origin/destination function, which its only purpose is to return data
		// from the chain and return it to the user. If storage fails, it should not disrupt this core purpose. Furthermore, we can assume that in
		// the case that storage of this data fails, it will be picked up by scribe and explorer in the next minute and stored correctly.
		logger.Errorf("could not store log while storing origin bridge watcher tx %v", storeErr)
	}
}
