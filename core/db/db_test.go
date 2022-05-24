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
}
