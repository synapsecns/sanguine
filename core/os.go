package core

import (
	"github.com/mitchellh/go-homedir"
	"os"
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
	return expanded
}

// GetEnv gets an environment variable. If it is not present, the default variable is used.
func GetEnv(name, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
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
