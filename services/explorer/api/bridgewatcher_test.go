package api_test

import (
	gosql "database/sql"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math/big"
)

func (g APISuite) TestExistingOriginTx() {
	chainID := uint32(1)

	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()

	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))

	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
		InsertTime:         1,
		ChainID:            chainID,
		Recipient:          gosql.NullString{String: address.String(), Valid: true},
		DestinationChainID: big.NewInt(int64(2)),
		BlockNumber:        1,
		TxHash:             txHash.String(),
		EventIndex:         gofakeit.Uint64(),
		Token:              tokenAddr,
		Sender:             tokenAddr,
	})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         chainID,
		TokenAddress:    tokenAddr,
		ContractAddress: contractAddress,
		TokenIndex:      1,
	})

	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, 1)
	Nil(g.T(), err)

	chainIDInt := int(chainID)
	txHashStr := txHash.String()
	bridgeType := model.BridgeTypeBridge
	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), &chainIDInt, &txHashStr, &bridgeType)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)

}

func (g APISuite) TestNonExistingOriginTx() {
	// Testing this tx: https://bscscan.com/tx/0x85f314fce071bec4109f054895f002fad84358bdb0eca31495958872a7d970e9
	txHash := "0x85f314fce071bec4109f054895f002fad84358bdb0eca31495958872a7d970e9"
	chainID := 56
	bridgeType := model.BridgeTypeBridge

	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), &chainID, &txHash, &bridgeType)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
}

func (g APISuite) TestExistingDestinationTx() {
	chainID := uint32(1)

	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	bridgeType := model.BridgeTypeBridge

	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
	kappa := "kappa"
	kappaSql := gosql.NullString{String: kappa, Valid: true}
	timestamp := uint64(1)
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
		InsertTime:         1,
		ChainID:            chainID,
		Recipient:          gosql.NullString{String: address.String(), Valid: true},
		DestinationChainID: big.NewInt(int64(2)),
		BlockNumber:        1,
		TxHash:             txHash.String(),
		EventIndex:         gofakeit.Uint64(),
		ContractAddress:    contractAddress,
		Token:              tokenAddr,
		Sender:             tokenAddr,
		Kappa:              kappaSql,
		TimeStamp:          &timestamp,
	})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         chainID,
		TokenAddress:    tokenAddr,
		ContractAddress: contractAddress,
		TokenIndex:      1,
	})

	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, 1)
	Nil(g.T(), err)

	chainIDInt := int(chainID)
	timestampInt := int(timestamp)
	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), &chainIDInt, &kappa, &contractAddress, &timestampInt, &bridgeType)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)

}
