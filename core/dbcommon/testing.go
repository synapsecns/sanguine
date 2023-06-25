package dbcommon

import (
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"os"
)

const (
	// EnableMysqlTestVar is the environment variable to enable mysql tests.
	EnableMysqlTestVar = "ENABLE_MYSQL_TEST"
	// MysqlDatabaseVar is the environment variable for the mysql database dsn.
	MysqlDatabaseVar = "MYSQL_DATABASE"
	// MysqlUserVar is the environment variable for the mysql user.
	MysqlUserVar = "MYSQL_USER"
	// MysqlPasswordVar is the environment variable for the mysql password.
	MysqlPasswordVar = "MYSQL_PASSWORD"
	// MysqlHostVar is the environment variable for the mysql host.
	MysqlHostVar = "MYSQL_HOST"
	// MysqlPortVar is the environment variable for the mysql port.
	MysqlPortVar = "MYSQL_PORT"
)

// GetTestConnString returns the connection string for the mysql test database.
// this is derived from environment variables.
// TODO: test this in ci.
func GetTestConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv(MysqlUserVar, "root"), os.Getenv(MysqlPasswordVar), core.GetEnv(MysqlHostVar, "127.0.0.1"), core.GetEnvInt(MysqlPortVar, 3306), os.Getenv(MysqlDatabaseVar))
}
