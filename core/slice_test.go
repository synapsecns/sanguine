package core_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/core"
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
