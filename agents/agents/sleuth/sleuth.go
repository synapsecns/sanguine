package sleuth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/sleuth/config"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	ethergoChain "github.com/synapsecns/sanguine/ethergo/chain"
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
func NewSleuth(ctx context.Context, config config.Config, clients map[uint32]Backend) (sleuth *Sleuth, err error) {
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

	return
}

// getMessage gets the message to investigate.
func (s *Sleuth) getMessage(ctx context.Context, txHash common.Hash, chainID uint32) (*types.Message, error) {
	receipt, err := s.clients[chainID].TransactionReceipt(ctx, txHash)
	if err != nil {
		if err == ethereum.NotFound {
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

	return nil, nil
}
