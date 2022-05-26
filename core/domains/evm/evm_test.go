package evm_test

import (
	"errors"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/synapse-node/pkg/common"
	"github.com/synapsecns/synapse-node/pkg/evm/client/mocks"
	"github.com/synapsecns/synapse-node/pkg/teller/backfiller"
	"github.com/synapsecns/synapse-node/testutils/utils"
	"math/big"
	"time"
)

func (e *EVMSuite) TestGetters() {
	name := "hi"
	testCfg := config.DomainConfig{
		DomainID: 1,
		RPCUrl:   e.testBackend.RPCAddress(),
	}

	testEvm, err := evm.NewEVM(e.GetTestContext(), name, testCfg, nil)
	Nil(e.T(), err)
	Equal(e.T(), testEvm.Config(), testCfg)
	Equal(e.T(), testEvm.Name(), name)
}

func (e *EVMSuite) TestFilterLogsMaxAttempts() {
	evm.SetMaxBackoff(time.Duration(0))
	evm.SetMinBackoff(time.Duration(0))

	mockFilterer := new(mocks.EVMClient)
	contractAddress := utils.NewMockAddress()

	// create a range filterer so we can test the filter logs method
	rangeFilter := backfiller.NewRangeFilter(contractAddress, mockFilterer, big.NewInt(1), big.NewInt(10), 1, true)

	mockFilterer.
		// on a filter logs call
		On("FilterLogs", mock.Anything, mock.Anything).
		// return an error
		Return(nil, errors.New("I'm a test error"))

	logInfo, err := rangeFilter.FilterLogs(e.GetTestContext(), &common.Chunk{
		StartBlock: big.NewInt(1),
		EndBlock:   big.NewInt(10),
	})

	Nil(e.T(), logInfo)
	NotNil(e.T(), err)
	mockFilterer.AssertNumberOfCalls(e.T(), "FilterLogs", evm.MaxAttempts)
}
