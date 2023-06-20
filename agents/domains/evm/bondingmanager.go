// Package evm is the package for evm contract stuff.
//
//nolint:dupl
package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/bondingmanager"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
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

//nolint:dupl
func (a bondingManagerContract) GetAgentStatus(ctx context.Context, bondedAgentSigner signer.Signer) (types.AgentStatus, error) {
	rawStatus, err := a.contract.AgentStatus(&bind.CallOpts{Context: ctx}, bondedAgentSigner.Address())
	if err != nil {
		return nil, fmt.Errorf("could not retrieve agent status: %w", err)
	}

	agentStatus := types.NewAgentStatus(rawStatus.Flag, rawStatus.Domain, rawStatus.Index)

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
func (a bondingManagerContract) GetProof(ctx context.Context, bondedAgentSigner signer.Signer) ([][32]byte, error) {
	proof, err := a.contract.GetProof(&bind.CallOpts{Context: ctx}, bondedAgentSigner.Address())
	if err != nil {
		return nil, fmt.Errorf("could not retrieve agent proof: %w", err)
	}

	return proof, nil
}
