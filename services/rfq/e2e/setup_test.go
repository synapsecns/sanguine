package e2e_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	apiConfig "github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db/sql"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
	"github.com/synapsecns/sanguine/services/rfq/relayer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"math/big"
	"net/http"
	"strconv"
	"sync"
)

func (i *IntegrationSuite) setupAPI() {
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

	// prdeploys are contracts we want to deploy before running the test to speed it up. Obviously, these can be deployed when we need them as well,
	// but this way we can do something while we're waiting for the other backend to startup.
	// no need to wait for these to deploy since they can happen in background as soon as the backend is up.
	predeploys := []contracts.ContractType{testutil.FastBridgeType, testutil.DAIType, testutil.USDTType, testutil.USDCType, testutil.WETH9Type}

	wg.Add(2)
	go func() {
		defer wg.Done()
		i.originBackend = geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(originBackendChainID))
		go func() {
			i.manager.BulkDeploy(i.GetTestContext(), core.ToSlice(i.originBackend), predeploys...)
		}()

	}()
	go func() {
		defer wg.Done()
		i.destBackend = geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(destBackendChainID))
		go func() {
			i.manager.BulkDeploy(i.GetTestContext(), core.ToSlice(i.destBackend), predeploys...)
		}()
	}()
	wg.Wait()

	i.omniServer = testhelper.NewOmnirpcServer(i.GetTestContext(), i.T(), i.originBackend, i.destBackend)
	i.omniClient = omnirpcClient.NewOmnirpcClient(i.omniServer, i.metrics, omnirpcClient.WithCaptureReqRes())
}

func (i *IntegrationSuite) setupRelayer() {
	var err error
	i.relayerWallet, err = wallet.FromRandom()
	i.NoError(err)

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
	dsn := filet.TmpDir(i.T(), "")
	cfg := relconfig.Config{
		// generated ex-post facto
		Tokens: map[int][]string{},
		Bridges: map[int]relconfig.ChainConfig{
			originBackendChainID: {
				Bridge:        i.manager.Get(i.GetTestContext(), i.originBackend, testutil.FastBridgeType).Address().String(),
				Confirmations: 0,
			},
			destBackendChainID: {
				Bridge:        i.manager.Get(i.GetTestContext(), i.destBackend, testutil.FastBridgeType).Address().String(),
				Confirmations: 0,
			},
		},
		OmnirpcURL: i.omniServer,
		// TODO: need to stop hardcoding
		DBConfig: dsn,
		// generated ex-post facto
		QuotableTokens: map[string][]string{},
		RelayerAddress: i.relayerWallet.Address(),
		RfqAPIURL:      i.apiServer,
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(i.T(), "", i.relayerWallet.PrivateKeyHex()).Name(),
		},
	}

	// in the first backend, we want to deploy a bunch of different tokens
	// TODO: functionalize me.
	for _, backend := range core.ToSlice(i.originBackend, i.destBackend) {
		tokenTypes := []contracts.ContractType{testutil.DAIType, testutil.USDTType, testutil.USDCType, testutil.WETH9Type}

		for _, tokenType := range tokenTypes {
			tokenAddress := i.manager.Get(i.GetTestContext(), backend, tokenType).Address().String()
			quotableTokenID := fmt.Sprintf("%d-%s", backend.GetChainID(), tokenAddress)

			// first the simple part, add the token to the token map
			cfg.Tokens[int(backend.GetChainID())] = append(cfg.Tokens[int(backend.GetChainID())], tokenAddress)

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

	// TODO: good chance we wanna leave actually starting this up to the indiividual test.
	i.relayer, err = relayer.NewRelayer(i.GetTestContext(), i.metrics, cfg)
	i.NoError(err)
	go func() {
		err = i.relayer.Start(i.GetTestContext())
	}()
}
