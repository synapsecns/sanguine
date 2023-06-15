package executor

import (
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/types"
)

// logToMessage converts the log to a leaf data.
func (e Executor) logToMessage(log ethTypes.Log, chainID uint32) (types.Message, error) {
	message, ok := e.chainExecutors[chainID].originParser.ParseSent(log)
	if !ok {
		return nil, fmt.Errorf("could not parse committed message")
	}

	if message == nil {
		//nolint:nilnil
		return nil, nil
	}

	return message, nil
}

// logToAttestation converts the log to an attestation.
func (e Executor) logToAttestation(log ethTypes.Log, chainID uint32) (types.Attestation, error) {
	attestation, ok := e.chainExecutors[chainID].lightInboxParser.ParseAttestationAccepted(log)
	if !ok {
		return nil, fmt.Errorf("could not parse attestation")
	}

	if attestation == nil {
		//nolint:nilnil
		return nil, nil
	}

	return attestation, nil
}

// logToSnapshot converts the log to a snapshot.
func (e Executor) logToSnapshot(log ethTypes.Log, chainID uint32) (types.Snapshot, error) {
	snapshot, domain, ok := e.chainExecutors[chainID].inboxParser.ParseSnapshotAccepted(log)
	if !ok {
		return nil, fmt.Errorf("could not parse snapshot")
	}

	if snapshot == nil || domain == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return snapshot, nil
}

func (e Executor) logToInterface(log ethTypes.Log, chainID uint32) (any, error) {
	switch {
	case e.isSnapshotAcceptedEvent(log, chainID):
		return e.logToSnapshot(log, chainID)
	case e.isSentEvent(log, chainID):
		return e.logToMessage(log, chainID)
	case e.isAttestationAcceptedEvent(log, chainID):
		return e.logToAttestation(log, chainID)
	default:
		//nolint:nilnil
		return nil, nil
	}
}

func (e Executor) isSnapshotAcceptedEvent(log ethTypes.Log, chainID uint32) bool {
	if e.chainExecutors[chainID].inboxParser == nil {
		return false
	}

	inboxEvent, ok := e.chainExecutors[chainID].inboxParser.EventType(log)
	return ok && inboxEvent == inbox.SnapshotAcceptedEvent
}

func (e Executor) isSentEvent(log ethTypes.Log, chainID uint32) bool {
	if e.chainExecutors[chainID].originParser == nil {
		return false
	}

	originEvent, ok := e.chainExecutors[chainID].originParser.EventType(log)
	return ok && originEvent == origin.SentEvent
}

func (e Executor) isAttestationAcceptedEvent(log ethTypes.Log, chainID uint32) bool {
	if e.chainExecutors[chainID].lightInboxParser == nil {
		return false
	}

	lightManagerEvent, ok := e.chainExecutors[chainID].lightInboxParser.EventType(log)
	return ok && lightManagerEvent == lightinbox.AttestationAcceptedEvent
}
