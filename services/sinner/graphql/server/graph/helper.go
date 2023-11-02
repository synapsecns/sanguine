package graph

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"
)

func ifNilString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func ifNilUint32(ptr *int) uint32 {
	if ptr != nil {
		if *ptr < 0 {
			return 0
		}
		return uint32(*ptr)
	}
	return 0
}

func (r *queryResolver) getOriginInfoWithMessageHash(ctx context.Context, messageHash string) ([]*graphqlModel.OriginInfo, error) {
	filter := model.OriginSent{
		MessageHash: messageHash,
	}
	originTxs, err := r.DB.RetrieveOriginSent(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving origin sent data: %w", err)
	}

	var output []*graphqlModel.OriginInfo
	for _, originTx := range originTxs {
		output = append(output, dbToGraphqlModelOrigin(originTx))
	}

	return output, nil
}

func dbToGraphqlModelOriginMultiple(events []model.OriginSent) []*graphqlModel.OriginInfo {
	var output []*graphqlModel.OriginInfo
	for _, event := range events {
		output = append(output, dbToGraphqlModelOrigin(event))
	}
	return output
}

func dbToGraphqlModelDestinationMultiple(events []model.Executed) []*graphqlModel.DestinationInfo {
	var output []*graphqlModel.DestinationInfo
	for _, event := range events {
		output = append(output, dbToGraphqlModelDestination(event))
	}
	return output
}
func dbToGraphqlModelOrigin(event model.OriginSent) *graphqlModel.OriginInfo {
	return &graphqlModel.OriginInfo{
		MessageHash:        &event.MessageHash,
		ContractAddress:    &event.ContractAddress,
		BlockNumber:        core.PtrTo(int(event.BlockNumber)),
		OriginTxHash:       &event.TxHash,
		Sender:             &event.Sender,
		Recipient:          &event.Recipient,
		OriginChainID:      core.PtrTo(int(event.ChainID)),
		DestinationChainID: core.PtrTo(int(event.DestinationChainID)),
		Nonce:              core.PtrTo(int(event.Nonce)),
		Message:            &event.Message,
		OptimisticSeconds:  core.PtrTo(int(event.OptimisticSeconds)),
		MessageFlag:        core.PtrTo(int(event.MessageFlag)),
		SummitTip:          &event.SummitTip,
		AttestationTip:     &event.AttestationTip,
		ExecutionTip:       &event.ExecutionTip,
		DeliveryTip:        &event.DeliveryTip,
		Version:            core.PtrTo(int(event.Version)),
		GasLimit:           core.PtrTo(int(event.GasLimit)),
		GasDrop:            &event.GasDrop,
	}
}

func dbToGraphqlModelDestination(event model.Executed) *graphqlModel.DestinationInfo {
	return &graphqlModel.DestinationInfo{
		ContractAddress: &event.ContractAddress,
		BlockNumber:     core.PtrTo(int(event.BlockNumber)),
		TxHash:          &event.TxHash,
		TxIndex:         core.PtrTo(int(event.TxIndex)),
		MessageHash:     &event.MessageHash,
		ChainID:         core.PtrTo(int(event.ChainID)),
		RemoteDomain:    core.PtrTo(int(event.RemoteDomain)),
		Success:         &event.Success,
	}
}
