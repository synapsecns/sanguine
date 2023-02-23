package types_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/types"
)

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

func TestEncodeDecodeAttestation(t *testing.T) {
	origin := gofakeit.Uint32()
	destination := origin + 1
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	attestKey := types.AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
	formattedData, err := types.EncodeAttestation(types.NewAttestation(attestKey.GetRawKey(), root))
	Nil(t, err)

	decodedAttestation, err := types.DecodeAttestation(formattedData)
	Nil(t, err)

	Equal(t, decodedAttestation.Origin(), origin)
	Equal(t, decodedAttestation.Nonce(), nonce)

	rawRoot := decodedAttestation.Root()
	Equal(t, common.BytesToHash(rawRoot[:]), root)
}

func TestEncodeDecodeSignedAttestation(t *testing.T) {
	domain := gofakeit.Uint32()
	destination := domain + 1
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	fakeGuardV1 := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeGuardR1 := big.NewInt(gofakeit.Int64())
	fakeGuardS1 := big.NewInt(gofakeit.Int64())
	fakeGuardSig1 := types.NewSignature(fakeGuardV1, fakeGuardR1, fakeGuardS1)

	fakeGuardV2 := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeGuardR2 := big.NewInt(gofakeit.Int64())
	fakeGuardS2 := big.NewInt(gofakeit.Int64())
	fakeGuardSig2 := types.NewSignature(fakeGuardV2, fakeGuardR2, fakeGuardS2)

	fakeGuardV3 := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeGuardR3 := big.NewInt(gofakeit.Int64())
	fakeGuardS3 := big.NewInt(gofakeit.Int64())
	fakeGuardSig3 := types.NewSignature(fakeGuardV3, fakeGuardR3, fakeGuardS3)

	fakeNotaryV1 := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeNotaryR1 := big.NewInt(gofakeit.Int64())
	fakeNotaryS1 := big.NewInt(gofakeit.Int64())
	fakeNotarySig1 := types.NewSignature(fakeNotaryV1, fakeNotaryR1, fakeNotaryS1)

	fakeNotaryV2 := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeNotaryR2 := big.NewInt(gofakeit.Int64())
	fakeNotaryS2 := big.NewInt(gofakeit.Int64())
	fakeNotarySig2 := types.NewSignature(fakeNotaryV2, fakeNotaryR2, fakeNotaryS2)

	attestKey := types.AttestationKey{
		Origin:      domain,
		Destination: destination,
		Nonce:       nonce,
	}
	signedAttestation := types.NewSignedAttestation(
		types.NewAttestation(attestKey.GetRawKey(), root),
		[]types.Signature{fakeGuardSig1, fakeGuardSig2, fakeGuardSig3},
		[]types.Signature{fakeNotarySig1, fakeNotarySig2},
	)

	encoded, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(t, err)

	decoded, err := types.DecodeSignedAttestation(encoded)
	Nil(t, err)

	Equal(t, decoded.Attestation().Root(), signedAttestation.Attestation().Root())
	Equal(t, decoded.Attestation().Origin(), signedAttestation.Attestation().Origin())
	Equal(t, decoded.Attestation().Destination(), signedAttestation.Attestation().Destination())
	Equal(t, decoded.Attestation().Nonce(), signedAttestation.Attestation().Nonce())

	Equal(t, len(decoded.GuardSignatures()), len(signedAttestation.GuardSignatures()))
	Equal(t, len(decoded.NotarySignatures()), len(signedAttestation.NotarySignatures()))

	for i := 0; i < len(decoded.GuardSignatures()); i++ {
		Equal(t, decoded.GuardSignatures()[i].V(), signedAttestation.GuardSignatures()[i].V())
		Equal(t, decoded.GuardSignatures()[i].R(), signedAttestation.GuardSignatures()[i].R())
		Equal(t, decoded.GuardSignatures()[i].S(), signedAttestation.GuardSignatures()[i].S())
	}

	for i := 0; i < len(decoded.NotarySignatures()); i++ {
		Equal(t, decoded.NotarySignatures()[i].V(), signedAttestation.NotarySignatures()[i].V())
		Equal(t, decoded.NotarySignatures()[i].R(), signedAttestation.NotarySignatures()[i].R())
		Equal(t, decoded.NotarySignatures()[i].S(), signedAttestation.NotarySignatures()[i].S())
	}
}

