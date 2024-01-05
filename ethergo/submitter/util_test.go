package submitter_test

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/util"
	"go.opentelemetry.io/otel/attribute"
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

func TestAddressPtrToString(t *testing.T) {
	// Test case 1: Address is nil
	var address *common.Address
	assert.Equal(t, submitter.AddressPtrToString(address), submitter.NullFieldAttribute)

	// Test case 2: Address is not nil
	address = core.PtrTo[common.Address](common.HexToAddress("0x1234567890123456789012345678901234567890"))
	assert.Equal(t, submitter.AddressPtrToString(address), "0x1234567890123456789012345678901234567890")
}

func TestBigPtrToString(t *testing.T) {
	// Test case: num is nil
	var num *big.Int
	expected := submitter.NullFieldAttribute
	result := submitter.BigPtrToString(num)
	if result != expected {
		t.Errorf("bigPtrToString(nil) = %q; want %q", result, expected)
	}

	// Test case: num is an integer
	num = big.NewInt(123)
	expected = "123"
	result = submitter.BigPtrToString(num)
	if result != expected {
		t.Errorf("bigPtrToString(123) = %q; want %q", result, expected)
	}
}

func (s *SubmitterSuite) TestTxToAttributesNullFields() {
	s.checkEmptyTx(types.NewTx(&types.DynamicFeeTx{}))
	s.checkEmptyTx(types.NewTx(&types.LegacyTx{}))
}

func (s *SubmitterSuite) checkEmptyTx(rawTx *types.Transaction) {
	tx := makeAttrMap(rawTx, uuid.New().String())

	s.Require().Equal(tx[submitter.HashAttr].AsString(), rawTx.Hash().Hex())
	s.Require().Equal(tx[submitter.NonceAttr].AsInt64(), int64(0))
	s.Require().Equal(tx[submitter.GasLimitAttr].AsInt64(), int64(0))
	s.Require().Equal(tx[submitter.ToAttr].AsString(), submitter.NullFieldAttribute)
	s.Require().Equal(tx[submitter.ValueAttr].AsString(), "0")
	s.Require().Equal(tx[submitter.DataAttr].AsString(), "")

	if rawTx.Type() == types.DynamicFeeTxType {
		s.Require().Equal(tx[submitter.GasTipCapAttr].AsString(), "0")
		s.Require().Equal(tx[submitter.GasFeeCapAttr].AsString(), "0")
	}
	if rawTx.Type() == types.LegacyTxType {
		s.Require().Equal(tx[submitter.GasPriceAttr].AsString(), "0")
	}
}

func (s *SubmitterSuite) TestTxToAttributesLegacyTX() {
	mockTX := mocks.GetMockTxes(s.GetTestContext(), s.T(), 1, types.LegacyTxType)[0]
	mapAttr := makeAttrMap(mockTX, uuid.New().String())

	s.Require().Equal(mapAttr[submitter.HashAttr].AsString(), mockTX.Hash().String())
	s.Require().Equal(mapAttr[submitter.NonceAttr].AsInt64(), int64(mockTX.Nonce()))
	s.Require().Equal(mapAttr[submitter.GasLimitAttr].AsInt64(), int64(mockTX.Gas()))
	s.Require().Equal(mapAttr[submitter.ToAttr].AsString(), mockTX.To().String())
	s.Require().Equal(mapAttr[submitter.ValueAttr].AsString(), mockTX.Value().String())
	s.Require().Equal(mapAttr[submitter.DataAttr].AsString(), "")

	s.Require().Equal(mapAttr[submitter.GasPriceAttr].AsString(), mockTX.GasPrice().String())
	_, hasFeeCap := mapAttr[submitter.GasFeeCapAttr]
	_, hasTipCap := mapAttr[submitter.GasTipCapAttr]
	s.Require().False(hasFeeCap)
	s.Require().False(hasTipCap)
	s.Require().NotNil(mapAttr[submitter.FromAttr])
}

func (s *SubmitterSuite) TestTxToAttributesDynamicTX() {
	mockTX := mocks.GetMockTxes(s.GetTestContext(), s.T(), 1, types.DynamicFeeTxType)[0]
	mapAttr := makeAttrMap(mockTX, uuid.New().String())

	s.Require().Equal(mapAttr[submitter.HashAttr].AsString(), mockTX.Hash().String())
	s.Require().Equal(mapAttr[submitter.NonceAttr].AsInt64(), int64(mockTX.Nonce()))
	s.Require().Equal(mapAttr[submitter.GasLimitAttr].AsInt64(), int64(mockTX.Gas()))
	s.Require().Equal(mapAttr[submitter.ToAttr].AsString(), mockTX.To().String())
	s.Require().Equal(mapAttr[submitter.ValueAttr].AsString(), mockTX.Value().String())
	s.Require().Equal(mapAttr[submitter.DataAttr].AsString(), "")

	s.Require().Equal(mapAttr[submitter.GasFeeCapAttr].AsString(), mockTX.GasFeeCap().String())
	s.Require().Equal(mapAttr[submitter.GasTipCapAttr].AsString(), mockTX.GasTipCap().String())
	_, hasGasPrice := mapAttr[submitter.GasPriceAttr]
	s.Require().False(hasGasPrice)
	s.Require().NotNil(mapAttr[submitter.FromAttr])
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

func makeAttrMap(tx *types.Transaction, UUID string) map[string]attribute.Value {
	mapAttr := make(map[string]attribute.Value)
	attr := submitter.TxToAttributes(tx, UUID)
	for _, a := range attr {
		mapAttr[string(a.Key)] = a.Value
	}
	return mapAttr
}
