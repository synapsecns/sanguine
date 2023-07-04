package api_test

import (
	"testing"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql/base"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql/sqlite"
)

type RelayerAPISuite struct {
	*testsuite.TestSuite
	// testStore is the test store for the test
	testStore *base.Store
}

// NewTestSuite creates a new test suite.
func NewTestSuite(tb testing.TB) *RelayerAPISuite {
	tb.Helper()
	return &RelayerAPISuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *RelayerAPISuite) SetupSuite() {
	s.TestSuite.SetupSuite()
}

func (s *RelayerAPISuite) SetupTest() {
	s.TestSuite.SetupTest()

	// create the test store
	path := filet.TmpDir(s.T(), "")
	db, err := sqlite.NewSqliteStore(s.GetTestContext(), path, nil, false)
	s.Require().NoError(err)
	s.testStore = base.NewStore(db.DB(), nil)
}

func TestRelayerAPISuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
