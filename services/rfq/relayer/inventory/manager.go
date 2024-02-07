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
	// GetCommittableBalance gets the total balance available for quotes
	// this does not include on-chain balances committed in previous quotes that may be
	// refunded in the event of a revert.
	GetCommittableBalance(ctx context.Context, chainID int, token common.Address, options ...BalanceFetchArgOption) (*big.Int, error)
	// GetCommittableBalances gets the total balances committable for all tracked tokens.
	GetCommittableBalances(ctx context.Context, options ...BalanceFetchArgOption) (map[int]map[common.Address]*big.Int, error)
	// ApproveAllTokens approves all tokens for the relayer address.
	ApproveAllTokens(ctx context.Context, submitter submitter.TransactionSubmitter) error
	// HasSufficientGas checks if there is sufficient gas for a given route.
	HasSufficientGas(ctx context.Context, origin, dest int) (bool, error)
}

type inventoryManagerImpl struct {
	// map chainID->address->tokenMetadata
	tokens map[int]map[common.Address]*tokenMetadata
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
	db          reldb.Service
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
			res[chainID][address] = core.CopyBigInt(tokenData.balance)
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

type tokenMetadata struct {
	name           string
	balance        *big.Int
	decimals       uint8
	startAllowance *big.Int
	isGasToken     bool
}

var (
	funcBalanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")
	funcName      = w3.MustNewFunc("name()", "string")
	funcDecimals  = w3.MustNewFunc("decimals()", "uint8")
	funcAllowance = w3.MustNewFunc("allowance(address,address)", "uint256")
)

// TODO: replace w/ config.
const defaultPollPeriod = 5

// NewInventoryManager creates a list of tokens we should use.
func NewInventoryManager(ctx context.Context, clientFetcher submitter.ClientFetcher, handler metrics.Handler, cfg relconfig.Config, relayer common.Address, db reldb.Service) (Manager, error) {
	i := inventoryManagerImpl{
		relayerAddress: relayer,
		handler:        handler,
		cfg:            cfg,
		chainClient:    clientFetcher,
		db:             db,
	}

	err := i.initializeTokens(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not initialize tokens: %w", err)
	}

	// TODO: move
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(defaultPollPeriod * time.Second):
				// this returning an error isn't really possible unless a config error happens
				// TODO: need better error handling.
				err = i.refreshBalances(ctx)
				if err != nil {
					logger.Errorf("could not refresh balances")
					return
				}
			}
		}
	}()

	return &i, nil
}

const maxBatchSize = 10

// ApproveAllTokens approves all checks if allowance is set and if not approves.
func (i *inventoryManagerImpl) ApproveAllTokens(ctx context.Context, submitter submitter.TransactionSubmitter) error {
	i.mux.RLock()
	defer i.mux.RUnlock()

	for chainID, tokenMap := range i.tokens {
		backendClient, err := i.chainClient.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		for address, token := range tokenMap {
			// if startAllowance is 0
			if address != chain.EthAddress && token.startAllowance.Cmp(big.NewInt(0)) == 0 {
				chainID := chainID // capture func literal
				address := address // capture func literal
				// init an approval in submitter. Note: in the case where submitter hasn't finished from last boot, this will double submit approvals unfortanutely
				_, err = submitter.SubmitTransaction(ctx, big.NewInt(int64(chainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
					erc20, err := ierc20.NewIERC20(address, backendClient)
					if err != nil {
						return nil, fmt.Errorf("could not get erc20: %w", err)
					}

					approveAmount, err := erc20.Approve(transactor, common.HexToAddress(i.cfg.Chains[chainID].Bridge), abi.MaxInt256)
					if err != nil {
						return nil, fmt.Errorf("could not approve: %w", err)
					}

					return approveAmount, nil
				})
				if err != nil {
					return fmt.Errorf("could not submit approval: %w", err)
				}
			}
		}
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
	i.tokens = make(map[int]map[common.Address]*tokenMetadata)
	i.gasBalances = make(map[int]*big.Int)

	type registerCall func() error
	// TODO: this can be pre-capped w/ len(cfg.Tokens) for each chain id.
	// here we register metrics for exporting through otel. We wait to call these functions until are tokens have been initialized to avoid nil issues.
	var deferredRegisters []registerCall
	deferredCalls := make(map[int][]w3types.Caller)

	// iterate through all tokens to get the metadata
	for chainID, chainCfg := range cfg.GetChains() {
		i.tokens[chainID] = map[common.Address]*tokenMetadata{}

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
			rtoken := &tokenMetadata{
				isGasToken: tokenName == nativeToken,
			}

			var token common.Address
			if rtoken.isGasToken {
				token = chain.EthAddress
			} else {
				token = common.HexToAddress(tokenCfg.Address)
			}
			i.tokens[chainID][token] = rtoken

			// requires non-nil pointer
			rtoken.balance = new(big.Int)
			rtoken.startAllowance = new(big.Int)

			if rtoken.isGasToken {
				rtoken.decimals = 18
				rtoken.name = tokenName
				rtoken.balance = i.gasBalances[chainID]
				// TODO: start allowance?
			} else {
				deferredCalls[chainID] = append(deferredCalls[chainID],
					eth.CallFunc(funcBalanceOf, token, i.relayerAddress).Returns(rtoken.balance),
					eth.CallFunc(funcDecimals, token).Returns(&rtoken.decimals),
					eth.CallFunc(funcName, token).Returns(&rtoken.name),
					eth.CallFunc(funcAllowance, token, i.relayerAddress, common.HexToAddress(i.cfg.Chains[chainID].Bridge)).Returns(rtoken.startAllowance),
				)
			}

			chainID := chainID // capture func literal
			deferredRegisters = append(deferredRegisters, func() error {
				//nolint: wrapcheck
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
			if !token.isGasToken {
				deferredCalls = append(deferredCalls, eth.CallFunc(funcBalanceOf, tokenAddress, i.relayerAddress).Returns(token.balance))
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
			attribute.String("token_name", tokenData.name), attribute.Int("decimals", int(tokenData.decimals)),
			attribute.String("token_address", token.String()))

		observer.ObserveFloat64(balanceGauge, core.BigToDecimals(tokenData.balance, tokenData.decimals), metric.WithAttributeSet(attributes))

		return nil
	}, balanceGauge); err != nil {
		return fmt.Errorf("could not register callback: %w", err)
	}
	return nil
}

// Ultimately this should produce a list of all balances and remove the
// quoted amounts from the database
