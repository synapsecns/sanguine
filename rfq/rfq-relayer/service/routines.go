package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/listener"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"
	"golang.org/x/sync/errgroup"

	"math/big"

	"time"
)

func (r *relayerImpl) RunListeners(parentCtx context.Context) error {
	g, gCtx := errgroup.WithContext(parentCtx)

	// Start the event listener on every chain
	for _, chainCfg := range r.config.Chains {
		chainID := chainCfg.ChainID
		g.Go(func() error {
			chainListener, err := listener.NewChainListener(r.chainConfigs[chainID], r.db, r.eventChan, r.seenChan)
			if err != nil {
				return fmt.Errorf("could not create chain listener: %w", err)
			}
			err = chainListener.StartListening(gCtx)
			if err != nil {
				return fmt.Errorf("could not start chain listener: %w", err)
			}
			return nil
		})
	}
	return nil
}
func (r *relayerImpl) HandleClaimEvents(parentCtx context.Context) error {
	for {
		select {
		case <-parentCtx.Done():
			return nil
		case <-time.After(time.Duration(r.config.QueuePollInterval) * time.Second):
			// Check if head of queue is ready
			now := time.Now().Unix()
			deadlinePassed, err := r.claimQueue.HasLiveElements(now)
			if err != nil {
				return fmt.Errorf("could not peek head of queue: %w", err)
			}

			// If deadline has passed, dequeue and execute claim
			if deadlinePassed {
				transactionID, qErr := r.claimQueue.Dequeue(parentCtx)
				if qErr != nil {
					return fmt.Errorf("could not dequeue event: %w", qErr)
				}
				originBridgeEvent, qErr := r.db.GetOriginBridgeEvent(parentCtx, transactionID)
				if qErr != nil {
					return fmt.Errorf("could not get origin bridge event: %w", qErr)
				}
				requestBytes := common.Hex2Bytes(originBridgeEvent.Request)
				nonce, qErr := r.txSubmitter.SubmitTransaction(parentCtx, big.NewInt(int64(originBridgeEvent.OriginChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
					transactor.Value = big.NewInt(0)
					tx, err = r.contracts[originBridgeEvent.OriginChainID].Claim(transactor, requestBytes, common.HexToAddress(r.config.RelayerAddress))
					if err != nil {
						logger.Errorf("could not submit transaction: %w", err)
						return nil, fmt.Errorf("could not submit transaction: %w", err)
					}

					return tx, nil
				})
				if qErr != nil {
					return fmt.Errorf("could not submit claim transaction: %w", qErr)
				}

				// Log success
				logger.Infof("submitted claim at nonce %d", nonce)
			}
		}
	}
}
func (r *relayerImpl) HandleUnconfirmedEvents(parentCtx context.Context) error {
	for {
		select {
		case <-parentCtx.Done():
			return nil
		case wrappedLog := <-r.seenChan:
			if utils.IsBridgeRequested(wrappedLog.Log.Topics[0], r.chainConfigs[wrappedLog.OriginChainID].ABI) {
				bridgeEvent, err := utils.ParseBridgeRequested(wrappedLog.Log, r.chainConfigs[wrappedLog.OriginChainID].ABI)
				if err != nil {
					localError := fmt.Errorf("error parsing bridge requested event %v for tx %s on chain %d", wrappedLog.Log, wrappedLog.Log.TxHash.Hex(), wrappedLog.OriginChainID)
					logger.Error(localError)
					return errors.Wrap(err, localError.Error()) // This will trigger a restart of the relayer, and make it re-index the range.
				}
				err = r.quoter.HandleUnconfirmedBridgeRequest(bridgeEvent)
				if err != nil {
					localError := fmt.Errorf("error while handling unconfirmed bridge event %v for tx %s on chain %d", wrappedLog.Log, wrappedLog.Log.TxHash.Hex(), wrappedLog.OriginChainID)
					logger.Error(err)
					return errors.Wrap(err, localError.Error()) // This will trigger a restart of the relayer, and make it re-index the range.
				}
			}
		}
	}
}

func (r *relayerImpl) HandleConfirmedEvents(parentCtx context.Context) error {
	for {
		select {
		case <-parentCtx.Done():
			return nil
		case wrappedLog := <-r.eventChan:
			topic := wrappedLog.Log.Topics[0]
			abi := r.chainConfigs[wrappedLog.OriginChainID].ABI
			// TODO: rewrite as switch
			//nolint: gocritic, nestif
			if utils.IsBridgeRequested(topic, abi) {
				// Parse BridgeRequested
				bridgeEvent, err := utils.ParseBridgeRequested(wrappedLog.Log, r.chainConfigs[wrappedLog.OriginChainID].ABI)
				if err != nil {
					localErr := fmt.Errorf("error parsing bridge requested event %v for tx %s on chain %d", wrappedLog.Log, wrappedLog.Log.TxHash.Hex(), wrappedLog.OriginChainID)
					logger.Error(localErr)
					return errors.Wrap(err, localErr.Error()) // This will trigger a restart of the relayer, and make it re-index the range.
				}

				// Process BridgeRequested
				err = r.TryProcessBridgeRequested(parentCtx, bridgeEvent)
				if err != nil { // non fatal
					logger.Errorf("error while handling confirmed bridge event %v for tx %s on chain %d", wrappedLog.Log, wrappedLog.Log.TxHash.Hex(), wrappedLog.OriginChainID)
				}
			} else if utils.IsBridgeRelayed(topic, abi) {
				// Parse BridgeRelayed
				bridgeEvent, err := utils.ParseBridgeRelayed(wrappedLog.Log, r.chainConfigs[wrappedLog.OriginChainID].ABI)
				if err != nil {
					localErr := fmt.Errorf("error parsing bridge relayed event %v for tx %s on chain %d", wrappedLog.Log, wrappedLog.Log.TxHash.Hex(), wrappedLog.OriginChainID)
					logger.Error(localErr)
					return errors.Wrap(err, localErr.Error()) // This will trigger a restart of the relayer, and make it re-index the range.
				}
				// Process BridgeRelayed
				err = r.TryProcessBridgeRelayed(parentCtx, bridgeEvent)
				if err != nil { // non fatal
					logger.Errorf("error while handling confirmed bridge relay event %v for tx %s on chain %d", wrappedLog.Log, wrappedLog.Log.TxHash.Hex(), wrappedLog.OriginChainID)
				}
			} else {
				return fmt.Errorf("unknown event type")
			}
		}
	}
}

// TryProcessBridgeRequested processes a BridgeRequested event, and executes a relay if it matches a current relayer.
// nolint: cyclop
func (r *relayerImpl) TryProcessBridgeRequested(ctx context.Context, req *bindings.FastBridgeBridgeRequested) error {
	transactionID := common.Bytes2Hex(req.TransactionId[:]) // keccak256 hash of the request
	event, err := utils.Decode(req.Request)
	if err != nil {
		return fmt.Errorf("could not decode bridge request: %w", err)
	}

	// Check if the requested bridge can be fulrelayed by this relayer
	quoteID := utils.GenerateQuoteID(event.OriginChainId, event.OriginToken, event.DestChainId, event.DestToken)
	destTokenID := utils.GenerateTokenID(event.DestChainId, event.DestToken)
	_, err = r.quoter.GetValidQuote(quoteID, destTokenID, event.DestAmount)
	// Quote is not valid, volume out of bounds or there's not a quote for that token pair on this relayer
	if err != nil {
		return fmt.Errorf("could not furelay seen bridge event: %w", err)
	}

	// Execute relay
	var logs []*types.Log
	nonce, err := r.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(event.OriginChainId)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		transactor.Value = big.NewInt(0)
		tx, err = r.contracts[event.OriginChainId].Relay(transactor, req.Request)
		if err != nil {
			logger.Errorf("could not submit transaction: %w", err)
			// roll back quotes/balance
			err = r.quoter.HandleUncompletedBridge(transactionID, event)
			if err != nil {
				logger.Errorf("could not roll back balanced during relay error: %v", err)
			}
			return nil, fmt.Errorf("could not submit transaction: %w", err)
		}

		// Get receipt for storing event
		receipt, err := r.evmClients[event.OriginChainId].TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, fmt.Errorf("could not get transaction receipt: %w", err)
		}
		// Get logs from receipt
		logs = receipt.Logs
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit relay transaction: %w", err)
	}

	// Log success
	logger.Infof("submitted transaction at nonce %d", nonce)

	// Update Balance
	err = r.quoter.HandleCompletedBridge(transactionID, event)
	if err != nil {
		return fmt.Errorf("could not update balance after relaying: %w", err)
	}

	// Insert into database
	for _, log := range logs {
		if utils.IsBridgeRequested(log.Topics[0], r.chainConfigs[event.OriginChainId].ABI) {
			if r.db.StoreOriginBridgeEvent(ctx, event.OriginChainId, log, req) != nil {
				return fmt.Errorf("could not store origin bridge event: %w", err)
			}
			return nil
		}
	}
	return nil
}

