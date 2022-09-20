package http_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
)

var jsonOptions = &gofakeit.JSONOptions{
	Type: "array",
	Fields: []gofakeit.Field{
		{Name: "id", Function: "autoincrement"},
		{Name: "first_name", Function: "firstname"},
	},
	RowCount: gofakeit.Number(5, 20),
	Indent:   true,
}

func (c *HTTPSuite) TestClient() {
	for _, client := range c.clients {
		headers := c.MockHeaders(10)

		mockBody, err := gofakeit.JSON(jsonOptions)
		Nil(c.T(), err)

		mockResp, err := gofakeit.JSON(jsonOptions)
		Nil(c.T(), err)

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for key, val := range headers {
				Equal(c.T(), val, r.Header.Get(key))
			}
			fullBody, err := io.ReadAll(r.Body)
			Nil(c.T(), err)

			Equal(c.T(), fullBody, mockBody)

			_, err = w.Write(mockResp)
			Nil(c.T(), err)
		}))

		req := client.NewRequest()
		req.SetRequestURI(svr.URL)
		req.SetBody(mockBody)
		req.SetContext(c.GetTestContext())
		for key, val := range headers {
			if gofakeit.Bool() {
				req.SetHeader(key, val)
			} else {
				req.SetHeaderBytes([]byte(key), []byte(val))
			}
		}

		resp, err := req.Do()
		Nil(c.T(), err)

		Equal(c.T(), resp.Body(), mockResp)
	}
}
