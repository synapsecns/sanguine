package guard_test

import (
	"testing"
	"time"

	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
)

// GuardSuite tests the guard agent.
type GuardSuite struct {
	*testutil.SimulatedBackendsTestSuite
	attestationDomainClient domains.DomainClient
}

// NewGuardSuite creates a new guard suite.
func NewGuardSuite(tb testing.TB) *GuardSuite {
	tb.Helper()

	return &GuardSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (u *GuardSuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	u.SimulatedBackendsTestSuite.SetupTest()

	var err error
	u.originDomainClient, err = evm.NewEVM(u.GetTestContext(), "origin_guard", config.DomainConfig{
		DomainID:                    uint32(u.TestBackendOrigin.GetBigChainID().Uint64()),
		Type:                        types.EVM.String(),
		OriginAddress:               u.OriginContract.Address().String(),
		AttestationCollectorAddress: u.AttestationContract.Address().String(),
		RPCUrl:                      u.TestBackendOrigin.RPCAddress(),
	})
	Nil(u.T(), err)

	u.destinationDomainClient, err = evm.NewEVM(u.GetTestContext(), "destination_guard", config.DomainConfig{
		DomainID:                    uint32(u.TestBackendDestination.GetBigChainID().Uint64()),
		Type:                        types.EVM.String(),
		DestinationAddress:          u.DestinationContract.Address().String(),
		AttestationCollectorAddress: u.AttestationContract.Address().String(),
		RPCUrl:                      u.TestBackendDestination.RPCAddress(),
	})
	Nil(u.T(), err)
}

func TestGuardSuite(t *testing.T) {
	suite.Run(t, NewGuardSuite(t))
}
