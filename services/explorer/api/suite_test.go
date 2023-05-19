package api_test

import (
	gosql "database/sql"
	"fmt"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/services/explorer/api"
	explorerclient "github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	gqlServer "github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"math/big"
	"net/http"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/graphql/client"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server"
	"go.uber.org/atomic"
)

type MvBridgeEvent struct {
	InsertTime uint64 `gorm:"column:insert_time"`
	// FInsertTime is the time the event was inserted into the database.
	FInsertTime uint64 `gorm:"column:finsert_time"`
	// FContractAddress is the address of the contract that generated the event.
	FContractAddress string `gorm:"column:fcontract_address"`
	// FChainID is the chain id of the contract that generated the event.
	FChainID uint32 `gorm:"column:fchain_id"`
	// FEventType is the type of the event.
	FEventType uint8 `gorm:"column:fevent_type"`
	// FBlockNumber is the block number of the event.
	FBlockNumber uint64 `gorm:"column:fblock_number"`
	// FTxHash is the transaction hash of the event.
	FTxHash string `gorm:"column:ftx_hash"`
	// FToken is the address of the token.
	FToken string `gorm:"column:ftoken"`
	// FAmount is the amount of tokens.
	FAmount *big.Int `gorm:"column:famount;type:UInt256"`
	// FEventIndex is the index of the log.
	FEventIndex uint64 `gorm:"column:fevent_index"`
	// FDestinationKappa is the destination kappa.
	FDestinationKappa string `gorm:"column:fdestination_kappa"`
	// FSender is the address of the sender.
	FSender string `gorm:"column:fsender"`

	// FRecipient is the address to send the tokens to.
	FRecipient gosql.NullString `gorm:"column:frecipient"`
	// FRecipientBytes is the recipient address in bytes.
	FRecipientBytes gosql.NullString `gorm:"column:frecipient_bytes"`
	// FDestinationChainID is the chain id of the chain to send the tokens to.
	FDestinationChainID *big.Int `gorm:"column:fdestination_chain_id;type:UInt256"`
	// FFee is the fee.
	FFee *big.Int `gorm:"column:ffee;type:UInt256"`
	// FKappa is theFee keccak256 hash of the transaction.
	FKappa gosql.NullString `gorm:"column:fkappa"`
	// FTokenIndexFrom is the index of the from token in the pool.
	FTokenIndexFrom *big.Int `gorm:"column:ftoken_index_from;type:UInt256"`
	// FTokenIndexTo is the index of the to token in the pool.
	FTokenIndexTo *big.Int `gorm:"column:ftoken_index_to;type:UInt256"`
	// FMinDy is the minimum amount of tokens to receive.
	FMinDy *big.Int `gorm:"column:fmin_dy;type:UInt256"`
	// FDeadline is the deadline of the transaction.
	FDeadline *big.Int `gorm:"column:fdeadline;type:UInt256"`
	// FSwapSuccess is whether the swap was successful.
	FSwapSuccess *big.Int `gorm:"column:fswap_success;type:UInt256"`
	// FSwapTokenIndex is the index of the token in the pool.
	FSwapTokenIndex *big.Int `gorm:"column:fswap_token_index;type:UInt256"`
	// FSwapMinAmount is the minimum amount of tokens to receive.
	FSwapMinAmount *big.Int `gorm:"column:fswap_min_amount;type:UInt256"`
	// FSwapDeadline is the deadline of the swap transaction.
	FSwapDeadline *big.Int `gorm:"column:fswap_deadline;type:UInt256"`
	// FTokenID is the token's ID.
	FTokenID gosql.NullString `gorm:"column:ftoken_id"`
	// FAmountUSD is the amount in USD.
	FAmountUSD *float64 `gorm:"column:famount_usd;type:Float64"`
	// FFeeAmountUSD is the fee amount in USD.
	FFeeAmountUSD *float64 `gorm:"column:ffee_amount_usd;type:Float64"`
	// FTokenDecimal is the token's decimal.
	FTokenDecimal *uint8 `gorm:"column:ftoken_decimal"`
	// FTokenSymbol is the token's symbol from coin gecko.
	FTokenSymbol gosql.NullString `gorm:"column:ftoken_symbol"`
	// FTimeStamp is the timestamp of the block in which the event occurred.
	FTimeStamp *uint64 `gorm:"column:ftimestamp"`
	// TInsertTime is the time the event was inserted into the database.
	TInsertTime uint64 `gorm:"column:finsert_time"`
	// TContractAddress is the address of the contract that generated the event.
	TContractAddress string `gorm:"column:tcontract_address"`
	// TChainID is the chain id of the contract that generated the event.
	TChainID uint32 `gorm:"column:tchain_id"`
	// TEventType is the type of the event.
	TEventType uint8 `gorm:"column:tevent_type"`
	// TBlockNumber is the block number of the event.
	TBlockNumber uint64 `gorm:"column:tblock_number"`
	// TTxHash is the transaction hash of the event.
	TTxHash string `gorm:"column:ttx_hash"`
	// TToken is the address of the token.
	TToken string `gorm:"column:ttoken"`
	// TAmount is the amount of tokens.
	TAmount *big.Int `gorm:"column:tamount;type:UInt256"`
	// TEventIndex is the index of the log.
	TEventIndex uint64 `gorm:"column:tevent_index"`
	// TDestinationKappa is the destination kappa.
	TDestinationKappa string `gorm:"column:tdestination_kappa"`
	// TSender is the address of the sender.
	TSender string `gorm:"column:tsender"`

	// TRecipient is the address to send the tokens to.
	TRecipient gosql.NullString `gorm:"column:trecipient"`
	// TRecipientBytes is the recipient address in bytes.
	TRecipientBytes gosql.NullString `gorm:"column:trecipient_bytes"`
	// TDestinationChainID is the chain id of the chain to send the tokens to.
	TDestinationChainID *big.Int `gorm:"column:tdestination_chain_id;type:UInt256"`
	// TFee is the fee.
	TFee *big.Int `gorm:"column:tfee;type:UInt256"`
	// TKappa is theFee keccak256 hash of the transaction.
	TKappa gosql.NullString `gorm:"column:tkappa"`
	// TTokenIndexFrom is the index of the from token in the pool.
	TTokenIndexFrom *big.Int `gorm:"column:ttoken_index_from;type:UInt256"`
	// TTokenIndexTo is the index of the to token in the pool.
	TTokenIndexTo *big.Int `gorm:"column:ttoken_index_to;type:UInt256"`
	// TMinDy is the minimum amount of tokens to receive.
	TMinDy *big.Int `gorm:"column:tmin_dy;type:UInt256"`
	// TDeadline is the deadline of the transaction.
	TDeadline *big.Int `gorm:"column:tdeadline;type:UInt256"`
	// TSwapSuccess is whether the swap was successful.
	TSwapSuccess *big.Int `gorm:"column:tswap_success;type:UInt256"`
	// TSwapTokenIndex is the index of the token in the pool.
	TSwapTokenIndex *big.Int `gorm:"column:tswap_token_index;type:UInt256"`
	// TSwapMinAmount is the minimum amount of tokens to receive.
	TSwapMinAmount *big.Int `gorm:"column:tswap_min_amount;type:UInt256"`
	// TSwapDeadline is the deadline of the swap transaction.
	TSwapDeadline *big.Int `gorm:"column:tswap_deadline;type:UInt256"`
	// TTokenID is the token's ID.
	TTokenID gosql.NullString `gorm:"column:ttoken_id"`
	// TAmountUSD is the amount in USD.
	TAmountUSD *float64 `gorm:"column:tamount_usd;type:Float64"`
	// TFeeAmountUSD is the fee amount in USD.
	TFeeAmountUSD *float64 `gorm:"column:tfee_amount_usd;type:Float64"`
	// TTokenDecimal is the token's decimal.
	TTokenDecimal *uint8 `gorm:"column:ttoken_decimal"`
	// TTokenSymbol is the token's symbol from coin gecko.
	TTokenSymbol gosql.NullString `gorm:"column:ttoken_symbol"`
	// TTimeStamp is the timestamp of the block in which the event occurred.
	TTimeStamp *uint64 `gorm:"column:ttimestamp"`
}

