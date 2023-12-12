package testutil_test

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/testutil"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"

	"time"
)

func TestNewTestContractHandlerImpl(t *testing.T) {
	testCtx := context.Background()

	// Origin Chain
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, testChainID, t)

	// Wallet
	testWallet, _ := wallet.FromRandom()
	testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
	Nil(t, err)
	NotNil(t, testContractHandler)
}

func TestTokens(t *testing.T) {
	testCtx := context.Background()

	// Origin Chain
	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, testChainID, t)

	// Wallet
	testWallet, _ := wallet.FromRandom()
	testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
	Nil(t, err)
	NotNil(t, testContractHandler)

	tokens := testContractHandler.Tokens()
	for _, token := range tokens {
		NotNil(t, token.Erc20Address)
		NotNil(t, token.Erc20Contract)

		// Check Allowance
		allowance, err := token.Erc20Contract.Allowance(nil, testWallet.Address(), testContractHandler.FastBridgeAddress())
		Nil(t, err)
		Equal(t, big.NewInt(testutil.DefaultMintAmount), allowance)

		// Check Balance
		balance, err := token.Erc20Contract.BalanceOf(nil, testWallet.Address())
		Nil(t, err)
		Equal(t, big.NewInt(testutil.DefaultMintAmount), balance)
	}
}

// TestBridge tests the entire bridge process (bridge, prove, claim).
func TestBridge(t *testing.T) {
	testCtx := context.Background()
	// Wallet
	testWallet, _ := wallet.FromRandom()

	// Origin Chain
	originChainID := uint32(42161)
	anvilOriginBackend := testutil.NewAnvilBackend(testCtx, originChainID, t)
	originContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilOriginBackend, testWallet, originChainID)
	Nil(t, err)
	NotNil(t, originContractHandler)
	originTokens := originContractHandler.Tokens()

	// Destination Chain
	destinationChainID := uint32(1)
	anvilDestinationBackend := testutil.NewAnvilBackend(testCtx, destinationChainID, t)
	destinationContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilDestinationBackend, testWallet, destinationChainID)
	Nil(t, err)
	NotNil(t, destinationContractHandler)
	destinationTokens := destinationContractHandler.Tokens()

	// Bridge ---------------------
	bridgeParams := bindings.IFastBridgeBridgeParams{
		DstChainId:   destinationChainID,
		To:           testWallet.Address(),
		OriginToken:  originTokens[0].Erc20Address,
		DestToken:    destinationTokens[0].Erc20Address,
		OriginAmount: big.NewInt(params.GWei),
		DestAmount:   big.NewInt(params.GWei),
		// deadline must be 30 minutes in the future (based on current contract)
		Deadline: big.NewInt(time.Now().Unix() + 4000),
	}
	tx, err := originContractHandler.FBExecuteBridge(testCtx, bridgeParams)
	Nil(t, err)
	NotNil(t, tx)
	Equal(t, int64(originChainID), tx.ChainId().Int64())
	Equal(t, originContractHandler.FastBridgeAddress(), *tx.To())

	// Get the transaction receipt to get logs to get request for next test
	receipt, err := anvilOriginBackend.TransactionReceipt(testCtx, tx.Hash())
	Nil(t, err)

	// Get ABI
	parsedABI, err := abi.JSON(strings.NewReader(bindings.FastBridgeMetaData.ABI))
	Nil(t, err)

	event := new(bindings.FastBridgeBridgeRequested)

	for _, log := range receipt.Logs {
		// Check if the log is a BridgeRequested event
		if log.Topics[0] == parsedABI.Events["BridgeRequested"].ID {
			// Unpack the event
			logErr := parsedABI.UnpackIntoInterface(event, "BridgeRequested", log.Data)
			Nil(t, logErr)
		}
	}

	// Check Parity
	bridgeTransaction, err := utils.Decode(event.Request)
	Nil(t, err)
	Equal(t, bridgeParams.DstChainId, bridgeTransaction.DestChainId)
	Equal(t, bridgeParams.OriginToken, bridgeTransaction.OriginToken)
	Equal(t, bridgeParams.DestToken, bridgeTransaction.DestToken)
	Equal(t, bridgeParams.OriginAmount, bridgeTransaction.OriginAmount)
	Equal(t, bridgeParams.DestAmount, bridgeTransaction.DestAmount)
	Equal(t, bridgeParams.Deadline, bridgeTransaction.Deadline)

	// Relay ---------------------
	destTx, err := destinationContractHandler.FBExecuteRelay(testCtx, event.Request)
	Nil(t, err)

	// Prove ---------------------
	var txHash [32]byte
	txHash = destTx.Hash()
	tx, err = originContractHandler.FBExecuteProve(testCtx, event.Request, txHash)
	Nil(t, err)
	NotNil(t, tx)

	// Not going to test claim here because it's going to take 30 minutes
}
