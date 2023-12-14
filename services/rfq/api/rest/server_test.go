package rest_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *ServerSuite) TestNewAPIServer() {
	resp, err := http.Get("http://localhost:9000/quotes")
	c.Nil(err)
	defer resp.Body.Close()
	c.Equal(http.StatusOK, resp.StatusCode)
	c.GetTestContext().Done()
}

// TestEIP191_SuccessfulSignature tests the EIP191 signature process for successful authentication.
func (c *ServerSuite) TestEIP191_SuccessfulSignature() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startAPIServer()

	// Prepare the authorization header with a signed timestamp.
	header, err := c.prepareAuthHeader()
	if err != nil {
		c.Error(err)
		return
	}

	// Perform a PUT request to the API server with the authorization header.
	resp, err := c.sendPutRequest(header)
	if err != nil {
		c.Error(err)
		return
	}
	defer resp.Body.Close()

	// Log the response body for debugging.
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Assert that the response status code is HTTP 200 OK.
	c.Equal(http.StatusOK, resp.StatusCode)
}

// startAPIServer starts the API server and waits for it to initialize.
func (c *ServerSuite) startAPIServer() {
	go func() {
		err := c.APIServer.Run(c.GetTestContext())
		c.Nil(err)
	}()
	time.Sleep(2 * time.Second) // Wait for the server to start.
}

// prepareAuthHeader generates an authorization header using EIP191 signature.
func (c *ServerSuite) prepareAuthHeader() (string, error) {
	// Get the current Unix timestamp as a string.
	now := strconv.Itoa(int(time.Now().Unix()))

	// Prepare the data to be signed.
	data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now
	digest := crypto.Keccak256([]byte(data))

	// Sign the data.
	sig, err := crypto.Sign(digest, c.testWallet.PrivateKey())
	if err != nil {
		return "", err
	}
	signature := hexutil.Encode(sig)

	// Return the combined header value.
	return now + ":" + signature, nil
}

// sendPutRequest sends a PUT request to the server with the given authorization header.
func (c *ServerSuite) sendPutRequest(header string) (*http.Response, error) {
	// Prepare the PUT request with JSON data.
	client := &http.Client{}
	jsonData := `{"dest_chain_id":"1"}`
	req, err := http.NewRequest("PUT", "http://localhost:9000/quotes", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", header)

	// Send the request to the server.
	return client.Do(req)
}
