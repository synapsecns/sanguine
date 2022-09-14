package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"math/big"
)

func (t *DBSuite) TestBridgeWrite() {
	deposit := bridge.SynapseBridgeTokenDeposit{
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
	err := t.db.StoreEvent(t.GetTestContext(), deposit, gofakeit.Uint32())
	Nil(t.T(), err)
}
