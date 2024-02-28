package listener_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ipfs/go-log"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/ethergo/chain/listener/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/chain/listener"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"gorm.io/driver/sqlite"
)

const chainID = 10

type ListenerTestSuite struct {
	*testsuite.TestSuite
	manager            *testutil.DeployManager
	backend            backends.SimulatedTestBackend
	store              db.ChainListenerDB
	metrics            metrics.Handler
	fastBridge         *fastbridge.FastBridgeRef
	fastBridgeMetadata contracts.DeployedContract
}

func NewListenerSuite(tb testing.TB) *ListenerTestSuite {
	return &ListenerTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestListenerSuite(t *testing.T) {
	suite.Run(t, NewListenerSuite(t))
}

func (l *ListenerTestSuite) SetupTest() {
	l.TestSuite.SetupTest()

	l.manager = testutil.NewDeployManager(l.T())
	l.backend = geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(chainID))
	var err error
	l.metrics = metrics.NewNullHandler()
	l.store, err = NewSqliteStore(l.GetTestContext(), filet.TmpDir(l.T(), ""), l.metrics)
	l.Require().NoError(err)

	l.fastBridgeMetadata, l.fastBridge = l.manager.GetFastBridge(l.GetTestContext(), l.backend)
}

func (l *ListenerTestSuite) TestGetMetadataNoStore() {
	deployBlock, err := l.fastBridge.DeployBlock(&bind.CallOpts{Context: l.GetTestContext()})
	l.NoError(err)

	// nothing stored, should use start block
	cl := listener.NewTestChainListener(listener.TestChainListenerArgs{
		Address:      l.fastBridge.Address(),
		InitialBlock: deployBlock.Uint64(),
		Client:       l.backend,
		Store:        l.store,
		Handler:      l.metrics,
	})

	startBlock, myChainID, err := cl.GetMetadata(l.GetTestContext())
	l.NoError(err)
	l.Equal(myChainID, uint64(chainID))
	l.Equal(startBlock, deployBlock.Uint64())
}

func (l *ListenerTestSuite) TestStartBlock() {
	cl := listener.NewTestChainListener(listener.TestChainListenerArgs{
		Address: l.fastBridge.Address(),
		Client:  l.backend,
		Store:   l.store,
		Handler: l.metrics,
	})

	deployBlock, err := l.fastBridge.DeployBlock(&bind.CallOpts{Context: l.GetTestContext()})
	l.NoError(err)

	expectedLastIndexed := deployBlock.Uint64() + 10
	err = l.store.PutLatestBlock(l.GetTestContext(), chainID, expectedLastIndexed)
	l.NoError(err)

	startBlock, cid, err := cl.GetMetadata(l.GetTestContext())
	l.Equal(cid, uint64(chainID))
	l.Equal(startBlock, expectedLastIndexed)
}

func (l *ListenerTestSuite) TestListen() {

}

// NewSqliteStore creates a new sqlite data store.
func NewSqliteStore(parentCtx context.Context, dbPath string, handler metrics.Handler) (_ *db.Store, err error) {
	logger := log.Logger("sqlite-store")

	logger.Debugf("creating sqlite store at %s", dbPath)

	ctx, span := handler.Tracer().Start(parentCtx, "start-sqlite")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// create the directory to the store if it doesn't exist
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite store")
	}

	logger.Warnf("submitter database is at %s/synapse.db", dbPath)

	namingStrategy := schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", gofakeit.Int64(), time.Now().Unix()),
	}

	gdb, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/%s", dbPath, "synapse.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   common_base.GetGormLogger(logger),
		FullSaveAssociations:                     true,
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           namingStrategy,
	})
	if err != nil {
		return nil, fmt.Errorf("could not connect to db %s: %w", dbPath, err)
	}

	handler.AddGormCallbacks(gdb)

	err = gdb.WithContext(ctx).AutoMigrate(txdb.GetAllModels()...)
	if err != nil {
		return nil, fmt.Errorf("could not migrate models: %w", err)
	}
	return db.NewChainListenerStore(gdb, handler), nil
}
