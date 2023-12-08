package swap_test

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/swap/mocks"
	"math/big"
)

func (s *SwapFetcherSuite) TestSwapMocks() {
	testService := new(mocks.ISwapFetcher)

	testChainID := uint32(gofakeit.Number(1, 1000))
	testTokenIndex := uint8(gofakeit.Number(0, 255))
	testContractAddress := common.HexToAddress(big.NewInt(gofakeit.Int64()).String())
	testTokenAddress := common.HexToAddress(big.NewInt(gofakeit.Int64()).String())

	testService.On("GetToken", mock.Anything, testContractAddress).Return(nil, nil, errors.New("error occurred")).Once()
	testService.On("GetToken", mock.Anything, testContractAddress).Return(nil, nil, nil).Once()
	testService.On("GetTokenAddress", mock.Anything, testTokenIndex).Return(nil, errors.New("error occurred")).Once()
	testService.On("GetTokenAddress", mock.Anything, testTokenIndex).Return(&testTokenAddress, nil).Once()
	testService.On("ChainID", mock.Anything, testChainID).Return(&testChainID, nil).Once()
}
