package listener_test

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/listener"
)

func (l *ListenerTestSuite) TestListenForEvents() {
	_, handle := l.manager.GetCounter(l.GetTestContext(), l.backend)
	var wg sync.WaitGroup
	const iterations = 10
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
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
		}()
	}

	wg.Wait()

	startBlock, err := handle.DeployBlock(&bind.CallOpts{Context: l.GetTestContext()})
	l.NoError(err)

	cl, err := listener.NewChainListener(
		l.backend,
		l.store,
		handle.Address(),
		uint64(startBlock.Int64()),
		l.metrics,
		listener.WithNewBlockHandler(func(ctx context.Context, block uint64) error {
			fmt.Println(block)
			return nil
		}),
	)
	l.NoError(err)

	clSafe, err := listener.NewChainListener(
		l.backend,
		l.store,
		handle.Address(),
		uint64(startBlock.Int64()),
		l.metrics,
		listener.WithNewBlockHandler(func(ctx context.Context, block uint64) error {
			fmt.Println(block)
			return nil
		}),
		listener.WithFinalityMode("safe"),
		listener.WithBlockWait(10),
	)
	l.NoError(err)

	clFinalized, err := listener.NewChainListener(
		l.backend,
		l.store,
		handle.Address(),
		uint64(startBlock.Int64()),
		l.metrics,
		listener.WithNewBlockHandler(func(ctx context.Context, block uint64) error {
			fmt.Println(block)
			return nil
		}),
		listener.WithFinalityMode("finalized"),
		listener.WithBlockWait(10),
	)
	l.NoError(err)

	eventCount := 0

	// TODO: check for timeout,but it will be extremely obvious if it gets hit.
	listenCtx, cancel := context.WithCancel(l.GetTestContext())
	_ = cl.Listen(listenCtx, func(ctx context.Context, log types.Log) error {
		eventCount++

		if eventCount == iterations*2 {
			cancel()
		}

		return nil
	})

	_ = clSafe.Listen(listenCtx, func(ctx context.Context, log types.Log) error {
		eventCount++

		if eventCount == iterations*2 {
			cancel()
		}

		return nil
	})

	_ = clFinalized.Listen(listenCtx, func(ctx context.Context, log types.Log) error {
		eventCount++

		if eventCount == iterations*2 {
			cancel()
		}

		return nil
	})

	l.NotEqual(cl.LatestBlock(), clFinalized.LatestBlock(), clSafe.LatestBlock())

}
