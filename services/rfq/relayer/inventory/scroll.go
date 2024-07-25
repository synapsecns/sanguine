package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l1gateway"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l1scrollmessenger"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l2gateway"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
	// boundL1ScrollMessenger is the L1ScrollMessenger contract
	boundL1ScrollMessenger *l1scrollmessenger.L1ScrollMessenger
	// boundL2Gateway is the L2GatewayRouter contract
	boundL2Gateway *l2gateway.L2GatewayRouter
	// l1ETHGatewayListener is the listener for the L1GatewayRouter contract
	l1ETHGatewayListener listener.ContractListener
	// l1ERC20GatewayListener is the listener for the L1GatewayRouter contract
	l1ERC20GatewayListener listener.ContractListener
	// l2ETHGatewayListener is the listener for the L2GatewayRouter contract
	l2ETHGatewayListener listener.ContractListener
	// l2ERC20GatewayListener is the listener for the L2GatewayRouter contract
	l2ERC20GatewayListener listener.ContractListener
	// l1ChainID is the chain ID for the L1 chain
	l1ChainID int
	// l2ChainID is the chain ID for the L2 chain
	l2ChainID int
	// l1ERC20Address is the address of the ERC20 to rebalance on the L1.
	l1ERC20Address common.Address
	// l2ERC20Address is the address of the ERC20 to rebalance on the L2.
	l2ERC20Address common.Address
	// db is the database
	db reldb.Service
	// apiURL is the URL for the scroll API
	apiURL *string
	// httpClient is the client for http requests
	httpClient *http.Client
}

func newRebalanceManagerScroll(cfg relconfig.Config, handler metrics.Handler, chainClient submitter.ClientFetcher, txSubmitter submitter.TransactionSubmitter, relayerAddress common.Address, db reldb.Service) *rebalanceManagerScroll {
	return &rebalanceManagerScroll{
		cfg:            cfg,
		handler:        handler,
		chainClient:    chainClient,
		txSubmitter:    txSubmitter,
		relayerAddress: relayerAddress,
		db:             db,
		httpClient:     &http.Client{},
	}
}

const mainnetChainID = 1
const scrollChainID = 534352
const sepoliaChainID = 11155111
const scrollSepoliaChainID = 534351

func isScrollChain(chainID int) bool {
	return chainID == scrollChainID || chainID == scrollSepoliaChainID
}

func isEthereumChain(chainID int) bool {
	return chainID == mainnetChainID || chainID == sepoliaChainID
}

func isTestnetChain(chainID int) bool {
	return chainID == scrollSepoliaChainID || chainID == sepoliaChainID
}

const claimCheckInterval = 30

func (c *rebalanceManagerScroll) Start(ctx context.Context) (err error) {
	fmt.Println("starting rebalance manager scroll")
	err = c.initContracts(ctx)
	if err != nil {
		fmt.Printf("could not initialize contracts: %v\n", err)
		return fmt.Errorf("could not initialize contracts: %w", err)
	}
	fmt.Println("initialized contracts")

	err = c.initListeners(ctx)
	if err != nil {
		fmt.Printf("init listener err: %v\n", err)
		return fmt.Errorf("could not initialize listeners: %w", err)
	}

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		if !c.txSubmitter.Started() {
			err := c.txSubmitter.Start(ctx)
			if err != nil && !errors.Is(err, submitter.ErrSubmitterAlreadyStarted) {
				return fmt.Errorf("could not start submitter: %w", err)
			}
			return nil
		}
		return nil
	})
	g.Go(func() error {
		return c.listenL1ETHGateway(ctx)
	})
	g.Go(func() error {
		return c.listenL1ERC20Gateway(ctx)
	})
	g.Go(func() error {
		return c.listenL2ETHGateway(ctx)
	})
	g.Go(func() error {
		return c.listenL2ERC20Gateway(ctx)
	})
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(claimCheckInterval * time.Second):
				err := c.claimL2ToL1(ctx)
				if err != nil {
					logger.Warnf("could not claim: %v", err)
				}
			}
		}
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not listen: %w", err)
	}

	return nil
}

const mainnetScrollAPIURL = "https://mainnet-api-bridge-v2.scroll.io/api/l2"
const testnetScrollAPIURL = "https://sepolia-api-bridge-v2.scroll.io/api/l2"
const erc20Name = "USDC"

