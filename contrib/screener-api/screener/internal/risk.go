// Package internal provides internal functionality for the screener-api.
package internal

import (
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"strconv"
	"strings"
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
	HasAddressIndicators(riskIndicators ...trmlabs.AddressRiskIndicator) (bool, error)
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
func (cr *CallerRuler) HasAddressIndicators(riskIndicators ...trmlabs.AddressRiskIndicator) (bool, error) {
	for _, ri := range riskIndicators {
		incoming, err := strconv.ParseFloat(ri.IncomingVolumeUsd, 32)
		if err != nil {
			return false, fmt.Errorf("could not parse incoming volume: %w", err)
		}

		outgoing, err := strconv.ParseFloat(ri.OutgoingVolumeUsd, 32)
		if err != nil {
			return false, fmt.Errorf("could not parse outgoing volume: %w", err)
		}

		riskParam := MakeParam(ri.Category, ri.RiskType)
		fmt.Println("param is", riskParam)
		isBlocked, found := cr.riskRules[riskParam]
		if isBlocked && found && (incoming > 0 || outgoing > 0) {
			return true, nil
		}
	}
	fmt.Println("risk rules are", cr.riskRules)

	return false, nil
}

func MakeParam(category string, riskType string) string {
	return strings.ToLower(fmt.Sprintf("%s_%s", category, riskType))
}
