package sleuth

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/sleuth/config"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	ethergoChain "github.com/synapsecns/sanguine/ethergo/chain"
	"math/big"
)

// Sleuth is the sleuth agent.
type Sleuth struct {
	// config is the sleuth config.
	config config.Config
	// clients is a map of chain ID -> chain client.
	clients map[uint32]Backend
	// originParsers is a map of chain ID -> origin parser.
	originParsers map[uint32]origin.Parser
	// summit is the summit contract.
	summit domains.SummitContract
	// destinations is a map of chain ID -> destination contract.
	destinations map[uint32]domains.DestinationContract
}

// NewSleuth creates a new sleuth agent.
func NewSleuth(ctx context.Context, config config.Config, clients map[uint32]Backend) (*Sleuth, error) {
	var sleuth Sleuth

	sleuth.clients = clients
	sleuth.config = config
	sleuth.originParsers = make(map[uint32]origin.Parser)
	sleuth.destinations = make(map[uint32]domains.DestinationContract)

	for _, chain := range config.Chains {
		if _, ok := sleuth.clients[chain.ChainID]; !ok {
			return nil, fmt.Errorf("chain %d does not have a client", chain.ChainID)
		}

		chainOriginParser, err := origin.NewParser(common.HexToAddress(chain.OriginAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create origin parser: %w", err)
		}

		sleuth.originParsers[chain.ChainID] = chainOriginParser

		// chainRPCURL := fmt.Sprintf("%s/1/rpc/%d", config.BaseOmnirpcURL, chain.ChainID)
		//
		underlyingClient, err := ethergoChain.NewFromURL(ctx, chain.TempRPC)
		if err != nil {
			return nil, fmt.Errorf("could not get evm: %w", err)
		}

		if config.SummitChainID == chain.ChainID {
			summit, err := evm.NewSummitContract(ctx, underlyingClient, common.HexToAddress(config.SummitAddress))
			if err != nil {
				return nil, fmt.Errorf("could not create summit contract: %w", err)
			}

			sleuth.summit = summit
		}

		destination, err := evm.NewDestinationContract(ctx, underlyingClient, common.HexToAddress(chain.DestinationAddress))
		if err != nil {
			return nil, fmt.Errorf("could not create destination contract: %w", err)
		}

		sleuth.destinations[chain.ChainID] = destination
	}

	return &sleuth, nil
}

// checkMessage gets the message to investigate.
func (s *Sleuth) checkMessage(ctx context.Context, txHash common.Hash, chainID uint32) (*types.Message, error) {
	receipt, err := s.clients[chainID].TransactionReceipt(ctx, txHash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			//nolint:nilnil
			return nil, nil
		}
		return nil, fmt.Errorf("could not get receipt: %w", err)
	}

	for _, log := range receipt.Logs {
		if log == nil {
			continue
		}

		if log.Address == common.HexToAddress(s.config.SummitAddress) {
			committedMessage, ok := s.originParsers[chainID].ParseDispatched(*log)
			if !ok {
				continue
			}

			message, err := types.DecodeMessage(committedMessage.Message())
			if err != nil {
				return nil, fmt.Errorf("could not decode message: %w", err)
			}

			return &message, nil
		}
	}

	//nolint:nilnil
	return nil, nil
}

// checkState sees if a message has been included in a state.
func (s *Sleuth) checkState(ctx context.Context, nonce, chainID uint32) (bool, error) {
	latestState, err := s.summit.GetLatestState(ctx, chainID)
	if err != nil {
		return false, fmt.Errorf("could not get latest state: %w", err)
	}

	if latestState == nil {
		return false, nil
	}

	if latestState.Nonce() < nonce {
		return false, nil
	}

	return true, nil
}

// checkSnapshot sees if a message has been included in a snapshot.
func (s *Sleuth) checkSnapshot(ctx context.Context, nonce, chainID uint32) (*[][32]byte, error) {
	var snapshotRoots [][32]byte

	latestSnapshotNonce, err := s.findLatestSnapshotNonce(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get latest snapshot nonce: %w", err)
	}

	if new(big.Int).Mul(latestSnapshotNonce, big.NewInt(0)) == big.NewInt(0) {
		//nolint:nilnil
		return nil, nil
	}

	for i := latestSnapshotNonce; i.Cmp(big.NewInt(0)) >= 0; i.Sub(i, big.NewInt(1)) {
		snapshot, err := s.summit.GetNotarySnapshot(ctx, i)
		if err != nil {
			return nil, fmt.Errorf("could not get snapshot: %w", err)
		}
		if snapshot == nil {
			return nil, fmt.Errorf("no snapshot for nonce %s", i.String())
		}

		for _, state := range (*snapshot).States() {
			if state.Origin() == chainID {
				if state.Nonce() >= nonce {
					snapshotRoot, _, err := (*snapshot).SnapshotRootAndProofs()
					if err != nil {
						return nil, fmt.Errorf("could not get snapshot root: %w", err)
					}

					snapshotRoots = append(snapshotRoots, snapshotRoot)
				} else {
					return &snapshotRoots, nil
				}
			}
		}
	}

	return &snapshotRoots, nil
}

// checkAttestation sees if a message has been involved in an attestation.
func (s *Sleuth) checkAttestation(ctx context.Context, snapshotRoots [][32]byte, destinationDomain uint32) (bool, error) {
	// Convert the snapshot roots to a map.
	snapshotRootsMap := make(map[[32]byte]bool)
	for _, snapshotRoot := range snapshotRoots {
		snapshotRootsMap[snapshotRoot] = true
	}

	attestationsAmount, err := s.destinations[destinationDomain].AttestationsAmount(ctx)
	if err != nil {
		return false, fmt.Errorf("could not get attestations amount: %w", err)
	}

	for i := attestationsAmount - 1; i >= 0; i-- {
		root, _, err := s.destinations[destinationDomain].GetAttestation(ctx, i)
		if err != nil {
			return false, fmt.Errorf("could not get attestation: %w", err)
		}

		if snapshotRootsMap[root] {
			return true, nil
		}
	}

	return false, nil
}

// checkExecuted sees if a message has been executed.
func (s *Sleuth) checkExecuted(ctx context.Context, messageHash [32]byte, destinationDomain uint32) (bool, error) {
	executed, err := s.destinations[destinationDomain].MessageStatus(ctx, messageHash)
	if err != nil {
		return false, fmt.Errorf("could not get message status: %w", err)
	}

	return executed, nil
}

func (s *Sleuth) findLatestSnapshotNonce(ctx context.Context) (*big.Int, error) {
	one := big.NewInt(1)

	high := new(big.Int).Sub(new(big.Int).Lsh(one, 256), one)
	low := big.NewInt(1)
	lastNonNilNonce := big.NewInt(0)

	for low.Cmp(high) <= 0 {
		mid := new(big.Int).Add(low, new(big.Int).Rsh(new(big.Int).Sub(high, low), 1))
		snapshot, err := s.summit.GetNotarySnapshot(ctx, mid)
		if err != nil {
			return nil, fmt.Errorf("could not get snapshot: %w", err)
		}

		if snapshot != nil {
			lastNonNilNonce = mid
			low = new(big.Int).Add(mid, one)
		} else {
			high = new(big.Int).Sub(mid, one)
		}
	}

	return lastNonNilNonce, nil
}
