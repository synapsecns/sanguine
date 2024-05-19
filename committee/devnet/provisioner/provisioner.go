package provisioner

import (
	"context"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/committee/devnet/config"
	"github.com/synapsecns/sanguine/committee/devnet/contracts/gasoracle"
	"github.com/synapsecns/sanguine/core/metrics"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

type Provisioner struct {
	client     omnirpcClient.RPCClient
	gasOracles map[int]*gasoracle.SynapseGasOracleV1
}

func NewProvisioner(ctx context.Context, handler metrics.Handler, cfg config.ProvisionerConfig) (*Provisioner, error) {

	gasOracles := make(map[int]*gasoracle.SynapseGasOracleV1)

	client := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())

	return &Provisioner{
		gasOracles: gasOracles,
		client:     client,
	}, nil
}

// Run removes verifiers from the synapse module by impersonating the owner, and adds our own.
func (p *Provisioner) Run(ctx context.Context, cfg config.ProvisionerConfig) error {

	// TODO: make this cleaner
	for chainID, synapseModuleAddress := range cfg.SynapseModuleDeployments {
		chainClient, err := p.client.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		anvilClient, err := anvil.Dial(ctx, fmt.Sprintf("http://omnirpc:9001/rpc/%d", chainID))
		if err != nil {
			return fmt.Errorf("could not get anvil client: %w", err)
		}

		synapseModule, err := synapsemodule.NewSynapseModule(common.HexToAddress(synapseModuleAddress), chainClient)
		if err != nil {
			return fmt.Errorf("could not create synapse module: %w", err)
		}

		if err = p.deleteVerifiers(ctx, synapseModule, chainID, anvilClient); err != nil {
			return fmt.Errorf("could not delete verifiers on chain %d: %v", chainID, err)
		}

		if err = p.addVerifiers(ctx, synapseModule, chainID, anvilClient, cfg); err != nil {
			return fmt.Errorf("could not add verifiers on chain %d: %v", chainID, err)
		}

		if err = p.changeThreshold(ctx, synapseModule, chainID, big.NewInt(1), anvilClient); err != nil {
			return fmt.Errorf("could not change threshold on chain %d: %v", chainID, err)
		}

		fmt.Println("---------------------------------------------------------")
	}

	// map each chain to a gas oracle
	for chainID, gasOracleAddress := range cfg.GasOracleDeployments {
		chainClient, err := p.client.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		gasOracle, err := gasoracle.NewSynapseGasOracleV1(common.HexToAddress(gasOracleAddress), chainClient)
		if err != nil {
			return fmt.Errorf("could not create gas oracle: %w", err)
		}
		p.gasOracles[chainID] = gasOracle
	}

	// ================== Gas Oracle ==================
	chainids := make([]int, 0, len(p.gasOracles))
	for chainid := range p.gasOracles {
		chainids = append(chainids, chainid)
	}

	// chain a <> chain b pairwise gas oracle setup
	for i := 0; i < len(chainids); i++ {
		for j := 0; j < len(chainids); j++ {
			if i != j {
				anvilClient, err := anvil.Dial(ctx, fmt.Sprintf("http://omnirpc:9001/rpc/%d", chainids[i]))
				if err != nil {
					return fmt.Errorf("could not get anvil client: %w", err)
				}

				err = p.setLocalNativePrice(ctx, p.gasOracles[chainids[i]], chainids[j], big.NewInt(100), anvilClient)
				if err != nil {
					return fmt.Errorf("could not add local native price: %v", err)
				}

				err = p.setRemoteGasData(
					ctx,
					p.gasOracles[chainids[i]],
					chainids[j],
					gasoracle.ISynapseGasOracleV1RemoteGasData{
						GasPrice:      big.NewInt(10),
						CalldataPrice: big.NewInt(1),
						NativePrice:   big.NewInt(10),
					},
					anvilClient,
				)
				if err != nil {
					return fmt.Errorf("could not set remote gas data: %v", err)
				}
			}
		}
	}

	fmt.Println("Successfully provisioned SynapseModule contracts.")

	return nil
}

