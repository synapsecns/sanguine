package evm_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
)

func (i ContractSuite) TestClientSendMessage() {
	destinationDomain := uint32(i.TestBackendDestination.GetChainID())

	auth := i.TestBackendOrigin.GetTxContext(i.GetTestContext(), nil)

	recipient := i.TestClientMetadataOnDestination.Address()
	optimisticSeconds := uint32(10)
	message := []byte{byte(gofakeit.Uint32())}
	gasLimit := uint64(10000000)
	version := uint32(1)
	tx, err := i.TestClientOnOrigin.SendMessage(auth.TransactOpts, destinationDomain, recipient, optimisticSeconds, gasLimit, version, message)
	Nil(i.T(), err)

	i.TestBackendOrigin.WaitForConfirmation(i.GetTestContext(), tx)
}
