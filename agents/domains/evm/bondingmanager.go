// Package evm is the package for evm contract stuff.
//
//nolint:dupl
package evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
)

// NewBondingManagerContract returns a bound bonding manager contract.
//
//nolint:staticcheck
func NewBondingManagerContract(ctx context.Context, client chain.Chain, bondingManagerAddress common.Address) (domains.BondingManagerContract, error) {
	boundCountract, err := bondingmanager.NewBondingManagerRef(bondingManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &bondingmanager.BondingManagerRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return bondingManagerContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type bondingManagerContract struct {
	// contract contains the conract handle
	contract *bondingmanager.BondingManagerRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

func (a bondingManagerContract) AddAgent(transactor *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.AddAgent(transactor, domain, agent, proof)
	if err != nil {
		return nil, fmt.Errorf("could not add agent: %w", err)
	}

	return tx, nil
}

//nolint:dupl
func (a bondingManagerContract) GetAgentStatus(ctx context.Context, address common.Address) (types.AgentStatus, error) {
	rawStatus, err := a.contract.AgentStatus(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve agent status: %w", err)
	}

	agentStatus := types.NewAgentStatus(types.AgentFlagType(rawStatus.Flag), rawStatus.Domain, rawStatus.Index)

	return agentStatus, nil
}

//nolint:dupl
func (a bondingManagerContract) GetAgentRoot(ctx context.Context) ([32]byte, error) {
	agentRoot, err := a.contract.AgentRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not retrieve agent root: %w", err)
	}

	return agentRoot, nil
}

//nolint:dupl
func (a bondingManagerContract) GetProof(ctx context.Context, address common.Address) ([][32]byte, error) {
	proof, err := a.contract.GetProof(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve agent proof: %w", err)
	}

	return proof, nil
}

func (a bondingManagerContract) DisputeStatus(ctx context.Context, address common.Address) (disputeStatus domains.DisputeStatus, err error) {
	rawDispute, err := a.contract.DisputeStatus(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return domains.DisputeStatus{}, fmt.Errorf("could not retrieve dispute status: %w", err)
	}

	return domains.DisputeStatus{
		DisputeFlag: rawDispute.Flag,
		Rival:       rawDispute.Rival,
		FraudProver: rawDispute.FraudProver,
		DisputePtr:  rawDispute.DisputePtr,
	}, nil
}

func (a bondingManagerContract) GetDispute(ctx context.Context, index *big.Int) (err error) {
	_, err = a.contract.GetDispute(&bind.CallOpts{Context: ctx}, index)
	if err != nil {
		return fmt.Errorf("could not retrieve dispute: %w", err)
	}

	return nil
}

func (a bondingManagerContract) CompleteSlashing(transactor *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (tx *ethTypes.Transaction, err error) {
	tx, err = a.contract.CompleteSlashing(transactor, domain, agent, proof)
	if err != nil {
		return nil, fmt.Errorf("could not submit state report: %w", err)
	}

	return tx, nil
}

func (a bondingManagerContract) GetDisputeStatus(ctx context.Context, agent common.Address) (disputeStatus types.DisputeStatus, err error) {
	rawStatus, err := a.contract.DisputeStatus(&bind.CallOpts{Context: ctx}, agent)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve dispute status: %w", err)
	}

	disputeStatus = types.NewDisputeStatus(types.DisputeFlagType(rawStatus.Flag), rawStatus.Rival, rawStatus.FraudProver, rawStatus.DisputePtr)
	return disputeStatus, nil
}

func (a bondingManagerContract) GetAgent(ctx context.Context, index *big.Int) (types.AgentStatus, common.Address, error) {
	rawStatus, err := a.contract.GetAgent(&bind.CallOpts{Context: ctx}, index)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("could not retrieve agent status: %w", err)
	}

	agentStatus := types.NewAgentStatus(types.AgentFlagType(rawStatus.Status.Flag), rawStatus.Status.Domain, rawStatus.Status.Index)

	return agentStatus, rawStatus.Agent, nil
}
