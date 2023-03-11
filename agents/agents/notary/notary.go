package notary

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"strconv"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
)

// Notary checks the Summit for that latest states signed by guards, validates those states on origin,
// then signs and submits the snapshot to Summit.
type Notary struct {
	bondedSigner            signer.Signer
	unbondedSigner          signer.Signer
	domains                 []domains.DomainClient
	summitDomain            domains.DomainClient
	destinationDomain       domains.DomainClient
	refreshInterval         time.Duration
	summitMyLatestStates    map[uint32]types.State
	summitGuardLatestStates map[uint32]types.State
	summitParser            summit.Parser
}

// NewNotary creates a new notary.
//
//nolint:cyclop
func NewNotary(ctx context.Context, cfg config.AgentConfig, scribeClient client.ScribeClient) (_ Notary, err error) {
	notary := Notary{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalSeconds),
	}
	notary.domains = []domains.DomainClient{}

	notary.bondedSigner, err = config.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with bondedSigner, could not create notary: %w", err)
	}

	notary.unbondedSigner, err = config.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with unbondedSigner, could not create notary: %w", err)
	}

	for domainName, domain := range cfg.Domains {
		var domainClient domains.DomainClient
		domainClient, err = evm.NewEVM(ctx, domainName, domain)
		if err != nil {
			return Notary{}, fmt.Errorf("failing to create evm for domain, could not create notary for: %w", err)
		}
		notary.domains = append(notary.domains, domainClient)
		if domain.DomainID == cfg.SummitDomainID {
			notary.summitDomain = domainClient
		}
		if domain.DomainID == cfg.DomainID {
			notary.destinationDomain = domainClient
		}
	}

	notary.summitMyLatestStates = make(map[uint32]types.State, len(notary.domains))
	notary.summitGuardLatestStates = make(map[uint32]types.State, len(notary.domains))

	notary.summitParser, err = summit.NewParser(common.HexToAddress(notary.summitDomain.Config().SummitAddress))
	if err != nil {
		return Notary{}, fmt.Errorf("could not create summit parser: %w", err)
	}

	// Scribe gRPC setup.
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return Notary{}, fmt.Errorf("could not dial grpc: %w", err)
	}

	grpcClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := grpcClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return Notary{}, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return Notary{}, fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	return notary, nil
}

func (n Notary) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient) error {
	// TODO: How to set your `fromBlock` without a database?
	fromBlock := strconv.FormatUint(0, 10)

	stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: n.summitDomain.Config().SummitAddress}},
			ChainId:         n.summitDomain.Config().DomainID,
		},
		FromBlock: fromBlock,
		ToBlock:   "latest",
	})
	if err != nil {
		return fmt.Errorf("could not stream logs: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			log := response.Log.ToLog()
			if log == nil {
				return fmt.Errorf("could not convert to log")
			}

			attestation, err := n.logToAttestation(*log)
			if err != nil {
				return fmt.Errorf("could not convert to attestation: %w", err)

			}
			if attestation == nil {
				return fmt.Errorf("could not convert to attestation")
			}

			// Do your stuff with the attestation here!

		}
	}
}

func (n Notary) logToAttestation(log ethTypes.Log) (*types.Attestation, error) {
	attestationEvent, ok := n.summitParser.ParseAttestationSaved(log)
	if !ok {
		return nil, fmt.Errorf("could not parse attestation")
	}

	attestation, err := types.DecodeAttestation(attestationEvent)
	if err != nil {
		return nil, fmt.Errorf("could not decode attestation: %w", err)
	}

	return &attestation, nil
}

//nolint:cyclop
func (n Notary) loadSummitMyLatestStates(ctx context.Context) {
	for _, domain := range n.domains {
		originID := domain.Config().DomainID
		myLatestState, err := n.summitDomain.Summit().GetLatestAgentState(ctx, originID, n.bondedSigner)
		if err != nil {
			myLatestState = nil
			logger.Errorf("Failed calling GetLatestAgentState for originID on the Summit contract: %d, err = %v", originID, err)
		}
		if myLatestState != nil && myLatestState.Nonce() > uint32(0) {
			n.summitMyLatestStates[originID] = myLatestState
		}
	}
}

