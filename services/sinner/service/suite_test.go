package service_test

import (
	"context"
	"database/sql"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	"fmt"

	"github.com/alecthomas/assert"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"k8s.io/apimachinery/pkg/util/wait"

	"math/big"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/scribe/api"
	scribeSqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/sql/mysql"
	"github.com/synapsecns/sanguine/services/sinner/db/sql/sqlite"
	"github.com/synapsecns/sanguine/services/sinner/fetcher"
	"github.com/synapsecns/sanguine/services/sinner/fetcher/client"
	gqlServer "github.com/synapsecns/sanguine/services/sinner/graphql/server"
	"github.com/synapsecns/sanguine/services/sinner/metadata"
	"gorm.io/gorm/schema"
)

// ServiceSuite is the test suite for the db package.
type ServiceSuite struct {
	*testsuite.TestSuite
	dbs                    []db.TestEventDB
	logIndex               atomic.Int64
	scribeDB               scribedb.EventDB
	scribeDBPath           string
	scribeFetcherPath      string
	metrics                metrics.Handler
	scribeFetcher          fetcher.ScribeFetcher
	testBackend            backends.SimulatedTestBackend
	destinationTestBackend backends.SimulatedTestBackend
	originChainID          uint32
	destinationChainID     uint32
	originTestLog          types.Log
	destinationTestLog     types.Log
	originTestTx           *types.Transaction
	destinationTestTx      *types.Transaction
}

// NewEventServiceSuite creates a new EventServiceSuite.
func NewEventServiceSuite(tb testing.TB) *ServiceSuite {
	tb.Helper()

	return &ServiceSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.TestEventDB{},
	}
}

func (t *ServiceSuite) SetupTest() {
	t.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)

	t.dbs = []db.TestEventDB{sqliteStore}
	t.setupMysqlDB()
}

func (t *ServiceSuite) SetupSuite() {
	t.TestSuite.SetupSuite()
	t.logIndex.Store(0)

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(t.GetSuiteContext(), t.T())
		metricsHandler = metrics.Jaeger
	}
	var err error
	t.metrics, err = metrics.NewByType(t.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(t.T(), err)

	t.scribeDBPath = filet.TmpDir(t.T(), "")
	sqliteStore, err := sqlite.NewSqliteStore(t.GetSuiteContext(), t.scribeDBPath, t.metrics, false)
	Nil(t.T(), err)

	t.dbs = []db.TestEventDB{sqliteStore}
	t.originChainID = 421614
	t.destinationChainID = 444
	t.scribeDB, t.scribeFetcher = t.CreateScribeFetcher(t.GetSuiteContext())
	t.testBackend = simulated.NewSimulatedBackendWithChainID(t.GetSuiteContext(), t.T(), big.NewInt(int64(t.originChainID)))
	t.destinationTestBackend = simulated.NewSimulatedBackendWithChainID(t.GetSuiteContext(), t.T(), big.NewInt(int64(t.destinationChainID)))
	t.BuildAndSetTestData()
}

func (t *ServiceSuite) setupMysqlDB() {
	// skip if mysql test disabled, this really only needs to be run in ci
	// skip if mysql test disabled
	if os.Getenv(dbcommon.EnableMysqlTestVar) == "" {
		return
	}
	// sets up the conn string to the default database
	connString := dbcommon.GetTestConnString()
	// sets up the myqsl db
	testDB, err := sql.Open("mysql", connString)
	Nil(t.T(), err)
	// close the db once the connection is odne
	defer func() {
		Nil(t.T(), testDB.Close())
	}()

	// override the naming strategy to prevent tests from messing with each other.
	// todo this should be solved via a proper teardown process or transactions.
	mysql.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", t.GetTestID(), time.Now().Unix()),
	}

	mysql.MaxIdleConns = 10
	mysql.MaxOpenConns = 10

	// create the sql store
	mysqlStore, err := mysql.NewMysqlStore(t.GetTestContext(), connString, t.metrics, false)
	Nil(t.T(), err)
	// add the db
	t.dbs = append(t.dbs, mysqlStore)
}

func (t *ServiceSuite) RunOnAllDBs(testFunc func(testDB db.TestEventDB)) {
	t.T().Helper()
	wg := sync.WaitGroup{}
	for _, testDB := range t.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.TestEventDB) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

func (t *ServiceSuite) CreateScribeFetcher(ctx context.Context) (scribeDB scribedb.EventDB, scribeFetcher fetcher.ScribeFetcher) {
	t.T().Helper()
	dbPath := filet.TmpDir(t.T(), "")
	sqliteStore, err := scribeSqlite.NewSqliteStore(ctx, dbPath, t.metrics, false)
	Nil(t.T(), err)

	scribeDB = sqliteStore

	freePort := freeport.GetPort()

	go func() {
		err = api.Start(ctx, api.Config{
			Port:     uint16(freePort),
			Database: "sqlite",
			Path:     dbPath,
		}, t.metrics)
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", freePort)
	gqlClient := client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, gqlServer.GraphqlEndpoint))
	t.scribeFetcherPath = fmt.Sprintf("%s%s", baseURL, gqlServer.GraphqlEndpoint)
	checkConnection := func() bool {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		assert.Nil(t.T(), err)
		res, err := gqlClient.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	}
	cancellableCtx, cancel := context.WithCancel(ctx)
	isTrue := false
	wait.UntilWithContext(cancellableCtx, func(cancellableCtx context.Context) {
		if checkConnection() {
			isTrue = true
			cancel()
		}
	}, time.Millisecond)

	// make sure the context didn't cancel
	if !isTrue {
		t.T().Errorf("expected %T to be true before test context timed out", checkConnection)
	}

	scribeFetcher = fetcher.NewFetcher(gqlClient, t.metrics)

	return scribeDB, scribeFetcher
}

