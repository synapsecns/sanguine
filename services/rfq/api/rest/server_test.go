package rest_test

import (
	"fmt"
	"net/http"
)

func (c *ServerSuite) TestNewAPIServer() {
	fmt.Println("I'M HERe")
	err := c.APIServer.Run(c.GetTestContext())
	c.Nil(err)
	fmt.Println("RAN THE SERVER MFER")
	resp, err := http.Get("http://localhost:9000/quotes")
	c.Nil(err)
	defer resp.Body.Close()
	c.Equal(http.StatusOK, resp.StatusCode)
	c.GetTestContext().Done()
}
