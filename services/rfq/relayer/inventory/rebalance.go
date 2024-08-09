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

// getRebalances gets the best rebalance action for each token.
func getRebalances(ctx context.Context, cfg relconfig.Config, inv map[int]map[common.Address]*TokenMetadata) (rebalances map[string]*RebalanceData, err error) {
	rebalances = map[string]*RebalanceData{}

	rebalanceCandidates, err := getRebalanceCandidates(ctx, cfg, inv)
	if err != nil {
		return nil, fmt.Errorf("could not get rebalance candidates: %w", err)
	}

	for tokenName, methodCandidates := range rebalanceCandidates {
		methodCandidatesSlice := []RebalanceData{}
		for _, candidate := range methodCandidates {
			if candidate == nil {
				continue
			}
			methodCandidatesSlice = append(methodCandidatesSlice, *candidate)
		}
		rebalances[tokenName], err = getBestRebalance(ctx, methodCandidatesSlice)
		if err != nil {
			return nil, fmt.Errorf("could not get best rebalance: %w", err)
		}
	}

	return rebalances, nil
}

// getRebalanceCandidates gets the best rebalance for each token and rebalance method supported by the config.
func getRebalanceCandidates(ctx context.Context, cfg relconfig.Config, inv map[int]map[common.Address]*TokenMetadata) (rebalances map[string]map[relconfig.RebalanceMethod]*RebalanceData, err error) {
	rebalances = map[string]map[relconfig.RebalanceMethod]*RebalanceData{}
	for chainID, chainCfg := range cfg.Chains {
		for tokenName, tokenCfg := range chainCfg.Tokens {
			if len(tokenCfg.RebalanceMethods) == 0 {
				continue
			}

			_, ok := rebalances[tokenName]
			if !ok {
				rebalances[tokenName] = map[relconfig.RebalanceMethod]*RebalanceData{}
			}

			methods, err := cfg.GetRebalanceMethods(chainID, tokenCfg.Address)
			if err != nil {
				return nil, fmt.Errorf("could not get rebalance methods: %w", err)
			}
			for _, method := range methods {
				rebalances[tokenName][method], err = getRebalanceForMethod(ctx, cfg, inv, method, tokenName)
				if err != nil {
					return nil, fmt.Errorf("could not get rebalance for method %s: %w", method.String(), err)
				}
			}
		}
	}

	return rebalances, nil
}

func getRebalanceForMethod(ctx context.Context, cfg relconfig.Config, inv map[int]map[common.Address]*TokenMetadata, method relconfig.RebalanceMethod, tokenName string) (rebalance *RebalanceData, err error) {
	candidateChains := map[int]*TokenMetadata{}
	for chainID, chainCfg := range cfg.Chains {
		var validCandidate bool
		var candidateMetadata *TokenMetadata
		for name, tokenCfg := range chainCfg.Tokens {
			if name != tokenName {
				continue
			}

			// check that the token supports given rebalance method
			supportedMethods, err := cfg.GetRebalanceMethods(chainID, tokenCfg.Address)
			if err != nil {
				return nil, fmt.Errorf("could not get rebalance methods: %w", err)
			}
			var supported bool
			for _, m := range supportedMethods {
				if m == method {
					supported = true
					break
				}
			}
			if supported {
				validCandidate = true
				candidateMetadata = inv[chainID][common.HexToAddress(tokenCfg.Address)]
				if candidateMetadata == nil {
					return nil, fmt.Errorf("could not get token metadata for chain %d and addr %s", chainID, tokenCfg.Address)
				}
				break
			}
		}
		if validCandidate {
			candidateChains[chainID] = candidateMetadata
		}
	}

	// now we have candidate chains, produce the rebalance data for each permutation of the chains
	rebalanceCandidates := []RebalanceData{}
	for i := range candidateChains {
		for j := range candidateChains {
			if i == j {
				continue
			}

			candidate := RebalanceData{
				OriginMetadata: candidateChains[i],
				DestMetadata:   candidateChains[j],
				Method:         method,
			}
			rebalanceCandidates = append(rebalanceCandidates, candidate)
		}
	}

	rebalance, err = getBestRebalance(ctx, rebalanceCandidates)
	if err != nil {
		return nil, fmt.Errorf("could not get best rebalance: %w", err)
	}

	if rebalance != nil {
		rebalance.Amount, err = getRebalanceAmount(ctx, cfg, inv, rebalance.OriginMetadata, rebalance.DestMetadata)
		if err != nil {
			return nil, fmt.Errorf("could not get rebalance amount: %w", err)
		}
	}

	return rebalance, nil
}

func getBestRebalance(ctx context.Context, candidates []RebalanceData) (best *RebalanceData, err error) {
	var maxDelta *big.Int
	best = nil

	for _, candidate := range candidates {
		originBalance := candidate.OriginMetadata.Balance
		destBalance := candidate.DestMetadata.Balance

		delta := new(big.Int).Sub(originBalance, destBalance)

		if maxDelta == nil || delta.Cmp(maxDelta) > 0 {
			maxDelta = delta
			candidateCopy := candidate
			best = &candidateCopy
		}
	}

	return best, nil
}

// getRebalanceAmount calculates the amount to rebalance based on the configured thresholds.
//
//nolint:cyclop,nilnil
func getRebalanceAmount(ctx context.Context, cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, originTokenData, destTokenData *TokenMetadata) (amount *big.Int, err error) {
	span := trace.SpanFromContext(ctx)

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
