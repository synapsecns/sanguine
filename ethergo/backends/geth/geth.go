package geth

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	ethCore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/graphql"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/chain"
	legacyClient "github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/teivah/onecontext"
	"k8s.io/apimachinery/pkg/util/wait"
)

const gasLimit = 10000000

// NewEmbeddedBackendForChainID gets the embedded backend for a specific chain id.
func NewEmbeddedBackendForChainID(ctx context.Context, t *testing.T, chainID *big.Int) *Backend {
	t.Helper()
	// get the default config
	config := ethCore.DeveloperGenesisBlock(0, gasLimit, common.Address{}).Config
	config.ChainID = chainID
	return NewEmbeddedBackendWithConfig(ctx, t, config)
}

// NewEmbeddedBackend gets the default embedded backend.
func NewEmbeddedBackend(ctx context.Context, t *testing.T) *Backend {
	t.Helper()
	return NewEmbeddedBackendForChainID(ctx, t, params.AllCliqueProtocolChanges.ChainID)
}

// see: https://git.io/JGsC1
// taken from geth, used to speed up tests.
const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

// NewEmbeddedBackendWithConfig gets a full node backend to test against and returns the rpc url
// can be canceled with the past in context object.
func NewEmbeddedBackendWithConfig(ctx context.Context, t *testing.T, config *params.ChainConfig) *Backend {
	t.Helper()
	setupEthLogger()

	embedded := Backend{}

	logger.Debug("creating eth node")

	kstr := keystore.NewKeyStore(filet.TmpDir(t, ""), veryLightScryptN, veryLightScryptP)
	password := gofakeit.Password(true, true, true, false, false, 10)
	acct, err := kstr.NewAccount(password)
	assert.Nil(t, err)

	data, err := os.ReadFile(acct.URL.Path)
	assert.Nil(t, err)

	key, err := keystore.DecryptKey(data, password)
	assert.Nil(t, err)

	embedded.faucetAddr = key
	assert.Nil(t, err)

	embedded.Node, err = node.New(makeNodeConfig(t))
	assert.Nil(t, err)

	ethConfig := makeEthConfig(embedded.faucetAddr.Address, config)

	embedded.ethBackend, err = eth.New(embedded.Node, ethConfig)
	assert.Nil(t, err)

	embedded.Node.RegisterAPIs(toPublic(tracers.APIs(embedded.ethBackend.APIBackend)))

	// Configure log filter RPC API.
	filterSystem := utils.RegisterFilterAPI(embedded.Node, embedded.ethBackend.APIBackend, ethConfig)

	// TODO: this service should be optional. We use it enough in debugging right now to enable globally
	err = graphql.New(embedded.Node, embedded.ethBackend.APIBackend, filterSystem, embedded.Node.Config().GraphQLCors, embedded.Node.Config().GraphQLVirtualHosts)
	assert.Nil(t, err)

	assert.Nil(t, embedded.Node.Start())

	// import the faucet account (etherbase)
	// add the scrypt test backend
	keystoreBackend := keystore.NewKeyStore(filet.TmpDir(t, ""), veryLightScryptN, veryLightScryptP)

	file, err := os.ReadFile(acct.URL.Path)
	assert.Nil(t, err)

	acct, err = keystoreBackend.Import(file, password, password)
	assert.Nil(t, err)

	assert.Nil(t, keystoreBackend.Unlock(acct, password))

	// set the backend
	embedded.ethBackend.AccountManager().AddBackend(keystoreBackend)
	embedded.ethBackend.SetEtherbase(acct.Address)

	embedded.ethBackend.TxPool().SetGasPrice(big.NewInt(0))
	err = embedded.ethBackend.APIBackend.StartMining(0)
	assert.Nil(t, err)

	// add debugger for node stop
	go func() {
		embedded.Node.Wait()
		logger.Debug("eth node stopped")
	}()

	go func() {
		<-ctx.Done()
		assert.Nil(t, embedded.Node.Close())
	}()

	// wait until the simulated node has started mining
	isMiningCtx, cancelMiningCtx := context.WithCancel(ctx)
	wait.UntilWithContext(isMiningCtx, func(ctx context.Context) {
		if embedded.ethBackend.IsMining() {
			cancelMiningCtx()
		} else {
			_ = embedded.ethBackend.APIBackend.StartMining(0)
		}
	}, time.Millisecond*50)

	baseClient := embedded.makeClient(t)

	chn, err := chain.NewFromClient(ctx, &legacyClient.Config{ChainID: int(config.ChainID.Int64()), RPCUrl: []string{embedded.Node.HTTPEndpoint()}}, baseClient)

	assert.Nil(t, err)
	chn.SetChainConfig(config)

	embedded.Backend, err = base.NewBaseBackend(ctx, t, chn)
	assert.Nil(t, err)

	return &embedded
}

// Backend is a full geth backend equivalent to running geth --dev.
type Backend struct {
	// Chain is the creates chain object
	*base.Backend
	// Node is the eth node
	*node.Node
	// faucet addr is they key store used for etherbase
	faucetAddr *keystore.Key
	// ethBackend is the eth backend
	ethBackend *eth.Ethereum
}

func (f *Backend) BatchWithContext(ctx context.Context, calls ...w3types.Caller) error {
	return f.BatchContext(ctx, calls...)
}

// ChainConfig gets the chain config for the backend.
func (f *Backend) ChainConfig() *params.ChainConfig {
	return f.ethBackend.BlockChain().Config()
}

// Signer gets the signer for the chain.
func (f *Backend) Signer() types.Signer {
	latestBlock, err := f.BlockByNumber(f.Context(), nil)
	assert.Nil(f.T(), err)

	return types.MakeSigner(f.ChainConfig(), latestBlock.Number())
}

