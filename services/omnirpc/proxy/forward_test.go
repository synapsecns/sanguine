package proxy_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	chainManagerMocks "github.com/synapsecns/sanguine/services/omnirpc/chainmanager/mocks"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"github.com/synapsecns/sanguine/services/omnirpc/http/mocks"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	proxyMocks "github.com/synapsecns/sanguine/services/omnirpc/proxy/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

func (p *ProxySuite) TestServeRequestNoChain() {
	prxy := proxy.NewProxy(config.Config{}, p.metrics)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	prxy.Forward(c, 1, nil)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

func (p *ProxySuite) TestCannotReadBody() {
	prxy := proxy.NewProxy(config.Config{}, p.metrics)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mockBody := new(proxyMocks.BodyReader)
	mockBody.On("Read", mock.Anything).Return(0, errors.New("could not read body"))

	c.Request = &http.Request{
		Body: mockBody,
	}

	prxy.Forward(c, 1, nil)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

func (p *ProxySuite) generateFakeJSON() []byte {
	rawBody, err := gofakeit.JSON(&gofakeit.JSONOptions{
		Type: "array",
		Fields: []gofakeit.Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
		},
		RowCount: gofakeit.Number(5, 20),
		Indent:   true,
	})
	Nil(p.T(), err)

	return rawBody
}

func (p *ProxySuite) TestMalformedRequestBody() {
	prxy := proxy.NewProxy(config.Config{}, p.metrics)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(p.generateFakeJSON()))

	prxy.Forward(c, 1, nil)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

// TestAcquireReleaseForwarder makes sure the forwarder is cleared afte r being released.
func (p *ProxySuite) TestAcquireReleaseForwarder() {
	prxy := proxy.NewProxy(config.Config{}, p.metrics)

	forwarder := prxy.AcquireForwarder()
	forwarder.SetChain(new(chainManagerMocks.Chain))
	forwarder.SetC(&gin.Context{})
	forwarder.SetClient(omniHTTP.NewClient(omniHTTP.Resty))
	forwarder.SetR(prxy)
	forwarder.SetBody(gofakeit.ImagePng(5, 5))
	forwarder.SetRequestID([]byte(uuid.New().String()))
	forwarder.SetRequiredConfirmations(gofakeit.Uint16())
	forwarder.SetBlankResMap()
	forwarder.SetRPCRequest([]rpc.Request{{
		ID:     gofakeit.Number(1, 2),
		Method: gofakeit.Word(),
	}})
	prxy.ReleaseForwarder(forwarder)

	forwarder = prxy.AcquireForwarder()
	// should be set by acquirer or recycled
	NotNil(p.T(), forwarder.R())
	Nil(p.T(), forwarder.C())
	Nil(p.T(), forwarder.Chain())
	Nil(p.T(), forwarder.Body())
	Zero(p.T(), forwarder.RequiredConfirmations())
	// should be set by acquirer
	NotNil(p.T(), forwarder.Client())
	Nil(p.T(), forwarder.ResMap())
	Nil(p.T(), forwarder.RPCRequest())
}

func (p *ProxySuite) TestForwardRequestDisallowWS() {
	prxy := proxy.NewProxy(config.Config{}, p.metrics)

	invalidSchemes := []string{"wss", "ws"}
	for _, scheme := range invalidSchemes {
		forwarder := prxy.AcquireForwarder()

		testURL := gofakeit.URL()
		parsedURL, err := url.Parse(testURL)
		Nil(p.T(), err)

		// change the scheme to use wss to see if
		parsedURL.Scheme = scheme
		testURL = parsedURL.String()

		rawRes, err := forwarder.ForwardRequest(p.GetTestContext(), testURL)
		Nil(p.T(), rawRes)
		NotNil(p.T(), err)

		prxy.ReleaseForwarder(forwarder)
	}
}

func (p *ProxySuite) TestForwardRequest() {
	prxy := proxy.NewProxy(config.Config{}, p.metrics)

	methodName := "test"
	testRes := p.MustMarshall(proxy.JSONRPCMessage{
		Version: strconv.Itoa(gofakeit.Number(1, 2)),
		Method:  methodName,
		Params:  nil,
		Error:   nil,
		Result:  nil,
	})

	captureClient := omniHTTP.NewCaptureClient(func(c *omniHTTP.CapturedRequest) (omniHTTP.Response, error) {
		bodyRes := new(mocks.Response)
		bodyRes.On("Body").Return(testRes)
		bodyRes.On("StatusCode").Return(200)
		return bodyRes, nil
	})
	prxy.SetClient(captureClient)

	testURL := gofakeit.URL()
	testRequestID := gofakeit.UUID()
	testBody := gofakeit.ImagePng(10, 10)
	forwarder := prxy.AcquireForwarder()
	forwarder.SetBody(testBody)
	forwarder.SetRequestID([]byte(testRequestID))
	forwarder.SetRPCRequest([]rpc.Request{{Method: methodName}})

	_, err := forwarder.ForwardRequest(p.GetTestContext(), testURL)
	Nil(p.T(), err)

	requests := captureClient.Requests()
	Equal(p.T(), len(requests), 1)

	request := requests[0]

	Equal(p.T(), request.RequestURI, testURL)
	idHeader, ok := request.ByteHeaders.Get(omniHTTP.XRequestID)
	True(p.T(), ok)

	Equal(p.T(), idHeader, []byte(testRequestID))
	Equal(p.T(), request.Body, testBody)
}

func (p *ProxySuite) TestOverrideConfirmability() {
	prxy := proxy.NewProxy(config.Config{}, p.metrics)
	forwarder := prxy.AcquireForwarder()
	_, span := p.metrics.Tracer().Start(p.GetTestContext(), fmt.Sprintf("test-%d", p.GetTestID()))
	forwarder.SetSpan(span)

	const chainConfirmations = uint16(10)
	const overridedConfirmations = uint16(2)
	// make the chain require 10 confirmations
	var urls []string
	for i := 0; i < int(chainConfirmations); i++ {
		urls = append(urls, gofakeit.URL())
	}

	chainManager := new(chainManagerMocks.Chain)
	chainManager.On("ConfirmationsThreshold").Return(chainConfirmations)
	chainManager.On("URLs").Return(urls)
	forwarder.SetChain(chainManager)

	forwarder.SetBody(p.MustMarshall(rpc.Request{
		ID:     1,
		Method: string(client.BlockByNumberMethod),
		Params: []json.RawMessage{[]byte("\"1\"")},
	}))
	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	forwarder.SetC(testContext)

	// try an override
	forwarder.SetRequiredConfirmations(overridedConfirmations)

	True(p.T(), forwarder.CheckAndSetConfirmability())
	Equal(p.T(), forwarder.RequiredConfirmations(), overridedConfirmations)
}
