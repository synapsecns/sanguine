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
	db                 db.TestEventDB
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
	Nil(t.T(), err)

	sqliteStore, err := sqlite.NewSqliteStore(t.GetSuiteContext(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)

	t.db = sqliteStore
	t.originChainID = 1
	t.destinationChainID = 2
	sentTopic := common.HexToHash("0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87")
	sentTopic2 := common.HexToHash("0xc6e19a3538fbd9b7a4f9bd8d45e08a95ff23e7e03b6a3bc9d9db9b8869b55c94")
	sentTopic3 := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")
	sentTopic4 := common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000002c")

	executeTopic := common.HexToHash("0x39c48fd1b2185b07007abc7904a8cdf782cfe449fd0e9bba1c2223a691e15f0b")
	executeTopic2 := common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000002b")
	executeTopic3 := common.HexToHash("0x481244fb9db711b88ab9bfe081311cbed0b50dc547a71151aef55a38871fc9bd")

	t.originTestLog = types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		BlockNumber: gofakeit.Uint64(),
		Topics:      []common.Hash{sentTopic, sentTopic2, sentTopic3, sentTopic4},
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
		Topics:      []common.Hash{executeTopic, executeTopic2, executeTopic3},
		Data:        []byte{},
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		TxIndex:     uint(gofakeit.Int8()),
		BlockHash:   common.HexToHash(big.NewInt(gofakeit.Int64()).String()),
		Index:       uint(gofakeit.Int8()),
		Removed:     false,
	}
}

// TestContractsSuite tests the db suite.
func TestEventContractsSuite(t *testing.T) {
	suite.Run(t, NewEventContractsSuite(t))
}
