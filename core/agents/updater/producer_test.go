package updater_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/updater"
	"github.com/synapsecns/sanguine/core/db/datastore/pebble"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

func (u UpdaterSuite) TestUpdateProducer() {
	u.T().Skip("todo: test update producer alone")

	testDb, err := pebble.NewMessageDB(filet.TmpDir(u.T(), ""), "home1")
	Nil(u.T(), err)

	testWallet, err := wallet.FromRandom()
	Nil(u.T(), err)

	signer := localsigner.NewSigner(testWallet.PrivateKey())

	updateProducer := updater.NewUpdateProducer(u.domainClient, testDb, signer)
	_ = updateProducer

	go func() {
		Nil(u.T(), updateProducer.Start(u.GetTestContext()))
	}()

	_, home := u.deployManager.GetHome(u.GetTestContext(), u.testBackend)
	_ = home
}
