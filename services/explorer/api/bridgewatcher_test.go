package api_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
)

//	func (g APISuite) TestExistingOriginTx() {
//		chainID := uint32(1)
//
//		contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//
//		address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
//		tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
//
//		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&MvBridgeEvent{
//			TChainID:         chainID,
//			TContractAddress: contractAddress,
//			TEventType:       1,
//			TBlockNumber:     1,
//			TEventIndex:      gofakeit.Uint64(),
//			TTxHash:          txHash.String(),
//
//			TRecipient:          gosql.NullString{String: address.String(), Valid: true},
//			TDestinationChainID: big.NewInt(int64(2)),
//			TToken:              tokenAddr,
//			TSender:             tokenAddr,
//			TInsertTime:         1,
//
//			FChainID:         chainID,
//			FContractAddress: contractAddress,
//			FEventType:       1,
//			FBlockNumber:     1,
//			FEventIndex:      gofakeit.Uint64(),
//			FTxHash:          txHash.String(),
//
//			FInsertTime:         1,
//			FRecipient:          gosql.NullString{String: address.String(), Valid: true},
//			FDestinationChainID: big.NewInt(int64(2)),
//			FToken:              tokenAddr,
//			FSender:             tokenAddr,
//		})
//		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
//			ChainID:         chainID,
//			TokenAddress:    tokenAddr,
//			ContractAddress: contractAddress,
//			TokenIndex:      1,
//		})
//
//		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, 1)
//		Nil(g.T(), err)
//
//		chainIDInt := int(chainID)
//		txHashStr := txHash.String()
//		bridgeType := model.BridgeTypeBridge
//		result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), chainIDInt, txHashStr, bridgeType)
//		Nil(g.T(), err)
//		NotNil(g.T(), result)
//		Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)
//	}
//
// // nolint:gosec
//
//	func (g APISuite) TestNonExistingOriginTx() {
//		// Testing this tx: https://arbiscan.io/tx/0xa890211029aed050d94b9c1fb9c9864d68067d59a26194bdd04c1410d3e925ec
//		txHash := "0xa890211029aed050d94b9c1fb9c9864d68067d59a26194bdd04c1410d3e925ec"
//		chainID := 42161
//		bridgeType := model.BridgeTypeBridge
//		arbAddr := "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"
//		inputAmount := "277000000000000000"
//		swapContract := "0xa067668661C84476aFcDc6fA5D758C4c01C34352"
//		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
//			ChainID:         uint32(chainID),
//			TokenAddress:    arbAddr,
//			TokenIndex:      1,
//			ContractAddress: swapContract,
//		})
//		result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), chainID, txHash, bridgeType)
//		Nil(g.T(), err)
//		NotNil(g.T(), result)
//		Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
//
//		// check if data from swap logs were collected
//		Equal(g.T(), arbAddr, *result.Response.BridgeTx.TokenAddress)
//		Equal(g.T(), inputAmount, *result.Response.BridgeTx.Value)
//	}
//
// // nolint:gosec
//
//	func (g APISuite) TestNonExistingCCTPOriginTx() {
//		// Testing this tx: https://etherscan.io/tx/0x23392252f6afc660169bad0101d4c4b3bb9be8c7cca146dd1a7a9ce08f2281be
//		txHash := "0x23392252f6afc660169bad0101d4c4b3bb9be8c7cca146dd1a7a9ce08f2281be"
//		value := "976246870"
//		token := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
//		kappa := "336e45f3bae1d1477f219ae2a0c77ad2e84eba2d8da5859603a1759b9d9e536f"
//		chainID := 1
//		bridgeType := model.BridgeTypeCctp
//
//		result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), chainID, txHash, bridgeType)
//		Nil(g.T(), err)
//		NotNil(g.T(), result)
//		Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
//		Equal(g.T(), value, *result.Response.BridgeTx.Value)
//		Equal(g.T(), token, *result.Response.BridgeTx.TokenAddress)
//		Equal(g.T(), kappa, *result.Response.Kappa)
//	}
//
//	func (g APISuite) TestExistingDestinationTx() {
//		chainID := uint32(1)
//
//		contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//		bridgeType := model.BridgeTypeBridge
//
//		address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
//		tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
//		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
//		kappa := "kappa"
//		kappaSQL := gosql.NullString{String: kappa, Valid: true}
//		timestamp := uint64(1)
//		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&MvBridgeEvent{
//			TChainID:         chainID,
//			TContractAddress: contractAddress,
//			TEventType:       1,
//			TBlockNumber:     1,
//			TEventIndex:      gofakeit.Uint64(),
//			TTxHash:          txHash.String(),
//
//			TKappa:              kappaSQL,
//			TRecipient:          gosql.NullString{String: address.String(), Valid: true},
//			TDestinationChainID: big.NewInt(int64(2)),
//			TToken:              tokenAddr,
//			TSender:             tokenAddr,
//			TInsertTime:         1,
//
//			FChainID:         chainID,
//			FContractAddress: contractAddress,
//			FEventType:       1,
//			FBlockNumber:     1,
//			FEventIndex:      gofakeit.Uint64(),
//			FTxHash:          txHash.String(),
//
//			FInsertTime:         1,
//			FRecipient:          gosql.NullString{String: address.String(), Valid: true},
//			FDestinationChainID: big.NewInt(int64(2)),
//			FToken:              tokenAddr,
//			FSender:             tokenAddr,
//		})
//		var t []sql.HybridBridgeEvent
//		test := g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Raw("SELECT * FROM mv_bridge_events").Scan(&t)
//		fmt.Println("HOO", len(t), t[0].TKappa, t[0].TTxHash, test)
//		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
//			ChainID:         chainID,
//			TokenAddress:    tokenAddr,
//			ContractAddress: contractAddress,
//			TokenIndex:      1,
//		})
//
//		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, 1)
//		Nil(g.T(), err)
//
//		timestampInt := int(timestamp)
//		historical := false
//
//		result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), int(chainID), kappa, contractAddress, timestampInt, bridgeType, &historical)
//		Nil(g.T(), err)
//		NotNil(g.T(), result)
//		Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)
//	}
//
// // nolint:gosec
//
//	func (g APISuite) TestNonExistingDestinationTx() {
//		// Testing this tx: https://optimistic.etherscan.io/tx/0x7021a6046a39b3f5bd8956b83e0f6aa2b59c316e180e7fc41425d463cda35ae6
//		txHash := "0x7021a6046a39b3f5bd8956b83e0f6aa2b59c316e180e7fc41425d463cda35ae6"
//		kappa := "23C54D703DEA0451B74B40FFD22E1C1CA5A9F90CEF48BC322182491A386501AF"
//		address := "0x2d5a17539943a8c1a753578af3b4f91c9eb85eb9"
//		timestamp := 1692378548
//
//		chainID := 10
//		bridgeType := model.BridgeTypeBridge
//		historical := true // set to false if this tx is within the last hour or so
//		result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), chainID, kappa, address, timestamp, bridgeType, &historical)
//		Nil(g.T(), err)
//		NotNil(g.T(), result)
//		Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
//	}
//
// nolint:gosec
func (g APISuite) TestNonExistingDestinationTxHistorical() {
	// Testing this tx: https://optimistic.etherscan.io/tx/0x7021a6046a39b3f5bd8956b83e0f6aa2b59c316e180e7fc41425d463cda35ae6
	txHash := "0x7021a6046a39b3f5bd8956b83e0f6aa2b59c316e180e7fc41425d463cda35ae6"
	kappa := "23C54D703DEA0451B74B40FFD22E1C1CA5A9F90CEF48BC322182491A386501AF"
	address := "0x2d5a17539943a8c1a753578af3b4f91c9eb85eb9"
	timestamp := 1692378548

	chainID := 10
	bridgeType := model.BridgeTypeBridge
	historical := true
	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), chainID, kappa, address, timestamp, bridgeType, &historical)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
}

