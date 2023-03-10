package mocktesting_test

import (
	"github.com/synapsecns/sanguine/core/mocktesting"
	"testing"
)

func TestMockTester(t *testing.T) {
	mockTester := mocktesting.NewMockTester("TestMockTester")

	// Test Log and Logf
	mockTester.Log("Test Log")
	mockTester.Logf("Test Logf with argument %d", 42)

	// Test Error and Errorf
	mockTester.Error("Test Error")
	mockTester.Errorf("Test Errorf with argument %s", "hello")

	// Test Fail and FailNow
	mockTester.Fail()
	if !mockTester.Failed() {
		t.Errorf("Expected MockTester to have failed, but it didn't")
	}
	mockTester.FailNow()

	// Test Failed
	mockTester2 := mocktesting.NewMockTester("TestFailed")
	if mockTester2.Failed() {
		t.Errorf("Expected MockTester2 to not have failed, but it did")
	}

	// Test Skip and Skipf
	mockTester.Skip("Test Skip")
	mockTester.Skipf("Test Skipf with argument %d", 42)

	// Test Skipped
	mockTester3 := mocktesting.NewMockTester("TestSkipped")
	if mockTester3.Skipped() {
		t.Errorf("Expected MockTester3 to not have been skipped, but it did")
	}
	mockTester3.SkipNow()
	if !mockTester3.Skipped() {
		t.Errorf("Expected MockTester3 to have been skipped, but it wasn't")
	}

	// Test Name
	if mockTester.Name() != "TestMockTester" {
		t.Errorf("Expected MockTester name to be %q, but got %q", "TestMockTester", mockTester.Name())
	}

	// Test Helper
	mockTester.Helper()

	// Test output
	expectedOutput := []string{
		"Test Log\n",
		"Test Logf with argument 42\n",
		"Test Error\n",
		"Test Errorf with argument hello\n",
		"Test Skip\n",
		"Test Skipf with argument 42\n",
	}
	for i, expectedLine := range expectedOutput {
		if mockTester.Output()[i] != expectedLine {
			t.Errorf("Expected output line %d to be %q but got %q", i, expectedLine, mockTester.Output())
		}
	}
}
