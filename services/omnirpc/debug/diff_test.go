package debug_test

import (
	"github.com/synapsecns/sanguine/services/omnirpc/debug"
	"testing"
	// used for embedding the test file
	_ "embed"
)

//go:embed error.json
var testFile []byte

func TestDebug(T *testing.T) {
	_ = make([]byte, 10<<30)

	debug.HashDiff(testFile)
}
