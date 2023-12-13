package quote_test

import (
	"fmt"
	"math/big"

	. "github.com/stretchr/testify/assert"
	EVMClient "github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/quote"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/testutil"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"
)

func (t *QuoteSuite) TestNewQuoter() {
	testCtx := t.GetTestContext()
	clients := make(map[uint32]EVMClient.EVM)
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, t.T(), testChainID)
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

	quoter, err := quote.NewQuoter(testCtx, clients, assets, testWallet.Address(), "", t.signer)
	Nil(t.T(), err)
	NotNil(t.T(), quoter)
}

func (t *QuoteSuite) TestGetValidQuote() {
	testCtx := t.GetTestContext()

	// Wallet
	testWallet, _ := wallet.FromRandom()

	clients := make(map[uint32]EVMClient.EVM)
	originChainID := uint32(42161)
	originAnvilBackend := testutil.NewAnvilBackend(testCtx, t.T(), originChainID)
	originEvmClient, err := testutil.NewEVMClientFromAnvil(testCtx, originAnvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[originChainID] = originEvmClient
	originContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, originAnvilBackend, testWallet, originChainID)
	Nil(t.T(), err)
	NotNil(t.T(), originContractHandler)

	destChainID := uint32(1)
	destAnvilBackend := testutil.NewAnvilBackend(testCtx, t.T(), destChainID)
	destEvmClient, err := testutil.NewEVMClientFromAnvil(testCtx, destAnvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[destChainID] = destEvmClient
	destContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, destAnvilBackend, testWallet, destChainID)
	Nil(t.T(), err)
	NotNil(t.T(), destContractHandler)

	var assets []config.AssetConfig
	// Get Origin Tokens
	originTokens := originContractHandler.Tokens()
	for _, token := range originTokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: originChainID,
		}
		assets = append(assets, asset)
	}

	destTokens := destContractHandler.Tokens()
	for _, token := range destTokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: destChainID,
		}
		assets = append(assets, asset)
	}

	quoter, err := quote.NewQuoter(testCtx, clients, assets, testWallet.Address(), "", t.signer)
	Nil(t.T(), err)
	NotNil(t.T(), quoter)

	originTestToken := originTokens[0].Erc20Address
	destTestToken := destTokens[0].Erc20Address
	destTokenID := utils.GenerateTokenID(destChainID, destTestToken)
	quoteID := utils.GenerateQuoteID(originChainID, originTestToken, destChainID, destTestToken)

	testAmount := big.NewInt(1000000)

	q, err := quoter.GetValidQuote(quoteID, destTokenID, testAmount)
	Nil(t.T(), err)
	NotNil(t.T(), q)
}

func (t *QuoteSuite) TestQuoteToAPIQuote() {
	testCtx := t.GetTestContext()

	// Wallet
	testWallet, _ := wallet.FromRandom()

	clients := make(map[uint32]EVMClient.EVM)
	originChainID := uint32(42161)
	originAnvilBackend := testutil.NewAnvilBackend(testCtx, t.T(), originChainID)
	originEvmClient, err := testutil.NewEVMClientFromAnvil(testCtx, originAnvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[originChainID] = originEvmClient
	originContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, originAnvilBackend, testWallet, originChainID)
	Nil(t.T(), err)
	NotNil(t.T(), originContractHandler)

	destChainID := uint32(1)
	destAnvilBackend := testutil.NewAnvilBackend(testCtx, t.T(), destChainID)
	destEvmClient, err := testutil.NewEVMClientFromAnvil(testCtx, destAnvilBackend, t.metrics)
	Nil(t.T(), err)
	clients[destChainID] = destEvmClient
	destContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, destAnvilBackend, testWallet, destChainID)
	Nil(t.T(), err)
	NotNil(t.T(), destContractHandler)

	var assets []config.AssetConfig
	// Get Origin Tokens
	originTokens := originContractHandler.Tokens()
	for _, token := range originTokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: originChainID,
		}
		assets = append(assets, asset)
	}

	destTokens := destContractHandler.Tokens()
	for _, token := range destTokens {
		asset := config.AssetConfig{
			Address: token.Erc20Address.String(),
			ChainID: destChainID,
		}
		assets = append(assets, asset)
	}

	quoter, err := quote.NewQuoter(testCtx, clients, assets, testWallet.Address(), "", t.signer)
	Nil(t.T(), err)
	NotNil(t.T(), quoter)

	originTestToken := originTokens[0].Erc20Address
	destTestToken := destTokens[0].Erc20Address
	quoteID := utils.GenerateQuoteID(originChainID, originTestToken, destChainID, destTestToken)
	quotes := quoter.GetQuotes(quoteID)
	NotNil(t.T(), quotes)
	Equal(t.T(), 1, len(quotes))

	qt := quotes[0]
	NotNil(t.T(), qt)
	apiQuote, err := quoter.QuoteToAPIQuote(qt)
	Nil(t.T(), err)
	NotNil(t.T(), apiQuote)
	Equal(t.T(), qt.OriginToken.String(), apiQuote.OriginToken)
	Equal(t.T(), qt.DestToken.String(), apiQuote.DestToken)
	Equal(t.T(), uint(qt.OriginChainID), apiQuote.OriginChainID)
	Equal(t.T(), uint(qt.DestChainID), apiQuote.DestChainID)
	Equal(t.T(), fmt.Sprintf("%d", int(testutil.DefaultMintAmount)), apiQuote.OriginAmount)
	Equal(t.T(), fmt.Sprintf("%d", int(testutil.DefaultMintAmount)), apiQuote.DestAmount)
	Equal(t.T(), uint8(18), apiQuote.OriginDecimals)
	Equal(t.T(), uint8(18), apiQuote.DestDecimals)
}
