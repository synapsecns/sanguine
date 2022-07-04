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
