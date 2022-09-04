package integration_test

import (
	"fmt"
	"net/http"

	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/server"
)

func (i IntegrationSuite) TestGqlServer() {
	// fill w/ fake data
	// etc

	port := freeport.GetPort()

	go func() {
		Nil(i.T(), server.Start(uint16(port)))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)

	i.Eventually(func() bool {
		// TODO: use context here
		_, err := http.Get(fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint))
		return err == nil
	})

	// TODO: use conext
	gqlClient := client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	res, err := gqlClient.GetLogs(i.GetTestContext())
	Nil(i.T(), err)

	// TODO: this will panic if response is nil
	Equal(i.T(), res.Response[0].BlockNumber, 131)
}
