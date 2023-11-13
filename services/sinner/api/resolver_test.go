package api_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
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

	result, err := t.sinnerAPI.GetOriginInfo(t.GetTestContext(), core.PtrTo(messageHash), nil, nil)
	Nil(t.T(), err)
	NotNil(t.T(), result)
	Equal(t.T(), txHash, *result.Response[0].OriginTxHash)

	results, err := t.sinnerAPI.GetOriginInfo(t.GetTestContext(), nil, core.PtrTo(int(chainID)), core.PtrTo(txHash))
	Nil(t.T(), err)
	NotNil(t.T(), results)
	Equal(t.T(), txHash, *results.Response[0].OriginTxHash)
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

	result, err := t.sinnerAPI.GetDestinationInfo(t.GetTestContext(), core.PtrTo(messageHash), nil, nil)
	Nil(t.T(), err)
	NotNil(t.T(), result)
	Equal(t.T(), txHash, *result.Response[0].TxHash)

	results, err := t.sinnerAPI.GetDestinationInfo(t.GetTestContext(), nil, core.PtrTo(int(chainID)), core.PtrTo(txHash))
	Nil(t.T(), err)
	NotNil(t.T(), results)
	Equal(t.T(), txHash, *results.Response[0].TxHash)
}

func (t *APISuite) TestMessageStatus() {
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	desTxHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

	err := t.db.StoreOrUpdateMessageStatus(t.GetTestContext(), txHash, messageHash, types.Origin)
	Nil(t.T(), err)

	result, err := t.sinnerAPI.GetMessageStatus(t.GetTestContext(), core.PtrTo(messageHash), nil, nil)
	Nil(t.T(), err)
	NotNil(t.T(), result)
	Equal(t.T(), txHash, *result.Response.OriginTxHash)
	Equal(t.T(), graphqlModel.MessageStateLastSeenOrigin, *result.Response.LastSeen)

	// Add destination
	err = t.db.StoreOrUpdateMessageStatus(t.GetTestContext(), desTxHash, messageHash, types.Destination)
	Nil(t.T(), err)

	desResult, err := t.sinnerAPI.GetMessageStatus(t.GetTestContext(), core.PtrTo(messageHash), nil, nil)
	Nil(t.T(), err)
	NotNil(t.T(), desResult)
	Equal(t.T(), desTxHash, *desResult.Response.DestinationTxHash)
	Equal(t.T(), graphqlModel.MessageStateLastSeenDestination, *desResult.Response.LastSeen)

	// Test query by origin tx hash
	originSent := &model.OriginSent{
		ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		BlockNumber:     gofakeit.Uint64(),
		TxHash:          txHash,
		MessageHash:     messageHash,
		ChainID:         gofakeit.Uint32(),
	}
	err = t.db.StoreOriginSent(t.GetTestContext(), originSent)
	Nil(t.T(), err)
	desResult, err = t.sinnerAPI.GetMessageStatus(t.GetTestContext(), nil, core.PtrTo(int(originSent.ChainID)), core.PtrTo(originSent.TxHash))
	Nil(t.T(), err)
	NotNil(t.T(), desResult)
	Equal(t.T(), desTxHash, *desResult.Response.DestinationTxHash)
	Equal(t.T(), graphqlModel.MessageStateLastSeenDestination, *desResult.Response.LastSeen)
}

func (t *APISuite) TestPendingMessageStatus() {
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	desTxHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
	messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

	err := t.db.StoreOrUpdateMessageStatus(t.GetTestContext(), txHash, messageHash, types.Origin)
	Nil(t.T(), err)

	result, err := t.sinnerAPI.GetMessagesByStatus(t.GetTestContext(), graphqlModel.MessageStatePending, 1)
	Nil(t.T(), err)
	NotNil(t.T(), result)
	Equal(t.T(), 1, len(result.Response))
	Equal(t.T(), txHash, *result.Response[0].OriginTxHash)
	Equal(t.T(), graphqlModel.MessageStateLastSeenOrigin, *result.Response[0].LastSeen)

	// Add destination
	err = t.db.StoreOrUpdateMessageStatus(t.GetTestContext(), desTxHash, messageHash, types.Destination)
	Nil(t.T(), err)

	desResult, err := t.sinnerAPI.GetMessagesByStatus(t.GetTestContext(), graphqlModel.MessageStatePending, 1)
	Nil(t.T(), err)
	NotNil(t.T(), desResult)
	Equal(t.T(), 0, len(desResult.Response))

	// Check if completed messages query gets two messages.
	completedMessagesResult, err := t.sinnerAPI.GetMessagesByStatus(t.GetTestContext(), graphqlModel.MessageStateCompleted, 1)
	Nil(t.T(), err)
	NotNil(t.T(), desResult)
	Equal(t.T(), 2, len(completedMessagesResult.Response))
}
