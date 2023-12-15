package client

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
)

// Client is an interface for the RFQ API.
// It provides methods for creating, retrieving and updating quotes.
type Client interface {
	PutQuote(q *APIQuotePutRequest) error
	GetAllQuotes() ([]*db.Quote, error)
	GetSpecificQuote(q *APIQuoteSpecificGetRequest) ([]*db.Quote, error)
	GetQuoteByRelayerAddress(relayerAddr string) ([]*db.Quote, error)
}

type clientImpl struct {
	rClient *resty.Client
}

// NewClient creates a new client for the RFQ quoting API.
// TODO: @aurelius,  you don't actually need to be authed for GET Requests
func NewClient(rfqURL string, reqSigner signer.Signer) (Client, error) {
	client := resty.New().
		SetBaseURL(rfqURL).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			// if request.Method == "PUT" && request.URL == rfqURL+rest.QUOTE_ROUTE {
			// i.e. signature (hex encoded) = keccak(bytes.concat("\x19Ethereum Signed Message:\n", len(strconv.Itoa(time.Now().Unix()), strconv.Itoa(time.Now().Unix())))
			// so that full auth header string: auth = strconv.Itoa(time.Now().Unix()) + ":" + signature
			// Get the current Unix timestamp as a string.
			now := strconv.Itoa(int(time.Now().Unix()))

			// Prepare the data to be signed.
			data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now

			sig, err := reqSigner.SignMessage(request.Context(), []byte(data), true)

			if err != nil {
				return fmt.Errorf("failed to sign request: %w", err)
			}

			res := fmt.Sprintf("%s:%s", now, hexutil.Encode(signer.Encode(sig)))
			request.SetHeader("Authorization", res)

			return nil
			// } else {
			// 	return nil
			// }
		})

	return &clientImpl{
		rClient: client,
	}, nil
}

// CreateQuote creates a new quote in the RFQ quoting API.
func (c *clientImpl) PutQuote(q *APIQuotePutRequest) error {
	res, err := c.rClient.R().
		SetBody(q).
		Put(rest.QuoteRoute)

	// TODO: Figure out if there's anything to do with the response, right now it's result: Status Code 200 OK
	_ = res

	return err
}

// CreateQuote creates a new quote in the RFQ quoting API.
func (c *clientImpl) GetAllQuotes() ([]*db.Quote, error) {
	var quotes []*db.Quote
	resp, err := c.rClient.R().
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", resp.Status())
	}

	return quotes, nil
}

// CreateQuote creates a new quote in the RFQ quoting API.
func (c *clientImpl) GetSpecificQuote(q *APIQuoteSpecificGetRequest) ([]*db.Quote, error) {
	var quotes []*db.Quote
	resp, err := c.rClient.R().
		SetQueryParams(map[string]string{
			"originChainId":   q.OriginChainID,
			"originTokenAddr": q.OriginTokenAddr,
			"destChainId":     q.DestChainID,
			"destTokenAddr":   q.DestTokenAddr,
		}).
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", resp.Status())
	}

	return quotes, nil
}

func (c *clientImpl) GetQuoteByRelayerAddress(relayerAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote
	resp, err := c.rClient.R().
		SetQueryParams(map[string]string{
			"relayerAddr": relayerAddr,
		}).
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", resp.Status())
	}

	return quotes, nil
}

// APIQuotePutRequest is the struct for the quote API.
type APIQuotePutRequest struct {
	OriginChainID   string `json:"origin_chain_id"`
	OriginTokenAddr string `json:"origin_token_addr"`
	DestChainID     string `json:"dest_chain_id"`
	DestTokenAddr   string `json:"dest_token_addr"`
	DestAmount      string `json:"dest_amount"`
	Price           string `json:"price"`
	MaxOriginAmount string `json:"max_origin_amount"`
}

// APIQuoteSpecificGetRequest is the struct for the quote API.
type APIQuoteSpecificGetRequest struct {
	OriginChainID   string `json:"originChainId"`
	OriginTokenAddr string `json:"originTokenAddr"`
	DestChainID     string `json:"destChainId"`
	DestTokenAddr   string `json:"destTokenAddr"`
}
