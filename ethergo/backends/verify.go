package backends

import (
	"github.com/synapsecns/sanguine/ethergo/contracts"
)

// ContractVerifier is a hook used to verify contracts with a test provider.
type ContractVerifier interface {
	VerifyContract(contractType contracts.ContractType, contract contracts.DeployedContract) error
}
