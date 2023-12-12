package balance_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	EVMClient "github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/balance"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/testutil"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"
)

func (t *BalanceSuite) TestNewBalanceManager() {
	testCtx := t.GetTestContext()
	clients := make(map[uint32]EVMClient.EVM)
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, testChainID, t.T())
	evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[testChainID] = evmClient

	// Wallet
	testWallet, _ := wallet.FromRandom()
	testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
	Nil(t.T(), err)
	NotNil(t.T(), testContractHandler)

	// Get Tokens
	tokens := testContractHandler.Tokens()
	var assets []config.AssetConfig
	for _, token := range tokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: testChainID,
		}
		assets = append(assets, asset)
	}

	balanceManager, err := balance.NewBalanceManager(clients, assets, testWallet.Address())
	Nil(t.T(), err)
	NotNil(t.T(), balanceManager)
}

func (t *BalanceSuite) TestGetBalance() {
	testCtx := t.GetTestContext()
	clients := make(map[uint32]EVMClient.EVM)
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, testChainID, t.T())
	evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[testChainID] = evmClient

	// Wallet
	testWallet, _ := wallet.FromRandom()
	testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
	Nil(t.T(), err)
	NotNil(t.T(), testContractHandler)

	// Get Tokens
	tokens := testContractHandler.Tokens()
	var assets []config.AssetConfig
	for _, token := range tokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: testChainID,
		}
		assets = append(assets, asset)
	}

	balanceManager, err := balance.NewBalanceManager(clients, assets, testWallet.Address())
	Nil(t.T(), err)
	NotNil(t.T(), balanceManager)

	// Get Balance
	err = balanceManager.GetOnChainBalance(testCtx, utils.GenerateTokenID(testChainID, tokens[0].Erc20Address))
	Nil(t.T(), err)

	// Check Balance
	tokenID := utils.GenerateTokenID(testChainID, tokens[0].Erc20Address)
	currBalance := balanceManager.GetBalance(tokenID)
	Equal(t.T(), big.NewInt(int64(testutil.DefaultMintAmount)), currBalance.Amount)
}
func (t *BalanceSuite) TestIncrementBalance() {
	// Setup test environment
	testCtx := t.GetTestContext()
	clients := make(map[uint32]EVMClient.EVM)
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, testChainID, t.T())
	evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[testChainID] = evmClient

	// Wallet
	testWallet, _ := wallet.FromRandom()
	testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
	Nil(t.T(), err)
	NotNil(t.T(), testContractHandler)

	// Get Tokens
	tokens := testContractHandler.Tokens()
	var assets []config.AssetConfig
	for _, token := range tokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: testChainID,
		}
		assets = append(assets, asset)
	}
	balanceManager, err := balance.NewBalanceManager(clients, assets, testWallet.Address())
	Nil(t.T(), err)
	NotNil(t.T(), balanceManager)

	// Populate
	err = balanceManager.GetAllOnChainBalances(testCtx)
	Nil(t.T(), err)

	tokenID := utils.GenerateTokenID(assets[0].ChainID, common.HexToAddress(assets[0].Address))
	initialBalance := big.NewInt(testutil.DefaultMintAmount)

	// Increment
	incrementAmount := big.NewInt(500)
	balanceManager.IncrementBalance(tokenID, incrementAmount)

	// Check the new balance
	newBalance := balanceManager.GetBalance(tokenID)
	expectedBalance := new(big.Int).Add(initialBalance, incrementAmount)
	Equal(t.T(), expectedBalance, newBalance.Amount)
}

func (t *BalanceSuite) TestDecrementBalance() {
	// Setup test environment
	testCtx := t.GetTestContext()
	clients := make(map[uint32]EVMClient.EVM)
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, testChainID, t.T())
	evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[testChainID] = evmClient

	// Wallet
	testWallet, _ := wallet.FromRandom()
	testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
	Nil(t.T(), err)
	NotNil(t.T(), testContractHandler)

	// Get Tokens
	tokens := testContractHandler.Tokens()
	var assets []config.AssetConfig
	for _, token := range tokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: testChainID,
		}
		assets = append(assets, asset)
	}
	balanceManager, err := balance.NewBalanceManager(clients, assets, testWallet.Address())
	NotNil(t.T(), balanceManager)

	// Populate
	err = balanceManager.GetAllOnChainBalances(testCtx)
	Nil(t.T(), err)

	tokenID := utils.GenerateTokenID(assets[0].ChainID, common.HexToAddress(assets[0].Address))
	initialBalance := big.NewInt(testutil.DefaultMintAmount)

	// Decrement
	incrementAmount := big.NewInt(500)
	balanceManager.DecrementBalance(tokenID, incrementAmount)

	// Check the new balance
	newBalance := balanceManager.GetBalance(tokenID)
	expectedBalance := new(big.Int).Sub(initialBalance, incrementAmount)
	Equal(t.T(), expectedBalance, newBalance.Amount)
}

func (t *BalanceSuite) TestBalanceFloat64() {
	// Setup test environment
	testCtx := t.GetTestContext()
	clients := make(map[uint32]EVMClient.EVM)
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, testChainID, t.T())
	evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[testChainID] = evmClient

	// Wallet
	testWallet, _ := wallet.FromRandom()
	testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
	Nil(t.T(), err)
	NotNil(t.T(), testContractHandler)

	// Get Tokens
	tokens := testContractHandler.Tokens()
	var assets []config.AssetConfig
	for _, token := range tokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: testChainID,
		}
		assets = append(assets, asset)
	}
	balanceManager, err := balance.NewBalanceManager(clients, assets, testWallet.Address())
	NotNil(t.T(), balanceManager)

	// Populate
	err = balanceManager.GetAllOnChainBalances(testCtx)
	Nil(t.T(), err)

	tokenID := utils.GenerateTokenID(assets[0].ChainID, common.HexToAddress(assets[0].Address))

	// Get balance
	balance := balanceManager.GetBalance(tokenID)
	NotNil(t.T(), balance)
	normBalance, err := balance.ToFloat64()
	Nil(t.T(), err)
	Equal(t.T(), float64(1), normBalance)
}
