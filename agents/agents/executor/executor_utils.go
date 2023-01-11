package executor

import (
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

// logType determines whether a log is a `Dispatch` from Origin.sol or `AttestationAccepted` from Destination.sol.
func (e Executor) logType(log ethTypes.Log, chainID uint32) contractType {
	contract := other

	if eventType, ok := e.chainExecutors[chainID].originParser.EventType(log); ok && eventType == origin.DispatchEvent {
		contract = originContract
	}

	if eventType, ok := e.chainExecutors[chainID].destinationParser.EventType(log); ok && eventType == destination.AttestationAcceptedEvent {
		contract = destinationContract
	}

	return contract
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

// binarySearchAttestationsForNonce performs a binary search for the attestation with the given nonce, or the attestation
// that has the minimum nonce that is greater than the target nonce.
func binarySearchAttestationsForNonce(attestations []execTypes.DBAttestation, nonce uint32) *execTypes.DBAttestation {
	low := 0
	high := len(attestations) - 1

	for low <= high {
		mid := (low + high) / 2
		switch {
		case *attestations[mid].Nonce == nonce:
			return &attestations[mid]
		case *attestations[mid].Nonce < nonce:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	if low < len(attestations) {
		return &attestations[low]
	}

	return nil
}
