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

	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.HybridBridgeEvent{
		FInsertTime:         1,
		FChainID:            chainID,
		FRecipient:          gosql.NullString{String: address.String(), Valid: true},
		FDestinationChainID: big.NewInt(int64(2)),
		FBlockNumber:        1,
		FTxHash:             txHash.String(),
		FEventIndex:         gofakeit.Uint64(),
		FToken:              tokenAddr,
		FSender:             tokenAddr,
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
	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), chainIDInt, txHashStr, bridgeType)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)
}

// nolint:gosec
func (g APISuite) TestNonExistingOriginTx() {
	// Testing this tx: https://bscscan.com/tx/0x0478fa7e15d61498ed00bdde6254368df416bbc66a11a2aed88f4ce2983b5470
	txHash := "0x0478fa7e15d61498ed00bdde6254368df416bbc66a11a2aed88f4ce2983b5470"
	chainID := 56
	bridgeType := model.BridgeTypeBridge
	bscusdAddr := "0x55d398326f99059fF775485246999027B3197955"
	inputAmount := "7500003889000000000000"
	swapContract := "0x28ec0B36F0819ecB5005cAB836F4ED5a2eCa4D13"
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         uint32(chainID),
		TokenAddress:    bscusdAddr,
		TokenIndex:      3,
		ContractAddress: swapContract,
	})
	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), chainID, txHash, bridgeType)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)

	// check if data from swap logs were collected
	Equal(g.T(), bscusdAddr, *result.Response.BridgeTx.TokenAddress)
	Equal(g.T(), inputAmount, *result.Response.BridgeTx.Value)
}

// nolint:gosec
func (g APISuite) TestNonExistingCCTPOriginTx() {
	// Testing this tx: https://etherscan.io/tx/0x23392252f6afc660169bad0101d4c4b3bb9be8c7cca146dd1a7a9ce08f2281be
	txHash := "0x23392252f6afc660169bad0101d4c4b3bb9be8c7cca146dd1a7a9ce08f2281be"
	value := "976246870"
	token := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	kappa := "336e45f3bae1d1477f219ae2a0c77ad2e84eba2d8da5859603a1759b9d9e536f"
	chainID := 1
	bridgeType := model.BridgeTypeCctp

	result, err := g.client.GetOriginBridgeTx(g.GetTestContext(), chainID, txHash, bridgeType)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
	Equal(g.T(), value, *result.Response.BridgeTx.Value)
	Equal(g.T(), token, *result.Response.BridgeTx.TokenAddress)
	Equal(g.T(), kappa, *result.Response.Kappa)
}

func (g APISuite) TestExistingDestinationTx() {
	chainID := uint32(1)

	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	bridgeType := model.BridgeTypeBridge

	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
	kappa := "kappa"
	kappaSQL := gosql.NullString{String: kappa, Valid: true}
	timestamp := uint64(1)
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.HybridBridgeEvent{
		TInsertTime:         1,
		TChainID:            chainID,
		TRecipient:          gosql.NullString{String: address.String(), Valid: true},
		TDestinationChainID: big.NewInt(int64(2)),
		TBlockNumber:        1,
		TTxHash:             txHash.String(),
		TEventIndex:         gofakeit.Uint64(),
		TContractAddress:    contractAddress,
		TToken:              tokenAddr,
		TSender:             tokenAddr,
		TKappa:              kappaSQL,
		TTimeStamp:          &timestamp,
	})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         chainID,
		TokenAddress:    tokenAddr,
		ContractAddress: contractAddress,
		TokenIndex:      1,
	})

	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, 1)
	Nil(g.T(), err)

	timestampInt := int(timestamp)
	historical := false

	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), int(chainID), kappa, contractAddress, timestampInt, bridgeType, &historical)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash.String(), *result.Response.BridgeTx.TxnHash)
}

// nolint:gosec
func (g APISuite) TestNonExistingDestinationTx() {
	// Testing this tx: https://bscscan.com/tx/0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67
	txHash := "0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67"
	kappa := "e16367a638236d4c1e942aba379fcc9babf468b76908253cc7797ed2df691e57"
	address := "0x76160a62E9142552c4a1eeAe935Ed5cd3001f7fd"
	timestamp := 1692099540

	chainID := 56
	bridgeType := model.BridgeTypeBridge
	historical := true // set to false if this tx is within the last hour or so
	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), chainID, kappa, address, timestamp, bridgeType, &historical)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
}

// nolint:gosec
func (g APISuite) TestNonExistingDestinationTxHistorical() {
	// Testing this tx: https://bscscan.com/tx/0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67
	txHash := "0xa8697dd51ffaa025c5a7449e1f70a8f0776e78bbc92993bae18bf4eb1be99f67"
	kappa := "e16367a638236d4c1e942aba379fcc9babf468b76908253cc7797ed2df691e57"
	address := "0x76160a62E9142552c4a1eeAe935Ed5cd3001f7fd"
	timestamp := 1692099540

	chainID := 56
	bridgeType := model.BridgeTypeBridge
	historical := true
	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), chainID, kappa, address, timestamp, bridgeType, &historical)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
}

// nolint:gosec
func (g APISuite) TestNonExistingDestinationTxCCTP() {
	// Testing this tx: https://etherscan.io/tx/0xc0fc8fc8b13856ede8862439c2ac9705005a1c7f2610f52446ae7c3f9d52d360
	txHash := "0xc0fc8fc8b13856ede8862439c2ac9705005a1c7f2610f52446ae7c3f9d52d360"
	kappa := "1d41f047267fdaf805234d76c998bd0fa63558329c455f2419d81fa26167214d"
	address := "0xfE332ab9f3a0F4424c8Cb03b621120319E7b5f53"
	timestamp := 1692110880
	value := "3699210873"
	chainID := 1
	bridgeType := model.BridgeTypeCctp
	historical := false
	result, err := g.client.GetDestinationBridgeTx(g.GetTestContext(), chainID, kappa, address, timestamp, bridgeType, &historical)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), txHash, *result.Response.BridgeTx.TxnHash)
	Equal(g.T(), value, *result.Response.BridgeTx.Value)
}
