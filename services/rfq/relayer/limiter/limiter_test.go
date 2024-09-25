package limiter_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	evmMock "github.com/synapsecns/sanguine/ethergo/client/mocks"
	listenerMock "github.com/synapsecns/sanguine/ethergo/listener/mocks"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/limiter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
)

var (
	originAmount, _ = new(big.Float).SetFloat64(11983199386503458).Int(nil)
	destAmount, _   = new(big.Float).SetFloat64(11980157553261996).Int(nil)
	user            = common.HexToAddress("0x40D566c2581890B74a10bE7C5d3b5dDfE3F1062F")
)

func (l *LimiterSuite) TestOverLimitEnoughConfirmations() {
	quote := reldb.QuoteRequest{
		TransactionID: common.HexToHash("0xdeadbeef"),
		BlockNumber:   4,
		Sender:        user,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId:   10,
			DestChainId:     81457,
			OriginSender:    user,
			DestRecipient:   user,
			OriginToken:     util.EthAddress,
			DestToken:       util.EthAddress,
			OriginAmount:    originAmount,
			DestAmount:      destAmount,
			OriginFeeAmount: big.NewInt(0),
			SendChainGas:    false,
			Deadline:        big.NewInt(1727200795),
			Nonce:           big.NewInt(86730),
		},
	}
	mockQuoter := buildMockQuoter(100_000_000)
	mockListener := buildMockListener(6)
	packedBridgeTx, err := packBridgeTransaction(quote.Transaction)
	l.NoError(err)
	mockClient := buildMockEVMClient(quoteRequestToReceipt(quote, packedBridgeTx, false))

	l.limiter = limiter.NewRateLimiter(
		l.cfg,
		mockListener,
		mockQuoter,
		l.metrics,
		l.cfg.Chains[1].Tokens,
		mockClient,
	)

	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), &quote)
	l.NoError(err)
	l.True(allowed)

	// now test the case where the transaction is reverted (we find out from the log)
	mockClient = buildMockEVMClient(quoteRequestToReceipt(quote, packedBridgeTx, true))

	l.limiter = limiter.NewRateLimiter(
		l.cfg,
		mockListener,
		mockQuoter,
		l.metrics,
		l.cfg.Chains[1].Tokens,
		mockClient,
	)

	allowed, err = l.limiter.IsAllowed(l.GetTestContext(), &quote)
	l.NoError(err)
	l.False(allowed)
}

func (l *LimiterSuite) TestUnderLimitEnoughConfirmations() {
	quote := reldb.QuoteRequest{
		TransactionID: common.HexToHash("0xdeadbeef"),
		BlockNumber:   5,
		Sender:        user,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId:   10,
			DestChainId:     81457,
			OriginSender:    user,
			DestRecipient:   user,
			OriginToken:     util.EthAddress,
			DestToken:       util.EthAddress,
			OriginAmount:    big.NewInt(100),
			DestAmount:      big.NewInt(100),
			OriginFeeAmount: big.NewInt(0),
			SendChainGas:    false,
			Deadline:        big.NewInt(1727200795),
			Nonce:           big.NewInt(86730),
		},
	}
	mockQuoter := buildMockQuoter(100)
	mockListener := buildMockListener(10)
	packedBridgeTx, err := packBridgeTransaction(quote.Transaction)
	l.NoError(err)
	mockClient := buildMockEVMClient(quoteRequestToReceipt(quote, packedBridgeTx, false))

	l.limiter = limiter.NewRateLimiter(
		l.cfg,
		mockListener,
		mockQuoter,
		l.metrics,
		l.cfg.Chains[1].Tokens,
		mockClient,
	)

	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), &quote)
	l.NoError(err)
	l.True(allowed)
}

func (l *LimiterSuite) TestUnderLimitNotEnoughConfirmations() {
	quote := reldb.QuoteRequest{
		TransactionID: common.HexToHash("0xdeadbeef"),
		BlockNumber:   1, // same block number,but shouldnt matter
		Sender:        user,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId:   10,
			DestChainId:     81457,
			OriginSender:    user,
			DestRecipient:   user,
			OriginToken:     util.EthAddress,
			DestToken:       util.EthAddress,
			OriginAmount:    big.NewInt(100),
			DestAmount:      big.NewInt(100),
			OriginFeeAmount: big.NewInt(0),
			SendChainGas:    false,
			Deadline:        big.NewInt(1727200795),
			Nonce:           big.NewInt(86730),
		},
	}
	mockQuoter := buildMockQuoter(100)
	mockListener := buildMockListener(1)
	packedBridgeTx, err := packBridgeTransaction(quote.Transaction)
	l.NoError(err)
	mockClient := buildMockEVMClient(quoteRequestToReceipt(quote, packedBridgeTx, false))

	l.limiter = limiter.NewRateLimiter(
		l.cfg,
		mockListener,
		mockQuoter,
		l.metrics,
		l.cfg.Chains[1].Tokens,
		mockClient,
	)

	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), &quote)
	l.NoError(err)
	l.True(allowed)
}

