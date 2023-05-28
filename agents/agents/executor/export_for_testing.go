package executor

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	agentsConfig "github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/summit"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/merkle"
	"github.com/synapsecns/sanguine/core/metrics"
	ethergoChain "github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// -------- [ UTILS ] -------- \\

// NewExecutorInjectedBackend creates a new Executor suitable for testing since it does not need a valid omnirpcURL.
//
//nolint:cyclop
func NewExecutorInjectedBackend(ctx context.Context, config config.Config, executorDB db.ExecutorDB, scribeClient client.ScribeClient, clients map[uint32]Backend, urls map[uint32]string, handler metrics.Handler) (*Executor, error) {
	chainExecutors := make(map[uint32]*chainExecutor)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not dial grpc: %w", err)
	}

	grpcClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := grpcClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return nil, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return nil, fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	executorSigner, err := agentsConfig.SignerFromConfig(ctx, config.UnbondedSigner)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
	}

	if config.ExecuteInterval == 0 {
		config.ExecuteInterval = 2
	}

	if config.SetMinimumTimeInterval == 0 {
		config.SetMinimumTimeInterval = 2
	}

	for _, chain := range config.Chains {
		originParser, err := origin.NewParser(common.HexToAddress(chain.OriginAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create origin parser: %w", err)
		}

		destinationParser, err := destination.NewParser(common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create destination parser: %w", err)
		}

		var summitParserRef *summit.Parser
		var inboxParserRef *inbox.Parser
		var lightInboxParserRef *lightinbox.Parser

		if config.SummitChainID == chain.ChainID {
			summitParser, err := summit.NewParser(common.HexToAddress(config.SummitAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create summit parser: %w", err)
			}

			summitParserRef = &summitParser

			inboxParser, err := inbox.NewParser(common.HexToAddress(config.InboxAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create bonding manager parser: %w", err)
			}

			inboxParserRef = &inboxParser
		} else {
			lightInboxParser, err := lightinbox.NewParser(common.HexToAddress(chain.LightInboxAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create destination parser: %w", err)
			}

			lightInboxParserRef = &lightInboxParser
		}

		underlyingClient, err := ethergoChain.NewFromURL(ctx, urls[chain.ChainID])
		if err != nil {
			return nil, fmt.Errorf("could not get evm: %w", err)
		}

		boundDestination, err := evm.NewDestinationContract(ctx, underlyingClient, common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind destination contract: %w", err)
		}

		tree, err := newTreeFromDB(ctx, chain.ChainID, executorDB)
		if err != nil {
			return nil, fmt.Errorf("could not get tree from db: %w", err)
		}

		chainExecutors[chain.ChainID] = &chainExecutor{
			chainID: chain.ChainID,
			lastLog: &logOrderInfo{
				blockNumber: 0,
				blockIndex:  0,
			},
			closeConnection:   make(chan bool, 1),
			stopListenChan:    make(chan bool, 1),
			originParser:      originParser,
			destinationParser: destinationParser,
			summitParser:      summitParserRef,
			lightInboxParser:  lightInboxParserRef,
			inboxParser:       inboxParserRef,
			logChan:           make(chan *ethTypes.Log, logChanSize),
			merkleTree:        tree,
			rpcClient:         clients[chain.ChainID],
			boundDestination:  boundDestination,
			executed:          make(map[[32]byte]bool),
		}
	}

	return &Executor{
		config:         config,
		executorDB:     executorDB,
		scribeClient:   scribeClient,
		grpcConn:       conn,
		grpcClient:     grpcClient,
		signer:         executorSigner,
		chainExecutors: chainExecutors,
		handler:        handler,
	}, nil
}

// NewTreeFromDB builds a merkle tree from the db.
func NewTreeFromDB(ctx context.Context, chainID uint32, executorDB db.ExecutorDB) (*merkle.HistoricalTree, error) {
	return newTreeFromDB(ctx, chainID, executorDB)
}

// -------- [ EXECUTOR ] -------- \\

//// SetMinimumTimes goes through a list of messages and sets the minimum time for each message
//// that has an associated attestation.
// func (e Executor) SetMinimumTimes(ctx context.Context, messages []types.Message, attestations []execTypes.DBAttestation) error {
//	return e.setMinimumTimes(ctx, messages, attestations)
//}

// GetLogChan gets a log channel.
func (e Executor) GetLogChan(chainID uint32) chan *ethTypes.Log {
	return e.chainExecutors[chainID].logChan
}

// StartAndListenOrigin starts and listens to a chain.
func (e Executor) StartAndListenOrigin(ctx context.Context, chainID uint32, address string) error {
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chainID, address, nil, contractEventType{
			contractType: originContract,
			eventType:    sentEvent,
		})
	})

	g.Go(func() error {
		return e.receiveLogs(ctx, chainID)
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error in executor agent: %w", err)
	}

	return nil
}

// GetMerkleTree gets a merkle tree.
func (e Executor) GetMerkleTree(chainID uint32) *merkle.HistoricalTree {
	return e.chainExecutors[chainID].merkleTree
}

// GetExecuted gets the executed mapping.
func (e Executor) GetExecuted(chainID uint32) map[[32]byte]bool {
	return e.chainExecutors[chainID].executed
}

// VerifyMessageMerkleProof verifies message merkle proof.
func (e Executor) VerifyMessageMerkleProof(message types.Message) (bool, error) {
	return e.verifyMessageMerkleProof(message)
}

// VerifyStateMerkleProof verifies state merkle proof.
func (e Executor) VerifyStateMerkleProof(ctx context.Context, state types.State) (bool, error) {
	return e.verifyStateMerkleProof(ctx, state)
}

// VerifyMessageOptimisticPeriod verifies message optimistic period.
func (e Executor) VerifyMessageOptimisticPeriod(ctx context.Context, message types.Message) (*uint32, error) {
	return e.verifyMessageOptimisticPeriod(ctx, message)
}

// OverrideMerkleTree overrides the merkle tree for the chainID and domain.
func (e Executor) OverrideMerkleTree(chainID uint32, tree *merkle.HistoricalTree) {
	e.chainExecutors[chainID].merkleTree = tree
}

// SetMinimumTime sets the minimum times.
func (e Executor) SetMinimumTime(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.setMinimumTime(ctx, chain.ChainID)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error when setting minimum time: %w", err)
	}

	return nil
}

// ExecuteExecutable executes executable messages in the database.
func (e Executor) ExecuteExecutable(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.executeExecutable(ctx, chain.ChainID)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error when executing executable messages: %w", err)
	}

	return nil
}