// APISuite defines the basic test suite.
type APISuite struct {
	*testsuite.TestSuite
	db     db.ConsumerDB
	client *client.Client
	// grpcClient *rest.APIClient
	eventDB       scribedb.EventDB
	gqlClient     *explorerclient.Client
	logIndex      atomic.Int64
	cleanup       func()
	testBackend   backends.SimulatedTestBackend
	deployManager *testutil.DeployManager
	chainIDs      []uint32
	scribeMetrics metrics.Handler
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *APISuite {
	tb.Helper()
	return &APISuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (g *APISuite) SetupSuite() {
	g.TestSuite.SetupSuite()
	localmetrics.SetupTestJaeger(g.GetSuiteContext(), g.T())

	var err error
	g.scribeMetrics, err = metrics.NewByType(g.GetSuiteContext(), metadata.BuildInfo(), metrics.Jaeger)
	g.Require().Nil(err)
}

func (g *APISuite) SetupTest() {
	g.TestSuite.SetupTest()

	g.db, g.eventDB, g.gqlClient, g.logIndex, g.cleanup, g.testBackend, g.deployManager = testutil.NewTestEnvDB(g.GetTestContext(), g.T(), g.scribeMetrics)

	httpport := freeport.GetPort()
	cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
	NotNil(g.T(), cleanup)
	NotNil(g.T(), port)
	Nil(g.T(), err)
	if port == nil || err != nil {
		g.TearDownTest()
		return
	}

	address := "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test"
	g.db, err = sql.OpenGormClickhouse(g.GetTestContext(), address, false)
	Nil(g.T(), err)
	err = g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Set("gorm:table_options", "ENGINE=ReplacingMergeTree(finsert_time) ORDER BY (fevent_index, fblock_number, fevent_type, ftx_hash, fchain_id, fcontract_address)").AutoMigrate(&MvBridgeEvent{})
	Nil(g.T(), err)

	g.chainIDs = []uint32{1, 10, 25, 56, 137}
	go func() {
		Nil(g.T(), api.Start(g.GetSuiteContext(), api.Config{
			HTTPPort:  uint16(httpport),
			Address:   address,
			ScribeURL: g.gqlClient.Client.BaseURL,
		}))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", httpport)

	g.client = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, gqlServer.GraphqlEndpoint))

	g.Eventually(func() bool {
		request, err := http.NewRequestWithContext(g.GetTestContext(), http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		Nil(g.T(), err)
		res, err := g.client.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	})
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
