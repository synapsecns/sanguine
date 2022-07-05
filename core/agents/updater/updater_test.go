package updater_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/updater"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db/datastore/sql"
)

func (u UpdaterSuite) TestUpdaterE2E() {
	u.T().Skip("could not update")

	testConfig := config.Config{
		Domains: map[string]config.DomainConfig{
			"test": u.domainClient.Config(),
		},
		Signer: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.wallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       sql.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
	}
	ud, err := updater.NewUpdater(u.GetTestContext(), testConfig)
	Nil(u.T(), err)

	go func() {
		err = ud.Start(u.GetTestContext())
		Nil(u.T(), err)
	}()
}
