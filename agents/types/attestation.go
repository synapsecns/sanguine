package types

import (
	"encoding/binary"
	"math/big"
	"time"
)

// AttestationState is the state the attestation is in, in terms of being processed by an agent.
type AttestationState uint32

const (
	// AttestationStateNotaryUnsigned is when attestation has been fetched from origin but not yet signed.
	AttestationStateNotaryUnsigned AttestationState = iota // 0
	// AttestationStateNotarySignedUnsubmitted is when attestation has been signed but not yet submitted to the attestation collector.
	AttestationStateNotarySignedUnsubmitted // 1
	// AttestationStateNotarySubmittedUnconfirmed is when the attestation was signed and submitted to the attestation collector but not yet confirmed.
	AttestationStateNotarySubmittedUnconfirmed // 2
	// AttestationStateNotaryConfirmed is when the attestation was confirmed as posted on the attestation collector.
	AttestationStateNotaryConfirmed // 3
	// AttestationStateGuardInitialState is when the attestation is fetched from origin. This is a temporary hack state.
	AttestationStateGuardInitialState // 4 // 4
	// AttestationStateGuardUnsignedAndVerified is when the attestation was signed by Notary but not yet by the Guard, but Guard verified it on origin.
	AttestationStateGuardUnsignedAndVerified // 5
	// AttestationStateGuardSignedUnsubmitted is when the attestation was signed by Guard (and Notary) but not yet submitted.
	AttestationStateGuardSignedUnsubmitted // 6
	// AttestationStateGuardSubmittedToCollectorUnconfirmed is when the attestation was signed by Guard and Notary and submitted to the attestation collector but not destination,
	// but we have yet to confirm it on the AttestationCollector.
	AttestationStateGuardSubmittedToCollectorUnconfirmed // 7
	// AttestationStateGuardConfirmedOnCollector is when the attestation was signed by Guard and Notary and submitted to the attestation collector but not destination,
	// and we have confirmed it on the AttestationCollector.
	AttestationStateGuardConfirmedOnCollector // 8
	// AttestationStateSubmittedToDestinationUnconfirmed is when the attestation was signed by Guard and Notary and submitted to the attestation collector and destination but not yet confirmed on destination.
	AttestationStateSubmittedToDestinationUnconfirmed // 9
	// AttestationStateConfirmedOnDestination is when the attestation was confirmed as posted on the destination.
	AttestationStateConfirmedOnDestination // 10
)

const sizeOfUint256 = uint32(32)
const sizeOfUint32 = uint32(4)
const sizeOfUint16 = uint32(2)
const sizeOfUint8 = uint32(1)

const attestationKeyNonceStartingByte = uint32(0)
const attestationKeyDestinationStartingByte = uint32(4)
const attestationKeyOriginStartingByte = uint32(8)
const attestationKeySize = uint32(12)

const attestedDomainsDestinationStartingByte = uint32(0)
const attestedDomainsOriginStartingByte = uint32(4)
const attestedDomainsSize = uint32(8)

const attestationAgentCountsNotaryCountStartingByte = uint32(1)
const attestationAgentCountsGuardCountStartingByte = uint32(0)
const attestationAgentCountsSize = uint32(2)

// Attestation is the attestation.
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

// AttestationAgentCounts is the tuple (GuardCount, NotaryCount) so we know how many of each type of signature we have.
type AttestationAgentCounts struct {
	// GuardCount is the number of guard signatures collected in the SignedAttestation.
	GuardCount uint32
	// NotaryCount is the number of notary signatures collected in the SignedAttestation.
	NotaryCount uint32
}

// NewAttestation creates a new attestation.
func NewAttestation(rawKey *big.Int, root [32]byte) Attestation {
	key := NewAttestationKey(rawKey)
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
	// GuardSignatures is the guard signatures of the attestation
	GuardSignatures() []Signature
	// NotarySignatures is the notary signatures of the attestation
	NotarySignatures() []Signature
}

// signedAttestation is a struct that conforms to SignedAttestation.
type signedAttestation struct {
	attestation      Attestation
	guardSignatures  []Signature
	notarySignatures []Signature
}