func (p *Provisioner) deleteVerifiers(
	ctx context.Context, synapseModule *synapsemodule.SynapseModule, chainid int, client *anvil.Client,
) error {
	owner, err := p.getSynapseModuleOwner(ctx, synapseModule)
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

	fmt.Printf("Successfully removed verifiers %v on chain %d\n", verifiers, chainid)
	return nil
}

func (p *Provisioner) addVerifiers(
	ctx context.Context,
	synapseModule *synapsemodule.SynapseModule,
	chainid int,
	client *anvil.Client,
	cfg config.ProvisionerConfig,
) error {
	owner, err := p.getSynapseModuleOwner(ctx, synapseModule)
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

	fmt.Printf("Successfully added verifiers %v on chain %d\n", cfg.ValidatorAddresses, chainid)
	return nil
}

func (p *Provisioner) changeThreshold(
	ctx context.Context,
	synapseModule *synapsemodule.SynapseModule,
	chainid int,
	threshold *big.Int,
	client *anvil.Client,
) error {
	owner, err := p.getSynapseModuleOwner(ctx, synapseModule)
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

	fmt.Printf("successfully changed threshold to %d on chain %d\n", threshold, chainid)
	return nil
}

func (p *Provisioner) setLocalNativePrice(
	ctx context.Context,
	originGasOracle *gasoracle.SynapseGasOracleV1,
	originChainId int,
	gasPrice *big.Int,
	client *anvil.Client,
) error {
	// call the gasprice oracle and update for the respective chainids
	owner, err := p.getGasOracleOwner(ctx, originChainId)
	if err != nil {
		return fmt.Errorf("addVerifiers: could not get synapse module owner: %w", err)
	}

	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	nativePrice, err := originGasOracle.GetLocalNativePrice(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("setLocalNativePrice: could not get local native price: %w", err)
	}
	if nativePrice.Cmp(gasPrice) == 0 {
		fmt.Printf("local native price is already %d on chain %d\n", gasPrice, originChainId)
		return nil
	}

	tx, err := originGasOracle.SetLocalNativePrice(
		&bind.TransactOpts{
			From:   owner,
			Value:  big.NewInt(0),
			NoSend: true,
			Signer: anvil.ImpersonatedSigner,
			Nonce:  nil,
		}, gasPrice)
	if err != nil {
		return fmt.Errorf("setLocalNativePrice: could not build tx: %w", err)
	}

	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("setLocalNativePrice: could not send tx: %w", err)
	}

	fmt.Printf("Successfully set local native price to %d on chain %d\n", gasPrice, originChainId)

	return nil
}

func (p *Provisioner) setRemoteGasData(
	ctx context.Context,
	originGasOracle *gasoracle.SynapseGasOracleV1,
	remoteChainId int,
	data gasoracle.ISynapseGasOracleV1RemoteGasData,
	client *anvil.Client,
) error {
	owner, err := p.getGasOracleOwner(ctx, 42)
	if err != nil {
		return fmt.Errorf("addVerifiers: could not get synapse module owner: %w", err)
	}

	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	tx, err := originGasOracle.SetRemoteGasData(
		&bind.TransactOpts{
			From:   owner,
			Value:  big.NewInt(0),
			NoSend: true,
			Signer: anvil.ImpersonatedSigner,
			Nonce:  nil,
		}, uint64(remoteChainId), data)

	if err != nil {
		return fmt.Errorf("setRemoteGasData: could not build tx: %w", err)
	}

	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("setRemoteGasData: could not send tx: %w", err)
	}

	fmt.Printf("Successfully set remote gas data to %v on chain %d\n", data, remoteChainId)
	return nil
}

func (p *Provisioner) getGasOracleOwner(ctx context.Context, chainid int) (common.Address, error) {
	owner, err := p.gasOracles[chainid].Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}

func (p *Provisioner) getSynapseModuleOwner(ctx context.Context, synapseModule *synapsemodule.SynapseModule) (common.Address, error) {
	owner, err := synapseModule.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}
