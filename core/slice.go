package core

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

// RandomItem is a generic function to get a random item from a slice.
func RandomItem[T any](items []T) (res T, _ error) {
	if len(items) == 0 {
		return res, fmt.Errorf("empty slice")
	}

	index, err := randInt(len(items))
	if err != nil {
		return res, fmt.Errorf("error generating random index: %w", err)
	}

	return items[index], nil
}

// randInt generates a random integer between 0 (inclusive) and max (exclusive).
func randInt(max int) (int, error) {
	var buf [4]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return 0, fmt.Errorf("error reading random bytes: %w", err)
	}

	// Interpret the buffer as an uint32 and convert to int.
	n := int(binary.BigEndian.Uint32(buf[:]))

	// Limit the range to [0, max).
	return n % max, nil
}
