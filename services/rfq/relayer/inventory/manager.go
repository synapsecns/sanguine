package inventory

import (
	"context"
	"fmt"
	"math/big"
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
	HasSufficientGas(ctx context.Context, origin, dest int) (bool, error)
	// Rebalance checks whether a given token should be rebalanced, and
	// executes the rebalance if necessary.
	Rebalance(ctx context.Context, chainID int, token common.Address) error
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
	Name               string
	Balance            *big.Int
	Decimals           uint8
	StartAllowanceRFQ  *big.Int
	StartAllowanceCCTP *big.Int
	IsGasToken         bool
	ChainID            int
	Addr               common.Address
}

var (
	funcBalanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")
	funcName      = w3.MustNewFunc("name()", "string")
	funcDecimals  = w3.MustNewFunc("decimals()", "uint8")
	funcAllowance = w3.MustNewFunc("allowance(address,address)", "uint256")
)

// TODO: replace w/ config.
const defaultPollPeriod = 5

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
		case relconfig.RebalanceMethodCCTP:
			rebalanceManagers[method] = newRebalanceManagerCCTP(cfg, handler, clientFetcher, txSubmitter, relayer, db)
		default:
			return nil, fmt.Errorf("unsupported rebalance method: %s", method)
		}
	}

	i := inventoryManagerImpl{
		relayerAddress:    relayer,
		handler:           handler,
		cfg:               cfg,
		chainClient:       clientFetcher,
		txSubmitter:       txSubmitter,
		rebalanceManagers: rebalanceManagers,
		db:                db,
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
			if address != chain.EthAddress && token.StartAllowanceRFQ.Cmp(big.NewInt(0)) == 0 {
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

			// approve CCTP contract
			if address != chain.EthAddress && token.StartAllowanceCCTP.Cmp(big.NewInt(0)) == 0 {
				tokenAddr := address // capture func literal
				contractAddr, err := i.cfg.GetCCTPAddress(chainID)
				if err != nil {
					return fmt.Errorf("could not get CCTP address: %w", err)
				}
				err = i.approve(ctx, tokenAddr, common.HexToAddress(contractAddr), backendClient)
				if err != nil {
					return fmt.Errorf("could not approve CCTP contract: %w", err)
				}
			}
		}
	}
	return nil
}