func (c *rebalanceManagerScroll) initContracts(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initContracts-scroll")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	for chainID := range c.cfg.Chains {
		fmt.Printf("inspecting chain: %d\n", chainID)
		if isEthereumChain(chainID) {
			fmt.Printf("found ethereum chain: %d\n", chainID)
			c.l1ChainID = chainID
			chainClient, err := c.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
			if err != nil {
				return fmt.Errorf("could not get chain client: %w", err)
			}
			addr, err := c.cfg.GetL1GatewayAddress(chainID)
			if err != nil {
				return fmt.Errorf("could not get l1 gateway address: %w", err)
			}
			c.boundL1Gateway, err = l1gateway.NewL1GatewayRouter(common.HexToAddress(addr), chainClient)
			if err != nil {
				return fmt.Errorf("could not get l1 gateway contract: %w", err)
			}
			fmt.Printf("assigned l1 gateway on chain %v at address %v\n", chainID, addr)
			messengerAddr, err := c.cfg.GetL1ScrollMessengerAddress(chainID)
			if err != nil {
				return fmt.Errorf("could not get l1 scroll messenger address: %w", err)
			}
			c.boundL1ScrollMessenger, err = l1scrollmessenger.NewL1ScrollMessenger(common.HexToAddress(messengerAddr), chainClient)
			if err != nil {
				return fmt.Errorf("could not get l1 scroll messenger contract: %w", err)
			}
			span.SetAttributes(
				attribute.String(fmt.Sprintf("l1_gateway_%d", chainID), addr),
				attribute.String(fmt.Sprintf("scroll_messenger_%d", chainID), messengerAddr),
			)
			fmt.Printf("assigned scroll messenger on chain %v at address %v\n", chainID, addr)
		} else if isScrollChain(chainID) {
			fmt.Printf("found scroll chain: %d\n", chainID)
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
			span.SetAttributes(
				attribute.String(fmt.Sprintf("l2_gateway_%d", chainID), addr),
			)
		}
	}
	if c.boundL1Gateway == nil {
		return fmt.Errorf("l1 gateway contract not set")
	}
	if c.boundL1ScrollMessenger == nil {
		return fmt.Errorf("l1 scroll messenger not set")
	}
	if c.boundL2Gateway == nil {
		return fmt.Errorf("l2 gateway contract not set")
	}
	if isTestnetChain(c.l1ChainID) != isTestnetChain(c.l2ChainID) {
		return fmt.Errorf("testnet chain mismatch: %d %d", c.l1ChainID, c.l2ChainID)
	}
	fmt.Println("contracts ok")

	// set ERC20 addresses
	for chainID, chainCfg := range c.cfg.Chains {
		for tokenName, tokenCfg := range chainCfg.Tokens {
			if tokenName != erc20Name {
				continue
			}
			if chainID == c.l1ChainID {
				c.l1ERC20Address = common.HexToAddress(tokenCfg.Address)
			}
			if chainID == c.l2ChainID {
				c.l2ERC20Address = common.HexToAddress(tokenCfg.Address)
			}
		}
	}
	zeroAddress := common.Address{}
	if c.l1ERC20Address == zeroAddress {
		return fmt.Errorf("l1 erc20 address not set")
	}
	if c.l2ERC20Address == zeroAddress {
		return fmt.Errorf("l2 erc20 address not set")
	}
	fmt.Printf("erc20 addresses ok: %v %v\n", c.l1ERC20Address, c.l2ERC20Address)

	// set API URL
	baseURL := mainnetScrollAPIURL
	if isTestnetChain(c.l1ChainID) {
		baseURL = testnetScrollAPIURL
	}
	url := fmt.Sprintf("%s/unclaimed/withdrawals?address=%s", baseURL, c.relayerAddress.Hex())
	c.apiURL = &url

	return nil
}

var zeroAddress = common.Address{}

