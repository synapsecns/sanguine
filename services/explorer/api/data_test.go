package api_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math/big"
	"time"
)

//nolint:cyclop
func (g APISuite) TestGetCountByChainID() {
	chainID := gofakeit.Uint32()
	destinationChainIDA := gofakeit.Uint32()
	destinationChainIDB := gofakeit.Uint32()
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var destinationChainID uint32
		if blockNumber%2 == 0 {
			destinationChainID = destinationChainIDA
		} else {
			destinationChainID = destinationChainIDB
		}
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            chainID,
			ContractAddress:    address.String(),
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
		})
		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDA, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDB, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}
	addressRef := address.String()
	directionRef := model.DirectionOut
	resultOut, err := g.client.GetCountByChainID(g.GetTestContext(), nil, &addressRef, &directionRef, nil)
	Nil(g.T(), err)
	// There should be 3 chains, 2 for the destination chain IDs and 1 for the source chain ID.
	Equal(g.T(), len(resultOut.Response), 3)
	// The source chain ID should have 10 events out, and the destination chain IDs should have 0 events out.
	var reached = 0
	for _, res := range resultOut.Response {
		switch *res.ChainID {
		case int(chainID):
			Equal(g.T(), *res.Count, 10)
			reached++
		case int(destinationChainIDA):
			Equal(g.T(), *res.Count, 0)
			reached++
		case int(destinationChainIDB):
			Equal(g.T(), *res.Count, 0)
			reached++
		}
	}
	Equal(g.T(), reached, 3)

	directionRef = model.DirectionIn
	resultIn, err := g.client.GetCountByChainID(g.GetTestContext(), nil, &addressRef, &directionRef, nil)
	Nil(g.T(), err)
	// Again, there should be 3 chains, 2 for the destination chain IDs and 1 for the source chain ID.
	Equal(g.T(), len(resultIn.Response), 3)
	// The source chain ID should have 0 events in, and the destination chain IDs should have 5 events in.
	reached = 0
	for _, res := range resultIn.Response {
		switch *res.ChainID {
		case int(chainID):
			Equal(g.T(), *res.Count, 0)
			reached++
		case int(destinationChainIDA):
			Equal(g.T(), *res.Count, 5)
			reached++
		case int(destinationChainIDB):
			Equal(g.T(), *res.Count, 5)
			reached++
		}
	}
	Equal(g.T(), reached, 3)
}
