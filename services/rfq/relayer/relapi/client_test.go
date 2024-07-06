package relapi_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (c *RelayerClientSuite) TestHealth() {
	ok, err := c.Client.Health(c.GetTestContext())
	c.NoError(err)
	c.True(ok)
}

func (c *RelayerClientSuite) TestGetQuoteRequestStatusByTxHash() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.GetQuoteRequestStatusByTxHash(c.GetTestContext(), testReq.OriginTxHash.String())
	c.Require().NoError(err)

	c.Equal(resp.Status, testReq.Status.String())
	c.Equal(resp.TxID, hexutil.Encode(testReq.TransactionID[:]))
	c.Equal(resp.DestTxHash, testReq.DestTxHash.String())
	c.Equal(resp.OriginTxHash, testReq.OriginTxHash.String())
}

func (c *RelayerClientSuite) TestGetQuoteRequestStatusByTxID() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.GetQuoteRequestStatusByTxID(c.GetTestContext(), hexutil.Encode(testReq.TransactionID[:]))
	c.Require().NoError(err)

	c.Equal(resp.Status, testReq.Status.String())
	c.Equal(resp.TxID, hexutil.Encode(testReq.TransactionID[:]))
	c.Equal(resp.DestTxHash, testReq.DestTxHash.String())
	c.Equal(resp.OriginTxHash, testReq.OriginTxHash.String())
}

func (c *RelayerClientSuite) TestRetryTransaction() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.RetryTransaction(c.GetTestContext(), testReq.OriginTxHash.String())
	c.Require().NoError(err)

	c.Equal(resp.TxID, hexutil.Encode(testReq.TransactionID[:]))
}

func (c *RelayerClientSuite) TestGetQuoteByTX() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.GetQuoteRequestByTXID(c.GetTestContext(), hexutil.Encode(testReq.TransactionID[:]))
	c.Require().NoError(err)

	c.Equal(len(common.Hex2Bytes(resp.QuoteRequestRaw)), len(testReq.RawRequest))
}
