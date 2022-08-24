package evm_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"math/big"
)

func (i ContractSuite) TestSubmitAttestation() {
	attestationCollector, err := evm.NewAttestationCollectorContract(i.GetTestContext(), i.attestationBackend, i.attestationContract.Address())
	Nil(i.T(), err)

	localDomain := uint32(attestationDomain)
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	unsignedAttestation := types.NewAttestation(localDomain, nonce, root)
	hashedAttestation, err := notary.HashAttestation(unsignedAttestation)
	Nil(i.T(), err)

	signature, err := i.signer.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(i.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, signature)

	err = attestationCollector.SubmitAttestation(i.GetTestContext(), i.signer, signedAttestation)
	Nil(i.T(), err)

	latestNonce, err := attestationCollector.LatestNonce(i.GetTestContext(), localDomain, i.signer)
	Nil(i.T(), err)
	Equal(i.T(), nonce, latestNonce)
}
