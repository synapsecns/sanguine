package testsuite

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/wait"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

// Here, we separate out some functions that may be useful even outside of the test suite context
// for example, in utility functions.

// MustMarshall is a helper method that attempts to marshall, otherwise it
// fails the test.
func MustMarshall(tb testing.TB, v any) []byte {
	tb.Helper()

	res, err := json.Marshal(v)
	assert.Nil(tb, err)
	return res
}

// Eventually asserts willBeTrue is true before the test context times out.
func Eventually(ctx context.Context, tb testing.TB, willBeTrue func() bool) {
	tb.Helper()

	ctx, cancel := context.WithCancel(ctx)
	isTrue := false
	wait.UntilWithContext(ctx, func(ctx context.Context) {
		if willBeTrue() {
			isTrue = true
			cancel()
		}
	}, time.Millisecond)

	// make sure the context just didn't cancel
	if !isTrue {
		tb.Errorf("expected %T to be true before test context timed out", willBeTrue)
	}
}

// GetFunctionName returns the name of the function passed in.
// this is useful for getting the name of the test function that is running.
func GetFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	functionNameWithSuffix := parts[len(parts)-1]
	functionName := strings.TrimSuffix(functionNameWithSuffix, "-fm")

	return functionName
}
