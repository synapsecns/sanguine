package inventory

import (
	"context"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

type rebalanceManagerSynapseCCTP struct {
	// cfg is the config
	cfg relconfig.Config
	// handler is the metrics handler
	handler metrics.Handler
	// chainClient is an omnirpc client
	chainClient submitter.ClientFetcher
	// txSubmitter is the transaction submitter
	txSubmitter submitter.TransactionSubmitter
	// cctpContracts is the map of cctp contracts (used for rebalancing)
	cctpContracts map[int]*cctp.SynapseCCTP
	// relayerAddress contains the relayer address
	relayerAddress common.Address
	// chainListeners is the map of chain listeners for CCTP events
	chainListeners map[int]listener.ContractListener
	// db is the database
	db reldb.Service
}

func newRebalanceManagerSynapseCCTP(cfg relconfig.Config, handler metrics.Handler, chainClient submitter.ClientFetcher, txSubmitter submitter.TransactionSubmitter, relayerAddress common.Address, db reldb.Service) *rebalanceManagerSynapseCCTP {
	return &rebalanceManagerSynapseCCTP{
		cfg:            cfg,
		handler:        handler,
		chainClient:    chainClient,
		txSubmitter:    txSubmitter,
		cctpContracts:  make(map[int]*cctp.SynapseCCTP),
		relayerAddress: relayerAddress,
		chainListeners: make(map[int]listener.ContractListener),
		db:             db,
	}
}

func (c *rebalanceManagerSynapseCCTP) Start(ctx context.Context) (err error) {
	err = c.initContracts(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize contracts: %w", err)
	}

	err = c.initListeners(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize listeners: %w", err)
	}

	g, _ := errgroup.WithContext(ctx)
	for cid := range c.cfg.Chains {
		// capture func literal
		chainID := cid
		g.Go(func() error {
			return c.listen(ctx, chainID)
		})
	}

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("error listening to contract: %w", err)
	}
	return nil
}

func (c *rebalanceManagerSynapseCCTP) initContracts(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initContracts")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	for chainID := range c.cfg.Chains {
		contractAddr, err := c.cfg.GetSynapseCCTPAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp address: %w", err)
		}
		chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		contract, err := cctp.NewSynapseCCTP(contractAddr, chainClient)
		if err != nil {
			return fmt.Errorf("could not get cctp: %w", err)
		}
		c.cctpContracts[chainID] = contract
	}
	return nil
}

func (c *rebalanceManagerSynapseCCTP) initListeners(ctx context.Context) (err error) {
	for chainID := range c.cfg.GetChains() {
		cctpAddr, err := c.cfg.GetSynapseCCTPAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp address: %w", err)
		}
		chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		initialBlock, err := c.cfg.GetRebalanceStartBlock(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp start block: %w", err)
		}
		chainListener, err := listener.NewChainListener(chainClient, c.db, cctpAddr, initialBlock, c.handler)
		if err != nil {
			return fmt.Errorf("could not get chain listener: %w", err)
		}
		c.chainListeners[chainID] = chainListener
	}
	return nil
}

func (c *rebalanceManagerSynapseCCTP) Execute(parentCtx context.Context, rebalance *RebalanceData) (err error) {
	contract, ok := c.cctpContracts[rebalance.OriginMetadata.ChainID]
	if !ok {
		return fmt.Errorf("could not find cctp contract for chain %d", rebalance.OriginMetadata.ChainID)
	}
	ctx, span := c.handler.Tracer().Start(parentCtx, "rebalance.Execute", trace.WithAttributes(
		attribute.Int("rebalance_origin", rebalance.OriginMetadata.ChainID),
		attribute.Int("rebalance_dest", rebalance.DestMetadata.ChainID),
		attribute.String("rebalance_amount", rebalance.Amount.String()),
	))
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	// perform rebalance by calling sendCircleToken()
	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = contract.SendCircleToken(
			transactor,
			c.relayerAddress,
			big.NewInt(int64(rebalance.DestMetadata.ChainID)),
			rebalance.OriginMetadata.Addr,
			rebalance.Amount,
			0,        // TODO: inspect
			[]byte{}, // TODO: inspect
		)
		if err != nil {
			return nil, fmt.Errorf("could not send circle token: %w", err)
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit CCTP rebalance: %w", err)
	}

	// store the rebalance in the db
	model := reldb.Rebalance{
		Origin:          uint64(rebalance.OriginMetadata.ChainID), //nolint:gosec // ChainID is always positive and within uint64 range
		Destination:     uint64(rebalance.DestMetadata.ChainID),   //nolint:gosec // ChainID is always positive and within uint64 range
		OriginAmount:    rebalance.Amount,
		Status:          reldb.RebalanceInitiated,
		OriginTokenAddr: rebalance.OriginMetadata.Addr,
		TokenName:       rebalance.OriginMetadata.Name,
	}
	err = c.db.StoreRebalance(ctx, model)
	if err != nil {
		return fmt.Errorf("could not store rebalance: %w", err)
	}
	return nil
}

// nolint:cyclop
func (c *rebalanceManagerSynapseCCTP) listen(parentCtx context.Context, chainID int) (err error) {
	listener, ok := c.chainListeners[chainID]
	if !ok {
		return fmt.Errorf("could not find listener for chain %d", chainID)
	}
	ethClient, err := c.chainClient.GetClient(parentCtx, big.NewInt(int64(chainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	cctpAddr, err := c.cfg.GetSynapseCCTPAddress(chainID)
	if err != nil {
		return fmt.Errorf("could not get cctp address: %w", err)
	}
	parser, err := cctp.NewSynapseCCTPEvents(cctpAddr, ethClient)
	if err != nil {
		return fmt.Errorf("could not get cctp events: %w", err)
	}

	err = listener.Listen(parentCtx, func(parentCtx context.Context, log types.Log) (err error) {
		ctx, span := c.handler.Tracer().Start(parentCtx, "rebalance.Listen", trace.WithAttributes(
			attribute.Int(metrics.ChainID, chainID),
		))
		defer func(err error) {
			metrics.EndSpanWithErr(span, err)
		}(err)

		switch log.Topics[0] {
		case cctp.CircleRequestSentTopic:
			parsedEvent, err := parser.ParseCircleRequestSent(log)
			if err != nil {
				logger.Warnf("could not parse circle request sent: %v", err)
				return nil
			}
			if parsedEvent.Sender != c.relayerAddress {
				return nil
			}
			span.SetAttributes(
				attribute.String("log_type", "CircleRequestSent"),
				attribute.String("request_id", hexutil.Encode(parsedEvent.RequestID[:])),
			)

			// update rebalance model in db
			requestIDHex := hexutil.Encode(parsedEvent.RequestID[:])
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &requestIDHex,
				Origin:          uint64(chainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: parsedEvent.Token,
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case cctp.CircleRequestFulfilledTopic:
			parsedEvent, err := parser.ParseCircleRequestFulfilled(log)
			if err != nil {
				logger.Warnf("could not parse circle request fulfilled: %v", err)
				return nil
			}
			if parsedEvent.Recipient != c.relayerAddress {
				return nil
			}
			span.SetAttributes(
				attribute.String("log_type", "CircleRequestFulfilled"),
				attribute.String("request_id", hexutil.Encode(parsedEvent.RequestID[:])),
			)

			// update rebalance model in db
			requestIDHex := hexutil.Encode(parsedEvent.RequestID[:])
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &requestIDHex,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(parentCtx, rebalanceModel, false)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		default:
			logger.Warnf("unknown event %s", log.Topics[0])
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not listen to contract: %w", err)
	}
	return nil
}
