// Package internal provides internal functionality for the screener-api.
package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
)

// rulesetManager manages the rulesets.
type rulesetManager struct {
	rulesets map[string]map[string]bool
}

// RulesetManager interface defines methods to work with rulesets.
type RulesetManager interface {
	// GetRuleset returns a RuleSet for the specified caller type.
	GetRuleset(rulesetName string) RuleSet
	// AddRuleset adds a new ruleset to the manager.
	AddRuleset(rulesetName string, rules map[string]bool) error
}

func (rm *rulesetManager) AddRuleset(rulesetName string, rules map[string]bool) error {
	if _, ok := rm.rulesets[rulesetName]; ok {
		return fmt.Errorf("ruleset %s already exists", rulesetName)
	}
	rm.rulesets[rulesetName] = rules
	return nil
}

// NewRulesetManager creates a new rulesetManager with the given rulesets.
func NewRulesetManager(rulesets map[string]map[string]bool) RulesetManager {
	if rulesets == nil {
		rulesets = make(map[string]map[string]bool)
	}
	return &rulesetManager{
		rulesets: rulesets,
	}
}

// GetRuleset returns a RuleSet for the specified caller type.
func (rm *rulesetManager) GetRuleset(callerType string) RuleSet {
	riskRules, exists := rm.rulesets[callerType]
	if !exists {
		return nil // or handle this case as per your application logic
	}
	return NewRuleset(riskRules)
}

// RuleSet interface defines methods to work with risk rules.
type RuleSet interface {
	HasRisk(riskType string) bool
	HasAddressIndicators(thresholds []config.VolumeThreshold, riskIndicators ...trmlabs.AddressRiskIndicator) (bool, error)
}

// CallerRuler implements the RuleSet interface for a specific caller type.
type CallerRuler struct {
	riskRules map[string]bool
}

// NewRuleset creates a new CallerRuler with the given risk rules.
func NewRuleset(riskRules map[string]bool) *CallerRuler {
	return &CallerRuler{
		riskRules: riskRules,
	}
}

// HasRisk checks if the specified risk type is present.
func (cr *CallerRuler) HasRisk(riskType string) bool {
	return cr.riskRules[riskType]
}

// HasAddressIndicators returns a list of addressRiskIndicator.
//
//nolint:cyclop
func (cr *CallerRuler) HasAddressIndicators(thresholds []config.VolumeThreshold, riskIndicators ...trmlabs.AddressRiskIndicator) (bool, error) {
	// Initialize a variable to track if any indicator is blocked
	anyIndicatorBlocked := false

	for _, ri := range riskIndicators {
		incoming, err := strconv.ParseFloat(ri.IncomingVolumeUsd, 32)
		if err != nil {
			return false, fmt.Errorf("could not parse incoming volume: %w", err)
		}

		outgoing, err := strconv.ParseFloat(ri.OutgoingVolumeUsd, 32)
		if err != nil {
			return false, fmt.Errorf("could not parse outgoing volume: %w", err)
		}

		// Check against thresholds
		for _, threshold := range thresholds {
			if strings.EqualFold(ri.Category, threshold.Category) && strings.EqualFold(ri.RiskType, threshold.TypeOfRisk) {
				// If either incoming or outgoing volume exceeds the threshold, the indicator is blocked
				if (threshold.Incoming > 0 && incoming > threshold.Incoming) || (threshold.Outgoing > 0 && outgoing > threshold.Outgoing) {
					anyIndicatorBlocked = true
					break // No need to check other thresholds, this indicator is blocked
				}
			}
		}

		if anyIndicatorBlocked {
			break // No need to check further indicators, at least one indicator is blocked
		}
	}

	// Return true if any indicator is blocked, otherwise false
	return anyIndicatorBlocked, nil
}

// MakeParam creates a risk param from the given category and risk type in a standardized format.
func MakeParam(category string, riskType string) string {
	return strings.ToLower(fmt.Sprintf("%s_%s", category, riskType))
}
