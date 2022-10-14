package db_test

import (
	"database/sql"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	messageTypes "github.com/synapsecns/sanguine/services/explorer/types/message"
	"math/big"
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
	err := t.db.StoreEvent(t.GetTestContext(), bridgeEvent, nil, nil)
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
	err := t.db.StoreEvent(t.GetTestContext(), nil, swapEvent, nil)
	Nil(t.T(), err)
}

func (t *DBSuite) TestMessageWrite() {
	defer t.cleanup()
	messageEvent := &model.MessageEvent{
		InsertTime:      gofakeit.Uint64(),
		ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		ChainID:         gofakeit.Uint32(),
		EventType:       messageTypes.MessageSentEvent.Int(),
		BlockNumber:     gofakeit.Uint64(),
		TxHash:          common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		MessageID:       gofakeit.Sentence(10),
		SourceChainID:   big.NewInt(gofakeit.Int64()),

		SourceAddress:      sql.NullString{String: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(), Valid: true},
		DestinationChainID: big.NewInt(gofakeit.Int64()),
		Nonce:              big.NewInt(gofakeit.Int64()),
		Fee:                big.NewInt(gofakeit.Int64()),
		Options:            sql.NullString{String: gofakeit.Sentence(10), Valid: true},
		Message:            sql.NullString{String: gofakeit.Sentence(10), Valid: true},
		Receiver:           sql.NullString{String: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(), Valid: true},
	}
	err := t.db.StoreEvent(t.GetTestContext(), nil, nil, messageEvent)
	Nil(t.T(), err)
}
