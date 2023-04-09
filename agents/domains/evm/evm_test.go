package evm_test

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain/client/mocks"
	etherMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"golang.org/x/sync/errgroup"
)

func (e *RPCSuite) TestGetters() {
	originDomain := uint32(e.TestBackendOrigin.GetChainID())

	name := "hi"

	testCfg := config.DomainConfig{
		DomainID:      originDomain,
		RPCUrl:        e.TestBackendOrigin.RPCAddress(),
		OriginAddress: e.OriginContract.Address().String(),
	}

	testEvm, err := evm.NewEVM(e.GetTestContext(), name, testCfg)
	Nil(e.T(), err)
	Equal(e.T(), testEvm.Config(), testCfg)
	Equal(e.T(), testEvm.Name(), name)

	// get latest block from rpc
	latestBlock, err := e.TestBackendOrigin.BlockNumber(e.GetTestContext())
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
	contractAddress := etherMocks.MockAddress()

	// create a range filterer so we can test the filter logs method
	rangeFilter := evm.NewRangeFilter(contractAddress, mockFilterer, big.NewInt(1), big.NewInt(10), 1, true)

	mockFilterer.
		// on a filter logs call
		On("FilterLogs", mock.Anything, mock.Anything).
		// return an error
		Return(nil, errors.New("I'm a test error"))

	logInfo, err := rangeFilter.FilterLogs(e.GetTestContext(), &util.Chunk{
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
	newAddress := etherMocks.MockAddress().Bytes()
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
	if os.Getenv("CI") != "" {
		e.T().Skip("flakes on ci: since this will be replaced by scribe, we can deprecate this")
	}

	dispatches := GenerateDispatches(10)

	var lastTx *ethTypes.Transaction
	for _, dispatch := range dispatches {
		auth := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), nil)

		enodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
		Nil(e.T(), err)

		// TODO (joe): Figure out what to set for content
		addedDispatch, err := e.OriginContract.SendBaseMessage(auth.TransactOpts, dispatch.destinationDomain, dispatch.recipientAddress, dispatch.optimisticSeconds, enodedTips, dispatch.messageBody, []byte{})
		Nil(e.T(), err)

		e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), addedDispatch)
		lastTx = addedDispatch
	}

	receipt, err := e.TestBackendOrigin.TransactionReceipt(e.GetTestContext(), lastTx.Hash())
	Nil(e.T(), err)

	rangeFilter := evm.NewRangeFilter(e.OriginContract.Address(), e.TestBackendOrigin, big.NewInt(0), receipt.BlockNumber, 1, true)

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