func (c *rebalanceManagerScroll) initListeners(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initListeners")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	// setup l1 listeners
	l1Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l1ChainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	l1InitialBlock, err := c.cfg.GetCCTPStartBlock(c.l1ChainID)
	if err != nil {
		return fmt.Errorf("could not get cctp start block: %w", err)
	}
	l1ETHAddr, err := c.boundL1Gateway.EthGateway(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get L1ETHGateway address: %w", err)
	}
	c.l1ETHGatewayListener, err = listener.NewChainListener(l1Client, c.db, l1ETHAddr, l1InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get L1ETHGateway listener: %w", err)
	}
	l1ERC20Addr, err := c.boundL1Gateway.GetERC20Gateway(&bind.CallOpts{Context: ctx}, c.l1ERC20Address)
	if err != nil {
		return fmt.Errorf("could not get L1ERC20Gateway address: %w", err)
	}
	fmt.Printf("got l1ERC20Addr %v from token addr %v\n", l1ERC20Addr, c.l1ERC20Address)
	if l1ERC20Addr == zeroAddress {
		return fmt.Errorf("got zero address for L1ERC20Gateway and token address %v", c.l1ERC20Address)
	}
	c.l1ERC20GatewayListener, err = listener.NewChainListener(l1Client, c.db, l1ERC20Addr, l1InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get L1ERC20Gateway listener: %w", err)
	}

	// setup l2 listeners
	l2Client, err := c.chainClient.GetClient(ctx, big.NewInt(int64(c.l2ChainID)))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	l2InitialBlock, err := c.cfg.GetCCTPStartBlock(c.l2ChainID)
	if err != nil {
		return fmt.Errorf("could not get cctp start block: %w", err)
	}
	l2ETHAddr, err := c.boundL2Gateway.EthGateway(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get L2ETHGateway address: %w", err)
	}
	c.l2ETHGatewayListener, err = listener.NewChainListener(l2Client, c.db, l2ETHAddr, l2InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get L2ETHGateway listener: %w", err)
	}
	l2ERC20Addr, err := c.boundL2Gateway.GetERC20Gateway(&bind.CallOpts{Context: ctx}, c.l2ERC20Address)
	if err != nil {
		return fmt.Errorf("could not get L2ERC20Gateway address: %w", err)
	}
	fmt.Printf("got l2ERC20Addr %v from token addr %v\n", l2ERC20Addr, c.l2ERC20Address)
	if l2ERC20Addr == zeroAddress {
		return fmt.Errorf("got zero address for L2ERC20Gateway and token address %v", c.l2ERC20Address)
	}
	c.l2ERC20GatewayListener, err = listener.NewChainListener(l2Client, c.db, l2ERC20Addr, l2InitialBlock, c.handler)
	if err != nil {
		return fmt.Errorf("could not get L2ERC20Gateway listener: %w", err)
	}

	span.SetAttributes(
		attribute.String("l1_eth_gateway", l1ETHAddr.String()),
		attribute.String("l1_erc20_gateway", l1ERC20Addr.String()),
		attribute.String("l2_eth_gateway", l2ETHAddr.String()),
		attribute.String("l2_erc20_gateway", l2ERC20Addr.String()),
	)

	return nil
}

