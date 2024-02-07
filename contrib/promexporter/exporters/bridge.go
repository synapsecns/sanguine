package exporters

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/decoders"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	ethergoClient "github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"math/big"
	"time"
)

func (e *exporter) getBridgeConfig(ctx context.Context) (*bridgeconfig.BridgeConfigRef, error) {
	client, err := e.omnirpcClient.GetConfirmationsClient(ctx, e.cfg.BridgeConfig.ChainID, 1)
	if err != nil {
		return nil, fmt.Errorf("could not get confirmations client: %w", err)
	}

	// note this will not update
	configContract, err := bridgeconfig.NewBridgeConfigRef(common.HexToAddress(e.cfg.BridgeConfig.Address), client)
	if err != nil {
		return nil, fmt.Errorf("could not get bridge config contract: %w", err)
	}
	return configContract, nil
}

// Will be a lot faster w/: https://github.com/open-telemetry/opentelemetry-go/issues/3034
// nolint: cyclop
func (e *exporter) vpriceStats(ctx context.Context, chainID int, tokenID string) error {
	meter := e.metrics.Meter(meterName)
	vpriceMetric, err := meter.Float64ObservableGauge("vpriceMetric")
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	attributes := attribute.NewSet(attribute.Int(metrics.ChainID, chainID), attribute.String("tokenID", tokenID))

	client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
	if err != nil {
		return fmt.Errorf("could not get confirmations client: %w", err)
	}

	bridgeConfig, err := e.getBridgeConfig(ctx)
	if err != nil {
		return err
	}

	token, err := bridgeConfig.GetToken(&bind.CallOpts{Context: ctx}, tokenID, big.NewInt(int64(chainID)))
	if err != nil {
		return fmt.Errorf("could not get token: %w", err)
	}

	poolConfig, err := bridgeConfig.GetPoolConfig(&bind.CallOpts{Context: ctx}, common.HexToAddress(token.TokenAddress), big.NewInt(int64(chainID)))
	if err != nil {
		return errPoolNotExist
	}

	// pool doesn't exist, no metrics to record!
	if poolConfig.PoolAddress == common.BigToAddress(big.NewInt(0)) {
		return errPoolNotExist
	}

	tokenContract, err := bridge.NewERC20(common.HexToAddress(token.TokenAddress), client)
	if err != nil {
		return fmt.Errorf("could not get tokenID contract: %w", err)
	}

	decimals, err := tokenContract.Decimals(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get decimals: %w", err)
	}

	iswap, err := swap.NewISwap(poolConfig.PoolAddress, client)
	if err != nil {
		return fmt.Errorf("could not get iswap contract: %w", err)
	}

	if _, err := meter.RegisterCallback(func(parentCtx context.Context, o metric.Observer) (err error) {
		ctx, span := e.metrics.Tracer().Start(parentCtx, "vprice_stats", trace.WithAttributes(
			attribute.Int(metrics.ChainID, chainID), attribute.String("tokenID", tokenID),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		realvPrice, err := iswap.GetVirtualPrice(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get virtual price: %w", err)
		}

		// Use floatVPrice as required
		o.ObserveFloat64(vpriceMetric, core.BigToDecimals(realvPrice, decimals), metric.WithAttributeSet(attributes))

		return nil
	}, vpriceMetric); err != nil {
		return fmt.Errorf("registering callback on instruments: %w", err)
	}

	return nil
}

var errPoolNotExist = errors.New("pool does not exist")

// nolint: cyclop
func (e *exporter) getTokenBalances(ctx context.Context) error {
	allTokens, err := e.getAllTokens(ctx)
	if err != nil {
		return fmt.Errorf("could not get all tokens: %w", err)
	}

	for chainID, bridgeContract := range e.cfg.BridgeChecks {
		chainID := chainID
		bridgeContract := bridgeContract // capture func literals
		meter := e.metrics.Meter(meterName)

		bridgeBalanceMetric, err := meter.Float64ObservableGauge("bridgeBalanceMetric")
		if err != nil {
			return fmt.Errorf("could not create gauge: %w", err)
		}

		feeBalanceMetric, err := meter.Float64ObservableCounter("feeBalance")
		if err != nil {
			return fmt.Errorf("could not create counter: %w", err)
		}

		totalSupplyMetric, err := meter.Float64ObservableGauge("totalSupply")
		if err != nil {
			return fmt.Errorf("could not create gauge: %w", err)
		}

		gasBalanceMetric, err := meter.Float64ObservableGauge("gasBalance")
		if err != nil {
			return fmt.Errorf("could not create gauge: %w", err)
		}

		if _, err := meter.RegisterCallback(func(parentCtx context.Context, o metric.Observer) (err error) {
			ctx, span := e.metrics.Tracer().Start(parentCtx, "tokenbalances", trace.WithAttributes(
				attribute.Int(metrics.ChainID, chainID),
			))

			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
			if err != nil {
				return fmt.Errorf("could not get confirmations client: %w", err)
			}

			var realGasBalance big.Int
			calls := []w3types.Caller{
				eth.Balance(common.HexToAddress(bridgeContract), nil).Returns(&realGasBalance),
			}

			type tokenData struct {
				metadata        TokenConfig
				contractBalance *big.Int
				totalSuppply    *big.Int
				feeBalance      *big.Int
			}

			allTokenData := make([]tokenData, len(allTokens.GetForChainID(chainID)))

			for i, tokenConfig := range allTokens.GetForChainID(chainID) {
				// initialize empty struct
				allTokenData[i] = tokenData{
					metadata:        tokenConfig,
					contractBalance: new(big.Int),
					totalSuppply:    new(big.Int),
					feeBalance:      new(big.Int),
				}

				calls = append(calls,
					eth.CallFunc(decoders.FuncBalanceOf(), tokenConfig.TokenAddress, common.HexToAddress(bridgeContract)).Returns(allTokenData[i].contractBalance),
					eth.CallFunc(decoders.FuncTotalSupply(), tokenConfig.TokenAddress).Returns(allTokenData[i].totalSuppply),
					eth.CallFunc(decoders.FuncFeeBalance(), common.HexToAddress(bridgeContract), tokenConfig.TokenAddress).Returns(allTokenData[i].feeBalance),
				)
			}

			err = e.batchCalls(ctx, client, calls)
			if err != nil {
				return fmt.Errorf("could not get token balances: %w", err)
			}

			// eth is always 18 decimals
			o.ObserveFloat64(gasBalanceMetric, core.BigToDecimals(&realGasBalance, 18), metric.WithAttributes(attribute.Int(metrics.ChainID, chainID)))

			for _, td := range allTokenData {
				tokenAttributes := attribute.NewSet(attribute.String("tokenID", td.metadata.TokenID), attribute.Int(metrics.ChainID, td.metadata.ChainID))
				o.ObserveFloat64(bridgeBalanceMetric, core.BigToDecimals(td.contractBalance, td.metadata.TokenDecimals), metric.WithAttributeSet(tokenAttributes))
				o.ObserveFloat64(feeBalanceMetric, core.BigToDecimals(td.feeBalance, td.metadata.TokenDecimals), metric.WithAttributeSet(tokenAttributes))
				o.ObserveFloat64(totalSupplyMetric, core.BigToDecimals(td.totalSuppply, td.metadata.TokenDecimals), metric.WithAttributeSet(tokenAttributes))
			}

			return nil
		}, bridgeBalanceMetric, feeBalanceMetric, totalSupplyMetric, gasBalanceMetric); err != nil {
			return fmt.Errorf("could not register")
		}
	}

	return nil
}

func (e *exporter) getAllTokens(parentCtx context.Context) (allTokens Tokens, err error) {
	allTokens = []TokenConfig{}

	ctx, span := e.metrics.Tracer().Start(parentCtx, "get_all_tokens")

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	bridgeConfig, err := e.getBridgeConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get bridge config: %w", err)
	}

	// TODO: multicall is preferable here, but I ain't got time for that
	tokenIDs, err := bridgeConfig.GetAllTokenIDs(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("could not get all token ids: %w", err)
	}

	bridgeConfigClient, err := e.omnirpcClient.GetConfirmationsClient(ctx, e.cfg.BridgeConfig.ChainID, 1)
	if err != nil {
		return nil, fmt.Errorf("could not get confirmations client: %w", err)
	}

	bridgeTokens := make([]*bridgeconfig.BridgeConfigV3Token, len(tokenIDs)*len(e.cfg.BridgeChecks))
	tokenIDS := make([]string, len(tokenIDs)*len(e.cfg.BridgeChecks))

	var calls []w3types.Caller

	i := 0
	for _, tokenID := range tokenIDs {
		for chainID := range e.cfg.BridgeChecks {
			token := &bridgeconfig.BridgeConfigV3Token{}
			calls = append(calls, eth.CallFunc(decoders.TokenConfigGetToken(), bridgeConfig.Address(), tokenID, big.NewInt(int64(chainID))).Returns(token))
			bridgeTokens[i] = token
			tokenIDS[i] = tokenID
			i++
		}
	}

	// TODO: once go 1.21 is introduced do min(cfg.BatchCallLimit, 2)
	err = e.batchCalls(ctx, bridgeConfigClient, calls)
	if err != nil {
		return nil, fmt.Errorf("could not get token balances: %w", err)
	}

	for i, token := range bridgeTokens {
		tokenID := tokenIDS[i]

		if token.TokenAddress == "" {
			continue
		}

		allTokens = append(allTokens, TokenConfig{
			TokenID:       tokenID,
			ChainID:       int(token.ChainId.Int64()),
			TokenAddress:  common.HexToAddress(token.TokenAddress),
			TokenDecimals: token.TokenDecimals,
			HasUnderlying: token.HasUnderlying,
			IsUnderlying:  token.IsUnderlying,
		})
	}

	return allTokens, nil
}

func (e *exporter) batchCalls(ctx context.Context, evmClient ethergoClient.EVM, calls []w3types.Caller) (err error) {
	tasks := core.ChunkSlice(calls, e.cfg.BatchCallLimit)

	g, ctx := errgroup.WithContext(ctx)
	for _, task := range tasks {
		task := task // capture func literal
		g.Go(func() error {
			err = evmClient.BatchWithContext(ctx, task...)
			if err != nil {
				return fmt.Errorf("could not batch calls: %w", err)
			}

			return nil
		})
	}

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not get token balances: %w", err)
	}

	return nil
}

// Tokens is a list of token configs.
type Tokens []TokenConfig

// GetForChainID returns all tokens for a given chainID.
func (t Tokens) GetForChainID(chainID int) Tokens {
	var chainTokens []TokenConfig
	for _, token := range t {
		if token.ChainID == chainID {
			chainTokens = append(chainTokens, token)
		}
	}

	return chainTokens
}

// TokenConfig is a cleaned up token config.
type TokenConfig struct {
	TokenID       string
	ChainID       int
	TokenAddress  common.Address
	TokenDecimals uint8
	HasUnderlying bool
	IsUnderlying  bool
}
