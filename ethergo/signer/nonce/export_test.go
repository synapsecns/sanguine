package nonce

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/chain"
)

// TestManager exports the nonce manager with additional methods required only for testing.
type TestManager interface {
	Manager
	AssertNoncesEqual(address common.Address, equalTo int64)
	GetChainID() *big.Int
}

// testManagerImpl implements the nonce manager.
type testManagerImpl struct {
	// the underlying nonce manager object
	*nonceManagerImp
	// tb is used for test assertions
	tb      testing.TB
	chainID *big.Int
}

func (t *testManagerImpl) GetChainID() *big.Int {
	return t.chainID
}

// NewTestNonceManger wraps NewNonceManager w/ newly exported methods for testing.
//
//nolint:staticcheck
func NewTestNonceManger(ctx context.Context, tb testing.TB, chain chain.Chain) TestManager {
	tb.Helper()
	manager := NewNonceManager(ctx, chain, chain.GetBigChainID())
	castManager, ok := manager.(*nonceManagerImp)
	True(tb, ok)
	return &testManagerImpl{
		nonceManagerImp: castManager,
		tb:              tb,
		chainID:         chain.GetBigChainID(),
	}
}

// AssertNoncesEqual asserts the nonce for an address is equal to the passed in nonce.
func (t *testManagerImpl) AssertNoncesEqual(address common.Address, equalTo int64) {
	nonce, err := t.GetNextNonce(address)
	Nil(t.tb, err)

	// should be 0
	equalityCheck := nonce.Cmp(big.NewInt(equalTo))
	Equal(t.tb, equalityCheck, 0)
}
