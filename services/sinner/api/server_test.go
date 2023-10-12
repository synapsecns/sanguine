package api_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/services/sinner/api"
	"github.com/synapsecns/sanguine/services/sinner/db"
)

type MockSqliteStore struct {
	mock.Mock
}

func (m *MockSqliteStore) DB() db.EventDB {
	args := m.Called()
	return args.Get(0).(db.EventDB)
}

type MockMysqlStore struct {
	mock.Mock
}

func (m *MockMysqlStore) DB() db.EventDB {
	args := m.Called()
	return args.Get(0).(db.EventDB)
}

func (t *APISuite) TestInitDB() {
	// Mock SQLite and MySQL stores
	sqliteStore := new(MockSqliteStore)
	sqliteStore.On("DB").Return(nil)

	mysqlStore := new(MockMysqlStore)
	mysqlStore.On("DB").Return(nil)

	// Test with sqlite
	db, err := api.InitDB(t.GetTestContext(), "sqlite", filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)
	NotNil(t.T(), db)

	// Test with invalid db type
	db, err = api.InitDB(t.GetTestContext(), "invalidDBType", filet.TmpDir(t.T(), ""), t.metrics, false)
	NotNil(t.T(), err)
	Nil(t.T(), db)
	Equal(t.T(), "invalid databaseType type: invalidDBType", err.Error())
}
