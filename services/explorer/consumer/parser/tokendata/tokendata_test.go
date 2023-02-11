package tokendata_test

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/mocks"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"math/big"
)

func (t *TokenDataSuite) TestTokenDataRetrieve() {
	testService := new(mocks.Service)

	testTokenID := gofakeit.Word()
	testTokenDecimals := gofakeit.Uint8()

	// error once
	testService.On("GetTokenID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("im an error")).Once()
	// make sure we don't try to fetch more than once
	testService.On("GetTokenID", mock.Anything, mock.Anything, mock.Anything).Return(&testTokenID, nil).Once()

	testService.On("GetToken", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("im an error")).Once()
	testService.On("GetToken", mock.Anything, mock.Anything, mock.Anything).Return(&bridgeconfig.BridgeConfigV3Token{
		TokenDecimals: testTokenDecimals,
	}, nil).Once()

	tokenService, err := tokendata.NewTokenDataService(testService, nil)
	Nil(t.T(), err)

	tokenData, err := tokenService.GetTokenData(t.GetTestContext(), 1, common.BigToAddress(big.NewInt(1)))
	Nil(t.T(), err)

	Equal(t.T(), testTokenID, tokenData.TokenID())
	Equal(t.T(), testTokenDecimals, tokenData.Decimals())

	// if we try to bypass cache, this will break
	NotPanics(t.T(), func() {
		_, _ = tokenService.GetTokenData(t.GetTestContext(), 1, common.BigToAddress(big.NewInt(1)))
	})
}
