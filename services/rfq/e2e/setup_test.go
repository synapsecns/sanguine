package e2e_test

import (
	"fmt"
	"math/big"
	"net/http"
	"slices"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	apiConfig "github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db/sql"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/service"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
)

func (i *IntegrationSuite) setupQuoterAPI() {
	dbPath := filet.TmpDir(i.T(), "")
	apiPort, err := freeport.GetFreePort()
	i.NoError(err)

	apiStore, err := sql.Connect(i.GetTestContext(), dbcommon.Sqlite, dbPath, i.metrics)
	i.NoError(err)

	// make the api without bridges
	apiCfg := apiConfig.Config{
		Database: apiConfig.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  dbPath,
		},
		OmniRPCURL: i.omniServer,
		Bridges: map[uint32]string{
			originBackendChainID: i.manager.Get(i.GetTestContext(), i.originBackend, testutil.FastBridgeType).Address().String(),
			destBackendChainID:   i.manager.Get(i.GetTestContext(), i.destBackend, testutil.FastBridgeType).Address().String(),
		},
		Port: strconv.Itoa(apiPort),
	}
	api, err := rest.NewAPI(i.GetTestContext(), apiCfg, i.metrics, i.omniClient, apiStore)
	i.NoError(err)

	i.apiServer = fmt.Sprintf("http://localhost:%d", apiPort)

	go func() {
		err = api.Run(i.GetTestContext())
		i.NoError(err)
	}()

	// make sure api server hast started
	testsuite.Eventually(i.GetTestContext(), i.T(), func() bool {
		var req *http.Request
		req, err = http.NewRequestWithContext(i.GetTestContext(), http.MethodGet, i.apiServer, nil)
		i.NoError(err)

		//nolint: bodyclose
		_, err = http.DefaultClient.Do(req)
		if err == nil {
			return true
		}
		return false
	})
}

// setupBackends sets up the ether backends and the omnirpc client/server
func (i *IntegrationSuite) setupBackends() {
	var wg sync.WaitGroup

	// Note: we're intentionally not gonna give these guys any tokens to allow the test to do it. What we will do is give them some eth and store the keys.
	var err error
	i.relayerWallet, err = wallet.FromRandom()
	i.NoError(err)

	i.userWallet, err = wallet.FromRandom()
	i.NoError(err)

	// Technically, we can use anvil for origin and geth for destination since only origin needs to use a block
	wg.Add(2)
	go func() {
		defer wg.Done()
		options := anvil.NewAnvilOptionBuilder()
		options.SetChainID(1)
		i.originBackend = anvil.NewAnvilBackend(i.GetTestContext(), i.T(), options)
		i.setupBE(i.originBackend)
	}()
	go func() {
		defer wg.Done()
		i.destBackend = geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(destBackendChainID))
		i.setupBE(i.destBackend)
	}()
	wg.Wait()

	i.omniServer = testhelper.NewOmnirpcServer(i.GetTestContext(), i.T(), i.originBackend, i.destBackend)
	i.omniClient = omnirpcClient.NewOmnirpcClient(i.omniServer, i.metrics, omnirpcClient.WithCaptureReqRes())
}

// setupBe sets up one backend
func (i *IntegrationSuite) setupBE(backend backends.SimulatedTestBackend) {
	// prdeploys are contracts we want to deploy before running the test to speed it up. Obviously, these can be deployed when we need them as well,
	// but this way we can do something while we're waiting for the other backend to startup.
	// no need to wait for these to deploy since they can happen in background as soon as the backend is up.
	predeployTokens := []contracts.ContractType{testutil.DAIType, testutil.USDTType, testutil.USDCType, testutil.WETH9Type}
	predeploys := append(predeployTokens, testutil.FastBridgeType)
	slices.Reverse(predeploys) // return fast bridge first

	ethAmount := *new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10))

	// store the keys
	backend.Store(base.WalletToKey(i.T(), i.relayerWallet))
	backend.Store(base.WalletToKey(i.T(), i.userWallet))

	// fund each of the wallets
	backend.FundAccount(i.GetTestContext(), i.relayerWallet.Address(), ethAmount)
	backend.FundAccount(i.GetTestContext(), i.userWallet.Address(), ethAmount)

	go func() {
		i.manager.BulkDeploy(i.GetTestContext(), core.ToSlice(backend), predeploys...)
	}()

	// TODO: in the case of relayer this not finishing before the test starts can lead to race conditions since
	// nonce may be shared between submitter and relayer. Think about how to deal w/ this.
	for _, user := range []wallet.Wallet{i.relayerWallet, i.userWallet} {
		go func(userWallet wallet.Wallet) {
			for _, token := range predeployTokens {
				i.Approve(backend, i.manager.Get(i.GetTestContext(), backend, token), userWallet)
			}
		}(user)
	}

}

// Approve checks if the token is approved and approves it if not.
func (i *IntegrationSuite) Approve(backend backends.SimulatedTestBackend, token contracts.DeployedContract, user wallet.Wallet) {
	erc20, err := ierc20.NewIERC20(token.Address(), backend)
	i.NoError(err)

	_, fastBridge := i.manager.GetFastBridge(i.GetTestContext(), backend)

	allowance, err := erc20.Allowance(&bind.CallOpts{Context: i.GetTestContext()}, user.Address(), fastBridge.Address())
	i.NoError(err)

	// TODO: can also use in mem cache
	if allowance.Cmp(big.NewInt(0)) == 0 {
		txOpts := backend.GetTxContext(i.GetTestContext(), user.AddressPtr())
		tx, err := erc20.Approve(txOpts.TransactOpts, fastBridge.Address(), core.CopyBigInt(abi.MaxUint256))
		i.NoError(err)
		backend.WaitForConfirmation(i.GetTestContext(), tx)
	}
}

