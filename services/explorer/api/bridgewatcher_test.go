package api_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
)

// 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38

// func (g APISuite) TestExistingOriginTx() {
//	chainID := uint32(1)
//
//	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//
//	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
//	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//	txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
//
//	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
//		InsertTime:         1,
//		ChainID:            chainID,
//		Recipient:          gosql.NullString{String: address.String(), Valid: true},
//		DestinationChainID: big.NewInt(int64(2)),
//		BlockNumber:        1,
//		TxHash:             txHash.String(),
//		EventIndex:         gofakeit.Uint64(),
//		Token:              tokenAddr,
//		Sender:             tokenAddr,
//	})
//	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
//		ChainID:         chainID,
//		TokenAddress:    tokenAddr,
//		ContractAddress: contractAddress,
//		TokenIndex:      1,
//	})
//
//	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, 1)
//	Nil(g.T(), err)
//
//	chainIDInt := int(chainID)
//	txHashStr := txHash.String()
//	bridgeType := model.BridgeTypeBridge
//	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), &chainIDInt, &txHashStr, &bridgeType)
//	Nil(g.T(), err)
//	NotNil(g.T(), result)
//	Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)
//
//}

// func (g APISuite) TestNonExistingOriginTx() {
//	// Testing this tx: https://bscscan.com/tx/0x0478fa7e15d61498ed00bdde6254368df416bbc66a11a2aed88f4ce2983b5470
//	txHash := "0x0478fa7e15d61498ed00bdde6254368df416bbc66a11a2aed88f4ce2983b5470"
//	chainID := 56
//	bridgeType := model.BridgeTypeBridge
//	bscusdAddr := "0x55d398326f99059fF775485246999027B3197955"
//	inputAmount := "7500003889000000000000"
//	swapContract := "0x28ec0B36F0819ecB5005cAB836F4ED5a2eCa4D13"
//	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
//		ChainID:         uint32(chainID),
//		TokenAddress:    bscusdAddr,
//		TokenIndex:      3,
//		ContractAddress: swapContract,
//	})
//	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), &chainID, &txHash, &bridgeType)
//	Nil(g.T(), err)
//	NotNil(g.T(), result)
//	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
//
//	// check if data from swap logs were collected
//	Equal(g.T(), bscusdAddr, *result.Response.BridgeTx.TokenAddress)
//	Equal(g.T(), inputAmount, *result.Response.BridgeTx.Value)
//
//}

//
// func (g APISuite) TestExistingDestinationTx() {
//	chainID := uint32(1)
//
//	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//	bridgeType := model.BridgeTypeBridge
//
//	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
//	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//	txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
//	kappa := "kappa"
//	kappaSql := gosql.NullString{String: kappa, Valid: true}
//	timestamp := uint64(1)
//	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
//		InsertTime:         1,
//		ChainID:            chainID,
//		Recipient:          gosql.NullString{String: address.String(), Valid: true},
//		DestinationChainID: big.NewInt(int64(2)),
//		BlockNumber:        1,
//		TxHash:             txHash.String(),
//		EventIndex:         gofakeit.Uint64(),
//		ContractAddress:    contractAddress,
//		Token:              tokenAddr,
//		Sender:             tokenAddr,
//		Kappa:              kappaSql,
//		TimeStamp:          &timestamp,
//	})
//	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
//		ChainID:         chainID,
//		TokenAddress:    tokenAddr,
//		ContractAddress: contractAddress,
//		TokenIndex:      1,
//	})
//
//	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, 1)
//	Nil(g.T(), err)
//
//	chainIDInt := int(chainID)
//	timestampInt := int(timestamp)
//	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), &chainIDInt, &kappa, &contractAddress, &timestampInt, &bridgeType)
//	Nil(g.T(), err)
//	NotNil(g.T(), result)
//	Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)
//
//}

// func (g APISuite) TestNonExistingDestinationTx() {
//	// Testing this tx: https://bscscan.com/tx/0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67
//	txHash := "0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67"
//	kappa := "e16367a638236d4c1e942aba379fcc9babf468b76908253cc7797ed2df691e57"
//	address := "0x76160a62E9142552c4a1eeAe935Ed5cd3001f7fd"
//	timestamp := 1692099540
//
//	chainID := 56
//	bridgeType := model.BridgeTypeBridge
//	historical := false
//	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), &chainID, &kappa, &address, &timestamp, &bridgeType, &historical)
//	Nil(g.T(), err)
//	NotNil(g.T(), result)
//	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
//}

func (g APISuite) TestNonExistingDestinationTxHistorical() {
	// Testing this tx: https://bscscan.com/tx/0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67
	txHash := "0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67"
	kappa := "e16367a638236d4c1e942aba379fcc9babf468b76908253cc7797ed2df691e57"
	address := "0x76160a62E9142552c4a1eeAe935Ed5cd3001f7fd"
	timestamp := 1692099540

	chainID := 56
	bridgeType := model.BridgeTypeBridge
	historical := true
	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), &chainID, &kappa, &address, &timestamp, &bridgeType, &historical)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
}