// NewSignedAttestation creates a new signed attestation.
func NewSignedAttestation(attestation Attestation, guardSignatures []Signature, notarySignatures []Signature) SignedAttestation {
	return signedAttestation{
		attestation:      attestation,
		guardSignatures:  guardSignatures,
		notarySignatures: notarySignatures,
	}
}

func (s signedAttestation) Attestation() Attestation {
	return s.attestation
}

func (s signedAttestation) GuardSignatures() []Signature {
	return s.guardSignatures
}

func (s signedAttestation) NotarySignatures() []Signature {
	return s.notarySignatures
}

var _ SignedAttestation = signedAttestation{}

// InProgressAttestation is an attestation to be processed by offline agent.
type InProgressAttestation interface {
	// SignedAttestation gets the signed attestation
	SignedAttestation() SignedAttestation
	// SubmittedToAttestationCollectorTime is time when signed attestation was submitted to AttestationCollector
	SubmittedToAttestationCollectorTime() *time.Time
	// SubmittedToDestinationTime is time when signed attestation was submitted to the Destination
	SubmittedToDestinationTime() *time.Time
	// AttestationState is the state the in-progress attestation is in right now
	AttestationState() AttestationState
}

// inProgressAttestation is a struct that conforms to InProgressAttestation.
type inProgressAttestation struct {
	signedAttestation                   SignedAttestation
	submittedToAttestationCollectorTime *time.Time
	submittedToDestinationTime          *time.Time
	attestationState                    AttestationState
}

// NewInProgressAttestation creates a new to process attestation.
func NewInProgressAttestation(signedAttestation SignedAttestation, submittedToAttestationCollectorTime *time.Time, state AttestationState) InProgressAttestation {
	return inProgressAttestation{
		signedAttestation:                   signedAttestation,
		submittedToAttestationCollectorTime: submittedToAttestationCollectorTime,
		attestationState:                    state,
	}
}

func (t inProgressAttestation) SignedAttestation() SignedAttestation {
	return t.signedAttestation
}

func (t inProgressAttestation) SubmittedToAttestationCollectorTime() *time.Time {
	return t.submittedToAttestationCollectorTime
}

func (t inProgressAttestation) SubmittedToDestinationTime() *time.Time {
	return t.submittedToDestinationTime
}

func (t inProgressAttestation) AttestationState() AttestationState {
	return t.attestationState
}

var _ InProgressAttestation = inProgressAttestation{}

// NewAttestationKey takes the raw AttestationKey serialized as a big endian big.Int
// and converts it to AttestationKey which is a tuple of (origin, destination, nonce).
func NewAttestationKey(rawKey *big.Int) AttestationKey {
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

// NewAttestationAgentCounts takes the raw AttestationAgentCounts serialized as a big endian big.Int
// and converts it to AttestationAgentCounts which is a tuple of (GuardCount, NotaryCount).
func NewAttestationAgentCounts(rawBytes []byte) AttestationAgentCounts {
	guardCountBytes := rawBytes[attestationAgentCountsGuardCountStartingByte:attestationAgentCountsNotaryCountStartingByte]
	notaryCountBytes := rawBytes[attestationAgentCountsNotaryCountStartingByte:attestationAgentCountsSize]

	guardCount := uint32(guardCountBytes[0])
	notaryCount := uint32(notaryCountBytes[0])

	return AttestationAgentCounts{
		GuardCount:  guardCount,
		NotaryCount: notaryCount,
	}
}

// GetRawAgentCounts returns the AttestationAgentCounts which is a tuple of (GuardCount, NotaryCount)
// as a serialized big.Int.
func (a AttestationAgentCounts) GetRawAgentCounts() []byte {
	guardCountBytes := make([]byte, sizeOfUint8)
	guardCountBytes[0] = uint8(a.GuardCount)
	notaryCountBytes := make([]byte, sizeOfUint8)
	notaryCountBytes[0] = uint8(a.NotaryCount)
	rawBytes := make([]byte, sizeOfUint16)
	copy(rawBytes[attestationAgentCountsGuardCountStartingByte:attestationAgentCountsNotaryCountStartingByte], guardCountBytes)
	copy(rawBytes[attestationAgentCountsNotaryCountStartingByte:attestationAgentCountsSize], notaryCountBytes)
	return rawBytes
}

var _ SignedAttestation = signedAttestation{}
