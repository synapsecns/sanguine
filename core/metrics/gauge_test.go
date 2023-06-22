package metrics

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdates(t *testing.T) {
	executed := false
	// Create new guage
	testFunc := func() error {
		executed = true
		return nil
	}
	guage := NewGauge(10000, testFunc)
	err := guage.Update(nil, 1)
	True(t, err != nil)
	True(t, executed)
}
