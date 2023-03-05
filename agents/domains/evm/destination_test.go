package evm_test

func (i ContractSuite) TestDestinationSubmitAttestation() {
	// TODO (joeallen): FIX ME
	i.T().Skip()
	// TODO (joeallen): FIX ME
	// originDomain := uint32(i.TestBackendOrigin.GetChainID())
	// destinationDomain := uint32(i.TestBackendDestination.GetChainID())

	// nonce := gofakeit.Uint32()
	// root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	// attestKey := types.AttestationKey{
	//	Origin:      originDomain,
	//	Destination: destinationDomain,
	//	Nonce:       nonce,
	//}
	// unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	// hashedAttestation, err := types.Hash(unsignedAttestation)
	// Nil(i.T(), err)

	// notarySignature, err := i.NotaryBondedSigner.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	// Nil(i.T(), err)
	// guardSignature, err := i.GuardBondedSigner.SignMessage(i.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	// Nil(i.T(), err)

	// signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	// rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	// Nil(i.T(), err)

	// auth := i.TestBackendDestination.GetTxContext(i.GetTestContext(), nil)

	// tx, err := i.DestinationContract.SubmitAttestation(auth.TransactOpts, rawSignedAttestation)
	// Nil(i.T(), err)

	// i.TestBackendDestination.WaitForConfirmation(i.GetTestContext(), tx)
}
