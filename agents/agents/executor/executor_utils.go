package executor

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
func (e Executor) logToAttestation(log ethTypes.Log, chainID uint32, summitAttestation bool) (*types.AttestationWithMetadata, error) {
	var attestationMetadata *types.AttestationWithMetadata

	var err error
	if summitAttestation {
		// This is a guard attestation.
		attestation, ok := e.chainExecutors[chainID].summitParser.ParseAttestationSaved(log)
		if !ok {
			return nil, fmt.Errorf("could not parse attestation")
		}
		attestationMetadata = &types.AttestationWithMetadata{Attestation: attestation}
	} else {
		// This is a notary attestation.
		attestationMetadata, err = e.chainExecutors[chainID].lightInboxParser.ParseAttestationAccepted(log)
		if err != nil {
			return nil, fmt.Errorf("could not parse attestation: %w", err)
		}
	}

	if attestationMetadata == nil {
		//nolint:nilnil
		return nil, nil
	}

	return attestationMetadata, nil
}

// logToSnapshot converts the log to a snapshot.
func (e Executor) logToSnapshot(log ethTypes.Log, chainID uint32) (types.Snapshot, error) {
	snapshotMetadata, err := e.chainExecutors[chainID].inboxParser.ParseSnapshotAccepted(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse snapshot: %w", err)
	}

	if snapshotMetadata.Snapshot == nil || snapshotMetadata.AgentDomain() == 0 {
		//nolint:nilnil
		return nil, nil
	}

	return snapshotMetadata.Snapshot, nil
}

