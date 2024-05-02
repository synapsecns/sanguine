package inventory

import (
	"context"
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
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
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
	// balanceHist is the histogram for balance
	balanceHist metric.Float64Histogram
	// pendingHist is the histogram for pending rebalances
	pendingHist metric.Float64Histogram
}

// GetCommittableBalance gets the committable balances.
func (i *inventoryManagerImpl) GetCommittableBalance(ctx context.Context, chainID int, token common.Address, options ...BalanceFetchArgOption) (*big.Int, error) {
	committableBalances, err := i.GetCommittableBalances(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("could not get balances: %w", err)
	}
	balance := committableBalances[chainID][token]
	// the gas token may not be registered in the inventory tokens map,
	// but it is always tracked in gasBalances.
	if balance == nil && token == chain.EthAddress {
		gasBalance, ok := i.gasBalances[chainID]
		if !ok || gasBalance == nil {
			return nil, fmt.Errorf("could not get gas balance for chain %d", chainID)
		}
		balance = i.gasBalances[chainID]
	}
	return balance, nil
}

func (i *inventoryManagerImpl) GetCommittableBalances(ctx context.Context, options ...BalanceFetchArgOption) (res map[int]map[common.Address]*big.Int, err error) {
	reqOptions := makeOptions(options)
	// TODO: hard fail if cache skip breaks
	if reqOptions.skipCache {
		// TODO; no need for this if refresh already in flight
		_ = i.refreshBalances(ctx)
	}
	// get db first
	// Add other committed, but incomplete statuses here
	// TODO: clean me up: you can do this by having a IsLiquidityCommitted() method on the type.
	inFlightQuotes, err := i.db.GetQuoteResultsByStatus(ctx, reldb.CommittedPending, reldb.CommittedConfirmed, reldb.RelayStarted)
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
	Name       string
	Balance    *big.Int
	Decimals   uint8
	Allowances map[spendableContract]*big.Int
	IsGasToken bool
	ChainID    int
	Addr       common.Address
}

type spendableContract int

const (
	contractRFQ = iota + 1
	contractSynapseCCTP
	contractTokenMessenger
)

var (
	funcBalanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")
	funcName      = w3.MustNewFunc("name()", "string")
	funcDecimals  = w3.MustNewFunc("decimals()", "uint8")
	funcAllowance = w3.MustNewFunc("allowance(address,address)", "uint256")
)

// TODO: replace w/ config.
const defaultPollPeriod = 5
const meterName = "github.com/synapsecns/sanguine/services/rfq/relayer/inventory"

// NewInventoryManager creates a new inventory manager.
// TODO: too many args here.
//
//nolint:gocognit
func NewInventoryManager(ctx context.Context, clientFetcher submitter.ClientFetcher, handler metrics.Handler, cfg relconfig.Config, relayer common.Address, txSubmitter submitter.TransactionSubmitter, db reldb.Service) (Manager, error) {
	rebalanceMethods, err := cfg.GetRebalanceMethods()
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
		default:
			return nil, fmt.Errorf("unsupported rebalance method: %s", method)
		}
	}

	meter := handler.Meter(meterName)
	balanceHist, err := meter.Float64Histogram("inventory_balance")
	if err != nil {
		return nil, fmt.Errorf("could not create balance histogram: %w", err)
	}
	pendingHist, err := meter.Float64Histogram("pending_rebalance_amount")
	if err != nil {
		return nil, fmt.Errorf("could not create pending rebalance histogram: %w", err)
	}

	i := inventoryManagerImpl{
		relayerAddress:    relayer,
		handler:           handler,
		cfg:               cfg,
		chainClient:       clientFetcher,
		txSubmitter:       txSubmitter,
		rebalanceManagers: rebalanceManagers,
		db:                db,
		meter:             meter,
		balanceHist:       balanceHist,
		pendingHist:       pendingHist,
	}

	err = i.initializeTokens(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not initialize tokens: %w", err)
	}

	return &i, nil
}

