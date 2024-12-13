package etherscan_test

import (
	"context"
	"time"

	"github.com/Flaque/filet"
	"github.com/synapsecns/sanguine/tools/abigen/internal/etherscan"
)

// PLACEHOLDER: EtherscanSuite is defined in suite_test.go

// PLACEHOLDER: TestEtherscanSuite function implementation

func (s *EtherscanSuite) TestRateLimiter() {
	waitTime := time.Second
	lockPath := filet.TmpDir(s.TestSuite.T(), "")

	rateLimiter, err := etherscan.NewFileRateLimiter(context.Background(), lockPath, waitTime)
	s.TestSuite.Require().NoError(err)

	for lockCount := 0; lockCount < 2; lockCount++ {
		expectedEndTime := time.Now().Add(waitTime)

		// obtain lock obtains the lock
		ok, err := rateLimiter.ObtainLock(context.Background())
		s.TestSuite.Assert().True(ok)
		s.TestSuite.Require().NoError(err)

		// release lock releases the lock
		ok, err = rateLimiter.ReleaseLock()
		s.TestSuite.Assert().True(ok)
		s.TestSuite.Require().NoError(err)

		s.TestSuite.Assert().GreaterOrEqual(expectedEndTime.UnixNano(), time.Now().UnixNano())
	}
}
