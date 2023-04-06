package ginhelper_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/ginhelper"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// GinHelperSuite defines the basic test suite.
type GinHelperSuite struct {
	*testsuite.TestSuite
	url    string
	logger *log.ZapEventLogger
}

func (g *GinHelperSuite) SetupTest() {
	g.TestSuite.SetupTest()
	g.logger = log.Logger(fmt.Sprintf("test-%d-logger", g.GetTestID()))

	testServer := ginhelper.New(g.logger)
	freePort, err := freeport.GetFreePort()
	Nil(g.T(), err)

	g.url = fmt.Sprintf("http://localhost:%d", freePort)

	go func() {
		connection := baseServer.Server{}
		err = connection.ListenAndServe(g.GetTestContext(), fmt.Sprintf(":%d", freePort), testServer)
		// we expect context cancellation errors at the end of the test
		if !errors.Is(err, context.Canceled) {
			Nil(g.T(), err)
		}
	}()
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *GinHelperSuite {
	tb.Helper()
	return &GinHelperSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		url:       "",
	}
}

func TestGinHelperSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
