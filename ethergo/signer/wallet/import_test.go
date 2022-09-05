package wallet_test

import (
	"encoding/hex"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/tyler-smith/go-bip39"
	"os"
	"path"
	"strings"
)

// ExampleFromSeedPhrase shows an example creating a wallet from a seed phrase.
func ExampleFromSeedPhrase() {
	const mnemonic = "tag volcano eight thank tide danger coast health above argue embrace heavy"
	testWallet, _ := wallet.FromSeedPhrase(mnemonic, accounts.DefaultBaseDerivationPath)

	fmt.Printf("Address: %s\n", testWallet.Address())
	fmt.Printf("Public Key (From Bytes): %s\n", hex.EncodeToString(testWallet.PrivateKeyBytes()))
	fmt.Printf("Public Key Hex: %s\n", testWallet.PublicKeyHex())
	fmt.Printf("Private Key (From Bytes): %s\n", testWallet.PrivateKeyHex())
	fmt.Printf("Private Key Hex: %s\n", hex.EncodeToString(testWallet.PrivateKeyBytes()))

	// output:
	// Address: 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
	// Public Key (From Bytes): 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
	// Public Key Hex: 6005c86a6718f66221713a77073c41291cc3abbfcd03aa4955e9b2b50dbf7f9b6672dad0d46ade61e382f79888a73ea7899d9419becf1d6c9ec2087c1188fa18
	// Private Key (From Bytes): 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
	// Private Key Hex: 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9
}

func (s WalletSuite) TestFromKeyFile() {
	for _, testCase := range s.makeKeyFileTests() {
		testFile := s.makeTestFile(testCase.contents, !testCase.shouldErr)
		testWallet, err := wallet.FromKeyFile(testFile)
		if testCase.shouldErr {
			NotNil(s.T(), err)
			continue
		}
		Nil(s.T(), err)

		Equal(s.T(), testWallet.Address(), testCase.address)
	}
}

func (s WalletSuite) TestFromRandom() {
	newWallet, err := wallet.FromRandom()
	Nil(s.T(), err)

	NotPanics(s.T(), func() {
		newWallet.Address()
		newWallet.PrivateKey()
		newWallet.PublicKey()
		newWallet.PrivateKeyBytes()
		newWallet.PublicKeyHex()
		newWallet.PublicKeyBytes()
	})
}

func (s WalletSuite) makeTestFile(contents []byte, fuzz bool) (fileName string) {
	if fuzz {
		contents = fuzzContents(contents)
	}

	testFile, err := os.Create(path.Join(filet.TmpDir(s.T(), ""), fmt.Sprintf("%s.wallet", gofakeit.UUID())))
	Nil(s.T(), err)

	defer func(testFile *os.File) {
		Nil(s.T(), testFile.Close())
	}(testFile)

	_, err = testFile.Write(contents)
	Nil(s.T(), err)

	return testFile.Name()
}

// makeFuzzFile randomly fuzzes a byte slice (representing a file) in a way that should not break
// (e.g. adding a new line at the end, spaces) in a way that does not change the contents.
func fuzzContents(contents []byte) (out []byte) {
	fuzzSetting := gofakeit.Number(0, 3)
	switch fuzzSetting {
	case 0:
		contents = append(contents, []byte(" ")...)
	case 1:
		contents = append(contents, []byte("\n\r")...)
	case 2:
		contents = append(contents, []byte("\f\r")...)
	}
	return contents
}

type KeyFileTest struct {
	// address is the expected output of the address
	address common.Address
	// shouldErr determines whether an error should be expected when running the test
	shouldErr bool
	// contents of the key file
	contents []byte
	// description of what's being tested
	description string
}

func (s WalletSuite) makeKeyFileTests() (tests []KeyFileTest) {
	// test an hdwallet, from https://pkg.go.dev/github.com/miguelmota/go-ethereum-hdwallet#section-readme
	tests = append(tests, KeyFileTest{
		address:     common.HexToAddress("0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947"),
		shouldErr:   false,
		contents:    []byte("tag volcano eight thank tide danger coast health above argue embrace heavy"),
		description: "test the default derivation path",
	})

	tests = append(tests, KeyFileTest{
		address:     common.HexToAddress("0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947"),
		shouldErr:   false,
		contents:    []byte("63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9"),
		description: "test the private key as a hex",
	})

	// fuzz
	tests = append(tests, KeyFileTest{
		address:     mocks.MockAddress(),
		shouldErr:   true,
		contents:    nil,
		description: "no file contents",
	})

	tests = append(tests, KeyFileTest{
		address:     mocks.MockAddress(),
		shouldErr:   true,
		contents:    []byte(strings.Join(bip39.GetWordList(), " ")),
		description: "long word list",
	})

	return tests
}
