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
		rebalances[tokenName] = getBestRebalance(methodCandidatesSlice)
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

// getRebalanceForMethod gets the best rebalance action for a given rebalance method.
//
//nolint:nilnil
func getRebalanceForMethod(ctx context.Context, cfg relconfig.Config, inv map[int]map[common.Address]*TokenMetadata, method relconfig.RebalanceMethod, tokenName string) (rebalance *RebalanceData, err error) {
	candidateChains, err := getCandidateChains(cfg, inv, method, tokenName)
	if err != nil {
		return nil, fmt.Errorf("could not get candidate chains: %w", err)
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

	rebalance = getBestRebalance(rebalanceCandidates)
	if rebalance != nil {
		rebalance.Amount, err = getRebalanceAmount(ctx, cfg, inv, rebalance)
		if err != nil {
			return nil, fmt.Errorf("could not get rebalance amount: %w", err)
		}
		if rebalance.Amount == nil {
			return nil, nil
		}
	}

	return rebalance, nil
}

// getCandidateChains gets the respective token metadata for each chain that supports the rebalance method.
func getCandidateChains(cfg relconfig.Config, inv map[int]map[common.Address]*TokenMetadata, method relconfig.RebalanceMethod, tokenName string) (map[int]*TokenMetadata, error) {
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

	return candidateChains, nil
}

func getBestRebalance(candidates []RebalanceData) (best *RebalanceData) {
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

	return best
}

// getRebalanceAmount calculates the amount to rebalance based on the configured thresholds.
//
//nolint:cyclop,nilnil
func getRebalanceAmount(ctx context.Context, cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, rebalance *RebalanceData) (amount *big.Int, err error) {
	span := trace.SpanFromContext(ctx)

	// get the maintenance and initial values for the destination chain
	maintenancePct, err := cfg.GetMaintenanceBalancePct(rebalance.DestMetadata.ChainID, rebalance.DestMetadata.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get maintenance pct: %w", err)
	}
	initialPct, err := cfg.GetInitialBalancePct(rebalance.DestMetadata.ChainID, rebalance.DestMetadata.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get initial pct: %w", err)
	}

	// calculate maintenance threshold relative to total balance
	totalBalance, err := getTotalBalance(cfg, tokens, rebalance.OriginMetadata.Name, rebalance.Method)
	if err != nil {
		return nil, fmt.Errorf("could not get total balance: %w", err)
	}
	maintenanceThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(maintenancePct/100)).Int(nil)
	if span != nil {
		span.SetAttributes(attribute.Float64("maintenance_pct", maintenancePct))
		span.SetAttributes(attribute.Float64("initial_pct", initialPct))
		span.SetAttributes(attribute.String("max_token_balance", rebalance.OriginMetadata.Balance.String()))
		span.SetAttributes(attribute.String("min_token_balance", rebalance.DestMetadata.Balance.String()))
		span.SetAttributes(attribute.String("total_balance", totalBalance.String()))
		span.SetAttributes(attribute.String("maintenance_thresh", maintenanceThresh.String()))
	}

	// no need to rebalance if we are not below maintenance threshold on destination
	if rebalance.DestMetadata.Balance.Cmp(maintenanceThresh) > 0 {
		fmt.Println("returning nil")
		return nil, nil
	}

	// calculate the amount to rebalance vs the initial threshold on destination
	initialThresh, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(initialPct/100)).Int(nil)
	amount = new(big.Int).Sub(rebalance.OriginMetadata.Balance, initialThresh)

	// no need to rebalance since amount would not be positive
	if amount.Cmp(big.NewInt(0)) <= 0 {
		//nolint:nilnil
		return nil, nil
	}

	// filter the rebalance amount by the configured min
	minAmount := cfg.GetMinRebalanceAmount(rebalance.OriginMetadata.ChainID, rebalance.OriginMetadata.Addr)
	if amount.Cmp(minAmount) < 0 {
		// no need to rebalance
		//nolint:nilnil
		return nil, nil
	}

	// clip the rebalance amount by the configured max
	maxAmount := cfg.GetMaxRebalanceAmount(rebalance.OriginMetadata.ChainID, rebalance.OriginMetadata.Addr)
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
	maintenancePctOrigin, err := cfg.GetMaintenanceBalancePct(rebalance.OriginMetadata.ChainID, rebalance.OriginMetadata.Addr.Hex())
	if err != nil {
		return nil, fmt.Errorf("could not get maintenance pct: %w", err)
	}
	maintenanceThreshOrigin, _ := new(big.Float).Mul(new(big.Float).SetInt(totalBalance), big.NewFloat(maintenancePctOrigin/100)).Int(nil)
	newBalanceOrigin := new(big.Int).Sub(rebalance.OriginMetadata.Balance, amount)
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

// getTotalBalance calculates the total balance for a token
// across all chains that support the given rebalance method.
func getTotalBalance(cfg relconfig.Config, tokens map[int]map[common.Address]*TokenMetadata, tokenName string, method relconfig.RebalanceMethod) (*big.Int, error) {
	totalBalance := big.NewInt(0)
	for _, tokenMap := range tokens {
		for _, tokenData := range tokenMap {
			if tokenData.Name != tokenName {
				continue
			}
			rebalanceMethods, err := cfg.GetRebalanceMethods(tokenData.ChainID, tokenData.Addr.Hex())
			if err != nil {
				return nil, fmt.Errorf("could not get rebalance methods: %w", err)
			}
			for _, m := range rebalanceMethods {
				if m == method {
					totalBalance.Add(totalBalance, tokenData.Balance)
					break
				}
			}
		}
	}
	return totalBalance, nil
}
