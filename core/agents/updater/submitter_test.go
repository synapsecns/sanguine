package updater_test

import (
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/updater"
	"github.com/synapsecns/sanguine/core/db/datastore/pebble"
	"github.com/synapsecns/sanguine/core/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/core/types"
	"time"
)

func (u UpdaterSuite) TestUpdateSubmitter() {
	txDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	messageDB, err := pebble.NewMessageDB(filet.TmpDir(u.T(), ""), "home1")
	Nil(u.T(), err)

	update := types.NewUpdate()
	messageDB.StoreProducedUpdate(common.Hash{}, types.NewSignedUpdate())

	updateSubmitter := updater.NewUpdateSubmitter(u.domainClient, messageDB, txDB, u.signer, time.Duration(0))
	_, err = updateSubmitter.Update(u.GetTestContext(), common.Hash{})
}
