package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/libs4go/crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/math"
)

const (
	uint32Len = 4
	uint40Len = 5
)

// EncodeState encodes a state.
func EncodeState(state State) ([]byte, error) {
	b := make([]byte, 0)
	originBytes := make([]byte, uint32Len)
	nonceBytes := make([]byte, uint32Len)

	binary.BigEndian.PutUint32(originBytes, state.Origin())
	binary.BigEndian.PutUint32(nonceBytes, state.Nonce())
	root := state.Root()

	b = append(b, root[:]...)
	b = append(b, originBytes...)
	b = append(b, nonceBytes...)
	b = append(b, math.PaddedBigBytes(state.BlockNumber(), uint40Len)...)
	b = append(b, math.PaddedBigBytes(state.Timestamp(), uint40Len)...)

	return b, nil
}

// DecodeState decodes a state.
func DecodeState(toDecode []byte) (State, error) {
	if len(toDecode) != stateSize {
		return nil, fmt.Errorf("invalid state length, expected %d, got %d", stateSize, len(toDecode))
	}

	root := toDecode[stateOffsetRoot:stateOffsetOrigin]
	origin := binary.BigEndian.Uint32(toDecode[stateOffsetOrigin:stateOffsetNonce])
	nonce := binary.BigEndian.Uint32(toDecode[stateOffsetNonce:stateOffsetBlockNumber])
	blockNumber := new(big.Int).SetBytes(toDecode[stateOffsetBlockNumber:stateOffsetTimestamp])
	timestamp := new(big.Int).SetBytes(toDecode[stateOffsetTimestamp:stateSize])

	var rootB32 [32]byte
	copy(rootB32[:], root)

	return state{
		root:        rootB32,
		origin:      origin,
		nonce:       nonce,
		blockNumber: blockNumber,
		timestamp:   timestamp,
	}, nil
}

// EncodeSnapshot encodes a snapshot.
func EncodeSnapshot(snapshot Snapshot) ([]byte, error) {
	states := snapshot.States()

	if len(states) == 0 {
		return nil, fmt.Errorf("no states to encode")
	}

	encodedStates := make([]byte, 0)

	for _, state := range states {
		encodedState, err := EncodeState(state)
		if err != nil {
			return nil, fmt.Errorf("could not encode state: %w", err)
		}
		encodedStates = append(encodedStates, encodedState...)
	}

	return encodedStates, nil
}

// DecodeSnapshot decodes a snapshot.
func DecodeSnapshot(toDecode []byte) (Snapshot, error) {
	var states []State

	if len(toDecode)%stateSize != 0 {
		return nil, fmt.Errorf("invalid snapshot length, expected multiple of %d, got %d", stateSize, len(toDecode))
	}

	for i := 0; i < len(toDecode); i += stateSize {
		state, err := DecodeState(toDecode[i : i+stateSize])
		if err != nil {
			return nil, fmt.Errorf("could not decode state: %w", err)
		}
		states = append(states, state)
	}

	return snapshot{
		states: states,
	}, nil
}

// EncodeAttestation encodes an attestation.
func EncodeAttestation(attestation Attestation) ([]byte, error) {
	b := make([]byte, 0)
	height := []byte{attestation.Height()}
	nonceBytes := make([]byte, uint32Len)

	binary.BigEndian.PutUint32(nonceBytes, attestation.Nonce())
	snapshotRoot := attestation.SnapshotRoot()

	b = append(b, snapshotRoot[:]...)
	b = append(b, height...)
	b = append(b, nonceBytes...)
	b = append(b, math.PaddedBigBytes(attestation.BlockNumber(), uint40Len)...)
	b = append(b, math.PaddedBigBytes(attestation.Timestamp(), uint40Len)...)

	return b, nil
}

// DecodeAttestation decodes an attestation.
func DecodeAttestation(toDecode []byte) (Attestation, error) {
	if len(toDecode) != attestationSize {
		return nil, fmt.Errorf("invalid attestation length, expected %d, got %d", attestationSize, len(toDecode))
	}

	snapshotRoot := toDecode[attestationOffsetRoot:attestationOffsetDepth]
	height := toDecode[attestationOffsetDepth:attestationOffsetNonce][0]
	nonce := binary.BigEndian.Uint32(toDecode[attestationOffsetNonce:attestationOffsetBlockNumber])
	blockNumber := new(big.Int).SetBytes(toDecode[attestationOffsetBlockNumber:attestationOffsetTimestamp])
	timestamp := new(big.Int).SetBytes(toDecode[attestationOffsetTimestamp:attestationSize])

	var snapshotRootB32 [32]byte
	copy(snapshotRootB32[:], snapshotRoot)

	return attestation{
		snapshotRoot: snapshotRootB32,
		height:       height,
		nonce:        nonce,
		blockNumber:  blockNumber,
		timestamp:    timestamp,
	}, nil
}

