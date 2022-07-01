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

		encoded, err := realMessage.Encode()
		Nil(m.T(), err)

		committedMessage := types.NewCommittedMessage(100, common.BigToHash(big.NewInt(gofakeit.Int64())), encoded)

		realLeaf, err := realMessage.ToLeaf()
		Nil(m.T(), err)

		Equal(m.T(), realLeaf, committedMessage.Leaf())

		err = newDB.StoreCommittedMessage(committedMessage)
		Nil(m.T(), err)

		// try by nonce
		byNonce, err := newDB.MessageByNonce(realMessage.Destination(), realMessage.Nonce())
		Nil(m.T(), err)

		Equal(m.T(), byNonce.Message(), encoded)
		Equal(m.T(), byNonce.CommitedRoot(), committedMessage.CommitedRoot())
		Equal(m.T(), byNonce.Leaf(), committedMessage.Leaf())
		Equal(m.T(), byNonce.LeafIndex(), committedMessage.LeafIndex())

		// try by leaf
		byLeaf, err := newDB.MessageByLeaf(realLeaf)
		Nil(m.T(), err)

		Equal(m.T(), byNonce.Message(), byLeaf.Message())
		Equal(m.T(), byNonce.CommitedRoot(), byLeaf.CommitedRoot())
		Equal(m.T(), byNonce.Leaf(), byLeaf.Leaf())
		Equal(m.T(), byNonce.LeafIndex(), byLeaf.LeafIndex())

		// try by leaf index
		byLeafIndex, err := newDB.MessageByLeafIndex(byNonce.LeafIndex())
		Nil(m.T(), err)

		Equal(m.T(), byLeafIndex.Message(), byNonce.Message())
	})
}

func (m *MessageSuite) TestStoresAndRetreivesProofs() {
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

		retreivedRoot, err := newDB.RetrieveLatestRoot()
		Nil(m.T(), err)

		Equal(m.T(), retreivedRoot, latestRoot)
	})
}

func (m *MessageSuite) TestStoreGetProducedUpdate() {
	m.T().Skip("TODO:  teststore/retrieve produced update")
}
