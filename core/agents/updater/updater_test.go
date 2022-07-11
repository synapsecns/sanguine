package updater_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/updater"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db/datastore/sql"
)

func (u UpdaterSuite) TestUpdaterE2E() {
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

	auth := u.testBackend.GetTxContext(u.GetTestContext(), nil)
	tx, err := u.homeContract.Dispatch(auth.TransactOpts, gofakeit.Uint32(), [32]byte{}, gofakeit.Uint32(), []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	go func() {
		err = ud.Start(u.GetTestContext())
		Nil(u.T(), err)
	}()

	u.Eventually(func() bool {
		committedRoot, err := u.homeContract.CommittedRoot(&bind.CallOpts{Context: u.GetTestContext()})
		Nil(u.T(), err)

		return committedRoot != common.Hash{}
	})
}
