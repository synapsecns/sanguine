package core

import (
	"os"
	"strconv"
)

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
