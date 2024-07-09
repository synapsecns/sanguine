package inventory

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l1gateway"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l2gateway"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"golang.org/x/sync/errgroup"
)

type rebalanceManagerScroll struct {
	// cfg is the config
	cfg relconfig.Config
	// handler is the metrics handler
	handler metrics.Handler
	// chainClient is an omnirpc client
	chainClient submitter.ClientFetcher
	// txSubmitter is the transaction submitter
	txSubmitter submitter.TransactionSubmitter
	// relayerAddress contains the relayer address
	relayerAddress common.Address
	// boundL1Gateway is the L1GatewayRouter contract
	boundL1Gateway *l1gateway.L1GatewayRouter
	// boundL2Gateway is the L2GatewayRouter contract
	boundL2Gateway *l2gateway.L2GatewayRouter
	// l1GatewayListener is the listener for the L1GatewayRouter contract
	l1GatewayListener listener.ContractListener
	// l2GatewayListener is the listener for the L2GatewayRouter contract
	l2GatewayListener listener.ContractListener
	// l1ChainID is the chain ID for the L1 chain
	l1ChainID int
	// l2ChainID is the chain ID for the L2 chain
	l2ChainID int
	// db is the database
	db reldb.Service
}

func newRebalanceManagerScroll(cfg relconfig.Config, handler metrics.Handler, chainClient submitter.ClientFetcher, txSubmitter submitter.TransactionSubmitter, relayerAddress common.Address, db reldb.Service) *rebalanceManagerScroll {
	return &rebalanceManagerScroll{
		cfg:            cfg,
		handler:        handler,
		chainClient:    chainClient,
		txSubmitter:    txSubmitter,
		relayerAddress: relayerAddress,
		db:             db,
	}
}

func isScrollChain(chainID int) bool {
	return chainID == 534352
}

func isMainnetChain(chainID int) bool {
	return chainID == 1
}

func (c *rebalanceManagerScroll) Start(ctx context.Context) (err error) {
	err = c.initContracts(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize contracts: %w", err)
	}

	err = c.initListeners(ctx)
	if err != nil {
		return fmt.Errorf("could not initialize listeners: %w", err)
	}

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		l1Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l1ChainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		return c.listenL1Gateway(ctx, l1Client)
	})
	g.Go(func() error {
		l2Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l2ChainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		return c.listenL2Gateway(ctx, l2Client)
	})

	return nil
}

func (c *rebalanceManagerScroll) initContracts(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initContracts")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	for chainID := range c.cfg.Chains {
		if isMainnetChain(chainID) {
			c.l1ChainID = chainID
			addr, err := c.cfg.GetL1GatewayAddress(chainID)
			if err != nil {
				return fmt.Errorf("could not get l1 gateway address: %w", err)
			}
			chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
			if err != nil {
				return fmt.Errorf("could not get chain client: %w", err)
			}
			c.boundL1Gateway, err = l1gateway.NewL1GatewayRouter(common.HexToAddress(addr), chainClient)
			if err != nil {
				return fmt.Errorf("could not get l1 gateway contract: %w", err)
			}
		} else if isScrollChain(chainID) {
			c.l2ChainID = chainID
			addr, err := c.cfg.GetL2GatewayAddress(chainID)
			if err != nil {
				return fmt.Errorf("could not get l2 gateway address: %w", err)
			}
			chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
			if err != nil {
				return fmt.Errorf("could not get chain client: %w", err)
			}
			c.boundL2Gateway, err = l2gateway.NewL2GatewayRouter(common.HexToAddress(addr), chainClient)
			if err != nil {
				return fmt.Errorf("could not get l2 gateway contract: %w", err)
			}
		}
	}
	if c.boundL1Gateway == nil {
		return fmt.Errorf("l1 gateway contract not set")
	}
	if c.boundL2Gateway == nil {
		return fmt.Errorf("l2 gateway contract not set")
	}

	return nil
}

func (c *rebalanceManagerScroll) initListeners(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initListeners")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	// setup l1 listener
	l1Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l1ChainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	l1InitialBlock, err := c.cfg.GetCCTPStartBlock(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get cctp start block: %w", err)
	}
	l1Addr, err := c.cfg.GetL1GatewayAddress(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get l1 gateway address: %w", err)
	}
	c.l1GatewayListener, err = listener.NewChainListener(l1Client, c.db, common.HexToAddress(l1Addr), l1InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get messenger listener: %w", err)
	}

	// setup l2 listener
	l2Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l2ChainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	l2InitialBlock, err := c.cfg.GetCCTPStartBlock(c.l2ChainID)
	if err != nil {
		return fmt.Errorf("could not get cctp start block: %w", err)
	}
	l2Addr, err := c.cfg.GetL2GatewayAddress(c.l2ChainID)
	if err != nil {
		return fmt.Errorf("could not get l2 gateway address: %w", err)
	}
	c.l2GatewayListener, err = listener.NewChainListener(l2Client, c.db, common.HexToAddress(l2Addr), l2InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get messenger listener: %w", err)
	}

	return nil
}

