package ginhelper_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics/logger"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// GinHelperSuite defines the basic test suite.
type GinHelperSuite struct {
	*testsuite.TestSuite
	urls   []urlTest
	logger *log.ZapEventLogger
}

type urlTest struct {
	url      string
	testType string
}

func (g *GinHelperSuite) SetupTest() {
	g.urls = []urlTest{}

	useExperimental := []bool{true, false}
	for _, useExp := range useExperimental {
		g.TestSuite.SetupTest()
		g.logger = log.Logger(fmt.Sprintf("test-%d-logger", g.GetTestID()))

		var testServer *gin.Engine
		if useExp {
			testServer = ginhelper.NewWithExperimentalLogger(g.GetTestContext(), logger.NewNullLogger())
		} else {
			testServer = ginhelper.New(g.logger)
		}
		freePort, err := freeport.GetFreePort()
		Nil(g.T(), err)

		g.urls = append(g.urls, urlTest{
			url:      fmt.Sprintf("http://localhost:%d", freePort),
			testType: fmt.Sprintf("experimental: %t", useExp),
		})

		go func() {
			connection := baseServer.Server{}
			err = connection.ListenAndServe(g.GetTestContext(), fmt.Sprintf(":%d", freePort), testServer)
			// we expect context cancellation errors at the end of the test
			if !errors.Is(err, context.Canceled) {
				Nil(g.T(), err)
			}
		}()
	}
}

func (g *GinHelperSuite) runOnAllURLs(f func(url string)) {
	for _, url := range g.urls {
		g.T().Run(url.testType, func(t *testing.T) {
			f(url.url)
		})
	}
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *GinHelperSuite {
	tb.Helper()
	return &GinHelperSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestGinHelperSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
