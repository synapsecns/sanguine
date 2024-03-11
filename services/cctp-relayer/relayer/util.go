package relayer

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/util"
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

// GetTxSender gets the sender of a transaction by fetching the transaction metadata.
func GetTxSender(ctx context.Context, txHash common.Hash, ethClient client.EVM) (sender common.Address, err error) {
	tx, _, err := ethClient.TransactionByHash(ctx, txHash)
	if err != nil {
		return sender, fmt.Errorf("could not get transaction by hash: %w", err)
	}
	call, err := util.TxToCall(tx)
	if err != nil {
		return sender, fmt.Errorf("could not convert transaction to call: %w", err)
	}
	return call.From, nil
}
