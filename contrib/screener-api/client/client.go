// Package client provides a client for the screener-api.
package client

import (
	"context"
	"fmt"
	"github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
)

// ScreenerClient is an interface for the Screener API.
type ScreenerClient interface {
	ScreenAddress(ctx context.Context, ruleset, address string) (bool, error)
}

type clientImpl struct {
	rClient *resty.Client
}

// NewClient creates a new client for the Screener API.
func NewClient(metricHandler metrics.Handler, screenerURL string) (ScreenerClient, error) {
	client := resty.New().
		SetBaseURL(screenerURL).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			request.Header.Add(ginhelper.RequestIDHeader, uuid.New().String())
			return nil
		})

	otelresty.TraceClient(client, otelresty.WithTracerProvider(metricHandler.GetTracerProvider()))
	return &clientImpl{client}, nil
}

type blockedResponse struct {
	Blocked bool `json:"risk"`
}

func (c clientImpl) ScreenAddress(ctx context.Context, ruleset, address string) (bool, error) {
	var blockedRes blockedResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetResult(&blockedRes).
		Get(fmt.Sprintf("/%s/address/%s", ruleset, address))
	if err != nil {
		return false, fmt.Errorf("error from server: %s: %w", resp.Status(), err)
	}

	if resp.IsError() {
		return false, fmt.Errorf("error from server: %s", resp.Status())
	}

	return blockedRes.Blocked, nil
}