func (i *IntegrationSuite) setupRelayer() {
	// add myself as a filler
	var wg sync.WaitGroup
	wg.Add(2)

	for _, backend := range core.ToSlice(i.originBackend, i.destBackend) {
		go func(backend backends.SimulatedTestBackend) {
			defer wg.Done()

			metadata, rfqContract := i.manager.GetFastBridge(i.GetTestContext(), backend)

			txContext := backend.GetTxContext(i.GetTestContext(), metadata.OwnerPtr())
			tx, err := rfqContract.AddRelayer(txContext.TransactOpts, i.relayerWallet.Address())
			i.NoError(err)

			backend.WaitForConfirmation(i.GetTestContext(), tx)
		}(backend)
	}
	wg.Wait()

	// construct the config
	relayerAPIPort, err := freeport.GetFreePort()
	i.NoError(err)
	dsn := filet.TmpDir(i.T(), "")
	cfg := relconfig.Config{
		// generated ex-post facto
		Chains: map[int]relconfig.ChainConfig{
			originBackendChainID: {
				RFQAddress:    i.manager.Get(i.GetTestContext(), i.originBackend, testutil.FastBridgeType).Address().String(),
				Confirmations: 0,
				Tokens: map[string]relconfig.TokenConfig{
					"ETH": {
						Address:  chain.EthAddress.String(),
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				NativeToken: "ETH",
			},
			destBackendChainID: {
				RFQAddress:    i.manager.Get(i.GetTestContext(), i.destBackend, testutil.FastBridgeType).Address().String(),
				Confirmations: 0,
				Tokens: map[string]relconfig.TokenConfig{
					"ETH": {
						Address:  chain.EthAddress.String(),
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				NativeToken: "ETH",
			},
		},
		OmniRPCURL: i.omniServer,
		// TODO: need to stop hardcoding
		Database: relconfig.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  dsn,
		},
		// generated ex-post facto
		QuotableTokens: map[string][]string{},
		RfqAPIURL:      i.apiServer,
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(i.T(), "", i.relayerWallet.PrivateKeyHex()).Name(),
		},
		RelayerAPIPort: strconv.Itoa(relayerAPIPort),
		BaseChainConfig: relconfig.ChainConfig{
			OriginGasEstimate: 500000,
			DestGasEstimate:   1000000,
		},
		FeePricer: relconfig.FeePricerConfig{
			GasPriceCacheTTLSeconds:   60,
			TokenPriceCacheTTLSeconds: 60,
		},
	}

	// in the first backend, we want to deploy a bunch of different tokens
	// TODO: functionalize me.
	for _, backend := range core.ToSlice(i.originBackend, i.destBackend) {
		tokenTypes := []contracts.ContractType{testutil.DAIType, testutil.USDTType, testutil.USDCType, testutil.WETH9Type}

		for _, tokenType := range tokenTypes {
			tokenAddress := i.manager.Get(i.GetTestContext(), backend, tokenType).Address().String()
			quotableTokenID := fmt.Sprintf("%d-%s", backend.GetChainID(), tokenAddress)

			tokenCaller, err := ierc20.NewIerc20Ref(common.HexToAddress(tokenAddress), backend)
			i.NoError(err)

			decimals, err := tokenCaller.Decimals(&bind.CallOpts{Context: i.GetTestContext()})
			i.NoError(err)

			// first the simple part, add the token to the token map
			cfg.Chains[int(backend.GetChainID())].Tokens[tokenType.Name()] = relconfig.TokenConfig{
				Address:               tokenAddress,
				Decimals:              decimals,
				PriceUSD:              1, // TODO: this will break on non-stables
				RebalanceMethod:       "cctp",
				MaintenanceBalancePct: 20,
				InitialBalancePct:     50,
			}

			compatibleTokens := []contracts.ContractType{tokenType}
			// DAI/USDT are fungible
			if tokenType == testutil.DAIType || tokenType == testutil.USDCType {
				compatibleTokens = []contracts.ContractType{testutil.DAIType, testutil.USDCType}
			}

			// now we need to add the token to the quotable tokens map
			for _, token := range compatibleTokens {
				otherBackend := i.getOtherBackend(backend)
				otherToken := i.manager.Get(i.GetTestContext(), otherBackend, token).Address().String()

				cfg.QuotableTokens[quotableTokenID] = append(cfg.QuotableTokens[quotableTokenID], fmt.Sprintf("%d-%s", otherBackend.GetChainID(), otherToken))
			}
		}
	}

	// Add ETH as quotable token from origin to destination
	cfg.QuotableTokens[fmt.Sprintf("%d-%s", originBackendChainID, chain.EthAddress)] = []string{
		fmt.Sprintf("%d-%s", destBackendChainID, chain.EthAddress),
	}
	cfg.QuotableTokens[fmt.Sprintf("%d-%s", destBackendChainID, chain.EthAddress)] = []string{
		fmt.Sprintf("%d-%s", originBackendChainID, chain.EthAddress),
	}

	// TODO: good chance we wanna leave actually starting this up to the indiividual test.
	i.relayer, err = service.NewRelayer(i.GetTestContext(), i.metrics, cfg)
	i.NoError(err)
	go func() {
		err = i.relayer.Start(i.GetTestContext())
	}()
}
