package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/connect"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	chainListeners map[int]listener.ChainListener
	inventory      inventory.InventoryManager
	quoter         *quoter.QuoterManager
	submitter      submitter.TransactionSubmitter
	signer         signer.Signer
	claimCache     *ttlcache.Cache[common.Hash, bool]
}

var logger = log.Logger("relayer")

// NewRelayer creates a new relayer.
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
	chainListeners := make(map[int]listener.ChainListener)

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
		return nil, fmt.Errorf("could not add imanager")
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

func (r *Relayer) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := r.startChainParser(ctx)
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
				return nil
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

	err := g.Wait()
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

		switch request.Status {
		case reldb.Seen:
			// TODO: check it deadline expired
			// get destination commitable balancs
			commitableBalance, err := r.inventory.GetCommittableBalance(ctx, destID, request.Transaction.DestToken)
			if err != nil {
				return fmt.Errorf("could not get commitable balance: %w", err)
			}
			// if commitableBalance > destAmount
			if commitableBalance.Cmp(request.Transaction.DestAmount) > 0 {
				err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.NotEnoughInventory)
			}
			err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.CommittedPending)
			if err != nil {
				return fmt.Errorf("could not update request status: %w", err)
			}

		case reldb.NotEnoughInventory:
			// TODO: recheck if there's enough inventory. Also if it's in this state, you can see if deadline expired

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

func (r *Relayer) startChainParser(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// TODO: good chance we wanna prepare these chain listeners up front and then listen later.
	for chainID, chainCFG := range r.cfg.Bridges {
		chainID := chainID   // capture func literal
		chainCFG := chainCFG // capture func literal

		g.Go(func() error {
			chainListener := r.chainListeners[chainID]

			parser, err := fastbridge.NewParser(common.HexToAddress(chainCFG.Bridge))
			if err != nil {
				return fmt.Errorf("could not parse: %w", err)
			}

			err = chainListener.Listen(ctx, func(ctx context.Context, log types.Log) error {
				_, parsedEvent, ok := parser.ParseEvent(log)
				// handle unknown event
				if !ok {
					if len(log.Topics) != 0 {
						logger.Warnf("unknown event %s", log.Topics[0])
					}
					return nil
				}

				switch event := parsedEvent.(type) {
				case *fastbridge.FastBridgeBridgeRequested:
					err = r.handleNewRequestLog(ctx, event, uint64(chainID))
					if err != nil {
						return fmt.Errorf("could not handle request: %w", err)
					}
				case *fastbridge.FastBridgeBridgeRelayed:
					err = r.handleNewRelayLog(ctx, event)
					if err != nil {
						return fmt.Errorf("could not handle relay: %w", err)
					}
				case *fastbridge.FastBridgeBridgeProofProvided:
					err = r.handleProofProvided(ctx, event)
					if err != nil {
						return fmt.Errorf("could not handle proof provided: %w", err)
					}
				case *fastbridge.FastBridgeBridgeDepositClaimed:
					r.handleDepositClaimed(ctx, event)
				}

				return nil
			})

			if err != nil {
				return fmt.Errorf("listener failed: %w", err)
			}
			return nil
		})
	}
	return nil
}

func (r *Relayer) handleDepositClaimed(ctx context.Context, event *fastbridge.FastBridgeBridgeDepositClaimed) {
	// TODO: err
	_ = r.db.UpdateQuoteRequestStatus(ctx, event.TransactionId, reldb.ClaimCompleted)
}

func (r *Relayer) handleProofProvided(ctx context.Context, req *fastbridge.FastBridgeBridgeProofProvided) (err error) {
	// TODO: this can still get re-orged
	// ALso: we should make sure the previous status  is ProvePosting
	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.ProvePosted)
	if err != nil {
		return fmt.Errorf("could not update request status: %v", err)
	}
	return nil
}

