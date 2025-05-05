package restclient_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/restclient"
)

func TestGetRoot(t *testing.T) {
	// Create a new client
	client, err := restclient.NewClient("https://api.synapseprotocol.com")
	assert.NoError(t, err)

	// Make a request to the root of the API
	resp, err := client.Get(context.Background())
	assert.NoError(t, err)
	defer func() {
		_ = resp.Body.Close()
	}()

	// Check that the status code is 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
