package base

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

// InMemoryKeyStore stores accounts.
type InMemoryKeyStore struct {
	keys []*keystore.Key
	mux  sync.RWMutex
}

// NewInMemoryKeyStore exposes a key store.
func NewInMemoryKeyStore() *InMemoryKeyStore {
	return &InMemoryKeyStore{
		keys: nil,
		mux:  sync.RWMutex{},
	}
}

// GetAccount gets an account from the in memory store. If no account is found, returns nil.
func (i *InMemoryKeyStore) GetAccount(address common.Address) *keystore.Key {
	i.mux.RLock()
	defer i.mux.RUnlock()
	for _, potentialKey := range i.keys {
		if address == potentialKey.Address {
			return potentialKey
		}
	}
	return nil
}

// Store stores an ccount.
func (i *InMemoryKeyStore) Store(key *keystore.Key) {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.keys = append(i.keys, key)
}
