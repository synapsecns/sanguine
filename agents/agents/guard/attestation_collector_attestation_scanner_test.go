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
	nonce := uint32(1)
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	attestKey := types.AttestationKey{
		Origin:      uint32(origin),
		Destination: uint32(destination),
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	Nil(u.T(), err)

	notarySignature, err := u.NotaryBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(u.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(u.T(), err)

	tx, err := u.AttestationContract.SubmitAttestation(auth.TransactOpts, rawSignedAttestation)
	Nil(u.T(), err)

	u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), tx)

	// call the update producing function
	attestationCollectorAttestationScanner := guard.NewAttestationCollectorAttestationScanner(
		u.AttestationDomainClient,
		uint32(origin),
		uint32(destination),
		testDB,
		u.UnbondedSigner,
		1*time.Second)

	err = attestationCollectorAttestationScanner.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	retrievedConfirmedInProgressAttestation, err := testDB.RetrieveOldestGuardUnsignedAndUnverifiedInProgressAttestation(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID)

	Nil(u.T(), err)
	NotNil(u.T(), retrievedConfirmedInProgressAttestation)

	retrievedSignedAttestation := retrievedConfirmedInProgressAttestation.SignedAttestation()
	Equal(u.T(), u.OriginDomainClient.Config().DomainID, retrievedSignedAttestation.Attestation().Origin())
	Equal(u.T(), u.DestinationDomainClient.Config().DomainID, retrievedSignedAttestation.Attestation().Destination())
	Equal(u.T(), root, common.Hash(retrievedSignedAttestation.Attestation().Root()))
	Len(u.T(), retrievedSignedAttestation.NotarySignatures(), 1)
	Len(u.T(), retrievedSignedAttestation.GuardSignatures(), 0)
	Equal(u.T(), types.AttestationStateGuardUnsignedAndUnverified, retrievedConfirmedInProgressAttestation.AttestationState())

	Nil(u.T(), err)
}
