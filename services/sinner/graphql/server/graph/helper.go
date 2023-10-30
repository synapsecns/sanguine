package graph

import (
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlmodel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"
)

// Helper function to convert int to *int.
func intPtr(val int) *int {
	return &val
}

// Helper function to convert string to *string.
func strPtr(val string) *string {
	return &val
}

func dbToGraphqlModelOrigin(event model.OriginSent) *graphqlmodel.OriginInfo {
	return &graphqlmodel.OriginInfo{
		MessageHash:        &event.MessageHash,
		ContractAddress:    &event.ContractAddress,
		BlockNumber:        intPtr(int(event.BlockNumber)),
		OriginTxHash:       &event.TxHash,
		Sender:             &event.Sender,
		Recipient:          &event.Recipient,
		OriginChainID:      intPtr(int(event.ChainID)),
		DestinationChainID: intPtr(int(event.DestinationChainID)),
		Nonce:              intPtr(int(event.Nonce)),
		Message:            strPtr(event.Message),
		OptimisticSeconds:  intPtr(int(event.OptimisticSeconds)),
		MessageFlag:        intPtr(int(event.MessageFlag)),
		SummitTip:          &event.SummitTip,
		AttestationTip:     &event.AttestationTip,
		ExecutionTip:       &event.ExecutionTip,
		DeliveryTip:        &event.DeliveryTip,
		Version:            intPtr(int(event.Version)),
		GasLimit:           intPtr(int(event.GasLimit)),
		GasDrop:            &event.GasDrop,
	}
}

func dbToGraphqlModelDestination(event model.Executed) *graphqlmodel.DestinationInfo {
	return &graphqlmodel.DestinationInfo{
		ContractAddress: &event.ContractAddress,
		BlockNumber:     intPtr(int(event.BlockNumber)),
		TxHash:          &event.TxHash,
		TxIndex:         intPtr(int(event.TxIndex)),
		MessageHash:     &event.MessageHash,
		ChainID:         intPtr(int(event.ChainID)),
		RemoteDomain:    intPtr(int(event.RemoteDomain)),
		Success:         &event.Success,
	}
}
