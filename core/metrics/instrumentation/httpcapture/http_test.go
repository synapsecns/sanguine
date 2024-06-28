package httpcapture_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation/httpcapture"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/core/testsuite"
	"net/http"
	"strings"
	"testing"
)

type HTTPCaptureSuite struct {
	*testsuite.TestSuite
	server *gin.Engine
	port   int
}

// NewCaptureSuite creates a new HTTPCaptureSuite.
func NewCaptureSuite(t *testing.T) *HTTPCaptureSuite {
	t.Helper()
	return &HTTPCaptureSuite{
		TestSuite: testsuite.NewTestSuite(t),
	}
}

func (t *HTTPCaptureSuite) SetupTest() {
	t.TestSuite.SetupTest()
	var logger = log.Logger("httpcapture-test")
	t.server = ginhelper.New(logger)
	t.server.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})

	var err error
	t.port, err = freeport.GetFreePort()
	t.Require().NoError(err)

	go func() {
		connection := baseServer.Server{}
		_ = connection.ListenAndServe(t.GetTestContext(), fmt.Sprintf(":%d", t.port), t.server)
	}()

	err = ginhelper.WaitForStart(t.GetTestContext(), t.port)
	t.Require().NoError(err)
}

func (t *HTTPCaptureSuite) TestCaptureTransport() {
	httpClient := new(http.Client)

	mocktracer := metrics.NewTestTracer(t.GetTestContext(), t.T())

	httpClient.Transport = httpcapture.NewCaptureTransport(httpClient.Transport, mocktracer)

	const testRequestBody = "hi"
	req, err := http.NewRequestWithContext(t.GetTestContext(), http.MethodPost, fmt.Sprintf("http://localhost:%d/test", t.port), strings.NewReader(testRequestBody))
	t.Require().NoError(err)

	res, err := httpClient.Do(req)
	defer func() {
		if res != nil && res.Body != nil {
			_ = res.Body.Close()
			_ = req.Body.Close()
		}
	}()
	t.Require().NoError(err)
	t.Require().NotNil(res)

	span := mocktracer.GetSpansByName(httpcapture.RequestSpanName)[0]

	reqBody := metrics.SpanEventByName(span, httpcapture.RequestEventName)
	t.Equal(reqBody.AsString(), testRequestBody)

	respBody := metrics.SpanEventByName(span, httpcapture.ResponseEventName)
	t.Equal(respBody.AsString(), `{"message":"hello world"}`)
}

func TestCaptureSuite(t *testing.T) {
	suite.Run(t, NewCaptureSuite(t))
}
