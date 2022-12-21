package guard_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
)

func (u GuardSuite) TestAttestationCollectorAttestationScanner() {
	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)
	auth := u.TestBackendAttestation.GetTxContext(u.GetTestContext(), nil)

	origin := u.TestBackendOrigin.GetChainID()
	destination := u.TestBackendDestination.GetChainID()
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	attestKey := types.AttestationKey{
		Origin:      uint32(origin),
		Destination: uint32(destination),
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	Nil(u.T(), err)

	guardSignature, err := u.GuardSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(u.T(), err)

	notarySignature, err := u.NotarySigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(u.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(u.T(), err)

	tx, err := u.AttestationContract.SubmitAttestation(auth.TransactOpts, rawSignedAttestation)
	Nil(u.T(), err)

	u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), tx)

	// call the update producing function
	attestationCollectorAttestationScanner := guard.NewAttestationCollectorAttestationScanner(u.dom, u.destinationID, testDB, u.signer, 1*time.Second)

	err = originAttestationScanner.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedAttestation, err := testDB.RetrieveOldestUnsignedInProgressAttestation(u.GetTestContext(), u.domainClient.Config().DomainID, u.destinationID)
	Nil(u.T(), err)
	Equal(u.T(), producedAttestation.SignedAttestation().Attestation().Nonce(), uint32(1))
	Equal(u.T(), types.AttestationStateNotaryUnsigned, producedAttestation.AttestationState())
}
