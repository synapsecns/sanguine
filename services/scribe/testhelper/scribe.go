package testhelper

import (
	"context"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/ipfs/go-log"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	scribeAPI "github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"testing"
)

var logger = log.Logger("scribe-testhelper")

// NewTestScribe creates a new scribe server with all the test backends passed in.
// all contracts in the registry will be tracked.
func NewTestScribe(ctx context.Context, tb testing.TB, deployedContracts map[uint32][]contracts.DeployedContract, backends ...backends.SimulatedTestBackend) string {
	tb.Helper()

	const db = "sqlite"
	dbPath := filet.TmpDir(tb, "")

	omnirpcURL := testhelper.NewOmnirpcServer(ctx, tb, backends...)

	localmetrics.SetupTestJaeger(ctx, tb)
	metricsProvider, err := metrics.NewByType(ctx, metadata.BuildInfo(), metrics.Jaeger)
	assert.Nil(tb, err)

	eventDB, err := scribeAPI.InitDB(ctx, "sqlite", dbPath, metricsProvider, false)
	assert.Nil(tb, err)

	scribeClients := make(map[uint32][]backfill.ScribeBackend)

	var chainConfigs []config.ChainConfig

	for _, backend := range backends {
		// this backends chain id
		chainID := uint32(backend.GetChainID())

		// create the scribe backend client
		backendClient, err := backfill.DialBackend(ctx, testhelper.GetURL(omnirpcURL, backend), metricsProvider)
		assert.Nil(tb, err)

		// creat ethe scribe client for this chain
		scribeClients[chainID] = []backfill.ScribeBackend{backendClient}

		// loop through all deployed contracts for this chainid adding them to our config
		contractConfigs := getContractConfig(deployedContracts[chainID])

		// add the chain config to the list
		chainConfigs = append(chainConfigs, config.ChainConfig{
			ChainID:               uint32(backend.GetChainID()),
			RequiredConfirmations: 1,
			Contracts:             contractConfigs,
		})
	}

	scribeConfig := config.Config{
		Chains: chainConfigs,
		RPCURL: omnirpcURL,
	}

	scribe, err := node.NewScribe(eventDB, scribeClients, scribeConfig, metricsProvider)
	assert.Nil(tb, err)

	go func() {
		err = scribe.Start(ctx)
		if err != nil {
			logger.Warnf("scribe errored: %v, note this is not necessarily an error with scribe and could indicate the test finished", err)
		}
	}()

	embedded := client.NewEmbeddedScribe(db, dbPath, metricsProvider)
	go func() {
		err = embedded.Start(ctx)
		if err != nil {
			logger.Warnf("embedded scribe errored: %v, note this is not necessarily an error with scribe and could indicate the test finished", err)
		}
	}()

	return fmt.Sprintf("%s:%d", embedded.URL, embedded.Port)
}

func getContractConfig(contracts []contracts.DeployedContract) (contractConfigs config.ContractConfigs) {
	// loop through all deployed contracts for this chainid adding them to our config
	for _, contract := range contracts {
		contractConfigs = append(contractConfigs, config.ContractConfig{
			Address: contract.Address().String(),
			// Note: we could go ahead and get the deploy height
			// from the receipt, but this is more trouble than it's worth
			// considering everything goes through localhost and block numbers are
			// near 0
			StartBlock: 1,
		})
	}

	return contractConfigs
}