func TestEncodeDecodeTips(t *testing.T) {
	// we want to make sure we can deal w/ overflows
	notaryTip := randomUint96BigInt(t)
	broadcasterTip := randomUint96BigInt(t)
	proverTip := randomUint96BigInt(t)
	executorTip := randomUint96BigInt(t)

	encodedTips, err := types.EncodeTips(types.NewTips(notaryTip, broadcasterTip, proverTip, executorTip))
	Nil(t, err)

	decodedTips, err := types.DecodeTips(encodedTips)
	Nil(t, err)

	Equal(t, decodedTips.NotaryTip(), notaryTip)
	Equal(t, decodedTips.BroadcasterTip(), broadcasterTip)
	Equal(t, decodedTips.ProverTip(), proverTip)
	Equal(t, decodedTips.ExecutorTip(), executorTip)
}

func TestNewMessageEncodeDecode(t *testing.T) {
	origin := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	body := []byte(gofakeit.Sentence(gofakeit.Number(5, 15)))
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
	optimisticSeconds := gofakeit.Uint32()

	header := types.NewHeader(
		origin, sender, nonce, destination, recipient, optimisticSeconds)

	notaryTip := randomUint96BigInt(t)
	broadcasterTip := randomUint96BigInt(t)
	proverTip := randomUint96BigInt(t)
	executorTip := randomUint96BigInt(t)

	tips := types.NewTips(notaryTip, broadcasterTip, proverTip, executorTip)

	newMessage := types.NewMessage(header, tips, body)

	Equal(t, newMessage.OriginDomain(), origin)
	Equal(t, newMessage.Sender(), sender)
	Equal(t, newMessage.Nonce(), nonce)
	Equal(t, newMessage.DestinationDomain(), destination)
	Equal(t, newMessage.Body(), body)

	encodedMessage, err := types.EncodeMessage(newMessage)
	Nil(t, err)

	// make sure decode is same as encode
	decodedMessage, err := types.DecodeMessage(encodedMessage)
	Nil(t, err)

	Equal(t, newMessage.OriginDomain(), decodedMessage.OriginDomain())
	Equal(t, newMessage.Sender(), decodedMessage.Sender())
	Equal(t, newMessage.Nonce(), decodedMessage.Nonce())
	Equal(t, newMessage.DestinationDomain(), decodedMessage.DestinationDomain())
	Equal(t, newMessage.Body(), decodedMessage.Body())
}

func TestHeaderEncodeDecode(t *testing.T) {
	origin := gofakeit.Uint32()
	sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
	nonce := gofakeit.Uint32()
	destination := gofakeit.Uint32()
	recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
	optimisticSeconds := gofakeit.Uint32()

	ogHeader, err := types.EncodeHeader(types.NewHeader(origin, sender, nonce, destination, recipient, optimisticSeconds))
	Nil(t, err)

	decodedHeader, err := types.DecodeHeader(ogHeader)
	Nil(t, err)

	Equal(t, decodedHeader.OriginDomain(), origin)
	Equal(t, decodedHeader.Sender(), sender)
	Equal(t, decodedHeader.Nonce(), nonce)
	Equal(t, decodedHeader.DestinationDomain(), destination)
	Equal(t, decodedHeader.Recipient(), recipient)
	Equal(t, decodedHeader.OptimisticSeconds(), optimisticSeconds)
}

func TestDecodeAttestation(t *testing.T) {
	rawAttestation, err := hex.DecodeString("00000089000000890000000027ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757")
	Nil(t, err)

	decodedAttestation, err := types.DecodeAttestation(rawAttestation)
	Nil(t, err)
	NotNil(t, decodedAttestation)

	/*rootBytes32 := decodedAttestation.Root()
	rootStr := hex.EncodeToString(rootBytes32[:])
	fmt.Printf("\nCRONIN\n origin: %d,\n destination: %d,\n nonce: %d,\n root: %s\n \nCRONIN\n",
		decodedAttestation.Origin(), decodedAttestation.Destination(), decodedAttestation.Nonce(), rootStr)*/
}
