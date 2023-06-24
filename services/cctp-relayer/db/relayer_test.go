package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db"
	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
)

func (d *DBSuite) mockMessage(originChainID, destinationChainID, blockNumber uint32) types.Message {
	return types.Message{
		OriginTxHash:     mocks.NewMockHash(d.T()).String(),
		DestTxHash:       mocks.NewMockHash(d.T()).String(),
		OriginChainID:    originChainID,
		DestChainID:      destinationChainID,
		Message:          []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		MessageHash:      mocks.NewMockHash(d.T()).String(),
		Attestation:      []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		RequestVersion:   0,
		FormattedRequest: []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		BlockNumber:      uint64(blockNumber),
		State:            types.Pending,
	}
}

func (d *DBSuite) TestLastBlockNumber() {
	d.RunOnAllDBs(func(testDB db.CCTPRelayerDB) {
		originChainID := gofakeit.Uint32()
		destChainID := gofakeit.Uint32()

		// make sure w/ no messages inserted the last block number is 0
		lastBlockNumber, err := testDB.GetLastBlockNumber(d.GetTestContext(), originChainID)
		d.Nil(err)

		d.Equal(uint64(0), lastBlockNumber)

		// insert a message for origin chain id. Make sure the last block number matches for the origin chain
		// and not the dest chain.
		newBlockNumber := gofakeit.Uint32()
		message := d.mockMessage(originChainID, destChainID, newBlockNumber)

		err = testDB.StoreMessage(d.GetTestContext(), message)
		d.Nil(err)

		lastBlockNumber, err = testDB.GetLastBlockNumber(d.GetTestContext(), originChainID)
		d.Nil(err)
		d.Equal(uint64(newBlockNumber), lastBlockNumber)

		destBlockNumber, err := testDB.GetLastBlockNumber(d.GetTestContext(), destChainID)
		d.Nil(err)

		d.Equal(uint64(0), destBlockNumber)
	})
}

// TestUpsertStoreMessage asserts that the messages are correctly stored in the database
// regardless of what state they are in.
func (d *DBSuite) TestUpsertStoreMessage() {
	d.RunOnAllDBs(func(testDB db.CCTPRelayerDB) {
		message := d.mockMessage(gofakeit.Uint32(), gofakeit.Uint32(), gofakeit.Uint32())
		// first insert w/ no error
		message.Attestation = nil
		message.State = types.Pending

		err := testDB.StoreMessage(d.GetTestContext(), message)
		d.Nil(err)

		d.T().Skip("TODO: finish test, this needs better getters to make sure it's working correctly")
	})
}
