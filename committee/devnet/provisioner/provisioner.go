package provisioner

import (
	"context"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/core/metrics"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

type Provisioner struct {
	a              *anvil.Client
	b              *anvil.Client
	c              *anvil.Client
	client         omnirpcClient.RPCClient
	synapseModules map[int]*synapsemodule.SynapseModule
}

func NewProvisioner(ctx context.Context, handler metrics.Handler, cfg config.Config) (*Provisioner, error) {

	a, err := anvil.Dial(ctx, "http://localhost:9001/rpc/42")
	if err != nil {
		return nil, err
	}

	b, err := anvil.Dial(ctx, "http://localhost:9001/rpc/43")
	if err != nil {
		return nil, err
	}

	c, err := anvil.Dial(ctx, "http://localhost:9001/rpc/44")
	if err != nil {
		return nil, err
	}

	synapseModules := make(map[int]*synapsemodule.SynapseModule)

	client := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())

	return &Provisioner{
		a:              a,
		b:              b,
		c:              c,
		synapseModules: synapseModules,
		client:         client,
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
			common.HexToAddress(address), chainClient,
		)
		if err != nil {
			return fmt.Errorf("could not create synapse module: %w", err)
		}
		p.synapseModules[chainID] = synapseModuleDeployment
	}

	for chainID, synapseModule := range p.synapseModules {
		fmt.Println(chainID, synapseModule)
	}

	err := p.deleteVerifiers(ctx)
	if err != nil {
		return fmt.Errorf("could not delete verifiers: %v", err)
	}

	err = p.addVerifiers(ctx, cfg)
	if err != nil {
		return fmt.Errorf("could not add verifiers: %v", err)
	}

	err = p.changeThreshold(ctx)
	if err != nil {
		return fmt.Errorf("could not change threshold: %v", err)
	}

	return nil
}

func (p *Provisioner) deleteVerifiers(ctx context.Context) error {

	for chainid, synapseModule := range p.synapseModules {
		var err error

		owner, err := p.getSynapseModuleOwner(ctx, chainid)
		if err != nil {
			return fmt.Errorf("deleteVerifier: could not get synapse module owner: %w", err)
		}
		fmt.Println("owner: ", owner.String())
		verifiers, err := synapseModule.GetVerifiers(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("deleteVerifier: could not get verifiers: %w", err)
		}

		for _, verifier := range verifiers {
			fmt.Println("verifier: ", verifier.String())
		}

		_ = p.a.ImpersonateAccount(ctx, owner)

		tx, err := synapseModule.RemoveVerifiers(
			&bind.TransactOpts{
				From:     owner,
				Value:    big.NewInt(0),
				NoSend:   true,
				Signer:   anvil.ImpersonatedSigner,
				GasLimit: 0,
			},
			verifiers,
		)
		if err != nil {
			return fmt.Errorf("deleteVerifier: could not build tx: %w", err)
		}

		err = p.a.SendUnsignedTransaction(ctx, owner, tx)
		if err != nil {
			return fmt.Errorf("deleteVerifier: could not remove verifiers: %w", err)
		}

		_ = p.a.StopImpersonatingAccount(ctx, owner)

		break

	}

	fmt.Println("Successfully removed verifiers")
	return nil
}

func (p *Provisioner) addVerifiers(ctx context.Context, cfg config.Config) error {

	var i int
	for chainid, synapseModule := range p.synapseModules {
		owner, err := p.getSynapseModuleOwner(ctx, chainid)
		if err != nil {
			return fmt.Errorf("addVerifiers: could not get synapse module owner: %w", err)
		}

		if chainid == 42 {
			err = p.a.ImpersonateAccount(ctx, owner)
			if err != nil {
				return err
			}
			defer p.a.StopImpersonatingAccount(ctx, owner)
		} else if chainid == 43 {
			err = p.b.ImpersonateAccount(ctx, owner)
			if err != nil {
				return err
			}
			defer p.b.StopImpersonatingAccount(ctx, owner)
		} else if chainid == 44 {
			err = p.c.ImpersonateAccount(ctx, owner)
			if err != nil {
				return err
			}
			defer p.c.StopImpersonatingAccount(ctx, owner)
		}

		_, err = synapseModule.AddVerifiers(&bind.TransactOpts{
			From:   owner,
			Value:  big.NewInt(0),
			NoSend: false,
			Signer: anvil.ImpersonatedSigner,
		}, cfg.ValidatorAddresses)

		if err != nil {
			return fmt.Errorf("addVerifiers: could not add verifiers: %w", err)
		}

		// 		err = p.a.SendUnsignedTransaction(ctx, owner, tx)
		//
		// 		if err != nil {
		// 			return fmt.Errorf("addVerifiers: could not add verifiers: %w", err)
		// 		}
		//
		i++

	}

	return nil
}

func (p *Provisioner) changeThreshold(ctx context.Context) error {
	var i int
	for chainid, synapseModule := range p.synapseModules {
		owner, err := p.getSynapseModuleOwner(ctx, chainid)
		if err != nil {
			return fmt.Errorf("changeThreshold: could not get synapse module owner: %w", err)
		}

		if i == 0 {
			err = p.a.ImpersonateAccount(ctx, owner)
			if err != nil {
				return err
			}
			defer p.a.StopImpersonatingAccount(ctx, owner)
		} else if i == 1 {
			err = p.b.ImpersonateAccount(ctx, owner)
			if err != nil {
				return err
			}
			defer p.b.StopImpersonatingAccount(ctx, owner)
		} else if i == 2 {
			err = p.c.ImpersonateAccount(ctx, owner)
			if err != nil {
				return err
			}
			defer p.c.StopImpersonatingAccount(ctx, owner)
		}

		_, err = synapseModule.SetThreshold(&bind.TransactOpts{
			From:   owner,
			Value:  big.NewInt(0),
			NoSend: true,
			Signer: anvil.ImpersonatedSigner,
			Nonce:  nil,
		}, big.NewInt(2))

		if err != nil {
			return fmt.Errorf("changeThreshold: could not change threshold: %w", err)
		}

		i++

	}

	return nil
}

func (p *Provisioner) getSynapseModuleOwner(ctx context.Context, chainid int) (common.Address, error) {
	owner, err := p.synapseModules[chainid].Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}
