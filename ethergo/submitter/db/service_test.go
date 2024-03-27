package db_test

import (
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"golang.org/x/exp/slices"
	"math"
	"testing"
)

// TestScanValid tests that invalid statuses can be scanned.
func TestScanInvalid(t *testing.T) {
	allStatuses := getAllStatusInts()
	var i uint8

	for i = 1; i < math.MaxUint8; i++ {
		status := db.Status(0)

		// skip valid statuses
		if slices.Contains(allStatuses, i) {
			continue
		}

		err := status.Scan(int32(i))
		if err == nil {
			t.Fatalf("expected error, got nil for %d", i)
		}
	}
}

func getAllStatusInts() []uint8 {
	allStatuses := db.GetAllStatusTypes()
	var allStatusInts []uint8
	for _, status := range allStatuses {
		allStatusInts = append(allStatusInts, status.Int())
	}
	return allStatusInts
}
