package nonce_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core"
	evmMocks "github.com/synapsecns/sanguine/ethergo/chain/mocks"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"k8s.io/apimachinery/pkg/util/sets"
	"math/big"
	"sync"
	"testing"
)

// MockAccount implements some methods to make testing easier.
type MockAccount struct {
	// Key is the account key
	*keystore.Key
	// nonceManager is the nonce manager, used for testing in methods
	nonceManager nonce.TestManager
	// tb is the testing objcet
	tb testing.TB
	// currentNonce is the correct current nonce
	currentNonce int64
	// nonceSet contains the nonce set
	nonceSet sets.Int64
	// muxSet contains the mutex set
	muxSet sync.Mutex
}

// NewMockAccount gets a new mock account.
func (n NonceSuite) NewMockAccount(nonceManager nonce.TestManager) *MockAccount {
	key := mocks.MockAccount(n.T())

	return &MockAccount{
		nonceManager: nonceManager,
		tb:           n.T(),
		Key:          key,
		currentNonce: 0,
		nonceSet:     sets.NewInt64(),
	}
}

// CreateMockAccounts creates count mock accounts.
func (n NonceSuite) CreateMockAccounts(nonceManager nonce.TestManager, count int) (mockAccounts []*MockAccount) {
	for i := 0; i < count; i++ {
		mockAccounts = append(mockAccounts, n.NewMockAccount(nonceManager))
	}
	Equal(n.T(), len(mockAccounts), count)
	return mockAccounts
}

// GetSignedtx gets a signed tx for the given account. It checks that the nonce does not change
// before signing occurs. This tests both signer types (testing Transactor on even and SignTx on odd).
func (m *MockAccount) GetSignedTx() {
	presignNonce, err := m.nonceManager.GetNextNonce(m.Address)
	Nil(m.tb, err)

	// save the nonce
	currentNonce := uint64(m.currentNonce)

	rawTX := types.NewTx(&types.LegacyTx{
		To:       &m.Address,
		Gas:      params.GWei,
		GasPrice: big.NewInt(params.GWei),
	})

	var signedTX *types.Transaction
	if presignNonce.Uint64()%2 == 0 {
		transactor, err := m.nonceManager.NewKeyedTransactorFromKey(m.Key.PrivateKey)
		Nil(m.tb, err)

		signedTX, err = transactor.Signer(m.Address, rawTX)
		Nil(m.tb, err)
	} else {
		signedTX, err = m.nonceManager.SignTx(rawTX, types.LatestSignerForChainID(m.nonceManager.GetChainID()), m.Key.PrivateKey)
		Nil(m.tb, err)
	}

	fmt.Printf("nonce %d and %d\n", signedTX.Nonce(), m.currentNonce)
	// first one should be 0
	if currentNonce != 0 {
		Greater(m.tb, signedTX.Nonce(), currentNonce)
	}

	m.currentNonce = int64(signedTX.Nonce())

	// use the mux set to make sure we don't have conflicts
	m.muxSet.Lock()
	m.nonceSet.Insert(int64(signedTX.Nonce()))
	m.muxSet.Unlock()
}

func (n NonceSuite) TestNonceManager() {
	mockChain := evmMocks.Chain{}
	// since this should only be alled on the first try, this always returns 0
	mockChain.On("PendingNonceAt", mock.Anything, mock.Anything).Return(uint64(0), nil)
	mockChain.On("GetBigChainID").Return(core.CopyBigInt(params.MainnetChainConfig.ChainID))

	nonceManager := nonce.NewTestNonceManger(n.GetTestContext(), n.T(), &mockChain)

	// generate nonceCount nonces for accountCount accounts
	const accountCount = 10
	const nonceCount = 50
	// make sure this works paralellized
	mockAccounts := n.CreateMockAccounts(nonceManager, accountCount)

	// paralellize to test locking mechanism
	var wg sync.WaitGroup
	for _, mockAccount := range mockAccounts {
		for i := 0; i < nonceCount; i++ {
			wg.Add(1)
			go func(mockAccount *MockAccount) {
				mockAccount.GetSignedTx()
				wg.Done()
			}(mockAccount)
		}
	}

	wg.Wait()

	// make sure all nonces are present
	for _, mockAccount := range mockAccounts {
		Equal(n.T(), nonceCount, mockAccount.nonceSet.Len())
		for i := 0; i < nonceCount; i++ {
			True(n.T(), mockAccount.nonceSet.Has(int64(i)))
		}
	}
}
