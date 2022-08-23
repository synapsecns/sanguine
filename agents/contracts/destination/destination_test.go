package destination_test

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/synapse-node/contracts/bridge"
)

func (d DestinationSuite) TestDestinationSuite() {
	// Set up contexts for both Origin and Destination, also getting owner for Destination for reassigning notary role.
	txContextOrigin := d.testBackendOrigin.GetTxContext(d.GetTestContext(), nil)
	txContextDestination := d.testBackendDestination.GetTxContext(d.GetTestContext(), d.destinationContractMetadata.OwnerPtr())

	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	attestationSink := make(chan *destination.DestinationAttestationAccepted)
	subAttestation, err := d.destinationContract.WatchAttestationAccepted(&bind.WatchOpts{Context: d.GetTestContext()}, attestationSink, []uint32{}, []uint32{}, [][32]byte{})
	Nil(d.T(), err)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(d.T(), err)

	// Dispatch an event from the Origin contract to be accepted on the Destination contract.
	tx, err := d.originContract.Dispatch(txContextOrigin.TransactOpts, 1, [32]byte{}, 1, encodedTips, nil)
	Nil(d.T(), err)
	d.testBackendOrigin.WaitForConfirmation(d.GetTestContext(), tx)

	// Create an attestation
	localDomain := uint32(d.testBackendOrigin.Config().ChainID)
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	unsignedAttestation := types.NewAttestation(localDomain, nonce, root)
	hashedAttestation, err := notary.HashAttestation(unsignedAttestation)
	Nil(d.T(), err)

	signature, err := d.signer.SignMessage(d.GetTestContext(), bridge.KappaToSlice(hashedAttestation), false)
	Nil(d.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, signature)
	encodedSig, err := types.EncodeSignature(signedAttestation.Signature())
	Nil(d.T(), err)

	attestation, err := d.attestationHarness.FormatAttestation(
		&bind.CallOpts{Context: d.GetTestContext()},
		signedAttestation.Attestation().Domain(),
		signedAttestation.Attestation().Nonce(),
		signedAttestation.Attestation().Root(),
		encodedSig,
	)
	Nil(d.T(), err)

	// Set notary to the testing address so we can submit attestations.
	tx, err = d.destinationContract.SetNotary(txContextDestination.TransactOpts, uint32(d.testBackendOrigin.GetChainID()), d.signer.Address())
	Nil(d.T(), err)
	d.testBackendDestination.WaitForConfirmation(d.GetTestContext(), tx)

	// Submit the attestation to get an AttestationAccepted event.
	tx, err = d.destinationContract.SubmitAttestation(txContextDestination.TransactOpts, attestation)
	Nil(d.T(), err)
	d.testBackendDestination.WaitForConfirmation(d.GetTestContext(), tx)

	watchCtx, cancel := context.WithTimeout(d.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		d.T().Error(d.T(), fmt.Errorf("test context completed %w", d.GetTestContext().Err()))
	case <-subAttestation.Err():
		d.T().Error(d.T(), subAttestation.Err())
	// get attestation accepted event
	case item := <-attestationSink:
		parser, err := destination.NewParser(d.destinationContract.Address())
		Nil(d.T(), err)

		// Check to see if the event was an AttestationAccepted event.
		eventType, ok := parser.EventType(item.Raw)
		True(d.T(), ok)
		Equal(d.T(), eventType, destination.AttestationAcceptedEvent)

		break
	}
}
