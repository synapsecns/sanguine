package evm_test

import (
	"github.com/synapsecns/sanguine/core/merkle"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
)

func (i ContractSuite) TestDestinationSubmitAttestation() {
	originDomain := uint32(i.TestBackendOrigin.GetChainID())
	destinationDomain := uint32(i.TestBackendDestination.GetChainID())

	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	attestKey := types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	Nil(i.T(), err)

	notarySignature, err := i.NotarySigner.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(i.T(), err)
	guardSignature, err := i.GuardSigner.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(i.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(i.T(), err)

	auth := i.TestBackendDestination.GetTxContext(i.GetTestContext(), nil)

	tx, err := i.DestinationContract.SubmitAttestation(auth.TransactOpts, rawSignedAttestation)
	Nil(i.T(), err)

	i.TestBackendDestination.WaitForConfirmation(i.GetTestContext(), tx)
}

func (i ContractSuite) TestDestinationExecute() {
	tree := merkle.NewTree()
	originDomain := uint32(i.TestBackendOrigin.GetChainID())
	destinationDomain := uint32(i.TestBackendDestination.GetChainID())

	recipient := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	optimisticSeconds := uint32(0)
	tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
	encodedTips, err := types.EncodeTips(tips)
	Nil(i.T(), err)
	messageBody := []byte("This is a test message")

	txContextOrigin := i.TestBackendOrigin.GetTxContext(i.GetTestContext(), i.OriginContractMetadata.OwnerPtr())

	tx, err := i.OriginContract.Dispatch(txContextOrigin.TransactOpts, destinationDomain, recipient.Hash(), optimisticSeconds, encodedTips, messageBody)

	sender, err := i.TestBackendOrigin.Signer().Sender(tx)
	Nil(i.T(), err)

	header := types.NewHeader(originDomain, sender.Hash(), 1, destinationDomain, recipient.Hash(), optimisticSeconds)

	message := types.NewMessage(header, tips, messageBody)

	leaf, err := message.ToLeaf()
	Nil(i.T(), err)

	tree.Insert(leaf[:])

	root, err := tree.Root(1)
	Nil(i.T(), err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	nonce := gofakeit.Uint32()
	attestKey := types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), rootB32)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	Nil(i.T(), err)

	notarySignature, err := i.NotarySigner.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(i.T(), err)
	guardSignature, err := i.GuardSigner.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(i.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(i.T(), err)

	auth := i.TestBackendDestination.GetTxContext(i.GetTestContext(), nil)

	tx, err = i.DestinationContract.SubmitAttestation(auth.TransactOpts, rawSignedAttestation)
	Nil(i.T(), err)

	i.TestBackendDestination.WaitForConfirmation(i.GetTestContext(), tx)

	proof, err := tree.MerkleProof(0, 1)
	Nil(i.T(), err)

	var proofB32B32 [32][32]byte
	for i, p := range proof {
		copy(proofB32B32[i][:], p)
	}

	encodedMessage, err := types.EncodeMessage(message)
	Nil(i.T(), err)

	tx, err = i.DestinationContract.Execute(auth.TransactOpts, encodedMessage, proofB32B32, big.NewInt(0))
	Nil(i.T(), err)

	i.TestBackendDestination.WaitForConfirmation(i.GetTestContext(), tx)
}
