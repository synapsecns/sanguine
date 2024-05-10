package inventory

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// RebalanceData contains metadata for a rebalance action.
type RebalanceData struct {
	OriginMetadata *TokenMetadata
	DestMetadata   *TokenMetadata
	Amount         *big.Int
	Method         relconfig.RebalanceMethod
}

// RebalanceManager is the interface for the rebalance manager.
type RebalanceManager interface {
	// Start starts the rebalance manager.
	Start(ctx context.Context) (err error)
	// Execute executes a rebalance action.
	Execute(ctx context.Context, rebalance *RebalanceData) error
}

//nolint:cyclop,gocognit,nilnil
func getRebalance(span trace.Span, cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, chainID int, token common.Address) (rebalance *RebalanceData, err error) {
	maintenancePct, err := cfg.GetMaintenanceBalancePct(chainID, token.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get maintenance pct: %w", err)
	}

	// get rebalance method
	method, err := cfg.GetRebalanceMethod(chainID, token.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get rebalance method: %w", err)
	}
	if method == relconfig.RebalanceMethodNone {
		return nil, nil
	}

	// get token metadata
	var rebalanceTokenData *TokenMetadata
	for address, tokenData := range tokens[chainID] {
		if address == token {
			rebalanceTokenData = tokenData
			break
		}
	}

	// evaluate the origin and dest of the rebalance based on min/max token balances
	var destTokenData, originTokenData *TokenMetadata
	for _, tokenMap := range tokens {
		for _, tokenData := range tokenMap {
			if tokenData.Name == rebalanceTokenData.Name {
				// make sure that the token is compatible with our rebalance method
				tokenMethod, tokenErr := cfg.GetRebalanceMethod(tokenData.ChainID, tokenData.Addr.Hex())
				if tokenErr != nil {
					logger.Errorf("could not get token rebalance method: %v", tokenErr)
					continue
				}
				if tokenMethod != method {
					continue
				}

				// assign dest / origin metadata based on min / max balances
				if destTokenData == nil || tokenData.Balance.Cmp(destTokenData.Balance) < 0 {
					destTokenData = tokenData
				}
				if originTokenData == nil || tokenData.Balance.Cmp(originTokenData.Balance) > 0 {
					originTokenData = tokenData
				}
			}
		}
	}

	// if the given chain is not the origin of the rebalance, no need to do anything
	if originTokenData == nil {
		span.SetAttributes(attribute.Bool("no_rebalance_origin", true))
		return nil, nil
	}
	if destTokenData == nil {
		span.SetAttributes(attribute.Bool("no_rebalance_dest", true))
		return nil, nil
	}
	if originTokenData.ChainID != chainID {
		span.SetAttributes(attribute.Int("rebalance_origin", originTokenData.ChainID))
		return nil, nil
	}

	// get the initialPct for the origin chain
	initialPct, err := cfg.GetInitialBalancePct(originTokenData.ChainID, originTokenData.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get initial pct: %w", err)
	}

	// calculate maintenance threshold relative to total balance
	totalBalance := big.NewInt(0)
	for _, tokenMap := range tokens {
		for _, tokenData := range tokenMap {
			if tokenData.Name == rebalanceTokenData.Name {
				totalBalance.Add(totalBalance, tokenData.Balance)
			}
		}
	}
	maintenanceThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(maintenancePct/100)).Int(nil)
	if span != nil {
		span.SetAttributes(attribute.Float64("maintenance_pct", maintenancePct))
		span.SetAttributes(attribute.Float64("initial_pct", initialPct))
		span.SetAttributes(attribute.String("max_token_balance", originTokenData.Balance.String()))
		span.SetAttributes(attribute.String("min_token_balance", destTokenData.Balance.String()))
		span.SetAttributes(attribute.String("total_balance", totalBalance.String()))
		span.SetAttributes(attribute.String("maintenance_thresh", maintenanceThresh.String()))
	}

	// check if the minimum balance is below the threshold and trigger rebalance
	if destTokenData.Balance.Cmp(maintenanceThresh) > 0 {
		return rebalance, nil
	}

	// calculate the amount to rebalance vs the initial threshold on origin
	initialThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(initialPct/100)).Int(nil)
	amount := new(big.Int).Sub(originTokenData.Balance, initialThresh)

	// no need to rebalance since amount would not be positive
	if amount.Cmp(big.NewInt(0)) <= 0 {
		//nolint:nilnil
		return nil, nil
	}

	// filter the rebalance amount by the configured min
	minAmount := cfg.GetMinRebalanceAmount(originTokenData.ChainID, originTokenData.Addr)
	if amount.Cmp(minAmount) < 0 {
		// no need to rebalance
		//nolint:nilnil
		return nil, nil
	}

	// clip the rebalance amount by the configured max
	maxAmount := cfg.GetMaxRebalanceAmount(originTokenData.ChainID, originTokenData.Addr)
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
		OriginMetadata: originTokenData,
		DestMetadata:   destTokenData,
		Amount:         amount,
		Method:         method,
	}
	return rebalance, nil
}
