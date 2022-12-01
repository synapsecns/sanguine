package types

import (
	"encoding/binary"
	"math/big"
	"time"
)

const sizeOfUint256 = uint32(32)
const sizeOfUint32 = uint32(4)

const attestationNonceStartingByte = uint32(0)
const attestationDestinationStartingByte = uint32(4)
const attestationOriginStartingByte = uint32(8)
const attestationRootStartingByte = uint32(12)
const attestationSize = uint32(44)
const attestationRawKeyStartingByte = uint32(0)

const attestationKeyNonceStartingByte = uint32(0)
const attestationKeyDestinationStartingByte = uint32(4)
const attestationKeyOriginStartingByte = uint32(8)
const attestationKeySize = uint32(12)

const attestedDomainsDestinationStartingByte = uint32(0)
const attestedDomainsOriginStartingByte = uint32(4)
const attestedDomainsSize = uint32(8)

// Attestation is the attestation.
// TODO (joe): Add tests for the AttestationKey and AttestedDomains,
// converting to and from serialized big.Int form.
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

// AttestationKey is the tuple (origin, destination, nonce).
type AttestationKey struct {
	// Origin of the attestation
	Origin uint32
	// Destination of the attestation
	Destination uint32
	// Nonce of the attestation
	Nonce uint32
}

// AttestedDomains is the tuple (origin, destination).
type AttestedDomains struct {
	// Origin of the attestation
	Origin uint32
	// Destination of the attestation
	Destination uint32
}

