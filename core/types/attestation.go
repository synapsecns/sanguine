package types

// Attestation is the attestation.
type Attestation interface {
	// Domain gets the domain of the attestation
	Domain() uint32
	// Nonce gets the nonce of the attestation
	Nonce() uint32
	// Root gets the root of the contract
	Root() [32]byte
}

type attestation struct {
	// domain of the attestation
	domain uint32
	// nonce of the attestation
	nonce uint32
	// root - the merkle root
	root [32]byte
}

// NewAttestation creates a new attesation.
func NewAttestation(localDomain, nonce uint32, root [32]byte) Attestation {
	return attestation{
		domain: localDomain,
		nonce:  nonce,
		root:   root,
	}
}

func (a attestation) Domain() uint32 {
	return a.domain
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