func (t *ServiceSuite) BuildAndSetTestData() {
	// https://sepolia-explorer.arbitrum.io/tx/0x5763ef1207b1e043a94cd0cc74e98b743c7be11e826143c64eed3606c84d1222
	sentTopic := common.HexToHash("0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87")
	sentTopic2 := common.HexToHash("0xa486a31a83939a86bab5088e333624ec6cd232a257a14295086f6c319917955b")
	sentTopic3 := common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000f0")
	sentTopic4 := common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000001bc")

	// https://explorerl2-synapse-sepolia-testnet-1mdqkm651f.t.conduit.xyz/tx/0x09680d9dd6585a7608470ed29acd0a845a5498d93136ad2ea6a8f271e55de577/logs
	executeTopic := common.HexToHash("0x39c48fd1b2185b07007abc7904a8cdf782cfe449fd0e9bba1c2223a691e15f0b")
	executeTopic2 := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000066eee")
	executeTopic3 := common.HexToHash("0xa486a31a83939a86bab5088e333624ec6cd232a257a14295086f6c319917955b")

	t.originTestLog = types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		BlockNumber: 625780,
		Topics:      []common.Hash{sentTopic, sentTopic2, sentTopic3, sentTopic4},
		Data:        []byte{},
		TxHash:      common.HexToHash("0x5763ef1207b1e043a94cd0cc74e98b743c7be11e826143c64eed3606c84d1222"),
		TxIndex:     1,
		BlockHash:   common.HexToHash(big.NewInt(625780).String()),
		Index:       202,
		Removed:     false,
	}
	t.destinationTestLog = types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		BlockNumber: 1975778,
		Topics:      []common.Hash{executeTopic, executeTopic2, executeTopic3},
		Data:        common.FromHex("0000000000000000000000000000000000000000000000000000000000000001"),
		TxHash:      common.HexToHash("0x09680d9dd6585a7608470ed29acd0a845a5498d93136ad2ea6a8f271e55de577"),
		TxIndex:     1,
		BlockHash:   common.HexToHash(big.NewInt(1975778).String()),
		Index:       40,
		Removed:     false,
	}

	fromAddress := common.HexToAddress("0x7A193a5f45ff4cDE43708101b3C03793155A152F")
	toAddress := common.HexToAddress("0xA944636Ac279e0346AF96Ef7e236025C6cBFE609")

	t.originTestTx = types.NewTx(&types.LegacyTx{
		Nonce:    39,
		GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
		Gas:      380833,
		To:       &fromAddress,
		Value:    big.NewInt(0),
		Data:     common.Hex2Bytes("0xaa402039000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000001bc0000000000000000000000007a193a5f45ff4cde43708101b3c03793155a152f0000000000000000000000000000000000000000000000000000000000000000"),
	})
	t.destinationTestTx = types.NewTx(&types.LegacyTx{
		Nonce:    110,
		GasPrice: new(big.Int).SetUint64(gofakeit.Uint64()),
		Gas:      229658,
		To:       &toAddress,
		Value:    big.NewInt(0),
		Data:     common.Hex2Bytes("0x32ff14d200000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000005e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000000e90000066eee000000f0000001bc0000003b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000007a193a5f45ff4cde43708101b3c03793155a152f0000000000000000000000007a193a5f45ff4cde43708101b3c03793155a152f000000000000000000000000000000000007a1200000000000000000000000000000000000000000000000000000000000000000000000c10000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020da872e557575ae62cf159393b3a342fc413d76227f01e1ccdf6af4157abb828cfee5bf32581a72e90f2b825c3a3b8fad35d2fd63505a31bfda397c9dbcfb76278825cb53695cf8f65db11a70429f0ae2a5ca5ce14ebc8db5ae3d01a55372aead451e5216d110e7d0d4f4bd7e9722fda6764ef838a88e199ee7dfb965d215e70aee5d577d26f867ec2a8166d1d5c723fbf65c668ddf79f100f4ce2d79123a9f9ff49481f1cad21fba3effd258b8968dcd7cea6aa2df9dfb56ffa45267f8cf0ce1c8ff49c84eb9e9f0c0d4d77b3e98ce26cf49ccb7cbbf5ecd53683e4a0f999c36f07c4f9ef51544aa70b6fa1d55bb2cfeea01b413286e89e285810b529cfbc1db1c73fa2ed74e0f884be2f6d4e24a1b5f84f90fcaaf9f081c84a84dab098eba5a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006201ecd046812de623dc2ed30164eb340f08018623405ea2da9d5a3131df9bb2100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	})
}

// TestServiceSuite tests the db suite.
func TestEventServiceSuite(t *testing.T) {
	suite.Run(t, NewEventServiceSuite(t))
}
