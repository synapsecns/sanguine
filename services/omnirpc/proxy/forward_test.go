package proxy_test

import (
	"bytes"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager/mocks"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	proxyMocks "github.com/synapsecns/sanguine/services/omnirpc/proxy/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

func (p *ProxySuite) TestServeRequestNoChain() {
	prxy := proxy.NewProxy(config.Config{}, omniHTTP.FastHTTP)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	prxy.Forward(c, 1)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

func (p *ProxySuite) TestCannotReadBody() {
	prxy := proxy.NewProxy(config.Config{}, omniHTTP.FastHTTP)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mockBody := new(proxyMocks.BodyReader)
	mockBody.On("Read", mock.Anything).Return(0, errors.New("could not read body"))

	c.Request = &http.Request{
		Body: mockBody,
	}

	prxy.Forward(c, 1)
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
	prxy := proxy.NewProxy(config.Config{}, omniHTTP.FastHTTP)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(p.generateFakeJSON()))

	prxy.Forward(c, 1)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

// TestAcquireReleaseForwarder makes sure the forwarder is cleared afte r being released.
func (p *ProxySuite) TestAcquireReleaseForwarder() {
	prxy := proxy.NewProxy(config.Config{}, omniHTTP.FastHTTP)

	forwarder := prxy.AcquireForwarder()
	forwarder.SetChain(new(mocks.Chain))
	forwarder.SetC(&gin.Context{})
	forwarder.SetClient(omniHTTP.NewClient(omniHTTP.Resty))
	forwarder.SetR(prxy)
	forwarder.SetBody(gofakeit.ImagePng(5, 5))
	forwarder.SetRequestID(uuid.New().String())
	forwarder.SetRequiredConfirmations(gofakeit.Uint16())
	forwarder.SetBlankResMap()
	forwarder.SetRPCRequest(&proxy.RPCRequest{
		ID:     []byte(strconv.Itoa(gofakeit.Number(1, 2))),
		Method: gofakeit.Word(),
	})
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
	prxy := proxy.NewProxy(config.Config{}, omniHTTP.FastHTTP)

	invalidSchemes := []string{"wss", "ws"}
	for _, scheme := range invalidSchemes {
		forwarder := prxy.AcquireForwarder()

		testURL := gofakeit.URL()
		parsedURL, err := url.Parse(testURL)
		Nil(p.T(), err)

		// change the scheme to use wss to see if
		parsedURL.Scheme = scheme
		testURL = parsedURL.String()

		rawRes, err := forwarder.ForwardRequest(p.GetTestContext(), testURL, gofakeit.UUID())
		Nil(p.T(), rawRes)
		NotNil(p.T(), err)

		prxy.ReleaseForwarder(forwarder)
	}
}

func (p *ProxySuite) TestForwardRequest() {

}
