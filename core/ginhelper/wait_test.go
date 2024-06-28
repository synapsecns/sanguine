package ginhelper_test

import (
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/ginhelper"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
	"time"
)

// WaitSuite defines the basic test suite.
type WaitSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewUtilSuite(tb testing.TB) *WaitSuite {
	tb.Helper()
	return &WaitSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, NewUtilSuite(t))
}

func (t *WaitSuite) TearDownTest() {
	ginhelper.ResetServerTimeout()
}

func (t *WaitSuite) TestWaitForStartFail() {
	port, err := freeport.GetFreePort()
	t.Require().NoError(err)

	ginhelper.SetServerTimeout(time.Millisecond)

	// make sure this errors on non existent server
	t.Require().NotNil(ginhelper.WaitForStart(t.GetTestContext(), port))
}

// make sure a working server passes.
func (t *WaitSuite) TestWaitForStartSucceed() {
	testLogger := log.Logger("test")

	tmpPort, err := freeport.GetFreePort()
	t.Require().NoError(err)

	router := ginhelper.New(testLogger)

	// start a server
	go func() {
		connection := baseServer.Server{}
		_ = connection.ListenAndServe(t.GetTestContext(), fmt.Sprintf(":%d", tmpPort), router)
	}()

	err = ginhelper.WaitForStart(t.GetTestContext(), tmpPort)
	t.Require().NoError(err)
}
