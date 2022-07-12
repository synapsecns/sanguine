package evm_test

import (
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/synapse-node/pkg/common"
	"github.com/synapsecns/synapse-node/pkg/evm/client/mocks"
	"github.com/synapsecns/synapse-node/testutils/utils"
	"golang.org/x/sync/errgroup"
	"math/big"
	"time"
)

func (e *RPCSuite) TestGetters() {
	name := "hi"

	_, xAppContract := e.deployManager.GetXAppConfig(e.GetTestContext(), e.testBackend)

	testCfg := config.DomainConfig{
		DomainID:          1,
		RPCUrl:            e.testBackend.RPCAddress(),
		XAppConfigAddress: xAppContract.Address().String(),
	}

	testEvm, err := evm.NewEVM(e.GetTestContext(), name, testCfg)
	Nil(e.T(), err)
	Equal(e.T(), testEvm.Config(), testCfg)
	Equal(e.T(), testEvm.Name(), name)

	// get latest block from rpc
	latestBlock, err := e.testBackend.BlockNumber(e.GetTestContext())
	Nil(e.T(), err)

	// make sure it's equal to the client backend
	domainBlock, err := testEvm.BlockNumber(e.GetTestContext())
	Nil(e.T(), err)

	Equal(e.T(), latestBlock, uint64(domainBlock))
}

func (e *RPCSuite) TestFilterLogsMaxAttempts() {
	evm.SetMaxBackoff(time.Duration(0))
	evm.SetMinBackoff(time.Duration(0))

	mockFilterer := new(mocks.EVMClient)
	contractAddress := utils.NewMockAddress()

	// create a range filterer so we can test the filter logs method
	rangeFilter := evm.NewRangeFilter(contractAddress, mockFilterer, big.NewInt(1), big.NewInt(10), 1, true)

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

type Dispatch struct {
	destinationDomain uint32
	recipientAddress  [32]byte
	messageBody       []byte
	optimisticSeconds uint32
}

// GenerateDispatch generates a mock dispatch for testing.
func GenerateDispatch() Dispatch {
	newAddress := utils.NewMockAddress().Bytes()
	var recipient [32]byte

	copy(recipient[:], newAddress)

	return Dispatch{
		destinationDomain: gofakeit.Uint32(),
		recipientAddress:  recipient,
		messageBody:       []byte(gofakeit.Paragraph(3, 2, 1, " ")),
		optimisticSeconds: gofakeit.Uint32(),
	}
}

// GenerateDispatches generates a slice of dispatches.
func GenerateDispatches(dispatchCount int) (arr []Dispatch) {
	for i := 0; i < dispatchCount; i++ {
		arr = append(arr, GenerateDispatch())
	}
	return arr
}

func (e *RPCSuite) TestFilterer() {
	deployHelper := testutil.NewDeployManager(e.T())

	deployedContract, handle := deployHelper.GetHome(e.GetTestContext(), e.testBackend)

	dispatches := GenerateDispatches(10)

	var lastTx *types.Transaction
	for _, dispatch := range dispatches {
		auth := e.testBackend.GetTxContext(e.GetTestContext(), nil)

		addedDispatch, err := handle.Dispatch(auth.TransactOpts, dispatch.destinationDomain, dispatch.recipientAddress, dispatch.optimisticSeconds, dispatch.messageBody)
		Nil(e.T(), err)

		e.testBackend.WaitForConfirmation(e.GetTestContext(), addedDispatch)
		lastTx = addedDispatch
	}

	receipt, err := e.testBackend.TransactionReceipt(e.GetTestContext(), lastTx.Hash())
	Nil(e.T(), err)

	rangeFilter := evm.NewRangeFilter(deployedContract.Address(), e.testBackend, big.NewInt(0), receipt.BlockNumber, 1, true)

	g, ctx := errgroup.WithContext(e.GetTestContext())
	g.Go(func() error {
		//nolint: wrapcheck
		return rangeFilter.Start(e.GetTestContext())
	})

	_ = ctx

	g.Go(func() error {
		logChan := rangeFilter.GetLogChan()

		for log := range logChan {
			// TODO: assert log
			fmt.Println(log)
			if rangeFilter.Done() {
				return nil
			}
		}
		return nil
	})

	Nil(e.T(), g.Wait())
}

func (e *RPCSuite) TestProduceUpdate() {
	e.T().Skip("TODO: add produce update test")
}
