package pingpongclient_test

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
)

func (h PingPongClientSuite) TestPing() {
	pingSentSink := make(chan *pingpongclient.PingPongClientPingSent)
	pingSentSub, err := h.pingPongClientContract.WatchPingSent(&bind.WatchOpts{Context: h.GetTestContext()}, pingSentSink)
	Nil(h.T(), err)

	txContextPingPongClientOrigin := h.testBackend.GetTxContext(h.GetTestContext(), h.pingPongClientMetadata.OwnerPtr())

	recipient := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	pingPongClientOnOriginTx, err := h.pingPongClientContract.DoPing(txContextPingPongClientOrigin.TransactOpts, h.destinationID, recipient, uint16(1))
	h.Nil(err)
	h.testBackend.WaitForConfirmation(h.GetTestContext(), pingPongClientOnOriginTx)

	watchCtx, cancel := context.WithTimeout(h.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		h.T().Error(h.T(), fmt.Errorf("test context completed %w", h.GetTestContext().Err()))
	case <-pingSentSub.Err():
		h.T().Error(h.T(), pingSentSub.Err())
	// get sent event
	case item := <-pingSentSink:
		h.NotNil(item)
		break
	}
}
