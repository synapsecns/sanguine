package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/types"
)

func (t *DBSuite) TestRetrieveLatestNonce() {
	const domainID = 1

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		_, err := testDB.RetrieveLatestCommittedMessageNonce(t.GetTestContext(), domainID)
		ErrorIs(t.T(), err, db.ErrNoNonceForDomain)

		nonce := 0
		leafIndex := uint32(1)

		for i := 0; i < 10; i++ {
			realTips := types.NewTips(big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
			realHeader := types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), uint32(nonce), gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
			realMessage := types.NewMessage(realHeader, realTips, []byte(gofakeit.Sentence(10)))

			encoded, err := types.EncodeMessage(realMessage)
			Nil(t.T(), err)

			err = testDB.StoreCommittedMessage(t.GetTestContext(), domainID, types.NewCommittedMessage(leafIndex, encoded))
			Nil(t.T(), err)

			newNonce, err := testDB.RetrieveLatestCommittedMessageNonce(t.GetTestContext(), domainID)
			Nil(t.T(), err)
			Equal(t.T(), uint32(nonce), newNonce)

			nonce++
			leafIndex++
		}
	})
}
