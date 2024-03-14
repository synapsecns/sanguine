package relayer

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/client"
	tokenmessenger "github.com/synapsecns/sanguine/services/cctp-relayer/contracts/tokenmessenger"
)

// destinationDomainIndex is the index of the destination domain in the formatted CCTP message.
const destinationDomainIndex = 8

// ParseDestDomain parses the destination domain from a CCTP message.
func ParseDestDomain(message []byte) (uint32, error) {
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

// GetCircleRequestID returns a request ID from a source domain and nonce.
func GetCircleRequestID(sourceDomain uint32, nonce uint64) string {
	return fmt.Sprintf("%d-%d", sourceDomain, nonce)
}

// GetMessageTransmitterAddress gets the message transmitter address from a token messenger contract.
func GetMessageTransmitterAddress(ctx context.Context, tokenMessengerAddr common.Address, ethClient client.EVM) (transmitterAddr common.Address, err error) {
	messengerContract, err := tokenmessenger.NewTokenMessenger(tokenMessengerAddr, ethClient)
	if err != nil {
		return transmitterAddr, fmt.Errorf("could not get token messenger contract: %w", err)
	}
	transmitterAddr, err = messengerContract.LocalMessageTransmitter(&bind.CallOpts{Context: ctx})
	if err != nil {
		return transmitterAddr, fmt.Errorf("could not get local message transmitter: %w", err)
	}
	return transmitterAddr, nil
}
