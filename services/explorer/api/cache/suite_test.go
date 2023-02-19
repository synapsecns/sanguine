package cache_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type APICacheSuite struct {
	*testsuite.TestSuite
}

func NewAPICacheSuite(tb testing.TB) *APICacheSuite {
	tb.Helper()

	return &APICacheSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestAPICacheDataSuite(t *testing.T) {
	suite.Run(t, NewAPICacheSuite(t))
}
