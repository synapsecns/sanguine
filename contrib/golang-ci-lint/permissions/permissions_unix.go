//go:build !windows

package permissions

import (
	"os"
	"syscall"
)

// setUnixPermissions sets secure file permissions on Unix systems.
func setUnixPermissions(path string, perm os.FileMode) error {
	oldUmask := syscall.Umask(DefaultUmask)
	defer syscall.Umask(oldUmask)
	return os.Chmod(path, perm)
}

// setUnixUmask sets a secure umask on Unix systems.
func setUnixUmask() (func(), error) {
	oldUmask := syscall.Umask(DefaultUmask)
	return func() {
		syscall.Umask(oldUmask)
	}, nil
}
