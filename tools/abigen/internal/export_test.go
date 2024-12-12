package internal

import (
	"github.com/ethereum/go-ethereum/common/compiler"
)

// CheckForDocker exports checkForDocker for testing.
func CheckForDocker() error {
	return checkForDocker()
}

// CompileSolidity exports compileSolidity for testing.
func CompileSolidity(version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	return compileSolidity(version, filePath, optimizeRuns, evmVersion)
}

// FilePathsAreEqual exports filePathsAreEqual for testing.
func FilePathsAreEqual(file1 string, file2 string) (equal bool, err error) {
	return filePathsAreEqual(file1, file2)
}
