package base_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/backends/base/mocks"
	"github.com/synapsecns/synapse-node/testutils/utils/eth"
	"time"
)

func (b *BaseSuite) TestWaitForConfirmation() {
	mockClient := new(mocks.ConfirmationClient)

	mockTx := eth.GetMockTxes(b.GetTestContext(), b.T(), 1, types.LegacyTxType)[0]

	const minConfirmTime = 60 * time.Millisecond
	confirmStart := time.Now()
	timer := time.NewTimer(minConfirmTime)

	mockClient.On("TransactionByHash", mock.Anything, mock.Anything).Once().Return(nil, true, nil)
	mockClient.On("TransactionByHash", mock.Anything, mock.Anything).WaitUntil(timer.C).Return(mockTx, false, nil)

	base.WaitForConfirmation(b.GetTestContext(), mockClient, mockTx, time.Millisecond*5)
	False(b.T(), time.Since(confirmStart) < minConfirmTime, "tx could not have been confirmed yet")
}
