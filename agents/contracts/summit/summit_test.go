package summit_test

import (
	. "github.com/stretchr/testify/assert"
)

//nolint:unused
func (a SummitSuite) launchTest(amountGuards, amountNotaries int) {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	GreaterOrEqual(a.T(), amountGuards+amountNotaries, 1)
	LessOrEqual(a.T(), amountGuards, 1)
	LessOrEqual(a.T(), amountNotaries, 1)
}

// TestAttestationCollectorSuite tests submitting an attesation with one guard and one notary.
func (a SummitSuite) TestSubmitAttestationOneGuardOneNotary() {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	a.launchTest(1, 1)
}

// TestSubmitAttestationOnlyOneNotary tests submitting an attesation with only one notary.
func (a SummitSuite) TestSubmitAttestationOnlyOneNotary() {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	a.launchTest(0, 1)
}

// TestSubmitAttestationOnlyOneGuard tests submitting an attesation with only one guard.
func (a SummitSuite) TestSubmitAttestationOnlyOneGuard() {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	a.launchTest(1, 0)
}
