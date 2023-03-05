package attestationcollector_test

import (
	. "github.com/stretchr/testify/assert"
)

//nolint:unused
func (a AttestationCollectorSuite) launchTest(amountGuards, amountNotaries int) {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	GreaterOrEqual(a.T(), amountGuards+amountNotaries, 1)
	LessOrEqual(a.T(), amountGuards, 1)
	LessOrEqual(a.T(), amountNotaries, 1)

	// TODO (joeallen): FIX ME
	// txContextAttestationCollector := a.TestBackendAttestation.GetTxContext(a.GetTestContext(), a.AttestationContractMetadata.OwnerPtr())

	// TODO (joeallen): FIX ME
	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	// attestationSink := make(chan *attestationcollector.AttestationCollectorAttestationAccepted)
	// subAttestation, err := a.AttestationContract.WatchAttestationAccepted(&bind.WatchOpts{Context: a.GetTestContext()}, attestationSink)
	// Nil(a.T(), err)

	// Create an attestation
	// origin := uint32(a.TestBackendOrigin.GetBigChainID().Uint64())
	// destination := uint32(a.TestBackendDestination.GetChainID())
	// destination := origin + 1
	// nonce := gofakeit.Uint32()
	// root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	// attestKey := types.AttestationKey{
	//	Origin:      origin,
	//	Destination: destination,
	//	Nonce:       nonce,
	//}
	// unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	//hashedAttestation, err := types.Hash(unsignedAttestation)
	//Nil(a.T(), err)

	// TODO (joeallen): FIX ME
	// encodedAttestation, err := types.EncodeAttestation(unsignedAttestation)
	// Nil(a.T(), err)

	// notarySignatures := []types.Signature{}
	// if amountNotaries == 1 {
	//	notarySignature, err := a.NotaryBondedSigner.SignMessage(a.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	//	Nil(a.T(), err)
	//	notarySignatures = append(notarySignatures, notarySignature)
	//}
	// guardSignatures := []types.Signature{}
	// if amountGuards == 1 {
	//	guardSignature, err := a.GuardBondedSigner.SignMessage(a.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	//	Nil(a.T(), err)
	//	guardSignatures = append(guardSignatures, guardSignature)
	//}
	// TODO (joeallen): FIX ME
	// signedAttestation := types.NewSignedAttestation(
	//	unsignedAttestation,
	//	guardSignatures,
	//	notarySignatures)
	// encodedGuardSignatures, err := types.EncodeSignatures(signedAttestation.GuardSignatures())
	// Nil(a.T(), err)
	//encodedNotarySignatures, err := types.EncodeSignatures(signedAttestation.NotarySignatures())
	//Nil(a.T(), err)

	// attestation, err := a.AttestationHarness.FormatAttestation(
	//	&bind.CallOpts{Context: a.GetTestContext()},
	//	encodedAttestation,
	//	encodedGuardSignatures,
	//	encodedNotarySignatures,
	//)
	// Nil(a.T(), err)

	// Submit the attestation to get an AttestationSubmitted event.
	// txSubmitAttestation, err := a.AttestationContract.SubmitAttestation(txContextAttestationCollector.TransactOpts, attestation)
	// Nil(a.T(), err)
	// a.TestBackendAttestation.WaitForConfirmation(a.GetTestContext(), txSubmitAttestation)

	// watchCtx, cancel := context.WithTimeout(a.GetTestContext(), time.Second*10)
	// defer cancel()

	// select {
	// check for errors and fail
	// case <-watchCtx.Done():
	//	a.T().Error(a.T(), fmt.Errorf("test context completed %w", a.GetTestContext().Err()))
	// case <-subAttestation.Err():
	//	a.T().Error(a.T(), subAttestation.Err())
	// get AttestationSubmitted event
	// case item := <-attestationSink:
	//	parser, err := attestationcollector.NewParser(a.AttestationContract.Address())
	//	Nil(a.T(), err)
	// Check to see if the event was an AttestationSubmitted event.
	//	eventType, ok := parser.EventType(item.Raw)
	//	True(a.T(), ok)
	//	Equal(a.T(), eventType, attestationcollector.AttestationAcceptedEvent)

	//	break
	//}
}

// TestAttestationCollectorSuite tests submitting an attesation with one guard and one notary.
func (a AttestationCollectorSuite) TestSubmitAttestationOneGuardOneNotary() {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	a.launchTest(1, 1)
}

// TestSubmitAttestationOnlyOneNotary tests submitting an attesation with only one notary.
func (a AttestationCollectorSuite) TestSubmitAttestationOnlyOneNotary() {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	a.launchTest(0, 1)
}

// TestSubmitAttestationOnlyOneGuard tests submitting an attesation with only one guard.
func (a AttestationCollectorSuite) TestSubmitAttestationOnlyOneGuard() {
	// TODO (joeallen): FIX ME
	a.T().Skip()
	a.launchTest(1, 0)
}
