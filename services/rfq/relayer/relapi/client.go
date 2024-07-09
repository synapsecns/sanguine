package relapi

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

// RelayerClient is the interface for the relayer client.
type RelayerClient interface {
	Health(ctx context.Context) (ok bool, err error)
	GetQuoteRequestStatusByTxHash(ctx context.Context, hash string) (*GetQuoteRequestStatusResponse, error)
	GetQuoteRequestStatusByTxID(ctx context.Context, hash string) (*GetQuoteRequestStatusResponse, error)
	RetryTransaction(ctx context.Context, txhash string) (*GetTxRetryResponse, error)
	Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error)
	GetTxHashByNonce(ctx context.Context, req *GetTxByNonceRequest) (*TxHashByNonceResponse, error)
	GetQuoteRequestByTXID(ctx context.Context, txid string) (*GetQuoteRequestResponse, error)
}

type relayerClient struct {
	client *resty.Client
}

// NewRelayerClient creates a new RelayerClient.
func NewRelayerClient(handler metrics.Handler, url string) RelayerClient {
	client := resty.New()
	client.SetBaseURL(url)
	client.SetHeader(headers.UserAgent, "relayer-client")

	otelresty.TraceClient(client, otelresty.WithTracerProvider(handler.GetTracerProvider()))

	return &relayerClient{
		client: client,
	}
}

// Health checks if the relayer is healthy.
func (r *relayerClient) Health(ctx context.Context) (ok bool, err error) {
	resp, err := r.client.R().SetContext(ctx).Get(getHealthRoute)
	if err != nil {
		return false, fmt.Errorf("failed to check health: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	ok = fastjson.GetString(resp.Body(), "status") == "ok"

	return ok, nil
}

func (r *relayerClient) GetQuoteRequestStatusByTxHash(ctx context.Context, hash string) (*GetQuoteRequestStatusResponse, error) {
	var res GetQuoteRequestStatusResponse

	resp, err := r.client.R().SetContext(ctx).
		SetQueryParam("hash", hash).
		SetResult(&res).
		Get(getQuoteStatusByTxHashRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to get quote request status by tx hash: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &res, nil
}

func (r *relayerClient) GetQuoteRequestStatusByTxID(ctx context.Context, txid string) (*GetQuoteRequestStatusResponse, error) {
	var res GetQuoteRequestStatusResponse

	resp, err := r.client.R().SetContext(ctx).
		SetQueryParam("id", txid).
		SetResult(&res).
		Get(getQuoteStatusByTxIDRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to get quote request status by tx hash: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &res, nil
}

func (r *relayerClient) RetryTransaction(ctx context.Context, txhash string) (*GetTxRetryResponse, error) {
	var res GetTxRetryResponse
	resp, err := r.client.R().SetContext(ctx).
		SetQueryParam("hash", txhash).
		SetResult(&res).
		Get(getRetryRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to retry transaction: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &res, nil
}

// WithdrawResponse is the response for the withdraw request.
type WithdrawResponse struct {
	Nonce uint64 `json:"nonce"`
}

// Withdraw withdraws an ERC20 or Ether from the relayer.
func (r *relayerClient) Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error) {
	var res WithdrawResponse
	resp, err := r.client.R().SetContext(ctx).
		SetResult(&res).
		SetBody(req).
		Post(postWithdrawRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw transaction: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &res, nil
}

// TxHashByNonceResponse is the request for getting a transaction hash by nonce.
type TxHashByNonceResponse struct {
	Hash string `json:"withdrawTxHash"`
}

// GetTxByNonceRequest is the request for getting a transaction hash by nonce.
func (r *relayerClient) GetTxHashByNonce(ctx context.Context, req *GetTxByNonceRequest) (*TxHashByNonceResponse, error) {
	var res TxHashByNonceResponse

	resp, err := r.client.R().SetContext(ctx).
		SetResult(&res).
		SetQueryParam("chain_id", fmt.Sprintf("%d", req.ChainID)).
		SetQueryParam("nonce", fmt.Sprintf("%d", req.Nonce)).
		Get(getTxHashByNonceRoute)

	if err != nil {
		return nil, fmt.Errorf("failed to get tx hash by nonce: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &res, nil
}

func (r *relayerClient) GetQuoteRequestByTXID(ctx context.Context, txid string) (*GetQuoteRequestResponse, error) {
	var res GetQuoteRequestResponse
	resp, err := r.client.R().SetContext(ctx).
		SetQueryParam("id", txid).
		SetResult(&res).
		Get(getRequestByTxID)
	if err != nil {
		return nil, fmt.Errorf("failed to get quote request by tx id: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &res, nil
}