//nolint:gocognit,cyclop
func (i *inventoryManagerImpl) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)
	for _, rebalanceManager := range i.rebalanceManagers {
		rebalanceManager := rebalanceManager
		g.Go(func() error {
			err := rebalanceManager.Start(ctx)
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
			case <-ctx.Done():
				return fmt.Errorf("context canceled: %w", ctx.Err())
			case <-time.After(250 * time.Millisecond):
				// this returning an error isn't really possible unless a config error happens
				// TODO: need better error handling.
				err := i.refreshBalances(ctx)
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
				case <-ctx.Done():
					return fmt.Errorf("context canceled: %w", ctx.Err())
				case <-time.After(rebalanceInterval):
					err := i.refreshBalances(ctx)
					if err != nil {
						return fmt.Errorf("could not refresh balances: %w", err)
					}
					for chainID, chainConfig := range i.cfg.Chains {
						for tokenName, tokenConfig := range chainConfig.Tokens {
							err = i.Rebalance(ctx, chainID, common.HexToAddress(tokenConfig.Address))
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

		for address, token := range tokenMap {
			// approve RFQ contract.
			// Note: in the case where submitter hasn't finished from last boot,
			// this will double submit approvals unfortunately.
			if address != chain.EthAddress && token.Allowances[contractRFQ].Cmp(big.NewInt(0)) == 0 {
				tokenAddr := address // capture func literal
				contractAddr, err := i.cfg.GetRFQAddress(chainID)
				if err != nil {
					return fmt.Errorf("could not get RFQ address: %w", err)
				}
				err = i.approve(ctx, tokenAddr, common.HexToAddress(contractAddr), backendClient)
				if err != nil {
					return fmt.Errorf("could not approve RFQ contract: %w", err)
				}
			}

			// approve SynapseCCTP contract
			if address != chain.EthAddress && token.Allowances[contractSynapseCCTP].Cmp(big.NewInt(0)) == 0 {
				tokenAddr := address // capture func literal
				contractAddr, err := i.cfg.GetSynapseCCTPAddress(chainID)
				if err != nil {
					return fmt.Errorf("could not get CCTP address: %w", err)
				}
				err = i.approve(ctx, tokenAddr, common.HexToAddress(contractAddr), backendClient)
				if err != nil {
					return fmt.Errorf("could not approve SynapseCCTP contract: %w", err)
				}
			}

			// approve TokenMessenger contract
			if address != chain.EthAddress && token.Allowances[contractTokenMessenger].Cmp(big.NewInt(0)) == 0 {
				tokenAddr := address // capture func literal
				contractAddr, err := i.cfg.GetTokenMessengerAddress(chainID)
				if err != nil {
					return fmt.Errorf("could not get CCTP address: %w", err)
				}
				err = i.approve(ctx, tokenAddr, common.HexToAddress(contractAddr), backendClient)
				if err != nil {
					return fmt.Errorf("could not approve TokenMessenger contract: %w", err)
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

	if contractAddr == (common.Address{}) {
		span.AddEvent("not approving to zero address")
		return nil
	}

	erc20, err := ierc20.NewIERC20(tokenAddr, backendClient)
	if err != nil {
		return fmt.Errorf("could not get erc20: %w", err)
	}
	chainID, err := backendClient.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("could not get chain id: %w", err)
	}

	_, err = i.txSubmitter.SubmitTransaction(ctx, chainID, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = erc20.Approve(transactor, contractAddr, abi.MaxInt256)
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

	gasBalance, err := i.GetCommittableBalance(ctx, chainID, chain.EthAddress)
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
func (i *inventoryManagerImpl) Rebalance(parentCtx context.Context, chainID int, token common.Address) error {
	// evaluate the rebalance method
	method, err := i.cfg.GetRebalanceMethod(chainID, token.Hex())
	if err != nil {
		return fmt.Errorf("could not get rebalance method: %w", err)
	}
	if method == relconfig.RebalanceMethodNone {
		return nil
	}
	ctx, span := i.handler.Tracer().Start(parentCtx, "Rebalance", trace.WithAttributes(
		attribute.Int(metrics.ChainID, chainID),
		attribute.String("token", token.Hex()),
		attribute.String("rebalance_method", method.String()),
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
	for _, pendingReb := range pendingRebalances {
		registerErr := i.registerPendingRebalance(ctx, pendingReb)
		if registerErr != nil {
			span.AddEvent("could not register pending rebalance", trace.WithAttributes(attribute.String("error", registerErr.Error())))
		}
	}

	// execute the rebalance
	manager, ok := i.rebalanceManagers[method]
	if !ok {
		return fmt.Errorf("no rebalance manager for method: %s", method)
	}
	err = manager.Execute(ctx, rebalance)
	if err != nil {
		return fmt.Errorf("could not execute rebalance: %w", err)
	}
	return nil
}

// registerPendingRebalance registers a callback to update the pending rebalance amount gauge.
func (i *inventoryManagerImpl) registerPendingRebalance(ctx context.Context, rebalance *reldb.Rebalance) (err error) {
	if rebalance == nil || i.meter == nil || i.pendingHist == nil {
		return nil
	}

	attributes := attribute.NewSet(
		attribute.Int(metrics.Origin, int(rebalance.Origin)),
		attribute.Int(metrics.Destination, int(rebalance.Destination)),
		attribute.String("status", rebalance.Status.String()),
		attribute.String("relayer", i.relayerAddress.Hex()),
	)
	tokenMetadata, err := i.GetTokenMetadata(int(rebalance.Origin), rebalance.OriginTokenAddr)
	if err != nil {
		return fmt.Errorf("could not get token metadata: %w", err)
	}
	i.pendingHist.Record(ctx, core.BigToDecimals(rebalance.OriginAmount, tokenMetadata.Decimals), metric.WithAttributeSet(attributes))
	return nil
}

//nolint:cyclop,gocognit
func getRebalance(span trace.Span, cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, chainID int, token common.Address) (rebalance *RebalanceData, err error) {
	maintenancePct, err := cfg.GetMaintenanceBalancePct(chainID, token.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get maintenance pct: %w", err)
	}

	// get token metadata
	var rebalanceTokenData *TokenMetadata
	for address, tokenData := range tokens[chainID] {
		if address == token {
			rebalanceTokenData = tokenData
			break
		}
	}

	// get total balance for given token across all chains
	totalBalance := big.NewInt(0)
	for _, tokenMap := range tokens {
		for _, tokenData := range tokenMap {
			if tokenData.Name == rebalanceTokenData.Name {
				totalBalance.Add(totalBalance, tokenData.Balance)
			}
		}
	}

	// check if any balances are below maintenance threshold
	var minTokenData, maxTokenData *TokenMetadata
	for _, tokenMap := range tokens {
		for _, tokenData := range tokenMap {
			if tokenData.Name == rebalanceTokenData.Name {
				if minTokenData == nil || tokenData.Balance.Cmp(minTokenData.Balance) < 0 {
					minTokenData = tokenData
				}
				if maxTokenData == nil || tokenData.Balance.Cmp(maxTokenData.Balance) > 0 {
					maxTokenData = tokenData
				}
			}
		}
	}

	// get the initialPct for the origin chain
	initialPct, err := cfg.GetInitialBalancePct(maxTokenData.ChainID, maxTokenData.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get initial pct: %w", err)
	}
	maintenanceThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(maintenancePct/100)).Int(nil)
	if span != nil {
		span.SetAttributes(attribute.Float64("maintenance_pct", maintenancePct))
		span.SetAttributes(attribute.Float64("initial_pct", initialPct))
		span.SetAttributes(attribute.String("max_token_balance", maxTokenData.Balance.String()))
		span.SetAttributes(attribute.String("min_token_balance", minTokenData.Balance.String()))
		span.SetAttributes(attribute.String("total_balance", totalBalance.String()))
		span.SetAttributes(attribute.String("maintenance_thresh", maintenanceThresh.String()))
	}

	// check if the minimum balance is below the threshold and trigger rebalance
	if minTokenData.Balance.Cmp(maintenanceThresh) > 0 {
		return rebalance, nil
	}

	// calculate the amount to rebalance vs the initial threshold on origin
	initialThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(initialPct/100)).Int(nil)
	amount := new(big.Int).Sub(maxTokenData.Balance, initialThresh)

	// no need to rebalance since amount would not be positive
	if amount.Cmp(big.NewInt(0)) <= 0 {
		//nolint:nilnil
		return nil, nil
	}

	// filter the rebalance amount by the configured min
	minAmount := cfg.GetMinRebalanceAmount(maxTokenData.ChainID, maxTokenData.Addr)
	if amount.Cmp(minAmount) < 0 {
		// no need to rebalance
		//nolint:nilnil
		return nil, nil
	}

	// clip the rebalance amount by the configured max
	maxAmount := cfg.GetMaxRebalanceAmount(maxTokenData.ChainID, maxTokenData.Addr)
	if amount.Cmp(maxAmount) > 0 {
		amount = maxAmount
	}
	if span != nil {
		span.SetAttributes(
			attribute.String("initial_thresh", initialThresh.String()),
			attribute.String("rebalance_amount", amount.String()),
			attribute.String("max_rebalance_amount", maxAmount.String()),
		)
	}

	rebalance = &RebalanceData{
		OriginMetadata: maxTokenData,
		DestMetadata:   minTokenData,
		Amount:         amount,
	}
	return rebalance, nil
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

	type registerCall func() error
	// TODO: this can be pre-capped w/ len(cfg.Tokens) for each chain id.
	// here we register metrics for exporting through otel. We wait to call these functions until are tokens have been initialized to avoid nil issues.
	var deferredRegisters []registerCall
	deferredCalls := make(map[int][]w3types.Caller)

	// iterate through all tokens to get the metadata
	for cid, chainCfg := range cfg.GetChains() {
		chainID := cid //capture func literal
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
				IsGasToken: tokenName == nativeToken,
				ChainID:    chainID,
				Allowances: make(map[spendableContract]*big.Int),
			}

			var token common.Address
			if rtoken.IsGasToken {
				token = chain.EthAddress
			} else {
				token = common.HexToAddress(tokenCfg.Address)
			}
			i.tokens[chainID][token] = rtoken
			rtoken.Addr = token

			// requires non-nil pointer
			rtoken.Balance = new(big.Int)
			for _, contract := range []spendableContract{contractRFQ, contractSynapseCCTP, contractTokenMessenger} {
				rtoken.Allowances[contract] = new(big.Int)
			}

			if rtoken.IsGasToken {
				rtoken.Decimals = 18
				rtoken.Name = tokenName
				rtoken.Balance = i.gasBalances[chainID]
				// TODO: start allowance?
			} else {
				rfqAddr, err := cfg.GetRFQAddress(chainID)
				if err != nil {
					return fmt.Errorf("could not get rfq address: %w", err)
				}
				deferredCalls[chainID] = append(deferredCalls[chainID],
					eth.CallFunc(funcBalanceOf, token, i.relayerAddress).Returns(rtoken.Balance),
					eth.CallFunc(funcDecimals, token).Returns(&rtoken.Decimals),
					eth.CallFunc(funcName, token).Returns(&rtoken.Name),
					eth.CallFunc(funcAllowance, token, i.relayerAddress, common.HexToAddress(rfqAddr)).Returns(rtoken.Allowances[contractRFQ]),
				)
				cctpAddr, _ := cfg.GetSynapseCCTPAddress(chainID)
				if len(cctpAddr) > 0 {
					deferredCalls[chainID] = append(deferredCalls[chainID],
						eth.CallFunc(funcAllowance, token, i.relayerAddress, common.HexToAddress(cctpAddr)).Returns(rtoken.Allowances[contractSynapseCCTP]),
					)
				}
				messengerAddr, _ := cfg.GetTokenMessengerAddress(chainID)
				if len(messengerAddr) > 0 {
					deferredCalls[chainID] = append(deferredCalls[chainID],
						eth.CallFunc(funcAllowance, token, i.relayerAddress, common.HexToAddress(messengerAddr)).Returns(rtoken.Allowances[contractTokenMessenger]),
					)
				}
			}

			deferredRegisters = append(deferredRegisters, func() error {
				//nolint:wrapcheck
				return i.registerBalance(ctx, chainID, token)
			})
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

	for _, register := range deferredRegisters {
		err = register()
		if err != nil {
			return fmt.Errorf("could not register func: %w", err)
		}
	}

	return nil
}

var logger = log.Logger("inventory")

// refreshBalances refreshes all the token balances.
func (i *inventoryManagerImpl) refreshBalances(ctx context.Context) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	var wg sync.WaitGroup
	wg.Add(len(i.tokens))

	type registerCall func() error
	// TODO: this can be pre-capped w/ len(cfg.Tokens) for each chain id.
	// here we register metrics for exporting through otel. We wait to call these functions until are tokens have been initialized to avoid nil issues.
	var deferredRegisters []registerCall

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
		deferredRegisters = append(deferredRegisters, func() error {
			//nolint:wrapcheck
			return i.registerBalance(ctx, chainID, chain.EthAddress)
		})

		// queue token balance fetches
		for tokenAddress, token := range tokenMap {
			// TODO: make sure Returns does nothing on error
			if !token.IsGasToken {
				deferredCalls = append(deferredCalls, eth.CallFunc(funcBalanceOf, tokenAddress, i.relayerAddress).Returns(token.Balance))
				deferredRegisters = append(deferredRegisters, func() error {
					//nolint:wrapcheck
					return i.registerBalance(ctx, chainID, tokenAddress)
				})
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

	for _, register := range deferredRegisters {
		err := register()
		if err != nil {
			logger.Warnf("could not register func: %v", err)
		}
	}

	return nil
}

func (i *inventoryManagerImpl) registerBalance(ctx context.Context, chainID int, token common.Address) (err error) {
	if i.meter == nil || i.balanceHist == nil {
		return nil
	}

	// TODO: make sure this doesn't get called until we're done
	tokenData, ok := i.tokens[chainID][token]
	if !ok {
		return fmt.Errorf("could not find token in chainTokens for chainID: %d, token: %s", chainID, token)
	}

	attributes := attribute.NewSet(
		attribute.Int(metrics.ChainID, chainID),
		attribute.String("relayer_address", i.relayerAddress.String()),
		attribute.String("token_name", tokenData.Name),
		attribute.Int("decimals", int(tokenData.Decimals)),
		attribute.String("token_address", token.String()),
		attribute.String("raw_balance", tokenData.Balance.String()),
		attribute.String("relayer", i.relayerAddress.Hex()),
	)

	i.balanceHist.Record(ctx, core.BigToDecimals(tokenData.Balance, tokenData.Decimals), metric.WithAttributeSet(attributes))
	return nil
}

// Ultimately this should produce a list of all balances and remove the
// quoted amounts from the database
