package internal_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type GeneratorSuite struct {
	*testutils.TestSuite
}

// NewGeneratorSuite creates a end-to-end test suite.
func NewGeneratorSuite(tb testing.TB) *GeneratorSuite {
	tb.Helper()
	return &GeneratorSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestGeneratorSuite(t *testing.T) {
	suite.Run(t, NewGeneratorSuite(t))
}
