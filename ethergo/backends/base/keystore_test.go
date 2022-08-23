package base_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/synapse-node/testutils/utils"
	"testing"
)

func TestStoreGetAccount(t *testing.T) {
	inMemoryKeyStore := base.NewInMemoryKeyStore()

	key, err := utils.NewMockAuthConfig(t).Key()
	Nil(t, err)

	inMemoryKeyStore.Store(key)

	retreivedKey := inMemoryKeyStore.GetAccount(key.Address)
	Equal(t, retreivedKey, key)
}
