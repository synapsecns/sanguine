package tenderly

import (
	"fmt"
	"os"

	"github.com/tenderly/tenderly-cli/model"
	"github.com/tenderly/tenderly-cli/providers"
	"github.com/tenderly/tenderly-cli/truffle"
)

// LocalTXDeploymentProviderName is the custom synapse deployment provider.
const LocalTXDeploymentProviderName providers.DeploymentProviderName = "LocalTXDeploymentProvider"

// LocalTXDeploymentProvider provides the development provider for synapse.
type LocalTXDeploymentProvider struct {
	*truffle.DeploymentProvider
}

// GetConfig gets the config for the synapse deployment TODO the rest of this probably needs to be configured somehow.
func (p *LocalTXDeploymentProvider) GetConfig(configName string, configDir string) (*providers.Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not get working directory: %w", err)
	}
	return &providers.Config{
		ProjectDirectory: dir,
		BuildDirectory:   dir,
		Networks:         nil,
		Solc:             nil,
		Compilers:        nil,
		ConfigType:       "",
		Paths:            providers.Paths{},
	}, nil
}

// MustGetConfig does nothing but conform to the interface.
//
//nolint:wrapcheck
func (p *LocalTXDeploymentProvider) MustGetConfig() (*providers.Config, error) {
	return p.DeploymentProvider.MustGetConfig()
}

// CheckIfProviderStructure checks if the provider is currently this provider, which it always is
// note: there may be some issues here if multiple providers answer true (idk what happens)
func (p *LocalTXDeploymentProvider) CheckIfProviderStructure(directory string) bool {
	return true
}

// GetContracts gets contracts currently deployed on the provider. Since this is for unit tests, this is none always since this method is only called on initialization.
func (p *LocalTXDeploymentProvider) GetContracts(buildDir string, networkIDs []string, objects ...*model.StateObject) ([]providers.Contract, int, error) {
	return []providers.Contract{}, 0, nil
}

// NewDeploymentProvider creates a provider for this project.
func NewDeploymentProvider() *LocalTXDeploymentProvider {
	return &LocalTXDeploymentProvider{
		truffle.NewDeploymentProvider(),
	}
}

var _ providers.DeploymentProvider = (*LocalTXDeploymentProvider)(nil)

// GetProviderName of the LocalTXDeploymentProvider.
func (*LocalTXDeploymentProvider) GetProviderName() providers.DeploymentProviderName {
	return LocalTXDeploymentProviderName
}
