package updater_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/updater"
	"github.com/synapsecns/sanguine/core/db/datastore/pebble"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"time"
)

func (u UpdaterSuite) TestUpdateProducer() {
	testDB, err := pebble.NewMessageDB(filet.TmpDir(u.T(), ""), "home1")
	Nil(u.T(), err)

	testWallet, err := wallet.FromRandom()
	Nil(u.T(), err)

	signer := localsigner.NewSigner(testWallet.PrivateKey())

	// dispatch a random update
	auth := u.testBackend.GetTxContext(u.GetTestContext(), nil)
	tx, err := u.homeContract.Dispatch(auth.TransactOpts, gofakeit.Uint32(), [32]byte{}, gofakeit.Uint32(), []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	// call the update producing function
	updateProducer := updater.NewUpdateProducer(u.domainClient, testDB, signer, 1*time.Second)

	err = updateProducer.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedUpdate, err := testDB.RetrieveProducedUpdate(common.Hash{})
	Nil(u.T(), err)
	Equal(u.T(), producedUpdate.Update().PreviousRoot(), common.Hash{})
}
