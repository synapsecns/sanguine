package evm_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/updater"
	pebble2 "github.com/synapsecns/sanguine/core/db/datastore/pebble"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/types"
)

func (i ContractSuite) TestSubmitAttestation() {
	attestationCollector, err := evm.NewAttestationCollectorContract(i.GetTestContext(), i.attestationBackend, i.attestationContract.Address())
	Nil(i.T(), err)

	localDomain := attestationDomain
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	unsignedAttestation := types.NewAttestation(uint32(localDomain), nonce, root)
	hashedAttestation, err := updater.HashAttestation(unsignedAttestation)
	Nil(i.T(), err)

	signature, err := i.signer.SignMessage(i.GetTestContext(), pebble2.ToSlice(hashedAttestation), false)
	Nil(i.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, signature)

	err = attestationCollector.SubmitAttestation(i.signer, signedAttestation)
	Nil(i.T(), err)

	// TODO retrieve once we have a way to
}
