package executor

import (
	"context"
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
)

// logToMessage converts the log to a leaf data.
func (e Executor) logToMessage(log ethTypes.Log, chainID uint32) (*types.Message, error) {
	committedMessage, ok := e.chainExecutors[chainID].originParser.ParseDispatch(log)
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
func (e Executor) logToSnapshot(log ethTypes.Log) (*types.Snapshot, error) {
	snapshot, domain, ok := e.summitParser.ParseSnapshotAccepted(log)
	if !ok {
		return nil, fmt.Errorf("could not parse snapshot")
	}

	if domain == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return &snapshot, nil
}

// logType determines whether a log is a `Dispatch` from Origin.sol or `AttestationAccepted` from Destination.sol.
func (e Executor) logType(log ethTypes.Log, chainID uint32) contractEventType {
	contractEvent := contractEventType{
		contractType: other,
		eventType:    otherEvent,
	}

	if eventType, ok := e.chainExecutors[chainID].originParser.EventType(log); ok && eventType == origin.DispatchEvent {
		contractEvent.contractType = originContract
		contractEvent.eventType = dispatchEvent
	} else if eventType, ok := e.chainExecutors[chainID].destinationParser.EventType(log); ok {
		contractEvent.contractType = destinationContract
		if eventType == destination.AttestationAcceptedEvent {
			contractEvent.eventType = attestationAcceptedEvent
		} else if eventType == destination.ExecutedEvent {
			contractEvent.eventType = executedEvent
		}
	}

	// TODO: Add for summit.

	return contractEvent
}

// setMinimumTimes goes through a list of messages and sets the minimum time for each message
// that has an associated attestation.
// The messages need to be sorted by nonce, and the attestations by their destination submission time (which can be via block number or block time).
func (e Executor) setMinimumTimes(ctx context.Context, messages []types.Message, attestations []execTypes.DBAttestation) error {
	messageIndex := 0
	attestationIndex := 0
	for messageIndex < len(messages) && attestationIndex < len(attestations) {
		if messages[messageIndex].Nonce() <= *attestations[attestationIndex].Nonce {
			minimumTime := *attestations[attestationIndex].DestinationBlockTime + uint64(messages[messageIndex].OptimisticSeconds())
			originDomain := messages[messageIndex].OriginDomain()
			destinationDomain := messages[messageIndex].DestinationDomain()
			nonce := messages[messageIndex].Nonce()
			messageMask := execTypes.DBMessage{
				ChainID:     &originDomain,
				Destination: &destinationDomain,
				Nonce:       &nonce,
			}
			err := e.executorDB.SetMinimumTime(ctx, messageMask, minimumTime)
			if err != nil {
				return fmt.Errorf("could not set minimum time: %w", err)
			}

			messageIndex++
		} else {
			attestationIndex++
		}
	}

	return nil
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
