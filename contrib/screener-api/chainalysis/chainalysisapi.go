package chainalysis

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/TwiN/gocache/v2"
	"github.com/dubonzi/otelresty"
	"github.com/valyala/fastjson"

	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
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
	client            *resty.Client
	apiKey            string
	url               string
	riskLevels        []string
	registrationCache *gocache.Cache
}

const (
	maxCacheSizeGB            = 3
	bytesInGB                 = 1024 * 1024 * 1024
	chainalysisRequestTimeout = 30 * time.Second
)

// NewClient creates a new Chainalysis API client.
func NewClient(metricHandler metrics.Handler, riskLevels []string, apiKey, url string) Client {
	client := resty.New().
		SetBaseURL(url).
		SetHeader("Content-Type", "application/json").
		SetHeader("Token", apiKey).
		SetTimeout(chainalysisRequestTimeout)

	// max cache size 3gb
	// TODO: make this configurable.
	registrationCache := gocache.NewCache().WithEvictionPolicy(gocache.LeastRecentlyUsed).WithMaxMemoryUsage(maxCacheSizeGB * bytesInGB)

	otelresty.TraceClient(client, otelresty.WithTracerProvider(metricHandler.GetTracerProvider()))

	return &clientImpl{
		client:            client,
		apiKey:            apiKey,
		url:               url,
		riskLevels:        riskLevels,
		registrationCache: registrationCache,
	}
}

// ScreenAddress screens an address from the Chainalysis API.
func (c *clientImpl) ScreenAddress(parentCtx context.Context, address string) (bool, error) {
	// make sure to cancel the context when we're done.
	// this ensures if we didn't need pessimistic register, we don't wait on it.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	address = strings.ToLower(address)

	// check the cache before we make any network calls.
	if _, ok := c.registrationCache.Get(address); ok {
		return true, nil
	}

	// we don't even wait on pessimistic register since if the address is already registered, but not in the in-memory cache
	// this will just get canceled.
	go func() {
		// Register the address in the cache.
		if err := c.pessimisticRegister(ctx, address); err != nil && !errors.Is(err, context.Canceled) {
			fmt.Printf("could not register address: %v\n", err)
		}
	}()

	result, err := c.checkBlacklist(ctx, address)
	// timeout hotfix
	if err != nil && errors.Is(err, errCouldNotGetResponse) {
		return false, nil
	}

	return result, err
}

// pessimisticRegister registers an address if its not in memory cache. This happens regardless it was registered before.
func (c *clientImpl) pessimisticRegister(ctx context.Context, address string) error {
	if _, isPresent := c.registrationCache.Get(address); !isPresent {
		if err := c.registerAddress(ctx, address); err != nil {
			return fmt.Errorf("could not register address: %w", err)
		}
	}
	return nil
}

var errCouldNotGetResponse = errors.New("could not get response")

func (c *clientImpl) checkBlacklist(ctx context.Context, address string) (bool, error) {
	var resp *resty.Response
	// Retry until the user is registered.
	err := retry.WithBackoff(ctx,
		func(ctx context.Context) (err error) {
			resp, err = c.client.R().
				SetContext(ctx).
				SetPathParam("address", address).
				Get(EntityEndpoint + "/" + address)
			if err != nil {
				return fmt.Errorf("could not get response: %w", err)
			}

			fmt.Println(string(resp.Body()))

			if resp.StatusCode() != http.StatusOK {
				return fmt.Errorf("could not get response: %s", resp.Status())
			}

			return nil
		}, retry.WithMax(time.Second), retry.WithMaxTotalTime(time.Second*10))
	if err != nil {
		return false, errCouldNotGetResponse
	}

	// address has been registered and retrieved, let's screen it and cache whether it is risky or not.
	risk := fastjson.GetString(resp.Body(), "risk")

	if slices.Contains(c.riskLevels, risk) {
		c.registrationCache.Set(address, struct{}{})
		return true, nil
	}

	return false, nil
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
