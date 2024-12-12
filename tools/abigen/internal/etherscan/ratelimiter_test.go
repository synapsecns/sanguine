package etherscan_test

import (
	"context"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/tools/abigen/internal/etherscan"
)

// PLACEHOLDER: EtherscanSuite is defined in suite_test.go

// PLACEHOLDER: TestEtherscanSuite function implementation

func (s *EtherscanSuite) TestRateLimiter() {
	waitTime := time.Second
	lockPath := filet.TmpDir(s.T(), "")

	rateLimiter, err := etherscan.NewFileRateLimiter(context.Background(), lockPath, waitTime)
	s.Require().NoError(err)

	for lockCount := 0; lockCount < 2; lockCount++ {
		expectedEndTime := time.Now().Add(waitTime)

		// obtain lock obtains the lock
		ok, err := rateLimiter.ObtainLock(context.Background())
		s.Assert().True(ok)
		s.Require().NoError(err)

		// release lock releases the lock
		ok, err = rateLimiter.ReleaseLock()
		s.Assert().True(ok)
		s.Require().NoError(err)

		s.Assert().GreaterOrEqual(expectedEndTime.UnixNano(), time.Now().UnixNano())
	}
}
