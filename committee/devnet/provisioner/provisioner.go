package provisioner

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/core/metrics"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

type Provisioner struct {
	*anvil.Backend
	client         omnirpcClient.RPCClient
	synapseModules map[int]*synapsemodule.SynapseModule
}

func NewProvisioner(ctx context.Context, handler metrics.Handler, cfg config.Config) (*Provisioner, error) {

	synapseModules := make(map[int]*synapsemodule.SynapseModule)

	// TODO: make it configurable

	c := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())

	return &Provisioner{
		synapseModules: synapseModules,
		client:         c,
	}, nil
}

// Run removes verifiers from the synapse module by impersonating the owner, and adds our own.
func (p *Provisioner) Run(ctx context.Context, cfg config.Config) error {

	for chainID, address := range cfg.Chains {
		chainClient, err := p.client.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		synapseModuleDeployment, err := synapsemodule.NewSynapseModule(
			common.HexToAddress(address), chainClient)
		if err != nil {
			return fmt.Errorf("could not create synapse module: %w", err)
		}
		p.synapseModules[chainID] = synapseModuleDeployment
	}

	for chainID, synapseModule := range p.synapseModules {
		fmt.Println(chainID, synapseModule)
	}

	fmt.Println("about to call owner penis")
	owner, _ := p.synapseModules[42].Owner(nil)
	fmt.Println("owner penis", owner)

	err := p.deleteVerifiers(ctx)
	if err != nil {
		return fmt.Errorf("could not delete verifiers: %v", err)
	}

	err = p.addVerifiers(ctx, cfg)
	if err != nil {
		return fmt.Errorf("could not add verifiers: %v", err)
	}

	return nil
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
	fmt.Println("fauk")
	fmt.Println(p.synapseModules)
	owner, err := p.synapseModules[chainid].Owner(&bind.CallOpts{})
	fmt.Println(err)
	fmt.Println("ok we passed")
	if err != nil {
		fmt.Println("err was not nil")
		return common.Address{}, err
	}

	return owner, nil
}
