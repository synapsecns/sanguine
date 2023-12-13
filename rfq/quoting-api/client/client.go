package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/db/models"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/rest"
	"strconv"
	"time"
)

type Client interface {
	CreateQuote(q models.Quote) error
}

type clientImpl struct {
	rClient *resty.Client
}

// NewClient creates a new client for the RFQ quoting API
func NewClient(rfqURL string, reqSigner signer.Signer) (Client, error) {
	client := resty.New().
		SetBaseURL(rfqURL).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			timestamp := strconv.Itoa(int(time.Now().Unix()))
			hash, _ := signer.SignHash([]byte(timestamp))
			sig, err := reqSigner.SignMessage(request.Context(), hash, true)
			if err != nil {
				return fmt.Errorf("failed to sign request: %w", err)
			}

			request.SetHeader("Authentication", signer.EncodeHex(sig))

			return nil
		})

	return clientImpl{
		rClient: client,
	}, nil
}

// CreateQuote creates a new quote in the RFQ quoting API
func (c clientImpl) CreateQuote(q models.Quote) error {
	res, err := c.rClient.R().
		SetBody(q).
		Post(rest.QUOTE_ROUTE)
	// TODO: Figure out if there's anyhting to do with the response
	_ = res

	return err
}
