package types_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
	"testing"
)

func TestEncodeSignedUpdate(t *testing.T) {
	fakeDomain := gofakeit.Uint32()
	previousRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nextRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))

	fakeV := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeR := new(big.Int).SetUint64(gofakeit.Uint64())
	fakeS := new(big.Int).SetUint64(gofakeit.Uint64())

	signature := types.NewSignature(fakeV, fakeR, fakeS)
	update := types.NewUpdate(fakeDomain, previousRoot, nextRoot)

	signedUpdate := types.NewSignedUpdate(update, signature)

	res, err := types.EncodeSignedUpdate(signedUpdate)
	Nil(t, err)

	signedUp, err := types.DecodeSignedUpdate(res)
	Nil(t, err)

	Equal(t, signedUp.Update().NewRoot(), nextRoot)
	Equal(t, signedUp.Update().PreviousRoot(), previousRoot)
	Equal(t, signedUp.Update().HomeDomain(), fakeDomain)

	Equal(t, signedUp.Signature().S(), fakeS)
	Equal(t, signedUp.Signature().R(), fakeR)
	Equal(t, signedUp.Signature().V(), fakeV)
}

func TestEncodeDecodeSignature(t *testing.T) {
	fakeV := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeR := big.NewInt(gofakeit.Int64())
	fakeS := big.NewInt(gofakeit.Int64())

	sig := types.NewSignature(fakeV, fakeR, fakeS)
	res, err := types.EncodeSignature(sig)
	Nil(t, err)

	decodedSig, err := types.DecodeSignature(res)
	Nil(t, err)

	Equal(t, sig.V(), decodedSig.V())
	Equal(t, sig.R(), decodedSig.R())
	Equal(t, sig.S(), decodedSig.S())
}

func TestEncodeDecodeUpdate(t *testing.T) {
	fakeDomain := gofakeit.Uint32()
	previousRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nextRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))

	update := types.NewUpdate(fakeDomain, previousRoot, nextRoot)

	res, err := types.EncodeUpdate(update)
	Nil(t, err)

	_, err = types.DecodeUpdate(res)
	Nil(t, err)
}
