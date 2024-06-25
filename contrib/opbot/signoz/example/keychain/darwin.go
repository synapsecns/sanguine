//go:build darwin
// +build darwin

package keychain

import keychainHelper "github.com/keybase/go-keychain"

// NewGenericPassword creates a new generic password item.
func NewGenericPassword(service string, account string, label string, data []byte, accessGroup string) interface{} {
	return keychainHelper.NewGenericPassword(service, account, label, data, accessGroup)
}

// AddItem adds an item to the keychain.
func AddItem(item interface{}) error {
	return keychainHelper.AddItem(item.(keychainHelper.Item))
}

// GetGenericPassword retrieves a generic password from the keychain.
func GetGenericPassword(service string, account string, label string, accessGroup string) ([]byte, error) {
	return keychainHelper.GetGenericPassword(service, account, label, accessGroup)
}
