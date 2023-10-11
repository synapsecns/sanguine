package api_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"
	"github.com/synapsecns/sanguine/services/sinner/types"
	"math/big"
)

//nolint:cyclop
func (t *APISuite) TestGetOrigin() {
	chainID := gofakeit.Uint32()
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

	originSent := &model.OriginSent{
		ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		BlockNumber:     gofakeit.Uint64(),
		TxHash:          txHash,
		MessageHash:     messageHash,
		ChainID:         chainID,
	}

	err := t.db.StoreOriginSent(t.GetTestContext(), originSent)
	Nil(t.T(), err)

	result, err := t.sinnerAPI.GetOriginInfo(t.GetTestContext(), txHash, int(chainID))
	Nil(t.T(), err)
	NotNil(t.T(), result)
	Equal(t.T(), txHash, *result.Response.OriginTxHash)
}

func (t *APISuite) TestGetExecuted() {
	chainID := gofakeit.Uint32()
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

	executed := &model.Executed{
		ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		BlockNumber:     gofakeit.Uint64(),
		TxHash:          txHash,
		MessageHash:     messageHash,
		ChainID:         chainID,
		RemoteDomain:    gofakeit.Uint32(),
		Success:         gofakeit.Bool(),
	}

	err := t.db.StoreExecuted(t.GetTestContext(), executed)
	Nil(t.T(), err)

	result, err := t.sinnerAPI.GetDestinationInfo(t.GetTestContext(), txHash, int(chainID))
	Nil(t.T(), err)
	NotNil(t.T(), result)
	Equal(t.T(), txHash, *result.Response.TxHash)
}

func (t *APISuite) TestMessageStatus() {
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	desTxHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

	err := t.db.StoreOrUpdateMessageStatus(t.GetTestContext(), txHash, messageHash, types.Origin)
	Nil(t.T(), err)

	result, err := t.sinnerAPI.GetMessageStatus(t.GetTestContext(), messageHash)
	Nil(t.T(), err)
	NotNil(t.T(), result)
	Equal(t.T(), txHash, *result.Response.OriginTxHash)
	fmt.Println("r", result.Response.LastSeen, result.Response)
	Equal(t.T(), graphqlModel.MessageStateLastSeenOrigin, *result.Response.LastSeen)

	// Add destination
	err = t.db.StoreOrUpdateMessageStatus(t.GetTestContext(), desTxHash, messageHash, types.Destination)
	Nil(t.T(), err)

	desResult, err := t.sinnerAPI.GetMessageStatus(t.GetTestContext(), messageHash)
	Nil(t.T(), err)
	NotNil(t.T(), desResult)
	Equal(t.T(), desTxHash, *desResult.Response.DestinationTxHash)
	Equal(t.T(), graphqlModel.MessageStateLastSeenDestination, *desResult.Response.LastSeen)
}
