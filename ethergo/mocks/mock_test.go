package mocks_test

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/richardwilkes/toolbox/collection"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"testing"
)

func TestMockAccount(t *testing.T) {
	account := mocks.MockAccount(t)

	address := crypto.PubkeyToAddress(account.PrivateKey.PublicKey)
	Equal(t, address, account.Address)
}

func TestMockAddress(t *testing.T) {
	// make sure addresses are unique
	stringSet := collection.Set[string]{}

	const genCount = 100

	for i := 0; i < genCount; i++ {
		stringSet.Add(mocks.MockAddress().String())
	}

	Equal(t, genCount, len(stringSet.Values()))
}

func TestGetMockTxes(t *testing.T) {
	const testTxCount = 10
	mockTxes := mocks.GetMockTxes(context.Background(), t, testTxCount, 0)

	// make sure txes are unique
	txSet := collection.Set[string]{}
	jsonSet := collection.Set[string]{}

	for _, tx := range mockTxes {
		raw, err := tx.MarshalJSON()
		Nil(t, err)

		jsonSet.Add(string(raw))

		txSet.Add(tx.Hash().String())
	}

	Equal(t, testTxCount, len(txSet.Values()))
	Equal(t, testTxCount, len(jsonSet.Values()))
}
