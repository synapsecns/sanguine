package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/jftuga/ellipsis"
	backoffHelper "github.com/jpillora/backoff"
	"gitlab.com/1f320/x/duration"
	"io"
	"net/http"
	"strconv"
	"time"
)

// PublicRPCMapURL is the url we pull the rpc list from.
// TODO: this has some rate-limits, they're relatively aggressive but something like gitcdn.xyz would be good here.
const PublicRPCMapURL = "https://raw.githubusercontent.com/DefiLlama/chainlist/master/constants/extraRpcs.json"

// GetPublicRPCConfig gets the rpc map from defillama. This should be done at startup time.
// this will retry on a backoffHelper until context cancellation.
func GetPublicRPCConfig(ctx context.Context) (c Config, err error) {
	backoff := &backoffHelper.Backoff{
		Factor: 1.3,
		Jitter: true,
		Min:    time.Second * 1,
		Max:    time.Second * 10,
	}

	var waitTime time.Duration
	httpClient := &http.Client{}
	for {
		select {
		case <-ctx.Done():
			return Config{}, fmt.Errorf("could not get rpc map: %w", ctx.Err())
		case <-time.After(waitTime):
			waitTime = backoff.Duration()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, PublicRPCMapURL, nil)
			if err != nil {
				return Config{}, fmt.Errorf("could not create request: %w", err)
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				logger.Errorf("could not retrieve rpc list from %s, waiting %s before trying again (error: %v)", PublicRPCMapURL, duration.Format(waitTime), err)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Errorf("could not read body from %s, waiting %s before trying again (error: %v)", PublicRPCMapURL, duration.Format(waitTime), err)
				continue
			}

			_ = resp.Body.Close()

			c, err = parseConfig(body)
			if err != nil {
				logger.Error(err)
				continue
			}

			return c, nil
		}
	}
}

// parseConfig parses a chain map from a json payload.
func parseConfig(rawData []byte) (c Config, err error) {
	c = Config{
		Chains: make(map[uint32]ChainConfig),
	}

	// iterate over chain ids which are strings
	err = jsonparser.ObjectEach(rawData, func(key []byte, value []byte, dataType jsonparser.ValueType, _ int) error {
		if dataType != jsonparser.Object {
			return fmt.Errorf("expected %s got %s when parsing %s", jsonparser.Object, dataType, ellipsis.Shorten(string(rawData), 10))
		}

		chainID, err := strconv.Atoi(string(key))
		if err != nil {
			return fmt.Errorf("could not parse %s to int", key)
		}

		rawRPCList, dataType, _, err := jsonparser.Get(value, "rpcs")
		// skip this key
		if err != nil {
			logger.Debugf("key %s missing", key)
			//nolint: nilerr
			return nil
		}

		if dataType != jsonparser.Array {
			return fmt.Errorf("expected %s got %s when parsing %s", jsonparser.Array, dataType, ellipsis.Shorten(string(value), 10))
		}

		var rpcArr []string
		err = json.Unmarshal(rawRPCList, &rpcArr)
		if err != nil {
			return fmt.Errorf("could not unmarshal array: %w", err)
		}

		// skip nil
		if len(rpcArr) == 0 {
			return nil
		}

		// public rpcs always use
		c.Chains[uint32(chainID)] = ChainConfig{
			RPCs:   rpcArr,
			Checks: 1,
		}

		return nil
	})

	if err != nil {
		return c, fmt.Errorf("could not parse rpc map: %w", err)
	}

	return c, nil
}
