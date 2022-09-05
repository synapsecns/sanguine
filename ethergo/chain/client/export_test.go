package client

import (
	"context"
	"fmt"
	. "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// ChainConfigs exports chainConfigs for testing.
var ChainConfigs = chainConfigs

func init() {
	defaultConnectionResetTimeout = connectionResetTimeout
}

// defaultConnectionResetTimeout stores the connection reset timeout so it can be reset before each test.
var defaultConnectionResetTimeout time.Duration

// GetDefaultResetTimeout gets the default connection reset timeout.
func GetDefaultResetTimeout() time.Duration {
	return defaultConnectionResetTimeout
}

// SetResetTimeout sets the connection reset timeout for testing.
func SetResetTimeout(timeout time.Duration) {
	connectionResetTimeout = timeout
}

// TestClient exports client methods for testing.
type TestClient interface {
	EVMClient
	StartConnectionResetTicker(ctx context.Context)
	AttemptReconnect() error
}

// NewTestClient gets a client with additional methods exported for testing.
func NewTestClient(ctx context.Context, tb testing.TB, rpcURL string) (TestClient, error) {
	tb.Helper()

	testClient, err := NewClient(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("could not create client: %w", err)
	}

	marshalledClient, ok := testClient.(*clientImpl)
	True(tb, ok)
	return marshalledClient, nil
}
