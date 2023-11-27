package swap_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type SwapFetcherSuite struct {
	*testsuite.TestSuite
}

func NewSwapFetcherSuite(tb testing.TB) *SwapFetcherSuite {
	tb.Helper()
	return &SwapFetcherSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestSwapFetcherSuite(t *testing.T) {
	suite.Run(t, NewSwapFetcherSuite(t))
}
