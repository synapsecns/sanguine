package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
)

func (c *ServerSuite) TestNewAPIServer() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startAPIServer()
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/quotes", c.port), nil)
	c.Require().NoError(err)
	resp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		err = resp.Body.Close()
		c.Require().NoError(err)
	}()
	c.Equal(http.StatusOK, resp.StatusCode)
	c.GetTestContext().Done()
}

// TestEIP191_SuccessfulSignature tests the EIP191 signature process for successful authentication.
func (c *ServerSuite) TestEIP191_SuccessfulSignature() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startAPIServer()

	// Prepare the authorization header with a signed timestamp.
	header, err := c.prepareAuthHeader(c.testWallet)
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
	defer func() {
		err = resp.Body.Close()
		c.Require().NoError(err)
	}()

	// Log the response body for debugging.
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Assert that the response status code is HTTP 200 OK.
	c.Equal(http.StatusOK, resp.StatusCode)
}

// TestEIP191_UnsuccessfulSignature tests the EIP191 signature process with an incorrect wallet signature.
func (c *ServerSuite) TestEIP191_UnsuccessfulSignature() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startAPIServer()

	// Prepare the authorization header with a signed timestamp using an incorrect wallet.
	randomWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	header, err := c.prepareAuthHeader(randomWallet)
	if err != nil {
		c.Error(err)
		return
	}

	// Perform a PUT request to the API server with the incorrect authorization header.
	resp, err := c.sendPutRequest(header)
	if err != nil {
		c.Error(err)
		return
	}
	defer func() {
		err = resp.Body.Close()
		c.Require().NoError(err)
	}()
	// Log the response body for debugging.
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Assert that the response status code is HTTP 400 Bad Request.
	c.Equal(http.StatusBadRequest, resp.StatusCode)
}

// TestEIP191_SuccessfulPutSubmission tests a successful PUT request submission.
func (c *ServerSuite) TestEIP191_SuccessfulPutSubmission() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startAPIServer()

	// Prepare the authorization header with a signed timestamp.
	header, err := c.prepareAuthHeader(c.testWallet)
	c.Require().NoError(err)

	// Perform a PUT request to the API server with the authorization header.
	resp, err := c.sendPutRequest(header)
	c.Require().NoError(err)
	defer func() {
		_ = resp.Body.Close()
	}()

	// Log the response body for debugging.
	body, err := io.ReadAll(resp.Body)
	c.Require().NoError(err)
	fmt.Println(string(body))

	// Assert that the response status code is HTTP 200 OK.
	c.Assert().Equal(http.StatusOK, resp.StatusCode)
}

func (c *ServerSuite) TestPutAndGetQuote() {
	c.startAPIServer()

	header, err := c.prepareAuthHeader(c.testWallet)
	c.Require().NoError(err)

	// Send PUT request
	putResp, err := c.sendPutRequest(header)
	c.Require().NoError(err)
	defer func() {
		err = putResp.Body.Close()
		c.Require().NoError(err)
	}()
	c.Assert().Equal(http.StatusOK, putResp.StatusCode)

	// Send GET request to verify the PUT
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/quotes?originChainId=1&originTokenAddr=0xOriginTokenAddrdestChainId=42161&destTokenAddr=0xDestTokenAddr", c.port), nil)
	c.Require().NoError(err)

	getResp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		_ = getResp.Body.Close()
	}()
	c.Assert().Equal(http.StatusOK, getResp.StatusCode)

	var quotes []*db.Quote
	err = json.NewDecoder(getResp.Body).Decode(&quotes)
	c.Require().NoError(err)

	// Check if the newly added quote is present
	found := false
	for _, q := range quotes {
		if q.FixedFee.String() == "10" {
			found = true
			break
		}
	}
	c.Assert().True(found, "Newly added quote not found")
}

func (c *ServerSuite) TestPutAndGetQuoteByRelayer() {
	c.startAPIServer()

	header, err := c.prepareAuthHeader(c.testWallet)
	c.Require().NoError(err)

	// Send PUT request
	putResp, err := c.sendPutRequest(header)
	c.Require().NoError(err)
	defer func() {
		err = putResp.Body.Close()
		c.Require().NoError(err)
	}()
	c.Assert().Equal(http.StatusOK, putResp.StatusCode)

	// Send GET request to verify the PUT
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/quotes?relayerAddress=%s", c.port, c.testWallet.Address().Hex()), nil)
	c.Require().NoError(err)

	getResp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		_ = getResp.Body.Close()
	}()
	c.Assert().Equal(http.StatusOK, getResp.StatusCode)

	var quotes []*db.Quote
	err = json.NewDecoder(getResp.Body).Decode(&quotes)
	c.Require().NoError(err)

	// Check if the newly added quote is present
	found := false
	for _, q := range quotes {
		if q.FixedFee.String() == "10" {
			found = true
			break
		}
	}
	c.Assert().True(found, "Newly added quote not found")
}

// startAPIServer starts the API server and waits for it to initialize.
func (c *ServerSuite) startAPIServer() {
	go func() {
		err := c.APIServer.Run(c.GetTestContext())
		c.Require().NoError(err)
	}()
	time.Sleep(2 * time.Second) // Wait for the server to start.
}

// prepareAuthHeader generates an authorization header using EIP191 signature with the given private key.
func (c *ServerSuite) prepareAuthHeader(wallet wallet.Wallet) (string, error) {
	// Get the current Unix timestamp as a string.
	now := strconv.Itoa(int(time.Now().Unix()))

	// Prepare the data to be signed.
	data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now
	digest := crypto.Keccak256([]byte(data))

	// Sign the data with the provided private key.
	sig, err := crypto.Sign(digest, wallet.PrivateKey())
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}
	signature := hexutil.Encode(sig)

	// Return the combined header value.
	return now + ":" + signature, nil
}

// sendPutRequest sends a PUT request to the server with the given authorization header.
func (c *ServerSuite) sendPutRequest(header string) (*http.Response, error) {
	// Prepare the PUT request with JSON data.
	client := &http.Client{}
	putData := rest.PutRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100.0",
		MaxOriginAmount: "200.0",
		FixedFee:        "10.0",
	}
	jsonData, err := json.Marshal(putData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal putData: %w", err)
	}

	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodPut, fmt.Sprintf("http://localhost:%d/quotes", c.port), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create PUT request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", header)

	// Send the request to the server.
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send PUT request: %w", err)
	}
	return resp, nil
}
