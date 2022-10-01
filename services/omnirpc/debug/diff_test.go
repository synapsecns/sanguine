package debug_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/debug"
	"testing"
	// used for embedding the test file.
	_ "embed"
)

//go:embed error.json
var testFile []byte

func TestDebug(t *testing.T) {
	NotPanics(t, func() {
		Nil(t, debug.HashDiff(testFile))
	})
}
