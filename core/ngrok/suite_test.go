package ngrok_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// NgrokSuite defines the basic test suite.
type NgrokSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewNgrokSuite(tb testing.TB) *NgrokSuite {
	tb.Helper()
	return &NgrokSuite{
		testsuite.NewTestSuite(tb),
	}
}

func (n *NgrokSuite) SetupSuite() {
	n.TestSuite.SetupSuite()

}

func TestNgrokSuite(t *testing.T) {
	suite.Run(t, NewNgrokSuite(t))
}
