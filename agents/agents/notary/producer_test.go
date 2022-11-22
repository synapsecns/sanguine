package notary_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

func (u NotarySuite) TestUpdateProducer() {
	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	testWallet, err := wallet.FromRandom()
	Nil(u.T(), err)

	signer := localsigner.NewSigner(testWallet.PrivateKey())

	// dispatch a random update
	auth := u.testBackend.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	tx, err := u.originContract.Dispatch(auth.TransactOpts, u.domainClient.Config().DomainID, [32]byte{}, gofakeit.Uint32(), encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	// call the update producing function
	attestationProducer := notary.NewAttestationProducer(u.domainClient, testDB, signer, 1*time.Second)

	err = attestationProducer.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	// TODO (joe): fix this after the global registry stuff
	/*producedAttestation, err := testDB.RetrieveSignedAttestationByNonce(u.GetTestContext(), u.domainClient.Config().DomainID, 0)
	Nil(u.T(), err)
	Equal(u.T(), producedAttestation.Attestation().Nonce(), uint32(1))*/
}
