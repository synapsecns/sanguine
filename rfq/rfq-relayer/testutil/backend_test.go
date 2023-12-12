package testutil_test

import (
	"context"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/testutil"
)

func TestBackend(t *testing.T) {
	testCtx := context.Background()
	testChainID := uint32(42161)

	// Test Anvil
	anvilBackend := testutil.NewAnvilBackend(testCtx, t, testChainID)
	NotNil(t, anvilBackend)
	Equal(t, "anvil", anvilBackend.BackendName())
	chain, err := anvilBackend.ChainID(testCtx)
	Nil(t, err)
	Equal(t, testChainID, uint32(chain.Int64()))

	// Test EVM client
	handler := metrics.NewNullHandler()
	evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, handler)
	Nil(t, err)
	NotNil(t, evmClient)
	clientChain, err := evmClient.ChainID(testCtx)
	Nil(t, err)
	Equal(t, testChainID, uint32(clientChain.Int64()))
}
