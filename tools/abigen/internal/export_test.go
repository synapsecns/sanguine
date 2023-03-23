package internal

import (
	"github.com/ethereum/go-ethereum/common/compiler"
	"os"
)

// CreateRunFile exports create run file for testing.
func CreateRunFile(version string) (runFile *os.File, err error) {
	return createRunFile(version)
}

// CompileSolidity exports compileSolidity for testingw.
func CompileSolidity(version string, filePath string, optimizeRuns int) (map[string]*compiler.Contract, error) {
	return compileSolidity(version, filePath, optimizeRuns)
}

// FilePathsAreEqual exports filePathsAreEqual for testing.
func FilePathsAreEqual(file1 string, file2 string) (equal bool, err error) {
	return filePathsAreEqual(file1, file2)
}
