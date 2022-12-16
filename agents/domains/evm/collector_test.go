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
	// TODO (joe): Get this test working
	/*attestationCollector*/
	_, err := evm.NewAttestationCollectorContract(i.GetTestContext(), i.attestationBackend, i.attestationContract.Address())
	Nil(i.T(), err)

	originDomain, err := i.originContract.OriginHarness.LocalDomain(&bind.CallOpts{Context: i.GetTestContext()})
	Nil(i.T(), err)
	localDomain := originDomain
	destination := testDestinationDomain
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	attestKey := types.AttestationKey{
		Origin:      localDomain,
		Destination: destination,
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := notary.HashAttestation(unsignedAttestation)
	Nil(i.T(), err)

	// TODO (joe): Get this working
	/*signature*/
	_, err = i.signer.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(i.T(), err)

	/*signedAttestation := types.NewSignedAttestation(unsignedAttestation, signature)
	err = attestationCollector.SubmitAttestation(i.GetTestContext(), i.signer, signedAttestation)
	Nil(i.T(), err)

	latestNonce, err := attestationCollector.GetLatestNonce(i.GetTestContext(), localDomain, destination, i.signer)
	Nil(i.T(), err)
	Equal(i.T(), nonce, latestNonce)*/
}
