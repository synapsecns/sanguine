package core_test

import (
	"github.com/synapsecns/sanguine/core"
	"reflect"
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
