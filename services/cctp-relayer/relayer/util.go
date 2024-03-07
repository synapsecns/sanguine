package relayer

import (
	"encoding/binary"
	"errors"
)

// destinationDomainIndex is the index of the destination domain in the formatted CCTP message.
const destinationDomainIndex = 8 * 9 // TODO: why multiply by 9 here?

func parseDestDomain(message []byte) (uint32, error) {
	return indexUint32(message, destinationDomainIndex, 4)
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
