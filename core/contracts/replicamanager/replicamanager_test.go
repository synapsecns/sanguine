package replicamanager_test

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
	"github.com/synapsecns/sanguine/core/contracts/replicamanager"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/synapse-node/contracts/bridge"
)

func (r ReplicaManagerSuite) TestReplicaManagerSuite() {
	// Set up contexts for both Home and Replica, also getting owner for Replica for reassigning updater role.
	txContextHome := r.testBackendHome.GetTxContext(r.GetTestContext(), nil)
	txContextReplica := r.testBackendReplica.GetTxContext(r.GetTestContext(), r.replicaContractMetadata.OwnerPtr())

	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	attestationSink := make(chan *replicamanager.ReplicaManagerAttestationAccepted)
	subAttestation, err := r.replicaContract.WatchAttestationAccepted(&bind.WatchOpts{Context: r.GetTestContext()}, attestationSink, []uint32{}, []uint32{}, [][32]byte{})
	Nil(r.T(), err)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(r.T(), err)

	// Dispatch an event from the Home contract to be accepted on the Replica contract.
	tx, err := r.homeContract.Dispatch(txContextHome.TransactOpts, 1, [32]byte{}, 1, encodedTips, nil)
	Nil(r.T(), err)
	r.testBackendHome.WaitForConfirmation(r.GetTestContext(), tx)

	// Create an attestation
	localDomain := uint32(r.testBackendHome.Config().ChainID)
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	unsignedAttestation := types.NewAttestation(localDomain, nonce, root)
	hashedAttestation, err := notary.HashAttestation(unsignedAttestation)
	Nil(r.T(), err)

	signature, err := r.signer.SignMessage(r.GetTestContext(), bridge.KappaToSlice(hashedAttestation), false)
	Nil(r.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, signature)
	encodedSig, err := types.EncodeSignature(signedAttestation.Signature())
	Nil(r.T(), err)

	attestation, err := r.attestationHarness.FormatAttestation(
		&bind.CallOpts{Context: r.GetTestContext()},
		signedAttestation.Attestation().Domain(),
		signedAttestation.Attestation().Nonce(),
		signedAttestation.Attestation().Root(),
		encodedSig,
	)
	Nil(r.T(), err)

	// Set updater to the testing address so we can submit attestations.
	tx, err = r.replicaContract.SetUpdater(txContextReplica.TransactOpts, uint32(r.testBackendHome.GetChainID()), r.signer.Address())
	Nil(r.T(), err)
	r.testBackendReplica.WaitForConfirmation(r.GetTestContext(), tx)

	// Submit the attestation to get an AttestationAccepted event.
	tx, err = r.replicaContract.SubmitAttestation(txContextReplica.TransactOpts, attestation)
	Nil(r.T(), err)
	r.testBackendReplica.WaitForConfirmation(r.GetTestContext(), tx)

	watchCtx, cancel := context.WithTimeout(r.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		r.T().Error(r.T(), fmt.Errorf("test context completed %w", r.GetTestContext().Err()))
	case <-subAttestation.Err():
		r.T().Error(r.T(), subAttestation.Err())
	// get dispatch event
	case item := <-attestationSink:
		parser, err := replicamanager.NewParser(r.replicaContract.Address())
		Nil(r.T(), err)

		// Check to see if the event was an AttestationAccepted event.
		eventType, ok := parser.EventType(item.Raw)
		True(r.T(), ok)
		Equal(r.T(), eventType, replicamanager.AttestationAcceptedEvent)

		break
	}
}

func (r ReplicaManagerSuite) TestUpdateTopic() {
	r.T().Skip("TODO: test this. Mocker should be able to mock this out")
}
