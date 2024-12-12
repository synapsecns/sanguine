// Package permissions provides platform-specific file permission handling
package permissions

import (
	"fmt"
	"os"
	"runtime"
)

const (
	// DefaultUmask is the default umask value for secure file operations.
	DefaultUmask = 0077
	// FilePerms is the default permission for regular files.
	FilePerms = 0400
	// DirPerms is the default permission for directories.
	DirPerms = 0750
	// ExecPerms is the default permission for executables.
	ExecPerms = 0500
)

// SetSecureUmask sets a secure umask and returns a cleanup function.
func SetSecureUmask() (func(), error) {
	if runtime.GOOS == "windows" {
		return func() {}, nil
	}
	return setUnixUmask()
}

// SetFilePermissions sets secure file permissions in a platform-specific way.
func SetFilePermissions(path string, perm os.FileMode) error {
	if runtime.GOOS == "windows" {
		if err := os.Chmod(path, perm); err != nil {
			return fmt.Errorf("failed to set file permissions on Windows: %w", err)
		}
		return nil
	}
	return setUnixPermissions(path, perm)
}
