package attestationcollector_test

import (
	"testing"

	"github.com/synapsecns/sanguine/agents/testutil"

	"github.com/stretchr/testify/suite"
)

// AttestationCollectorSuite is the attestation collector test suite.
type AttestationCollectorSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewAttestationCollectorSuite creates an end-to-end test suite.
func NewAttestationCollectorSuite(tb testing.TB) *AttestationCollectorSuite {
	tb.Helper()
	return &AttestationCollectorSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

// SetupTest sets up the test.
func (a *AttestationCollectorSuite) SetupTest() {
	a.SimulatedBackendsTestSuite.SetupTest()
}

// TestAttestationCollectorSuite runs the integration test suite.
func TestAttestationCollectorSuite(t *testing.T) {
	suite.Run(t, NewAttestationCollectorSuite(t))
}