// GethBackendName is the name of the geth backend.
const GethBackendName = "GethBackend"

// BackendName gets the backend name.
func (f *Backend) BackendName() string {
	return GethBackendName
}

// GetTxContext gets a signed transaction from full backend.
func (f *Backend) GetTxContext(ctx context.Context, address *common.Address) (res backends.AuthType) {
	ctx, cancel := onecontext.Merge(ctx, f.Context())
	defer cancel()

	var acct *keystore.Key
	// TODO handle storing accounts to conform to get tx context
	if address != nil {
		acct = f.GetAccount(*address)
		if acct == nil {
			f.T().Errorf("could not get account %s", address.String())
			return res
		}
	} else {
		acct = f.GetFundedAccount(ctx, big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(10)))
		f.Store(acct)
	}

	auth, err := f.NewKeyedTransactorFromKey(acct.PrivateKey)
	assert.Nil(f.T(), err)

	latestBlock, err := f.BlockByNumber(ctx, nil)
	assert.Nil(f.T(), err)

	err = f.Chain.GasSetter().SetGasFee(ctx, auth, latestBlock.NumberU64(), core.CopyBigInt(gasprice.DefaultMaxPrice))
	assert.Nil(f.T(), err)

	auth.GasLimit = ethCore.DeveloperGenesisBlock(0, gasLimit, f.faucetAddr.Address).GasLimit / 2

	return backends.AuthType{
		TransactOpts: auth,
		PrivateKey:   acct.PrivateKey,
	}
}

// wrappedClient wraps the legacyClient in one that contains a chain config.
type wrappedClient struct {
	*ethclient.Client
	rpcClient   *rpc.Client
	w3Client    *w3.Client
	chainConfig *params.ChainConfig
}

// ChainConfig gets the chain config from the wrapped legacyClient.
func (w wrappedClient) ChainConfig() *params.ChainConfig {
	return w.chainConfig
}

// CallContext calls the call context method on the underlying legacyClient.
func (w wrappedClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	//nolint:wrapcheck
	return w.rpcClient.CallContext(ctx, result, method, args...)
}

// BatchCallContext calls the batch call method on the underlying legacyClient.
func (w wrappedClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	//nolint:wrapcheck
	return w.rpcClient.BatchCallContext(ctx, b)
}

func (w wrappedClient) BatchContext(ctx context.Context, calls ...w3types.Caller) error {
	//nolint:wrapcheck
	return w.w3Client.CallCtx(ctx, calls...)
}

// EVMClient gets a legacyClient for the backend.
func (f *Backend) makeClient(tb testing.TB) *wrappedClient {
	tb.Helper()
	handler, err := f.RPCHandler()
	assert.Nil(tb, err)

	rpcClient := rpc.DialInProc(handler)
	rawClient := ethclient.NewClient(rpcClient)
	w3Client := w3.NewClient(rpcClient)

	return &wrappedClient{Client: rawClient, chainConfig: f.ChainConfig(), rpcClient: rpcClient, w3Client: w3Client}
}

// getFaucetTxContext gets a signed transaction from the faucet address.
func (f *Backend) getFaucetTxContext(ctx context.Context) *bind.TransactOpts {
	ctx, cancel := onecontext.Merge(ctx, f.Context())
	defer cancel()

	auth, err := f.NewKeyedTransactorFromKey(f.faucetAddr.PrivateKey)
	assert.Nil(f.T(), err)

	auth.GasPrice, err = f.Client().SuggestGasPrice(ctx)
	assert.Nil(f.T(), err)

	auth.GasLimit = ethCore.DeveloperGenesisBlock(0, gasLimit, f.faucetAddr.Address).GasLimit / 2

	return auth
}

// FaucetSignTx will sign a tx with the faucet addr.
func (f *Backend) FaucetSignTx(tx *types.Transaction) *types.Transaction {
	tx, err := f.Backend.SignTx(tx, f.Signer(), f.faucetAddr.PrivateKey)
	assert.Nil(f.T(), err)
	return tx
}

// FundAccount fundsa  new account.
func (f *Backend) FundAccount(ctx context.Context, address common.Address, amount big.Int) {
	rawTx := f.getFaucetTxContext(ctx)

	tx := f.FaucetSignTx(types.NewTx(&types.LegacyTx{
		To:       &address,
		Value:    &amount,
		Gas:      rawTx.GasLimit,
		GasPrice: rawTx.GasPrice,
	}))

	assert.Nil(f.T(), f.Client().SendTransaction(f.Context(), tx))

	// wait for tx confirmation
	f.WaitForConfirmation(ctx, tx)

	// wait for value of account to be  > 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// attempt to get around insufficient funds for gas * price + value
			newBalance, err := f.Client().BalanceAt(ctx, address, nil)
			if err != nil {
				continue
			}

			if newBalance.Cmp(big.NewInt(0)) != 0 {
				return
			}
		}
	}
}

// GetFundedAccount returns an account with the requested balance. (Note: if genesis acount has an insufficient
// balance, blocks may be mined here).
func (f *Backend) GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key {
	key := f.MockAccount()

	f.Store(key)

	f.FundAccount(ctx, key.Address, *requestBalance)

	return key
}

// toPublic converts the analytics to public apis.
func toPublic(apis []rpc.API) (publicApis []rpc.API) {
	for _, api := range apis {
		api.Public = true
		publicApis = append(publicApis, api)
	}
	return publicApis
}

var _ backends.SimulatedTestBackend = &Backend{}
