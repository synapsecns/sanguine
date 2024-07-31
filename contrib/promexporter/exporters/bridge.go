package exporters

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/decoders"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
)

func (e *exporter) getBridgeConfig(ctx context.Context) (*bridgeconfig.BridgeConfigRef, error) {
	// client, err := e.omnirpcClient.GetClient(ctx, big.NewInt(int64(e.cfg.BridgeConfig.ChainID)))
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

	poolConfig, err := bridgeConfig.GetPoolConfig(
		&bind.CallOpts{Context: ctx},
		common.HexToAddress(token.TokenAddress),
		big.NewInt(int64(chainID)),
	)
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

	realvPrice, err := iswap.GetVirtualPrice(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get virtual price: %w", err)
	}

	e.otelRecorder.RecordVPrice(chainID, core.BigToDecimals(realvPrice, decimals))

	return nil
}

var errPoolNotExist = errors.New("pool does not exist")

// nolint: cyclop
func (e *exporter) getTokenBalancesStats(ctx context.Context) error {
	allTokens, err := e.getAllTokens(ctx)

	if err != nil {
		return fmt.Errorf("could not get all tokens: %w", err)
	}

	for chainID, bridgeContract := range e.cfg.BridgeChecks {
		client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
		if err != nil {
			return fmt.Errorf("could not get confirmations client: %w", err)
		}

		var realGasBalance big.Int
		calls := []w3types.Caller{
			eth.Balance(common.HexToAddress(bridgeContract), nil).Returns(&realGasBalance),
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
		_ = e.batchCalls(ctx, client, calls)

		e.otelRecorder.RecordBridgeGasBalance(chainID, core.BigToDecimals(&realGasBalance, 18)*params.Ether)

		for _, td := range allTokenData {
			e.otelRecorder.RecordTokenBalance(chainID, td)
		}

	}

	return nil
}
