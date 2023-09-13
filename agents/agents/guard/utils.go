package guard

import (
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
)

func isSnapshotAcceptedEvent(parser inbox.Parser, log ethTypes.Log) bool {
	inboxEvent, ok := parser.EventType(log)
	return ok && inboxEvent == inbox.SnapshotAcceptedEvent
}

func isAttestationAcceptedEvent(parser lightinbox.Parser, log ethTypes.Log) bool {
	lightManagerEvent, ok := parser.EventType(log)
	return ok && lightManagerEvent == lightinbox.AttestationAcceptedEvent
}

func isReceiptAcceptedEvent(parser inbox.Parser, log ethTypes.Log) bool {
	inboxEvent, ok := parser.EventType(log)
	return ok && inboxEvent == inbox.ReceiptAcceptedEvent
}

func isStatusUpdatedEvent(parser bondingmanager.Parser, log ethTypes.Log) bool {
	bondingManagerEvent, ok := parser.EventType(log)
	return ok && bondingManagerEvent == bondingmanager.StatusUpdatedEvent
}

func isRootUpdatedEvent(bondingParser bondingmanager.Parser, log ethTypes.Log) bool {
	bondingManagerEvent, ok := bondingParser.EventType(log)
	if ok && bondingManagerEvent == bondingmanager.RootUpdatedEvent {
		return true
	}
	return false
}
