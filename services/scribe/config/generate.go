package config

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/parser/hardhat"
	"os"
	"strconv"
	"time"

	"bitbucket.org/tentontrain/math"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slices"
)

// ClientGenerator generates an ethclient from a context and a url, this is used so we can override
// ethclient.DialContext for testing.
type ClientGenerator func(ctx context.Context, rawURL string) (ReceiptClient, error)

// ReceiptClient is an client that implements receipt fetching.
type ReceiptClient interface {
	// TransactionReceipt returns the receipt of a mined transaction. Note that the
	// transaction may not be included in the current canonical chain even if a receipt
	// exists.
	// TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// DefaultClientGenerator generates the default ethclient.
func DefaultClientGenerator(ctx context.Context, rawURL string) (ReceiptClient, error) {
	//nolint: wrapcheck
	return ethclient.DialContext(ctx, rawURL)
}

type configList map[int]map[string]ContractConfig

// ContractsForChain gets all contraacts for a given chain.
func (c configList) ContractsForChain(chainID int) (configs []ContractConfig) {
	chainConfigs, ok := c[chainID]
	if !ok {
		return configs
	}

	for _, contractConfigs := range chainConfigs {
		configs = append(configs, contractConfigs)
	}
	return configs
}

// GenerateConfig generates a config using a hardhat deployment and scribe.
// this requires scribe to be live.
//
//nolint:cyclop
func GenerateConfig(ctx context.Context, omniRPCUrl, deployPath string, outputPath string, skippedChainIDS []int, cg ClientGenerator) error {
	contracts, err := hardhat.GetDeployments(deployPath)
	if err != nil {
		return fmt.Errorf("could not get deployments: %w", err)
	}

	configList := make(configList)

	for _, contract := range contracts {
		for chainIDStr, network := range contract.Networks {
			chainID, err := strconv.Atoi(chainIDStr)
			if err != nil {
				return fmt.Errorf("could not parse chainid from string %s: %w", chainIDStr, err)
			}

			// skip chainids
			if slices.Contains(skippedChainIDS, chainID) {
				continue
			}

			deployBlock, err := getDeployBlock(ctx, cg, omniRPCUrl, common.HexToHash(network.TransactionHash), uint32(chainID))
			if err != nil {
				return fmt.Errorf("could not get deploy block for contract %s on network %d", contract.Name, chainID)
			}

			// initialize the chain map
			_, hasChain := configList[chainID]
			if !hasChain {
				configList[chainID] = make(map[string]ContractConfig)
			}

			chainContract, hasContract := configList[chainID][network.Address]
			// if the contract already exist, just use lesser of the two start blocks
			if hasContract {
				chainContract.StartBlock = math.Min[uint64](deployBlock, chainContract.StartBlock)
				configList[chainID][network.Address] = chainContract
				continue
			}

			configList[chainID][network.Address] = ContractConfig{
				Address:    network.Address,
				StartBlock: deployBlock,
			}
		}
	}

	config := Config{}
	for chainID := range configList {
		config.Chains = append(config.Chains, ChainConfig{
			ChainID:   uint32(chainID),
			Contracts: configList.ContractsForChain(chainID),
		})
	}

	encodedConfig, err := config.Encode()
	if err != nil {
		return fmt.Errorf("could not create encoded config: %w", err)
	}

	err = os.WriteFile(outputPath, encodedConfig, 0600)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	return nil
}

func getDeployBlock(ctx context.Context, cg ClientGenerator, omniRPCUrl string, txhash common.Hash, chainID uint32) (block uint64, err error) {
	rpcClient, err := cg(ctx, fmt.Sprintf("%s/rpc/%d", omniRPCUrl, chainID))
	if err != nil {
		return block, fmt.Errorf("could not get client: %w", err)
	}

	var txReceipt *types.Receipt

OUTER:
	for attempt := 0; attempt < 20; attempt++ {
		txReceipt, err = rpcClient.TransactionReceipt(ctx, txhash)
		if err != nil {
			if attempt < 20 {
				_ = awsTime.SleepWithContext(ctx, time.Second*2)
				continue
			}
			return block, fmt.Errorf("could not get tx receipt: %w", err)
		}
		break OUTER
	}

	return txReceipt.BlockNumber.Uint64(), nil
}
