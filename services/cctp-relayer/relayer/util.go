package relayer

import (
	"encoding/binary"
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

// destinationDomainIndex is the index of the destination domain in the formatted CCTP message.
const destinationDomainIndex = 8

// ParseDestDomain parses the destination domain from a CCTP message.
func ParseDestDomain(message []byte) (uint32, error) {
	return indexUint32(message, destinationDomainIndex, 4)
}

// senderIndex is the index of the sender in the formatted CCTP message.
const senderIndex = 20

// ParseSender parses the sender from a CCTP message.
func ParseSender(message []byte) (common.Address, error) {
	return indexAddress(message, senderIndex)
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

const addressLength = 20 // Ethereum address length
const bytes32Length = 32 // bytes32 length

// indexAddress parses a bytes32 value from a byte slice starting at index and converts it to a common.Address.
// Assumes that the bytes32 value represents an Ethereum address right-padded to 32 bytes.
func indexAddress(memView []byte, index int) (common.Address, error) {
	// Ensure the slice has enough bytes to extract bytes32 value
	if index+bytes32Length > len(memView) {
		return common.Address{}, errors.New("slice does not have enough bytes to extract bytes32")
	}

	// Extract the rightmost 20 bytes of the bytes32 value
	addressBytes := memView[index+bytes32Length-addressLength : index+bytes32Length]

	// Convert the extracted bytes to common.Address
	var address common.Address
	copy(address[:], addressBytes)

	return address, nil
}

// AddressToBytes32 converts an Ethereum address to a bytes32 value.
func AddressToBytes32(addr common.Address) [32]byte {
	var buf [32]byte
	copy(buf[12:], addr.Bytes())
	return buf
}

// Bytes32ToAddress converts a bytes32 value to an Ethereum address.
func Bytes32ToAddress(bytes32 [32]byte) common.Address {
	return common.BytesToAddress(bytes32[12:])
}
