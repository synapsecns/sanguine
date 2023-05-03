package testsuite_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/mocktesting"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
	"time"
)

func TestMustMarshall(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	p := Person{Name: "John Doe", Age: 30}

	// Test that MustMarshall successfully marshalls the struct to JSON
	result := testsuite.MustMarshall(t, p)
	expectedResult := `{"name":"John Doe","age":30}`
	if string(result) != expectedResult {
		t.Errorf("Expected %q but got %q", expectedResult, result)
	}

	// Test that MustMarshall fails when given an unmarshallable input
	invalidInput := make(chan int)
	mockTester := mocktesting.NewMockTester("mock tester")
	mockTester.SetOutputHandler(func(i ...interface{}) {
		// do nothing
	})
	testsuite.MustMarshall(mockTester, invalidInput)
	True(t, mockTester.Failed())
}

func TestEventually(t *testing.T) {
	// Test that Eventually returns without error when the condition is true before the context times out
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	count := 0
	testFunc := func() bool {
		count++
		return count >= 5
	}
	testsuite.Eventually(ctx, t, testFunc)
	if count != 5 {
		t.Errorf("Expected testFunc to be called 5 times but got %d", count)
	}

	// Test that Eventually returns an error when the context times out before the condition is true
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	testFunc = func() bool {
		return false
	}

	mockTester := mocktesting.NewMockTester("mock tester")
	testsuite.Eventually(ctx, mockTester, testFunc)
	True(t, mockTester.Failed())
}

// foo is a function that is used to test GetFunctionName.
func foo() {

}

// bar is a struct that is used to test GetFunctionName.
type bar struct {
}

// foo is a method that is used to test GetFunctionName.
func (b *bar) foo() {}

func TestGetFunctionName(t *testing.T) {
	// Test that GetFunctionName returns the name of a function
	functionName := testsuite.GetFunctionName(foo)
	if functionName != "foo" {
		t.Errorf("Expected GetFunctionName to return %q but got %q", "foo", functionName)
	}

	// Test that GetFunctionName returns the name of a method
	b := bar{}
	functionName = testsuite.GetFunctionName(b.foo)
	if functionName != "foo" {
		t.Errorf("Expected GetFunctionName to return %q but got %q", "foo", functionName)
	}
}
