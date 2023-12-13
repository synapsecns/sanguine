package rest_test

import (
	"net/http"
	"time"
)

func (c *ServerSuite) TestNewAPIServer() {
	go func() {
		err := c.APIServer.Run(c.GetTestContext())
		c.Nil(err)
	}()
	time.Sleep(time.Second * 2) // wait for server to start
	resp, err := http.Get("http://localhost:9000/quotes")
	c.Nil(err)
	defer resp.Body.Close()
	c.Equal(http.StatusOK, resp.StatusCode)
	c.GetTestContext().Done()
}
