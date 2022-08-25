package wallet_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type WalletSuite struct {
	*testsuite.TestSuite
}

// NewWalletSuite creates a end-to-end test suite.
func NewWalletSuite(tb testing.TB) *WalletSuite {
	tb.Helper()
	return &WalletSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestWalletSuite(t *testing.T) {
	suite.Run(t, NewWalletSuite(t))
}
