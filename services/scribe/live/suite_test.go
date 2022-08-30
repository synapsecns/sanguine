package live_test

import (
	"testing"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
)

type LiveSuite struct {
	*testsuite.TestSuite
	testDB  db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
}

// NewLiveSuite creates a new live test suite.
func NewLiveSuite(tb testing.TB) *LiveSuite {
	tb.Helper()
	return &LiveSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (b *LiveSuite) SetupTest() {
	b.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(b.GetTestContext(), filet.TmpDir(b.T(), ""))
	Nil(b.T(), err)

	b.testDB = sqliteStore

	b.manager = testutil.NewDeployManager(b.T())

	b.wallet, err = wallet.FromRandom()
	Nil(b.T(), err)
	b.signer = localsigner.NewSigner(b.wallet.PrivateKey())
}

// TestLiveSuite tests the live suite.
func TestLiveSuite(t *testing.T) {
	suite.Run(t, NewLiveSuite(t))
}
