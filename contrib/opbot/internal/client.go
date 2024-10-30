// Package internal provides the RFQ client implementation.
package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dubonzi/otelresty"
	"github.com/go-http-utils/headers"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/core/metrics"
)

const (
	getRFQRoute = "/transaction-id/%s"
)

// RFQClient is the interface for the RFQ client.
type RFQClient interface {
	// GetRFQ gets a quote request by transaction ID.
	GetRFQ(ctx context.Context, txIdentifier string) (resp *GetRFQByTxIDResponse, status string, err error)
}

type rfqClientImpl struct {
	client *resty.Client
}

// NewRFQClient creates a new RFQClient.
func NewRFQClient(handler metrics.Handler, indexerURL string) RFQClient {
	client := resty.New()
	client.SetBaseURL(indexerURL)
	client.SetHeader(headers.UserAgent, "rfq-client")

	otelresty.TraceClient(client, otelresty.WithTracerProvider(handler.GetTracerProvider()))

	return &rfqClientImpl{
		client: client,
	}
}

// GetRFQByTxID gets a quote request by transaction ID or transaction hash.
func (r *rfqClientImpl) GetRFQ(ctx context.Context, txIdentifier string) (*GetRFQByTxIDResponse, string, error) {
	var res GetRFQByTxIDResponse
	resp, err := r.client.R().SetContext(ctx).
		SetResult(&res).
		Get(fmt.Sprintf(getRFQRoute, txIdentifier))
	if err != nil {
		return nil, "", fmt.Errorf("failed to get quote request by tx ID: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, "", fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var status string
	switch {
	case res.BridgeClaim != (BridgeClaim{}):
		status = "Claimed"
	case res.BridgeProof != (BridgeProof{}):
		status = "Proven"
	case res.BridgeRelay != (BridgeRelay{}):
		status = "Relayed"
	case res.BridgeRequest != (BridgeRequest{}):
		status = "Requested"
	default:
		status = "Unknown"
	}

	return &res, status, nil
}
