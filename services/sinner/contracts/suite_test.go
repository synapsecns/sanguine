package contracts_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"math/big"
	"sync/atomic"
	"testing"

	"github.com/synapsecns/sanguine/services/sinner/db/sql/sqlite"

	"github.com/synapsecns/sanguine/services/sinner/metadata"
)

type ContractsSuite struct {
	*testsuite.TestSuite
	db                 db.EventDB
	logIndex           atomic.Int64
	metrics            metrics.Handler
	originChainID      uint32
	destinationChainID uint32
	originTestLog      types.Log
	desTestLog         types.Log
}

// NewEventContractsSuite creates a new EventContractsSuite.
func NewEventContractsSuite(tb testing.TB) *ContractsSuite {
	tb.Helper()

	return &ContractsSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (t *ContractsSuite) SetupSuite() {
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
	t.Require().Nil(err)

	sqliteStore, err := sqlite.NewSqliteStore(t.GetSuiteContext(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)

	t.db = sqliteStore
	t.originChainID = 1
	t.destinationChainID = 2
	sentTopic := common.HexToHash("0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87")
	executeTopic := common.HexToHash("0x39c48fd1b2185b07007abc7904a8cdf782cfe449fd0e9bba1c2223a691e15f0b")

	// TODO add byte data for each log to test parse
	//originData, err := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000e90000066eee0000000700aa36a70000000f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000007a193a5f45ff4cde43708101b3c03793155a152f0000000000000000000000007a193a5f45ff4cde43708101b3c03793155a152f000000000000000000000000000000000007a120000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	//sentData, err := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000e90000066eee0000000700aa36a70000000f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000007a193a5f45ff4cde43708101b3c03793155a152f0000000000000000000000007a193a5f45ff4cde43708101b3c03793155a152f000000000000000000000000000000000007a120000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	Nil(t.T(), err)
	t.originTestLog = types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		BlockNumber: gofakeit.Uint64(),
		Topics:      []common.Hash{sentTopic},
		Data:        []byte{},
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		TxIndex:     uint(gofakeit.Int8()),
		BlockHash:   common.HexToHash(big.NewInt(gofakeit.Int64()).String()),
		Index:       uint(gofakeit.Int8()),
		Removed:     false,
	}
	t.desTestLog = types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		BlockNumber: gofakeit.Uint64(),
		Topics:      []common.Hash{executeTopic},
		Data:        []byte{},
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		TxIndex:     uint(gofakeit.Int8()),
		BlockHash:   common.HexToHash(big.NewInt(gofakeit.Int64()).String()),
		Index:       uint(gofakeit.Int8()),
		Removed:     false,
	}
	t.Require().Nil(err)
}

// TestContractsSuite tests the db suite.
func TestEventContractsSuite(t *testing.T) {
	suite.Run(t, NewEventContractsSuite(t))
}
