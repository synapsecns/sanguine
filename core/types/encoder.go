package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/libs4go/crypto/ecdsa"
)

// EncodeSignedUpdate encodes a signed update in binary.
func EncodeSignedUpdate(signed SignedUpdate) ([]byte, error) {
	encodedUpdate, err := EncodeUpdate(signed.Update())
	if err != nil {
		return nil, fmt.Errorf("could not encode update: %w", err)
	}
	encodedSignature, err := EncodeSignature(signed.Signature())
	if err != nil {
		return nil, fmt.Errorf("could not encode signature: %w", err)
	}

	return append(encodedUpdate, encodedSignature...), nil
}

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

	signedAttesationSize := binary.Size(decAttestation)

	attestationBin := toDecode[0:signedAttesationSize]
	signBin := toDecode[signedAttesationSize:]

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

// DecodeSignedUpdate decodes a signed update.
func DecodeSignedUpdate(toDecode []byte) (SignedUpdate, error) {
	var decUpdate updateEncoder

	decUpdateSize := binary.Size(decUpdate)

	updateBin := toDecode[0:decUpdateSize]
	signBin := toDecode[decUpdateSize:]

	upd, err := DecodeUpdate(updateBin)
	if err != nil {
		return nil, fmt.Errorf("could not decode update: %w", err)
	}

	sig, err := DecodeSignature(signBin)
	if err != nil {
		return nil, fmt.Errorf("could not decode signature: %w", err)
	}

	return NewSignedUpdate(upd, sig), nil
}

// updateEncoder encodes/decodes updates in binary.
type updateEncoder struct {
	HomeDomain   uint32
	PreviousRoot [32]byte
	NewRoot      [32]byte
}

// EncodeUpdate encodes an update to a byte slice.
func EncodeUpdate(update Update) ([]byte, error) {
	buf := new(bytes.Buffer)

	encodedUpdate := updateEncoder{
		HomeDomain:   update.HomeDomain(),
		PreviousRoot: update.PreviousRoot(),
		NewRoot:      update.NewRoot(),
	}

	err := binary.Write(buf, binary.BigEndian, encodedUpdate)
	if err != nil {
		return nil, fmt.Errorf("could not write binary: %w", err)
	}

	return buf.Bytes(), nil
}

// DecodeUpdate decodes an update.
func DecodeUpdate(toDecode []byte) (Update, error) {
	reader := bytes.NewReader(toDecode)

	var encodedUpdate updateEncoder
	dataSize := binary.Size(encodedUpdate)

	if dataSize > len(toDecode) {
		return nil, fmt.Errorf("message too small, expected at least %d, got %d", dataSize, len(toDecode))
	}

	err := binary.Read(reader, binary.BigEndian, &encodedUpdate)
	if err != nil {
		return nil, fmt.Errorf("could not read: %w", err)
	}

	return update{
		homeDomain:   encodedUpdate.HomeDomain,
		previousRoot: encodedUpdate.PreviousRoot,
		newRoot:      encodedUpdate.NewRoot,
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

// attestationEncoder encodes attestations.
type attestationEncoder struct {
	Domain, Nonce uint32
	Root          [32]byte
}

// EncodeAttestation encodes an attestation.
func EncodeAttestation(attestation Attestation) ([]byte, error) {
	buf := new(bytes.Buffer)

	encodedUpdate := attestationEncoder{
		Domain: attestation.Domain(),
		Nonce:  attestation.Nonce(),
		Root:   attestation.Root(),
	}

	err := binary.Write(buf, binary.BigEndian, encodedUpdate)
	if err != nil {
		return nil, fmt.Errorf("could not write binary: %w", err)
	}

	return buf.Bytes(), nil
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
		domain: encodedAttestation.Domain,
		nonce:  encodedAttestation.Nonce,
		root:   encodedAttestation.Root,
	}, nil
}
