// Package testsuite contains the standard test suite
package testsuite

import (
	"context"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/stretchr/testify/suite"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// cancellableContext is a context object
// note: we don't really have a choice here, but note this is vulnerable to context leaks if setup/teardown
// events are not called.
type cancellableContext struct {
	// ctx is the context object
	//nolint: containedctx
	ctx context.Context
	// cancelFunc allows the context to be canceled
	cancelFunc context.CancelFunc
}

// newCancellableContext creates a new cancellable context.
func newCancellableContext(parentContext context.Context) cancellableContext {
	ctx, cancel := context.WithCancel(parentContext)
	return cancellableContext{
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

// TestSuite defines the basic test suite.
// TODO: we should make sure global vars don't get mutated. Namely eth params.
type TestSuite struct {
	suite.Suite
	// logDir is the directory where logs will be written for the docker containers that host the anvil nodes
	// this allows you to do tai -f /path/to/logs/*.combined.log to see all logs
	LogDir string
	// context is the context object for the test suite. All other context objects inherit
	// from it. unlike suiteContext and testContext it is not canceled and functionally does
	// not have a lifecycle. TODO this should probably replace suiteContext
	//nolint: containedctx
	context context.Context
	// suiteContext is the context object for the duration of the test. It is canceled at the end
	// of the test
	suiteContext cancellableContext
	// testContext is the context object for the test. It is canceled at the end of the test
	// inheritance of test contexts is as follows:
	//
	// context->suiteContext->testContext
	//
	testContext cancellableContext
	// testID is an autoincrementing that can be used as a unique, per test, identifier.
	testID int
	// runAfterTest are functions that should be run after the test
	runAfterTest []func()
	// runAfterSuite are functions that are run after the suite
	runAfterSuite []func()
	// structMux protects non thread safe operations
	structMux sync.Mutex
	// setupSuiteCalled is an assertion that makes sure setupSuite was called on the parent
	setupSuiteCalled atomic.Bool
	// setupTestCalled is an assertion that setupTest was called on the parent
	setupTestCalled atomic.Bool
	// beforeTestCalled is wether or not before test was called
	beforeTestCalled atomic.Bool
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *TestSuite {
	tb.Helper()
	log.SetAllLoggers(log.LevelWarn)
	ctx := context.Background()

	return &TestSuite{
		context: ctx,
		testID:  0,
	}
}

// Eventually asserts something is eventually true.
// Note: this uses the suite context instead of the test context.
// TODO: either separate suite and test context calls or deprecte and require generic.
func (s *TestSuite) Eventually(willBeTrue func() bool) {
	Eventually(s.GetSuiteContext(), s.T(), willBeTrue)
}

// SetupSuite sets up the test suite.
func (s *TestSuite) SetupSuite() {
	s.structMux.Lock()
	defer s.structMux.Unlock()

	s.runAfterSuite = nil
	s.suiteContext = newCancellableContext(s.context)
	s.setupSuiteCalled.Store(true)
}

// BeforeTest performs some assertions and sets up the test.
func (s *TestSuite) BeforeTest(_, _ string) {
	if !s.setupTestCalled.Load() {
		panic("make sure SetupTest is called on the parent")
	}

	if !s.setupSuiteCalled.Load() {
		panic("make sure SetupSuite is called on the parent")
	}

	s.beforeTestCalled.Store(true)
	s.DeferAfterTest(func() {
		if !s.beforeTestCalled.Load() {
			panic("make sure you call before test on the parent")
		}
	})
}

// runDefferedFunctions runs deferred functions.
func runDeferredFunctions(deferredFuncs []func()) {
	for _, deferredFunc := range deferredFuncs {
		deferredFunc()
	}
}

// TearDownSuite tears down the test suite.
func (s *TestSuite) TearDownSuite() {
	runDeferredFunctions(s.runAfterSuite)
	s.suiteContext.cancelFunc()

	s.setupSuiteCalled.Store(false)
}

// SetupTest runs checks at the end of the test suite.
func (s *TestSuite) SetupTest() {
	s.runAfterTest = nil
	s.testContext = newCancellableContext(s.suiteContext.ctx)
	s.setupTestCalled.Store(true)

	fmt.Printf("running test %s with id %d \n", s.T().Name(), s.testID)
}

// SetTestTimeout will create a test timout override for the context.
// this will wrap s.testContext.
// TODO: consider enabling a value by default.
func (s *TestSuite) SetTestTimeout(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(s.testContext.ctx, timeout)
	oldCancel := s.testContext.cancelFunc
	s.testContext = cancellableContext{ctx: ctx, cancelFunc: func() {
		// cancel parent
		cancel()
		// then underlying
		oldCancel()
	}}
}

// GetTestID gets the unique test id for the current test.
// uniqueness is per-suite.
func (s *TestSuite) GetTestID() int {
	return s.testID
}

// TearDownTest runs checks at the end of the test suite.
func (s *TestSuite) TearDownTest() {
	s.testID++
	runDeferredFunctions(s.runAfterTest)
	// this will panic if you failed to call SetupTest() from an inheriting suite
	s.testContext.cancelFunc()
	fmt.Printf("finished running test %s with id %d \n", s.T().Name(), s.testID)

	s.setupTestCalled.Store(false)
}

// DeferAfterTest runs a function after the test. This will run before context cancellation
// if you'd like to do otherwise you can from a new goroutine that watches TestContext()
// TODO: in cases of crashes this will not be done so fix this.
func (s *TestSuite) DeferAfterTest(newFunc func()) {
	s.structMux.Lock()
	defer s.structMux.Unlock()
	s.runAfterTest = append(s.runAfterTest, newFunc)
}

// DeferAfterSuite runs a function after the suite. This will run before context cancellation
// if you'd like to do otherwise you can from a new goroutine that watches SuiteContext()
// TODO: in cases of crashes this will not be done so fix this.
func (s *TestSuite) DeferAfterSuite(newFunc func()) {
	s.structMux.Lock()
	defer s.structMux.Unlock()
	s.runAfterSuite = append(s.runAfterSuite, newFunc)
}

// GetTestContext gets the context object for the suite. This is an alias for GetTestContext()
// TODO: right now this is run as a test because of naming. this is mostly harmless.
func (s *TestSuite) GetTestContext() context.Context {
	return s.testContext.ctx
}

// GetSuiteContext returns the context for the test suite.
func (s *TestSuite) GetSuiteContext() context.Context {
	return s.suiteContext.ctx
}

// MustMarshall is a helper method that attempts to marshall, otherwise it
// fails the test.
func (s *TestSuite) MustMarshall(v any) []byte {
	return MustMarshall(s.T(), v)
}
