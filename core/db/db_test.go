package db_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
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
