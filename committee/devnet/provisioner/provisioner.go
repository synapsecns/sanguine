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
	"github.com/ethereum/go-ethereum/params"
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
	gasOracles     map[int]*gasoracle.SynapseGasOracleV1
}

func NewProvisioner(ctx context.Context, handler metrics.Handler, cfg config.ProvisionerConfig) (*Provisioner, error) {

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
	gasOracles := make(map[int]*gasoracle.SynapseGasOracleV1)

	client := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())

	return &Provisioner{
		a:              a,
		b:              b,
		c:              c,
		synapseModules: synapseModules,
		gasOracles:     gasOracles,
		client:         client,
	}, nil
}

// Run removes verifiers from the synapse module by impersonating the owner, and adds our own.
func (p *Provisioner) Run(ctx context.Context, cfg config.ProvisionerConfig) error {

	for chainID, deployedSynapseModule := range cfg.SynapseModuleDeployments {
		chainClient, err := p.client.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		synapseModuleDeployment, err := synapsemodule.NewSynapseModule(
			common.HexToAddress(deployedSynapseModule), chainClient,
		)
		if err != nil {
			return fmt.Errorf("could not create synapse module: %w", err)
		}
		p.synapseModules[chainID] = synapseModuleDeployment
	}

	for chainID, gasOracleAddress := range cfg.GasOracleDeployments {
		chainClient, err := p.client.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}
		gasOracle, err := gasoracle.NewSynapseGasOracleV1(
			common.HexToAddress(gasOracleAddress), chainClient,
		)
		if err != nil {
			return fmt.Errorf("could not create gas oracle: %w", err)
		}
		p.gasOracles[chainID] = gasOracle
	}

	// TODO: make this cleaner
	// do it programmatically, for now idc bbecause jus twantti tto work
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

	// ================== Gas Oracle ==================
	err = p.setLocalNativePrice(ctx, p.gasOracles[42], 42, big.NewInt(params.Wei), p.a)
	if err != nil {
		return fmt.Errorf("could not add local native price: %v", err)
	}
	err = p.setLocalNativePrice(ctx, p.gasOracles[43], 43, big.NewInt(params.Wei), p.b)
	if err != nil {
		return fmt.Errorf("could not add local native price: %v", err)
	}

	err = p.setRemoteCallDataPrice(ctx, p.gasOracles[42], 42, 43, big.NewInt(0), p.a)
	if err != nil {
		return fmt.Errorf("could not set remote calldata price: %v", err)
	}
	err = p.setRemoteCallDataPrice(ctx, p.gasOracles[43], 43, 42, big.NewInt(0), p.b)
	if err != nil {
		return fmt.Errorf("could not set remote calldata price %v", err)
	}

	err = p.setRemoteGasPrice(ctx, p.gasOracles[42], 42, 43, big.NewInt(1000), p.a)
	if err != nil {
		return fmt.Errorf("could not add remote chain gas price: %v", err)
	}
	err = p.setRemoteGasPrice(ctx, p.gasOracles[43], 43, 42, big.NewInt(1000), p.b)
	if err != nil {
		return fmt.Errorf("could not add remote chain gas price: %v", err)
	}

	err = p.setRemoteNativePrice(ctx, p.gasOracles[42], 42, 43, big.NewInt(params.Wei), p.a)
	if err != nil {
		return fmt.Errorf("could not add remote native price: %v", err)
	}
	err = p.setRemoteNativePrice(ctx, p.gasOracles[43], 43, 42, big.NewInt(params.Wei), p.b)
	if err != nil {
		return fmt.Errorf("could not add remote native price: %v", err)
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
	owner, err := p.getSynapseModuleOwner(ctx, chainid)
	if err != nil {
		return fmt.Errorf("changeThreshold: could not get synapse module owner: %w", err)
	}

	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	received, err := synapseModule.GetThreshold(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("changeThreshold: could not get threshold: %w", err)
	}
	if received.Cmp(threshold) == 0 {
		fmt.Printf("threshold is already %d on chain %d\n", threshold, chainid)
		return nil
	}

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

func (p *Provisioner) setRemoteCallDataPrice(
	ctx context.Context,
	originGasOracle *gasoracle.SynapseGasOracleV1,
	originChainId int,
	remoteChainId int,
	gasPrice *big.Int,
	client *anvil.Client,
) error {

	owner, err := p.getGasOracleOwner(ctx, originChainId)
	if err != nil {
		return err
	}
	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	remoteGasData, err := originGasOracle.GetRemoteGasData(
		&bind.CallOpts{Context: ctx},
		uint64(remoteChainId),
	)
	if err != nil {
		return fmt.Errorf("setRemoteGasPrice: could not get remote gas price: %w", err)
	}
	if remoteGasData.CalldataPrice.Cmp(gasPrice) == 0 {
		fmt.Printf("remote gas price is already %d on chain %d\n", gasPrice, remoteChainId)
		return nil
	}

	tx, err := originGasOracle.SetRemoteCallDataPrice(
		&bind.TransactOpts{
			From:   owner,
			Value:  big.NewInt(0),
			NoSend: true,
			Signer: anvil.ImpersonatedSigner,
			Nonce:  nil,
		}, uint64(remoteChainId), gasPrice)
	if err != nil {
		return fmt.Errorf("setRemoteCallDataPrice: could not build tx: %w", err)
	}
	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("setRemoteCallDataPrice: could not send tx: %w", err)
	}

	fmt.Printf("Successfully set remote call data price to %d on chain %d\n", gasPrice, remoteChainId)
	return nil
}