func (c *rebalanceManagerScroll) Execute(ctx context.Context, rebalance *RebalanceData) (err error) {
	switch rebalance.OriginMetadata.ChainID {
	case c.l1ChainID:
		err = c.initiateL1ToL2(ctx, rebalance)
	case c.l2ChainID:
		err = c.initiateL2ToL1(ctx, rebalance)
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
const scrollMsgFee = 1e17

func (c *rebalanceManagerScroll) initiateL1ToL2(parentCtx context.Context, rebalance *RebalanceData) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initiateL1ToL2", trace.WithAttributes(
		attribute.Int(metrics.Origin, rebalance.OriginMetadata.ChainID),
		attribute.Int(metrics.Destination, rebalance.DestMetadata.ChainID),
		attribute.String("origin_token", rebalance.OriginMetadata.Name),
		attribute.String("dest_token", rebalance.OriginMetadata.Name),
		attribute.String("amount", rebalance.Amount.String()),
		attribute.Int("msg_fee", scrollMsgFee),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(rebalance.OriginMetadata.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		transactor.Value = big.NewInt(int64(scrollMsgFee))
		if chain.IsGasToken(rebalance.OriginMetadata.Addr) {
			tx, err = c.boundL1Gateway.DepositETH(transactor, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
			if err != nil {
				return nil, fmt.Errorf("could not deposit gas token: %w", err)
			}
		} else {
			fmt.Printf("calling depositERC20 on chain %v contract %v with token %v, amount %v, gas limit %v\n", rebalance.OriginMetadata.ChainID, c.boundL1Gateway, rebalance.OriginMetadata.Addr, rebalance.Amount, big.NewInt(int64(scrollGasLimit)))
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

func (c *rebalanceManagerScroll) initiateL2ToL1(parentCtx context.Context, rebalance *RebalanceData) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "initiateL2ToL1", trace.WithAttributes(
		attribute.Int(metrics.Origin, rebalance.OriginMetadata.ChainID),
		attribute.Int(metrics.Destination, rebalance.DestMetadata.ChainID),
		attribute.String("origin_token", rebalance.OriginMetadata.Name),
		attribute.String("dest_token", rebalance.OriginMetadata.Name),
		attribute.String("amount", rebalance.Amount.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
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

func (c *rebalanceManagerScroll) listenL1ETHGateway(ctx context.Context) (err error) {
	addr, err := c.boundL1Gateway.EthGateway(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get ETHGateway address: %w", err)
	}
	parser, err := l1gateway.NewParser(addr)
	if err != nil {
		return fmt.Errorf("could not get l1 gateway parser: %w", err)
	}
	err = c.l1ETHGatewayListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		_, parsedEvent, ok := parser.ParseEvent(log)
		if !ok {
			return nil
		}

		switch event := parsedEvent.(type) {
		case *l1gateway.L1GatewayRouterDepositETH:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleDepositETH", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:          uint64(c.l1ChainID),
				Destination:     uint64(c.l2ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: chain.EthAddress,
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l1gateway.L1GatewayRouterFinalizeWithdrawETH:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleFinalizeWithdrawETH", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:      uint64(c.l2ChainID),
				Destination: uint64(c.l1ChainID),
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
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

func (c *rebalanceManagerScroll) listenL1ERC20Gateway(ctx context.Context) (err error) {
	addr, err := c.boundL1Gateway.GetERC20Gateway(&bind.CallOpts{Context: ctx}, c.l1ERC20Address)
	if err != nil {
		return fmt.Errorf("could not get ERC20Gateway address: %w", err)
	}
	parser, err := l1gateway.NewParser(addr)
	if err != nil {
		return fmt.Errorf("could not get l1 gateway parser: %w", err)
	}
	err = c.l1ERC20GatewayListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		_, parsedEvent, ok := parser.ParseEvent(log)
		if !ok {
			return nil
		}

		switch event := parsedEvent.(type) {
		case *l1gateway.L1GatewayRouterDepositERC20:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleDepositERC20", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:          uint64(c.l1ChainID),
				Destination:     uint64(c.l2ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: event.L1Token,
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l1gateway.L1GatewayRouterFinalizeWithdrawERC20:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleFinalizeWithdrawERC20", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:      uint64(c.l2ChainID),
				Destination: uint64(c.l1ChainID),
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
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

func (c *rebalanceManagerScroll) listenL2ETHGateway(ctx context.Context) (err error) {
	addr, err := c.boundL2Gateway.EthGateway(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get L2ETHGateway address: %w", err)
	}
	parser, err := l2gateway.NewParser(addr)
	if err != nil {
		return fmt.Errorf("could not get l2 gateway parser: %w", err)
	}
	err = c.l2ETHGatewayListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		_, parsedEvent, ok := parser.ParseEvent(log)
		if !ok {
			return nil
		}

		switch event := parsedEvent.(type) {
		case *l2gateway.L2GatewayRouterWithdrawETH:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleWithdrawETH", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:          uint64(c.l2ChainID),
				Destination:     uint64(c.l1ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: chain.EthAddress,
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l2gateway.L2GatewayRouterFinalizeDepositETH:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleFinalizeDepositETH", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:      uint64(c.l1ChainID),
				Destination: uint64(c.l2ChainID),
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
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

func (c *rebalanceManagerScroll) listenL2ERC20Gateway(ctx context.Context) (err error) {
	addr, err := c.boundL2Gateway.GetERC20Gateway(&bind.CallOpts{Context: ctx}, c.l2ERC20Address)
	if err != nil {
		return fmt.Errorf("could not get L2ERC20Gateway address: %w", err)
	}
	parser, err := l2gateway.NewParser(addr)
	if err != nil {
		return fmt.Errorf("could not get l2 gateway parser: %w", err)
	}
	err = c.l2ERC20GatewayListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		_, parsedEvent, ok := parser.ParseEvent(log)
		if !ok {
			return nil
		}

		switch event := parsedEvent.(type) {
		case *l2gateway.L2GatewayRouterWithdrawERC20:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleWithdrawERC20", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:          uint64(c.l2ChainID),
				Destination:     uint64(c.l1ChainID),
				OriginTxHash:    log.TxHash,
				OriginTokenAddr: event.L2Token,
				Status:          reldb.RebalancePending,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
			if err != nil {
				logger.Warnf("could not update rebalance status: %v", err)
				return nil
			}
		case *l2gateway.L2GatewayRouterFinalizeDepositERC20:
			if event.To != c.relayerAddress || event.From != c.relayerAddress {
				return
			}

			ctx, span := c.handler.Tracer().Start(parentCtx, "handleFinalizeDepositERC20", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
				attribute.String(metrics.Contract, log.Address.String()),
				attribute.String("block_hash", log.BlockHash.String()),
				attribute.Int64("block_number", int64(log.BlockNumber)),
			))
			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			rebalanceModel := reldb.Rebalance{
				Origin:      uint64(c.l2ChainID),
				Destination: uint64(c.l1ChainID),
				DestTxHash:  log.TxHash,
				Status:      reldb.RebalanceCompleted,
			}
			err = c.db.UpdateLatestRebalance(ctx, rebalanceModel)
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

type scrollAPIResponse struct {
	Data struct {
		Results []struct {
			ClaimInfo ClaimInfo `json:"claim_info"`
		} `json:"results"`
	} `json:"data"`
}

type ClaimInfo struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Value      string `json:"value"`
	Nonce      string `json:"nonce"`
	BatchHash  string `json:"batch_hash"`
	Message    string `json:"message"`
	Proof      string `json:"proof"`
	BatchIndex string `json:"batch_index"`
}

func (c *rebalanceManagerScroll) claimL2ToL1(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "claimL2ToL1")
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	if c.apiURL == nil {
		return fmt.Errorf("api URL not set")
	}
	span.SetAttributes(attribute.String("api_url", *c.apiURL))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, *c.apiURL, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not get response: %w", err)
	}
	//nolint:errcheck
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var claimableResp scrollAPIResponse
	err = json.Unmarshal(body, &claimableResp)
	if err != nil {
		return fmt.Errorf("could not unmarshal body: %w", err)
	}

	for _, result := range claimableResp.Data.Results {
		err = c.submitClaim(ctx, result.ClaimInfo)
		if err != nil {
			return fmt.Errorf("could not submit transaction: %w", err)
		}
	}
	return nil
}

func (c *rebalanceManagerScroll) submitClaim(parentCtx context.Context, claimInfo ClaimInfo) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "submitClaim", trace.WithAttributes(
		attribute.String("from", claimInfo.From),
		attribute.String("to", claimInfo.To),
		attribute.String("value", claimInfo.Value),
		attribute.String("nonce", claimInfo.Nonce),
		attribute.String("batch_index", claimInfo.BatchIndex),
	))
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(c.l1ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		if transactor == nil {
			return nil, fmt.Errorf("transactor is nil")
		}
		// Note: we hardcode the 'to' parameter as our own relayerAddress as a safety measure.
		value, ok := new(big.Int).SetString(claimInfo.Value, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse value: %w", err)
		}
		nonce, ok := new(big.Int).SetString(claimInfo.Nonce, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse nonce: %w", err)
		}
		batchIndex, ok := new(big.Int).SetString(claimInfo.BatchIndex, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse batch index: %w", err)
		}
		message, err := hexutil.Decode(claimInfo.Message)
		if err != nil {
			return nil, fmt.Errorf("could not decode message: %w", err)
		}
		merkleProof, err := hexutil.Decode(claimInfo.Proof)
		if err != nil {
			return nil, fmt.Errorf("could not decode merkle proof: %w", err)
		}
		proof := l1scrollmessenger.IL1ScrollMessengerL2MessageProof{
			BatchIndex:  batchIndex,
			MerkleProof: merkleProof,
		}
		tx, err = c.boundL1ScrollMessenger.RelayMessageWithProof(transactor, common.HexToAddress(claimInfo.From), c.relayerAddress, value, nonce, message, proof)
		if err != nil {
			return nil, fmt.Errorf("could not relay message: %w", err)
		}
		return tx, nil
	})
	return nil
}
