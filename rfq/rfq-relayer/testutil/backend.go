package testutil

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/client"
	"testing"
	"time"
)

const (
	AnvilMaxRetries    = 3
	AnvilMaxTimeToWait = 3
)

// NewAnvilBackend creates a new instance of the Anvil backend with a given chain.
// Because NewAnvilBackend sometimes fails or times out, it will retry a few times if it does not return within 5 seconds.
func NewAnvilBackend(ctx context.Context, chainID uint32, t *testing.T) backends.SimulatedTestBackend {
	var backend backends.SimulatedTestBackend

	for i := 0; i < AnvilMaxRetries; i++ {

		done := make(chan bool, 1)
		go func() {
			anvilOptsOrigin := anvil.NewAnvilOptionBuilder()
			anvilOptsOrigin.SetChainID(uint64(chainID))
			anvilOptsOrigin.SetBlockTime(1 * time.Second)
			backend = anvil.NewAnvilBackend(ctx, t, anvilOptsOrigin)
			done <- true
		}()

		select {
		case <-ctx.Done():
			// Context cancelled, return with nil or some error
			return nil
		case <-time.After(AnvilMaxTimeToWait * time.Second):
			// Timeout occurred, retry
			fmt.Println("Retrying Anvil backend creation, on attempt", i, "of", AnvilMaxRetries, "attempts.")
			continue
		case <-done:
			// Function returned within the timeout, return the result
			return backend
		}
	}
	// Return the last attempt's result or error
	return backend
}

// NewEVMClientFromAnvil creates a new instance of the EVM client with a given (anvil) backend.
func NewEVMClientFromAnvil(ctx context.Context, backend backends.SimulatedTestBackend, handler metrics.Handler) (client.EVM, error) {
	evmClient, err := client.DialBackend(ctx, backend.RPCAddress(), handler)
	if err != nil {
		return nil, err
	}
	return evmClient, err

}
