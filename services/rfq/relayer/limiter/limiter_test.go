package limiter_test

import (
	"math/big"
	"time"

	"github.com/stretchr/testify/mock"
	clientMocks "github.com/synapsecns/sanguine/ethergo/client/mocks"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/limiter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter/mocks"
		"github.com/synapsecns/sanguine/services/rfq/util"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (l *LimiterSuite) TestOverLimitEnoughConfirmations() {
	mockQuoter := new(mocks.Quoter)
	mockQuoter.On("GetPrice", mock.Anything, mock.Anything).Once().Return(float64(10001), nil)

	mockClient := new(clientMocks.EVM)
	// confirmations are enough here
	mockClient.On("BlockNumber", mock.Anything, mock.Anything).Once().Return(uint64(6), nil)

	l.limiter = limiter.NewRateLimiter(l.cfg, mockQuoter, mockClient, l.metrics, l.cfg.Chains[1].Tokens)

	quote := reldb.QuoteRequest{
		BlockNumber: 5,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: 1,
			DestChainId:   2,
			OriginToken:   util.EthAddress,
			OriginAmount:  big.NewInt(100),
			DestAmount:    big.NewInt(100),
			Deadline:      big.NewInt(time.Now().Unix()),
			Nonce:         big.NewInt(0),
		},
	}

	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), quote)
	l.NoError(err)
	l.True(allowed)
}

func (l *LimiterSuite) TestUnderLimitEnoughConfirmations() {
	mockQuoter := new(mocks.Quoter)
	mockQuoter.On("GetPrice", mock.Anything, mock.Anything).Once().Return(float64(100), nil)

	mockClient := new(clientMocks.EVM)
	// confirmations are enough here
	mockClient.On("BlockNumber", mock.Anything, mock.Anything).Once().Return(uint64(10), nil)
	l.limiter = limiter.NewRateLimiter(l.cfg, mockQuoter, mockClient, l.metrics, l.cfg.Chains[1].Tokens)

	quote := reldb.QuoteRequest{
		BlockNumber: 5,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: 1,
			DestChainId:   2,
			OriginToken:   util.EthAddress,
			OriginAmount:  big.NewInt(100),
			DestAmount:    big.NewInt(100),
		},
	}
	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), quote)
	l.NoError(err)
	l.True(allowed)
}

func (l *LimiterSuite) TestUnderLimitNotEnoughConfirmations() {
	mockQuoter := new(mocks.Quoter)
	mockQuoter.On("GetPrice", mock.Anything, mock.Anything).Once().Return(float64(100), nil)

	mockClient := new(clientMocks.EVM)

	mockClient.On("BlockNumber", mock.Anything, mock.Anything).Once().Return(uint64(1), nil)
	l.limiter = limiter.NewRateLimiter(l.cfg, mockQuoter, mockClient, l.metrics, l.cfg.Chains[1].Tokens)

	quote := reldb.QuoteRequest{
		BlockNumber: 1, // same block number,but shouldnt matter
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: 1,
			DestChainId:   2,
			OriginToken:   util.EthAddress,
			OriginAmount:  big.NewInt(100),
			DestAmount:    big.NewInt(100),
		},
	}

	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), quote)
	l.NoError(err)
	l.True(allowed)
}

func (l *LimiterSuite) TestOverLimitNotEnoughConfirmations() {
	mockQuoter := new(mocks.Quoter)
	mockQuoter.On("GetPrice", mock.Anything, mock.Anything).Once().Return(float64(69420), nil)

	mockClient := new(clientMocks.EVM)

	mockClient.On("BlockNumber", mock.Anything, mock.Anything).Once().Return(uint64(4), nil)
	l.limiter = limiter.NewRateLimiter(l.cfg, mockQuoter, mockClient, l.metrics, l.cfg.Chains[1].Tokens)

	quote := reldb.QuoteRequest{
		BlockNumber: 4,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: 1,
			DestChainId:   2,
			OriginToken:   util.EthAddress,
			OriginAmount:  big.NewInt(100),
			DestAmount:    big.NewInt(100),
		},
	}

	allowed, err := l.limiter.IsAllowed(l.GetTestContext(), quote)
	l.NoError(err)
	l.False(allowed)
}

func (l *LimiterSuite) TestRateLimitOverWindow() {
	l.T().Skip("TODO: implement the sliding window: queue up requests and process them in order if cumulative volume is above limit")
}
