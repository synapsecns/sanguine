package listener_test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"sync"
)

func (l *ListenerTestSuite) TestListenForEvents() {
	_, handle := l.manager.GetCounter(l.GetTestContext(), l.backend)
	var wg sync.WaitGroup
	const iterations = 50
	for i := 0; i < iterations; i++ {
		i := i
		wg.Add(1)
		go func(_ int) {
			defer wg.Done()

			auth := l.backend.GetTxContext(l.GetTestContext(), nil)

			//nolint:typecheck
			bridgeRequestTX, err := handle.IncrementCounter(auth.TransactOpts)
			l.NoError(err)
			l.NotNil(bridgeRequestTX)

			l.backend.WaitForConfirmation(l.GetTestContext(), bridgeRequestTX)

			bridgeResponseTX, err := handle.DecrementCounter(auth.TransactOpts)
			l.NoError(err)
			l.NotNil(bridgeResponseTX)
			l.backend.WaitForConfirmation(l.GetTestContext(), bridgeResponseTX)
		}(i)
	}

	wg.Wait()

	startBlock, err := handle.DeployBlock(&bind.CallOpts{Context: l.GetTestContext()})
	l.NoError(err)

	cl, err := listener.NewChainListener(l.backend, l.store, handle.Address(), uint64(startBlock.Int64()), l.metrics)
	l.NoError(err)

	eventCount := 0

	// TODO: check for timeout,but it will be extremely obvious if it gets hit.
	listenCtx, cancel := context.WithCancel(l.GetTestContext())
	err = cl.Listen(listenCtx, func(ctx context.Context, log types.Log) error {
		eventCount++

		if eventCount == iterations*2 {
			cancel()
		}

		return nil
	})
	l.NoError(err)
}
