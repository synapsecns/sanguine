// Package evm is the package for evm contract stuff.
//
//nolint:dupl
package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/contracts/lightmanager"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewLightManagerContract returns a bound light manager contract.
//
//nolint:staticcheck
func NewLightManagerContract(ctx context.Context, client chain.Chain, lightManagerAddress common.Address) (domains.LightManagerContract, error) {
	boundCountract, err := lightmanager.NewLightManagerRef(lightManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", &lightmanager.LightManagerRef{}, err)
	}

	nonceManager := nonce.NewNonceManager(ctx, client, client.GetBigChainID())
	return lightManagerContract{
		contract:     boundCountract,
		client:       client,
		nonceManager: nonceManager,
	}, nil
}

type lightManagerContract struct {
	// contract contains the contract handle
	contract *lightmanager.LightManagerRef
	// client contains the evm client
	//nolint: staticcheck
	client chain.Chain
	// nonceManager is the nonce manager used for transacting with the chain
	nonceManager nonce.Manager
}

//nolint:dupl
func (a lightManagerContract) GetAgentStatus(ctx context.Context, bondedAgentSigner signer.Signer) (types.AgentStatus, error) {
	rawStatus, err := a.contract.AgentStatus(&bind.CallOpts{Context: ctx}, bondedAgentSigner.Address())
	if err != nil {
		return nil, fmt.Errorf("could not retrieve agent status: %w", err)
	}

	agentStatus := types.NewAgentStatus(rawStatus.Flag, rawStatus.Domain, rawStatus.Index)

	return agentStatus, nil
}
