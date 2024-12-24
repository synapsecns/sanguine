//go:build !windows

package permissions

import (
	"fmt"
	"os"
	"syscall"
)

// setUnixPermissions sets secure file permissions on Unix systems.
func setUnixPermissions(path string, perm os.FileMode) error {
	oldUmask := syscall.Umask(DefaultUmask)
	defer syscall.Umask(oldUmask)
	if err := os.Chmod(path, perm); err != nil {
		return fmt.Errorf("failed to set file permissions on Unix: %w", err)
	}
	return nil
}

// setUnixUmask sets a secure umask on Unix systems.
func setUnixUmask() (func(), error) {
	oldUmask := syscall.Umask(DefaultUmask)
	return func() {
		syscall.Umask(oldUmask)
	}, nil
}
