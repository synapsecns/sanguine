package inventory

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l1gateway"
	"github.com/synapsecns/sanguine/services/rfq/contracts/l2gateway"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// Manager is the interface for the inventory manager.
//
//go:generate go run github.com/vektra/mockery/v2 --name Manager --output ./mocks --case=underscore
type Manager interface {
	// Start starts the inventory manager.
	Start(ctx context.Context) (err error)
	// GetCommittableBalance gets the total balance available for quotes
	// this does not include on-chain balances committed in previous quotes that may be
	// refunded in the event of a revert.
	GetCommittableBalance(ctx context.Context, chainID int, token common.Address, options ...BalanceFetchArgOption) (*big.Int, error)
	// GetCommittableBalances gets the total balances committable for all tracked tokens.
	GetCommittableBalances(ctx context.Context, options ...BalanceFetchArgOption) (map[int]map[common.Address]*big.Int, error)
	// ApproveAllTokens approves all tokens for the relayer address.
	ApproveAllTokens(ctx context.Context) error
	// HasSufficientGas checks if there is sufficient gas for a given route.
	HasSufficientGas(ctx context.Context, chainID int, gasValue *big.Int) (bool, error)
	// Rebalance checks whether a given token should be rebalanced, and
	// executes the rebalance if necessary.
	Rebalance(ctx context.Context, chainID int, token common.Address) error
	// GetTokenMetadata gets the metadata for a token.
	GetTokenMetadata(chainID int, token common.Address) (*TokenMetadata, error)
}

type inventoryManagerImpl struct {
	// map chainID->address->TokenMetadata
	tokens map[int]map[common.Address]*TokenMetadata
	// map chainID->balance
	gasBalances map[int]*big.Int
	// mux contains the mutex
	mux sync.RWMutex
	// handler is the metrics handler
	handler metrics.Handler
	// cfg is the config
	cfg relconfig.Config
	// relayerAddress contains the relayer address
	relayerAddress common.Address
	// chainClient is an omnirpc client
	chainClient submitter.ClientFetcher
	// txSubmitter is the transaction submitter
	txSubmitter submitter.TransactionSubmitter
	// rebalanceManagers is the map of rebalance managers
	rebalanceManagers map[relconfig.RebalanceMethod]RebalanceManager
	// db is the database
	db reldb.Service
	// meter is the metrics meter for this package
	meter metric.Meter
	// balanceGauge is the histogram for balance
	balanceGauge metric.Float64ObservableGauge
	// inFlightQuoteManager is the cache for in flight quotes
	inFlightQuoteManager *inFlightManager
}

// ErrUnsupportedChain is the error for an unsupported chain.
var ErrUnsupportedChain = errors.New("could not get gas balance for unsupported chain")

// GetCommittableBalance gets the committable balances.
func (i *inventoryManagerImpl) GetCommittableBalance(ctx context.Context, chainID int, token common.Address, options ...BalanceFetchArgOption) (*big.Int, error) {
	committableBalances, err := i.GetCommittableBalances(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("could not get balances: %w", err)
	}
	balance := committableBalances[chainID][token]
	// the gas token may not be registered in the inventory tokens map,
	// but it is always tracked in gasBalances.
	if balance == nil && token == util.EthAddress {
		gasBalance, ok := i.gasBalances[chainID]
		if !ok || gasBalance == nil {
			return nil, ErrUnsupportedChain
		}
		balance = i.gasBalances[chainID]
	}
	return balance, nil
}

