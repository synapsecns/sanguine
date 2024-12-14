//go:build windows

package permissions

import (
	"os"
)

// setUnixPermissions is a stub for Windows
func setUnixPermissions(path string, perm os.FileMode) error {
	return os.Chmod(path, perm)
}

// setUnixUmask is a stub for Windows that maintains interface compatibility
func setUnixUmask() (func(), error) {
	return func() {}, nil
}
