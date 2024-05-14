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

	err := p.deleteVerifiers(ctx, p.synapseModules[42], 42, p.a)
	if err != nil {
		return fmt.Errorf("could not delete verifiers on chain 42: %v", err)
	}
	err = p.deleteVerifiers(ctx, p.synapseModules[43], 43, p.b)
	if err != nil {
		return fmt.Errorf("could not delete verifiers on chain 43: %v", err)
	}
	err = p.deleteVerifiers(ctx, p.synapseModules[44], 44, p.c)
	if err != nil {
		return fmt.Errorf("could not delete verifiers on chain 44: %v", err)
	}

	err = p.addVerifiers(ctx, p.synapseModules[42], 42, p.a, cfg)
	if err != nil {
		return fmt.Errorf("could not add verifiers on chain 42: %v", err)
	}
	err = p.addVerifiers(ctx, p.synapseModules[43], 43, p.b, cfg)
	if err != nil {
		return fmt.Errorf("could not add verifiers on chain 43: %v", err)
	}
	err = p.addVerifiers(ctx, p.synapseModules[44], 44, p.c, cfg)
	if err != nil {
		return fmt.Errorf("could not add verifiers on chain 44: %v", err)
	}

	err = p.changeThreshold(ctx, p.synapseModules[42], 42, big.NewInt(2), p.a)
	if err != nil {
		return fmt.Errorf("could not change threshold: %v", err)
	}
	err = p.changeThreshold(ctx, p.synapseModules[43], 43, big.NewInt(2), p.b)
	if err != nil {
		return fmt.Errorf("could not change threshold: %v", err)
	}
	err = p.changeThreshold(ctx, p.synapseModules[44], 44, big.NewInt(2), p.c)
	if err != nil {
		return fmt.Errorf("could not change threshold: %v", err)
	}

	fmt.Println("Successfully provisioned SynapseModule contracts.")

	return nil
}

func (p *Provisioner) deleteVerifiers(
	ctx context.Context, synapseModule *synapsemodule.SynapseModule, chainid int, client *anvil.Client,
) error {

	owner, err := p.getSynapseModuleOwner(ctx, chainid)
	if err != nil {
		return fmt.Errorf("deleteVerifier: could not get synapse module owner: %w", err)
	}
	verifiers, err := synapseModule.GetVerifiers(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("deleteVerifier: could not get verifiers: %w", err)
	}

	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return fmt.Errorf("deleteVerifier: could not impersonate account: %w", err)
	}
	defer client.StopImpersonatingAccount(ctx, owner)

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

	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("deleteVerifier: could not remove verifiers: %w", err)
	}

	fmt.Println("successfully removed verifiers")
	return nil
}

func (p *Provisioner) addVerifiers(ctx context.Context, synapseModule *synapsemodule.SynapseModule, chainid int, client *anvil.Client, cfg config.Config) error {

	owner, err := p.getSynapseModuleOwner(ctx, chainid)
	if err != nil {
		return fmt.Errorf("addVerifiers: could not get synapse module owner: %w", err)
	}

	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	tx, err := synapseModule.AddVerifiers(&bind.TransactOpts{
		From:   owner,
		Value:  big.NewInt(0),
		NoSend: true,
		Signer: anvil.ImpersonatedSigner,
	}, cfg.ValidatorAddresses)
	if err != nil {
		return fmt.Errorf("addVerifiers: could not add verifiers: %w", err)
	}

	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("addVerifiers: could not send tx: %w", err)
	}

	fmt.Println("successfully added verifiers")
	return nil
}

func (p *Provisioner) changeThreshold(ctx context.Context, synapseModule *synapsemodule.SynapseModule, chainid int, threshold *big.Int, client *anvil.Client) error {
	owner, err := p.getSynapseModuleOwner(ctx, chainid)
	if err != nil {
		return fmt.Errorf("changeThreshold: could not get synapse module owner: %w", err)
	}

	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	tx, err := synapseModule.SetThreshold(&bind.TransactOpts{
		From:   owner,
		Value:  big.NewInt(0),
		NoSend: true,
		Signer: anvil.ImpersonatedSigner,
		Nonce:  nil,
	}, threshold)
	if err != nil {
		return fmt.Errorf("changeThreshold: could not change threshold: %w", err)
	}

	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("changeThreshold: could not send tx: %w", err)
	}

	fmt.Println("successfully changed threshold")
	return nil
}

func (p *Provisioner) getSynapseModuleOwner(ctx context.Context, chainid int) (common.Address, error) {
	owner, err := p.synapseModules[chainid].Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}