func (c *rebalanceManagerScroll) Execute(ctx context.Context, rebalance *RebalanceData) (err error) {
	switch rebalance.OriginMetadata.ChainID {
	case c.l1ChainID:
		err = c.executeL1ToL2(ctx, rebalance)
	case c.l2ChainID:
		err = c.executeL2ToL1(ctx, rebalance)
	default:
		return fmt.Errorf("unexpected origin: %d", rebalance.OriginMetadata.ChainID)
	}
	if err != nil {
		return fmt.Errorf("could not execute rebalance: %w", err)
	}

	// store the rebalance in the db
	rebalanceModel := reldb.Rebalance{
		Origin:       uint64(rebalance.OriginMetadata.ChainID),
		Destination:  uint64(rebalance.DestMetadata.ChainID),
		OriginAmount: rebalance.Amount,
		Status:       reldb.RebalanceInitiated,
	}
	err = c.db.StoreRebalance(ctx, rebalanceModel)
	if err != nil {
		return fmt.Errorf("could not store rebalance: %w", err)
	}
	return nil
}

// TODO: configurable?
const scrollGasLimit = 200_000

func (c *rebalanceManagerScroll) executeL1ToL2(ctx context.Context, rebalance *RebalanceData) (err error) {
	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		if chain.IsGasToken(rebalance.OriginMetadata.Addr) {
			tx, err = c.boundL1Gateway.DepositETH(transactor, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not deposit gas token: %w", err)
			}
		} else {
			tx, err = c.boundL1Gateway.DepositERC20(transactor, rebalance.OriginMetadata.Addr, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not deposit erc20 token: %w", err)
			}
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}
	return nil
}

func (c *rebalanceManagerScroll) executeL2ToL1(ctx context.Context, rebalance *RebalanceData) (err error) {
	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		if chain.IsGasToken(rebalance.OriginMetadata.Addr) {
			tx, err = c.boundL2Gateway.WithdrawETH(transactor, c.relayerAddress, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not withdraw gas token: %w", err)
			}
		} else {
			tx, err = c.boundL2Gateway.WithdrawERC20(transactor, rebalance.OriginMetadata.Addr, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not withdraw erc20 token: %w", err)
			}
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}
	return nil
}

func getScrollRebalanceID(eventData []byte) string {
	return common.BytesToHash(eventData).Hex()
}

func (c *rebalanceManagerScroll) listenL1Gateway(ctx context.Context, ethClient client.EVM) (err error) {
	addr, err := c.cfg.GetL1GatewayAddress(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get l1 gateway address: %w", err)
	}
	parser, err := l1gateway.NewL1GatewayRouterFilterer(common.HexToAddress(addr), ethClient)
	if err != nil {
		return fmt.Errorf("could not get l1 gateway parser: %w", err)
	}
	err = c.l1GatewayListener.Listen(ctx, func(ctx context.Context, log types.Log) (err error) {
		switch log.Topics[0] {
		case l1gateway.DepositETHTopic:
			event, err := parser.ParseDepositETH(log)
			if err != nil {
				return fmt.Errorf("could not parse DepositETH event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l1ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: chain.EthAddress,
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case l1gateway.DepositERC20Topic:
			event, err := parser.ParseDepositERC20(log)
			if err != nil {
				return fmt.Errorf("could not parse DepositERC20 event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l1ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: event.L1Token,
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case l1gateway.FinalizeWithdrawETHTopic:
			event, err := parser.ParseFinalizeWithdrawETH(log)
			if err != nil {
				return fmt.Errorf("could not parse FinalizeWithdrawETH event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case l1gateway.FinalizeWithdrawERC20Topic:
			event, err := parser.ParseFinalizeWithdrawERC20(log)
			if err != nil {
				return fmt.Errorf("could not parse FinalizeWithdrawERC20 event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("could not listen for L1GatewayRouter events: %w", err)
	}
	return nil
}

func (c *rebalanceManagerScroll) listenL2Gateway(ctx context.Context, ethClient client.EVM) (err error) {
	addr, err := c.cfg.GetL2GatewayAddress(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get l2 gateway address: %w", err)
	}
	parser, err := l2gateway.NewL2GatewayRouterFilterer(common.HexToAddress(addr), ethClient)
	if err != nil {
		return fmt.Errorf("could not get l2 gateway parser: %w", err)
	}
	err = c.l2GatewayListener.Listen(ctx, func(ctx context.Context, log types.Log) (err error) {
		switch log.Topics[0] {
		case l2gateway.WithdrawETHTopic:
			event, err := parser.ParseWithdrawETH(log)
			if err != nil {
				return fmt.Errorf("could not parse WithdrawETH event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l2ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: chain.EthAddress,
				Destination:     uint64(c.l1ChainID),
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case l2gateway.WithdrawERC20Topic:
			event, err := parser.ParseWithdrawERC20(log)
			if err != nil {
				return fmt.Errorf("could not parse WithdrawERC20 event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID:     &rebalanceID,
				Origin:          uint64(c.l2ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: event.L2Token,
				Destination:     uint64(c.l1ChainID),
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case l2gateway.FinalizeDepositETHTopic:
			event, err := parser.ParseFinalizeDepositETH(log)
			if err != nil {
				return fmt.Errorf("could not parse FinalizeDepositETH event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case l2gateway.FinalizeDepositERC20Topic:
			event, err := parser.ParseFinalizeDepositERC20(log)
			if err != nil {
				return fmt.Errorf("could not parse FinalizeDepositERC20 event: %w", err)
			}
			rebalanceID := getScrollRebalanceID(event.Data)
			rebalanceModel := reldb.Rebalance{
				RebalanceID: &rebalanceID,
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateRebalance(ctx, rebalanceModel, true)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("could not listen for L2GatewayRouter events: %w", err)
	}
	return nil
}
