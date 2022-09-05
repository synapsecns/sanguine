package mocks

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chzyer/test"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// see: https://git.io/JGsC1
// taken from geth, used to speed up tests.
const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

// MockAccount mocks a new account.
func MockAccount(tb testing.TB) *keystore.Key {
	tb.Helper()

	kstr := keystore.NewKeyStore(filet.TmpDir(tb, ""), veryLightScryptN, veryLightScryptP)
	password := gofakeit.Password(true, true, true, false, false, 10)
	acct, err := kstr.NewAccount(password)
	assert.Nil(tb, err)

	data, err := os.ReadFile(acct.URL.Path)
	assert.Nil(tb, err)

	key, err := keystore.DecryptKey(data, password)
	assert.Nil(tb, err)

	return key
}

// MockAddress generates a random ethereum address for testing.
func MockAddress() common.Address {
	return common.BytesToAddress(test.RandBytes(common.HashLength))
}
