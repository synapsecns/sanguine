// Package client provides a client for the screener-api.
package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
)

// ScreenerClient is an interface for the Screener API.
type ScreenerClient interface {
	ScreenAddress(ctx context.Context, address string) (blocked bool, err error)
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

type notFoundResponse struct {
	Message string `json:"message"`
}

// ScreenAddress checks if an address is blocked by the screener API.
func (c clientImpl) ScreenAddress(ctx context.Context, address string) (bool, error) {
	var blockedRes blockedResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetResult(&blockedRes).
		Get("/address/" + address)
	if err != nil {
		return false, fmt.Errorf("error from server: %s: %w", resp.Status(), err)
	}

	if resp.IsError() {
		// The address was not found
		if err := json.Unmarshal(resp.Body(), &notFoundResponse{}); err == nil {
			return false, nil
		}

		return false, fmt.Errorf("error from server: %s %w", resp, err)
	}

	return blockedRes.Blocked, nil
}

// BlacklistAddress blacklists an address with the screener API.
func (c clientImpl) BlacklistAddress(ctx context.Context, appsecret string, appid string, body BlackListBody) (string, error) {
	var blacklistRes blacklistResponse

	nonce := strings.ReplaceAll(uuid.New().String(), "-", "")[:32]
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	bodyBz, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("error marshaling body: %w", err)
	}

	bodyStr, err := core.BytesToJSONString(bodyBz)
	if err != nil {
		return "", fmt.Errorf("could not convert bytes to json: %w", err)
	}

	message := fmt.Sprintf("%s;%s;%s;%s;%s;%s",
		appid, timestamp, nonce, "POST", "/api/data/sync", bodyStr)

	signature := GenerateSignature(appsecret, message)

	resp, err := c.rClient.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Signature-appid", appid).
		SetHeader("X-Signature-timestamp", timestamp).
		SetHeader("X-Signature-nonce", nonce).
		SetHeader("X-Signature-signature", signature).
		SetBody(body).
		SetResult(&blacklistRes).
		Post("/api/data/sync/")

	if err != nil {
		return resp.Status(), fmt.Errorf("error from server: %s: %w", resp.String(), err)
	}

	if resp.IsError() {
		return resp.Status(), fmt.Errorf("error from server: %s", resp.String())
	}

	return blacklistRes.Status, nil
}

// BlackListBody is the json payload that represents a blacklisted address.
type BlackListBody struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Data Data   `json:"data"`
}

// Data is the data field in the BlackListBody.
type Data struct {
	Address string `json:"address"`
	Network string `json:"network"`
	Tag     string `json:"tag"`
	Remark  string `json:"remark"`
}

type blacklistResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// GenerateSignature generates a signature for the request.
func GenerateSignature(secret, message string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// NewNoOpClient creates a new no-op client for the Screener API.
// it returns false for every address.
func NewNoOpClient() (ScreenerClient, error) {
	return &noOpClient{}, nil
}

type noOpClient struct{}

func (n noOpClient) ScreenAddress(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (n noOpClient) RegisterAddress(_ context.Context, _ string) error {
	return nil
}

func (n noOpClient) BlacklistAddress(_ context.Context, _ string, _ string, _ BlackListBody) (string, error) {
	return "", nil
}

var _ ScreenerClient = noOpClient{}
