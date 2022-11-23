package simulated

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ipfs/go-log"
	"github.com/stretchr/testify/assert"
	commonBackend "github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated/multibackend"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/teivah/onecontext"
	"math/big"
	"testing"
	"time"
)

var logger = log.Logger("simulated-logger")

// Backend is the simulated backend.
type Backend struct {
	t *testing.T
	// base backend is the base backend
	*base.Backend
	// backend handles commits
	simulatedBackend *multibackend.SimulatedBackend
	// faucetAddr is the address funde at genesis
	faucetAddr *keystore.Key
	// gasLimit is the block gas limit
	gasLimit uint64
	// store stores the accounts
	store *base.InMemoryKeyStore
	// chainConfig is the chainConfig for this chain
	chainConfig *params.ChainConfig
}

// Signer gets the signer for the backend.
func (s *Backend) Signer() types.Signer {
	latestBlock, err := s.BlockByNumber(s.Context(), nil)
	assert.Nil(s.T(), err)

	return types.MakeSigner(s.chainConfig, latestBlock.Number())
}

// T gets the testing object for the backend.
func (s *Backend) T() *testing.T {
	s.t.Helper()
	return s.t
}

// SetT sets the testing object on the backend.
func (s *Backend) SetT(t *testing.T) {
	t.Helper()
	s.t = t
}

// BackendName is the name of the simulated backend.
const BackendName = "SimulatedBackend"

// BackendName gets the name of SimulatedBackend.
func (s *Backend) BackendName() string {
	return BackendName
}

// EnableTenderly tells the user tenderly is not currently enabled for simulated backend type.
func (s *Backend) EnableTenderly() (enabled bool) {
	logger.Warnf("tenderly cannot be enabled on backend %s", BackendName)
	return false
}

// Commit commits pending txes to the backend. Does not thing if no txes are pending.
func (s *Backend) Commit() {
	s.simulatedBackend.Commit()
}

// EmptyBlock mines an empty block.
func (s *Backend) EmptyBlock(blockTime time.Time) {
	s.simulatedBackend.EmptyBlock(blockTime)
}

// AdjustTime adjusts the time of the most recent block.
func (s *Backend) AdjustTime(adjustment time.Duration) error {
	//nolint: wrapcheck
	return s.simulatedBackend.AdjustTime(adjustment)
}

// getFaucetTxContext gets a signed transaction from the faucet address.
func (s *Backend) getFaucetTxContext(ctx context.Context) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(s.faucetAddr.PrivateKey, s.Chain.GetBigChainID())
	assert.Nil(s.T(), err)

	//nolint: ineffassign
	gasPrice, err := s.SuggestGasPrice(ctx)
	assert.Nil(s.T(), err)

	// standard eth value tx price
	auth.GasLimit = 21000
	auth.GasPrice = gasPrice

	return auth
}

// faucetSignTx will sign a tx with the faucet addr.
func (s *Backend) faucetSignTx(tx *types.Transaction) *types.Transaction {
	tx, err := s.SignTx(tx, s.Signer(), s.faucetAddr.PrivateKey)
	assert.Nil(s.T(), err)
	return tx
}

// ChainConfig gets the chain config for the simulated backend.
func (s *Backend) ChainConfig() *params.ChainConfig {
	return s.chainConfig
}

// GetFundedAccount returns an account with the requested balance. (Note: if genesis acount has an insufficient
// balance, blocks may be mined here).
func (s *Backend) GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key {
	key := s.MockAccount()

	s.store.Store(key)

	s.FundAccount(ctx, key.Address, *requestBalance)

	return key
}

// FundAccount fundsa  new account.
func (s *Backend) FundAccount(ctx context.Context, address common.Address, amount big.Int) {
	rawTx := s.getFaucetTxContext(ctx)

	tx := s.faucetSignTx(types.NewTx(&types.LegacyTx{
		To:       &address,
		Value:    &amount,
		Gas:      rawTx.GasLimit,
		GasPrice: rawTx.GasPrice,
	}))

	assert.Nil(s.T(), s.SendTransaction(ctx, tx))
	s.Commit()

	// wait for tx confirmation
	s.WaitForConfirmation(ctx, tx)
}

