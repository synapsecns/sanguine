package etherscan

import (
	"context"
	"fmt"
	"github.com/nanmu42/etherscan-api"
	"github.com/synapsecns/sanguine/core/config"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Client implements an etherscan client.
type Client struct {
	*etherscan.Client
	// rateLimiter is a file based rate limiter
	rateLimiter *fileRateLimiter
}

// timeout is the http timeout for a request.
const timeout = time.Second * 30

// newEtherscanABIClient creates a new etherscan client.
func newEtherscanABIClient(parentCtx context.Context, chainID uint32, url string, disableRateLimiter bool) (*Client, error) {
	var client Client
	ctx, cancel := context.WithCancel(parentCtx)

	apiKeyEnv := strings.ToUpper(fmt.Sprintf("%d_KEY", chainID))
	apiKey := os.Getenv(apiKeyEnv)

	customization := etherscan.Customization{
		Client: &http.Client{
			Timeout: timeout,
		},
		BaseURL: url,
	}

	// waitBetweenRequest is how long to wait between requests. If an analytics key is enabled, rate limiting is disabled
	waitBetweenRequests := time.Second * 5
	rateLimiterEnabled := apiKey == "" && !disableRateLimiter

	if rateLimiterEnabled {
		configDir, err := config.GetConfigDir()
		if err != nil {
			cancel()
			return nil, fmt.Errorf("could not create file rate limiter: %w", err)
		}

		rateLimitDir := filepath.Join(configDir, strconv.Itoa(int(chainID)))

		client.rateLimiter, err = newFileRateLimiter(ctx, rateLimitDir, waitBetweenRequests)
		if err != nil {
			cancel()
			return nil, fmt.Errorf("could not create file rate limiter: %w", err)
		}

		customization.BeforeRequest = func(_, _ string, _ map[string]interface{}) error {
			_, err := client.rateLimiter.obtainLock(ctx)
			return err
		}

		customization.AfterRequest = func(_, action string, _ map[string]interface{}, _ interface{}, _ error) {
			_, err = client.rateLimiter.releaseLock()
			if err != nil {
				logger.Error(err)
				cancel()
			}
		}
	} else {
		// context cancellation is handled by the parent
		_ = cancel
	}

	client.Client = etherscan.NewCustomized(customization)
	return &client, nil
}
