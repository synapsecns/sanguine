package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
)

// StoresAndRetrievesMessages tests storage/retreival.
func (m *MessageSuite) TestStoresAndRetrievesMessages() {
	m.RunOnAllDBs(func(newDB db.MessageDB) {
		realMessage := types.NewMessage(10, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32(), gofakeit.Uint32(), []byte(gofakeit.Sentence(10)), common.BigToHash(big.NewInt(gofakeit.Int64())))

		encoded, err := types.EncodeMessage(realMessage)
		Nil(m.T(), err)

		committedMessage := types.NewCommittedMessage(100, encoded)

		realLeaf, err := realMessage.ToLeaf()
		Nil(m.T(), err)

		Equal(m.T(), realLeaf, committedMessage.Leaf())

		err = newDB.StoreCommittedMessage(committedMessage)
		Nil(m.T(), err)

		// try by nonce
		byNonce, err := newDB.MessageByNonce(realMessage.Destination(), realMessage.Nonce())
		Nil(m.T(), err)

		Equal(m.T(), byNonce.Message(), encoded)
		Equal(m.T(), byNonce.Leaf(), committedMessage.Leaf())
		Equal(m.T(), byNonce.LeafIndex(), committedMessage.LeafIndex())

		// try by leaf
		byLeaf, err := newDB.MessageByLeaf(realLeaf)
		Nil(m.T(), err)

		Equal(m.T(), byNonce.Message(), byLeaf.Message())
		Equal(m.T(), byNonce.Leaf(), byLeaf.Leaf())
		Equal(m.T(), byNonce.LeafIndex(), byLeaf.LeafIndex())

		// try by leaf index
		byLeafIndex, err := newDB.MessageByLeafIndex(byNonce.LeafIndex())
		Nil(m.T(), err)

		Equal(m.T(), byLeafIndex.Message(), byNonce.Message())
	})
}

func (m *MessageSuite) TestStoresAndretrievesProofs() {
	m.RunOnAllDBs(func(newDB db.MessageDB) {
		leaf := common.BigToHash(big.NewInt(gofakeit.Int64()))
		index := gofakeit.Uint32()
		path := common.Hash{}

		proof := types.NewProof(leaf, index, path)

		err := newDB.StoreProof(13, proof)
		Nil(m.T(), err)

		byIndex, err := newDB.ProofByLeafIndex(13)
		Nil(m.T(), err)

		Equal(m.T(), byIndex.Index(), index)
		Equal(m.T(), byIndex.Path(), path)
		Equal(m.T(), byIndex.Leaf(), leaf)
	})
}

func (m *MessageSuite) TestStoreGetMessageLatestBlockEnd() {
	m.RunOnAllDBs(func(newDB db.MessageDB) {
		_, err := newDB.GetMessageLatestBlockEnd()
		Error(m.T(), err, pebble.ErrNotFound)

		fakeBlock := gofakeit.Uint32()

		err = newDB.StoreMessageLatestBlockEnd(fakeBlock)
		Nil(m.T(), err)

		latestHeight, err := newDB.GetMessageLatestBlockEnd()
		Nil(m.T(), err)

		Equal(m.T(), latestHeight, fakeBlock)
	})
}

func (m *MessageSuite) TestStoreAndRetrieveLatestRoot() {
	m.RunOnAllDBs(func(newDB db.MessageDB) {
		_, err := newDB.RetrieveLatestRoot()
		Error(m.T(), err, pebble.ErrNotFound)

		latestRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))

		err = newDB.StoreLatestRoot(latestRoot)
		Nil(m.T(), err)

		retrievedRoot, err := newDB.RetrieveLatestRoot()
		Nil(m.T(), err)

		Equal(m.T(), retrievedRoot, latestRoot)
	})
}

func (m *MessageSuite) TestStoreGetProducedUpdate() {
	m.T().Skip("TODO:  teststore/retrieve produced update")
}

func (t *DBSuite) TestRetrieveLatestNonce() {
	const domainID = 1

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		_, err := testDB.RetrieveLatestNonce(t.GetTestContext(), domainID)
		ErrorIs(t.T(), err, db.ErrNoNonceForDomain)

		nonce := 0
		leafIndex := uint32(1)

		for i := 0; i < 10; i++ {
			realMessage := types.NewMessage(10, common.BigToHash(big.NewInt(gofakeit.Int64())), uint32(nonce), gofakeit.Uint32(), []byte(gofakeit.Sentence(10)), common.BigToHash(big.NewInt(gofakeit.Int64())))

			encoded, err := types.EncodeMessage(realMessage)
			Nil(t.T(), err)

			err = testDB.StoreCommittedMessage(t.GetTestContext(), domainID, types.NewCommittedMessage(leafIndex, encoded))
			Nil(t.T(), err)

			newNonce, err := testDB.RetrieveLatestNonce(t.GetTestContext(), domainID)
			Equal(t.T(), uint32(nonce), newNonce)

			nonce += 1
			leafIndex += 1
		}

	})
}
