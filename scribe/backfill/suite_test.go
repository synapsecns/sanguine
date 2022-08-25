package backfill_test

import (
	"testing"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/scribe/db"
	"github.com/synapsecns/sanguine/scribe/db/datastore/sql/sqlite"
)

type BackfillSuite struct {
	*testsuite.TestSuite
	testDB  db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
}

// NewBackfillSuite creates a new backfill test suite.
func NewBackfillSuite(tb testing.TB) *BackfillSuite {
	tb.Helper()
	return &BackfillSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (b *BackfillSuite) SetupTest() {
	b.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(b.GetTestContext(), filet.TmpDir(b.T(), ""))
	Nil(b.T(), err)

	b.testDB = sqliteStore

	b.manager = testutil.NewDeployManager(b.T())

	b.wallet, err = wallet.FromRandom()
	Nil(b.T(), err)
	b.signer = localsigner.NewSigner(b.wallet.PrivateKey())
}

// TestBackfillSuite tests the backfill suite.
func TestBackfillSuite(t *testing.T) {
	suite.Run(t, NewBackfillSuite(t))
}
