package livefill_test

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

type LivefillSuite struct {
	*testsuite.TestSuite
	testDB  db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
}

// NewLivefillSuite creates a new backfill test suite.
func NewLivefillSuite(tb testing.TB) *LivefillSuite {
	tb.Helper()
	return &LivefillSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (l *LivefillSuite) SetupTest() {
	l.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(l.GetTestContext(), filet.TmpDir(l.T(), ""))
	Nil(l.T(), err)

	l.testDB = sqliteStore

	l.manager = testutil.NewDeployManager(l.T())

	l.wallet, err = wallet.FromRandom()
	Nil(l.T(), err)
	l.signer = localsigner.NewSigner(l.wallet.PrivateKey())
}

// TestLivefillSuite tests the backfill suite.
func TestLivefillSuite(t *testing.T) {
	suite.Run(t, NewLivefillSuite(t))
}