//nolint:cyclop
func (n Notary) loadSummitGuardLatestStates(ctx context.Context) {
	for _, domain := range n.domains {
		originID := domain.Config().DomainID

		guardLatestState, err := n.summitDomain.Summit().GetLatestState(ctx, originID)
		if err != nil {
			guardLatestState = nil
			logger.Errorf("Failed calling GetLatestState for originID %d on the Summit contract: err = %v", originID, err)
		}
		if guardLatestState != nil && guardLatestState.Nonce() > uint32(0) {
			n.summitGuardLatestStates[originID] = guardLatestState
		}
	}
}

//nolint:cyclop
func (n Notary) isValidOnOrigin(ctx context.Context, state types.State, domain domains.DomainClient) bool {
	stateOnOrigin, err := domain.Origin().SuggestState(ctx, state.Nonce())
	if err != nil {
		logger.Errorf("Failed calling SuggestState for originID %d on the Origin contract: err = %v", err)
		// return false since we weren't able to validate the state on the origin
		return false
	}

	if stateOnOrigin.Root() != state.Root() {
		logger.Errorf("State roots do not equal")
		return false
	}

	if stateOnOrigin.Origin() != state.Origin() {
		logger.Errorf("State origins do not equal")
		return false
	}

	if stateOnOrigin.Nonce() != state.Nonce() {
		logger.Errorf("State nonces do not equal")
		return false
	}

	if stateOnOrigin.BlockNumber() == nil {
		logger.Errorf("State on origin had nil block number")
		return false
	}

	if state.BlockNumber() == nil {
		logger.Errorf("State to validate had nil block number")
		return false
	}

	if stateOnOrigin.BlockNumber().Uint64() != state.BlockNumber().Uint64() {
		logger.Errorf("State block numbers do not equal")
		return false
	}

	if stateOnOrigin.Timestamp() == nil {
		logger.Errorf("State on origin had nil time stamp")
		return false
	}

	if state.Timestamp() == nil {
		logger.Errorf("State to validate had nil time stamp")
		return false
	}

	if stateOnOrigin.Timestamp().Uint64() != state.Timestamp().Uint64() {
		logger.Errorf("State timestamps do not equal")
		return false
	}

	stateOnOriginHash, err := stateOnOrigin.Hash()
	if err != nil {
		logger.Errorf("Error computing state on origin hash")
		return false
	}

	stateHash, err := state.Hash()
	if err != nil {
		logger.Errorf("Error computing state on summit hash")
		return false
	}

	if stateOnOriginHash != stateHash {
		logger.Errorf("State hashes do not equal")
		return false
	}

	return true
}

//nolint:cyclop
func (n Notary) getLatestSnapshot(ctx context.Context) (types.Snapshot, map[uint32]types.State) {
	statesToSubmit := make(map[uint32]types.State, len(n.domains))
	for _, domain := range n.domains {
		originID := domain.Config().DomainID
		summitMyLatest, ok := n.summitMyLatestStates[originID]
		if !ok || summitMyLatest == nil || summitMyLatest.Nonce() == 0 {
			summitMyLatest = nil
		}
		summitGuardLatest, ok := n.summitGuardLatestStates[originID]
		if !ok || summitGuardLatest == nil || summitGuardLatest.Nonce() == 0 {
			continue
		}

		if summitMyLatest != nil && summitMyLatest.Nonce() >= summitGuardLatest.Nonce() {
			// Here this notary already submitted this state
			continue
		}
		if !n.isValidOnOrigin(ctx, summitGuardLatest, domain) {
			logger.Errorf("State not valid on origin %d, nonce %d",
				summitGuardLatest.Origin(),
				summitGuardLatest.Nonce())
			continue
		}
		statesToSubmit[originID] = summitGuardLatest
	}
	snapshotStates := make([]types.State, 0, len(statesToSubmit))
	for _, state := range statesToSubmit {
		if state.Nonce() == 0 {
			continue
		}
		snapshotStates = append(snapshotStates, state)
	}
	if len(snapshotStates) > 0 {
		return types.NewSnapshot(snapshotStates), statesToSubmit
	}
	//nolint:nilnil
	return nil, nil
}

