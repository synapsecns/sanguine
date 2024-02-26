package inventory

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"golang.org/x/sync/errgroup"
)

// RebalanceData contains metadata for a rebalance action.
type RebalanceData struct {
	OriginMetadata *TokenMetadata
	DestMetadata   *TokenMetadata
	Amount         *big.Int
}

// RebalanceManager is the interface for the rebalance manager.
type RebalanceManager interface {
	// Start starts the rebalance manager.
	Start(ctx context.Context) (err error)
	// Execute executes a rebalance action.
	Execute(ctx context.Context, rebalance *RebalanceData) error
}

type rebalanceManagerCCTP struct {
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

func newRebalanceManagerCCTP(cfg relconfig.Config, handler metrics.Handler, chainClient submitter.ClientFetcher, txSubmitter submitter.TransactionSubmitter, relayerAddress common.Address, db reldb.Service) *rebalanceManagerCCTP {
	return &rebalanceManagerCCTP{
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

func (c *rebalanceManagerCCTP) Start(ctx context.Context) error {
	// initialize contracts
	for chainID := range c.cfg.Chains {
		contractAddr, err := c.cfg.GetCCTPAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp address: %w", err)
		}
		chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		fmt.Printf("CCTP contract addr for chain %d: %v\n", chainID, contractAddr)
		contract, err := cctp.NewSynapseCCTP(common.HexToAddress(contractAddr), chainClient)
		if err != nil {
			return fmt.Errorf("could not get cctp: %w", err)
		}
		c.cctpContracts[chainID] = contract
	}

	// initialize chain listeners
	for chainID := range c.cfg.GetChains() {
		cctpAddr, err := c.cfg.GetCCTPAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get cctp address: %w", err)
		}
		chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, c.db, common.HexToAddress(cctpAddr), c.handler)
		if err != nil {
			return fmt.Errorf("could not get chain listener: %w", err)
		}
		c.chainListeners[chainID] = chainListener
	}

	g, _ := errgroup.WithContext(ctx)
	for cid := range c.cfg.Chains {
		// capture func literal
		chainID := cid
		g.Go(func() error {
			return c.listen(ctx, chainID)
		})
	}

	return g.Wait()
}

func (c *rebalanceManagerCCTP) Execute(ctx context.Context, rebalance *RebalanceData) (err error) {
	contract, ok := c.cctpContracts[rebalance.OriginMetadata.ChainID]
	if !ok {
		return fmt.Errorf("could not find cctp contract for chain %d", rebalance.OriginMetadata.ChainID)
	}

	// perform rebalance by calling sendCircleToken()
	fmt.Printf("[cctp] execute rebalance on chain %d and token %v\n", rebalance.OriginMetadata.ChainID, rebalance.OriginMetadata.Addr)
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
		fmt.Printf("[cctp] rebalance tx: %s\n", tx.Hash().Hex())
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit CCTP rebalance: %w", err)
	}

	// store the rebalance in the db
	model := reldb.Rebalance{
		Origin:       uint64(rebalance.OriginMetadata.ChainID),
		Destination:  uint64(rebalance.DestMetadata.ChainID),
		OriginAmount: rebalance.Amount,
		Status:       reldb.RebalanceInitiated,
	}
	err = c.db.StoreRebalance(ctx, model)
	if err != nil {
		return fmt.Errorf("could not store rebalance: %w", err)
	}
	return nil
}

// nolint:cyclop
func (c *rebalanceManagerCCTP) listen(ctx context.Context, chainID int) (err error) {
	fmt.Printf("listen on chain %d\n", chainID)
	listener, ok := c.chainListeners[chainID]
	if !ok {
		return fmt.Errorf("could not find listener for chain %d", chainID)
	}
	ethClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	cctpAddr := common.HexToAddress(c.cfg.Chains[chainID].CCTPAddress)
	fmt.Printf("CCTP listen on chain %d: %s\n", chainID, cctpAddr.String())
	parser, err := cctp.NewSynapseCCTPEvents(cctpAddr, ethClient)
	if err != nil {
		return fmt.Errorf("could not get cctp events: %w", err)
	}

	return listener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		switch log.Topics[0] {
		case cctp.CircleRequestSentTopic:
			parsedEvent, err := parser.ParseCircleRequestSent(log)
			if err != nil {
				return fmt.Errorf("could not parse circle request sent: %w", err)
			}
			origin := uint64(parsedEvent.ChainId.Int64())
			err = c.db.UpdateRebalanceStatus(parentCtx, parsedEvent.RequestID, &origin, reldb.RebalancePending)
			if err != nil {
				return fmt.Errorf("could not update rebalance status: %w", err)
			}
		case cctp.CircleRequestFulfilledTopic:
			parsedEvent, err := parser.ParseCircleRequestFulfilled(log)
			if err != nil {
				return fmt.Errorf("could not parse circle request fulfilled: %w", err)
			}
			err = c.db.UpdateRebalanceStatus(parentCtx, parsedEvent.RequestID, nil, reldb.RebalanceCompleted)
			if err != nil {
				return fmt.Errorf("could not update rebalance status: %w", err)
			}
		default:
			logger.Warnf("unknown event %s", log.Topics[0])
		}
		return nil
	})
}
