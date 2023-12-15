package core

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ExpandOrReturnPath expands a homedir path, if it can't be expanded, the original dir is returned
// since on these systems homedir cannot be used anyway.
func ExpandOrReturnPath(path string) string {
	expanded, err := homedir.Expand(path)
	if err != nil {
		return path
	}
	absPath, err := filepath.Abs(expanded)
	if err != nil {
		return expanded
	}
	return absPath
}

// GetEnv gets an environment variable. If it is not present, the default variable is used.
func GetEnv(name, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
}

// GetEnvBool gets an environment variable as a bool. If not found the default value is used.
func GetEnvBool(name string, defaultVal bool) bool {
	val := os.Getenv(name)
	if val == name {
		return defaultVal
	}
	res, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}
	return res
}

// HasEnv checks if an environment variable is set.
func HasEnv(name string) bool {
	val := os.Getenv(name)
	return val != ""
}

// GetEnvInt gets an environment variable as an int. if not found or cast the
// defaultVal is returned.
func GetEnvInt(name string, defaultVal int) int {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return res
}

// IsTest returns true if the current process is a test.
func IsTest() bool {
	if len(os.Args) == 0 {
		return false
	}

	return strings.HasSuffix(os.Args[0], ".test")
}

// CopyFile copies a file from src to dest and preserves its permissions.
func CopyFile(src, dest string) error {
	// Open source file for reading
	//nolint: gosec
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func() {
		_ = srcFile.Close()
	}()

	// Obtain stat information from source
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %w", err)
	}

	// Create destination file with the source file permissions
	//nolint: gosec
	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to open destination file: %w", err)
	}
	defer func() {
		_ = destFile.Close()
	}()

	// Copy data from source to destination
	if _, err := io.Copy(destFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy data from source to destination: %w", err)
	}

	return nil
}
