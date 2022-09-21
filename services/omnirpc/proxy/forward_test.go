package proxy_test

import (
	"bytes"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy/mocks"
	"net/http"
	"net/http/httptest"
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
	mockBody := new(mocks.BodyReader)
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
