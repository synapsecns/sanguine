// Package client provides a client for the screener-api.
package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
)

var (
	// BlacklistEndpoint is the endpoint for blacklisting an address.
	BlacklistEndpoint = "/api/data/sync/"
)

// ScreenerClient is an interface for the Screener API.
type ScreenerClient interface {
	ScreenAddress(ctx context.Context, ruleset, address string) (blocked bool, err error)
	BlacklistAddress(ctx context.Context, appsecret string, appid string, body BlackListBody) (string, error)
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

// ScreenAddress checks if an address is blocked by the screener.
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

// BlackListBody is the json payload that represents a blacklisted address.
type BlackListBody struct {
	TypeReq string `json:"typereq"`
	ID      string `json:"id"`
	Data    string `json:"data"`
	Address string `json:"address"`
	Network string `json:"network"`
	Tag     string `json:"tag"`
	Remark  string `json:"remark"`
}

type blacklistResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func (c clientImpl) BlacklistAddress(ctx context.Context, appsecret string, appid string, body BlackListBody) (string, error) {
	var blacklistRes blacklistResponse

	nonce := strings.ReplaceAll(uuid.New().String(), "-", "")[:32]
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	queryString := ""

	signature := GenerateSignature(appsecret, appid, timestamp, nonce, queryString, body)

	resp, err := c.rClient.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeader("appid", appid).
		SetHeader("timestamp", timestamp).
		SetHeader("nonce", nonce).
		SetHeader("queryString", queryString).
		SetHeader("signature", signature).
		SetResult(&blacklistRes).
		SetBody(body).
		Post(BlacklistEndpoint)

	if err != nil {
		return resp.Status(), fmt.Errorf("error from server: %s: %w", resp.String(), err)
	}

	if resp.IsError() {
		return resp.Status(), fmt.Errorf("error from server: %s", resp.String())
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

func (n noOpClient) BlacklistAddress(_ context.Context, _ string, _ string, _ BlackListBody) (string, error) {
	return "", nil
}

// GenerateSignature generates a signature for the request.
func GenerateSignature(secret string,
	appid string,
	timestamp string,
	nonce string,
	queryString string,
	body BlackListBody,
) string {
	key := []byte(secret)

	// Concatenate the body.
	message := fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		appid,
		timestamp,
		nonce,
		"POST",
		BlacklistEndpoint,
		queryString,
		body,
	)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

var _ ScreenerClient = noOpClient{}
