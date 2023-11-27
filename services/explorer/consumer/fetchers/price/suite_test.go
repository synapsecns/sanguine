package price_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type TokenDataSuite struct {
	*testsuite.TestSuite
}

func NewTokenDataSuite(tb testing.TB) *TokenDataSuite {
	tb.Helper()
	return &TokenDataSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestTokenDataSuite(t *testing.T) {
	suite.Run(t, NewTokenDataSuite(t))
}
