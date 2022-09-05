package backfill_test

import (
	"math/big"

	"github.com/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/synapse-node/pkg/common"
	"github.com/synapsecns/synapse-node/pkg/evm/client/mocks"
	"github.com/synapsecns/synapse-node/testutils/utils"
)

// TestFilterLogsMaxAttempts ensures after the maximum number of attempts, an error is returned.
func (b BackfillSuite) TestFilterLogsMaxAttempts() {
	mockFilterer := new(mocks.EVMClient)
	contractAddress := utils.NewMockAddress()

	// create a range filterer so we can test the filter logs method
	rangeFilter := backfill.NewRangeFilter(contractAddress, mockFilterer, big.NewInt(1), big.NewInt(10), 1, true)

	mockFilterer.
		// on a filter logs call
		On("FilterLogs", mock.Anything, mock.Anything).
		// return an error
		Return(nil, errors.New("I'm a test error"))

	logInfo, err := rangeFilter.FilterLogs(b.GetTestContext(), &common.Chunk{
		StartBlock: big.NewInt(1),
		EndBlock:   big.NewInt(10),
	})

	Nil(b.T(), logInfo)
	NotNil(b.T(), err)
}
