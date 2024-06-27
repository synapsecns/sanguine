package core_test

import (
	"github.com/synapsecns/sanguine/core"
	"reflect"
	"strings"
	"testing"
)

func TestBytesToSlice(t *testing.T) {
	tests := []struct {
		name  string
		bytes [32]byte
		want  []byte
	}{
		{
			name:  "all zeros",
			bytes: [32]byte{},
			want:  make([]byte, 32),
		},
		{
			name:  "random bytes",
			bytes: [32]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			want:  []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			if got := core.BytesToSlice(tt.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToJSONString(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		want      string
		wantErr   bool
		errString string
	}{
		{
			name:    "Valid JSON",
			input:   []byte(`{"key1":"value1","key2":2}`),
			want:    `{"key1":"value1","key2":2}`,
			wantErr: false,
		},
		{
			name:      "Invalid JSON",
			input:     []byte(`{"key1":"value1",`),
			want:      "",
			wantErr:   true,
			errString: "failed to unmarshal JSON",
		},
		{
			name:    "Empty JSON",
			input:   []byte(`{}`),
			want:    `{}`,
			wantErr: false,
		},
		{
			name:    "Nested JSON",
			input:   []byte(`{"key1":{"nestedKey1":"nestedValue1"},"key2":[1,2,3]}`),
			want:    `{"key1":{"nestedKey1":"nestedValue1"},"key2":[1,2,3]}`,
			wantErr: false,
		},
		{
			name:    "JSON with special characters",
			input:   []byte(`{"key1":"value1\nvalue2","key2":"value3\tvalue4"}`),
			want:    `{"key1":"value1\nvalue2","key2":"value3\tvalue4"}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := core.BytesToJSONString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToJSONString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr && tt.errString != "" && !strings.Contains(err.Error(), tt.errString) {
				t.Errorf("BytesToJSONString() error = %v, expected error to contain %v", err, tt.errString)
			}
			if got != tt.want {
				t.Errorf("BytesToJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}
