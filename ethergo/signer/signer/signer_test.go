package signer_test

import (
	"github.com/brianvoe/gofakeit/v6"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"math/big"
	"testing"
)

func TestEncodeSigner(t *testing.T) {
	v := new(big.Int).SetUint64(gofakeit.Uint64())
	r := new(big.Int).SetUint64(gofakeit.Uint64())
	s := new(big.Int).SetUint64(gofakeit.Uint64())

	sig := types.NewSignature(v, r, s)

	rawSig := signer.Encode(sig)

	exampleSigner := ethTypes.FrontierSigner{}

	newR, newS, _, err := exampleSigner.SignatureValues(ethTypes.NewTx(&ethTypes.LegacyTx{
		V: v,
		R: r,
		S: s,
	}), rawSig)

	Nil(t, err)

	// skip v, this gets modified by eth
	Equal(t, s, newS)
	Equal(t, r, newR)
}