func (r *relayerImpl) TryProcessBridgeRelayed(ctx context.Context, req *bindings.FastBridgeBridgeRelayed) error {
	// Only process this event if the relayer that did the relay is this relayer.
	if req.Relayer.String() != r.config.RelayerAddress {
		return nil
	}
	transactionIDStr := common.Bytes2Hex(req.TransactionId[:])
	// Get origin bridge request
	originBridgeEvent, err := r.db.GetOriginBridgeEvent(ctx, transactionIDStr)
	if err != nil {
		return fmt.Errorf("could not get origin bridge event: %w", err)
	}
	requestBytes := common.Hex2Bytes(originBridgeEvent.Request)

	// Get tx hash
	var txHashBytes [32]byte
	copy(txHashBytes[:], common.Hex2Bytes(originBridgeEvent.TxHash))

	// Execute prove
	var logs []*types.Log
	nonce, err := r.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(originBridgeEvent.OriginChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		transactor.Value = big.NewInt(0)
		tx, err = r.contracts[originBridgeEvent.OriginChainID].Prove(transactor, requestBytes, txHashBytes)
		if err != nil {
			logger.Errorf("could not submit transaction: %w", err)
			return nil, fmt.Errorf("could not submit transaction: %w", err)
		}

		receipt, err := r.evmClients[originBridgeEvent.DestChainID].TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, fmt.Errorf("could not get transaction receipt: %w", err)
		}
		// Get logs from receipt
		logs = receipt.Logs
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit prove transaction: %w", err)
	}

	// Log success
	logger.Infof("submitted prove at nonce %d", nonce)

	// Add to claim queue
	err = r.claimQueue.Enqueue(ctx, transactionIDStr)

	// Insert into database
	for _, log := range logs {
		if utils.IsBridgeRelayed(log.Topics[0], r.chainConfigs[originBridgeEvent.DestChainID].ABI) {
			if r.db.StoreDestinationBridgeEvent(ctx, log, originBridgeEvent) != nil {
				return fmt.Errorf("could not store destination bridge event: %w", err)
			}
			return nil
		}
	}
	return nil
}