func (i *inventoryManagerImpl) GetCommittableBalances(ctx context.Context, options ...BalanceFetchArgOption) (res map[int]map[common.Address]*big.Int, err error) {
	reqOptions := makeOptions(options)
	// TODO: hard fail if cache skip breaks
	if reqOptions.shouldRefreshBalances {
		// TODO; no need for this if refresh already in flight
		_ = i.refreshBalances(ctx)
	}

	// get db first
	// Add other committed, but incomplete statuses here
	// TODO: clean me up: you can do this by having a IsLiquidityCommitted() method on the type.
	inFlightQuotes, err := i.inFlightQuoteManager.GetInFlightQuotes(ctx, reqOptions.skipDBCache)
	if err != nil {
		return nil, fmt.Errorf("could not get in flight quotes: %w", err)
	}

	// TODO: lock should be context aware
	i.mux.RLock()
	defer i.mux.RUnlock()
	res = make(map[int]map[common.Address]*big.Int)
	for chainID, tokenMap := range i.tokens {
		res[chainID] = map[common.Address]*big.Int{}
		for address, tokenData := range tokenMap {
			res[chainID][address] = core.CopyBigInt(tokenData.Balance)
			// now subtract by in flight quotes.
			// Yeah, this is an algorithmically atrocious for
			// TODO: fix, but we're really talking about 4 tokens
			for _, quote := range inFlightQuotes {
				if quote.Transaction.DestToken == address && quote.Transaction.DestChainId == uint32(chainID) {
					res[chainID][address] = new(big.Int).Sub(res[chainID][address], quote.Transaction.DestAmount)
				}
			}
		}
	}

	// TODO: db subtraction

	return res, nil
}

// TokenMetadata contains metadata for a token.
type TokenMetadata struct {
	// Name is the name of the token in the config
	Name string
	// ChainName is the name of the token onchain
	ChainName  string
	Balance    *big.Int
	Decimals   uint8
	IsGasToken bool
	ChainID    int
	Addr       common.Address
}

var (
	funcBalanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")
	funcName      = w3.MustNewFunc("name()", "string")
	funcDecimals  = w3.MustNewFunc("decimals()", "uint8")
)

// TODO: replace w/ config.
const defaultPollPeriod = 5
const meterName = "github.com/synapsecns/sanguine/services/rfq/relayer/inventory"

// NewInventoryManager creates a new inventory manager.
// TODO: too many args here.
//
//nolint:gocognit
func NewInventoryManager(ctx context.Context, clientFetcher submitter.ClientFetcher, handler metrics.Handler, cfg relconfig.Config, relayer common.Address, txSubmitter submitter.TransactionSubmitter, db reldb.Service) (Manager, error) {
	rebalanceMethods, err := cfg.GetAllRebalanceMethods()
	if err != nil {
		return nil, fmt.Errorf("could not get rebalance methods: %w", err)
	}
	rebalanceManagers := make(map[relconfig.RebalanceMethod]RebalanceManager)
	for method := range rebalanceMethods {
		//nolint:exhaustive
		switch method {
		case relconfig.RebalanceMethodSynapseCCTP:
			rebalanceManagers[method] = newRebalanceManagerSynapseCCTP(cfg, handler, clientFetcher, txSubmitter, relayer, db)
		case relconfig.RebalanceMethodCircleCCTP:
			rebalanceManagers[method] = newRebalanceManagerCircleCCTP(cfg, handler, clientFetcher, txSubmitter, relayer, db)
		case relconfig.RebalanceMethodScroll:
			rebalanceManagers[method] = newRebalanceManagerScroll(cfg, handler, clientFetcher, txSubmitter, relayer, db)
		default:
			return nil, fmt.Errorf("unsupported rebalance method: %s", method)
		}
	}

	i := inventoryManagerImpl{
		relayerAddress:       relayer,
		handler:              handler,
		cfg:                  cfg,
		chainClient:          clientFetcher,
		txSubmitter:          txSubmitter,
		rebalanceManagers:    rebalanceManagers,
		db:                   db,
		meter:                handler.Meter(meterName),
		inFlightQuoteManager: newInflightManager(db),
	}

	i.balanceGauge, err = i.meter.Float64ObservableGauge("inventory_balance")
	if err != nil {
		return nil, fmt.Errorf("could not create balance gauge: %w", err)
	}

	_, err = i.meter.RegisterCallback(i.recordBalances, i.balanceGauge)
	if err != nil {
		return nil, fmt.Errorf("could not register callback: %w", err)
	}

	err = i.initializeTokens(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not initialize tokens: %w", err)
	}

	return &i, nil
}

