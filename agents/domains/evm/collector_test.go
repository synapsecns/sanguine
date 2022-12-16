package evm_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
)

func (i ContractSuite) TestSubmitAttestation() {
	attestationCollector, err := evm.NewAttestationCollectorContract(i.GetTestContext(), i.attestationBackend, i.attestationContract.Address())
	Nil(i.T(), err)

	ownerPtr, err := i.attestationContract.AttestationCollectorCaller.Owner(&bind.CallOpts{Context: i.GetTestContext()})
	Nil(i.T(), err)

	attestationOwnerAuth := i.attestationBackend.GetTxContext(i.GetTestContext(), &ownerPtr)
	tx, err := i.attestationContract.AddAgent(attestationOwnerAuth.TransactOpts, testDestinationDomain, i.signer.Address())
	Nil(i.T(), err)
	i.attestationBackend.WaitForConfirmation(i.GetTestContext(), tx)

	notaries, err := i.attestationContract.AllAgents(&bind.CallOpts{Context: i.GetTestContext()}, testDestinationDomain)
	Nil(i.T(), err)
	Len(i.T(), notaries, 1)

	originDomain, err := i.originContract.OriginHarness.LocalDomain(&bind.CallOpts{Context: i.GetTestContext()})

	Nil(i.T(), err)
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	attestKey := types.AttestationKey{
		Origin:      originDomain,
		Destination: testDestinationDomain,
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := notary.HashAttestation(unsignedAttestation)
	Nil(i.T(), err)

	signature, err := i.signer.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(i.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{}, []types.Signature{signature})
	err = attestationCollector.SubmitAttestation(i.GetTestContext(), i.signer, signedAttestation)
	Nil(i.T(), err)

	latestNonce, err := attestationCollector.GetLatestNonce(i.GetTestContext(), originDomain, testDestinationDomain, i.signer)
	Nil(i.T(), err)
	Equal(i.T(), nonce, latestNonce)
}
