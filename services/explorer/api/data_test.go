package api_test

import (
	gosql "database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math/big"
	"sort"
	"time"
)

// nolint:cyclop
func (g APISuite) TestBridgeAmountStatistic() {
	chainID := gofakeit.Uint32()
	destinationChainIDA := gofakeit.Uint32()
	destinationChainIDB := gofakeit.Uint32()
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()

	cumulativePrice := []float64{}
	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var destinationChainID uint32
		if blockNumber%2 == 0 {
			destinationChainID = destinationChainIDA
		} else {
			destinationChainID = destinationChainIDB
		}
		price := float64(gofakeit.Number(1, 300))
		cumulativePrice = append(cumulativePrice, price)
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            chainID,
			Recipient:          gosql.NullString{String: address.String(), Valid: true},
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
			Token:              tokenAddr,
			Amount:             big.NewInt(int64(gofakeit.Number(1, 300))),
			AmountUSD:          &price,
		})
		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDA, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDB, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}

	total := 0.0
	for _, v := range cumulativePrice {
		total += v
	}
	count := float64(len(cumulativePrice))
	mean := total / count
	median := 0.0
	sort.Float64s(cumulativePrice)
	switch {
	case count == 0:
		median = 0.0
	case len(cumulativePrice)%2 == 0:
		median = (cumulativePrice[len(cumulativePrice)/2-1] + cumulativePrice[len(cumulativePrice)/2]) / 2
	default:
		median = cumulativePrice[len(cumulativePrice)/2]
	}

	statType := model.StatisticTypeTotal
	duration := model.DurationPastDay
	result, err := g.client.GetBridgeAmountStatistic(g.GetTestContext(), statType, &duration, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), fmt.Sprintf("%f", total), *result.Response.USDValue)

	statType = model.StatisticTypeCount
	result, err = g.client.GetBridgeAmountStatistic(g.GetTestContext(), statType, &duration, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), fmt.Sprintf("%f", count), *result.Response.USDValue)

	statType = model.StatisticTypeMean
	result, err = g.client.GetBridgeAmountStatistic(g.GetTestContext(), statType, &duration, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), fmt.Sprintf("%f", mean), *result.Response.USDValue)

	statType = model.StatisticTypeMedian
	result, err = g.client.GetBridgeAmountStatistic(g.GetTestContext(), statType, &duration, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), fmt.Sprintf("%f", median), *result.Response.USDValue)
}
