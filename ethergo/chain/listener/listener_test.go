package listener_test

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/ethergo/chain/listener"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemock"
)

func (l *ListenerTestSuite) TestListenForEvents() {
	_, handle := l.manager.GetMockFastBridge(l.GetTestContext(), l.backend)
	var wg sync.WaitGroup
	const iterations = 50
	for i := 0; i < iterations; i++ {
		i := i
		go func(num int) {
			wg.Add(1)
			defer wg.Done()

			testAddress := common.BigToAddress(big.NewInt(int64(i)))
			auth := l.backend.GetTxContext(l.GetTestContext(), nil)

			//nolint: typecheck
			txID := [32]byte(crypto.Keccak256(testAddress.Bytes()))
			bridgeRequestTX, err := handle.MockBridgeRequest(auth.TransactOpts, txID, testAddress, fastbridgemock.IFastBridgeBridgeParams{
				DstChainId:   gofakeit.Uint32(),
				Sender:       testAddress,
				To:           testAddress,
				OriginToken:  testAddress,
				DestToken:    testAddress,
				OriginAmount: new(big.Int).SetUint64(gofakeit.Uint64()),
				DestAmount:   new(big.Int).SetUint64(gofakeit.Uint64()),
				SendChainGas: false,
				Deadline:     new(big.Int).SetUint64(uint64(time.Now().Add(-1 * time.Second * time.Duration(gofakeit.Uint16())).Unix())),
			})
			l.NoError(err)
			l.NotNil(bridgeRequestTX)

			l.backend.WaitForConfirmation(l.GetTestContext(), bridgeRequestTX)

			bridgeResponseTX, err := handle.MockBridgeRelayer(auth.TransactOpts,
				// transactionID
				txID,
				// relayer
				testAddress,
				// to
				testAddress,
				// originChainID
				uint32(gofakeit.Uint16()),
				// originToken
				testAddress,
				// destToken
				testAddress,
				// originAmount
				new(big.Int).SetUint64(gofakeit.Uint64()),
				// destAmount
				new(big.Int).SetUint64(gofakeit.Uint64()),
				// gasAmount
				new(big.Int).SetUint64(gofakeit.Uint64()))
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
}
