package attestationcollector_test

import (
	"math/big"

	"github.com/synapsecns/sanguine/core"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/types"
)

func (a AttestationCollectorSuite) TestAttestationCollectorSuite() {
	// Set up the contexts for all contracts, including the destination and attestation collector to get owner.
	txContextOrigin := a.testBackendOrigin.GetTxContext(a.GetTestContext(), nil)
	txContextDestination := a.testBackendDestination.GetTxContext(a.GetTestContext(), a.destinationContractMetadata.OwnerPtr())
	txContextAttestationCollector := a.testBackendDestination.GetTxContext(a.GetTestContext(), a.attestationContractMetadata.OwnerPtr())

	// TODO (joe): re-enable this test
	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	/*attestationSink := make(chan *attestationcollector.AttestationCollectorAttestationAccepted)
	subAttestation, err := a.attestationContract.WatchAttestationAccepted(&bind.WatchOpts{Context: a.GetTestContext()}, attestationSink)
	Nil(a.T(), err)*/

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(a.T(), err)

	// Dispatch an event from the Origin contract to be accepted on the Destination contract.
	tx, err := a.originContract.Dispatch(txContextOrigin.TransactOpts, 1, [32]byte{}, 1, encodedTips, nil)
	Nil(a.T(), err)
	a.testBackendOrigin.WaitForConfirmation(a.GetTestContext(), tx)

	// Create an attestation
	origin := uint32(a.testBackendOrigin.GetBigChainID().Uint64())
	destination := uint32(a.testBackendDestination.GetChainID())
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	attestKey := types.AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := notary.HashAttestation(unsignedAttestation)
	Nil(a.T(), err)

	encodedAttestation, err := types.EncodeAttestation(unsignedAttestation)
	Nil(a.T(), err)

	notarySignature, err := a.notarySigner.SignMessage(a.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(a.T(), err)

	guardSignature, err := a.guardSigner.SignMessage(a.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(a.T(), err)

	signedAttestation := types.NewSignedAttestation(
		unsignedAttestation,
		[]types.Signature{guardSignature},
		[]types.Signature{notarySignature})
	encodedGuardSignatures, err := types.EncodeSignatures(signedAttestation.GuardSignatures())
	Nil(a.T(), err)
	encodedNotarySignatures, err := types.EncodeSignatures(signedAttestation.NotarySignatures())
	Nil(a.T(), err)

	attestation, err := a.attestationHarness.FormatAttestation(
		&bind.CallOpts{Context: a.GetTestContext()},
		encodedAttestation,
		encodedGuardSignatures,
		encodedNotarySignatures,
	)
	Nil(a.T(), err)

	// Set notary to the testing address so we can submit attestations.
	/*tx, err = a.destinationContract.AddNotary(txContextDestination.TransactOpts, uint32(a.testBackendOrigin.GetChainID()), a.notarySigner.Address())
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)*/

	// Submit the attestation to get an AttestationAccepted event.
	tx, err = a.destinationContract.SubmitAttestation(txContextDestination.TransactOpts, attestation)
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)

	// Set notary to the testing address so we can submit attestations.
	/*tx, err = a.attestationContract.AddNotary(txContextAttestationCollector.TransactOpts, uint32(a.testBackendDestination.GetChainID()), a.signer.Address())
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)*/

	// Submit the attestation to get an AttestationSubmitted event.
	tx, err = a.attestationContract.SubmitAttestation(txContextAttestationCollector.TransactOpts, attestation)
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)

	// TODO (joe): get this working
	/*watchCtx, cancel := context.WithTimeout(a.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		a.T().Error(a.T(), fmt.Errorf("test context completed %w", a.GetTestContext().Err()))
	case <-subAttestation.Err():
		a.T().Error(a.T(), subAttestation.Err())
	// get AttestationSubmitted event
	case item := <-attestationSink:
		parser, err := attestationcollector.NewParser(a.attestationContract.Address())
		Nil(a.T(), err)
		// Check to see if the event was an AttestationSubmitted event.
		eventType, ok := parser.EventType(item.Raw)
		True(a.T(), ok)
		Equal(a.T(), eventType, attestationcollector.AttestationAcceptedEvent)

		break
	}*/
}