// approve submits an ERC20 approval for a given token and contract address.
func (i *inventoryManagerImpl) approve(ctx context.Context, tokenAddr, contractAddr common.Address, backendClient client.EVM) (err error) {
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
func (i *inventoryManagerImpl) HasSufficientGas(ctx context.Context, origin, dest int) (sufficient bool, err error) {
	gasThresh, err := i.cfg.GetMinGasToken(dest)
	if err != nil {
		return false, fmt.Errorf("error getting min gas token: %w", err)
	}
	gasOrigin, err := i.GetCommittableBalance(ctx, origin, chain.EthAddress)
	if err != nil {
		return false, fmt.Errorf("error getting committable gas on origin: %w", err)
	}
	gasDest, err := i.GetCommittableBalance(ctx, dest, chain.EthAddress)
	if err != nil {
		return false, fmt.Errorf("error getting committable gas on dest: %w", err)
	}

	sufficient = gasOrigin.Cmp(gasThresh) >= 0 && gasDest.Cmp(gasThresh) >= 0
	return sufficient, nil
}

// Rebalance checks whether a given token should be rebalanced, and executes the rebalance if necessary.
// Note that if there are multiple tokens whose balance is below the maintenance balance, only the lowest balance
// will be rebalanced.
func (i *inventoryManagerImpl) Rebalance(ctx context.Context, chainID int, token common.Address) error {
	// evaluate the rebalance method
	method, err := i.cfg.GetRebalanceMethod(chainID, token.Hex())
	if err != nil {
		return fmt.Errorf("could not get rebalance method: %w", err)
	}
	if method == relconfig.RebalanceMethodNone {
		return nil
	}

	// build the rebalance action
	rebalance, err := getRebalance(i.cfg, i.tokens, chainID, token)
	if err != nil {
		return fmt.Errorf("could not get rebalance: %w", err)
	}
	if rebalance == nil {
		return nil
	}

	// make sure there are no pending rebalances that touch the given path
	pending, err := i.db.HasPendingRebalance(ctx, uint64(rebalance.OriginMetadata.ChainID), uint64(rebalance.DestMetadata.ChainID))
	if err != nil {
		return fmt.Errorf("could not check pending rebalance: %w", err)
	}
	if pending {
		return nil
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

//nolint:cyclop
func getRebalance(cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, chainID int, token common.Address) (rebalance *RebalanceData, err error) {
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

	// check if the minimum balance is below the threshold and trigger rebalance
	maintenanceThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(maintenancePct/100)).Int(nil)
	if minTokenData.Balance.Cmp(maintenanceThresh) < 0 {
		initialThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(initialPct/100)).Int(nil)
		amount := new(big.Int).Sub(maxTokenData.Balance, initialThresh)
		if amount.Cmp(big.NewInt(0)) < 0 {
			// do not rebalance since it would take us below initial threshold
			//nolint:nilnil
			return nil, nil
		}
		rebalance = &RebalanceData{
			OriginMetadata: maxTokenData,
			DestMetadata:   minTokenData,
			Amount:         amount,
		}
	}
	return rebalance, nil
}

// initializes tokens converts the configuration into a data structure we can use to determine inventory
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

	meter := i.handler.Meter("github.com/synapsecns/sanguine/services/rfq/relayer/inventory")

	// TODO: this needs to be a struct bound variable otherwise will be stuck.
	i.tokens = make(map[int]map[common.Address]*TokenMetadata)
	i.gasBalances = make(map[int]*big.Int)

	type registerCall func() error
	// TODO: this can be pre-capped w/ len(cfg.Tokens) for each chain id.
	// here we register metrics for exporting through otel. We wait to call these functions until are tokens have been initialized to avoid nil issues.
	var deferredRegisters []registerCall
	deferredCalls := make(map[int][]w3types.Caller)

	// iterate through all tokens to get the metadata
	for chainID, chainCfg := range cfg.GetChains() {
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
			rtoken.StartAllowanceRFQ = new(big.Int)
			rtoken.StartAllowanceCCTP = new(big.Int)

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
				cctpAddr, err := cfg.GetCCTPAddress(chainID)
				if err != nil {
					return fmt.Errorf("could not get cctp address: %w", err)
				}
				deferredCalls[chainID] = append(deferredCalls[chainID],
					eth.CallFunc(funcBalanceOf, token, i.relayerAddress).Returns(rtoken.Balance),
					eth.CallFunc(funcDecimals, token).Returns(&rtoken.Decimals),
					eth.CallFunc(funcName, token).Returns(&rtoken.Name),
					eth.CallFunc(funcAllowance, token, i.relayerAddress, common.HexToAddress(rfqAddr)).Returns(rtoken.StartAllowanceRFQ),
					eth.CallFunc(funcAllowance, token, i.relayerAddress, common.HexToAddress(cctpAddr)).Returns(rtoken.StartAllowanceCCTP),
				)
			}

			chainID := chainID // capture func literal
			deferredRegisters = append(deferredRegisters, func() error {
				//nolint:wrapcheck
				return i.registerMetric(meter, chainID, token)
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

	for chainID, tokenMap := range i.tokens {
		chainClient, err := i.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		// queue gas token balance fetch
		deferredCalls := []w3types.Caller{
			eth.Balance(i.relayerAddress, nil).Returns(i.gasBalances[chainID]),
		}

		// queue token balance fetches
		for tokenAddress, token := range tokenMap {
			// TODO: make sure Returns does nothing on error
			if !token.IsGasToken {
				deferredCalls = append(deferredCalls, eth.CallFunc(funcBalanceOf, tokenAddress, i.relayerAddress).Returns(token.Balance))
			}
		}

		chainID := chainID // capture func literal
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

func (i *inventoryManagerImpl) registerMetric(meter metric.Meter, chainID int, token common.Address) error {
	balanceGauge, err := meter.Float64ObservableGauge("inventory_balance")
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	if _, err := meter.RegisterCallback(func(ctx context.Context, observer metric.Observer) error {
		i.mux.RLock()
		defer i.mux.RUnlock()

		// TODO: make sure this doesn't get called until we're done
		tokenData, ok := i.tokens[chainID][token]
		if !ok {
			return fmt.Errorf("could not find token in chainTokens for chainID: %d, token: %s", chainID, token)
		}

		attributes := attribute.NewSet(attribute.Int(metrics.ChainID, chainID), attribute.String("relayer_address", i.relayerAddress.String()),
			attribute.String("token_name", tokenData.Name), attribute.Int("decimals", int(tokenData.Decimals)),
			attribute.String("token_address", token.String()))

		observer.ObserveFloat64(balanceGauge, core.BigToDecimals(tokenData.Balance, tokenData.Decimals), metric.WithAttributeSet(attributes))

		return nil
	}, balanceGauge); err != nil {
		return fmt.Errorf("could not register callback: %w", err)
	}
	return nil
}

// Ultimately this should produce a list of all balances and remove the
// quoted amounts from the database
