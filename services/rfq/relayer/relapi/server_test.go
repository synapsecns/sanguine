package relapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (c *RelayerServerSuite) TestNewAPIServer() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startAPIServer()
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/health", c.port), nil)
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

func getTestQuoteRequest(status reldb.QuoteRequestStatus) reldb.QuoteRequest {
	txIDRaw := hexutil.Encode(crypto.Keccak256([]byte("test")))
	var txID [32]byte
	copy(txID[:], txIDRaw)
	return reldb.QuoteRequest{
		OriginTokenDecimals: 6,
		DestTokenDecimals:   6,
		TransactionID:       txID,
		Status:              status,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginAmount: big.NewInt(100),
			DestAmount:   big.NewInt(100),
			Deadline:     big.NewInt(time.Now().Unix()),
			Nonce:        big.NewInt(1),
		},
		OriginTxHash: common.HexToHash("0x0000000"),
		DestTxHash:   common.HexToHash("0x0000001"),
	}
}

func (c *RelayerServerSuite) TestGetQuoteRequestByTxHash() {
	c.startAPIServer()

	// Insert quote request to db
	quoteRequest := getTestQuoteRequest(reldb.Seen)
	err := c.database.StoreQuoteRequest(c.GetTestContext(), quoteRequest)
	c.Require().NoError(err)

	// Fetch the quote request by tx hash
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/status?hash=%s", c.port, quoteRequest.OriginTxHash), nil)
	c.Require().NoError(err)
	resp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		err = resp.Body.Close()
		c.Require().NoError(err)
	}()
	c.Equal(http.StatusOK, resp.StatusCode)

	// Compare to expected result
	var result relapi.GetQuoteRequestStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	c.Require().NoError(err)
	expectedResult := relapi.GetQuoteRequestStatusResponse{
		Status:       quoteRequest.Status.String(),
		TxID:         hexutil.Encode(quoteRequest.TransactionID[:]),
		OriginTxHash: quoteRequest.OriginTxHash.String(),
		DestTxHash:   quoteRequest.DestTxHash.String(),
	}
	c.Equal(expectedResult, result)
	c.GetTestContext().Done()
}

func (c *RelayerServerSuite) TestGetQuoteRequestByTxID() {
	c.startAPIServer()

	// Insert quote request to db
	quoteRequest := getTestQuoteRequest(reldb.Seen)
	err := c.database.StoreQuoteRequest(c.GetTestContext(), quoteRequest)
	c.Require().NoError(err)

	// Fetch the quote request by tx hash
	client := &http.Client{}
	txIDStr := hexutil.Encode(quoteRequest.TransactionID[:])
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/status/by_tx_id?id=%s", c.port, txIDStr), nil)
	c.Require().NoError(err)
	resp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		err = resp.Body.Close()
		c.Require().NoError(err)
	}()
	c.Equal(http.StatusOK, resp.StatusCode)

	// Compare to expected result
	var result relapi.GetQuoteRequestStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	c.Require().NoError(err)
	expectedResult := relapi.GetQuoteRequestStatusResponse{
		Status:       quoteRequest.Status.String(),
		TxID:         hexutil.Encode(quoteRequest.TransactionID[:]),
		OriginTxHash: quoteRequest.OriginTxHash.String(),
		DestTxHash:   quoteRequest.DestTxHash.String(),
	}
	c.Equal(expectedResult, result)
	c.GetTestContext().Done()
}

func (c *RelayerServerSuite) TestGetTxRetry() {
	c.startAPIServer()

	// Insert quote request to db
	quoteRequest := getTestQuoteRequest(reldb.Seen)
	err := c.database.StoreQuoteRequest(c.GetTestContext(), quoteRequest)
	c.Require().NoError(err)

	// Send a retry request
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodPut, fmt.Sprintf("http://localhost:%d/retry?hash=%s", c.port, quoteRequest.OriginTxHash), nil)
	c.Require().NoError(err)
	resp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		err = resp.Body.Close()
		c.Require().NoError(err)
	}()
	c.Equal(http.StatusOK, resp.StatusCode)

	// Compare to expected result
	var result relapi.GetTxRetryResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	c.Require().NoError(err)
	expectedResult := relapi.GetTxRetryResponse{
		TxID:      hexutil.Encode(quoteRequest.TransactionID[:]),
		ChainID:   quoteRequest.Transaction.DestChainId,
		Nonce:     uint64(quoteRequest.Transaction.Nonce.Int64()),
		GasAmount: "0",
	}
	c.Equal(expectedResult, result)
	c.GetTestContext().Done()
}

// startAPIServer starts the API server and waits for it to initialize.
func (c *RelayerServerSuite) startAPIServer() {
	go func() {
		err := c.RelayerAPIServer.Run(c.GetTestContext())
		c.Require().NoError(err)
	}()

	// Wait for the server to start
	retry.WithBackoff(c.GetTestContext(), func(ctx context.Context) error {
		client := &http.Client{}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:%d/health", c.port), nil)
		c.Require().NoError(err)
		_, err = client.Do(req)
		return err
	}, retry.WithMaxTotalTime(10*time.Second))
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
