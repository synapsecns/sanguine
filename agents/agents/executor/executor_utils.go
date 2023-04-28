package executor

import (
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
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
	attestation, ok := e.chainExecutors[chainID].destinationParser.ParseAttestationAccepted(log)
	if !ok {
		return nil, fmt.Errorf("could not parse attestation")
	}

	return &attestation, nil
}

// logToSnapshot converts the log to a snapshot.
func (e Executor) logToSnapshot(log ethTypes.Log, chainID uint32) (*types.Snapshot, error) {
	snapshot, domain, ok := (*e.chainExecutors[chainID].summitParser).ParseSnapshotAccepted(log)
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
	if e.chainExecutors[chainID].summitParser != nil {
		if summitEvent, ok := (*e.chainExecutors[chainID].summitParser).EventType(log); ok && summitEvent == summit.SnapshotAcceptedEvent {
			contractEvent.contractType = summitContract
			contractEvent.eventType = snapshotAcceptedEvent
		}
	} else if originEvent, ok := e.chainExecutors[chainID].originParser.EventType(log); ok && originEvent == origin.SentEvent {
		contractEvent.contractType = originContract
		contractEvent.eventType = sentEvent
	} else if destinationEvent, ok := e.chainExecutors[chainID].destinationParser.EventType(log); ok {
		contractEvent.contractType = destinationContract
		if destinationEvent == destination.AttestationAcceptedEvent {
			contractEvent.eventType = attestationAcceptedEvent
		} else if destinationEvent == destination.ExecutedEvent {
			contractEvent.eventType = executedEvent
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
