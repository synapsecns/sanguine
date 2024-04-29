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
	ScreenAddress(ctx context.Context, ruleset, address string) (blocked bool, err error)
	BlacklistAddress(ctx context.Context, body BlackListBody) (string, error)
}

type clientImpl struct {
	rClient *resty.Client
}

// NewClient creates a new client for the Screener API.
func NewClient(metricHandler metrics.Handler, screenerURL string) (ScreenerClient, error) {
	client := resty.New().
		SetBaseURL(screenerURL).
		OnBeforeRequest(func(_ *resty.Client, request *resty.Request) error {
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

type BlackListBody struct {
	TypeReq string `json:"type" binding:"required"`
	Id      string `json:"id" binding:"required"`
	Data    string `json:"data"`
	Address string `json:"address"`
	Network string `json:"network"`
	Tag     string `json:"tag"`
	Remark  string `json:"remark"`
}

type blacklistResponse struct {
	Status string `json:"status"`
}

func (c clientImpl) BlacklistAddress(ctx context.Context, body BlackListBody) (string, error) {
	var blacklistRes blacklistResponse

	resp, err := c.rClient.R().
		SetContext(ctx).
		SetResult(&blacklistRes).
		SetBody(body).
		Post("/api/data/sync/")

	if err != nil {
		return "", fmt.Errorf("error from server: %s: %w", resp.Status(), err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("error from server: %s", resp.Status())
	}

	return blacklistRes.Status, nil
}

// NewNoOpClient creates a new no-op client for the Screener API.
// it returns false for every address.
func NewNoOpClient() (ScreenerClient, error) {
	return &noOpClient{}, nil
}

type noOpClient struct{}

func (n noOpClient) ScreenAddress(_ context.Context, _, _ string) (bool, error) {
	return false, nil
}

func (n noOpClient) BlacklistAddress(_ context.Context, _ BlackListBody) (string, error) {
	return "", nil
}

var _ ScreenerClient = noOpClient{}
