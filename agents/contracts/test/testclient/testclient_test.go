package testclient_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
)

func (h TestClientSuite) TestSendMessage() {
	messageSentSink := make(chan *testclient.TestClientMessageSent)
	sub, err := h.testClientContract.WatchMessageSent(&bind.WatchOpts{Context: h.GetTestContext()}, messageSentSink)
	Nil(h.T(), err)

	txContextTestClientOrigin := h.testBackend.GetTxContext(h.GetTestContext(), h.testClientMetadata.OwnerPtr())

	optimisticSeconds := uint32(10)
	recipient := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	body := []byte{byte(gofakeit.Uint32())}
	testClientOnOriginTx, err := h.testClientContract.SendMessage(txContextTestClientOrigin.TransactOpts, h.destinationID, recipient, optimisticSeconds, body)
	h.Nil(err)
	h.testBackend.WaitForConfirmation(h.GetTestContext(), testClientOnOriginTx)

	sender := h.testClientContract.Address()

	fmt.Printf("\nCRONIN from getting sender\n dest = %d, nonce = %d, sender = %s, recipient = %s, message = %s \nCRONIN\n",
		h.destinationID,
		1,
		hex.EncodeToString(sender[:]),
		hex.EncodeToString(recipient[:]),
		hex.EncodeToString(body))

	watchCtx, cancel := context.WithTimeout(h.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		h.T().Error(h.T(), fmt.Errorf("test context completed %w", h.GetTestContext().Err()))
	case <-sub.Err():
		h.T().Error(h.T(), sub.Err())
	// get dispatch event
	case item := <-messageSentSink:
		h.NotNil(item)

		fmt.Printf("\nCRONIN from sink\n dest = %d, nonce = %d, sender = %s, recipient = %s, message = %s \nCRONIN\n",
			item.Destination,
			item.Nonce,
			hex.EncodeToString(item.Sender[:]),
			hex.EncodeToString(item.Recipient[:]),
			hex.EncodeToString(item.Message))

		break
	}
}
