package config_test

import (
	"context"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/integralist/go-findroot/find"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/parser/hardhat"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"math/big"
	"path/filepath"
	"strconv"
	"testing"
)

// TestClient returns fake data.
type TestClient struct {
	txes  map[common.Hash]*big.Int
	count uint64
}

// NewTestClient creates a new test client, this will be used for mocking big int returns and verifying
// the start block number matches the tx receipt block.
func NewTestClient() *TestClient {
	return &TestClient{txes: make(map[common.Hash]*big.Int), count: 60000}
}

func (t *TestClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	t.count--
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("could not complete request: %w", ctx.Err())
	default:
		// continue
	}

	t.txes[txHash] = new(big.Int).SetUint64(t.count)
	return &types.Receipt{
		// only used field in this config
		BlockNumber: t.txes[txHash],
	}, nil
}

var _ config.ReceiptClient = &TestClient{}

func (c *ConfigSuite) TestGenerateConfig() {
	testClient := NewTestClient()

	testGenerator := func(ctx context.Context, rawUrl string) (config.ReceiptClient, error) {
		return testClient, nil
	}

	repoRoot, err := find.Repo()
	Nil(c.T(), err)
	omnirpcURL := gofakeit.URL()
	outputPath := filepath.Join(filet.TmpDir(c.T(), ""), "test.yaml")

	deploymentsFolder := filepath.Join(repoRoot.Path, "ethergo", "internal", "test-data", "deployments")

	err = config.GenerateConfig(c.GetTestContext(), omnirpcURL, deploymentsFolder, outputPath, []int{5, 335, 43113, 1666700000}, testGenerator)
	Nil(c.T(), err)

	parsedDeployments, err := hardhat.GetDeployments(deploymentsFolder)
	Nil(c.T(), err)

	decodedConfig, err := config.DecodeConfig(outputPath)
	Nil(c.T(), err)

	for _, chain := range decodedConfig.Chains {
		for _, contract := range chain.Contracts {
			txHash := getDeploymentsTxHash(c.T(), parsedDeployments, chain.ChainID, common.HexToAddress(contract.Address))
			mockStartBlock := testClient.txes[txHash]
			LessOrEqual(c.T(), mockStartBlock.Uint64(), contract.StartBlock)
		}
	}
}

// getDeploymentsTxHash gets the tx receipt from a list of deployments by address and chainid.
func getDeploymentsTxHash(tb testing.TB, deployments []hardhat.Contract, chainID uint32, contractAddress common.Address) common.Hash {
	tb.Helper()

	for _, deployment := range deployments {
		// make sure contract belongs to this chain id
		network, ok := deployment.Networks[strconv.Itoa(int(chainID))]
		if !ok {
			continue
		}

		if common.HexToAddress(network.Address) == contractAddress {
			return common.HexToHash(network.TransactionHash)
		}
	}

	Fail(tb, fmt.Sprintf("no contract found for chain id %d and contract address %s", chainID, contractAddress))
	return common.Hash{}
}
