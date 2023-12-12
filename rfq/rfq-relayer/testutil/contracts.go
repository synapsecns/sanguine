package testutil

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
)

/*
contracts.go contains utils for interacting with the contracts on a given chain.
This makes testing the relayer easier, as we can just use the simple methods in the ITestContractHandler interface.

This file is purely for testing purposes. It is not used in the production code.
*/

const (
	DefaultMintAmount = params.Ether
	NumberOfTokens    = 2
)

// ITestContractHandler is the interface for the handling contracts on a given chain.
type ITestContractHandler interface {
	ChainID() uint32
	FastBridgeAddress() common.Address
	Tokens() []TokenContract
	FBExecuteBridge(ctx context.Context, bridgeParams bindings.IFastBridgeBridgeParams) (*types.Transaction, error)
	FBExecuteRelay(ctx context.Context, request []byte) (*types.Transaction, error)
	FBExecuteProve(ctx context.Context, request []byte, destTxHash [32]byte) (*types.Transaction, error)
	FBExecuteClaim(ctx context.Context, request []byte, to common.Address) (*types.Transaction, error)
}

// TokenContract is a struct containing the address and contract of a token.
type TokenContract struct {
	Erc20Address  common.Address
	Erc20Contract *bindings.MockERC20
}

// TestContractHandlerImpl is the implementation of the ITestContractHandler interface.
type TestContractHandlerImpl struct {
	fastBridgeAddress  common.Address
	fastBridgeContract *bindings.FastBridge
	tokens             []TokenContract
	anvilBackend       backends.SimulatedTestBackend
	testWallet         wallet.Wallet
	chainID            uint32
}

// NewTestContractHandlerImpl creates a new instance of the ITestContractHandler interface.
func NewTestContractHandlerImpl(ctx context.Context, anvilBackend backends.SimulatedTestBackend, testWallet wallet.Wallet, chainID uint32) (ITestContractHandler, error) {
	fundAmount := big.NewInt(params.Ether)
	anvilBackend.FundAccount(ctx, testWallet.Address(), *fundAmount)

	// Create an auth to interact with the blockchain
	auth, err := bind.NewKeyedTransactorWithChainID(testWallet.PrivateKey(), big.NewInt(int64(chainID)))
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %w", err)
	}

	// Deploy fast bridge contract
	fastBridgeAddress, tx, fastBridgeContract, err := bindings.DeployFastBridge(auth, anvilBackend, testWallet.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to deploy fast bridge contract: %w", err)
	}
	anvilBackend.WaitForConfirmation(ctx, tx)

	// Make wallet a relayer
	tx, err = fastBridgeContract.AddRelayer(auth, testWallet.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to add relayer: %w", err)
	}
	anvilBackend.WaitForConfirmation(ctx, tx)

	// Creates a new instance of the ERC20 contract, mints tokens to the test wallet, and approves the fast bridge contract.
	var tokens []TokenContract
	for i := 0; i < NumberOfTokens; i++ {
		tokenName := fmt.Sprintf("TESTTOKEN_%d", i)

		// Deploy Mock ERC 20 contract
		erc20address, erc20Tx, erc20contract, erc20Err := bindings.DeployMockERC20(auth, anvilBackend, tokenName, 18)
		if erc20Err != nil {
			return nil, fmt.Errorf("failed to deploy mock erc20 contract: %w", erc20Err)
		}
		anvilBackend.WaitForConfirmation(ctx, erc20Tx)

		// Mint to test wallet
		mintTx, erc20Err := erc20contract.Mint(auth, testWallet.Address(), big.NewInt(DefaultMintAmount))
		if erc20Err != nil {
			return nil, fmt.Errorf("failed to mint tokens: %w", erc20Err)
		}
		anvilBackend.WaitForConfirmation(ctx, mintTx)

		// Approve token + fast bridge contract
		approveTx, erc20Err := erc20contract.Approve(auth, fastBridgeAddress, big.NewInt(DefaultMintAmount))
		if erc20Err != nil {
			return nil, fmt.Errorf("failed to approve tokens: %w", erc20Err)
		}
		anvilBackend.WaitForConfirmation(ctx, approveTx)

		// Add to token list
		tokens = append(tokens, TokenContract{
			Erc20Address:  erc20address,
			Erc20Contract: erc20contract,
		})
	}

	return &TestContractHandlerImpl{
		fastBridgeAddress:  fastBridgeAddress,
		fastBridgeContract: fastBridgeContract,
		tokens:             tokens,
		anvilBackend:       anvilBackend,
		testWallet:         testWallet,
		chainID:            chainID,
	}, nil
}

// ChainID returns the chain ID of the contract handler.
func (t *TestContractHandlerImpl) ChainID() uint32 {
	return t.chainID
}

// FastBridgeAddress returns the address of the fast bridge contract.
func (t *TestContractHandlerImpl) FastBridgeAddress() common.Address {
	return t.fastBridgeAddress
}

// Tokens returns the list of tokens.
func (t *TestContractHandlerImpl) Tokens() []TokenContract {
	return t.tokens
}

// FBExecuteBridge executes the bridge function on the fast bridge contract.
func (t *TestContractHandlerImpl) FBExecuteBridge(ctx context.Context, bridgeParams bindings.IFastBridgeBridgeParams) (*types.Transaction, error) {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(t.testWallet.PrivateKey(), big.NewInt(int64(t.chainID)))
	tx, err := t.fastBridgeContract.Bridge(transactOpts, bridgeParams)
	if err != nil {
		return nil, err
	}
	t.anvilBackend.WaitForConfirmation(ctx, tx)
	return tx, nil
}

// FBExecuteRelay executes the relay function on the fast bridge contract.
func (t *TestContractHandlerImpl) FBExecuteRelay(ctx context.Context, request []byte) (*types.Transaction, error) {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(t.testWallet.PrivateKey(), big.NewInt(int64(t.chainID)))
	tx, err := t.fastBridgeContract.Relay(transactOpts, request)
	if err != nil {
		return nil, err
	}
	t.anvilBackend.WaitForConfirmation(ctx, tx)
	return tx, nil
}

// FBExecuteProve executes the prove function on the fast bridge contract.
func (t *TestContractHandlerImpl) FBExecuteProve(ctx context.Context, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(t.testWallet.PrivateKey(), big.NewInt(int64(t.chainID)))
	tx, err := t.fastBridgeContract.Prove(transactOpts, request, destTxHash)
	if err != nil {
		return nil, err
	}
	t.anvilBackend.WaitForConfirmation(ctx, tx)
	return tx, nil
}

// FBExecuteClaim executes the claim function on the fast bridge contract. Will likely be unused.
func (t *TestContractHandlerImpl) FBExecuteClaim(ctx context.Context, request []byte, to common.Address) (*types.Transaction, error) {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(t.testWallet.PrivateKey(), big.NewInt(int64(t.chainID)))
	tx, err := t.fastBridgeContract.Claim(transactOpts, request, to)
	if err != nil {
		return nil, err
	}
	t.anvilBackend.WaitForConfirmation(ctx, tx)
	return tx, nil
}
