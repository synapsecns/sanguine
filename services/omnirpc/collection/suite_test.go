package collection_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// CollectionSuite defines the basic test suite.
type CollectionSuite struct {
	*testsuite.TestSuite
}

// NewCollectionSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewCollectionSuite(tb testing.TB) *CollectionSuite {
	tb.Helper()
	return &CollectionSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestCollectionSuite(t *testing.T) {
	suite.Run(t, NewCollectionSuite(t))
}
