package submitter_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

func (s *SubmitterSuite) TestSortTxes() {
	expected := make(map[uint64][]*types.Transaction)
	var allTxes []db.TX
	var mapMux sync.Mutex
	var sliceMux sync.Mutex

	chainIDS := []int64{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(len(chainIDS))

	for i := range chainIDS {
		chainID := big.NewInt(chainIDS[i])
		go func() {
			defer wg.Done()
			backend := simulated.NewSimulatedBackendWithChainID(s.GetTestContext(), s.T(), chainID)

			testAddress := backend.GetTxContext(s.GetTestContext(), nil)
			testKey := &keystore.Key{PrivateKey: testAddress.PrivateKey, Address: testAddress.From}
			for i := 0; i < 50; i++ {
				mockTX := mocks.MockTx(s.GetTestContext(), s.T(), backend, testKey, types.DynamicFeeTxType)

				// add to map in order
				mapMux.Lock()
				expected[chainID.Uint64()] = append(expected[chainID.Uint64()], mockTX)
				mapMux.Unlock()

				sliceMux.Lock()
				tx := db.TX{
					Transaction: mockTX,
					Status:      db.Stored,
				}
				tx.UnsafeSetCreationTime(time.Now())

				allTxes = append(allTxes, tx)
				// shuffle the slice each time
				rand.Shuffle(len(allTxes), func(i, j int) {
					allTxes[i], allTxes[j] = allTxes[j], allTxes[i]
				})
				sliceMux.Unlock()
			}
		}()
	}
	wg.Wait()

	sorted := submitter.SortTxes(allTxes)
	assert.Equal(s.T(), len(sorted), len(expected))
	for chainID, txes := range expected {
		for i := range txes {
			assert.Equal(s.T(), sorted[chainID][i].Hash(), txes[i].Hash())
		}
	}
}

func (s *SubmitterSuite) TestGroupTxesByNonce() {
	ogTx := mocks.GetMockTxes(s.GetTestContext(), s.T(), 1, types.LegacyTxType)[0]
	var txes []db.TX
	// generate 1,000 txes with 100 different nonces
	for nonce := 0; nonce < 100; nonce++ {
		copiedTX, err := util.CopyTX(ogTx, util.WithNonce(uint64(nonce)))
		s.Require().NoError(err)

		for i := 0; i < 10; i++ {
			newTX, err := util.CopyTX(copiedTX, util.WithGasPrice(big.NewInt(int64(i))))
			s.Require().NoError(err)

			txes = append(txes, db.TX{
				Transaction: newTX,
				Status:      db.Pending,
			})
		}
	}

	nonceMap := submitter.GroupTxesByNonce(txes)
	for i := 0; i < 100; i++ {
		txList := nonceMap[uint64(i)]
		for _, tx := range txList {
			if tx.Nonce() != uint64(i) {
				s.Require().NoError(fmt.Errorf("expected nonce %d, got %d", i, tx.Nonce()))
			}
		}
	}
}
