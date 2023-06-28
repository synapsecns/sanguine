package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

// FromPrivateKey creates a new wallet from a private key.
func FromPrivateKey(privKey *ecdsa.PrivateKey) (Wallet, error) {
	var account walletImpl

	account.privateKey = privKey
	pubKey := privKey.Public()
	pubKeyEcdsa, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return account, fmt.Errorf("unable to cast public key of type %T to *ecdsa.PublicKey", pubKey)
	}

	accountAddr := crypto.PubkeyToAddress(*pubKeyEcdsa)

	return &walletImpl{
		address:    accountAddr,
		privateKey: privKey,
		publicKey:  pubKeyEcdsa,
	}, nil
}

// FromKeyFile creates a wallet from a file. The wallet detects if this is a
// seed phrase or a private key hex and returns the appropriate wallet. An error is
// returned if this cannot be determined.
//
// Currently, the default derivation path is used. This should be compatible with metamask
// and most wallets. It gets the first wallet in this derivation path.
// TODO: support json files.
func FromKeyFile(keyFile string) (Wallet, error) {
	//nolint:gosec
	rawKey, err := os.ReadFile(core.ExpandOrReturnPath(keyFile))
	if err != nil {
		return nil, fmt.Errorf("could not get seed phrase: %w", err)
	}
	key := strings.TrimSpace(string(rawKey))

	// if it's a mnemonic use that
	if bip39.IsMnemonicValid(key) {
		return FromSeedPhrase(key, accounts.DefaultBaseDerivationPath)
	}

	return FromHex(key)
}

// FromHex gets the wallet from the private key.
func FromHex(privateKey string) (Wallet, error) {
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode key: %w", err)
	}
	return FromPrivateKey(privKey)
}

// FromSeedPhrase gets the seed phrase for the wallet.
// Note: there seems to be some issue here w/ longer seed phraeses should investigate.
func FromSeedPhrase(seedPhrase string, derivationPath accounts.DerivationPath) (Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(seedPhrase)
	if err != nil {
		return nil, fmt.Errorf("could not get parse phrase: %w", err)
	}

	wallet.SetFixIssue172(true)

	account, err := wallet.Derive(derivationPath, true)
	if err != nil {
		return nil, fmt.Errorf("could not derive account: %w", err)
	}

	privKey, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, fmt.Errorf("could not get private key: %w", err)
	}
	return FromPrivateKey(privKey)
}

// FromRandom generates a new private key. Note: this should be used for testing only and has not been audited yet.
func FromRandom() (Wallet, error) {
	newSeed, err := hdwallet.NewSeed()
	if err != nil {
		return nil, fmt.Errorf("could not generate seed: %w", err)
	}

	newWallet, err := hdwallet.NewFromSeed(newSeed)
	if err != nil {
		return nil, fmt.Errorf("could not use seed: %w", err)
	}

	account, err := newWallet.Derive(accounts.DefaultBaseDerivationPath, true)
	if err != nil {
		return nil, fmt.Errorf("could not use derive account at %s: %w", accounts.DefaultBaseDerivationPath, err)
	}

	privKey, err := newWallet.PrivateKey(account)
	if err != nil {
		return nil, fmt.Errorf("could not get private key: %w", err)
	}

	return FromPrivateKey(privKey)
}
