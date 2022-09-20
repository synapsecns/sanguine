package testutil

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/testutil/clickhouse"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/scribe/api"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"go.uber.org/atomic"
	"math/big"
	"net/http"
)

// TestToken is a test token.
type TestToken struct {
	tokenID string
	bridgeconfig.BridgeConfigV3Token
}

var testTokens = []TestToken{{
	tokenID: gofakeit.FirstName(),
	BridgeConfigV3Token: bridgeconfig.BridgeConfigV3Token{
		ChainId:       big.NewInt(int64(gofakeit.Uint32())),
		TokenAddress:  mocks.MockAddress().String(),
		TokenDecimals: gofakeit.Uint8(),
		MaxSwap:       new(big.Int).SetUint64(gofakeit.Uint64()),
		// TODO: this should probably be smaller than maxswap
		MinSwap:       new(big.Int).SetUint64(gofakeit.Uint64()),
		SwapFee:       new(big.Int).SetUint64(gofakeit.Uint64()),
		MaxSwapFee:    new(big.Int).SetUint64(gofakeit.Uint64()),
		MinSwapFee:    new(big.Int).SetUint64(gofakeit.Uint64()),
		HasUnderlying: gofakeit.Bool(),
		IsUnderlying:  gofakeit.Bool(),
	},
},
}

// SetTokenConfig sets the token config.
func (c *TestToken) SetTokenConfig(bridgeConfigContract *bridgeconfig.BridgeConfigRef, opts backends.AuthType) (*ethTypes.Transaction, error) {
	tx, err := bridgeConfigContract.SetTokenConfig(opts.TransactOpts, c.tokenID, c.ChainId, common.HexToAddress(c.TokenAddress),
		c.TokenDecimals, c.MaxSwap, c.MinSwap, c.SwapFee, c.MaxSwapFee, c.MinSwapFee, c.HasUnderlying, c.IsUnderlying)
	if err != nil {
		return nil, fmt.Errorf("could not set token config: %w", err)
	}
	return tx, nil
}

// SetupDB sets up the db.
func SetupDB(t *testsuite.TestSuite) (db db.ConsumerDB, eventDB scribedb.EventDB, gqlClient *client.Client, logIndex atomic.Int64, cleanup func(), testBackend backends.SimulatedTestBackend, deployManager *DeployManager, bridgeConfigContract *bridgeconfig.BridgeConfigRef) {
	dbPath := filet.TmpDir(t.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), dbPath)
	assert.Nil(t.T(), err)

	eventDB = sqliteStore

	logIndex.Store(0)

	freePort := freeport.GetPort()

	go func() {
		assert.Nil(t.T(), api.Start(t.GetSuiteContext(), api.Config{
			HTTPPort: uint16(freePort),
			Database: "sqlite",
			Path:     dbPath,
			GRPCPort: uint16(freeport.GetPort()),
		}))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", freePort)

	gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	// var request *http.Request
	t.Eventually(func() bool {
		request, err := http.NewRequestWithContext(t.GetTestContext(), http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		assert.Nil(t.T(), err)
		res, err := gqlClient.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	})

	cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
	if cleanup == nil || *port == 0 || err != nil {
		return
	}
	assert.Equal(t.T(), err, nil)
	dbURL := "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test?read_timeout=10s&write_timeout=20s"
	consumerDB, err := sql.OpenGormClickhouse(t.GetTestContext(), dbURL)
	assert.Nil(t.T(), err)
	db = consumerDB

	// maybe newSimulatedBackendWithChainID?
	testBackend = simulated.NewSimulatedBackend(t.GetTestContext(), t.T())
	deployManager = NewDeployManager(t.T())

	var deployInfo contracts.DeployedContract
	deployInfo, bridgeConfigContract = deployManager.GetBridgeConfigV3(t.GetTestContext(), testBackend)

	for _, token := range testTokens {
		auth := testBackend.GetTxContext(t.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(bridgeConfigContract, auth)
		assert.Nil(t.T(), err)

		testBackend.WaitForConfirmation(t.GetTestContext(), tx)
	}

	return db, eventDB, gqlClient, logIndex, cleanup, testBackend, deployManager, bridgeConfigContract
}
