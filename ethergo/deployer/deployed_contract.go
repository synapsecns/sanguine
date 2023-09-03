package deployer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"math/big"
)

// DeployedContract represents a deployed contract. It is returned by a deployer after a successful deployment.
type DeployedContract struct {
	// address is the address where the contract has been deployed
	address common.Address
	// contractHandle is the actual handle returned by deploying the contract
	// this must be castt o be useful
	contractHandle interface{}
	// owner of the contract
	owner common.Address
	// deployTx is the transaction where the contract was created
	deployTx *types.Transaction
	// chainID is the chain id where the contract is deployed
	chainID *big.Int
}

// NewDeployedContract creates a new deployed contract. We take some shortcuts by making some assumptions:
// namely, that tx sender is owner.
func NewDeployedContract(handle vm.ContractRef, deployTx *types.Transaction) (DeployedContract, error) {
	// TODO: eip-2930 signer?
	msg, err := deployTx.AsMessage(types.LatestSignerForChainID(deployTx.ChainId()), nil)
	if err != nil {
		return DeployedContract{}, fmt.Errorf("failed to get message from deployTx: %w", err)
	}

	return DeployedContract{
		address:        handle.Address(),
		contractHandle: handle,
		owner:          msg.From(),
		deployTx:       deployTx,
		chainID:        deployTx.ChainId(),
	}, nil
}

// Address gets the address of the deployed contract.
func (d DeployedContract) Address() common.Address {
	return d.address
}

// ContractHandle is the contract handle of the deployed ocontract.
func (d DeployedContract) ContractHandle() interface{} {
	return d.contractHandle
}

// Owner gets the contract owner.
func (d DeployedContract) Owner() common.Address {
	return d.owner
}

// OwnerPtr returns a pointer to the owner (useful for GetTxContext() operations).
func (d DeployedContract) OwnerPtr() *common.Address {
	return &d.owner
}

// DeployTx gets the deploy transaction.
func (d DeployedContract) DeployTx() *types.Transaction {
	return d.deployTx
}

// ChainID is the chain id of the deployed contract.
func (d DeployedContract) ChainID() *big.Int {
	return core.CopyBigInt(d.chainID)
}

// String returns a string representation of the contract metadata.
func (d DeployedContract) String() string {
	return fmt.Sprintf("address: %s, owner: %s, chainID: %s, deployTX: %s", d.address.String(), d.owner.String(), d.chainID.String(), d.deployTx.Hash())
}

var _ contracts.DeployedContract = DeployedContract{}