//nolint:gocognit,cyclop
func (i *inventoryManagerImpl) Start(ctx context.Context) error {
	g, gctx := errgroup.WithContext(ctx)
	for _, rebalanceManager := range i.rebalanceManagers {
		rebalanceManager := rebalanceManager
		g.Go(func() error {
			err := rebalanceManager.Start(gctx)
			if err != nil {
				return fmt.Errorf("could not start rebalance manager: %w", err)
			}
			return nil
		})
	}

	// continuously refresh balances
	g.Go(func() error {
		for {
			select {
			case <-gctx.Done():
				return fmt.Errorf("context canceled: %w", gctx.Err())
			case <-time.After(250 * time.Millisecond):
				// this returning an error isn't really possible unless a config error happens
				// TODO: need better error handling.
				err := i.refreshBalances(gctx)
				if err != nil {
					logger.Errorf("could not refresh balances")
					//nolint:nilerr
					return nil
				}
			}
		}
	})

	// continuously check for rebalances
	rebalanceInterval := i.cfg.GetRebalanceInterval()
	if rebalanceInterval > 0 {
		g.Go(func() error {
			for {
				select {
				case <-gctx.Done():
					return fmt.Errorf("context canceled: %w", gctx.Err())
				case <-time.After(rebalanceInterval):
					err := i.refreshBalances(gctx)
					if err != nil {
						return fmt.Errorf("could not refresh balances: %w", err)
					}
					for chainID, chainConfig := range i.cfg.Chains {
						for tokenName, tokenConfig := range chainConfig.Tokens {
							err = i.Rebalance(gctx, chainID, common.HexToAddress(tokenConfig.Address))
							if err != nil {
								logger.Errorf("could not rebalance %s on chain %d: %v", tokenName, chainID, err)
							}
						}
					}
				}
			}
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("error starting inventory manager: %w", err)
	}
	return nil
}

const maxBatchSize = 10

// ApproveAllTokens approves all checks if allowance is set and if not approves.
// nolint:gocognit,nestif,cyclop
func (i *inventoryManagerImpl) ApproveAllTokens(ctx context.Context) error {
	i.mux.RLock()
	defer i.mux.RUnlock()

	for chainID, tokenMap := range i.tokens {
		backendClient, err := i.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		for tokenAddr := range tokenMap {
			// Note: in the case where submitter hasn't finished from last boot,
			// this will double submit approvals unfortunately.
			contractAddr, addrErr := i.cfg.GetRFQAddress(chainID)
			if addrErr == nil {
				err = i.approve(ctx, tokenAddr, contractAddr, backendClient)
				if err != nil {
					return fmt.Errorf("could not approve RFQ contract: %w", err)
				}
			}

			contractAddr, addrErr = i.cfg.GetSynapseCCTPAddress(chainID)
			if addrErr == nil {
				err = i.approve(ctx, tokenAddr, contractAddr, backendClient)
				if err != nil {
					return fmt.Errorf("could not approve SynapseCCTP contract: %w", err)
				}
			}

			contractAddr, addrErr = i.cfg.GetTokenMessengerAddress(chainID)
			if addrErr == nil {
				err = i.approve(ctx, tokenAddr, contractAddr, backendClient)
				if err != nil {
					return fmt.Errorf("could not approve TokenMessenger contract: %w", err)
				}
			}

			parentAddr, addrErr := i.cfg.GetL1GatewayAddress(chainID)
			if addrErr == nil {
				contract, err := l1gateway.NewL1GatewayRouter(parentAddr, backendClient)
				if err != nil {
					return fmt.Errorf("could not get L1Gateway contract: %w", err)
				}
				contractAddr, err = contract.ERC20Gateway(&bind.CallOpts{Context: ctx}, tokenAddr)
				if err != nil {
					return fmt.Errorf("could not get L1ERC20Gateway address: %w", err)
				}
				err = i.approve(ctx, tokenAddr, contractAddr, backendClient)
				if err != nil {
					return fmt.Errorf("could not approve L1ERC20Gateway contract: %w", err)
				}
			}

			parentAddr, addrErr = i.cfg.GetL2GatewayAddress(chainID)
			if addrErr == nil {
				contract, err := l2gateway.NewL2GatewayRouter(parentAddr, backendClient)
				if err != nil {
					return fmt.Errorf("could not get L2Gateway contract: %w", err)
				}
				contractAddr, err = contract.ERC20Gateway(&bind.CallOpts{Context: ctx}, tokenAddr)
				if err != nil {
					return fmt.Errorf("could not get L2ERC20Gateway address: %w", err)
				}
				err = i.approve(ctx, tokenAddr, contractAddr, backendClient)
				if err != nil {
					return fmt.Errorf("could not approve L2ERC20Gateway contract: %w", err)
				}
			}
		}
	}
	return nil
}

// approve submits an ERC20 approval for a given token and contract address.
func (i *inventoryManagerImpl) approve(parentCtx context.Context, tokenAddr, contractAddr common.Address, backendClient client.EVM) (err error) {
	ctx, span := i.handler.Tracer().Start(parentCtx, "approve", trace.WithAttributes(
		attribute.String("token_address", tokenAddr.Hex()),
		attribute.String("contract_address", contractAddr.Hex()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	if tokenAddr == util.EthAddress {
		span.AddEvent("not approving with eth address")
		return nil
	}

	if contractAddr == (common.Address{}) {
		span.AddEvent("not approving to zero address")
		return nil
	}

	erc20, err := ierc20.NewIERC20(tokenAddr, backendClient)
	if err != nil {
		return fmt.Errorf("could not get erc20: %w", err)
	}

	allowance, err := erc20.Allowance(&bind.CallOpts{Context: ctx}, i.relayerAddress, contractAddr)
	if err != nil {
		return fmt.Errorf("could not get allowance: %w", err)
	}
	if allowance.Cmp(big.NewInt(0)) > 0 {
		span.AddEvent("already has positive allowance")
		return nil
	}

	chainID, err := backendClient.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("could not get chain id: %w", err)
	}

	_, err = i.txSubmitter.SubmitTransaction(ctx, chainID, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = erc20.Approve(transactor, contractAddr, abi.MaxUint256)
		if err != nil {
			return nil, fmt.Errorf("could not approve: %w", err)
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit approval: %w", err)
	}
	return nil
}

// HasSufficientGas checks if there is sufficient gas for a given route.
func (i *inventoryManagerImpl) HasSufficientGas(parentCtx context.Context, chainID int, gasValue *big.Int) (sufficient bool, err error) {
	ctx, span := i.handler.Tracer().Start(parentCtx, "HasSufficientGas", trace.WithAttributes(
		attribute.Int(metrics.ChainID, chainID),
	))
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	gasThreshRaw, err := i.cfg.GetMinGasToken(chainID)
	if err != nil {
		return false, fmt.Errorf("error getting min gas token on origin: %w", err)
	}
	gasThresh := core.CopyBigInt(gasThreshRaw)
	if gasValue != nil {
		gasThresh = new(big.Int).Add(gasThresh, gasValue)
		span.SetAttributes(attribute.String("gas_value", gasValue.String()))
	}

	gasBalance, err := i.GetCommittableBalance(ctx, chainID, util.EthAddress)
	if err != nil {
		return false, fmt.Errorf("error getting committable gas on origin: %w", err)
	}

	sufficient = gasBalance.Cmp(gasThresh) >= 0
	span.SetAttributes(
		attribute.String("gas_threshold_raw", gasThreshRaw.String()),
		attribute.String("gas_threshold", gasThresh.String()),
		attribute.String("gas_balance", gasBalance.String()),
		attribute.Bool("sufficient", sufficient),
	)
	return sufficient, nil
}

// Rebalance checks whether a given token should be rebalanced, and executes the rebalance if necessary.
// Note that if there are multiple tokens whose balance is below the maintenance balance, only the lowest balance
// will be rebalanced.
//
//nolint:cyclop
func (i *inventoryManagerImpl) Rebalance(parentCtx context.Context, chainID int, token common.Address) (err error) {
	// short circuit if origin does not specify a rebalance method
	methodsOrigin, err := i.cfg.GetRebalanceMethods(chainID, token.Hex())
	if err != nil {
		return fmt.Errorf("could not get origin rebalance method: %w", err)
	}
	if len(methodsOrigin) == 0 {
		return nil
	}

	ctx, span := i.handler.Tracer().Start(parentCtx, "Rebalance", trace.WithAttributes(
		attribute.Int(metrics.ChainID, chainID),
		attribute.String("token", token.Hex()),
	))
	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	// build the rebalance action
	rebalance, err := getRebalance(span, i.cfg, i.tokens, chainID, token)
	if err != nil {
		return fmt.Errorf("could not get rebalance: %w", err)
	}
	if rebalance == nil || rebalance.Amount.Cmp(big.NewInt(0)) <= 0 {
		return nil
	}
	span.SetAttributes(
		attribute.String("rebalance_origin", strconv.Itoa(rebalance.OriginMetadata.ChainID)),
		attribute.String("rebalance_dest", strconv.Itoa(rebalance.DestMetadata.ChainID)),
		attribute.String("rebalance_amount", rebalance.Amount.String()),
		attribute.String("rebalance_method", rebalance.Method.String()),
	)

	// make sure there are no pending rebalances that touch the given path
	pendingRebalances, err := i.db.GetPendingRebalances(ctx, uint64(rebalance.OriginMetadata.ChainID), uint64(rebalance.DestMetadata.ChainID))
	if err != nil {
		return fmt.Errorf("could not check pending rebalance: %w", err)
	}
	pending := len(pendingRebalances) > 0
	span.SetAttributes(attribute.Bool("rebalance_pending", pending))
	if pending {
		return nil
	}

	// execute the rebalance
	manager, ok := i.rebalanceManagers[rebalance.Method]
	if !ok {
		return fmt.Errorf("no rebalance manager for method: %s", rebalance.Method)
	}
	span.AddEvent("executing")
	err = manager.Execute(ctx, rebalance)
	if err != nil {
		return fmt.Errorf("could not execute rebalance: %w", err)
	}
	return nil
}

func (i *inventoryManagerImpl) GetTokenMetadata(chainID int, token common.Address) (*TokenMetadata, error) {
	i.mux.RLock()
	defer i.mux.RUnlock()
	tokenData, ok := i.tokens[chainID][token]
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	return tokenData, nil
}

// initializeTokens converts the configuration into a data structure we can use to determine inventory
// it gets metadata like name, decimals, etc once and exports these to prometheus for ease of debugging.
func (i *inventoryManagerImpl) initializeTokens(parentCtx context.Context, cfg relconfig.Config) (err error) {
	i.mux.Lock()
	defer i.mux.Unlock()

	ctx, span := i.handler.Tracer().Start(parentCtx, "initializeTokens", trace.WithAttributes(
		attribute.String("relayer_address", i.relayerAddress.String()),
	))

	defer func(err error) {
		metrics.EndSpanWithErr(span, err)
	}(err)

	// TODO: this needs to be a struct bound variable otherwise will be stuck.
	i.tokens = make(map[int]map[common.Address]*TokenMetadata)
	i.gasBalances = make(map[int]*big.Int)

	// TODO: this can be pre-capped w/ len(cfg.Tokens) for each chain id.
	// here we register metrics for exporting through otel. We wait to call these functions until are tokens have been initialized to avoid nil issues.
	deferredCalls := make(map[int][]w3types.Caller)

	// iterate through all tokens to get the metadata
	for cid, chainCfg := range cfg.GetChains() {
		chainID := cid // capture func literal
		i.tokens[chainID] = map[common.Address]*TokenMetadata{}

		// set up balance fetching for this chain's gas token
		i.gasBalances[chainID] = new(big.Int)
		deferredCalls[chainID] = append(deferredCalls[chainID],
			eth.Balance(i.relayerAddress, nil).Returns(i.gasBalances[chainID]),
		)

		// assign metadata for each configured token
		for tokenName, tokenCfg := range chainCfg.Tokens {
			nativeToken, err := cfg.GetNativeToken(chainID)
			if err != nil {
				return fmt.Errorf("could not get native token: %w", err)
			}
			rtoken := &TokenMetadata{
				Name:       tokenName,
				IsGasToken: tokenName == nativeToken,
				ChainID:    chainID,
			}

			var token common.Address
			if rtoken.IsGasToken {
				token = util.EthAddress
			} else {
				token = common.HexToAddress(tokenCfg.Address)
			}
			i.tokens[chainID][token] = rtoken
			rtoken.Addr = token

			// requires non-nil pointer
			rtoken.Balance = new(big.Int)

			if rtoken.IsGasToken {
				rtoken.Decimals = 18
				rtoken.Balance = i.gasBalances[chainID]
				rtoken.ChainName = tokenName
			} else {
				deferredCalls[chainID] = append(deferredCalls[chainID],
					eth.CallFunc(funcBalanceOf, token, i.relayerAddress).Returns(rtoken.Balance),
					eth.CallFunc(funcDecimals, token).Returns(&rtoken.Decimals),
					eth.CallFunc(funcName, token).Returns(&rtoken.ChainName),
				)
			}
		}
	}

	// run through the deferred cals
	g, gctx := errgroup.WithContext(ctx)
	for chainID := range deferredCalls {
		chainID := chainID // capture func literal

		chainClient, err := i.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("can't initialize tokens, no chain client available for chain %d: %w", chainID, err)
		}

		g.Go(func() error {
			// TODO: add retries
			// TODO: we should see if we can move this to ethergo and deduplicate. We do this a lot, especially in
			// the prom exporter
			batches := core.ChunkSlice(deferredCalls[chainID], maxBatchSize)
			for _, batch := range batches {
				err = chainClient.BatchWithContext(gctx, batch...)
				if err != nil {
					return fmt.Errorf("could not batch: %w", err)
				}
			}
			return nil
		})
	}

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not get tx: %w", err)
	}

	return nil
}

var logger = log.Logger("inventory")

// refreshBalances refreshes all the token balances.
func (i *inventoryManagerImpl) refreshBalances(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(i.tokens))

	// TODO: this can be pre-capped w/ len(cfg.Tokens) for each chain id.
	// here we register metrics for exporting through otel. We wait to call these functions until are tokens have been initialized to avoid nil issues.
	for cid, tokenMap := range i.tokens {
		chainID := cid // capture func literal
		chainClient, err := i.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		// queue gas token balance fetch
		deferredCalls := []w3types.Caller{
			eth.Balance(i.relayerAddress, nil).Returns(i.gasBalances[chainID]),
		}

		// queue token balance fetches
		for ta, token := range tokenMap {
			tokenAddress := ta // capture func literal
			// TODO: make sure Returns does nothing on error
			if !token.IsGasToken {
				deferredCalls = append(deferredCalls, eth.CallFunc(funcBalanceOf, tokenAddress, i.relayerAddress).Returns(token.Balance))
			}
		}

		go func() {
			defer wg.Done()
			err = chainClient.BatchWithContext(ctx, deferredCalls...)
			if err != nil {
				logger.Warnf("could not refresh balances on %d: %v", chainID, err)
			}
		}()
	}
	wg.Wait()

	return nil
}

func (i *inventoryManagerImpl) recordBalances(ctx context.Context, observer metric.Observer) (err error) {
	if i.meter == nil || i.balanceGauge == nil {
		return nil
	}

	i.mux.RLock()
	defer i.mux.RUnlock()

	for chainID, tokens := range i.tokens {
		for token, tokenData := range tokens {
			opts := metric.WithAttributes(
				attribute.Int(metrics.ChainID, chainID),
				attribute.String("relayer_address", i.relayerAddress.String()),
				attribute.String("token_name", tokenData.Name),
				attribute.Int("decimals", int(tokenData.Decimals)),
				attribute.String("token_address", token.String()),
				attribute.String("raw_balance", tokenData.Balance.String()),
				attribute.String("relayer", i.relayerAddress.Hex()),
			)

			// Convert the balance and record it
			decimalBalance := core.BigToDecimals(tokenData.Balance, tokenData.Decimals)
			observer.ObserveFloat64(i.balanceGauge, decimalBalance, opts)
		}
	}

	return nil
}
