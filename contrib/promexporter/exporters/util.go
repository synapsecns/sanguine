package exporters

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/decoders"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	ethergoClient "github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"golang.org/x/sync/errgroup"
)

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

	//nolint: revive
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
