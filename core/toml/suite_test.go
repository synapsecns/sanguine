package toml_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// TomlSuite defines the basic test suite.
type TomlSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *TomlSuite {
	tb.Helper()
	return &TomlSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestTomlSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
