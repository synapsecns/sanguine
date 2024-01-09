package relapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core/retry"
	submitterdb "github.com/synapsecns/sanguine/ethergo/submitter/db"
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

func (c *RelayerServerSuite) TestGetQuoteRequestByTxHash() {
	c.startAPIServer()

	// Insert quote request to db
	quoteRequest := c.getTestQuoteRequest(reldb.Seen)
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
	quoteRequest := c.getTestQuoteRequest(reldb.Seen)
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
	quoteRequest := c.getTestQuoteRequest(reldb.Seen)
	err := c.database.StoreQuoteRequest(c.GetTestContext(), quoteRequest)
	c.Require().NoError(err)

	// Send a retry request
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/retry?hash=%s", c.port, quoteRequest.OriginTxHash), nil)
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

	// Verify that a transaction was submitted
	status, err := c.database.SubmitterDB().GetNonceStatus(c.GetTestContext(), c.wallet.Address(), big.NewInt(int64(quoteRequest.Transaction.DestChainId)), result.Nonce)
	c.Require().NoError(err)
	c.Equal(status, submitterdb.Stored)
}

// startAPIServer starts the API server and waits for it to initialize.
func (c *RelayerServerSuite) startAPIServer() {
	go func() {
		err := c.RelayerAPIServer.Run(c.GetTestContext())
		c.Require().NoError(err)
	}()

	// Wait for the server to start
	err := retry.WithBackoff(c.GetTestContext(), func(ctx context.Context) error {
		client := &http.Client{}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:%d/health", c.port), nil)
		c.Require().NoError(err)
		resp, err := client.Do(req)
		defer func() {
			err = resp.Body.Close()
			c.Require().NoError(err)
		}()
		return err
	}, retry.WithMaxTotalTime(10*time.Second))
	c.Require().NoError(err)
}

func (c *RelayerServerSuite) getTestQuoteRequest(status reldb.QuoteRequestStatus) reldb.QuoteRequest {
	txIDRaw := hexutil.Encode(crypto.Keccak256([]byte("test")))
	var txID [32]byte
	copy(txID[:], txIDRaw)
	return reldb.QuoteRequest{
		OriginTokenDecimals: 6,
		DestTokenDecimals:   6,
		TransactionID:       txID,
		Status:              status,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: c.originChainID,
			DestChainId:   c.destChainID,
			OriginAmount:  big.NewInt(100),
			DestAmount:    big.NewInt(100),
			Deadline:      big.NewInt(time.Now().Unix()),
			Nonce:         big.NewInt(0),
		},
		OriginTxHash: common.HexToHash("0x0000000"),
		DestTxHash:   common.HexToHash("0x0000001"),
	}
}