func (p *Provisioner) setRemoteGasPrice(
	ctx context.Context,
	originGasOracle *gasoracle.SynapseGasOracleV1,
	originChainId int,
	remoteChainId int,
	gasPrice *big.Int,
	client *anvil.Client,
) error {

	owner, err := p.getGasOracleOwner(ctx, originChainId)
	if err != nil {
		return err
	}
	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	remoteGasData, err := originGasOracle.GetRemoteGasData(
		&bind.CallOpts{Context: ctx},
		uint64(remoteChainId),
	)
	if err != nil {
		return fmt.Errorf("setRemoteGasPrice: could not get remote gas price: %w", err)
	}
	if remoteGasData.GasPrice.Cmp(gasPrice) == 0 {
		fmt.Printf("remote gas price is already %d on chain %d\n", gasPrice, remoteChainId)
		return nil
	}

	tx, err := originGasOracle.SetRemoteGasPrice(
		&bind.TransactOpts{
			From:   owner,
			Value:  big.NewInt(0),
			NoSend: true,
			Signer: anvil.ImpersonatedSigner,
			Nonce:  nil,
		}, uint64(remoteChainId), gasPrice)
	if err != nil {
		return fmt.Errorf("setRemoteGasPrice: could not build tx: %w", err)
	}
	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("setRemoteGasPrice: could not send tx: %w", err)
	}

	fmt.Printf("Successfully set remote gas price to %d on chain %d\n", gasPrice, remoteChainId)
	return nil
}

func (p *Provisioner) setRemoteNativePrice(
	ctx context.Context,
	originGasOracle *gasoracle.SynapseGasOracleV1,
	originChainId int,
	remoteChainId int,
	gasPrice *big.Int,
	client *anvil.Client,
) error {

	owner, err := p.getGasOracleOwner(ctx, originChainId)
	if err != nil {
		return err
	}
	err = client.ImpersonateAccount(ctx, owner)
	if err != nil {
		return err
	}
	defer client.StopImpersonatingAccount(ctx, owner)

	remoteGasData, err := originGasOracle.GetRemoteGasData(
		&bind.CallOpts{Context: ctx},
		uint64(remoteChainId),
	)
	if err != nil {
		return fmt.Errorf("setRemoteGasPrice: could not get remote gas price: %w", err)
	}
	if remoteGasData.NativePrice.Cmp(gasPrice) == 0 {
		fmt.Printf("remote gas price is already %d on chain %d\n", gasPrice, remoteChainId)
		return nil
	}

	tx, err := originGasOracle.SetRemoteNativePrice(
		&bind.TransactOpts{
			From:   owner,
			Value:  big.NewInt(0),
			NoSend: true,
			Signer: anvil.ImpersonatedSigner,
			Nonce:  nil,
		}, uint64(remoteChainId), gasPrice)
	if err != nil {
		return fmt.Errorf("setRemoteNativePrice: could not build tx: %w", err)
	}
	err = client.SendUnsignedTransaction(ctx, owner, tx)
	if err != nil {
		return fmt.Errorf("setRemoteNativePrice: could not send tx: %w", err)
	}

	fmt.Printf("Successfully set remote native price to %d on chain %d\n", gasPrice, remoteChainId)

	return nil
}

func (p *Provisioner) getGasOracleOwner(ctx context.Context, chainid int) (common.Address, error) {
	owner, err := p.gasOracles[chainid].Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}

func (p *Provisioner) getSynapseModuleOwner(ctx context.Context, chainid int) (common.Address, error) {
	owner, err := p.synapseModules[chainid].Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}
