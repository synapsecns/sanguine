package relapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core/retry"
	submitterdb "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (c *RelayerServerSuite) TestNewQuoterAPIServer() {
	// Start the API server in a separate goroutine and wait for it to initialize.
	c.startQuoterAPIServer()
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/health", c.port), nil)
	c.Require().NoError(err)
	resp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		if resp != nil && resp.Body != nil {
			err = resp.Body.Close()
			c.Require().NoError(err)
		}
	}()
	c.Equal(http.StatusOK, resp.StatusCode)
	c.GetTestContext().Done()
}

func (c *RelayerServerSuite) TestGetQuoteRequestByTxHash() {
	c.startQuoterAPIServer()

	// Insert quote request to db
	quoteRequest := c.getTestQuoteRequest(reldb.Seen)
	err := c.database.StoreQuoteRequest(c.GetTestContext(), quoteRequest)
	c.Require().NoError(err)

	// Fetch the quote request by tx hash
	client := &http.Client{}
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/request/by_tx_hash?hash=%s", c.port, quoteRequest.OriginTxHash), nil)
	c.Require().NoError(err)
	resp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		if resp != nil && resp.Body != nil {
			err = resp.Body.Close()
			c.Require().NoError(err)
		}
	}()
	c.Equal(http.StatusOK, resp.StatusCode)

	// Compare to expected result
	var result relapi.GetQuoteRequestResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	c.Require().NoError(err)
	expectedResult := relapi.GetQuoteRequestResponse{
		Status:          quoteRequest.Status.String(),
		TxID:            hexutil.Encode(quoteRequest.TransactionID[:]),
		QuoteRequestRaw: result.QuoteRequestRaw,
		OriginTxHash:    quoteRequest.OriginTxHash.String(),
		DestTxHash:      quoteRequest.DestTxHash.String(),
		OriginChainID:   quoteRequest.Transaction.OriginChainId,
		DestChainID:     quoteRequest.Transaction.DestChainId,
		OriginToken:     quoteRequest.Transaction.OriginToken.String(),
		DestToken:       quoteRequest.Transaction.DestToken.String(),
		Sender:          quoteRequest.Sender.String(),
	}
	c.Equal(expectedResult, result)
	c.GetTestContext().Done()
}

func (c *RelayerServerSuite) TestGetQuoteRequestByTxID() {
	c.startQuoterAPIServer()

	// Insert quote request to db
	quoteRequest := c.getTestQuoteRequest(reldb.Seen)
	err := c.database.StoreQuoteRequest(c.GetTestContext(), quoteRequest)
	c.Require().NoError(err)

	// Fetch the quote request by tx hash
	client := &http.Client{}
	txIDStr := hexutil.Encode(quoteRequest.TransactionID[:])
	req, err := http.NewRequestWithContext(c.GetTestContext(), http.MethodGet, fmt.Sprintf("http://localhost:%d/request/by_tx_id?id=%s", c.port, txIDStr), nil)
	c.Require().NoError(err)
	resp, err := client.Do(req)
	c.Require().NoError(err)
	defer func() {
		if resp != nil && resp.Body != nil {
			err = resp.Body.Close()
			c.Require().NoError(err)
		}
	}()
	c.Equal(http.StatusOK, resp.StatusCode)

	// Compare to expected result
	var result relapi.GetQuoteRequestResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	c.Require().NoError(err)
	expectedResult := relapi.GetQuoteRequestResponse{
		Status:          quoteRequest.Status.String(),
		TxID:            hexutil.Encode(quoteRequest.TransactionID[:]),
		QuoteRequestRaw: result.QuoteRequestRaw,
		OriginTxHash:    quoteRequest.OriginTxHash.String(),
		DestTxHash:      quoteRequest.DestTxHash.String(),
		OriginChainID:   quoteRequest.Transaction.OriginChainId,
		DestChainID:     quoteRequest.Transaction.DestChainId,
		OriginToken:     quoteRequest.Transaction.OriginToken.String(),
		DestToken:       quoteRequest.Transaction.DestToken.String(),
		Sender:          quoteRequest.Sender.String(),
	}
	c.Equal(expectedResult, result)
	c.GetTestContext().Done()
}

func (c *RelayerServerSuite) TestGetTxRetry() {
	c.startQuoterAPIServer()

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
		if resp != nil && resp.Body != nil {
			err = resp.Body.Close()
			c.Require().NoError(err)
		}
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

// startQuoterAPIServer starts the API server and waits for it to initialize.
func (c *RelayerServerSuite) startQuoterAPIServer() {
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
			if resp != nil && resp.Body != nil {
				closeErr := resp.Body.Close()
				c.NoError(closeErr)
			}
		}()
		if err != nil {
			return fmt.Errorf("server not ready: %w", err)
		}
		return nil
	}, retry.WithMaxTotalTime(60*time.Second))
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
		Transaction: fastbridgev2.IFastBridgeV2BridgeTransactionV2{
			OriginChainId:      c.originChainID,
			DestChainId:        c.destChainID,
			OriginAmount:       big.NewInt(100),
			DestAmount:         big.NewInt(100),
			Deadline:           big.NewInt(time.Now().Unix()),
			Nonce:              big.NewInt(0),
			ExclusivityEndTime: big.NewInt(0),
			OriginFeeAmount:    big.NewInt(0),
			ZapNative:          big.NewInt(0),
			ZapData:            []byte{},
		},
		OriginTxHash: common.HexToHash("0x0000000"),
		DestTxHash:   common.HexToHash("0x0000001"),
	}
}

func TestTokenIDExists(t *testing.T) {
	cfg := relconfig.Config{
		QuotableTokens: map[string][]string{
			fmt.Sprintf("1%s0x1234567890abcdef1234567890abcdef12345678", relconfig.TokenIDDelimiter): {},
			fmt.Sprintf("1%s0xabcdefabcdefabcdefabcdefabcdefabcdefabcd", relconfig.TokenIDDelimiter): {},
		},
	}

	tests := []struct {
		name         string
		tokenAddress common.Address
		chainID      int
		expected     bool
	}{
		{
			name:         "Valid token address",
			tokenAddress: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			chainID:      1,
			expected:     true,
		},
		{
			name:         "Invalid token address",
			tokenAddress: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			chainID:      1,
			expected:     false,
		},
		{
			name:         "Valid token address, different chain ID",
			tokenAddress: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			chainID:      2,
			expected:     false,
		},
		{
			name:         "Edge case: empty token address",
			tokenAddress: common.Address{},
			chainID:      1,
			expected:     false,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			result := relapi.TokenIDExists(cfg, tt.tokenAddress, tt.chainID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToAddressIsWhitelisted(t *testing.T) {
	cfg := relconfig.Config{
		WithdrawalWhitelist: []string{
			"0x1111111111111111111111111111111111111111",
			"0x2222222222222222222222222222222222222222",
		},
	}

	tests := []struct {
		name      string
		toAddress common.Address
		expected  bool
	}{
		{
			name:      "Address is whitelisted",
			toAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
			expected:  true,
		},
		{
			name:      "Address is not whitelisted",
			toAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
			expected:  false,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			result := relapi.ToAddressIsWhitelisted(cfg, tt.toAddress)
			assert.Equal(t, tt.expected, result)
		})
	}
}
