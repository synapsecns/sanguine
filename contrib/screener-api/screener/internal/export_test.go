package internal

type TestRulesetManager interface {
	RulesetManager
	Rulesets() map[string]map[string]bool
}

func NewTestRulesetManager(rulesets map[string]map[string]bool) TestRulesetManager {
	return NewRulesetManager(rulesets).(TestRulesetManager)
}

func (rm *rulesetManager) Rulesets() map[string]map[string]bool {
	return rm.rulesets
}

var _ TestRulesetManager = &rulesetManager{}
