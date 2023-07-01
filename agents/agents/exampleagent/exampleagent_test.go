package exampleagent_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

func (u ExampleAgentSuite) TestExampleAgentSimulatedTestSuite() {
	NotNil(u.T(), u.SimulatedBackendsTestSuite)

	notaryStatus, err := u.BondingManagerOnSummit.AgentStatus(&bind.CallOpts{Context: u.GetTestContext()}, u.NotaryBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), uint32(u.TestBackendDestination.GetChainID()), notaryStatus.Domain)

	notaryStatusFromEVM, err := u.SummitDomainClient.BondingManager().GetAgentStatus(u.GetTestContext(), u.NotaryBondedSigner)
	Nil(u.T(), err)
	Equal(u.T(), notaryStatusFromEVM.Domain(), uint32(u.TestBackendDestination.GetChainID()))

	guardStatus, err := u.SummitContract.AgentStatus(&bind.CallOpts{Context: u.GetTestContext()}, u.GuardBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), uint32(0), guardStatus.Domain)

	guardStatusFromEVM, err := u.SummitDomainClient.BondingManager().GetAgentStatus(u.GetTestContext(), u.GuardBondedSigner)
	Nil(u.T(), err)
	Equal(u.T(), guardStatusFromEVM.Domain(), uint32(0))
}

func (u ExampleAgentSuite) TestEncodeDecodeAttestationSignature() {
	var snapRoot [32]byte
	var dataHash [32]byte
	for i := 0; i < 32; i++ {
		snapRoot[i] = byte(i)
		dataHash[i] = byte(i)
	}

	badBlockNumber := uint64(100)
	badTimeStamp := uint64(1000)
	badBlockNumberBigInt := new(big.Int).SetUint64(badBlockNumber)
	badTimeStampBigInt := new(big.Int).SetUint64(badTimeStamp)
	attestation := types.NewAttestation(
		snapRoot,
		dataHash,
		uint32(200),
		badBlockNumberBigInt,
		badTimeStampBigInt)

	encodedAttestation, err := types.EncodeAttestation(attestation)
	Nil(u.T(), err)
	decodedAttestation, err := types.DecodeAttestation(encodedAttestation)
	Nil(u.T(), err)

	Equal(u.T(), attestation.Nonce(), decodedAttestation.Nonce())
	Equal(u.T(), attestation.Timestamp().Uint64(), decodedAttestation.Timestamp().Uint64())
	Equal(u.T(), attestation.BlockNumber().Uint64(), decodedAttestation.BlockNumber().Uint64())

	for i := 0; i < 32; i++ {
		Equal(u.T(), attestation.DataHash()[i], decodedAttestation.DataHash()[i])
		Equal(u.T(), attestation.SnapshotRoot()[i], decodedAttestation.DataHash()[i])
	}

	attestationSignature, _, _, err := attestation.SignAttestation(u.GetTestContext(), u.NotaryBondedSigner)
	Nil(u.T(), err)

	encodedSignature, err := types.EncodeSignature(attestationSignature)
	Nil(u.T(), err)

	decodedSignature, err := types.DecodeSignature(encodedSignature)

	Equal(u.T(), attestationSignature.R().Uint64(), decodedSignature.R().Uint64())
	Equal(u.T(), attestationSignature.V().Uint64(), decodedSignature.V().Uint64())
	Equal(u.T(), attestationSignature.S().Uint64(), decodedSignature.S().Uint64())

	encodedSignature2, err := types.EncodeSignature(decodedSignature)
	Nil(u.T(), err)

	Equal(u.T(), len(encodedSignature), len(encodedSignature2))

	for i := 0; i < len(encodedSignature); i++ {
		Equal(u.T(), encodedSignature[i], encodedSignature2[i])
	}
}
