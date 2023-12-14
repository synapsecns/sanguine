package inventory

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"math"
	"math/big"
	"sync"
)

// What we actaully want to be able to do here is
type InventoryManager interface {
	// GetCommittableBalance gets the total balance available for quotes
	// this does not include on-chain balances commited in previous quotes that may be
	// refunded in the event of a revert.
	GetCommittableBalance(ctx context.Context, chainID int, token common.Address, options ...BalanceFetchArgOption)
	// GetCommitableBalances gets the total balances commitable for all tracked tokens.
	GetCommitableBalances(ctx context.Context, options ...BalanceFetchArgOption) map[int]map[common.Address]*big.Int
}

type inventoryManagerImpl struct {
	// map chainID->address->Token
	tokens map[int]map[common.Address]*Token
	// mux contains the mutex
	mux sync.RWMutex
	// handler is the metrics handler
	handler metrics.Handler
	// relayerAddress contains the relayer address
	relayerAddress common.Address
	// chainClient is an omnirpc client
	chainClient omnirpcClient.RPCClient
}

type Token struct {
	name     string
	balance  *big.Int
	decimals uint8
}

var (
	funcBalanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")
	funcName      = w3.MustNewFunc("name()", "string")
	funcDecimals  = w3.MustNewFunc("decimals()", "uint8")
)

// NewInventoryManager creates a list of tokens we should use.
func NewInventoryManager(ctx context.Context, client omnirpcClient.RPCClient, handler metrics.Handler, cfg relconfig.Config, relayer common.Address) (interface{}, error) {
	i := inventoryManagerImpl{
		relayerAddress: relayer,
		handler:        handler,
		chainClient:    client,
	}

	err := i.initializeTokens(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not initialize tokens")
	}

	return nil, nil
}

const maxBatchSize = 10

// initlalizes tokens converts the configuration into a data structure we can use to determine inventory
// it gets metadata like name, decimals, etc once and exports these to prometheus for ease of debugging.
func (i *inventoryManagerImpl) initializeTokens(parentCtx context.Context, cfg relconfig.Config) (err error) {
	i.mux.Lock()
	defer i.mux.Unlock()

	ctx, span := i.handler.Tracer().Start(parentCtx, "initializeTokens", trace.WithAttributes(
		attribute.String("relayer_address", i.relayerAddress.String()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	meter := i.handler.Meter("github.com/synapsecns/sanguine/services/rfq/relayer/inventory")

	// TODO: this needs to be a struct bound variable otherwise will be stuck.
	i.tokens = make(map[int]map[common.Address]*Token)

	type registerCall func() error
	// TODO: this can be pre-capped w/ len(cfg.Tokens) for each chain id.
	// here we register metrics for exporting through otel. We wait to call these functions until are tokens have been initialized to avoid nil issues.
	var deferredRegisters []registerCall
	deferredCalls := make(map[int][]w3types.Caller)

	// iterate through all tokens to get the metadata
	for chainID, tokens := range cfg.Tokens {
		i.tokens[chainID] = map[common.Address]*Token{}

		for _, strToekn := range tokens {
			token := common.HexToAddress(strToekn)
			rtoken := &Token{}
			i.tokens[chainID][token] = rtoken
			// requires non-nil pointer
			rtoken.balance = new(big.Int)

			deferredCalls[chainID] = append(deferredCalls[chainID],
				eth.CallFunc(funcBalanceOf, token, i.relayerAddress).Returns(rtoken.balance),
				eth.CallFunc(funcDecimals, token).Returns(&rtoken.decimals),
				eth.CallFunc(funcName, token).Returns(&rtoken.name),
			)

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

		chainClient, err := i.chainClient.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("can't initialize tokens, no chain client available for chain %d: %w", chainID, err)
		}

		g.Go(func() error {
			// TODO: add retries
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

	return
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

		observer.ObserveFloat64(balanceGauge, bigToDecimals(tokenData.balance, tokenData.decimals), metric.WithAttributeSet(attributes))

		return nil
	}, balanceGauge); err != nil {
		return fmt.Errorf("could not register callback: %w", err)
	}
	return nil
}

// refreshBalances refreshes all the token balances.
func (i *inventoryManagerImpl) refreshBalances() {

}

// Ultimately this should produce a list of all balances and remove the
// quoted amounts from the database

// TODO: move me
func bigToDecimals(bigInt *big.Int, decimals uint8) float64 {
	// Convert vpriceMetric to *big.Float
	bigVPrice := new(big.Float).SetInt(bigInt)

	// Calculate the divisor for decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(int(decimals)))

	// Divide bigVPrice by the divisor to account for decimals
	realVPrice := new(big.Float).Quo(bigVPrice, divisor)

	// Convert the final value to float64
	floatVPrice, _ := realVPrice.Float64()
	return floatVPrice
}
