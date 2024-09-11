package metrics_test

import (
	"github.com/synapsecns/sanguine/core/metrics"
	"reflect"
	"testing"
)

func TestHeadersToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:  "basic input",
			input: "key1=value1,key2=value2",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: map[string]string{},
		},
		{
			name:  "input with extra spaces",
			input: "key1 = value1 , key2= value2 ",
			expected: map[string]string{
				"key1 ": " value1 ",
				" key2": " value2 ",
			},
		},
		{
			name:  "input with missing value",
			input: "key1=value1,key2=",
			expected: map[string]string{
				"key1": "value1",
				"key2": "",
			},
		},
		{
			name:  "input with missing key",
			input: "=value1,key2=value2",
			expected: map[string]string{
				"":     "value1",
				"key2": "value2",
			},
		},
		{
			name:  "input with multiple equal signs",
			input: "key1=value1=extra,key2=value2",
			expected: map[string]string{
				"key1": "value1=extra",
				"key2": "value2",
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			result := metrics.HeadersToMap(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("HeadersToMap(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