//// nolint:gosec
//func (g APISuite) TestNonExistingDestinationTxCCTP() {
//	// Testing this tx: https://etherscan.io/tx/0xc0fc8fc8b13856ede8862439c2ac9705005a1c7f2610f52446ae7c3f9d52d360
//	txHash := "0xc0fc8fc8b13856ede8862439c2ac9705005a1c7f2610f52446ae7c3f9d52d360"
//	kappa := "1d41f047267fdaf805234d76c998bd0fa63558329c455f2419d81fa26167214d"
//	address := "0xfE332ab9f3a0F4424c8Cb03b621120319E7b5f53"
//	timestamp := 1692110880
//	value := "3699210873"
//	chainID := 1
//	bridgeType := model.BridgeTypeCctp
//	historical := false
//	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), chainID, kappa, address, timestamp, bridgeType, &historical)
//	Nil(g.T(), err)
//	NotNil(g.T(), result)
//	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
//	Equal(g.T(), value, *result.Response.BridgeTx.Value)
//}

// nolint:gosec
func (g APISuite) TestNonExistingOriginTxOP() {
	// Testing this tx: https://optimistic.etherscan.io/tx/0x76263eb49042e6e5ff161b55d777eab6ba4f94fba8be8fafc3c950b0848ddebe
	txHash := "0x76263eb49042e6e5ff161b55d777eab6ba4f94fba8be8fafc3c950b0848ddebe"
	chainID := 10
	bridgeType := model.BridgeTypeBridge
	tokenAddr := "0x7F5c764cBc14f9669B88837ca1490cCa17c31607"
	inputAmount := "2000000"
	swapContract := "0xF44938b0125A6662f9536281aD2CD6c499F22004"
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         uint32(chainID),
		TokenAddress:    tokenAddr,
		TokenIndex:      1,
		ContractAddress: swapContract,
	})
	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), chainID, txHash, bridgeType)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)

	// check if data from swap logs were collected
	Equal(g.T(), tokenAddr, *result.Response.BridgeTx.TokenAddress)
	Equal(g.T(), inputAmount, *result.Response.BridgeTx.Value)
}
