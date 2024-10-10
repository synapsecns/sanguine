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
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
)

const (
	getRequestByTxHash = "/api/transaction-id/%s"
)

// RFQClient is the interface for the RFQ client.
type RFQClient interface {
	// GetRFQByTxID gets a quote request by transaction ID.
	GetRFQByTxID(ctx context.Context, txID string) (resp *GetRFQByTxIDResponse, status string, err error)
	// GetRFQByTxHash gets a quote request by transaction hash.
	GetRFQByTxHash(ctx context.Context, txHash string) (resp *relapi.GetQuoteRequestResponse, err error)
}

type rfqClientImpl struct {
	client         *resty.Client
	relayerClients []relapi.RelayerClient
}

// NewRFQClient creates a new RFQClient.
func NewRFQClient(handler metrics.Handler, indexerURL string, relayerURLs []string) RFQClient {
	client := resty.New()
	client.SetBaseURL(indexerURL)
	client.SetHeader(headers.UserAgent, "rfq-client")

	otelresty.TraceClient(client, otelresty.WithTracerProvider(handler.GetTracerProvider()))

	var relayerClients []relapi.RelayerClient
	for _, url := range relayerURLs {
		relayerClients = append(relayerClients, relapi.NewRelayerClient(handler, url))
	}

	return &rfqClientImpl{
		client:         client,
		relayerClients: relayerClients,
	}
}

// GetRFQByTxID gets a quote request by transaction ID.
func (r *rfqClientImpl) GetRFQByTxID(ctx context.Context, txID string) (*GetRFQByTxIDResponse, string, error) {
	var res GetRFQByTxIDResponse
	resp, err := r.client.R().SetContext(ctx).
		SetResult(&res).
		Get(fmt.Sprintf(getRequestByTxHash, txID))
	if err != nil {
		return nil, "", fmt.Errorf("failed to get quote request by tx hash: %w", err)
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

// GetRFQByTxHash gets a quote request by transaction hash.
func (r *rfqClientImpl) GetRFQByTxHash(ctx context.Context, txHash string) (*relapi.GetQuoteRequestResponse, error) {
	var resp *relapi.GetQuoteRequestResponse
	var err error
	for _, relayerClient := range r.relayerClients {
		resp, err = relayerClient.GetQuoteRequestByTxHash(ctx, txHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get quote request by tx hash: %w", err)
		}
	}

	return resp, nil
}

var _ RFQClient = &rfqClientImpl{}
