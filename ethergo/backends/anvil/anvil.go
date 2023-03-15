package anvil

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethCore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/google/uuid"
	"github.com/ipfs/go-log"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/teivah/onecontext"
	"math/big"
	"strings"
	"sync"
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
	// impersonationMux is used to lock the impersonation
	impersonationMux sync.Mutex
}

const backendName = "anvil"

// BackendName returns the name of the backend.
func (f *Backend) BackendName() string {
	return backendName
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

	go func() {
		err = tailLogs(ctx, resource, pool, true)
		logger.Warn(err)
	}()

	// Docker will hard kill the container in 4000 seconds (this is a test env).
	// containers should be removed on their own, but this is a safety net.
	// to prevent old containers from piling up, we set a timeout to remove the container.
	const resourceLifetime = uint(600)

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

	go func() {

	}()
	return &backend
}

type bufCloser struct {
	*bytes.Buffer
}

// Close implements io.Closer.
func (b bufCloser) Close() error {
	return nil
}

var logger = log.Logger("anvil-docker")

// tailLogs tails the logs of a docker container.
func tailLogs(ctx context.Context, resource *dockertest.Resource, pool *dockertest.Pool, follow bool) error {
	outStream := bufCloser{bytes.NewBuffer(nil)}
	errStream := bufCloser{bytes.NewBuffer(nil)}

	opts := docker.LogsOptions{
		Context: ctx,

		Stderr:      true,
		Stdout:      true,
		Follow:      follow,
		Timestamps:  true,
		RawTerminal: true,

		Container: resource.Container.ID,

		ErrorStream:  errStream,
		OutputStream: outStream,
	}

	_, err := processlog.StartLogs(processlog.WithStdOut(outStream), processlog.WithStdErr(errStream), processlog.WithCtx(ctx))
	if err != nil {
		return fmt.Errorf("failed to get container logs: %w", err)
	}

	//nolint: wrapcheck
	return pool.Client.Logs(opts)
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

	f.FundAccount(ctx, key.Address, *requestBalance)

	return key
}

// GetTxContext gets the tx context for the given address.
// TODO: dedupe w/ geth.
func (f *Backend) GetTxContext(ctx context.Context, address *common.Address) (res backends.AuthType) {
	ctx, cancel := onecontext.Merge(ctx, f.Context())
	defer cancel()

	var acct *keystore.Key
	// TODO handle storing accounts to conform to get tx context
	if address != nil {
		acct = f.store.GetAccount(*address)
		if acct == nil {
			f.T().Errorf("could not get account %s", address.String())
			return res
		}
	} else {
		acct = f.GetFundedAccount(ctx, big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(10)))
		f.store.Store(acct)
	}

	auth, err := f.NewKeyedTransactorFromKey(acct.PrivateKey)
	assert.Nil(f.T(), err)

	latestBlock, err := f.BlockByNumber(ctx, nil)
	assert.Nil(f.T(), err)

	err = f.Chain.GasSetter().SetGasFee(ctx, auth, latestBlock.NumberU64(), core.CopyBigInt(gasprice.DefaultMaxPrice))
	assert.Nil(f.T(), err)

	auth.GasLimit = ethCore.DeveloperGenesisBlock(0, gasLimit, acct.Address).GasLimit / 2

	return backends.AuthType{
		TransactOpts: auth,
		PrivateKey:   acct.PrivateKey,
	}
}

// ImpersonateAccount impersonates an account.
//
// Note *any* other call made to the backend will impersonate while this is being called
// in a future version, we'll wrap something like omnirpc to prevent other transactions submission calls from taking place
// in the meantime, this may cause race conditions.
//
// We also print a warning message to the console as an added precaution.
func (f *Backend) ImpersonateAccount(ctx context.Context, address common.Address, transact func(opts *bind.TransactOpts) *types.Transaction) {
	f.impersonationMux.Lock()
	defer f.impersonationMux.Unlock()

	f.warnImpersonation()

	anvilClient, err := Dial(ctx, f.RPCAddress())
	assert.Nil(f.T(), err)

	err = anvilClient.ImpersonateAccount(ctx, address)
	assert.Nil(f.T(), err)

	tx := transact(&bind.TransactOpts{
		Context: ctx,
		From:    address,
		Signer:  ImpersonatedSigner,
		NoSend:  true,
	})

	err = anvilClient.SendUnsignedTransaction(ctx, address, tx)

	defer func() {
		err = anvilClient.StopImpersonatingAccount(ctx, address)
		assert.Nil(f.T(), err)
	}()
}

func (f *Backend) warnImpersonation() {
	logOnce.Do(func() {
		f.T().Logf(`
				Using Account Impersonation.
				WARNING: This cannot be called concurrently with other impersonation calls.
				Please make sure your callers are concurrency safe against account impersonation.
				`)
	})
}

// ImpersonatedSigner is a signer that does nothing for use in account impersonation w/ contracts.
func ImpersonatedSigner(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return transaction, nil
}

// logOnce is used to log the impersonation warning message once.
// this is a global variable to prevent the message from being logged multiple times.
// normally, global variables are strongly discouraged, but we make an exception here
// considering how unexpected behavior can be if impersonate account is not used correctly.
var logOnce sync.Once

var _ backends.SimulatedTestBackend = &Backend{}
