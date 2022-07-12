package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

const homeDomain = "SYN"

// HomeDomainHash creates a home domain hash.
func HomeDomainHash(domain uint32) ([32]byte, error) {
	buf := new(bytes.Buffer)

	// pack the non-byte parts of the struct
	type packedDomain struct {
		Domain uint32
	}

	basePack := packedDomain{
		Domain: domain,
	}

	err := binary.Write(buf, binary.BigEndian, basePack)
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not write domain: %w", err)
	}

	packed := append(buf.Bytes(), []byte(homeDomain)...)
	return crypto.Keccak256Hash(packed), nil
}
