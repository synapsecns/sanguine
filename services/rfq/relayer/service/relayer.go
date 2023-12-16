package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/connect"
	"golang.org/x/sync/errgroup"
	"math/big"
	"time"
)

// Relayer is the core of the relayer application.
type Relayer struct {
	cfg            relconfig.Config
	metrics        metrics.Handler
	db             reldb.Service
	client         omnirpcClient.RPCClient
	chainListeners map[int]listener.ContractListener
	inventory      inventory.InventoryManager
	quoter         *quoter.QuoterManager
	submitter      submitter.TransactionSubmitter
	signer         signer.Signer
	claimCache     *ttlcache.Cache[common.Hash, bool]
}

var logger = log.Logger("relayer")

// NewRelayer creates a new relayer.
//
// The relayer is the core of the application. It is responsible for starting the listener and quoter event loops.
func NewRelayer(ctx context.Context, metricHandler metrics.Handler, cfg relconfig.Config) (*Relayer, error) {
	omniClient := omnirpcClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omnirpcClient.WithCaptureReqRes())

	// TODO: pull from config
	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}

	store, err := connect.Connect(ctx, dbType, cfg.Database.DSN, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}
	chainListeners := make(map[int]listener.ContractListener)

	// setup chain listeners
	for chainID, chainCFG := range cfg.Bridges {
		// TODO: consider getter for this convert step
		bridge := common.HexToAddress(chainCFG.Bridge)
		chainClient, err := omniClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, store, bridge, metricHandler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		chainListeners[chainID] = chainListener
	}

	sg, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer")
	}

	im, err := inventory.NewInventoryManager(ctx, omniClient, metricHandler, cfg, sg.Address(), store)
	if err != nil {
		return nil, fmt.Errorf("could not add imanager: %w", err)
	}

	q, err := quoter.NewQuoterManager(ctx, cfg.QuotableTokens, im, cfg.RfqAPIURL, sg)
	if err != nil {
		return nil, fmt.Errorf("could not get quoter")
	}

	sm := submitter.NewTransactionSubmitter(metricHandler, sg, omniClient, store.SubmitterDB(), &cfg.SubmitterConfig)

	cache := ttlcache.New[common.Hash, bool](ttlcache.WithTTL[common.Hash, bool](time.Second * 30))
	rel := Relayer{
		db:             store,
		client:         omniClient,
		quoter:         q,
		metrics:        metricHandler,
		claimCache:     cache,
		cfg:            cfg,
		inventory:      im,
		submitter:      sm,
		signer:         sg,
		chainListeners: chainListeners,
	}
	return &rel, nil
}

const defaultPostInterval = 1

// Start starts the relayer.
//
// This will:
// 1. Check if approvals need to be issued on chain & issue them if needed. This allows
// _pullToken to function correctly.
// 2. Start the chain parser: This will listen to and process any events on chain
// 3. Start the db selector: This will check the db for any requests that need to be processed.
// 4. Start the submitter: This will submit any transactions that need to be submitted.
func (r *Relayer) Start(ctx context.Context) error {
	err := r.inventory.ApproveAllTokens(ctx, r.submitter)
	if err != nil {
		return fmt.Errorf("could not approve all tokens: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := r.startChainIndexers(ctx)
		if err != nil {
			return fmt.Errorf("could not start chain parser: %w", err)
		}
		return nil
	})

	go r.claimCache.Start()
	go func() {
		select {
		case <-ctx.Done():
			r.claimCache.Stop()
		}
	}()

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return fmt.Errorf("could not start db selector: %w", ctx.Err())
			case <-time.After(defaultPostInterval * time.Second):
				err := r.quoter.SubmitAllQuotes()
				if err != nil {
					return fmt.Errorf("could not start db selector: %w", err)
				}
			}
		}
	})

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(defaultPostInterval * time.Second):
				err := r.runDBSelector(ctx)
				if err != nil {
					return fmt.Errorf("could not start db selector: %w", err)
				}
			}
		}
	})

	g.Go(func() error {
		err := r.submitter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not start submitter: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not start: %w", err)
	}

	return nil
}

// TODO: make this configurable.
const dbSelectorInterval = 1

