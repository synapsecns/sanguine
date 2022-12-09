package tokendata

// ImmutableTokenData contains immutable token data (decimals and tokenID)
// as long as the token ID of an address does not change, this data is guaranteed to be accurate.
type ImmutableTokenData interface {
	// Decimals gets the number of decimals for a token
	Decimals() uint8
	// TokenID gets the tokenID
	TokenID() string
}

type immutableTokenImpl struct {
	// decimals contains the number of decimals for the token
	decimals uint8
	// tokenID is the token ID of the token
	tokenID string
}

func (i immutableTokenImpl) Decimals() uint8 {
	return i.decimals
}

// TokenID gets the tokenID.
func (i immutableTokenImpl) TokenID() string {
	return i.tokenID
}
