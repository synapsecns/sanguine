package chainwatcher_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"time"
)

// simpleMockSubscriber is a mock subscriber designed to test channel draining.
// that is, if  we have a listener that starts at height 0 and ends at height 10,
// and we drain the channel, we should get 10.
type simpleMockSubscriber struct {
	height    uint64
	callCount uint64
}

var now = time.Now()

func (s *simpleMockSubscriber) LatestHeight(_ context.Context) (uint64, error) {
	defer func() {
		s.callCount++
	}()

	return s.height + uint64(time.Now().Sub(now).Seconds()*5), nil
}

func newSimpleMockSubscriber(startHeight uint64) *simpleMockSubscriber {
	return &simpleMockSubscriber{
		height: startHeight,
	}
}

var _ chainwatcher.BlockSubscriberClient = &simpleMockSubscriber{}

func (b *ChainWatcherSuite) TestBlockBroadcaster() {
	// don't wait
	chainwatcher.PollInterval = time.Nanosecond

	testChainID := gofakeit.Uint64()

	watcher := chainwatcher.NewBlockHeightWatcher(b.GetTestContext(), testChainID, newSimpleMockSubscriber(0))
	//time.Sleep(time.Second * 10)
	testSub := watcher.Subscribe()
	for {
		time.Sleep(time.Second)
		fmt.Println(getSubEndHeight(0, testSub))
	}
	//fmt.Println(getSubEndHeight(0, testSub))
	//fmt.Println(getSubEndHeight(0, testSub))

	//Equal(b.T(), getSubEndHeight(0, testSub), uint64(10))
}

func getSubEndHeight(startHeight uint64, heightChan <-chan uint64) (endHeight uint64) {
	// populate the default
	endHeight = startHeight

OUTER:
	for {
		select {
		case newHeight := <-heightChan:
			endHeight = newHeight
		default:
			break OUTER
		}
	}

	return endHeight
}
