package wallet_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type WalletSuite struct {
	*testutils.TestSuite
}

// NewWalletSuite creates a end-to-end test suite.
func NewWalletSuite(tb testing.TB) *WalletSuite {
	tb.Helper()
	return &WalletSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestWalletSuite(t *testing.T) {
	suite.Run(t, NewWalletSuite(t))
}
