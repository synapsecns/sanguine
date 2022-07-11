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