func (l *LimiterSuite) TestOverLimitNotEnoughConfirmations() {
	quote := reldb.QuoteRequest{
		TransactionID: common.HexToHash("0xdeadbeef"),
		BlockNumber:   4,
		Sender:        user,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId:   10,
			DestChainId:     81457,
			OriginSender:    user,
			DestRecipient:   user,
			OriginToken:     util.EthAddress,
			DestToken:       util.EthAddress,
			OriginAmount:    originAmount,
			DestAmount:      destAmount,
			OriginFeeAmount: big.NewInt(0),
			SendChainGas:    false,
			Deadline:        big.NewInt(1727200795),
			Nonce:           big.NewInt(86730),
		},
	}
	mockQuoter := buildMockQuoter(100_000_000) // eth price per
	mockListener := buildMockListener(4)

	packedBridgeTx, err := packBridgeTransaction(quote.Transaction)
	l.NoError(err)
	mockClient := buildMockEVMClient(quoteRequestToReceipt(quote, packedBridgeTx, false))

	l.limiter = limiter.NewRateLimiter(l.cfg,
		mockListener,
		mockQuoter,
		l.metrics,
		l.cfg.Chains[1].Tokens,
		mockClient,
	)

	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), &quote)
	l.NoError(err)
	l.False(allowed)
}

// returns a mock quoter that quotes the given price.
func buildMockQuoter(price float64) *mocks.Quoter {
	mockQuoter := new(mocks.Quoter)
	mockQuoter.On("GetPrice", mock.Anything, mock.Anything).Return(price, nil)
	return mockQuoter
}

// returns a mock listener that returns the given block number.
func buildMockListener(blockNumber uint64) *listenerMock.ContractListener {
	mockClient := new(listenerMock.ContractListener)
	mockClient.On("LatestBlock").Return(blockNumber, nil)
	return mockClient
}

// returns a mock EVM client that returns the given receipt.
func buildMockEVMClient(receipt *types.Receipt) *evmMock.EVM {
	mockClient := new(evmMock.EVM)
	mockClient.On("TransactionReceipt", mock.Anything, mock.Anything).Return(receipt, nil)
	return mockClient
}

func packBridgeTransaction(tx fastbridge.IFastBridgeBridgeTransaction) ([]byte, error) {
	bridgeStruct, err := abi.NewType(
		"tuple",
		"BridgeTransaction",
		[]abi.ArgumentMarshaling{
			{Name: "originChainId", Type: "uint32"},
			{Name: "destChainId", Type: "uint32"},
			{Name: "originSender", Type: "address"},
			{Name: "destRecipient", Type: "address"},
			{Name: "originToken", Type: "address"},
			{Name: "destToken", Type: "address"},
			{Name: "originAmount", Type: "uint256"},
			{Name: "destAmount", Type: "uint256"},
			{Name: "originFeeAmount", Type: "uint256"},
			{Name: "sendChainGas", Type: "bool"},
			{Name: "deadline", Type: "uint256"},
			{Name: "nonce", Type: "uint256"},
		})
	if err != nil {
		//nolint: wrapcheck
		return []byte{}, err
	}

	args := abi.Arguments{
		{Type: bridgeStruct, Name: "param_one"},
	}

	packed, err := args.Pack(&tx)
	if err != nil {
		//nolint: wrapcheck
		return []byte{}, err
	}
	return packed, nil
}

func quoteRequestToReceipt(q reldb.QuoteRequest, data []byte, removed bool) *types.Receipt {
	// Basically, I used some random RFQ transaction as the test case for this suite,
	// but mocked the block and price that's returned instead of changing the actual transaction fields.
	// This is why the data is hardcoded.
	// TODO: fix this once I figure out how to properly create the log data
	prefixDataStr := "0x00000000000000000000000000000000000000000000000000000000000000e0" +
		"0000000000000000000000000000000000000000000000000000000000013e31" +
		"000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" +
		"000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" +
		"000000000000000000000000000000000000000000000000002a92a806b48922" +
		"000000000000000000000000000000000000000000000000002a8fe3cb50bdac" +
		"0000000000000000000000000000000000000000000000000000000000000000" +
		"0000000000000000000000000000000000000000000000000000000000000180"

	prefixData := common.Hex2Bytes(prefixDataStr[2:])
	fullData := append(prefixData, data...)
	return &types.Receipt{
		Logs: []*types.Log{
			{
				Address: common.HexToAddress("0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E"),
				Topics: []common.Hash{
					common.HexToHash("0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a"),
					q.TransactionID,
					common.HexToHash(q.Transaction.OriginSender.String()),
				},
				Data:    fullData,
				Removed: removed,
				TxIndex: 69,
			},
		},
	}
}
