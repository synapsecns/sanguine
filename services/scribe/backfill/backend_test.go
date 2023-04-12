package backfill_test

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"k8s.io/apimachinery/pkg/util/sets"
	"math/big"
	"sync"
)

// startOmnirpcServer boots an omnirpc server for an rpc address.
// the url for this rpc is returned.
func (b *BackfillSuite) startOmnirpcServer(ctx context.Context, backend backends.SimulatedTestBackend) string {
	baseHost := testhelper.NewOmnirpcServer(ctx, b.T(), backend)
	return testhelper.GetURL(baseHost, backend)
}

// ReachBlockHeight reaches a block height on a backend.
func (b *BackfillSuite) ReachBlockHeight(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			b.T().Log(ctx.Err())
			return
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))

		latestBlock, err := backend.BlockNumber(ctx)
		Nil(b.T(), err)

		if latestBlock >= desiredBlockHeight {
			return
		}
	}
}

// ReachBlockHeight reaches a block height on a backend.
func (b *BackfillSuite) PopuluateWithLogs(ctx context.Context, backend backends.SimulatedTestBackend, desiredBlockHeight uint64) common.Address {
	i := 0
	var address common.Address
	for {
		select {
		case <-ctx.Done():
			b.T().Log(ctx.Err())
			return address
		default:
			// continue
		}
		i++
		backend.FundAccount(ctx, common.BigToAddress(big.NewInt(int64(i))), *big.NewInt(params.Wei))
		testContract, testRef := b.manager.GetTestContract(b.GetTestContext(), backend)
		address = testContract.Address()
		transactOpts := backend.GetTxContext(b.GetTestContext(), nil)
		tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
		Nil(b.T(), err)
		backend.WaitForConfirmation(b.GetTestContext(), tx)

		latestBlock, err := backend.BlockNumber(ctx)
		Nil(b.T(), err)

		if latestBlock >= desiredBlockHeight {
			return address
		}
	}
}

func (b *BackfillSuite) TestLogsInRange() {
	testBackend := geth.NewEmbeddedBackend(b.GetTestContext(), b.T())
	// start an omnirpc proxy and run 10 test tranactions so we can batch call blocks
	//  1-10
	var wg sync.WaitGroup
	wg.Add(2)

	const desiredBlockHeight = 10
	var commonAddress common.Address
	go func() {
		defer wg.Done()
		commonAddress = b.PopuluateWithLogs(b.GetTestContext(), testBackend, desiredBlockHeight)
	}()

	var host string
	go func() {
		defer wg.Done()
		host = b.startOmnirpcServer(b.GetTestContext(), testBackend)
	}()

	wg.Wait()

	scribeBackend, err := backfill.DialBackend(b.GetTestContext(), host, b.metrics)
	Nil(b.T(), err)

	chainID, err := scribeBackend.ChainID(b.GetTestContext())
	Nil(b.T(), err)

	res, err := backfill.GetLogsInRange(b.GetTestContext(), scribeBackend, 1, 10, 1, commonAddress, chainID.Uint64())
	Nil(b.T(), err)

	// use to make sure we don't double use values
	intSet := sets.NewInt64()

	itr := res.Iterator()

	for !itr.Done() {
		index, _ := itr.Next()

		Falsef(b.T(), intSet.Has(int64(index)), "%d appears at least twice", index)
		intSet.Insert(int64(index))
	}
}
