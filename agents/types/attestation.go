package types

import (
	"encoding/binary"
	"math/big"
)

// Attestation is the attestation.
// TODO (joe): Add tests for the AttestationKey and AttestedDomains,
// converting to and from serialized big.Int form
type Attestation interface {
	// Origin gets the origin of the attestation
	Origin() uint32
	// Destination gets the destination of the attestation
	Destination() uint32
	// Nonce gets the nonce of the attestation
	Nonce() uint32
	// Root gets the root of the contract
	Root() [32]byte
}

type attestation struct {
	// origin of the attestation
	origin uint32
	// destination of the attestation
	destination uint32
	// nonce of the attestation
	nonce uint32
	// root - the merkle root
	root [32]byte
}

// AttestationKey is the tuple (origin, destination, nonce)
type AttestationKey struct {
	// Origin of the attestation
	Origin uint32
	// Destination of the attestation
	Destination uint32
	// Nonce of the attestation
	Nonce uint32
}

// AttestatedDomains is the tuple (origin, destination)
type AttestatedDomains struct {
	// Origin of the attestation
	Origin uint32
	// Destination of the attestation
	Destination uint32
}

// NewAttestation creates a new attestation.
func NewAttestation(rawKey *big.Int, root [32]byte) Attestation {
	key := NewAttestionKey(rawKey)
	return attestation{
		origin:      key.Origin,
		destination: key.Destination,
		nonce:       key.Nonce,
		root:        root,
	}
}

func (a attestation) Origin() uint32 {
	return a.origin
}

func (a attestation) Destination() uint32 {
	return a.destination
}

func (a attestation) Nonce() uint32 {
	return a.nonce
}

func (a attestation) Root() [32]byte {
	return a.root
}

var _ Attestation = attestation{}

// SignedAttestation is a signed attestation.
type SignedAttestation interface {
	// Attestation gets the unsigned attestation
	Attestation() Attestation
	// Signature is the signature of the attestation
	Signature() Signature
}

// signedAttestation is a struct that conforms to SignedAttestation.
type signedAttestation struct {
	attestation Attestation
	signature   Signature
}

// NewSignedAttestation creates a new signed attestation.
func NewSignedAttestation(attestation Attestation, signature Signature) SignedAttestation {
	return signedAttestation{
		attestation: attestation,
		signature:   signature,
	}
}

func (s signedAttestation) Attestation() Attestation {
	return s.attestation
}

func (s signedAttestation) Signature() Signature {
	return s.signature
}

// NewAttestionKey takes the raw AttestationKey serialized as a big endian big.Int
// and converts it to AttestationKey which is a tuple of (origin, destination, nonce)
func NewAttestionKey(rawKey *big.Int) AttestationKey {
	rawBytes := make([]byte, 32)
	rawKey.FillBytes(rawBytes)
	originBytes := rawBytes[8:12]
	destinationBytes := rawBytes[4:8]
	nonceBytes := rawBytes[0:4]
	origin := binary.BigEndian.Uint32(originBytes)
	destination := binary.BigEndian.Uint32(destinationBytes)
	nonce := binary.BigEndian.Uint32(nonceBytes)
	return AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
}

// GetRawKey returns the AttestationKey as a serialized big.Int
func (a AttestationKey) GetRawKey() *big.Int {
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, a.Origin)
	destinationBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(destinationBytes, a.Destination)
	nonceBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nonceBytes, a.Nonce)
	rawBytes := make([]byte, 32)
	copy(rawBytes[0:4], nonceBytes[:])
	copy(rawBytes[4:8], destinationBytes[:])
	copy(rawBytes[8:12], originBytes[:])
	rawKey := new(big.Int)
	rawKey.SetBytes(rawBytes)
	return rawKey
}

// NewAttestatedDomains takes the raw AttestedDomains serialized as a big endian big.Int
// and converts it to AttestatedDomains which is a tuple of (origin, destination)
func NewAttestatedDomains(rawDomains *big.Int) AttestatedDomains {
	rawBytes := make([]byte, 32)
	rawDomains.FillBytes(rawBytes)
	originBytes := rawBytes[4:8]
	destinationBytes := rawBytes[0:4]

	origin := binary.BigEndian.Uint32(originBytes)
	destination := binary.BigEndian.Uint32(destinationBytes)

	return AttestatedDomains{
		Origin:      origin,
		Destination: destination,
	}
}

// GetRawDomains returns the AttestatedDomains which is a tuple of (origin, destination)
// as a serialized big.Int
func (a AttestatedDomains) GetRawDomains() *big.Int {
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, a.Origin)
	destinationBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(destinationBytes, a.Destination)
	rawBytes := make([]byte, 32)
	copy(rawBytes[0:4], destinationBytes[:])
	copy(rawBytes[4:8], originBytes[:])
	rawDomains := new(big.Int)
	rawDomains.SetBytes(rawBytes)
	return rawDomains
}

var _ SignedAttestation = signedAttestation{}
