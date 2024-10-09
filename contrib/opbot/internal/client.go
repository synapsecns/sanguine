package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dubonzi/otelresty"
	"github.com/go-http-utils/headers"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/valyala/fastjson"
)

const (
	getRequestByTxHash = "/api/transaction-id/%s"
)

// GetRFQByTxIDResponse represents the response of a quote request by transaction ID.
type RFQClient interface {
	// GetRFQByTxID gets a quote request by transaction ID.
	GetRFQByTxID(ctx context.Context, txID string) (resp *GetRFQByTxIDResponse, status string, err error)
}

type rfqClientImpl struct {
	client *resty.Client
}

// NewRFQClient creates a new RFQClient.
func NewRFQClient(handler metrics.Handler, url string) RFQClient {
	client := resty.New()
	client.SetBaseURL(url)
	client.SetHeader(headers.UserAgent, "rfq-client")

	otelresty.TraceClient(client, otelresty.WithTracerProvider(handler.GetTracerProvider()))

	return &rfqClientImpl{
		client: client,
	}
}

// GetRFQByTxID gets a quote request by transaction ID.
func (r *rfqClientImpl) GetRFQByTxID(ctx context.Context, txID string) (*GetRFQByTxIDResponse, string, error) {
	var res GetRFQByTxIDResponse
	resp, err := r.client.R().SetContext(ctx).
		SetQueryParam("hash", txID).
		SetResult(&res).
		Get(fmt.Sprintf(getRequestByTxHash, txID))
	if err != nil {
		return nil, "", fmt.Errorf("failed to get quote request by tx hash: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, "", fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var status string
	if fastjson.GetString(resp.Body(), "BridgeClaim") != "" {
		status = "Claimed"
	} else if fastjson.GetString(resp.Body(), "BridgeProof") != "" {
		status = "Proven"
	} else if fastjson.GetString(resp.Body(), "BridgeRelay") != "" {
		status = "Relayed"
	} else if fastjson.GetString(resp.Body(), "BridgeRequest") != "" {
		status = "Requested"
	} else {
		status = "Unknown"
	}

	return &res, status, nil

}

var _ RFQClient = &rfqClientImpl{}