// NewAttestationFromBytes creates a new attesation from raw bytes.
func NewAttestationFromBytes(rawBytes []byte) Attestation {
	rootBytes := rawBytes[attestationRootStartingByte:attestationSize]

	rawKeyBytes := rawBytes[attestationRawKeyStartingByte:attestationKeySize]
	originBytes := rawKeyBytes[attestationOriginStartingByte:attestationRootStartingByte]
	destinationBytes := rawKeyBytes[attestationDestinationStartingByte:attestationOriginStartingByte]
	nonceBytes := rawKeyBytes[attestationNonceStartingByte:attestationDestinationStartingByte]
	origin := binary.BigEndian.Uint32(originBytes)
	destination := binary.BigEndian.Uint32(destinationBytes)
	nonce := binary.BigEndian.Uint32(nonceBytes)
	var root [32]byte
	copy(root[:], rootBytes)
	return attestation{
		origin:      origin,
		destination: destination,
		nonce:       nonce,
		root:        root,
	}
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

var _ SignedAttestation = signedAttestation{}

// InProgressAttestation is an attestation to be processed by offline agent.
type InProgressAttestation interface {
	// SignedAttestation gets the signed attestation
	SignedAttestation() SignedAttestation
	// DispatchedBlockNumber when message was dispatched on origin
	OriginDispatchBlockNumber() uint64
	// SubmittedToAttestationCollectorTime is time when signed attestation was submitted to AttestationCollector
	SubmittedToAttestationCollectorTime() *time.Time
	// ConfirmedOnAttestationCollectorBlockNumber is block number when we confirmed the attesation posted on AttestationCollector
	ConfirmedOnAttestationCollectorBlockNumber() uint64
}

// inProgressAttestation is a struct that conforms to InProgressAttestation.
type inProgressAttestation struct {
	signedAttestation                          SignedAttestation
	originDispatchBlockNumber                  uint64
	submittedToAttestationCollectorTime        *time.Time
	confirmedOnAttestationCollectorBlockNumber uint64
}

// NewInProgressAttestation creates a new to process attestation.
func NewInProgressAttestation(signedAttestation SignedAttestation, originDispatchBlockNumber uint64, submittedToAttestationCollectorTime *time.Time, confirmedOnAttestationCollectorBlockNumber uint64) InProgressAttestation {
	return inProgressAttestation{
		signedAttestation:                          signedAttestation,
		originDispatchBlockNumber:                  originDispatchBlockNumber,
		submittedToAttestationCollectorTime:        submittedToAttestationCollectorTime,
		confirmedOnAttestationCollectorBlockNumber: confirmedOnAttestationCollectorBlockNumber,
	}
}

func (t inProgressAttestation) SignedAttestation() SignedAttestation {
	return t.signedAttestation
}

func (t inProgressAttestation) OriginDispatchBlockNumber() uint64 {
	return t.originDispatchBlockNumber
}

func (t inProgressAttestation) SubmittedToAttestationCollectorTime() *time.Time {
	return t.submittedToAttestationCollectorTime
}

func (t inProgressAttestation) ConfirmedOnAttestationCollectorBlockNumber() uint64 {
	return t.confirmedOnAttestationCollectorBlockNumber
}

var _ InProgressAttestation = inProgressAttestation{}

// NewAttestionKey takes the raw AttestationKey serialized as a big endian big.Int
// and converts it to AttestationKey which is a tuple of (origin, destination, nonce).
func NewAttestionKey(rawKey *big.Int) AttestationKey {
	rawBytes := make([]byte, sizeOfUint256)
	rawKey.FillBytes(rawBytes)
	originBytes := rawBytes[attestationKeyOriginStartingByte:attestationKeySize]
	destinationBytes := rawBytes[attestationKeyDestinationStartingByte:attestationKeyOriginStartingByte]
	nonceBytes := rawBytes[attestationKeyNonceStartingByte:attestationKeyDestinationStartingByte]
	origin := binary.BigEndian.Uint32(originBytes)
	destination := binary.BigEndian.Uint32(destinationBytes)
	nonce := binary.BigEndian.Uint32(nonceBytes)
	return AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
}

// GetRawKey returns the AttestationKey as a serialized big.Int.
func (a AttestationKey) GetRawKey() *big.Int {
	originBytes := make([]byte, sizeOfUint32)
	binary.BigEndian.PutUint32(originBytes, a.Origin)
	destinationBytes := make([]byte, sizeOfUint32)
	binary.BigEndian.PutUint32(destinationBytes, a.Destination)
	nonceBytes := make([]byte, sizeOfUint32)
	binary.BigEndian.PutUint32(nonceBytes, a.Nonce)
	rawBytes := make([]byte, sizeOfUint256)
	copy(rawBytes[attestationKeyNonceStartingByte:attestationKeyDestinationStartingByte], nonceBytes)
	copy(rawBytes[attestationKeyDestinationStartingByte:attestationKeyOriginStartingByte], destinationBytes)
	copy(rawBytes[attestationKeyOriginStartingByte:sizeOfUint256], originBytes)
	rawKey := new(big.Int)
	rawKey.SetBytes(rawBytes)
	return rawKey
}

// NewAttestedDomains takes the raw AttestedDomains serialized as a big endian big.Int
// and converts it to AttestedDomains which is a tuple of (origin, destination).
func NewAttestedDomains(rawDomains *big.Int) AttestedDomains {
	rawBytes := make([]byte, sizeOfUint256)
	rawDomains.FillBytes(rawBytes)
	originBytes := rawBytes[attestedDomainsOriginStartingByte:attestedDomainsSize]
	destinationBytes := rawBytes[attestedDomainsDestinationStartingByte:attestedDomainsOriginStartingByte]

	origin := binary.BigEndian.Uint32(originBytes)
	destination := binary.BigEndian.Uint32(destinationBytes)

	return AttestedDomains{
		Origin:      origin,
		Destination: destination,
	}
}

// GetRawDomains returns the AttestedDomains which is a tuple of (origin, destination)
// as a serialized big.Int.
func (a AttestedDomains) GetRawDomains() *big.Int {
	originBytes := make([]byte, sizeOfUint32)
	binary.BigEndian.PutUint32(originBytes, a.Origin)
	destinationBytes := make([]byte, sizeOfUint32)
	binary.BigEndian.PutUint32(destinationBytes, a.Destination)
	rawBytes := make([]byte, sizeOfUint256)
	copy(rawBytes[attestedDomainsDestinationStartingByte:attestedDomainsOriginStartingByte], destinationBytes)
	copy(rawBytes[attestedDomainsOriginStartingByte:attestedDomainsSize], originBytes)
	rawDomains := new(big.Int)
	rawDomains.SetBytes(rawBytes)
	return rawDomains
}
