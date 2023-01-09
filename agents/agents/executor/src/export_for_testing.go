package src

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	agentsConfig "github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/merkle"
	ethergoChain "github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetLogChan gets a log channel.
func (e Executor) GetLogChan(chainID uint32) chan *ethTypes.Log {
	return e.chainExecutors[chainID].logChan
}

// GetMerkleTree gets a merkle tree.
func (e Executor) GetMerkleTree(chainID uint32, domain uint32) *merkle.HistoricalTree {
	return e.chainExecutors[chainID].merkleTrees[domain]
}

// VerifyMessageMerkleProof verifies message merkle proof.
func (e Executor) VerifyMessageMerkleProof(message types.Message) (bool, error) {
	return e.verifyMessageMerkleProof(message)
}

// VerifyMessageOptimisticPeriod verifies message optimistic period.
func (e Executor) VerifyMessageOptimisticPeriod(ctx context.Context, message types.Message) (*uint32, error) {
	return e.verifyMessageOptimisticPeriod(ctx, message)
}

// NewTreeFromDB builds a merkle tree from the db.
func NewTreeFromDB(ctx context.Context, chainID uint32, domain uint32, executorDB db.ExecutorDB) (*merkle.HistoricalTree, error) {
	return newTreeFromDB(ctx, chainID, domain, executorDB)
}

// OverrideMerkleTree overrides the merkle tree for the chainID and domain.
func (e Executor) OverrideMerkleTree(chainID uint32, domain uint32, tree *merkle.HistoricalTree) {
	e.chainExecutors[chainID].merkleTrees[domain] = tree
}

// Start runs the executor.
func (e Executor) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain, originContract)
		})

		g.Go(func() error {
			return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chain, destinationContract)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error when streaming logs: %w", err)
	}

	return nil
}

// Listen scans for emitted logs from the various chains.
func (e Executor) Listen(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	for _, chain := range e.config.Chains {
		chain := chain

		g.Go(func() error {
			return e.receiveLogs(ctx, chain.ChainID)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error when receiving logs: %w", err)
	}

	return nil
}

// NewExecutorInjectedBackend creates a new Executor suitable for testing since it does not need a valid omnirpcURL.
//
//nolint:cyclop
func NewExecutorInjectedBackend(ctx context.Context, config config.Config, executorDB db.ExecutorDB, scribeClient client.ScribeClient, clients map[uint32]Backend, urls map[uint32]string) (*Executor, error) {
	chainExecutors := make(map[uint32]*chainExecutor)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	executorSigner, err := agentsConfig.SignerFromConfig(config.UnbondedSigner)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
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

		underlyingClient, err := ethergoChain.NewFromURL(ctx, urls[chain.ChainID])
		if err != nil {
			return nil, fmt.Errorf("could not get evm: %w", err)
		}

		boundDestination, err := evm.NewDestinationContract(ctx, underlyingClient, common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not bind destination contract: %w", err)
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
			logChan:           make(chan *ethTypes.Log, logChanSize),
			merkleTrees:       make(map[uint32]*merkle.HistoricalTree),
			rpcClient:         clients[chain.ChainID],
			boundDestination:  boundDestination,
		}

		for _, destinationChain := range config.Chains {
			if destinationChain.ChainID == chain.ChainID {
				continue
			}

			tree, err := newTreeFromDB(ctx, chain.ChainID, destinationChain.ChainID, executorDB)
			if err != nil {
				return nil, fmt.Errorf("could not get tree from db: %w", err)
			}

			chainExecutors[chain.ChainID].merkleTrees[destinationChain.ChainID] = tree
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
	}, nil
}
