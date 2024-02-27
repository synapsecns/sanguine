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
	cctpTest "github.com/synapsecns/sanguine/services/cctp-relayer/testutil"
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

	i.setupCCTP()
}

// setupBe sets up one backend
func (i *IntegrationSuite) setupBE(backend backends.SimulatedTestBackend) {
	// prdeploys are contracts we want to deploy before running the test to speed it up. Obviously, these can be deployed when we need them as well,
	// but this way we can do something while we're waiting for the other backend to startup.
	// no need to wait for these to deploy since they can happen in background as soon as the backend is up.
	predeployTokens := []contracts.ContractType{testutil.DAIType, testutil.USDTType, testutil.WETH9Type}
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

func (i *IntegrationSuite) setupCCTP() {
	// deploy the contract to all backends
	testBackends := core.ToSlice(i.originBackend, i.destBackend)
	i.cctpDeployManager.BulkDeploy(i.GetTestContext(), testBackends, cctpTest.SynapseCCTPType, cctpTest.MockMintBurnTokenType)

	// register remote deployments and tokens
	for _, backend := range testBackends {
		cctpContract, cctpHandle := i.cctpDeployManager.GetSynapseCCTP(i.GetTestContext(), backend)
		_, tokenMessengeHandle := i.cctpDeployManager.GetMockTokenMessengerType(i.GetTestContext(), backend)

		// on the above contract, set the remote for each backend
		for _, backendToSetFrom := range core.ToSlice(i.originBackend, i.destBackend) {
			// we don't need to set the backends own remote!
			if backendToSetFrom.GetChainID() == backend.GetChainID() {
				continue
			}

			remoteCCTP, _ := i.cctpDeployManager.GetSynapseCCTP(i.GetTestContext(), backendToSetFrom)
			remoteMessenger, _ := i.cctpDeployManager.GetMockTokenMessengerType(i.GetTestContext(), backendToSetFrom)

			txOpts := backend.GetTxContext(i.GetTestContext(), cctpContract.OwnerPtr())
			// set the remote cctp contract on this cctp contract
			// TODO: verify chainID / domain are correct
			remoteDomain := cctpTest.ChainIDDomainMap[uint32(remoteCCTP.ChainID().Int64())]

			tx, err := cctpHandle.SetRemoteDomainConfig(txOpts.TransactOpts,
				big.NewInt(remoteCCTP.ChainID().Int64()), remoteDomain, remoteCCTP.Address())
			i.Require().NoError(err)
			backend.WaitForConfirmation(i.GetTestContext(), tx)
			fmt.Printf("[cctp] backend: %d, backend to set from: %d, remote domain: %d\n", backend.GetChainID(), backendToSetFrom.GetChainID(), remoteDomain)
			fmt.Printf("[cctp] remote domain config setup tx on chain %d: %v\n", backendToSetFrom.GetChainID(), tx.Hash())

			// register the remote token messenger on the tokenMessenger contract
			_, err = tokenMessengeHandle.SetRemoteTokenMessenger(txOpts.TransactOpts, uint32(backendToSetFrom.GetChainID()), addressToBytes32(remoteMessenger.Address()))
			i.Nil(err)
			fmt.Printf("[cctp] remote token messenger on chain %d: %v\n", backendToSetFrom.GetChainID(), tx.Hash())
		}
	}
}

// addressToBytes32 converts an address to a bytes32.
func addressToBytes32(addr common.Address) [32]byte {
	var buf [32]byte
	copy(buf[:], addr[:])
	return buf
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
	cctpContractOrigin, _ := i.cctpDeployManager.GetSynapseCCTP(i.GetTestContext(), i.originBackend)
	cctpContractDest, _ := i.cctpDeployManager.GetSynapseCCTP(i.GetTestContext(), i.destBackend)
	cfg := relconfig.Config{
		// generated ex-post facto
		Chains: map[int]relconfig.ChainConfig{
			originBackendChainID: {
				RFQAddress:    i.manager.Get(i.GetTestContext(), i.originBackend, testutil.FastBridgeType).Address().String(),
				CCTPAddress:   cctpContractOrigin.Address().Hex(),
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
				CCTPAddress:   cctpContractDest.Address().Hex(),
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
	fmt.Printf("config cctp origin addr: %v\n", cfg.Chains[originBackendChainID].CCTPAddress)
	fmt.Printf("config cctp dest addr: %v\n", cfg.Chains[destBackendChainID].CCTPAddress)

	// in the first backend, we want to deploy a bunch of different tokens
	// TODO: functionalize me.
	for _, backend := range core.ToSlice(i.originBackend, i.destBackend) {
		tokenTypes := []contracts.ContractType{testutil.DAIType, testutil.USDTType, testutil.WETH9Type, cctpTest.MockMintBurnTokenType}

		for _, tokenType := range tokenTypes {
			useCCTP := tokenType == cctpTest.MockMintBurnTokenType
			var tokenAddress string
			if useCCTP {
				tokenAddress = i.cctpDeployManager.Get(i.GetTestContext(), backend, cctpTest.MockMintBurnTokenType).Address().String()
			} else {
				tokenAddress = i.manager.Get(i.GetTestContext(), backend, tokenType).Address().String()
			}
			quotableTokenID := fmt.Sprintf("%d-%s", backend.GetChainID(), tokenAddress)

			tokenCaller, err := ierc20.NewIerc20Ref(common.HexToAddress(tokenAddress), backend)
			i.NoError(err)

			decimals, err := tokenCaller.Decimals(&bind.CallOpts{Context: i.GetTestContext()})
			i.NoError(err)

			rebalanceMethod := ""
			if useCCTP {
				rebalanceMethod = "cctp"
			}

			// first the simple part, add the token to the token map
			cfg.Chains[int(backend.GetChainID())].Tokens[tokenType.Name()] = relconfig.TokenConfig{
				Address:               tokenAddress,
				Decimals:              decimals,
				PriceUSD:              1, // TODO: this will break on non-stables
				RebalanceMethod:       rebalanceMethod,
				MaintenanceBalancePct: 20,
				InitialBalancePct:     50,
			}

			compatibleTokens := []contracts.ContractType{tokenType}
			// DAI/USDC are fungible
			if tokenType == testutil.DAIType || tokenType == cctpTest.MockMintBurnTokenType {
				compatibleTokens = []contracts.ContractType{testutil.DAIType, cctpTest.MockMintBurnTokenType}
			}

			// now we need to add the token to the quotable tokens map
			for _, token := range compatibleTokens {
				otherBackend := i.getOtherBackend(backend)
				var otherToken string
				if token == cctpTest.MockMintBurnTokenType {
					otherToken = i.cctpDeployManager.Get(i.GetTestContext(), otherBackend, cctpTest.MockMintBurnTokenType).Address().String()
				} else {
					otherToken = i.manager.Get(i.GetTestContext(), otherBackend, token).Address().String()
				}

				cfg.QuotableTokens[quotableTokenID] = append(cfg.QuotableTokens[quotableTokenID], fmt.Sprintf("%d-%s", otherBackend.GetChainID(), otherToken))
			}

			// register the token with cctp contract
			cctpContract, cctpHandle := i.cctpDeployManager.GetSynapseCCTP(i.GetTestContext(), backend)
			txOpts := backend.GetTxContext(i.GetTestContext(), cctpContract.OwnerPtr())
			tokenName := fmt.Sprintf("CCTP.%s", tokenType.Name())
			tx, err := cctpHandle.AddToken(txOpts.TransactOpts, tokenName, tokenCaller.Address(), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
			i.Require().NoError(err)
			backend.WaitForConfirmation(i.GetTestContext(), tx)
			fmt.Printf("[cctp] added token %s on chain %d: %v, hash %v\n", tokenName, backend.GetChainID(), tokenCaller.Address().String(), tx.Hash())
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
