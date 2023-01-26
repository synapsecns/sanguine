package utils_test

import (
	"github.com/synapsecns/sanguine/contrib/tfcore/utils"
	"testing"
)

func TestCombineMaps(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"c": 3, "d": 4}
	expected := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	combinedMap := utils.MustCombineMaps(m1, m2)

	for key, value := range expected {
		if combinedMap[key] != value {
			t.Errorf("Expected value %d for key %s, but got %d", value, key, combinedMap[key])
		}
	}
}

func TestCombineMapsPanic(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 3, "d": 4}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	_ = utils.MustCombineMaps(m1, m2)
}
