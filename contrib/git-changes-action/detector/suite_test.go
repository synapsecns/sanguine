package detector_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// DetectorSuite defines the basic test suite.
type DetectorSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *DetectorSuite {
	tb.Helper()
	return &DetectorSuite{
		testsuite.NewTestSuite(tb),
	}
}

func (d *DetectorSuite) SetupTest() {
	d.TestSuite.SetupTest()

}

func TestDetectorSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
