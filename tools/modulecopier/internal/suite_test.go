package internal_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type GeneratorSuite struct {
	*testsuite.TestSuite
}

// NewGeneratorSuite creates a end-to-end test suite.
func NewGeneratorSuite(tb testing.TB) *GeneratorSuite {
	tb.Helper()
	return &GeneratorSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestGeneratorSuite(t *testing.T) {
	suite.Run(t, NewGeneratorSuite(t))
}
