package util_test

import (
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/gql/util"
	"github.com/synapsecns/sanguine/core/ginhelper"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// UtilSuite defines the basic test suite.
type UtilSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewUtilSuite(tb testing.TB) *UtilSuite {
	tb.Helper()
	return &UtilSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, NewUtilSuite(t))
}

func (t *UtilSuite) TestWaitForStartFail() {
	port, err := freeport.GetFreePort()
	t.Require().NoError(err)

	// make sure this errors on non existent server
	t.Require().NotNil(util.WaitForStart(t.GetTestContext(), port))
}

// make sure a working server passes.
func (t *UtilSuite) TestWaitForStartSucceed() {
	testLogger := log.Logger("test")

	tmpPort, err := freeport.GetFreePort()
	t.Require().NoError(err)

	router := ginhelper.New(testLogger)

	// start a server
	go func() {
		connection := baseServer.Server{}
		_ = connection.ListenAndServe(t.GetTestContext(), fmt.Sprintf(":%d", tmpPort), router)
	}()

	err = util.WaitForStart(t.GetTestContext(), tmpPort)
	t.Require().NoError(err)
}
