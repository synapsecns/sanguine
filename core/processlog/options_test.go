package processlog_test

import (
	"context"
	"github.com/synapsecns/sanguine/core/processlog"
	"io"
	"testing"
)

// TestMakeArgs tests the makeArgs function.
func TestMakeArgs(t *testing.T) {
	pr, _ := io.Pipe()

	tests := []struct {
		name      string
		opts      []processlog.StdStreamLogArgsOption
		expectErr bool
	}{
		{
			name: "valid options",
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithStdOut(pr),
				processlog.WithStdErr(pr),
				processlog.WithCtx(context.Background()),
			},
			expectErr: false,
		},
		{
			name: "missing Ctx",
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithStdOut(pr),
				processlog.WithStdErr(pr),
			},
			expectErr: true,
		},
		{
			name: "missing StdOut",
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithCtx(context.Background()),
				processlog.WithStdErr(pr),
			},
			expectErr: true,
		},
		{
			name: "missing StdOut (reader)",
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithCtx(context.Background()),
				processlog.WithReader(processlog.StdErr, pr),
			},
			expectErr: true,
		},
		{
			name: "missing StdErr",
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithCtx(context.Background()),
				processlog.WithStdOut(pr),
			},
			expectErr: true,
		},
		{
			name: "missing StdErr (reader)",
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithCtx(context.Background()),
				processlog.WithReader(processlog.StdOut, pr),
			},
			expectErr: true,
		},
		{
			name: "log directory must be set",
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithCtx(context.Background()),
				processlog.WithStdOut(pr),
				processlog.WithLogDir(""),
			},
			expectErr: true,
		},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T) {
			err := processlog.MakeArgs(tc.opts)
			if tc.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("expected nil error but got %v", err)
			}
		})
	}
}

func TestHasReader(t *testing.T) {
	setReader, _ := io.Pipe()

	// Set up test data.
	testCases := []struct {
		readerType processlog.ReaderType
		opts       []processlog.StdStreamLogArgsOption
		expected   bool
	}{
		{
			readerType: processlog.StdOut,
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithStdOut(nil),
				processlog.WithStdErr(nil),
			},
			expected: false,
		},
		{
			readerType: processlog.StdOut,
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithStdOut(setReader),
				processlog.WithStdErr(nil),
			},
			expected: true,
		},
		{
			readerType: processlog.StdErr,
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithStdOut(nil),
				processlog.WithStdErr(setReader),
			},
			expected: true,
		},
		{
			readerType: processlog.StdErr,
			opts: []processlog.StdStreamLogArgsOption{
				processlog.WithStdOut(nil),
				processlog.WithStdErr(nil),
			},
			expected: false,
		},
	}

	// Run tests.
	for _, tc := range testCases {
		actual := processlog.HasReader(tc.readerType, tc.opts...)
		if actual != tc.expected {
			t.Errorf("HasReader(%v, %v) = %v, expected %v", tc.readerType, tc.opts, actual, tc.expected)
		}
	}
}
