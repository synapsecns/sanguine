package submitter_test

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/go-cmp/cmp"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/util"
	"gotest.tools/assert"
)

func TestCopyTransactOpts(t *testing.T) {
	// Test case 1: All fields populated
	opts1 := &bind.TransactOpts{
		From:      common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Nonce:     big.NewInt(1),
		Signer:    nil,
		Value:     big.NewInt(100),
		GasPrice:  big.NewInt(200),
		GasFeeCap: big.NewInt(300),
		GasTipCap: big.NewInt(400),
		GasLimit:  500,
		Context:   context.Background(),
		NoSend:    true,
	}

	copyOpts1 := submitter.CopyTransactOpts(opts1)
	assertTransactOptsEquality(t, opts1, copyOpts1)

	// Test case 2: Some fields populated, others nil
	opts2 := &bind.TransactOpts{
		Nonce:     nil,
		Signer:    nil,
		Value:     nil,
		GasPrice:  nil,
		GasFeeCap: big.NewInt(300),
		GasTipCap: big.NewInt(400),
		GasLimit:  500,
		Context:   context.Background(),
		NoSend:    true,
	}

	copyOpts2 := submitter.CopyTransactOpts(opts2)
	assertTransactOptsEquality(t, opts2, copyOpts2)

	// Test case 3: All fields nil
	opts3 := &bind.TransactOpts{
		Nonce:     nil,
		Signer:    nil,
		Value:     nil,
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   nil,
		NoSend:    true,
	}

	copyOpts3 := submitter.CopyTransactOpts(opts3)
	assertTransactOptsEquality(t, opts3, copyOpts3)
}

func assertTransactOptsEquality(tb testing.TB, toA, toB *bind.TransactOpts) {
	tb.Helper()
	// Check that the pointer values of the big integer fields are different
	assertBigIntsCopiedEqual(tb, toA.Nonce, toB.Nonce, "Nonce")
	assertBigIntsCopiedEqual(tb, toA.Value, toB.Value, "Value")
	assertBigIntsCopiedEqual(tb, toA.GasPrice, toB.GasPrice, "GasPrice")
	assertBigIntsCopiedEqual(tb, toA.GasFeeCap, toB.GasFeeCap, "GasFeeCap")
	assertBigIntsCopiedEqual(tb, toA.GasTipCap, toB.GasTipCap, "GasFeeCap")
	assert.DeepEqual(tb, toA, toB, testsuite.BigIntComparer(), cmp.AllowUnexported(context.Background()))
}

// assertBigIntsCopiedEqual checks that the given big.Ints are equal and that
// they have different pointers.
func assertBigIntsCopiedEqual(tb testing.TB, original *big.Int, newVal *big.Int, fieldName string) {
	tb.Helper()
	if original == nil && newVal == nil {
		return
	}

	if core.ArePointersEqual(original, newVal) {
		tb.Errorf("%s has same pointer as original", fieldName)
	}

	if original.Cmp(newVal) != 0 {
		tb.Errorf("%s is not equal", fieldName)
	}
}

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

	sorted := submitter.SortTxes(allTxes, 50)
	assert.Equal(s.T(), len(sorted), len(expected))
	for chainID, txes := range expected {
		for i := range txes {
			assert.Equal(s.T(), sorted[chainID][i].Hash(), txes[i].Hash())
		}
	}

	// check tx cap
	numTxes := 10
	sorted = submitter.SortTxes(allTxes, numTxes)
	assert.Equal(s.T(), len(sorted), len(expected))
	for chainID, txes := range expected {
		chainTxes := txes[:numTxes]
		for i := range chainTxes {
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

func (s *SubmitterSuite) TestOpStackGas() {

	mockTx := mocks.GetMockTxes(s.GetTestContext(), s.T(), 1, types.LegacyTxType)[0]

	fmt.Printf("Original Transaction Gas Limit: %d\n", mockTx.Gas())

}

func TestBox(t *testing.T) {
	const testTxCount = 10
	mockTx := mocks.GetMockTxes(context.Background(), t, testTxCount, 0)

	fmt.Printf("Original Transaction Gas Limit: %d\n", mockTx[0].Gas())
}

// Test for the outersection function.
func TestOutersection(t *testing.T) {
	set := []*big.Int{
		big.NewInt(2),
		big.NewInt(4),
	}

	superset := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		big.NewInt(4),
		big.NewInt(5),
	}

	expected := []*big.Int{
		big.NewInt(1),
		big.NewInt(3),
		big.NewInt(5),
	}

	result := submitter.Outersection(set, superset)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d elements, but got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v.Cmp(expected[i]) != 0 {
			t.Errorf("Expected %s but got %s at index %d", expected[i], v, i)
		}
	}
}

// bigIntSlice is a type for sorting []*big.Int.
type bigIntSlice []*big.Int

func (p bigIntSlice) Len() int           { return len(p) }
func (p bigIntSlice) Less(i, j int) bool { return p[i].Cmp(p[j]) < 0 }
func (p bigIntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Test for the MapToBigIntSlice function with generics.
func TestMapToBigIntSlice(t *testing.T) {
	m := map[uint64]struct{}{
		1: {},
		2: {},
		3: {},
	}

	expected := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
	}

	result := submitter.MapToBigIntSlice(m)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d elements, but got %d", len(expected), len(result))
	}

	sort.Sort(bigIntSlice(result))
	sort.Sort(bigIntSlice(expected))

	for i, v := range result {
		if v.Cmp(expected[i]) != 0 {
			t.Errorf("Expected %s but got %s at index %d", expected[i], v, i)
		}
	}
}

func TestMapToBigIntSliceWithStruct(t *testing.T) {
	type MyStruct struct {
		Value int
	}
	m := map[uint64]MyStruct{
		1: {Value: 10},
		2: {Value: 20},
		3: {Value: 30},
	}

	expected := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
	}

	result := submitter.MapToBigIntSlice(m)

	if len(result) != len(expected) {
		t.Fatalf("Expected %d elements, but got %d", len(expected), len(result))
	}

	sort.Sort(bigIntSlice(result))
	sort.Sort(bigIntSlice(expected))

	for i, v := range result {
		if v.Cmp(expected[i]) != 0 {
			t.Errorf("Expected %s but got %s at index %d", expected[i], v, i)
		}
	}
}
