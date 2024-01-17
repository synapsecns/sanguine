package internal_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener/internal"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"testing"
)

func TestNewRulesetManager(t *testing.T) {
	rulesets := make(map[string]map[string]bool)
	rulesets["caller1"] = map[string]bool{"risk1": true}

	rm := internal.NewTestRulesetManager(rulesets)
	assert.NotNil(t, rm)
	assert.Equal(t, 1, len(rm.Rulesets()))
}

func TestAddRuleset(t *testing.T) {
	rm := internal.NewTestRulesetManager(nil)

	// Add a new ruleset
	err := rm.AddRuleset("newRuleset", map[string]bool{"risk1": true})
	assert.Nil(t, err)

	// Try to add a ruleset that already exists
	err = rm.AddRuleset("newRuleset", map[string]bool{"risk2": true})
	assert.NotNil(t, err)
}

func TestGetRuleset(t *testing.T) {
	rulesets := make(map[string]map[string]bool)
	rulesets["existing"] = map[string]bool{"risk1": true}

	rm := internal.NewTestRulesetManager(rulesets)

	// Get an existing ruleset
	rs := rm.GetRuleset("existing")
	assert.NotNil(t, rs)

	// Get a non-existing ruleset
	rs = rm.GetRuleset("nonExisting")
	assert.Nil(t, rs)
}

func TestHasAddressIndicators(t *testing.T) {
	riskRules := map[string]bool{
		"category1.risktype1": true,
	}

	cr := internal.NewRuleset(riskRules)

	// Test case where the indicator meets risk rules
	indicators := []trmlabs.AddressRiskIndicator{
		{IncomingVolumeUsd: "1000", OutgoingVolumeUsd: "500", Category: "Category1", RiskType: "RiskType1"},
	}
	result, err := cr.HasAddressIndicators(indicators...)
	assert.Nil(t, err)
	assert.True(t, result)

	// Test case where the indicator does not meet risk rules
	indicators = []trmlabs.AddressRiskIndicator{
		{IncomingVolumeUsd: "100", OutgoingVolumeUsd: "50", Category: "Category2", RiskType: "RiskType2"},
	}
	result, err = cr.HasAddressIndicators(indicators...)
	assert.Nil(t, err)
	assert.False(t, result)

	// Test case with invalid incoming volume
	indicators = []trmlabs.AddressRiskIndicator{
		{IncomingVolumeUsd: "invalid", OutgoingVolumeUsd: "500", Category: "Category1", RiskType: "RiskType1"},
	}
	_, err = cr.HasAddressIndicators(indicators...)
	assert.NotNil(t, err)

	// Test case with invalid outgoing volume
	indicators = []trmlabs.AddressRiskIndicator{
		{IncomingVolumeUsd: "1000", OutgoingVolumeUsd: "invalid", Category: "Category1", RiskType: "RiskType1"},
	}
	_, err = cr.HasAddressIndicators(indicators...)
	assert.NotNil(t, err)
}