func (e Executor) logToInterface(log ethTypes.Log, chainID uint32) (any, error) {
	switch {
	case e.isSnapshotAcceptedEvent(log, chainID):
		return e.logToSnapshot(log, chainID)
	case e.isSentEvent(log, chainID):
		return e.logToMessage(log, chainID)
	case e.isAttestationAcceptedEvent(log, chainID):
		return e.logToAttestation(log, chainID, false)
	case e.isAttestationSavedEvent(log, chainID):
		return e.logToAttestation(log, chainID, true)
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

func (e Executor) isAttestationSavedEvent(log ethTypes.Log, chainID uint32) bool {
	if e.chainExecutors[chainID].summitParser == nil {
		return false
	}

	summitEvent, ok := e.chainExecutors[chainID].summitParser.EventType(log)
	return ok && summitEvent == summit.AttestationSavedEvent
}

// processMessage processes and stores a message.
func (e Executor) processMessage(ctx context.Context, message types.Message, log ethTypes.Log) (err error) {
	types.LogTx("EXECUTOR", fmt.Sprintf("Processing message: %s", types.MessageToString(message)), message.OriginDomain(), nil)
	ctx, span := e.handler.Tracer().Start(ctx, "processMessage", trace.WithAttributes(
		attribute.String(metrics.TxHash, log.TxHash.String()),
		attribute.Int(metrics.BlockNumber, int(log.BlockNumber)),
		attribute.Int(metrics.Origin, int(message.OriginDomain())),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	leaf, err := message.ToLeaf()
	if err != nil {
		return fmt.Errorf("could not convert message to leaf: %w", err)
	}
	span.SetAttributes(attribute.String(metrics.MessageLeaf, common.BytesToHash(leaf[:]).String()))

	// Sanity check to make sure that the message has come from Origin.
	if log.Address.String() != e.chainConfigs[message.OriginDomain()].OriginAddress {
		span.AddEvent("message is not from origin", trace.WithAttributes(
			attribute.String("log_address", log.Address.String()),
			attribute.String("origin_address", e.config.Chains[message.OriginDomain()].OriginAddress),
		))
		return nil
	}

	merkleIndex := e.chainExecutors[message.OriginDomain()].merkleTree.NumOfItems()
	span.SetAttributes(attribute.Int("merkle_index", int(merkleIndex)))

	// Make sure the nonce of the message is being inserted at the right index.
	span.AddEvent("validating message nonce", trace.WithAttributes(
		attribute.Int(metrics.Nonce, int(message.Nonce())),
		attribute.Int("merkle_index_plus_1", int(merkleIndex+1)),
	))
	switch {
	case merkleIndex+1 > message.Nonce():
		return nil
	case merkleIndex+1 < message.Nonce():
		span.SetAttributes(
			attribute.Int("nonce", int(message.Nonce())),
			attribute.Int("merkle_index", int(merkleIndex)),
		)
		span.AddEvent("nonce is not correct")
		logger.Warnf("nonce is not correct. expected: %d, got: %d", merkleIndex+1, message.Nonce())
		return nil
	default:
	}

	e.chainExecutors[message.OriginDomain()].merkleTree.Insert(leaf[:])

	span.AddEvent("storing message", trace.WithAttributes(attribute.Int("nonce", int(message.Nonce()))))
	err = e.executorDB.StoreMessage(ctx, message, log.BlockNumber, false, 0, log.TxHash)
	if err != nil {
		return fmt.Errorf("could not store message: %w", err)
	}

	return nil
}

// processSnapshot processes and stores a snapshot.
func (e Executor) processSnapshot(ctx context.Context, snapshot types.Snapshot, log ethTypes.Log) (err error) {
	snapshotRoot, proofs, err := snapshot.SnapshotRootAndProofs()
	if err != nil {
		return fmt.Errorf("could not get snapshot root and proofs: %w", err)
	}

	ctx, span := e.handler.Tracer().Start(ctx, "processSnapshot", trace.WithAttributes(
		attribute.Int("logBlockNumber", int(log.BlockNumber)),
		attribute.String(metrics.SnapRoot, common.BytesToHash(snapshotRoot[:]).String()),
		attribute.String("txHash", log.TxHash.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	for _, s := range snapshot.States() {
		span.AddEvent("processing state", trace.WithAttributes(
			attribute.Int(metrics.Origin, int(s.Origin())),
			attribute.Int("nonce", int(s.Nonce())),
		))
		state := s
		statePayload, err := state.Encode()
		if err != nil {
			return fmt.Errorf("could not encode state: %w", err)
		}
		// Verify that the state is valid w.r.t. Origin.
		var valid bool
		contractCall := func(ctx context.Context) error {
			valid, err = e.chainExecutors[state.Origin()].boundOrigin.IsValidState(
				ctx,
				statePayload,
			)
			if err != nil {
				return fmt.Errorf("could not check validity of state: %w", err)
			}
			return nil
		}
		err = retry.WithBackoff(ctx, contractCall, e.retryConfig...)
		if err != nil {
			return fmt.Errorf("could not check validity of state: %w", err)
		}

		if !valid {
			stateRoot := state.Root()
			span.AddEvent("snapshot has invalid state", trace.WithAttributes(
				attribute.Int(metrics.Origin, int(state.Origin())),
				attribute.String(metrics.StateRoot, common.BytesToHash(stateRoot[:]).String()),
			))
			return nil
		}
	}

	span.AddEvent("storing states", trace.WithAttributes(attribute.Int("num_states", len(snapshot.States()))))
	err = e.executorDB.StoreStates(ctx, snapshot.States(), snapshotRoot, proofs, log.BlockNumber)
	if err != nil {
		return fmt.Errorf("could not store states: %w", err)
	}

	return nil
}

// processAttestation processes and stores an attestation.
func (e Executor) processAttestation(ctx context.Context, attestationMetadata types.AttestationWithMetadata, chainID uint32, log ethTypes.Log) (err error) {
	attestation := attestationMetadata.Attestation
	snapshotRoot := attestation.SnapshotRoot()
	ctx, span := e.handler.Tracer().Start(ctx, "processAttestation", trace.WithAttributes(
		attribute.Int("chainID", int(chainID)),
		attribute.Int("logBlockNumber", int(log.BlockNumber)),
		attribute.String(metrics.SnapRoot, common.BytesToHash(snapshotRoot[:]).String()),
		attribute.String("txHash", log.TxHash.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Make sure that we store the states corresponding to this attestation.
	attPayload := attestationMetadata.AttestationPayload()
	if attPayload != nil {
		snapshot, snapErr := e.chainExecutors[e.config.SummitChainID].boundSummit.GetNotarySnapshot(ctx, attPayload)
		if snapErr != nil {
			logger.Warnf("could not get snapshot for attestation with snapRoot %v: %v", snapshotRoot, err)
			span.AddEvent("could not get snapshot")
		} else {
			_, proofs, snapErr := snapshot.SnapshotRootAndProofs()
			if snapErr != nil {
				span.AddEvent("could not get proofs")
				logger.Warnf("could not get snapshot root and proofs for attestation with snapRoot %v: %v", snapshotRoot, err)
			} else {
				span.AddEvent("storing states", trace.WithAttributes(attribute.Int("num_states", len(snapshot.States()))))
				e.executorDB.StoreStates(ctx, snapshot.States(), snapshotRoot, proofs, log.BlockNumber)
			}
		}
	}

	// If the attestation is on the SynChain, we can directly use its block number and timestamp.
	if chainID == e.config.SummitChainID {
		span.AddEvent("storing summit attestation", trace.WithAttributes(attribute.Int("time", int(attestation.Timestamp().Uint64()))))
		err := e.executorDB.StoreAttestation(ctx, attestation, chainID, attestation.BlockNumber().Uint64(), attestation.Timestamp().Uint64())
		if err != nil {
			return fmt.Errorf("could not store attestation: %w", err)
		}

		return nil
	}

	// If the attestation is on a remote chain, we need to fetch the timestamp via an RPC call.
	var logHeader *ethTypes.Header
	contractCall := func(ctx context.Context) error {
		logHeader, err = e.chainExecutors[chainID].rpcClient.HeaderByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
		if err != nil {
			return fmt.Errorf("could not get log header: %w", err)
		}

		return nil
	}
	err = retry.WithBackoff(ctx, contractCall, e.retryConfig...)
	if err != nil {
		return fmt.Errorf("could not get log header: %w", err)
	}

	if logHeader == nil {
		return fmt.Errorf("could not get log header")
	}

	span.AddEvent("storing remote attestation", trace.WithAttributes(attribute.Int("time", int(logHeader.Time))))
	err = e.executorDB.StoreAttestation(ctx, attestation, chainID, log.BlockNumber, logHeader.Time)
	if err != nil {
		return fmt.Errorf("could not store attestation: %w", err)
	}

	return nil
}
