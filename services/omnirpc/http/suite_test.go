package http_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/omnirpc/http"
	"testing"
)

// clientSuite defines the basic test suite.
type HTTPSuite struct {
	*testsuite.TestSuite
	clients []http.Client
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewClientSuite(tb testing.TB) *HTTPSuite {
	tb.Helper()
	return &HTTPSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c *HTTPSuite) SetupTest() {
	c.TestSuite.SetupTest()

	for _, clientType := range http.AllClientTypes {
		c.clients = append(c.clients, http.NewClient(clientType))
	}
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, NewClientSuite(t))
}

func (c *HTTPSuite) MockHeaders(count int) (headers map[string]string) {
	headers = make(map[string]string)
	for i := 0; i < count; i++ {
		headers[gofakeit.FirstName()] = gofakeit.Sentence(10)
	}
	return headers
}
