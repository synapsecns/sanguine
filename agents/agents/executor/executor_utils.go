package executor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
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

// logType determines whether a log is a `Dispatch` from Origin.sol or `AttestationAccepted` from Destination.sol.
func (e Executor) logType(log ethTypes.Log, chainID uint32) contractEventType {
	contractEvent := contractEventType{
		contractType: other,
		eventType:    otherEvent,
	}

	//nolint:nestif
	if e.chainExecutors[chainID].summitParser != nil {
		if eventType, ok := (*e.chainExecutors[chainID].summitParser).EventType(log); ok {
			contractEvent.contractType = summitContract
			if eventType == summit.SnapshotAcceptedEvent {
				contractEvent.eventType = snapshotAcceptedEvent
			}
		}
	} else if eventType, ok := e.chainExecutors[chainID].originParser.EventType(log); ok && eventType == origin.DispatchEvent {
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

// getEarliestStateInRange returns the earliest state with the same snapshot root as an attestation within a nonce range.
func (e Executor) getEarliestStateInRange(ctx context.Context, origin, destination uint32, startNonce, endNonce uint32) (*types.State, error) {
	snapshotRoots, err := e.executorDB.GetSnapshotRootsInNonceRange(ctx, origin, startNonce, endNonce)
	if err != nil {
		return nil, fmt.Errorf("could not get snapshot roots: %w", err)
	}

	if len(snapshotRoots) == 0 {
		//nolint:nilnil
		return nil, nil
	}

	attestationMask := execTypes.DBAttestation{
		Destination: &destination,
	}

	earliestSnapshotRoot, err := e.executorDB.GetEarliestSnapshotFromAttestation(ctx, attestationMask, snapshotRoots)
	if err != nil {
		return nil, fmt.Errorf("could not get earliest snapshot root: %w", err)
	}

	if earliestSnapshotRoot == nil {
		//nolint:nilnil
		return nil, nil
	}

	earliestSnapshotRootString := common.BytesToHash((*earliestSnapshotRoot)[:]).String()

	stateMask := execTypes.DBState{
		SnapshotRoot: &earliestSnapshotRootString,
		ChainID:      &origin,
	}

	stateWithEarliestAttestation, err := e.executorDB.GetState(ctx, stateMask)
	if err != nil {
		return nil, fmt.Errorf("could not get state with earliest attestation: %w", err)
	}

	return stateWithEarliestAttestation, nil
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