// SendTransaction sends a transaction and commits it mining a new block
// in cases where you would not like to commit automatically, you can run
// s.Client().SendTransaction().
func (s *Backend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := s.Client().SendTransaction(ctx, tx)
	s.Commit()
	//nolint: wrapcheck
	return err
}

// GetAccount gets the private key for an account
// nil if the account doesn't exist.
func (s *Backend) GetAccount(address common.Address) *keystore.Key {
	return s.store.GetAccount(address)
}

// GetTxContext gets a signed transaction from full backend.
func (s *Backend) GetTxContext(ctx context.Context, address *common.Address) (res commonBackend.AuthType) {
	ctx, cancel := onecontext.Merge(ctx, s.Context())
	defer cancel()

	var acct *keystore.Key
	// TODO handle storing accounts to conform to get tx context
	if address != nil {
		acct = s.store.GetAccount(*address)
		if acct == nil {
			s.T().Errorf("could not get account %s", address.String())
			return res
		}
	} else {
		acct = s.GetFundedAccount(ctx, big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(10)))
		s.store.Store(acct)
	}

	auth, err := s.NewKeyedTransactorFromKey(acct.PrivateKey)
	assert.Nil(s.T(), err)

	gasBlock, err := s.BlockByNumber(ctx, nil)
	assert.Nil(s.T(), err)

	//nolint: ineffassign
	err = s.Chain.GasSetter().SetGasFee(ctx, auth, gasBlock.NumberU64(), gasprice.DefaultMaxPrice)
	assert.Nil(s.T(), err)

	return commonBackend.AuthType{
		TransactOpts: auth,
		PrivateKey:   acct.PrivateKey,
	}
}

// BlockGasLimit is the gas limit used for the block.
const BlockGasLimit = uint64(91712388)

// NewSimulatedBackend gets a simulated backend from geth for testing and creates an account
// with balance. ChainID is 1337.
func NewSimulatedBackend(ctx context.Context, t *testing.T) *Backend {
	t.Helper()
	return NewSimulatedBackendWithChainID(ctx, t, params.AllEthashProtocolChanges.ChainID)
}

// NewSimulatedBackendWithChainID gets a simulated backend from geth for testing and creates an account
// with balance.
func NewSimulatedBackendWithChainID(ctx context.Context, t *testing.T, chainID *big.Int) *Backend {
	t.Helper()
	return NewSimulatedBackendWithConfig(ctx, t, multibackend.NewConfigWithChainID(chainID))
}

// NewSimulatedBackendWithConfig gets a simulated backend from geth for testing and creates an account
// with balance.
func NewSimulatedBackendWithConfig(ctx context.Context, t *testing.T, config *params.ChainConfig) *Backend {
	t.Helper()
	// 100 million ether
	balance := big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(100000000))
	key := base.MockAccount(t)

	genesisAlloc := map[common.Address]core.GenesisAccount{
		key.Address: {
			Balance: balance,
		},
	}

	simulatedBackend := multibackend.NewSimulatedBackendWithConfig(genesisAlloc, BlockGasLimit, config)
	baseClient := Client{simulatedBackend}

	chn, err := chain.NewFromClient(ctx, &client.Config{
		RPCUrl:  []string{""},
		ChainID: int(config.ChainID.Uint64()),
	}, baseClient)
	chn.SetChainConfig(config)
	assert.Nil(t, err)

	baseBackend, err := base.NewBaseBackend(ctx, t, chn)
	assert.Nil(t, err)

	backend := Backend{
		Backend:          baseBackend,
		simulatedBackend: simulatedBackend,
		store:            base.NewInMemoryKeyStore(),
		chainConfig:      config,
	}
	backend.SetT(t)
	backend.Manager = nonce.NewNonceManager(ctx, &backend, backend.GetBigChainID())
	backend.faucetAddr = key
	backend.gasLimit = BlockGasLimit

	return &backend
}

var _ commonBackend.SimulatedTestBackend = &Backend{}
