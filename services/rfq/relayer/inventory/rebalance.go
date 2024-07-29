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

// getRebalance builds a rebalance action based on current token balances and configured thresholds.
// Note that only the given chain/token pair is considered for rebalance (as the destination chain).
//
//nolint:cyclop,nilnil
func getRebalance(span trace.Span, cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, chainID int, token common.Address) (rebalance *RebalanceData, err error) {
	// get rebalance method
	methods, err := cfg.GetRebalanceMethods(chainID, token.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get rebalance method: %w", err)
	}
	if len(methods) == 0 {
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
	originTokenData, destTokenData, method := getRebalanceMetadatas(cfg, tokens, rebalanceTokenData.Name, methods)
	if originTokenData == nil {
		if span != nil {
			span.SetAttributes(attribute.Bool("no_rebalance_origin", true))
		}
		return nil, nil
	}
	if destTokenData == nil {
		if span != nil {
			span.SetAttributes(attribute.Bool("no_rebalance_dest", true))
		}
		return nil, nil
	}
	if method == relconfig.RebalanceMethodNone {
		if span != nil {
			span.SetAttributes(attribute.Bool("no_rebalance_method", true))
		}
		return nil, nil
	}

	// if the given chain is not the destination of the rebalance, no need to do anything
	if destTokenData.ChainID != chainID {
		if span != nil {
			span.SetAttributes(attribute.Int("rebalance_dest", destTokenData.ChainID))
		}
		return nil, nil
	}

	amount, err := getRebalanceAmount(span, cfg, tokens, originTokenData, destTokenData)
	if err != nil {
		return nil, fmt.Errorf("could not get rebalance amount: %w", err)
	}
	if amount == nil {
		if span != nil {
			span.SetAttributes(attribute.Bool("no_rebalance_amount", true))
		}
		return nil, nil
	}

	rebalance = &RebalanceData{
		OriginMetadata: originTokenData,
		DestMetadata:   destTokenData,
		Amount:         amount,
		Method:         method,
	}
	return rebalance, nil
}

// getRebalanceMetadatas finds the origin and dest token metadata based on the configured rebalance method.
//
//nolint:nestif,cyclop
func getRebalanceMetadatas(cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, tokenName string, methods []relconfig.RebalanceMethod) (originTokenData, destTokenData *TokenMetadata, method relconfig.RebalanceMethod) {
	for _, method := range methods {
		for _, tokenMap := range tokens {
			for _, tokenData := range tokenMap {
				if tokenData.Name == tokenName {
					if !isTokenCompatible(tokenData, method, cfg) {
						continue
					}

					// assign origin / dest metadata based on min / max balances
					if originTokenData == nil || tokenData.Balance.Cmp(originTokenData.Balance) > 0 {
						originTokenData = tokenData
					}
					if destTokenData == nil || tokenData.Balance.Cmp(destTokenData.Balance) < 0 {
						destTokenData = tokenData
					}
				}
			}
		}
		if originTokenData != nil && destTokenData != nil {
			return originTokenData, destTokenData, method
		}
	}
	return nil, nil, relconfig.RebalanceMethodNone
}

func isTokenCompatible(tokenData *TokenMetadata, method relconfig.RebalanceMethod, cfg relconfig.Config) bool {
	// make sure that the token is compatible with our rebalance method
	tokenMethods, tokenErr := cfg.GetRebalanceMethods(tokenData.ChainID, tokenData.Addr.Hex())
	if tokenErr != nil {
		logger.Errorf("could not get token rebalance method: %v", tokenErr)
		return false
	}

	isCompatible := false
	for _, tm := range tokenMethods {
		if tm == method {
			isCompatible = true
		}
	}
	return isCompatible
}

// getRebalanceAmount calculates the amount to rebalance based on the configured thresholds.
//
//nolint:cyclop,nilnil
func getRebalanceAmount(span trace.Span, cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, originTokenData, destTokenData *TokenMetadata) (amount *big.Int, err error) {
	// get the maintenance and initial values for the destination chain
	maintenancePct, err := cfg.GetMaintenanceBalancePct(destTokenData.ChainID, destTokenData.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get maintenance pct: %w", err)
	}
	initialPct, err := cfg.GetInitialBalancePct(destTokenData.ChainID, destTokenData.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get initial pct: %w", err)
	}

	// calculate maintenance threshold relative to total balance
	tokenName := originTokenData.Name
	totalBalance := big.NewInt(0)
	for _, tokenMap := range tokens {
		for _, tokenData := range tokenMap {
			if tokenData.Name == tokenName {
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

	// no need to rebalance if we are not below maintenance threshold on destination
	if destTokenData.Balance.Cmp(maintenanceThresh) > 0 {
		return nil, nil
	}

	// calculate the amount to rebalance vs the initial threshold on destination
	initialThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(initialPct/100)).Int(nil)
	amount = new(big.Int).Sub(originTokenData.Balance, initialThresh)

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

	// make sure that the rebalance amount does not take origin below maintenance threshold
	maintenancePctOrigin, err := cfg.GetMaintenanceBalancePct(originTokenData.ChainID, originTokenData.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get maintenance pct: %w", err)
	}
	maintenanceThreshOrigin, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(maintenancePctOrigin/100)).Int(nil)
	newBalanceOrigin := new(big.Int).Sub(originTokenData.Balance, amount)
	if newBalanceOrigin.Cmp(maintenanceThreshOrigin) < 0 {
		if span != nil {
			span.SetAttributes(
				attribute.Float64("maintenance_pct_origin", maintenancePctOrigin),
				attribute.String("maintenance_thresh_origin", maintenanceThreshOrigin.String()),
				attribute.String("new_balance_origin", newBalanceOrigin.String()),
			)
		}
		return nil, nil
	}

	return amount, nil
}
