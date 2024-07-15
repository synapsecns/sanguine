package relayer_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
)

func TestParseDestDomain(t *testing.T) {
	// from transaction: https://sepolia.etherscan.io/tx/0x979b337d0e8ce86ffc5388d407e3033d359b0bc16564c0b7c101e1e17c82344b#eventlog
	messageBytes, err := hexutil.Decode("0x000000000000000000000006000000000003edd20000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa50000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa50000000000000000000000002703483b1a5a7c577e8680de9df8be03c6f30e3c000000000000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c7238000000000000000000000000d1a13c794c87122d700aeb8ece2391bd77ee32e300000000000000000000000000000000000000000000000000000000009896800000000000000000000000002703483b1a5a7c577e8680de9df8be03c6f30e3c")
	assert.NoError(t, err)

	domain, err := relayer.ParseDestDomain(messageBytes)
	assert.NoError(t, err)
	assert.Equal(t, uint32(6), domain)
}
