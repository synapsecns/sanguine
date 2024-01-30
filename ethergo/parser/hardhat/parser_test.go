package hardhat_test

import (
	"github.com/integralist/go-findroot/find"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/internal/testconsts"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/parser/hardhat"
	"path/filepath"
	"testing"
)

func TestParser(t *testing.T) {
	repoRoot, err := find.Repo()
	Nil(t, err)

	deploymentsFolder := filepath.Join(repoRoot.Path, "ethergo", "internal", "test-data", "deployments")
	contracts, err := hardhat.GetDeployments(deploymentsFolder)
	Nil(t, err)

	for _, contract := range contracts {
		True(t, len(contract.Networks) >= 1)
		for _, network := range contract.Networks {
			NotEmpty(t, network.TransactionHash)
			NotEmpty(t, network.Address)
		}
	}
}

func TestGetNetworks(t *testing.T) {
	repoRoot, err := find.Repo()
	Nil(t, err)

	address := mocks.MockAddress()

	contractDir := filepath.Join(repoRoot.Path, "ethergo", "internal", "test-data", "deployments", "arbitrum")
	networks, err := hardhat.GetNetworks(hardhat.HardhatContract{
		Contract: hardhat.Contract{},
		Address:  address.String(),
	}, contractDir)
	Nil(t, err)
	res, ok := networks["42161"]
	True(t, ok)

	Equal(t, res.Address, address.String())
}

func TestGetMetadata(t *testing.T) {
	metadata, err := hardhat.GetMetadata(hardhat.HardhatContract{Metadata: string(testconsts.AmplificationUtilsMetadata)}, "")
	Nil(t, err)

	Equal(t, metadata.GetCompiler().Version, "0.6.12+commit.27d51765")
	sources := []string{
		"@openzeppelin/contracts/token/ERC20/SafeERC20.sol",
		"contracts/MathUtils.sol",
		"contracts/SwapUtils.sol",
		"@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol",
		"@openzeppelin/contracts-upgradeable/proxy/Initializable.sol",
		"@openzeppelin/contracts-upgradeable/token/ERC20/ERC20BurnableUpgradeable.sol",
		"contracts/LPToken.sol",
	}

	Equal(t, len(metadata.GetSources()), 19)
	for _, source := range sources {
		_, ok := metadata.GetSources()[source]
		True(t, ok)
	}
}
