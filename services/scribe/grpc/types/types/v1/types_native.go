package pbscribe

import "github.com/ethereum/go-ethereum/common"

// FromNativeAddress converts a native address to a generated address.
func FromNativeAddress(address common.Address) *Address {
	return &Address{Bytes: address.Bytes()}
}

// ToAddress converts a byte to an address.
func (x *Address) ToAddress() common.Address {
	return common.BytesToAddress(x.GetBytes())
}

// FromNativeHash converts an eth hash to a bytes hash.
func FromNativeHash(nativeHash common.Hash) *Hash {
	return &Hash{Bytes: nativeHash.Bytes()}
}

// ToHash converts a native hash to a hash list.
func (x *Hash) ToHash() common.Hash {
	return common.BytesToHash(x.GetBytes())
}

// FromNativeHashes is a helper function for converting a slice of native hashes to a hash array.
func FromNativeHashes(nativeHashes []common.Hash) (castHashes []*Hash) {
	for _, hash := range nativeHashes {
		castHashes = append(castHashes, FromNativeHash(hash))
	}
	return castHashes
}

// ToNativeHashes is a helper function for converting a slice of hashes to a native hash array.
func ToNativeHashes(hashes []*Hash) (nativeHashes []common.Hash) {
	for _, hash := range hashes {
		nativeHashes = append(nativeHashes, hash.ToHash())
	}
	return nativeHashes
}
