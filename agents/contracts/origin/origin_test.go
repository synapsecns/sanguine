package origin_test

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
)

func (h OriginSuite) TestLocalDomain() {
	localDomain, err := h.originContract.LocalDomain(&bind.CallOpts{Context: h.GetTestContext()})
	Nil(h.T(), err)
	Equal(h.T(), uint32(h.testBackend.GetBigChainID().Uint64()), localDomain)
}

func (h OriginSuite) TestSentTopic() {
	// init the sent event
	txContext := h.testBackend.GetTxContext(h.GetTestContext(), nil)

	sentSink := make(chan *originharness.OriginHarnessSent)
	sub, err := h.originContract.WatchSent(&bind.WatchOpts{Context: h.GetTestContext()}, sentSink, [][32]byte{}, []uint32{}, []uint32{})
	Nil(h.T(), err)

	paddedRequest := big.NewInt(0)
	tx, err := h.originContract.SendBaseMessage(txContext.TransactOpts, h.destinationID, [32]byte{}, 1, paddedRequest, []byte{})
	Nil(h.T(), err)

	h.testBackend.WaitForConfirmation(h.GetTestContext(), tx)

	watchCtx, cancel := context.WithTimeout(h.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		h.T().Error(h.T(), fmt.Errorf("test context completed %w", h.GetTestContext().Err()))
	case <-sub.Err():
		h.T().Error(h.T(), sub.Err())
	// get sent event
	case item := <-sentSink:
		parser, err := origin.NewParser(h.originContract.Address())
		Nil(h.T(), err)

		eventType, ok := parser.EventType(item.Raw)
		True(h.T(), ok)
		Equal(h.T(), eventType, origin.SentEvent)

		break
	}
}

func (h OriginSuite) TestUpdateTopic() {
	h.T().Skip("TODO: test this. Mocker should be able to mock this out")
}
