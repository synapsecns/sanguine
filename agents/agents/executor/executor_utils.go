package executor

import (
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
)

// logToMessage converts the log to a leaf data.
func (e Executor) logToMessage(log ethTypes.Log, chainID uint32) (*types.Message, error) {
	committedMessage, ok := e.chainExecutors[chainID].originParser.ParseSent(log)
	if !ok {
		return nil, fmt.Errorf("could not parse committed message")
	}

	message, err := types.DecodeMessage(committedMessage.Message())
	if err != nil {
		return nil, fmt.Errorf("could not decode message: %w", err)
	}

	return &message, nil
}

// logToAttestation converts the log to an attestation.
func (e Executor) logToAttestation(log ethTypes.Log, chainID uint32) (*types.Attestation, error) {
	attestation, ok := (*e.chainExecutors[chainID].lightInboxParser).ParseAttestationAccepted(log)
	if !ok {
		return nil, fmt.Errorf("could not parse attestation")
	}

	return &attestation, nil
}

// logToSnapshot converts the log to a snapshot.
func (e Executor) logToSnapshot(log ethTypes.Log, chainID uint32) (*types.Snapshot, error) {
	snapshot, domain, ok := (*e.chainExecutors[chainID].inboxParser).ParseSnapshotAccepted(log)
	if !ok {
		return nil, fmt.Errorf("could not parse snapshot")
	}

	if domain == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &snapshot, nil
}

// logType determines whether a log is a `Sent` from Origin.sol or `AttestationAccepted` from Destination.sol.
func (e Executor) logType(log ethTypes.Log, chainID uint32) contractEventType {
	contractEvent := contractEventType{
		contractType: other,
		eventType:    otherEvent,
	}

	//nolint:nestif
	if e.chainExecutors[chainID].inboxParser != nil {
		if summitEvent, ok := (*e.chainExecutors[chainID].inboxParser).EventType(log); ok && summitEvent == inbox.SnapshotAcceptedEvent {
			contractEvent.contractType = inboxContract
			contractEvent.eventType = snapshotAcceptedEvent
		}

		return contractEvent
	}

	//nolint:nestif
	if originEvent, ok := e.chainExecutors[chainID].originParser.EventType(log); ok && originEvent == origin.SentEvent {
		contractEvent.contractType = originContract
		contractEvent.eventType = sentEvent
	} else if destinationEvent, ok := e.chainExecutors[chainID].destinationParser.EventType(log); ok {
		contractEvent.contractType = destinationContract
		if destinationEvent == destination.ExecutedEvent {
			contractEvent.eventType = executedEvent
		}
	} else if lightManagerEvent, ok := (*e.chainExecutors[chainID].lightInboxParser).EventType(log); ok {
		contractEvent.contractType = lightInboxContract
		if lightManagerEvent == lightinbox.AttestationAcceptedEvent {
			contractEvent.eventType = attestationAcceptedEvent
		}
	}

	return contractEvent
}

// verifyAfter guarantees the chronological ordering of logs.
func (l logOrderInfo) verifyAfter(log ethTypes.Log) bool {
	if log.BlockNumber < l.blockNumber {
		return false
	}

	if log.BlockNumber == l.blockNumber {
		return log.Index > l.blockIndex
	}

	return true
}
