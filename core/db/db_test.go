package db_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
)

// StoresAndRetrievesMessages tests storage/retreival.
func (d *DBSuite) TestStoresAndRetrievesMessages() {
	newDB, err := db.NewDB(filet.TmpDir(d.T(), ""), "home1")
	Nil(d.T(), err)

	realMessage := types.NewMessage(10, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32(), gofakeit.Uint32(), []byte(gofakeit.Sentence(10)), common.BigToHash(big.NewInt(gofakeit.Int64())))

	encoded, err := realMessage.Encode()
	Nil(d.T(), err)

	committedMessage := types.NewCommittedMessage(100, common.BigToHash(big.NewInt(gofakeit.Int64())), encoded)

	realLeaf, err := realMessage.ToLeaf()
	Nil(d.T(), err)

	Equal(d.T(), realLeaf, committedMessage.Leaf())

	err = newDB.StoreCommittedMessage(committedMessage)
	Nil(d.T(), err)

	// try by nonce
	byNonce, err := newDB.MessageByNonce(realMessage.Destination(), realMessage.Nonce())
	Nil(d.T(), err)

	Equal(d.T(), byNonce.Message(), encoded)
	Equal(d.T(), byNonce.CommitedRoot(), committedMessage.CommitedRoot())
	Equal(d.T(), byNonce.Leaf(), committedMessage.Leaf())
	Equal(d.T(), byNonce.LeafIndex(), committedMessage.LeafIndex())

	// try by leaf
	byLeaf, err := newDB.MessageByLeaf(realLeaf)
	Nil(d.T(), err)

	Equal(d.T(), byNonce.Message(), byLeaf.Message())
	Equal(d.T(), byNonce.CommitedRoot(), byLeaf.CommitedRoot())
	Equal(d.T(), byNonce.Leaf(), byLeaf.Leaf())
	Equal(d.T(), byNonce.LeafIndex(), byLeaf.LeafIndex())

	// try by leaf index
	byLeafIndex, err := newDB.MessageByLeafIndex(byNonce.LeafIndex())
	Nil(d.T(), err)

	Equal(d.T(), byLeafIndex.Message(), byNonce.Message())
}

func (d *DBSuite) TestStoresAndRetreivesProofs() {
	newDB, err := db.NewDB(filet.TmpDir(d.T(), ""), "home1")
	Nil(d.T(), err)

	leaf := common.BigToHash(big.NewInt(gofakeit.Int64()))
	index := gofakeit.Uint32()
	path := common.Hash{}

	proof := types.NewProof(leaf, index, path)

	err = newDB.StoreProof(13, proof)
	Nil(d.T(), err)

	byIndex, err := newDB.ProofByLeafIndex(13)
	Nil(d.T(), err)

	Equal(d.T(), byIndex.Index(), index)
	Equal(d.T(), byIndex.Path(), path)
	Equal(d.T(), byIndex.Leaf(), leaf)
}

func (d *DBSuite) TestStoreGetMessageLatestBlockEnd() {
	newDB, err := db.NewDB(filet.TmpDir(d.T(), ""), "home1")
	Nil(d.T(), err)

	_, err = newDB.GetMessageLatestBlockEnd()
	Error(d.T(), err, pebble.ErrNotFound)

	fakeBlock := gofakeit.Uint32()

	err = newDB.StoreMessageLatestBlockEnd(fakeBlock)
	Nil(d.T(), err)

	latestHeight, err := newDB.GetMessageLatestBlockEnd()
	Nil(d.T(), err)

	Equal(d.T(), latestHeight, fakeBlock)
}

func (d *DBSuite) TestStoreAndRetrieveLatestRoot() {
	newDB, err := db.NewDB(filet.TmpDir(d.T(), ""), "home1")
	Nil(d.T(), err)

	_, err = newDB.RetrieveLatestRoot()
	Error(d.T(), err, pebble.ErrNotFound)

	latestRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))

	err = newDB.StoreLatestRoot(latestRoot)
	Nil(d.T(), err)

	retreivedRoot, err := newDB.RetrieveLatestRoot()
	Nil(d.T(), err)

	Equal(d.T(), retreivedRoot, latestRoot)
}

func (d *DBSuite) TestStoreGetProducedUpdate() {
	d.T().Skip("TODO:  teststore/retrieve produced update")
}
