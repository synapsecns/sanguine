//go:build !darwin
// +build !darwin

package keychain

// NewGenericPassword creates a new generic password item.
func NewGenericPassword(service string, account string, label string, data []byte, accessGroup string) interface{} {
	panic("unsupported platform")
}

// AddItem adds an item to the keychain.
func AddItem(item interface{}) error {
	panic("unsupported platform")
}

// GetGenericPassword retrieves a generic password from the keychain.
func GetGenericPassword(service string, account string, label string, accessGroup string) ([]byte, error) {
	panic("unsupported platform")
}
