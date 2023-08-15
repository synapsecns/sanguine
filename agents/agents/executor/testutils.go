package executor

import (
	"context"
	"fmt"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/merkle"
	"golang.org/x/sync/errgroup"
)

// NewTreeFromDB builds a merkle tree from the db.
func NewTreeFromDB(ctx context.Context, chainID uint32, executorDB db.ExecutorDB) (*merkle.HistoricalTree, error) {
	return newTreeFromDB(ctx, chainID, executorDB)
}

// GetLogChan gets a log channel.
func (e Executor) GetLogChan(chainID uint32) chan *ethTypes.Log {
	return e.chainExecutors[chainID].logChan
}

// StartAndListenOrigin starts and listens to a chain.
func (e Executor) StartAndListenOrigin(ctx context.Context, chainID uint32, address string) error {
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		return e.streamLogs(ctx, e.grpcClient, e.grpcConn, chainID, address, execTypes.OriginContract)
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

// CheckIfExecuted checks if a message has been executed.
func (e Executor) CheckIfExecuted(ctx context.Context, message types.Message) (bool, error) {
	return e.checkIfExecuted(ctx, message)
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
