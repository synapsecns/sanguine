package util

import "gorm.io/gorm/schema"

// NamingStrategy is used to exported here to avoid a circular dependency.
var NamingStrategy = schema.NamingStrategy{}
