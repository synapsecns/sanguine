package dbcommon_test

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
	"testing"
)

func TestMysqlRealEscapeString(t *testing.T) {
	// ' to \\'
	// " to \\"
	// \r to \\r
	// \n to \\n
	// \\ to \\\\
	// \x00 (null byte) to \\x00
	// \x1a (substitute character) to \\Z
	tests := []struct {
		input  string
		expect string
	}{
		{"hello", "hello"},
		{"he'llo", "he\\'llo"},
		{"he\"llo", "he\\\"llo"},
		{"he\nllo", "he\\\nllo"},
		{"he\rllo", "he\\\rllo"},
		{"hello\\world", "hello\\\\world"},
		{"he\x00llo", "he\\\x00llo"},
		{"he\x1allo", "he\\Zllo"},
	}

	for _, tt := range tests {
		result := dbcommon.MysqlRealEscapeString(tt.input)
		if result != tt.expect {
			t.Errorf("For input '%s', expected '%s' but got '%s'", tt.input, tt.expect, result)
		}
	}
}
