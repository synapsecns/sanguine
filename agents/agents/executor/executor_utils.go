package executor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
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
func (e Executor) logType(log ethTypes.Log, chainID uint32) contractEventType {
	contractEvent := contractEventType{
		contractType:         other,
		destinationEventType: otherEvent,
	}

	if eventType, ok := e.chainExecutors[chainID].originParser.EventType(log); ok && eventType == origin.DispatchEvent {
		contractEvent.contractType = originContract
		contractEvent.destinationEventType = otherEvent
	}

	if eventType, ok := e.chainExecutors[chainID].destinationParser.EventType(log); ok {
		contractEvent.contractType = destinationContract
		if eventType == destination.AttestationAcceptedEvent {
			contractEvent.destinationEventType = attestationAcceptedEvent
		} else if eventType == destination.ExecutedEvent {
			contractEvent.destinationEventType = executedEvent
		}
	}

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

// hashTogether hashes a left and right item together.
func hashTogether(left, right []byte) []byte {
	return crypto.Keccak256(append(left, right...))
}

// getChainLeaf gets the leaf data for a chain within the snapshot root Merkle tree.
func (e Executor) getChainLeaf(originRoot []byte, chainID uint32, snapshot []byte) ([]byte, error) {
	// TODO: snapshot to snapshot type
	originRootAndChainID := hashTogether(originRoot, common.BigToHash(big.NewInt(int64(chainID))).Bytes())
	// TODO: extract actual snapshot data
	chainLeaf := hashTogether(originRootAndChainID, snapshot)

	return chainLeaf, nil
}
