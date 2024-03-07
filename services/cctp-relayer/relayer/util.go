package relayer

import (
	"encoding/binary"
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

// destinationDomainIndex is the index of the destination domain in the formatted CCTP message.
const destinationDomainIndex = 8

// recipientIndex is the index of the recipient in the formatted CCTP message.
// Note that we add 12 since eth address is only 10 bytes and have placeholder for 32.
const recipientIndex = 52 + 12

func parseDestDomain(message []byte) (uint32, error) {
	return indexUint32(message, destinationDomainIndex, 4)
}

func parseRecipient(message []byte) (common.Address, error) {
	return indexAddress(message, recipientIndex)
}

// indexUint32 parses an unsigned 32-bit integer from a byte slice starting at index with a specified length in bytes.
// Requires that the byte slice have enough bytes starting from the index.
func indexUint32(memView []byte, index int, bytes uint8) (uint32, error) {
	if int(bytes)+index > len(memView) || bytes > 4 {
		return 0, errors.New("invalid index or byte length")
	}

	// Ensure we are not trying to parse more than 4 bytes for a uint32.
	if bytes > 4 {
		return 0, errors.New("byte length too large for uint32")
	}

	// Pad the byte slice to 4 bytes if necessary.
	padded := make([]byte, 4)
	copy(padded[4-int(bytes):], memView[index:index+int(bytes)])

	// Convert the bytes to uint32.
	result := binary.BigEndian.Uint32(padded)
	return result, nil
}

// indexAddress parses an Ethereum address from a byte slice starting at index.
// Requires that the byte slice have at least 20 bytes starting from the index.
func indexAddress(memView []byte, index int) (common.Address, error) {
	const addressLength = 20 // Ethereum addresses are 20 bytes long

	if index+addressLength > len(memView) {
		return common.Address{}, errors.New("slice does not contain enough bytes to extract an address")
	}

	var addr common.Address
	copy(addr[:], memView[index:index+addressLength])
	return addr, nil
}
