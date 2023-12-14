package listener_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"math/big"
	"sync"
)

func (l *ListenerTestSuite) TestListenForEvents() {
	_, handle := l.manager.GetMockFastBridge(l.GetTestContext(), l.backend)
	var wg sync.WaitGroup
	const iterations = 50
	for i := 0; i < iterations; i++ {
		go func(num int) {
			wg.Add(1)
			defer wg.Done()

			testAddress := common.BigToAddress(big.NewInt(int64(i)))
			auth := l.backend.GetTxContext(l.GetTestContext(), nil)

			txID := [32]byte(crypto.Keccak256(testAddress.Bytes()))
			bridgeRequestTX, err := handle.MockBridgeRequestRaw(auth.TransactOpts, txID, testAddress, []byte(gofakeit.Sentence(10)))
			l.NoError(err)

			l.backend.WaitForConfirmation(l.GetTestContext(), bridgeRequestTX)

			bridgeResponseTX, err := handle.MockBridgeRelayer(auth.TransactOpts, txID, testAddress, testAddress, testAddress, new(big.Int).SetUint64(gofakeit.Uint64()))
			l.NoError(err)
			l.backend.WaitForConfirmation(l.GetTestContext(), bridgeResponseTX)
		}(i)
	}

	wg.Wait()

	cl, err := listener.NewChainListener(l.backend, l.store, handle.Address(), l.metrics)
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
}
