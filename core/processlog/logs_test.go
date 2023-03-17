package processlog_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/processlog"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// TestWritePrintFunc tests the writePrintFunc function.
func TestWritePrintFunc(t *testing.T) {
	// Create a new pipe for stdout and write some data to it.
	stdOutR, stdOutW := io.Pipe()
	go func() {
		defer func(stdOutW *io.PipeWriter) {
			err := stdOutW.Close()
			assert.Nil(t, err)
		}(stdOutW)
		_, err := stdOutW.Write([]byte("stdout data"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Create a new pipe for stderr and write some data to it.
	stdErrR, stdErrW := io.Pipe()
	go func() {
		defer func(stdErrW *io.PipeWriter) {
			err := stdErrW.Close()
			assert.Nil(t, err)
		}(stdErrW)
		_, err := stdErrW.Write([]byte("stderr data"))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	// Create a new buffer for the print function.
	var buf bytes.Buffer

	// Create a new context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Call the writePrintFunc function.
	processlog.WritePrintFunc(ctx, func(s []byte) {
		buf.WriteString(fmt.Sprintf("%s ", s))
	}, time.Second, stdOutR, stdErrR)

	// Check that the buffer contains the expected data.
	actual := strings.Split(strings.TrimSuffix(buf.String(), " "), " ")
	expected := strings.Split("stderr datastdout data", " ")

	if len(actual) != len(expected) {
		t.Errorf("expected %q, but got %q", expected, buf.String())
	}
}

type bufCloser struct {
	*bytes.Buffer
}

// Mock.
func (b bufCloser) Close() error {
	return nil
}

// TestPipeSingleOutput tests the pipeSingleOutput function.
func TestPipeSingleOutput(t *testing.T) {
	// Create a new temporary file.
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() {
		_ = os.Remove(tmpfile.Name())
	}()

	// Create a new buffer with some data.
	data := "test data"
	buf := bufCloser{bytes.NewBufferString(data)}

	// Call the pipeSingleOutput function.
	err = processlog.PipeSingleOutput(tmpfile, buf)

	// Check that the function returned no error.
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the file contains the expected data.
	fileData, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if string(fileData) != data {
		t.Errorf("expected %q, but got %q", data, string(fileData))
	}
}

// TestPipeCombinedOutput tests the pipeCombinedOutput function.
func TestPipeCombinedOutput(t *testing.T) {
	// Create a new temporary file.
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			return
		}
	}(tmpfile.Name())

	// Create some mock data.
	data1 := "test data 1"
	data2 := "test data 2"
	data3 := "test data 3"
	buf1 := bufCloser{bytes.NewBufferString(data1)}
	buf2 := bufCloser{bytes.NewBufferString(data2)}
	buf3 := bufCloser{bytes.NewBufferString(data3)}

	// Call the pipeCombinedOutput function with the mock data.
	err = processlog.PipeCombinedOutput(context.Background(), tmpfile, nil, buf1, buf2, buf3)

	// Check that the function returned no error.
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the file contains the expected data.
	fileData, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expectedData := data1 + "\n" + data2 + "\n" + data3 + "\n"
	if len(strings.Split(expectedData, "\n")) != len(strings.Split(string(fileData), "\n")) {
		t.Errorf("expected %q, but got %q", expectedData, string(fileData))
	}
}

// TestStartLogs tests the StartLogs function.
func TestStartLogs(t *testing.T) {
	// Create a new temporary directory.
	tmpdir, err := os.MkdirTemp("", "example")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tmpdir)

	// Create some mock data.
	data1 := "test data 1"
	data2 := "test data 2"
	data3 := "test data 3"
	buf1 := bytes.NewBufferString(data1)
	buf2 := bytes.NewBufferString(data2)
	acc := bytes.NewBufferString(data3)

	// Set up the arguments for the StartLogs function.
	logDir := filet.TmpDir(t, tmpdir)
	logFileName := "test"
	logFrequency := 1 * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	args := []processlog.StdStreamLogArgsOption{
		processlog.WithStdOut(bufCloser{buf1}),
		processlog.WithStdErr(bufCloser{buf2}),
		processlog.WithLogDir(logDir),
		processlog.WithLogFileName(logFileName),
		processlog.WithLogFrequency(logFrequency),
		processlog.WithAccumulator(acc),
		processlog.WithCtx(ctx),
	}

	// Call the StartLogs function.
	res, err := processlog.StartLogs(args...)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Wait for the log files to be written.
	time.Sleep(2 * time.Second)

	// Check that the separate stdout file contains the expected data.
	stdOutData, err := os.ReadFile(res.StdOutFile())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if string(stdOutData) != data1 {
		t.Errorf("expected %q, but got %q", data1+"\n", string(stdOutData))
	}

	// Check that the separate stderr file contains the expected data.
	stdErrData, err := os.ReadFile(res.StdErrFile())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if string(stdErrData) != data2 {
		t.Errorf("expected %q, but got %q", data2+"\n", string(stdErrData))
	}

	// Check that the combined log file contains the expected data.
	combinedData, err := os.ReadFile(res.CombinedFile())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !(len(string(combinedData)) >= len(stdErrData)+len(stdOutData)) {
		t.Error("failed to combine")
	}
}
