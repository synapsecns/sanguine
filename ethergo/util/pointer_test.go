package util_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/util"
	"testing"
)

func TestLogPointer(t *testing.T) {
	mockLogs := mocks.GetMockLogs(t, 2)
	mockPointerLogs := util.LogsPointer(mockLogs)

	for i := range mockLogs {
		Equal(t, mockLogs[i], *mockPointerLogs[i])
	}
}
