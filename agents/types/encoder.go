package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/libs4go/crypto/ecdsa"
)

// EncodeSignedAttestation encodes a signed attestation.
func EncodeSignedAttestation(signed SignedAttestation) ([]byte, error) {
	encodedAttestation, err := EncodeAttestation(signed.Attestation())
	if err != nil {
		return nil, fmt.Errorf("could not encode attestation: %w", err)
	}

	encodedSignature, err := EncodeSignature(signed.Signature())
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	return append(encodedAttestation, encodedSignature...), nil
}

// DecodeSignedAttestation decodes a signed attestation.
func DecodeSignedAttestation(toDecode []byte) (SignedAttestation, error) {
	var decAttestation attestation

	signedAttestationSize := binary.Size(decAttestation)

	attestationBin := toDecode[0:signedAttestationSize]
	signBin := toDecode[signedAttestationSize:]

	att, err := DecodeAttestation(attestationBin)
	if err != nil {
		return nil, fmt.Errorf("could not decode attestation: %w", err)
	}

	sig, err := DecodeSignature(signBin)
	if err != nil {
		return nil, fmt.Errorf("could not decode signature: %w", err)
	}

	return NewSignedAttestation(att, sig), nil
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

// attestationEncoder encodes attestations.
type attestationEncoder struct {
	Origin, Destination, Nonce uint32
	Root                       [32]byte
}

// EncodeAttestation encodes an attestation.
func EncodeAttestation(attestation Attestation) ([]byte, error) {
	buf := new(bytes.Buffer)

	encodedUpdate := attestationEncoder{
		Origin:      attestation.Origin(),
		Destination: attestation.Destination(),
		Nonce:       attestation.Nonce(),
		Root:        attestation.Root(),
	}

	err := binary.Write(buf, binary.BigEndian, encodedUpdate)
	if err != nil {
		return nil, fmt.Errorf("could not write binary: %w", err)
	}

	return buf.Bytes(), nil
}

func Hash(a Attestation) ([32]byte, error) {
	encodedAttestation, err := EncodeAttestation(a)
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not encode attestation: %w", err)
	}

	hashedDigest := crypto.Keccak256Hash(encodedAttestation)

	signedHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), hashedDigest.Bytes())
	return signedHash, nil
}

// DecodeAttestation decodes an attestation.
func DecodeAttestation(toDecode []byte) (Attestation, error) {
	reader := bytes.NewReader(toDecode)

	var encodedAttestation attestationEncoder
	dataSize := binary.Size(encodedAttestation)

	if dataSize > len(toDecode) {
		return nil, fmt.Errorf("message too small, expected at least %d, got %d", dataSize, len(toDecode))
	}

	err := binary.Read(reader, binary.BigEndian, &encodedAttestation)
	if err != nil {
		return nil, fmt.Errorf("could not read: %w", err)
	}

	return attestation{
		origin:      encodedAttestation.Origin,
		destination: encodedAttestation.Destination,
		nonce:       encodedAttestation.Nonce,
		root:        encodedAttestation.Root,
	}, nil
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
// nolint: makezero
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
