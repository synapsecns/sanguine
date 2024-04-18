package util

import (
	"gorm.io/gorm/schema"
	"os"
)

// NamingStrategy is used to exported here to avoid a circular dependency.
var NamingStrategy = schema.NamingStrategy{
	TablePrefix: os.Getenv("TABLE_PREFIX"),
}
