package config

import (
	"context"
	"fmt"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/synapsecns/sanguine/ethergo/parser"
	"golang.org/x/exp/slices"
	"io/fs"
	"os"
	"strconv"
	"time"
)

// ClientGenerator generates an ethclient from a context and a url, this is used so we can override
// ethclient.DialContext for testing.
type ClientGenerator func(ctx context.Context, rawUrl string) (ReceiptClient, error)

// ReceiptClient is an client that implements receipt fetching.
type ReceiptClient interface {
	// TransactionReceipt returns the receipt of a mined transaction. Note that the
	// transaction may not be included in the current canonical chain even if a receipt
	// exists.
	// TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// DefaultClientGenerator generates the deafult ethclient.
func DefaultClientGenerator(ctx context.Context, rawUrl string) (ReceiptClient, error) {
	return ethclient.DialContext(ctx, rawUrl)
}

// GenerateConfig generates a config using a hardhat deployment and scribe.
// this requires scribe to be live.
func GenerateConfig(ctx context.Context, omniRPCUrl, deployPath string, requiredConfirmations uint32, outputPath string, skippedChainIDS []int, cg ClientGenerator) error {
	contracts, err := parser.GetDeployments(deployPath)
	if err != nil {
		return fmt.Errorf("could not get deployments: %w", err)
	}

	configList := make(map[int][]ContractConfig)

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

			rpcClient, err := ethclient.DialContext(ctx, fmt.Sprintf("%s/rpc/%d", omniRPCUrl, chainID))
			if err != nil {
				return fmt.Errorf("could not get client: %w", err)
			}

			var txReceipt *types.Receipt

		OUTER:
			for attempt := 0; attempt < 20; attempt++ {
				txReceipt, err = rpcClient.TransactionReceipt(ctx, common.HexToHash(network.TransactionHash))
				if err != nil {
					if attempt < 20 {
						_ = awsTime.SleepWithContext(ctx, time.Second*2)
						continue
					} else {
						return fmt.Errorf("could not get tx receipt: %w", err)
					}
				} else {
					break OUTER
				}
			}

			configList[chainID] = append(configList[chainID], ContractConfig{
				Address:    network.Address,
				StartBlock: txReceipt.BlockNumber.Uint64(),
			})
		}
	}

	config := Config{}
	for chainID, chainContracts := range configList {
		config.Chains = append(config.Chains, ChainConfig{
			ChainID:               uint32(chainID),
			RPCUrl:                fmt.Sprintf("%s/rpc/%d", omniRPCUrl, chainID),
			RequiredConfirmations: requiredConfirmations,
			Contracts:             chainContracts,
		})
	}

	encodedConfig, err := config.Encode()
	if err != nil {
		return fmt.Errorf("could not create encoded config: %w", err)
	}

	err = os.WriteFile(outputPath, encodedConfig, fs.ModeType)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	return nil
}
