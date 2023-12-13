package inventory

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
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
	// map chainID->address->
	tokens map[int]map[common.Address]Token
	// mux contains the mutex
	mux sync.RWMutex
	// callbacksRegisterd is wether or not callbacks have been registered
	callbacksRegistered bool
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
func NewInventoryManager(ctx context.Context, client omnirpcClient.RPCClient, handler metrics.Handler, cfg config.Config, relayer common.Address) (interface{}, error) {
	meter := handler.Meter("github.com/synapsecns/sanguine/services/rfq/relayer/inventory")

	chainTokens := make(map[int]map[common.Address]*Token)
	for chainID, tokens := range cfg.Tokens {
		chainTokens[chainID] = map[common.Address]*Token{}

		for _, strToekn := range tokens {
			balanceGauge, err := meter.Float64ObservableGauge("inventory_balance")
			if err != nil {
				return nil, fmt.Errorf("could not create gauge: %w", err)
			}

			token := common.HexToAddress(strToekn)
			rtoken := &Token{}
			chainTokens[chainID][token] = rtoken

			eth.CallFunc(funcBalanceOf, token, relayer).Returns(rtoken.balance)
			eth.CallFunc(funcDecimals, token).Returns(&rtoken.decimals)
			eth.CallFunc(funcName, token).Returns(&rtoken.name)

			// register the callback. This only needs to happen once
			if _, err := meter.RegisterCallback(func(ctx context.Context, observer metric.Observer) error {
				tokenData, ok := chainTokens[chainID][token]
				if !ok {
					return fmt.Errorf("could not find token in chainTokens for chainID: %d, token: %s", chainID, token)
				}

				attributes := attribute.NewSet(attribute.Int(metrics.ChainID, chainID), attribute.String("relayer_address", relayer.String()),
					attribute.String("token_name", tokenData.name), attribute.Int("decimals", int(tokenData.decimals)),
					attribute.String("token_address", token.String()))

				observer.ObserveFloat64(balanceGauge, bigToDecimals(tokenData.balance, tokenData.decimals), metric.WithAttributeSet(attributes))

				return nil
			}, balanceGauge); err != nil {
				return nil, fmt.Errorf("could not register callback: %w", err)
			}
		}
	}

	return nil, nil
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
