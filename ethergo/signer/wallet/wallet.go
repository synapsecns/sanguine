package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet is an EVM-compatible wallet
// with an ECDSA private key + public key and
// standard ethereum hex address.
type Wallet interface {
	// Address gets the wallets address
	Address() common.Address
	// PublicKey returns this wallet's public key as an *ecdsa.PublicKey
	PublicKey() *ecdsa.PublicKey
	// PrivateKey returns this wallet's private key as an *ecdsa.PrivateKey
	PrivateKey() *ecdsa.PrivateKey
	// PublicKeyHex returns this wallet's public key in hex format,
	// without a leading 0x.
	PublicKeyHex() string
	// PrivateKeyHex returns this wallet's private key in hex format,
	// without a leading 0x.
	PrivateKeyHex() string
	// PublicKeyBytes is a helper function that returns the public key as a byte slice
	PublicKeyBytes() []byte
	// PrivateKeyBytes is a helper function that returns the private key as a byte slice
	PrivateKeyBytes() []byte
	// String is an alias for Address().
	String() string
}

// walletImpl is an implementation of ethereum wallet.
type walletImpl struct {
	// address of the wallet
	address common.Address
	// privateKey for the account
	privateKey *ecdsa.PrivateKey
	// publicKey for the account
	publicKey *ecdsa.PublicKey
}

func (w walletImpl) Address() common.Address {
	return w.address
}

func (w walletImpl) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w walletImpl) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w walletImpl) PublicKeyHex() string {
	return common.Bytes2Hex(w.PublicKeyBytes())
}

func (w walletImpl) PrivateKeyHex() string {
	return common.Bytes2Hex(w.PrivateKeyBytes())
}

func (w walletImpl) PublicKeyBytes() []byte {
	b := crypto.FromECDSAPub(w.publicKey)
	// remove any leading bytes
	if bytes.HasPrefix(b, []byte{04}) {
		b = bytes.TrimPrefix(b, []byte{04})
	}

	return b
}

func (w walletImpl) PrivateKeyBytes() []byte {
	return crypto.FromECDSA(w.privateKey)
}

func (w walletImpl) String() string {
	return w.Address().String()
}
