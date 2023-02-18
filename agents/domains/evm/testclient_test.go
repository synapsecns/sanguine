package evm_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
)

func (i ContractSuite) TestClientSendMessage() {
	destinationDomain := uint32(i.TestBackendDestination.GetChainID())

	auth := i.TestBackendDestination.GetTxContext(i.GetTestContext(), nil)

	recipient := i.TestClientMetadataOnDestination.Address()
	optimisticSeconds := uint32(10)
	message := []byte{byte(gofakeit.Uint32())}
	tx, err := i.TestClientOnOrigin.SendMessage(auth.TransactOpts, destinationDomain, recipient, optimisticSeconds, message)
	Nil(i.T(), err)

	i.TestBackendDestination.WaitForConfirmation(i.GetTestContext(), tx)
}
