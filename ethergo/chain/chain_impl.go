package chain

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/chain/gas"
	"github.com/synapsecns/sanguine/ethergo/chain/watcher"
)

var _ Chain = &baseChain{}

// baseChain is an extensible struct that implements chain. This is used for interacting with an underlying
// evm rpc.
type baseChain struct {
	// ctx: this is used on all non-client methods
	//nolint: containedctx
	ctx context.Context
	// EVMClient is the rpc client used to access the chain
	client.MeteredEVMClient
	// cfg is the configuration for the chain
	cfg *client.Config
	// structMux is the mutex
	structMux sync.RWMutex
	// chainConfig is the chain config
	chainConfig *params.ChainConfig
	// height watcher keeps track of the current chain tip height
	heightWatcher chainwatcher.BlockHeightWatcher
	// contractWatcher watches for events on the contract
	contractWatcher chainwatcher.ContractWatcher
	// chainID is the cached chain id
	chainID uint
	// cancel is used to close the session
	cancel context.CancelFunc
	// rpcAddress used by the chain
	rpcAddress string
	// estimator stores the gas estimator
	// Deprecated: use setter
	estimator gas.PriceEstimator
	// setter is used to set gas on a tx
	setter gas.Setter
}

func (b *baseChain) GasSetter() gas.Setter {
	return b.setter
}

// Name is the name of the chain.
// Deprecated: do not use.
func (b *baseChain) Name() string {
	return ""
}

// Config gets the config for the chain.
func (b *baseChain) Config() *client.Config {
	return b.cfg
}

// RPCAddress gets the rpc address the chain is currently connected to.
func (b *baseChain) RPCAddress() string {
	return b.rpcAddress
}

// ChainConfig gets the chain config if it's available.
func (b *baseChain) ChainConfig() *params.ChainConfig {
	b.structMux.RLock()
	defer b.structMux.RUnlock()

	return b.chainConfig
}

// Estimator gets the gas estimator.
func (b *baseChain) Estimator() gas.PriceEstimator {
	return b.estimator
}

// GetHeightWatcher gets a block height watcher.
func (b *baseChain) GetHeightWatcher() chainwatcher.BlockHeightWatcher {
	return b.heightWatcher
}

// SetChainConfig sets the chain config manually. This is used mostly for testing.
func (b *baseChain) SetChainConfig(config *params.ChainConfig) {
	b.structMux.Lock()
	defer b.structMux.Unlock()
	b.chainConfig = config
}

// NewFromMeteredClient creaates a new client from a metered evm client.
func NewFromMeteredClient(ctx context.Context, config *client.Config, meteredClient client.MeteredEVMClient) (chain Chain, err error) {
	b := baseChain{}
	b.ctx, b.cancel = context.WithCancel(ctx)
	b.cfg = config

	if len(config.RPCUrl) > 0 {
		b.rpcAddress = config.RPCUrl[0]
	}

	logger.Debugf("creating eth client %s", config.RPCUrl)

	// TODO, correctly set chain id here, only affects metrics
	b.MeteredEVMClient = meteredClient
	b.chainID = uint(config.ChainID)
	b.chainConfig = client.ConfigFromID(b.GetBigChainID())
	// initialize monitor/estimator

	b.heightWatcher = watcher.NewBlockHeightWatcher(ctx, uint64(b.chainID), b.MeteredEVMClient)
	b.contractWatcher = watcher.NewContractWatcher(ctx, &b, b.heightWatcher, b.cfg.RequiredConfirmations)
	b.estimator = gas.NewGasPriceEstimator(ctx, &b)
	b.setter = gas.NewGasSetter(ctx, &b)

	return &b, nil
}

// NewFromURL creates a new client from a url.
//
//nolint:ireturn
func NewFromURL(ctx context.Context, url string) (Chain, error) {
	tmpClient, err := client.NewClient(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("could not create new client at %s: %w", url, err)
	}

	chainID, err := tmpClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get chain id: %w", err)
	}

	meteredClient := client.NewMeteredClient(tmpClient, chainID, url, nil)
	return NewFromMeteredClient(ctx, &client.Config{
		RPCUrl:  []string{url},
		ChainID: int(chainID.Int64()),
	}, meteredClient)
}

// NewFromClient gets a chain from client.
func NewFromClient(ctx context.Context, config *client.Config, evmClient client.EVMClient) (chain Chain, err error) {
	rpcURL := ""
	if len(config.RPCUrl) > 0 {
		rpcURL = config.RPCUrl[0]
	}
	meteredClient := client.NewMeteredClient(evmClient, big.NewInt(0), rpcURL, config.LimiterConfig)
	return NewFromMeteredClient(ctx, config, meteredClient)
}

// New creates a new rpc client for querying an evm-based rpc server and attempts to connect to the chain.
func New(ctx context.Context, config *client.Config) (evmClient Chain, err error) {
	if len(config.RPCUrl) == 0 {
		return nil, errors.New("could not get ws url")
	}

	// RPCUrl gets the rpc url from the client
	baseClient, err := client.NewClientFromChainID(ctx, config.RPCUrl[0], big.NewInt(int64(config.ChainID)))
	if err != nil {
		return nil, fmt.Errorf("could not create pool client: %w", err)
	}

	return NewFromClient(ctx, config, baseClient)
}

// ChainName gets the chain name from the config.
func (b *baseChain) ChainName() string {
	return b.Name()
}

// GetChainID gets the cached chain id.
func (b *baseChain) GetChainID() uint {
	return b.chainID
}

// GetBigChainID gets the chain id as a big int.
func (b *baseChain) GetBigChainID() *big.Int {
	return big.NewInt(int64(b.chainID))
}

// ListenOnContract registers an event listener on a contract address
// the events emitted by the a contractAddress after confirmations (defined in the config).
func (b *baseChain) ListenOnContract(ctx context.Context, contractAddress string, eventLog chan interface{}) error {
	err := b.contractWatcher.ListenOnContract(ctx, contractAddress, eventLog)
	if err != nil {
		return fmt.Errorf("could not listen on contract: %w", err)
	}

	return nil
}
