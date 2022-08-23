package ganache

import (
	"encoding/json"
	"fmt"
	"os"
)

// Address is an object exported by --account_keys_path in ganache.
// these are used to authenticate with a test chain.
type Address struct {
	Secretkey struct {
		Type string `json:"type"`
		Data []int  `json:"data"`
	} `json:"secretKey"`
	Publickey struct {
		Type string `json:"type"`
		Data []int  `json:"data"`
	} `json:"publicKey"`
	Address string `json:"address"`
	Account struct {
		Nonce     string `json:"nonce"`
		Balance   string `json:"balance"`
		Stateroot string `json:"stateRoot"`
		Codehash  string `json:"codeHash"`
	} `json:"account"`
}

// Addresses is a list of ganache addresses.
type Addresses struct {
	Addresses   map[string]Address `json:"addresses"`
	PrivateKeys map[string]string  `json:"private_keys"`
}

// ParseAddresses parses the addresses out of a path and returns an object.
func ParseAddresses(path string) (addresses *Addresses, err error) {
	//nolint: gosec
	keyFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %s (got error: %w)", path, err)
	}

	err = json.Unmarshal(keyFile, &addresses)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal file %s (got error: %w)", keyFile, err)
	}
	return addresses, nil
}
