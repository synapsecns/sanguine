package db_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/validator/db"
)

// StoresAndRetrievesMessages tests storage/retreival
func (d *DBSuite) StoresAndRetrievesMessages() {
	newDB, err := db.NewDB(filet.TmpDir(d.T(), ""))
	Nil(d.T(), err)

	_ = newDB
}
