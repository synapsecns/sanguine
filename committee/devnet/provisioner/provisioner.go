package provisioner

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/committee/devnet/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
)

type Provisioner struct {
	*anvil.Backend
	synapseModules map[int]*synapsemodule.SynapseModule
}

func NewProvisioner(ctx context.Context, cfg config.Config) (*Provisioner, error) {

	synapseModuleDeployments := make(map[int]*synapsemodule.SynapseModule)

	for chainID, address := range cfg.Chains {
		synapseModuleDeployment, err := synapsemodule.NewSynapseModule(common.HexToAddress(address), nil)
		if err != nil {
			return nil, fmt.Errorf("could not create synapse module: %w", err)
		}
		synapseModuleDeployments[chainID] = synapseModuleDeployment
	}

	return &Provisioner{
		synapseModules: synapseModuleDeployments,
	}, nil
}

// Run removes verifiers from the synapse module by impersonating the owner, and adds our own.
func (p *Provisioner) Run(ctx context.Context, cfg config.Config) {

	err := p.deleteVerifiers(ctx)
	if err != nil {
		fmt.Printf("could not delete verifiers: %v\n", err)
		return
	}

	err = p.addVerifiers(ctx, cfg)
	if err != nil {
		fmt.Printf("could not add verifiers: %v\n", err)
		return
	}

}

func (p *Provisioner) deleteVerifiers(ctx context.Context) error {
	for chainid, synapseModule := range p.synapseModules {
		owner, err := p.getSynapseModuleOwner(chainid)
		if err != nil {
			return fmt.Errorf("could not get synapse module owner: %w", err)
		}
		verifiers, err := synapseModule.GetVerifiers(nil)
		if err != nil {
			return fmt.Errorf("could not get verifiers: %w", err)
		}

		err = p.ImpersonateAccount(ctx, owner, func(transactOpts *bind.TransactOpts) *types.Transaction {
			tx, err := synapseModule.RemoveVerifiers(transactOpts, verifiers)
			if err != nil {
				return nil
			}

			return tx
		})

		if err != nil {
			return fmt.Errorf("could not impersonate account: %w", err)
		}
	}

	return nil
}

func (p *Provisioner) addVerifiers(ctx context.Context, cfg config.Config) error {

	for chainid, synapseModule := range p.synapseModules {
		owner, err := p.getSynapseModuleOwner(chainid)
		if err != nil {
			return fmt.Errorf("could not get synapse module owner: %w", err)
		}
		err = p.ImpersonateAccount(ctx, owner, func(transactOpts *bind.TransactOpts) *types.Transaction {
			tx, err := synapseModule.AddVerifiers(transactOpts, cfg.ValidatorAddresses)
			if err != nil {
				return nil
			}
			return tx
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Provisioner) getSynapseModuleOwner(chainid int) (common.Address, error) {
	owner, err := p.synapseModules[chainid].Owner(nil)
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}
