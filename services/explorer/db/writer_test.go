package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"math/big"
)

func (t *DBSuite) TestBridgeWrite() {
	bridgeEvent := bridge.SynapseBridgeTokenDeposit{
		To:      common.BigToAddress(big.NewInt(gofakeit.Int64())),
		ChainId: big.NewInt(int64(gofakeit.Uint64())),
		Token:   common.BigToAddress(big.NewInt(gofakeit.Int64())),
		Amount:  big.NewInt(int64(gofakeit.Uint64())),
		Raw: ethTypes.Log{
			Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: gofakeit.Uint64(),
			TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
			TxIndex:     uint(gofakeit.Uint64()),
			BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
			Index:       uint(gofakeit.Uint64()),
			Removed:     false,
		},
	}
	chainID := gofakeit.Uint32()
	testTokenID := "testid"
	err := t.db.StoreEvent(t.GetTestContext(), bridgeEvent, nil, chainID, &testTokenID)
	Nil(t.T(), err)
	blockNumber, err := t.db.ReadBlockNumberByChainID(t.GetTestContext(), 0, chainID)
	Nil(t.T(), err)
	Equal(t.T(), *blockNumber, bridgeEvent.Raw.BlockNumber, "Returned data from bridge write incorrect.")
	t.cleanup()
}

func (t *DBSuite) TestSwapWrite() {
	swapEvent := swap.SwapFlashLoanTokenSwap{
		Buyer:        common.BigToAddress(big.NewInt(gofakeit.Int64())),
		TokensSold:   big.NewInt(int64(gofakeit.Uint64())),
		TokensBought: big.NewInt(int64(gofakeit.Uint64())),
		SoldId:       big.NewInt(int64(gofakeit.Uint64())),
		BoughtId:     big.NewInt(int64(gofakeit.Uint64())),
		Raw: ethTypes.Log{
			Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: gofakeit.Uint64(),
			TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
			TxIndex:     uint(gofakeit.Uint64()),
			BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
			Index:       uint(gofakeit.Uint64()),
			Removed:     false,
		},
	}
	chainID := gofakeit.Uint32()
	err := t.db.StoreEvent(t.GetTestContext(), nil, swapEvent, chainID, nil)
	Nil(t.T(), err)
	blockNumber, err := t.db.ReadBlockNumberByChainID(t.GetTestContext(), 1, chainID)
	Nil(t.T(), err)
	Equal(t.T(), *blockNumber, swapEvent.Raw.BlockNumber, "Returned data from swap write incorrect.")
	Nil(t.T(), err)
	t.cleanup()
}
