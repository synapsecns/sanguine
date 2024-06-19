package chainalysis

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// EntityEndpoint is the endpoint for the entity API.
	EntityEndpoint = "/api/risk/v2/entities"
)

// Client is the interface for the Chainalysis API client. It makes requests to the Chainalysis API.
type Client interface {
	ScreenAddress(ctx context.Context, address string) (bool, error)
}

// clientImpl is the implementation of the Chainalysis API client.
type clientImpl struct {
	client     *resty.Client
	apiKey     string
	url        string
	riskLevels []string
}

// NewClient creates a new Chainalysis API client.
func NewClient(riskLevels []string, apiKey, url string) Client {
	client := resty.New().
		SetBaseURL(url).
		SetHeader("Content-Type", "application/json").
		SetHeader("Token", apiKey).
		SetTimeout(30 * time.Second)

	return &clientImpl{
		client:     client,
		apiKey:     apiKey,
		url:        url,
		riskLevels: riskLevels,
	}
}

// ScreenAddress screens an address from the Chainalysis API.
func (c *clientImpl) ScreenAddress(ctx context.Context, address string) (bool, error) {
	address = strings.ToLower(address)
	// Get the response.
	resp, err := c.client.R().
		SetContext(ctx).
		SetPathParam("address", address).
		Get(EntityEndpoint + "/" + address)
	if err != nil {
		return false, fmt.Errorf("could not get response: %w", err)
	}

	return c.handleResponse(ctx, address, resp)
}

// handleResponse takes the Chainalysis response, and depending if the address is registered or not, returns the result.
// It will retry the request if the address is not registered.
func (c clientImpl) handleResponse(ctx context.Context, address string, resp *resty.Response) (bool, error) {
	// Response could differ based on if the address is registered or not.
	var rawResponse map[string]interface{}
	var err error
	if err := json.Unmarshal(resp.Body(), &rawResponse); err != nil {
		return false, fmt.Errorf("could not unmarshal response 1: %w", err)
	}

	// If the user is not registered, register them and try again.
	if userNotRegistered(rawResponse) {
		if err = c.registerAddress(ctx, address); err != nil {
			return false, fmt.Errorf("could not register address: %w", err)
		}

		// Then try again.
		time.Sleep(2 * time.Second)
		newResp, err := c.client.R().
			SetContext(ctx).
			SetPathParam("address", address).
			Get(EntityEndpoint + "/" + address)
		if err != nil {
			return false, fmt.Errorf("could not get response: %w", err)
		}
		if err := json.Unmarshal(newResp.Body(), &rawResponse); err != nil {
			return false, fmt.Errorf("could not unmarshal response 2: %w", err)
		}
	}

	risk, _ := rawResponse["risk"].(string)
	return slices.Contains(c.riskLevels, risk), nil
}

// registerAddress registers an address in the case that we try and screen for a nonexistent address.
func (c *clientImpl) registerAddress(ctx context.Context, address string) error {
	payload := map[string]interface{}{
		"address": address,
	}
	res, err := c.client.R().SetContext(ctx).SetBody(payload).Post(EntityEndpoint)
	if err != nil {
		return fmt.Errorf("could not register address: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("could not register address: %s", res.Status())
	}

	return nil
}

var _ Client = &clientImpl{}

func userNotRegistered(rawResponse map[string]interface{}) bool {
	_, ok := rawResponse["message"]
	return ok
}
