package anvil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethCore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/google/uuid"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/teivah/onecontext"
	"math/big"
	"strings"
	"testing"
	"time"
)

// Backend contains the anvil test backend.
type Backend struct {
	*base.Backend
	wallets []wallet.Wallet
	// walletMux is used to lock the wallets
	walletMux mapmutex.StringerMapMutex
	// store stores the accounts
	store *base.InMemoryKeyStore
	// chainConfig is the chain config
	chainConfig *params.ChainConfig
}

const gasLimit = 10000000

// NewAnvilBackend creates a test anvil backend.
func NewAnvilBackend(ctx context.Context, t *testing.T, args *OptionBuilder) *Backend {
	t.Helper()

	pool, err := dockertest.NewPool("")
	assert.Nil(t, err)

	pool.MaxWait = time.Minute * 2
	if err != nil {
		assert.Nil(t, err)
	}

	commandArgs, err := args.Build()
	assert.Nil(t, err)

	runOptions := &dockertest.RunOptions{
		Repository: "ghcr.io/foundry-rs/foundry",
		Tag:        "latest",
		Cmd:        []string{strings.Join(append([]string{"anvil"}, commandArgs...), " ")},
		Labels: map[string]string{
			"test-id": uuid.New().String(),
		},
		ExposedPorts: []string{"8545"},
	}

	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(t, err)

	// Docker will hard kill the container in 4000 seconds (this is a test env).
	// containers should be removed on their own, but this is a safety net.
	// to prevent old containers from piling up, we set a timeout to remove the container.
	const resourceLifetime = uint(4000)

	assert.Nil(t, resource.Expire(resourceLifetime))

	address := fmt.Sprintf("%s:%s", "http://localhost", resource.GetPort("8545/tcp"))

	var chainID *big.Int
	if err := pool.Retry(func() error {
		rpcClient, err := ethclient.DialContext(ctx, address)
		if err != nil {
			return fmt.Errorf("failed to connect")
		}
		chainID, err = rpcClient.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get chain id: %w", err)
		}
		return nil
	}); err != nil {
		assert.Nil(t, err)
	}

	chainConfig := args.GetHardfork().ToChainConfig(chainID)

	chn, err := chain.New(ctx, &client.Config{
		RPCUrl:  []string{address},
		ChainID: int(chainConfig.ChainID.Int64()),
	})
	chn.SetChainConfig(chainConfig)
	assert.Nil(t, err)

	baseBackend, err := base.NewBaseBackend(ctx, t, chn)
	assert.Nil(t, err)

	wallets, err := makeWallets(args)
	if err != nil {
		assert.Nil(t, err)
	}

	backend := Backend{
		Backend:     baseBackend,
		wallets:     wallets,
		walletMux:   mapmutex.NewStringerMapMutex(),
		store:       base.NewInMemoryKeyStore(),
		chainConfig: chainConfig,
	}

	go func() {
		<-ctx.Done()
		_ = pool.Purge(resource)
	}()
	return &backend
}

// makeWallets creates a list of preseeded wallets w/ balances.
// these are used for funding accounts.
//
// this may break in certain cases where we've ran out of funds.
func makeWallets(args *OptionBuilder) (wallets []wallet.Wallet, _ error) {
	derivationPath := args.GetDerivationPath()
	derivIter := accounts.DefaultIterator(derivationPath)
	maxAccounts := args.GetAccounts()
	for i := 0; i < int(maxAccounts); i++ {
		account := derivIter()

		wall, err := wallet.FromSeedPhrase(args.GetMnemonic(), account)
		if err != nil {
			return []wallet.Wallet{}, fmt.Errorf("could not get seed phrase: %w", err)
		}

		wallets = append(wallets, wall)
	}
	return
}

// ChainConfig gets the chain config.
func (f *Backend) ChainConfig() *params.ChainConfig {
	return f.chainConfig
}

// Signer gets the signer for the chain.
func (f *Backend) Signer() types.Signer {
	latestBlock, err := f.BlockByNumber(f.Context(), nil)
	assert.Nil(f.T(), err)

	return types.MakeSigner(f.ChainConfig(), latestBlock.Number())
}

// FundAccount funds an account with the given amount.
func (f *Backend) FundAccount(ctx context.Context, address common.Address, amount big.Int) {
	ctx, cancel := onecontext.Merge(ctx, f.Context())
	defer cancel()

	// get a funding wallet
	fundingWallet, err := core.RandomItem(f.wallets)
	assert.Nil(f.T(), err)

	// lock this wallet to avoid nonce issues (just in case, should be handled by nonce locker)
	locker := f.walletMux.Lock(fundingWallet)
	defer locker.Unlock()

	auth, err := f.NewKeyedTransactorFromKey(fundingWallet.PrivateKey())
	assert.Nil(f.T(), err)

	auth.GasPrice, err = f.Client().SuggestGasPrice(ctx)
	assert.Nil(f.T(), err)

	auth.GasLimit = ethCore.DeveloperGenesisBlock(0, gasLimit, auth.From).GasLimit

	tx, err := f.Backend.SignTx(types.NewTx(&types.LegacyTx{
		To:       &address,
		Value:    &amount,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
	}), f.Signer(), fundingWallet.PrivateKey())
	assert.Nil(f.T(), err)

	assert.Nil(f.T(), f.Client().SendTransaction(f.Context(), tx))

	f.WaitForConfirmation(ctx, tx)
}

// GetFundedAccount gets a funded account.
func (f *Backend) GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key {
	key := f.MockAccount()

	f.store.Store(key)

	panic("implement me")
}
