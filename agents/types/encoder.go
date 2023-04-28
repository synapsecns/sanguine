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
	nonceBytes := make([]byte, uint32Len)

	binary.BigEndian.PutUint32(nonceBytes, attestation.Nonce())
	snapshotRoot := attestation.SnapshotRoot()
	agentRoot := attestation.AgentRoot()

	b = append(b, snapshotRoot[:]...)
	b = append(b, agentRoot[:]...)
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

	snapshotRoot := toDecode[attestationOffsetRoot:attestationOffsetAgentRoot]
	agentRoot := toDecode[attestationOffsetAgentRoot:attestationOffsetNonce]
	nonce := binary.BigEndian.Uint32(toDecode[attestationOffsetNonce:attestationOffsetBlockNumber])
	blockNumber := new(big.Int).SetBytes(toDecode[attestationOffsetBlockNumber:attestationOffsetTimestamp])
	timestamp := new(big.Int).SetBytes(toDecode[attestationOffsetTimestamp:attestationSize])

	var snapshotRootB32, agentRootB32 [32]byte
	copy(snapshotRootB32[:], snapshotRoot)
	copy(agentRootB32[:], agentRoot)

	return attestation{
		snapshotRoot: snapshotRootB32,
		agentRoot:    agentRootB32,
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
	uint64Len = 8
)

// EncodeTips encodes a list of tips.
//
//nolint:makezero
func EncodeTips(tips Tips) ([]byte, error) {
	b := make([]byte, 0)

	b = append(b, math.PaddedBigBytes(tips.SummitTip(), uint64Len)...)
	b = append(b, math.PaddedBigBytes(tips.AttestationTip(), uint64Len)...)
	b = append(b, math.PaddedBigBytes(tips.ExecutionTip(), uint64Len)...)
	b = append(b, math.PaddedBigBytes(tips.DeliveryTip(), uint64Len)...)

	return b, nil
}

// DecodeTips decodes a tips typed mem view.
func DecodeTips(toDecode []byte) (Tips, error) {
	summitTip := new(big.Int).SetBytes(toDecode[0:8])
	attestationTip := new(big.Int).SetBytes(toDecode[8:16])
	executionTip := new(big.Int).SetBytes(toDecode[16:24])
	deliveryTip := new(big.Int).SetBytes(toDecode[24:])

	return NewTips(summitTip, attestationTip, executionTip, deliveryTip), nil
}

type headerEncoder struct {
	OriginDomain      uint32
	Nonce             uint32
	DestinationDomain uint32
	OptimisticSeconds uint32
}

// EncodeHeader encodes a message header.
func EncodeHeader(header Header) ([]byte, error) {
	newHeader := headerEncoder{
		OriginDomain:      header.OriginDomain(),
		Nonce:             header.Nonce(),
		DestinationDomain: header.DestinationDomain(),
		OptimisticSeconds: header.OptimisticSeconds(),
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, newHeader)
	if err != nil {
		return nil, fmt.Errorf("could not write binary: %w", err)
	}

	return buf.Bytes(), nil
}

// EncodeMessage encodes a message.
func EncodeMessage(m Message) ([]byte, error) {
	encodedHeader, err := EncodeHeader(m.Header())
	if err != nil {
		return []byte{}, fmt.Errorf("could not encode header: %w", err)
	}

	buf := new(bytes.Buffer)

	buf.Write([]byte{uint8(m.Flag())})
	buf.Write(encodedHeader)
	buf.Write(m.Body())

	return buf.Bytes(), nil
}

// EncodeAgentStatus encodes a agent status.
func EncodeAgentStatus(agentStatus AgentStatus) ([]byte, error) {
	b := make([]byte, 0)
	domainBytes := make([]byte, uint32Len)
	indexBytes := make([]byte, uint32Len)

	binary.BigEndian.PutUint32(domainBytes, agentStatus.Domain())
	binary.BigEndian.PutUint32(indexBytes, agentStatus.Index())

	b = append(b, agentStatus.Flag())
	b = append(b, domainBytes...)
	b = append(b, indexBytes...)

	return b, nil
}

// DecodeAgentStatus decodes an agent status.
func DecodeAgentStatus(toDecode []byte) (AgentStatus, error) {
	if len(toDecode) != agentStatusSize {
		return nil, fmt.Errorf("invalid agent status length, expected %d, got %d", agentStatusSize, len(toDecode))
	}

	flagBytes := toDecode[agentStatusOffsetFlag:agentStatusOffsetDomain]
	domain := binary.BigEndian.Uint32(toDecode[agentStatusOffsetDomain:agentStatusOffsetIndex])
	index := binary.BigEndian.Uint32(toDecode[agentStatusOffsetDomain:agentStatusSize])

	return agentStatus{
		flag:   flagBytes[0],
		domain: domain,
		index:  index,
	}, nil
}