// EncodeSignature encodes a signature.
func EncodeSignature(sig Signature) ([]byte, error) {
	return ecdsa.Sig2Bytes(secp256k1.S256(), sig.R(), sig.S(), sig.V()), nil
}

// DecodeSignature decodes a signature.
func DecodeSignature(toDecode []byte) (sig Signature, err error) {
	r, s, v, err := ecdsa.Bytes2Sig(secp256k1.S256(), toDecode)
	if err != nil {
		return nil, fmt.Errorf("could not decode signature: %w", err)
	}

	return NewSignature(v, r, s), nil
}

// HashRawBytes takes the raw bytes and produces a hash.
func HashRawBytes(rawBytes []byte) (common.Hash, error) {
	hashedDigest := crypto.Keccak256Hash(rawBytes)

	signedHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), hashedDigest.Bytes())
	return signedHash, nil
}

const (
	//nolint: staticcheck
	tipsVersion       uint16 = 1
	offsetNotary             = 2
	offsetBroadcaster        = 14
	offsetProver             = 26
	offsetExecutor           = 38
	uint96Len                = 12
)

// EncodeTips encodes a list of tips.
//
//nolint:makezero
func EncodeTips(tips Tips) ([]byte, error) {
	b := make([]byte, offsetNotary)
	binary.BigEndian.PutUint16(b, tipsVersion)

	b = append(b, math.PaddedBigBytes(tips.NotaryTip(), uint96Len)...)
	b = append(b, math.PaddedBigBytes(tips.BroadcasterTip(), uint96Len)...)
	b = append(b, math.PaddedBigBytes(tips.ProverTip(), uint96Len)...)
	b = append(b, math.PaddedBigBytes(tips.ExecutorTip(), uint96Len)...)

	return b, nil
}

// DecodeTips decodes a tips typed mem view.
func DecodeTips(toDecode []byte) (Tips, error) {
	notaryTip := new(big.Int).SetBytes(toDecode[offsetNotary:offsetBroadcaster])
	broadcasterTip := new(big.Int).SetBytes(toDecode[offsetBroadcaster:offsetProver])
	proverTip := new(big.Int).SetBytes(toDecode[offsetProver:offsetExecutor])
	executorTip := new(big.Int).SetBytes(toDecode[offsetExecutor:])

	return NewTips(notaryTip, broadcasterTip, proverTip, executorTip), nil
}

type headerEncoder struct {
	Version           uint16
	OriginDomain      uint32
	Sender            [32]byte
	Nonce             uint32
	DestinationDomain uint32
	Recipient         [32]byte
	OptimisticSeconds uint32
}

// EncodeHeader encodes a message header.
func EncodeHeader(header Header) ([]byte, error) {
	newHeader := headerEncoder{
		Version:           header.Version(),
		OriginDomain:      header.OriginDomain(),
		Sender:            header.Sender(),
		Nonce:             header.Nonce(),
		DestinationDomain: header.DestinationDomain(),
		Recipient:         header.Recipient(),
		OptimisticSeconds: header.OptimisticSeconds(),
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, newHeader)
	if err != nil {
		return nil, fmt.Errorf("could not write binary: %w", err)
	}

	return buf.Bytes(), nil
}

// messageEncoder contains the binary structore of the message.
type messageEncoder struct {
	Version      uint16
	HeaderLength uint16
	TipsLength   uint16
}

// EncodeMessage encodes a message.
func EncodeMessage(m Message) ([]byte, error) {
	encodedHeader, err := EncodeHeader(m.Header())
	if err != nil {
		return []byte{}, fmt.Errorf("could not encode header: %w", err)
	}

	encodedTips, err := EncodeTips(m.Tips())
	if err != nil {
		return []byte{}, fmt.Errorf("could not encode tips: %w", err)
	}

	newMessage := messageEncoder{
		Version:      m.Version(),
		HeaderLength: uint16(len(encodedHeader)),
		TipsLength:   uint16(len(encodedTips)),
	}

	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.BigEndian, newMessage)
	if err != nil {
		return nil, fmt.Errorf("could not write binary: %w", err)
	}

	buf.Write(encodedHeader)
	buf.Write(encodedTips)
	buf.Write(m.Body())

	return buf.Bytes(), nil
}
