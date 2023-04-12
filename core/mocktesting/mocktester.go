package mocktesting

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
)

// MockTester is a mock tester. This is used for making sure a test fails
// using the testing.TB interface and is useful for testing error cases in other test helpers.
type MockTester struct {
	name            string
	mu              sync.RWMutex
	output          []string
	skipped, failed bool
	// print errors is whether or not to print errors
	printErrors bool
	// whether or not to print logs
	printLogs bool
	// outputHandler specifies how to tread output
	outputHandler func(...interface{})
	// The testing.TB interface has an unexported method "to prevent users implementing the
	// interface and so future additions to it will not violate Go 1 compatibility."
	//
	// This may cause problems across Go versions, but let's ignore them and
	// work around that restriction by embedding a T so we satisfy the unexported methods of the interface.
	*testing.T
}

// NewMockTester creates a new process tester.
func NewMockTester(name string) *MockTester {
	return &MockTester{
		name:          name,
		mu:            sync.RWMutex{},
		output:        nil,
		skipped:       false,
		printErrors:   true,
		printLogs:     true,
		outputHandler: log.New(os.Stderr, "", 0).Println,
		failed:        false,
		T:             nil,
	}
}

// SetOutputHandler sets the output handler for the test output.
func (t *MockTester) SetOutputHandler(handler func(...interface{})) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.outputHandler = handler
}

var _ testing.TB = (*MockTester)(nil)

// Error registers an error.
func (t *MockTester) Error(args ...interface{}) {
	t.Log(args...)
	t.Fail()
}

// Errorf formats the error.
func (t *MockTester) Errorf(format string, args ...interface{}) {
	t.Logf(format, args...)
	t.Fail()
}

// Fail fails the process.
func (t *MockTester) Fail() {
	t.mu.Lock()
	t.failed = true
	t.mu.Unlock()
}

// FailNow immediately fails the process.
func (t *MockTester) FailNow() {
	t.Fail()
}

// Failed indicates whether or not the process failed.
func (t *MockTester) Failed() bool {
	t.mu.RLock()
	failed := t.failed
	t.mu.RUnlock()
	return failed
}

// Fatal fails instantly.
func (t *MockTester) Fatal(args ...interface{}) {
	t.Log(args...)
	t.FailNow()
}

// Fatalf formats a fail.
func (t *MockTester) Fatalf(format string, args ...interface{}) {
	t.Logf(format, args...)
	t.FailNow()
}

// Log logs the process.
func (t *MockTester) Log(args ...interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	formattedOut := fmt.Sprintln(args...)
	t.output = append(t.output, formattedOut)
	if t.printErrors {
		t.outputHandler(fmt.Sprintln(formattedOut))
	}
}

// Logf format logs.
func (t *MockTester) Logf(format string, args ...interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	// Ensure message ends with newline.
	if len(format) > 0 && format[len(format)-1] != '\n' {
		format += "\n"
	}
	formattedOut := fmt.Sprintf(format, args...)
	t.output = append(t.output, formattedOut)
	if t.printLogs {
		t.outputHandler(fmt.Sprintln(formattedOut))
	}
}

// Name gets the test name.
func (t *MockTester) Name() string {
	return t.name
}

// Skip skips the test.
func (t *MockTester) Skip(args ...interface{}) {
	t.Log(args...)
	t.SkipNow()
}

// SkipNow skips the test immediately.
func (t *MockTester) SkipNow() {
	t.mu.Lock()
	t.skipped = true
	t.mu.Unlock()
}

// Skipf skips the test with formaatting.
func (t *MockTester) Skipf(format string, args ...interface{}) {
	t.Logf(format, args...)
	t.SkipNow()
}

// Skipped gets whether or not a process was skipped.
func (t *MockTester) Skipped() bool {
	t.mu.Lock()
	skipped := t.skipped
	t.mu.Unlock()
	return skipped
}

// Helper is used to implement the test helper.
func (t *MockTester) Helper() {}

// Output gets the output of the test.
func (t *MockTester) Output() []string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.output
}
