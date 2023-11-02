package api_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/sinner/api"
)

func (t *APISuite) TestInitDB() {
	// Test with sqlite
	db, err := api.InitDB(t.GetTestContext(), dbcommon.Sqlite.String(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)
	NotNil(t.T(), db)

	// Test with invalid db type
	db, err = api.InitDB(t.GetTestContext(), "invalidDBType", filet.TmpDir(t.T(), ""), t.metrics, false)
	NotNil(t.T(), err)
	Nil(t.T(), db)
	Equal(t.T(), "invalid databaseType type: invalidDBType", err.Error())
}
