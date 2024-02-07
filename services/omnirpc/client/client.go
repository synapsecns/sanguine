package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"io"
	"math/big"
	"net/http"
)

// RPCClient is an interface for the omnirpc service.
type RPCClient interface {
	submitter.ClientFetcher
	// GetEndpoint returns the endpoint for the given chainID and confirmations.
	GetEndpoint(chainID, confirmations int) string
	// GetDefaultEndpoint returns the endpoint with the default confirmation count for the chain id.
	GetDefaultEndpoint(chainID int) string
	// GetConfirmationsClient returns a client for the given chainID and confirmations.
	// TODO: consider making confirmations iota or uint8, can be easy to confuse these params
	GetConfirmationsClient(ctx context.Context, chainID, confirmations int) (client.EVM, error)
	// GetChainClient returns a client for the given chainID.
	GetChainClient(ctx context.Context, chainID int) (client.EVM, error)
	// GetChainIDs returns all chain ids.
	GetChainIDs(ctx context.Context) ([]int, error)
}

type rpcClient struct {
	config   *rpcOptions
	endpoint string
	handler  metrics.Handler
	opts     []client.Options
}

// NewOmnirpcClient creates a new RPCClient.
func NewOmnirpcClient(endpoint string, handler metrics.Handler, options ...OptionsArgsOption) RPCClient {
	c := rpcClient{}
	c.config = makeOptions(options)
	c.endpoint = endpoint
	c.handler = handler
	c.opts = append(c.opts, client.Capture(c.config.captureReqRes))

	return &c
}

func (c *rpcClient) GetClient(ctx context.Context, chainID *big.Int) (client.EVM, error) {
	if !chainID.IsInt64() {
		return nil, errors.New("chain id is not a uint64")
	}

	return c.GetChainClient(ctx, int(chainID.Uint64()))
}

func (c *rpcClient) GetEndpoint(chainID, confirmations int) string {
	if confirmations == 0 {
		return fmt.Sprintf("%s/rpc/%d", c.endpoint, chainID)
	}
	return fmt.Sprintf("%s/confirmations/%d/rpc/%d", c.endpoint, confirmations, chainID)
}

func (c *rpcClient) GetDefaultEndpoint(chainID int) string {
	return c.GetEndpoint(chainID, c.config.confirmations)
}

func (c *rpcClient) GetConfirmationsClient(ctx context.Context, chainID, confirmations int) (client.EVM, error) {
	endpoint := c.GetEndpoint(chainID, confirmations)
	chainClient, err := client.DialBackend(ctx, endpoint, c.handler, c.opts...)
	if err != nil {
		return nil, fmt.Errorf("could not dial backend: %w", err)
	}
	return chainClient, nil
}

func (c *rpcClient) GetChainClient(ctx context.Context, chainID int) (client.EVM, error) {
	endpoint := c.GetDefaultEndpoint(chainID)
	chainClient, err := client.DialBackend(ctx, endpoint, c.handler, c.opts...)
	if err != nil {
		return nil, fmt.Errorf("could not dial backend: %w", err)
	}
	return chainClient, nil
}

func (c *rpcClient) GetChainIDs(ctx context.Context) (chainIDs []int, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/chain-ids", c.endpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	httpClient := new(http.Client)
	c.handler.ConfigureHTTPClient(httpClient)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get chain ids: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	readResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not get chain ids: %w", err)
	}

	err = json.Unmarshal(readResp, &chainIDs)
	if err != nil {
		return nil, fmt.Errorf("could not get chain ids: %w", err)
	}

	return chainIDs, nil
}
