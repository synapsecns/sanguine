package ganache

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/teivah/onecontext"
	"math/big"
	"testing"
)

// Backend is the ganache backend.
type Backend struct {
	*base.Backend
	// wsURL is the websocket url of the ganache chain
	wsURL string
	// chainName is the name of the chain
	chainName string
	// keyPath is the path to the ganache keys
	keyPath string
}

// Signer gets the signer for the chain.
func (b *Backend) Signer() types.Signer {
	return types.LatestSignerForChainID(big.NewInt(int64(b.GetChainID())))
}

// NewGanacheBackend creates a new ganache backend.
// Deprecated: this will be removed in a future version in favor of anvil.
func NewGanacheBackend(ctx context.Context, t *testing.T, chainConfig *params.ChainConfig, rpcURL, chainName, keyPath string) *Backend {
	t.Helper()

	chn, err := chain.New(ctx, &client.Config{
		RPCUrl:  []string{rpcURL},
		ChainID: int(chainConfig.ChainID.Uint64()),
	})
	chn.SetChainConfig(chainConfig)
	assert.Nil(t, err)

	baseBackend, err := base.NewBaseBackend(ctx, t, chn)
	assert.Nil(t, err)

	backend := &Backend{
		Backend:   baseBackend,
		wsURL:     rpcURL,
		chainName: chainName,
		keyPath:   keyPath,
	}

	return backend
}

// GetFundedAccount gets a funded account.
func (b *Backend) GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key {
	panic("new account creation is not yet supoprted for ganache")
}

// FundAccount funds an account with an amount.
func (b *Backend) FundAccount(ctx context.Context, address common.Address, amount big.Int) {
	auth := b.GetTxContext(ctx, nil)

	tx := types.NewTx(&types.LegacyTx{
		To:       &address,
		Value:    &amount,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
	})

	tx, err := b.SignTx(tx, b.Signer(), auth.PrivateKey)
	assert.Nil(b.T(), err)

	err = b.SendTransaction(ctx, tx)
	assert.Nil(b.T(), err)

	b.WaitForConfirmation(ctx, tx)
}

// GanacheBackendName is the name of the geth backend.
const GanacheBackendName = "Ganache"

// BackendName gets the backend name.
func (b *Backend) BackendName() string {
	return GanacheBackendName
}

// GetTxContext gets the transaction context.
func (b *Backend) GetTxContext(ctx context.Context, address *common.Address) (auth backends.AuthType) {
	ctx, cancel := onecontext.Merge(ctx, b.Context())
	defer cancel()

	//nolint: gosec
	keyChain, err := ParseAddresses(b.keyPath)
	assert.Nil(b.T(), err)

	// get the keys in ganache
	for _, privKey := range keyChain.PrivateKeys {
		// convert the private keys to an ecdsa key
		privateKey, err := crypto.HexToECDSA(privKey)
		assert.Nil(b.T(), err)

		// get the block by number
		blck, err := b.Client().BlockByNumber(ctx, nil)
		assert.Nil(b.T(), err)
		// set the gas limit to the total gas limit for the last block
		gasLimit := blck.GasLimit()

		//nolint: staticcheck
		tmpAuth, err := b.NewKeyedTransactorFromKey(privateKey)
		assert.Nil(b.T(), err)

		auth = backends.AuthType{TransactOpts: tmpAuth, PrivateKey: privateKey}

		auth.Value = big.NewInt(0) // in wei
		auth.GasLimit = gasLimit   // in unitser
		auth.GasPrice = big.NewInt(1)

		if address != nil {
			if address.String() == auth.From.String() {
				return auth
			}
		} else {
			return auth
		}
	}
	return auth
}

var _ backends.SimulatedTestBackend = &Backend{}
