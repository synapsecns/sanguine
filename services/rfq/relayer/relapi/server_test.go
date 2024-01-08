package relapi_test

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

func (c *RelayerServerSuite) TestNewAPIServer() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startAPIServer()
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/status", c.port), nil)
	c.Require().NoError(err)
	_, err = client.Do(req)
	c.Require().NoError(err)
	c.GetTestContext().Done()
}

func (c *RelayerServerSuite) TestPutAndGetQuote() {
	c.startAPIServer()
}

// startAPIServer starts the API server and waits for it to initialize.
func (c *RelayerServerSuite) startAPIServer() {
	go func() {
		err := c.RelayerAPIServer.Run(c.GetTestContext())
		c.Require().NoError(err)
	}()
	time.Sleep(2 * time.Second) // Wait for the server to start.
}

// prepareAuthHeader generates an authorization header using EIP191 signature with the given private key.
func (c *RelayerServerSuite) prepareAuthHeader(wallet wallet.Wallet) (string, error) {
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
