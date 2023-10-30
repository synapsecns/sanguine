package graph

import (
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlmodel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"
)

func dbToGraphqlModelOrigin(event model.OriginSent) *graphqlmodel.OriginInfo {
	return &graphqlmodel.OriginInfo{
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

func dbToGraphqlModelDestination(event model.Executed) *graphqlmodel.DestinationInfo {
	return &graphqlmodel.DestinationInfo{
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
