package api_test

import (
	gosql "database/sql"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/api"
	"github.com/synapsecns/sanguine/services/explorer/api/cache"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/graphql/client"
	"testing"

	"math/big"
)

func TestHandleJSONAmountStat(t *testing.T) {
	valueString := gofakeit.Word()
	// nolint
	valueStruct := gqlClient.GetAmountStatistic{
		Response: &struct {
			Value *string `json:"value" graphql:"value"`
		}{
			Value: &valueString,
		},
	}
	res := api.HandleJSONAmountStat(&valueStruct)
	NotNil(t, res)
	Equal(t, valueString, *res.Value)
}

func TestHandleJSONDailyStat(t *testing.T) {
	valueFloat := gofakeit.Float64()
	// nolint
	valueStruct := gqlClient.GetDailyStatisticsByChain{
		Response: []*struct {
			Date      *string  `json:"date" graphql:"date"`
			Ethereum  *float64 `json:"ethereum" graphql:"ethereum"`
			Optimism  *float64 `json:"optimism" graphql:"optimism"`
			Cronos    *float64 `json:"cronos" graphql:"cronos"`
			Bsc       *float64 `json:"bsc" graphql:"bsc"`
			Polygon   *float64 `json:"polygon" graphql:"polygon"`
			Fantom    *float64 `json:"fantom" graphql:"fantom"`
			Boba      *float64 `json:"boba" graphql:"boba"`
			Metis     *float64 `json:"metis" graphql:"metis"`
			Moonbeam  *float64 `json:"moonbeam" graphql:"moonbeam"`
			Moonriver *float64 `json:"moonriver" graphql:"moonriver"`
			Klaytn    *float64 `json:"klaytn" graphql:"klaytn"`
			Arbitrum  *float64 `json:"arbitrum" graphql:"arbitrum"`
			Avalanche *float64 `json:"avalanche" graphql:"avalanche"`
			Dfk       *float64 `json:"dfk" graphql:"dfk"`
			Aurora    *float64 `json:"aurora" graphql:"aurora"`
			Harmony   *float64 `json:"harmony" graphql:"harmony"`
			Canto     *float64 `json:"canto" graphql:"canto"`
			Dogechain *float64 `json:"dogechain" graphql:"dogechain"`
			Base      *float64 `json:"base" graphql:"base"`
			Total     *float64 `json:"total" graphql:"total"`
		}{
			{
				Total: &valueFloat,
			},
		},
	}
	res := api.HandleJSONDailyStat(&valueStruct)
	NotNil(t, res)
	Equal(t, valueFloat, *res[0].Total)
}

func (g APISuite) TestRehydrateCache() {
	responseCache, err := cache.NewAPICacheService()
	Nil(g.T(), err)
	chainID := g.chainIDs[0]
	chainID2 := g.chainIDs[1]
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	txHash := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	timestamp := uint64(1)
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	contractAddressSwap := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	tokenAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
		ChainID:            chainID,
		EventType:          gofakeit.Uint8(),
		DestinationChainID: big.NewInt(int64(chainID2)),
		Recipient:          gosql.NullString{String: address.String(), Valid: true},
		BlockNumber:        1,
		TxHash:             txHash.String(),
		EventIndex:         gofakeit.Uint64(),
		TimeStamp:          &timestamp,
		ContractAddress:    contractAddress,
		Token:              tokenAddress,
	})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         chainID,
		TokenAddress:    tokenAddress,
		ContractAddress: contractAddressSwap,
		TokenIndex:      1,
	})
	err = g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Table("mv_bridge_events").Set("gorm:table_options", "ENGINE=ReplacingMergeTree(insert_time) ORDER BY (fevent_index, fblock_number, fevent_type, ftx_hash, fchain_id, fcontract_address)").AutoMigrate(&MvBridgeEvent{})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&MvBridgeEvent{
		InsertTime:          1,
		FInsertTime:         0,
		FContractAddress:    "",
		FChainID:            0,
		FEventType:          0,
		FBlockNumber:        0,
		FTxHash:             "",
		FToken:              "",
		FAmount:             nil,
		FEventIndex:         0,
		FDestinationKappa:   "",
		FSender:             "",
		FRecipient:          gosql.NullString{},
		FRecipientBytes:     gosql.NullString{},
		FDestinationChainID: nil,
		FFee:                nil,
		FKappa:              gosql.NullString{},
		FTokenIndexFrom:     nil,
		FTokenIndexTo:       nil,
		FMinDy:              nil,
		FDeadline:           nil,
		FSwapSuccess:        nil,
		FSwapTokenIndex:     nil,
		FSwapMinAmount:      nil,
		FSwapDeadline:       nil,
		FTokenID:            gosql.NullString{},
		FAmountUSD:          nil,
		FFeeAmountUSD:       nil,
		FTokenDecimal:       nil,
		FTokenSymbol:        gosql.NullString{},
		FTimeStamp:          nil,
		TInsertTime:         0,
		TContractAddress:    "",
		TChainID:            0,
		TEventType:          0,
		TBlockNumber:        0,
		TTxHash:             "",
		TToken:              "",
		TAmount:             nil,
		TEventIndex:         0,
		TDestinationKappa:   "",
		TSender:             "",
		TRecipient:          gosql.NullString{},
		TRecipientBytes:     gosql.NullString{},
		TDestinationChainID: nil,
		TFee:                nil,
		TKappa:              gosql.NullString{},
		TTokenIndexFrom:     nil,
		TTokenIndexTo:       nil,
		TMinDy:              nil,
		TDeadline:           nil,
		TSwapSuccess:        nil,
		TSwapTokenIndex:     nil,
		TSwapMinAmount:      nil,
		TSwapDeadline:       nil,
		TTokenID:            gosql.NullString{},
		TAmountUSD:          nil,
		TFeeAmountUSD:       nil,
		TTokenDecimal:       nil,
		TTokenSymbol:        gosql.NullString{},
		TTimeStamp:          nil,
	})
	Nil(g.T(), err)
	err = api.RehydrateCache(g.GetTestContext(), g.client, responseCache, g.explorerMetrics)
	Nil(g.T(), err)
}
