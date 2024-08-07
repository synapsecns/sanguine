package service_test

import (
	"math/big"
	"sync"
	"testing"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
)

type RelayerTestSuite struct {
	*testsuite.TestSuite
	originBackend backends.SimulatedTestBackend
	destBackend   backends.SimulatedTestBackend
	manager       *testutil.DeployManager
	metrics       metrics.Handler
	cfg           relconfig.Config
}

func NewRelayerTestSuite(tb testing.TB) *RelayerTestSuite {
	return &RelayerTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestRelayerTestSuite(t *testing.T) {
	suite.Run(t, NewRelayerTestSuite(t))
}

func (r *RelayerTestSuite) SetupTest() {
	r.TestSuite.SetupTest()
	r.manager = testutil.NewDeployManager(r.T())
	r.metrics = metrics.NewNullHandler()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		r.originBackend = geth.NewEmbeddedBackendForChainID(r.GetTestContext(), r.T(), big.NewInt(1))
	}()
	go func() {
		defer wg.Done()
		r.destBackend = geth.NewEmbeddedBackendForChainID(r.GetTestContext(), r.T(), big.NewInt(2))
	}()
	wg.Wait()

	serverURL := omnirpcHelper.NewOmnirpcServer(r.GetTestContext(), r.T(), r.destBackend, r.originBackend)

	originContract, _ := r.manager.GetMockFastBridge(r.GetTestContext(), r.originBackend)
	destContract, _ := r.manager.GetMockFastBridge(r.GetTestContext(), r.destBackend)

	r.cfg = relconfig.Config{
		Database: relconfig.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  filet.TmpDir(r.T(), ""),
		},
		Chains: map[int]relconfig.ChainConfig{
			int(r.originBackend.GetChainID()): {
				RFQAddress: originContract.Address().String(),
			},
			int(r.destBackend.GetChainID()): {
				RFQAddress: destContract.Address().String(),
			},
		},
		OmniRPCURL: serverURL,
	}
}
