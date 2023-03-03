package destination_test

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/types"
)

func (d DestinationSuite) TestDestinationSuite() {
	// TODO (joeallen): FIX ME
	d.T().Skip()
	// TODO (joeallen): FIX ME
	var err error
	Nil(d.T(), err)
	// Set up contexts for both Origin and Destination, also getting owner for Destination for reassigning notary role.

	txContextOrigin := d.TestBackendOrigin.GetTxContext(d.GetTestContext(), nil)
	//txContextDestination := d.TestBackendDestination.GetTxContext(d.GetTestContext(), d.DestinationContractMetadata.OwnerPtr())

	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	attestationSink := make(chan *destinationharness.DestinationHarnessAttestationAccepted)
	subAttestation, err := d.DestinationContract.WatchAttestationAccepted(&bind.WatchOpts{
		Context: d.GetTestContext()},
		attestationSink)
	Nil(d.T(), err)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(d.T(), err)

	// Dispatch an event from the Origin contract to be accepted on the Destination contract.
	tx, err := d.OriginContract.Dispatch(txContextOrigin.TransactOpts, 1, [32]byte{}, 1, encodedTips, nil)
	Nil(d.T(), err)
	d.TestBackendOrigin.WaitForConfirmation(d.GetTestContext(), tx)

	// Create an attestation
	// TODO (joeallen): FIX ME
	//originDomain := uint32(d.TestBackendOrigin.GetBigChainID().Uint64())
	//destinationDomain := uint32(d.TestBackendDestination.GetBigChainID().Uint64())
	//nonce := gofakeit.Uint32()
	//attestationKey := types.AttestationKey{
	//	Origin:      originDomain,
	//	Destination: destinationDomain,
	//	Nonce:       nonce,
	//}

	//root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	//unsignedAttestation := types.NewAttestation(attestationKey.GetRawKey(), root)
	//hashedAttestation, err := types.Hash(unsignedAttestation)
	//Nil(d.T(), err)

	//notarySignature, err := d.NotaryBondedSigner.SignMessage(d.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	//Nil(d.T(), err)

	//guardSignature, err := d.GuardBondedSigner.SignMessage(d.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	//Nil(d.T(), err)

	//signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	//rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	//Nil(d.T(), err)

	//tx, err = d.DestinationContract.SubmitAttestation(txContextDestination.TransactOpts, rawSignedAttestation)
	//Nil(d.T(), err)

	d.TestBackendDestination.WaitForConfirmation(d.GetTestContext(), tx)

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
		parser, err := destination.NewParser(d.DestinationContract.Address())
		Nil(d.T(), err)

		// Check to see if the event was an AttestationAccepted event.
		eventType, ok := parser.EventType(item.Raw)
		True(d.T(), ok)
		Equal(d.T(), eventType, destination.AttestationAcceptedEvent)

		emittedSignedAttesation, err := types.DecodeSignedAttestation(item.Attestation)
		Nil(d.T(), err)

		Equal(d.T(), d.OriginDomainClient.Config().DomainID, emittedSignedAttesation.Attestation().Origin())
		Equal(d.T(), d.DestinationDomainClient.Config().DomainID, emittedSignedAttesation.Attestation().Destination())
		//Equal(d.T(), nonce, emittedSignedAttesation.Attestation().Nonce())
		//Equal(d.T(), [32]byte(root), emittedSignedAttesation.Attestation().Root())

		break
	}
}
