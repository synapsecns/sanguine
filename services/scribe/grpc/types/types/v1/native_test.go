package pbscribe_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"testing"
)

func TestAddressConversion(t *testing.T) {
	for i := 0; i < 500; i++ {
		randAddress := mocks.MockAddress()
		nativeAddress := pbscribe.FromNativeAddress(randAddress)
		reNativeAddress := nativeAddress.ToAddress()

		Equal(t, reNativeAddress, randAddress)
	}
}

func TestHashConversion(t *testing.T) {
	var nativeHashes []common.Hash
	for i := 0; i < 500; i++ {
		randHash := mocks.NewMockHash(t)
		// add to slice for group conversion test
		nativeHashes = append(nativeHashes, randHash)

		nativeHash := pbscribe.FromNativeHash(randHash)
		reNativeHash := nativeHash.ToHash()

		Equal(t, reNativeHash, randHash)
	}

	convertedHashes := pbscribe.FromNativeHashes(nativeHashes)
	ogHashes := pbscribe.ToNativeHashes(convertedHashes)

	Equal(t, ogHashes, nativeHashes)
}

func TestLogsConversion(t *testing.T) {
	mockLogs := LogsPointer(mocks.GetMockLogs(t, 500))
	convertedLogs := pbscribe.FromNativeLogs(mockLogs)
	reConvertedLogs := pbscribe.ToNativeLogs(convertedLogs)

	Equal(t, mockLogs, reConvertedLogs)
}

// LogsPointer wraps logs in a pointer.
func LogsPointer(logs []types.Log) (res []*types.Log) {
	for _, log := range logs {
		res = append(res, core.PtrTo(log))
	}
	return res
}
