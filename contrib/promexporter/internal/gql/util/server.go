package util

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/retry"
	"net/http"
	"time"
)

// WaitForStart waits for the connection to be ready on a ginhelper like server
// times out after serverStartTimeout.
func WaitForStart(ctx context.Context, port int) error {
	err := retry.WithBackoff(ctx, func(ctx context.Context) error {
		client := http.DefaultClient
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:%d%s", port, ginhelper.HealthCheck), nil)
		if err != nil {
			return fmt.Errorf("could not create request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("could not get response: %w", err)
		}

		if resp != nil {
			_ = resp.Body.Close()
		}

		return nil
	}, retry.WithMax(serverStartTimeout))
	if err != nil {
		return fmt.Errorf("could not start gqlServer: %w", err)
	}
	return nil
}

const serverStartTimeout = time.Second * 5
