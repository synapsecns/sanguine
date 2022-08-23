package backends

import (
	"github.com/synapsecns/sanguine/ethergo/deployer"
)

// ContractVerifier is a hook used to verify contracts with a test provider.
type ContractVerifier interface {
	VerifyContract(contractType deployer.ContractType, contract DeployedContract) error
}
