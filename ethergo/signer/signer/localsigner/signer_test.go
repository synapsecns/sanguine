package localsigner_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"testing"
)

func TestSigner(t *testing.T) {
	testWallet, err := wallet.FromHex("63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9")
	Nil(t, err)

	newSigner := localsigner.NewSigner(testWallet.PrivateKey())
	signature, err := newSigner.SignMessage(context.Background(), []byte("his"))
	Nil(t, err)
	NotNil(t, signature)
}
