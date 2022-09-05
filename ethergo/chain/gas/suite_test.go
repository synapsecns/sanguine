package gas_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// GasSuite defines the basic chain suite.
type GasSuite struct {
	*testsuite.TestSuite
}

// NewGasSuite creates a new chain testing suite.
func NewGasSuite(tb testing.TB) *GasSuite {
	tb.Helper()
	return &GasSuite{testsuite.NewTestSuite(tb)}
}

func TestContractGasSuite(t *testing.T) {
	suite.Run(t, NewGasSuite(t))
}
