package core_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"reflect"
	"testing"
)

func TestRandomItem(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"}

	randInt, err := core.RandomItem(intSlice)
	if err != nil {
		t.Fatalf("Error getting random item from intSlice: %v", err)
	}
	fmt.Printf("Random int: %v\n", randInt)

	randString, err := core.RandomItem(stringSlice)
	if err != nil {
		t.Fatalf("Error getting random item from stringSlice: %v", err)
	}
	fmt.Printf("Random string: %v\n", randString)

	var emptySlice []int
	_, err = core.RandomItem(emptySlice)
	if err == nil {
		t.Fatalf("Expected error when getting random item from empty slice, got nil")
	}
}

func TestChunkSlice(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		chunkSize int
		want      [][]int
	}{
		{
			name:      "Empty slice",
			slice:     []int{},
			chunkSize: 2,
			want:      [][]int{},
		},
		{
			name:      "Slice smaller than chunk size",
			slice:     []int{1, 2},
			chunkSize: 5,
			want:      [][]int{{1, 2}},
		},
		{
			name:      "Slice size equal to chunk size",
			slice:     []int{1, 2, 3},
			chunkSize: 3,
			want:      [][]int{{1, 2, 3}},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := core.ChunkSlice(tt.slice, tt.chunkSize); !reflect.DeepEqual(got, tt.want) {
				if len(got) == len(tt.want) && len(got) == 0 {
					return
				}
				t.Errorf("ChunkSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
