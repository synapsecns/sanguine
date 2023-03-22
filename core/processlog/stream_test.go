package processlog_test

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"github.com/synapsecns/sanguine/core/processlog"
	"io"
	"strings"
	"testing"
	"time"
)

func TestReadln(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Read normal single line",
			input:       "Hello, World!\n",
			expected:    "Hello, World!",
			expectError: false,
		},
		{
			name:        "Read empty line",
			input:       "\n",
			expected:    "",
			expectError: false,
		},
		{
			name:        "Read line without newline",
			input:       "Hello, World!",
			expected:    "Hello, World!",
			expectError: false,
		},
		{
			name:        "Read empty input",
			input:       "",
			expected:    "",
			expectError: true,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			r := bufio.NewReader(strings.NewReader(tt.input))
			line, err := processlog.ReadLine(r)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, but got: %v", err)
				}
			}

			if string(line) != tt.expected {
				t.Fatalf("expected line %q, but got %q", tt.expected, line)
			}
		})
	}
}

type readCloser struct {
	io.Reader
}

func (rc *readCloser) Close() error {
	return nil
}

// nolint: gocognit, cyclop
func TestCombineStreams(t *testing.T) {
	tests := []struct {
		name           string
		inputs         []string
		expectedOutput []string
	}{
		{
			name: "Combine two streams",
			inputs: []string{
				"Stream1 Line1\nStream1 Line2\n",
				"Stream2 Line1\nStream2 Line2\n",
			},
			expectedOutput: []string{
				"Stream1 Line1",
				"Stream1 Line2",
				"Stream2 Line1",
				"Stream2 Line2",
			},
		},
		{
			name: "Combine empty streams",
			inputs: []string{
				"",
				"",
			},
			expectedOutput: []string{},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			inputStreams := make([]io.ReadCloser, len(tt.inputs))
			for i, input := range tt.inputs {
				inputStreams[i] = &readCloser{strings.NewReader(input)}
			}

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			output, errChan := processlog.CombineStreams(ctx, inputStreams...)
			var results []string
			for line := range output {
				results = append(results, string(line))
			}

			for err := range errChan {
				if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) && !errors.Is(err, io.EOF) {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if len(results) != len(tt.expectedOutput) {
				t.Errorf("expected output length %d, but got %d", len(tt.expectedOutput), len(results))
			}

			for i := range tt.expectedOutput {
				expected := tt.expectedOutput[i]
				found := false
				for _, result := range results {
					if expected == result {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("Expected output %q not found in the combined results", expected)
				}
			}
		})
	}
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

func TestSplitStreams(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		splitCount int
	}{
		{
			name:       "Split input into two streams",
			input:      "Hello, World!",
			splitCount: 2,
		},
		{
			name:       "Split empty input",
			input:      "",
			splitCount: 3,
		},
	}

	for i := range tests {
		i := i
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			inputReader := nopCloser{bytes.NewReader([]byte(tt.input))}
			outputReaders := processlog.SplitStreams(inputReader, tt.splitCount)

			if len(outputReaders) != tt.splitCount {
				t.Errorf("expected %d output readers, but got %d", tt.splitCount, len(outputReaders))
			}

			for it, r := range outputReaders {
				data, err := io.ReadAll(r)
				if err != nil {
					t.Errorf("error reading from output reader %d: %v", it, err)
				}

				if string(data) != tt.input {
					t.Errorf("expected output reader %d to have data %q, but got %q", i, tt.input, string(data))
				}
			}
		})
	}
}

// nolint: dupl, gocognit, cyclop
func TestNewBufferedPipe(t *testing.T) {
	t.Run("Write and read data", func(t *testing.T) {
		pipe := processlog.NewBufferedPipe()

		testData := []byte("Hello, World!")

		go func() {
			n, err := pipe.Write(testData)
			if err != nil {
				t.Errorf("unexpected error while writing: %v", err)
				return
			}
			if n != len(testData) {
				t.Errorf("expected to write %d bytes, but wrote %d", len(testData), n)
				return
			}
			err = pipe.CloseWriter()
			if err != nil {
				t.Errorf("unexpected error while closing writer: %v", err)
				return
			}
		}()

		readData := make([]byte, len(testData))
		n, err := io.ReadFull(pipe, readData)
		if err != nil {
			t.Fatalf("unexpected error while reading: %v", err)
		}
		if n != len(testData) {
			t.Fatalf("expected to read %d bytes, but read %d", len(testData), n)
		}

		if !bytes.Equal(testData, readData) {
			t.Errorf("expected read data %q, but got %q", testData, readData)
		}

		err = pipe.CloseReader()
		if err != nil {
			t.Fatalf("unexpected error while closing reader: %v", err)
		}
	})

	t.Run("Close writer before reading", func(t *testing.T) {
		pipe := processlog.NewBufferedPipe()

		testData := []byte("Hello, World!")

		go func() {
			n, err := pipe.Write(testData)
			if err != nil {
				t.Errorf("unexpected error while writing: %v", err)
				return
			}
			if n != len(testData) {
				t.Errorf("expected to write %d bytes, but wrote %d", len(testData), n)
				return
			}
			err = pipe.CloseWriter()
			if err != nil {
				t.Errorf("unexpected error while closing: %v", err)
				return
			}
		}()

		readData := make([]byte, len(testData))
		n, err := io.ReadFull(pipe, readData)
		if err != nil {
			t.Fatalf("unexpected error while reading: %v", err)
		}
		if n != len(testData) {
			t.Fatalf("expected to read %d bytes, but read %d", len(testData), n)
		}

		if !bytes.Equal(testData, readData) {
			t.Errorf("expected read data %q, but got %q", testData, readData)
		}

		err = pipe.CloseReader()
		if err != nil {
			t.Fatalf("unexpected error while closing reader: %v", err)
		}
	})
}
