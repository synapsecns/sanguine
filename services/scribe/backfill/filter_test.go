package backfill_test

import (
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/chain/client/mocks"
	etherMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"

	"github.com/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
)

// TestFilterLogsMaxAttempts ensures after the maximum number of attempts, an error is returned.
func (b BackfillSuite) TestFilterLogsMaxAttempts() {
	b.T().Skip("flake")
	chainID := big.NewInt(int64(1))
	simulatedChain := geth.NewEmbeddedBackendForChainID(b.GetTestContext(), b.T(), chainID)
	simulatedClient, err := backfill.DialBackend(b.GetTestContext(), simulatedChain.RPCAddress(), b.metrics)
	Nil(b.T(), err)
	mockFilterer := new(mocks.EVMClient)
	contractAddress := etherMocks.MockAddress()
	rangeFilter := backfill.NewRangeFilter(contractAddress, simulatedClient, big.NewInt(1), big.NewInt(10), 1, true, 1, uint32(chainID.Uint64()))

	// Use the range filterer created above to create a mock log filter.
	mockFilterer.
		On("FilterLogs", mock.Anything, mock.Anything).
		Return(nil, errors.New("I'm a test error"))

	logInfo, err := rangeFilter.FilterLogs(b.GetTestContext(), &util.Chunk{
		StartBlock: big.NewInt(1),
		EndBlock:   big.NewInt(10),
	})
	Nil(b.T(), logInfo)
	NotNil(b.T(), err)
}