func (r *Relayer) runDBSelector(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("could not run db selector: %w", ctx.Err())
		case <-time.After(dbSelectorInterval * time.Second):
			// TODO: add context w/ timeout
			// TODO: add trigger
			err := r.processDB(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func (r *Relayer) processDB(ctx context.Context) error {
	requests, err := r.db.GetQuoteResultsByStatus(ctx, reldb.Seen, reldb.CommittedPending, reldb.CommittedConfirmed, reldb.RelayCompleted, reldb.ProvePosted)
	if err != nil {
		return nil
	}
	// Obviously, these are only seen.
	for _, request := range requests {
		originID := int(request.Transaction.OriginChainId)
		destID := int(request.Transaction.DestChainId)
		// TODO: check for deadline expired. if so mark and continue.

		originClient, err := r.client.GetChainClient(ctx, originID)
		if err != nil {
			logger.Errorf("could not get origin client: %v", err)
			continue
		}

		originFastBridge, err := fastbridge.NewFastBridgeRef(common.HexToAddress(r.cfg.Bridges[originID].Bridge), originClient)
		if err != nil {
			logger.Errorf("could not get origin fast bridge: %v", err)
			continue
		}

		destClient, err := r.client.GetChainClient(ctx, destID)
		if err != nil {
			logger.Errorf("could not get dest client: %v", err)
			continue
		}

		destFastBridge, err := fastbridge.NewFastBridgeRef(common.HexToAddress(r.cfg.Bridges[destID].Bridge), destClient)
		if err != nil {
			logger.Errorf("could not get dest fast bridge: %v", err)
			continue
		}

		// if deadline < now
		if request.Transaction.Deadline.Cmp(big.NewInt(time.Now().Unix())) < 0 && request.Status.Int() < reldb.RelayCompleted.Int() {
			err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.DeadlineExceeded)
		}

		switch request.Status {
		case reldb.Seen:
			if !r.quoter.ShouldProcess(request) {
				err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.WillNotProcess)
				if err != nil {
					return fmt.Errorf("could not update request status: %w", err)
				}
			}
			// TODO: check it deadline expired
			// TODO: check token validity
			// get destination commitable balancs
			commitableBalance, err := r.inventory.GetCommittableBalance(ctx, destID, request.Transaction.DestToken)
			if err != nil {
				return fmt.Errorf("could not get commitable balance: %w", err)
			}
			// if commitableBalance > destAmount
			if commitableBalance.Cmp(request.Transaction.DestAmount) < 0 {
				err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.NotEnoughInventory)
				if err != nil {
					return fmt.Errorf("could not update request status: %w", err)
				}
				continue
			}
			err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.CommittedPending)
			if err != nil {
				return fmt.Errorf("could not update request status: %w", err)
			}

		case reldb.NotEnoughInventory:
			commitableBalance, err := r.inventory.GetCommittableBalance(ctx, destID, request.Transaction.DestToken)
			if err != nil {
				return fmt.Errorf("could not get commitable balance: %w", err)
			}
			// if commitableBalance > destAmount
			if commitableBalance.Cmp(request.Transaction.DestAmount) > 0 {
				err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.NotEnoughInventory)
				if err != nil {
					return fmt.Errorf("could not update request status: %w", err)
				}
			}

		case reldb.CommittedPending:
			// TODO: build this in somehwere else  afte rwe commit
			// need to see if we can complete
			earliestConfirmBlock := request.BlockNumber + r.cfg.Bridges[originID].Confirmations

			if earliestConfirmBlock < r.chainListeners[originID].LatestBlock() {
				// can't complete, yet do nothing
				continue
			}

			// It's here: so at this point, we wanna check if it's still on chain.
			// TODO: this will cover cases where this got reorged out, but they will stay in the queue forever
			// should clean this eventually.
			//
			// Reorgs are rare enough that its questionable wether this is ever worth building or if we can just
			// leave these in the queue.
			bs, err := originFastBridge.BridgeStatuses(&bind.CallOpts{Context: ctx}, request.TransactionId)
			if err != nil {
				return fmt.Errorf("could not get bridge status: %w", err)
			}

			// sanity check to make sure it's still requested.
			if bs == fastbridge.REQUESTED.Int() {
				err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.CommittedConfirmed)
				if err != nil {
					return fmt.Errorf("could not update request status: %w", err)
				}
			}
		case reldb.CommittedConfirmed:
			nonce, err := r.submitter.SubmitTransaction(ctx, big.NewInt(int64(destID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
				tx, err = destFastBridge.Relay(transactor, request.RawRequest)
				if err != nil {
					return nil, fmt.Errorf("could not relay: %w", err)
				}

				return tx, nil
			})
			if err != nil {
				return fmt.Errorf("could not submit transaction: %w", err)
			}

			err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.RelayStarted)
			// TODO:
			_ = nonce

			if err != nil {
				return fmt.Errorf("could not update request status: %w", err)
			}
		case reldb.RelayCompleted:
			// relays been completed, it's time to go back to the origin chain and try to prove
			_, err := r.submitter.SubmitTransaction(ctx, big.NewInt(int64(originID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
				// MAJO MAJOR TODO should be dest tx hash
				tx, err = originFastBridge.Prove(transactor, request.RawRequest, request.TransactionId)
				if err != nil {
					return nil, fmt.Errorf("could not relay: %w", err)
				}

				return tx, nil
			})
			if err != nil {
				return fmt.Errorf("could not submit transaction: %w", err)
			}

			err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.ProvePosting)
			if err != nil {
				return fmt.Errorf("could not update request status: %w", err)
			}
		case reldb.ProvePosted:
			// we use claim cache to make sure we don't hit the rpc to check to often
			if r.claimCache.Has(request.TransactionId) {
				continue
			}

			r.claimCache.Set(request.TransactionId, true, 30*time.Second)

			canClaim, err := originFastBridge.CanClaim(&bind.CallOpts{Context: ctx}, request.TransactionId, r.signer.Address())
			if err != nil {
				return fmt.Errorf("could not check if can claim: %w", err)
			}

			if !canClaim {
				continue
			}
			_, err = r.submitter.SubmitTransaction(ctx, big.NewInt(int64(originID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
				tx, err = originFastBridge.Claim(transactor, request.RawRequest, transactor.From)
				if err != nil {
					return nil, fmt.Errorf("could not relay: %w", err)
				}

				return tx, nil
			})
			if err != nil {
				return fmt.Errorf("could not submit transaction: %w", err)
			}

			err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.ClaimPending)
			if err != nil {
				return fmt.Errorf("could not update request status: %w", err)
			}
		}
	}
	return nil
}
