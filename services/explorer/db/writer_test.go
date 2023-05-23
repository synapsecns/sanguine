package db_test

import (
	"database/sql"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
)

func (t *DBSuite) TestBridgeWrite() {
	defer t.cleanup()
	bridgeEvent := &model.BridgeEvent{
		InsertTime:      gofakeit.Uint64(),
		ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		ChainID:         gofakeit.Uint32(),
		EventType:       bridgeTypes.DepositEvent.Int(),
		BlockNumber:     gofakeit.Uint64(),
		TxHash:          common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),

		Token:              common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		Amount:             big.NewInt(gofakeit.Int64()),
		Recipient:          sql.NullString{String: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(), Valid: true},
		DestinationChainID: big.NewInt(gofakeit.Int64()),
	}
	err := t.db.StoreEvent(t.GetTestContext(), bridgeEvent)
	Nil(t.T(), err)
}

func (t *DBSuite) TestSwapWrite() {
	defer t.cleanup()
	swapEvent := &model.SwapEvent{
		InsertTime:      gofakeit.Uint64(),
		ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		ChainID:         gofakeit.Uint32(),
		EventType:       bridgeTypes.DepositEvent.Int(),
		BlockNumber:     gofakeit.Uint64(),
		TxHash:          common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),

		Buyer:        sql.NullString{String: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(), Valid: true},
		TokensSold:   big.NewInt(gofakeit.Int64()),
		TokensBought: big.NewInt(gofakeit.Int64()),
		SoldID:       big.NewInt(gofakeit.Int64()),
		BoughtID:     big.NewInt(gofakeit.Int64()),
	}
	err := t.db.StoreEvent(t.GetTestContext(), swapEvent)
	Nil(t.T(), err)
}

func (t *DBSuite) TestLastBlockWrite() {
	defer t.cleanup()
	chainID := gofakeit.Uint32()
	blockNumber := gofakeit.Uint64()
	contract := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	err := t.db.StoreLastBlock(t.GetTestContext(), chainID, blockNumber, contract)
	Nil(t.T(), err)
	blockNumber++
	err = t.db.StoreLastBlock(t.GetTestContext(), chainID, blockNumber, contract)
	Nil(t.T(), err)
	storedBlockNum, err := t.db.GetLastStoredBlock(t.GetTestContext(), chainID, contract)
	Nil(t.T(), err)
	Equal(t.T(), blockNumber, storedBlockNum)

	chainID2 := gofakeit.Uint32()
	blockNumber2 := gofakeit.Uint64()
	err = t.db.StoreLastBlock(t.GetTestContext(), chainID2, blockNumber2, contract)
	Nil(t.T(), err)
	blockNumber2--
	err = t.db.StoreLastBlock(t.GetTestContext(), chainID2, blockNumber2, contract)
	Nil(t.T(), err)
	storedBlockNum2, err := t.db.GetLastStoredBlock(t.GetTestContext(), chainID2, contract)
	Nil(t.T(), err)
	Equal(t.T(), blockNumber2+1, storedBlockNum2)

	storedBlockNumOg, err := t.db.GetLastStoredBlock(t.GetTestContext(), chainID, contract)
	Nil(t.T(), err)
	Equal(t.T(), blockNumber, storedBlockNumOg)
}