func (r *Relayer) handleNewRelayLog(ctx context.Context, req *fastbridge.FastBridgeBridgeRelayed) (err error) {
	reqID, err := r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	if err != nil {
		return fmt.Errorf("could not get quote request: %v", err)
	}
	// we might've accidentally gotten this later, if so we'll just ignore it
	if reqID.Status != reldb.RelayStarted {
		logger.Warnf("got relay log for request that was not relay started (transaction id: %s, txhash: %s)", hexutil.Encode(reqID.TransactionId[:]), req.Raw.TxHash)
		return nil
	}

	// TODO: this can still get re-orged
	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.RelayCompleted)
	if err != nil {
		return fmt.Errorf("could not update request status: %v", err)
	}
	return nil
}

func (r *Relayer) handleNewRequestLog(parentCtx context.Context, req *fastbridge.FastBridgeBridgeRequested, chainID uint64) (err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "getDecimals", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(req.TransactionId[:])),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: consider a mapmutex
	_, err = r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	// expect no results
	if !errors.Is(err, reldb.ErrNoQuoteForID) {
		// maybe a db err? if so error out & try again later
		if err != nil {
			return fmt.Errorf("could not call db: %w", err)
		}
	}

	// TODO: these should be premade
	originClient, err := r.client.GetChainClient(ctx, int(chainID))
	if err != nil {
		return fmt.Errorf("could not get correct omnirpc client: %w", err)
	}

	fastBridge, err := fastbridge.NewFastBridgeRef(req.Raw.Address, originClient)
	if err != nil {
		return fmt.Errorf("could not get correct fast bridge: %w", err)
	}

	bridgeTx, err := fastBridge.GetBridgeTransaction(&bind.CallOpts{Context: ctx}, req.Request)
	if err != nil {
		return fmt.Errorf("could not get bridge transaction: %w", err)
	}

	decimals, err := r.getDecimals(ctx, bridgeTx)
	if err != nil {
		return fmt.Errorf("could not get decimals: %w", err)
	}

	err = r.db.StoreQuoteRequest(ctx, reldb.QuoteRequest{
		BlockNumber:         req.Raw.BlockNumber,
		RawRequest:          req.Request,
		OriginTokenDecimals: decimals.originDecimals,
		DestTokenDecimals:   decimals.destDecimals,
		TransactionId:       req.TransactionId,
		Sender:              req.Sender,
		Transaction:         bridgeTx,
		Status:              reldb.Seen,
	})
	if err != nil {
		return fmt.Errorf("could not get db: %w", err)
	}

	return nil
}

func (r *Relayer) getDecimals(parentCtx context.Context, bridgeTx fastbridge.IFastBridgeBridgeTransaction) (_ *decimalsRes, err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "getDecimals", trace.WithAttributes(
		attribute.String("sender", bridgeTx.OriginSender.String()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: add a cache for decimals.
	res := decimalsRes{}

	// TODO: cleanup duplication, but keep paralellism.
	// this is  a bit of a pain since it deals w/ 3 different fields, but shouldn't take too long
	originClient, err := r.client.GetChainClient(ctx, int(bridgeTx.OriginChainId))
	if err != nil {
		return nil, fmt.Errorf("could not get origin client: %w", err)
	}

	destClient, err := r.client.GetChainClient(ctx, int(bridgeTx.DestChainId))
	if err != nil {
		return nil, fmt.Errorf("could not get dest client: %w", err)
	}

	originERC20, err := ierc20.NewIERC20(bridgeTx.OriginToken, originClient)
	if err != nil {
		return nil, fmt.Errorf("could not get origin token")
	}

	destERC20, err := ierc20.NewIERC20(bridgeTx.DestToken, destClient)
	if err != nil {
		return nil, fmt.Errorf("could not get dest token")
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		res.originDecimals, err = originERC20.Decimals(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get dest decimals: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		res.destDecimals, err = destERC20.Decimals(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get origin decimals: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get decimals: %w", err)
	}

	return &res, nil
}

type decimalsRes struct {
	originDecimals, destDecimals uint8
}
