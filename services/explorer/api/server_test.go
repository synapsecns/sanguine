package api_test

import (
	gosql "database/sql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"math/big"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/api"
	"github.com/synapsecns/sanguine/services/explorer/api/cache"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/graphql/client"
)

func TestHandleJSONAmountStat(t *testing.T) {

	valueString := gofakeit.Word()
	valueStruct := gqlClient.GetAmountStatistic{
		Response: &struct {
			Value *string "json:\"value\" graphql:\"value\""
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
	valueStruct := gqlClient.GetDailyStatisticsByChain{
		Response: []*struct {
			Date      *string  "json:\"date\" graphql:\"date\""
			Ethereum  *float64 "json:\"ethereum\" graphql:\"ethereum\""
			Optimism  *float64 "json:\"optimism\" graphql:\"optimism\""
			Cronos    *float64 "json:\"cronos\" graphql:\"cronos\""
			Bsc       *float64 "json:\"bsc\" graphql:\"bsc\""
			Polygon   *float64 "json:\"polygon\" graphql:\"polygon\""
			Fantom    *float64 "json:\"fantom\" graphql:\"fantom\""
			Boba      *float64 "json:\"boba\" graphql:\"boba\""
			Metis     *float64 "json:\"metis\" graphql:\"metis\""
			Moonbeam  *float64 "json:\"moonbeam\" graphql:\"moonbeam\""
			Moonriver *float64 "json:\"moonriver\" graphql:\"moonriver\""
			Klaytn    *float64 "json:\"klaytn\" graphql:\"klaytn\""
			Arbitrum  *float64 "json:\"arbitrum\" graphql:\"arbitrum\""
			Avalanche *float64 "json:\"avalanche\" graphql:\"avalanche\""
			Dfk       *float64 "json:\"dfk\" graphql:\"dfk\""
			Aurora    *float64 "json:\"aurora\" graphql:\"aurora\""
			Harmony   *float64 "json:\"harmony\" graphql:\"harmony\""
			Canto     *float64 "json:\"canto\" graphql:\"canto\""
			Dogechain *float64 "json:\"dogechain\" graphql:\"dogechain\""
			Total     *float64 "json:\"total\" graphql:\"total\""
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
	type MvBridgeEvent struct {
		InsertTime uint64 `gorm:"column:insert_time"`
		// FInsertTime is the time the event was inserted into the database.
		FInsertTime uint64 `gorm:"column:finsert_time"`
		// FContractAddress is the address of the contract that generated the event.
		FContractAddress string `gorm:"column:fcontract_address"`
		// FChainID is the chain id of the contract that generated the event.
		FChainID uint32 `gorm:"column:fchain_id"`
		// FEventType is the type of the event.
		FEventType uint8 `gorm:"column:fevent_type"`
		// FBlockNumber is the block number of the event.
		FBlockNumber uint64 `gorm:"column:fblock_number"`
		// FTxHash is the transaction hash of the event.
		FTxHash string `gorm:"column:ftx_hash"`
		// FToken is the address of the token.
		FToken string `gorm:"column:ftoken"`
		// FAmount is the amount of tokens.
		FAmount *big.Int `gorm:"column:famount;type:UInt256"`
		// FEventIndex is the index of the log.
		FEventIndex uint64 `gorm:"column:fevent_index"`
		// FDestinationKappa is the destination kappa.
		FDestinationKappa string `gorm:"column:fdestination_kappa"`
		// FSender is the address of the sender.
		FSender string `gorm:"column:fsender"`

		// FRecipient is the address to send the tokens to.
		FRecipient gosql.NullString `gorm:"column:frecipient"`
		// FRecipientBytes is the recipient address in bytes.
		FRecipientBytes gosql.NullString `gorm:"column:frecipient_bytes"`
		// FDestinationChainID is the chain id of the chain to send the tokens to.
		FDestinationChainID *big.Int `gorm:"column:fdestination_chain_id;type:UInt256"`
		// FFee is the fee.
		FFee *big.Int `gorm:"column:ffee;type:UInt256"`
		// FKappa is theFee keccak256 hash of the transaction.
		FKappa gosql.NullString `gorm:"column:fkappa"`
		// FTokenIndexFrom is the index of the from token in the pool.
		FTokenIndexFrom *big.Int `gorm:"column:ftoken_index_from;type:UInt256"`
		// FTokenIndexTo is the index of the to token in the pool.
		FTokenIndexTo *big.Int `gorm:"column:ftoken_index_to;type:UInt256"`
		// FMinDy is the minimum amount of tokens to receive.
		FMinDy *big.Int `gorm:"column:fmin_dy;type:UInt256"`
		// FDeadline is the deadline of the transaction.
		FDeadline *big.Int `gorm:"column:fdeadline;type:UInt256"`
		// FSwapSuccess is whether the swap was successful.
		FSwapSuccess *big.Int `gorm:"column:fswap_success;type:UInt256"`
		// FSwapTokenIndex is the index of the token in the pool.
		FSwapTokenIndex *big.Int `gorm:"column:fswap_token_index;type:UInt256"`
		// FSwapMinAmount is the minimum amount of tokens to receive.
		FSwapMinAmount *big.Int `gorm:"column:fswap_min_amount;type:UInt256"`
		// FSwapDeadline is the deadline of the swap transaction.
		FSwapDeadline *big.Int `gorm:"column:fswap_deadline;type:UInt256"`
		// FTokenID is the token's ID.
		FTokenID gosql.NullString `gorm:"column:ftoken_id"`
		// FAmountUSD is the amount in USD.
		FAmountUSD *float64 `gorm:"column:famount_usd;type:Float64"`
		// FFeeAmountUSD is the fee amount in USD.
		FFeeAmountUSD *float64 `gorm:"column:ffee_amount_usd;type:Float64"`
		// FTokenDecimal is the token's decimal.
		FTokenDecimal *uint8 `gorm:"column:ftoken_decimal"`
		// FTokenSymbol is the token's symbol from coin gecko.
		FTokenSymbol gosql.NullString `gorm:"column:ftoken_symbol"`
		// FTimeStamp is the timestamp of the block in which the event occurred.
		FTimeStamp *uint64 `gorm:"column:ftimestamp"`
		// TInsertTime is the time the event was inserted into the database.
		TInsertTime uint64 `gorm:"column:finsert_time"`
		// TContractAddress is the address of the contract that generated the event.
		TContractAddress string `gorm:"column:tcontract_address"`
		// TChainID is the chain id of the contract that generated the event.
		TChainID uint32 `gorm:"column:tchain_id"`
		// TEventType is the type of the event.
		TEventType uint8 `gorm:"column:tevent_type"`
		// TBlockNumber is the block number of the event.
		TBlockNumber uint64 `gorm:"column:tblock_number"`
		// TTxHash is the transaction hash of the event.
		TTxHash string `gorm:"column:ttx_hash"`
		// TToken is the address of the token.
		TToken string `gorm:"column:ttoken"`
		// TAmount is the amount of tokens.
		TAmount *big.Int `gorm:"column:tamount;type:UInt256"`
		// TEventIndex is the index of the log.
		TEventIndex uint64 `gorm:"column:tevent_index"`
		// TDestinationKappa is the destination kappa.
		TDestinationKappa string `gorm:"column:tdestination_kappa"`
		// TSender is the address of the sender.
		TSender string `gorm:"column:tsender"`

		// TRecipient is the address to send the tokens to.
		TRecipient gosql.NullString `gorm:"column:trecipient"`
		// TRecipientBytes is the recipient address in bytes.
		TRecipientBytes gosql.NullString `gorm:"column:trecipient_bytes"`
		// TDestinationChainID is the chain id of the chain to send the tokens to.
		TDestinationChainID *big.Int `gorm:"column:tdestination_chain_id;type:UInt256"`
		// TFee is the fee.
		TFee *big.Int `gorm:"column:tfee;type:UInt256"`
		// TKappa is theFee keccak256 hash of the transaction.
		TKappa gosql.NullString `gorm:"column:tkappa"`
		// TTokenIndexFrom is the index of the from token in the pool.
		TTokenIndexFrom *big.Int `gorm:"column:ttoken_index_from;type:UInt256"`
		// TTokenIndexTo is the index of the to token in the pool.
		TTokenIndexTo *big.Int `gorm:"column:ttoken_index_to;type:UInt256"`
		// TMinDy is the minimum amount of tokens to receive.
		TMinDy *big.Int `gorm:"column:tmin_dy;type:UInt256"`
		// TDeadline is the deadline of the transaction.
		TDeadline *big.Int `gorm:"column:tdeadline;type:UInt256"`
		// TSwapSuccess is whether the swap was successful.
		TSwapSuccess *big.Int `gorm:"column:tswap_success;type:UInt256"`
		// TSwapTokenIndex is the index of the token in the pool.
		TSwapTokenIndex *big.Int `gorm:"column:tswap_token_index;type:UInt256"`
		// TSwapMinAmount is the minimum amount of tokens to receive.
		TSwapMinAmount *big.Int `gorm:"column:tswap_min_amount;type:UInt256"`
		// TSwapDeadline is the deadline of the swap transaction.
		TSwapDeadline *big.Int `gorm:"column:tswap_deadline;type:UInt256"`
		// TTokenID is the token's ID.
		TTokenID gosql.NullString `gorm:"column:ttoken_id"`
		// TAmountUSD is the amount in USD.
		TAmountUSD *float64 `gorm:"column:tamount_usd;type:Float64"`
		// TFeeAmountUSD is the fee amount in USD.
		TFeeAmountUSD *float64 `gorm:"column:tfee_amount_usd;type:Float64"`
		// TTokenDecimal is the token's decimal.
		TTokenDecimal *uint8 `gorm:"column:ttoken_decimal"`
		// TTokenSymbol is the token's symbol from coin gecko.
		TTokenSymbol gosql.NullString `gorm:"column:ttoken_symbol"`
		// TTimeStamp is the timestamp of the block in which the event occurred.
		TTimeStamp *uint64 `gorm:"column:ttimestamp"`
	}
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
	err = api.RehydrateCache(g.GetTestContext(), g.client, responseCache)
	Nil(g.T(), err)

}