//nolint:cyclop
func (n Notary) submitLatestSnapshot(ctx context.Context) {
	snapshot, statesToSubmit := n.getLatestSnapshot(ctx)
	if snapshot == nil {
		return
	}

	snapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(ctx, n.bondedSigner)
	if err != nil {
		logger.Errorf("Error signing snapshot: %v", err)
	} else {
		logger.Infof("Notary submitting snapshot to summit")
		err := n.summitDomain.Summit().SubmitSnapshot(ctx, n.unbondedSigner, encodedSnapshot, snapshotSignature)
		if err != nil {
			logger.Errorf("Failed to submit snapshot to summit: %v", err)
		} else {
			for originID, state := range statesToSubmit {
				n.summitMyLatestStates[originID] = state
			}
		}
	}
}

//nolint:cyclop,unused
func (n Notary) submitAttestation(ctx context.Context, attBytes []byte) {
	hashedBytes, err := types.HashRawBytes(attBytes)
	if err != nil {
		logger.Errorf("could not hash attBytes: %w", err)
		return
	}
	signature, err := n.bondedSigner.SignMessage(ctx, core.BytesToSlice(hashedBytes), false)
	if err != nil {
		logger.Errorf("could not sign snapshot: %w", err)
		return
	}

	err = n.destinationDomain.Destination().SubmitAttestation(ctx, n.unbondedSigner, attBytes, signature)
	if err != nil {
		logger.Errorf("Failed to submit snapshot to summit: %v", err)
	}
}

// Start starts the notary.
//
//nolint:cyclop
func (n Notary) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	logger.Info("Starting the notary")

	/*attestationSavedSink := make(chan *summit.SummitAttestationSaved)
	var savedAttestation event.Subscription

	var err error
	savedAttestation, err = n.summitDomain.Summit().WatchAttestationSaved(ctx, attestationSavedSink)
	if err != nil {
		return fmt.Errorf("error setting up watcher for saved attestations: %w", err)
	}*/

	logger.Infof("Notary loadSummitMyLatestStates")
	n.loadSummitMyLatestStates(ctx)

	// First initialize a map to track what was the last state signed by this notary

	// TODO (add scribe listener for AttestationSaved events)
	// Whenever we get an event, the Notary would want to sign and submit to destination.
	// For MVP, its fine to just sign and submit.
	// Later, there will be validating the actual states.
	// On summit, the Notary will pass in the attestation payload to get the raw states associated with it.
	// It will then double check all the states on origin.
	// Then, it will sign and submit
	/*g.Go(func() error {
		watchCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		for {
			select {
			// check for errors and fail
			case <-watchCtx.Done():
				logger.Info("Notary Attestation Saved Watcher exiting without error")
				return nil
			case <-savedAttestation.Err():
				logger.Info("Notary Attestation Saved Watcher got an unexpected error: %v", savedAttestation.Err())
			// get message sent event
			case receivedAttestationSaved := <-attestationSavedSink:
				logger.Info("Notary received a saved attestation event, will sign and submit to destination")
				attToSubmit := receivedAttestationSaved.Attestation
				n.submitAttestation(ctx, attToSubmit)
			}
		}
	})*/

	g.Go(func() error {
		for {
			select {
			// parent loop terminated
			case <-ctx.Done():
				logger.Info("Notary exiting without error")
				return nil
			case <-time.After(n.refreshInterval):
				n.loadSummitGuardLatestStates(ctx)
				n.submitLatestSnapshot(ctx)
				logger.Info("Notary exiting after 1 run")
			}
		}
	})

	err := g.Wait()
	if err != nil {
		logger.Errorf("Notary exiting with error: %v", err)
		return fmt.Errorf("could not start the notary: %w", err)
	}

	logger.Info("Notary exiting without error")
	return nil
}
