package core

// BytesToSlice converts a 32 bit array to a slice slice.
func BytesToSlice(bytes [32]byte) []byte {
	rawBytes := make([]byte, len(bytes))
	copy(rawBytes, bytes[:])
	return rawBytes
}
