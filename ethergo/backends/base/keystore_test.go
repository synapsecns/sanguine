package base_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"testing"
)

func TestStoreGetAccount(t *testing.T) {
	inMemoryKeyStore := base.NewInMemoryKeyStore()

	key := mocks.MockAccount(t)

	inMemoryKeyStore.Store(key)

	retreivedKey := inMemoryKeyStore.GetAccount(key.Address)
	Equal(t, retreivedKey, key)
}
