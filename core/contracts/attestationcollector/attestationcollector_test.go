package attestationcollector_test

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/notary"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/synapse-node/contracts/bridge"
)

func (a AttestationCollectorSuite) TestAttestationCollectorSuite() {
	// Set up the contexts for all contracts, including the destination and attestation collector to get owner.
	txContextOrigin := a.testBackendOrigin.GetTxContext(a.GetTestContext(), nil)
	txContextDestination := a.testBackendDestination.GetTxContext(a.GetTestContext(), a.destinationContractMetadata.OwnerPtr())
	txContextAttestationCollector := a.testBackendDestination.GetTxContext(a.GetTestContext(), a.attestationContractMetadata.OwnerPtr())

	// Create a channel and subscription to receive AttestationSubmitted events as they are emitted.
	attestationSink := make(chan *attestationcollector.AttestationCollectorAttestationSubmitted)
	subAttestation, err := a.attestationContract.WatchAttestationSubmitted(&bind.WatchOpts{Context: a.GetTestContext()}, attestationSink, []common.Address{})
	Nil(a.T(), err)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(a.T(), err)

	// Dispatch an event from the Origin contract to be accepted on the Destination contract.
	tx, err := a.originContract.Dispatch(txContextOrigin.TransactOpts, 1, [32]byte{}, 1, encodedTips, nil)
	Nil(a.T(), err)
	a.testBackendOrigin.WaitForConfirmation(a.GetTestContext(), tx)

	// Create an attestation
	localDomain := uint32(a.testBackendOrigin.Config().ChainID)
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	unsignedAttestation := types.NewAttestation(localDomain, nonce, root)
	hashedAttestation, err := notary.HashAttestation(unsignedAttestation)
	Nil(a.T(), err)

	signature, err := a.signer.SignMessage(a.GetTestContext(), bridge.KappaToSlice(hashedAttestation), false)
	Nil(a.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, signature)
	encodedSig, err := types.EncodeSignature(signedAttestation.Signature())
	Nil(a.T(), err)

	attestation, err := a.attestationHarness.FormatAttestation(
		&bind.CallOpts{Context: a.GetTestContext()},
		signedAttestation.Attestation().Domain(),
		signedAttestation.Attestation().Nonce(),
		signedAttestation.Attestation().Root(),
		encodedSig,
	)
	Nil(a.T(), err)

	// Set notary to the testing address so we can submit attestations.
	tx, err = a.destinationContract.SetNotary(txContextDestination.TransactOpts, uint32(a.testBackendOrigin.GetChainID()), a.signer.Address())
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)

	// Submit the attestation to get an AttestationAccepted event.
	tx, err = a.destinationContract.SubmitAttestation(txContextDestination.TransactOpts, attestation)
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)

	// Set notary to the testing address so we can submit attestations.
	tx, err = a.attestationContract.AddNotary(txContextAttestationCollector.TransactOpts, uint32(a.testBackendOrigin.GetChainID()), a.signer.Address())
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)

	// Submit the attestation to get an AttestationSubmitted event.
	tx, err = a.attestationContract.SubmitAttestation(txContextAttestationCollector.TransactOpts, attestation)
	Nil(a.T(), err)
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), tx)

	watchCtx, cancel := context.WithTimeout(a.GetTestContext(), time.Second*10)
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
		Equal(a.T(), eventType, attestationcollector.AttestationSubmittedEvent)

		break
	}
}
