package core

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"strconv"
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
