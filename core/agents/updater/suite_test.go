package updater_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/contracts/xappconfig"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/preset"
	"testing"
)

// UpdaterSuite tests the updater agent.
type UpdaterSuite struct {
	*testutils.TestSuite
	testBackend   backends.TestBackend
	deployManager *testutil.DeployManager
	xappConfig    *xappconfig.XAppConfigRef
	domainClient  domains.DomainClient
}

// NewUpdaterSuite creates a new updater suite.
func NewUpdaterSuite(tb testing.TB) *UpdaterSuite {
	tb.Helper()

	return &UpdaterSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (u *UpdaterSuite) SetupTest() {
	u.TestSuite.SetupTest()

	u.testBackend = preset.GetRinkeby().Geth(u.GetTestContext(), u.T())
	u.deployManager = testutil.NewDeployManager(u.T())
	_, u.xappConfig = u.deployManager.GetXAppConfig(u.GetTestContext(), u.testBackend)

	var err error
	u.domainClient, err = evm.NewEVM(u.GetTestContext(), "updater", config.DomainConfig{
		DomainID:          1,
		Type:              types.EVM.String(),
		XAppConfigAddress: u.xappConfig.Address().String(),
		RPCUrl:            u.testBackend.RPCAddress(),
	})
	Nil(u.T(), err)
}

func TestUpdaterSuite(t *testing.T) {
	suite.Run(t, NewUpdaterSuite(t))
}
