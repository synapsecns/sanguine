package rest_test

import (
	"fmt"
)

func (c *ServerSuite) TestNewAPIServer() {
	fmt.Println("I'M HERe")
	err := c.APIServer.Run(c.GetTestContext())
	c.Nil(err)

	// resp, err := http.Get("http://localhost:8080/quotes")
	// c.Nil(err)
	// defer resp.Body.Close()
	// c.Equal(http.StatusOK, resp.StatusCode)
}
