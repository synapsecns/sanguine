package etherscan_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/tools/abigen/internal/etherscan"
	"time"
)

var waitTime = time.Second

func (s EtherscanSuite) TestRateLimiter() {
	lockPath := filet.TmpDir(s.T(), "")

	rateLimiter, err := etherscan.NewFileRateLimiter(s.GetTestContext(), lockPath, waitTime)
	Nil(s.T(), err)

	for lockCount := 0; lockCount < 2; lockCount++ {
		expectedEndTime := time.Now().Add(waitTime)
		// obtain lock obtains the lcok
		ok, err := rateLimiter.ObtainLock(s.GetTestContext())
		True(s.T(), ok)
		Nil(s.T(), err)

		// release lock releases the lcok
		ok, err = rateLimiter.ReleaseLock()
		True(s.T(), ok)
		Nil(s.T(), err)

		GreaterOrEqual(s.T(), expectedEndTime.UnixNano(), time.Now().UnixNano())
	}
}
