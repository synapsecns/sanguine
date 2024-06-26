package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"gorm.io/gorm"
)

func (d *DBSuite) TestBlacklist() {
	d.RunOnAllDBs(func(testDB db.DB) {
		testAddress := gofakeit.BitcoinAddress()

		blacklistBody := db.BlacklistedAddress{
			Type: "create",
			ID:   "testId",
			Data: db.Data{
				Address: testAddress,
				Network: "bitcoin",
				Tag:     "testTag",
				Remark:  "testRemark",
			},
		}

		// blacklist the address
		err := testDB.PutBlacklistedAddress(d.GetTestContext(), blacklistBody)
		d.Require().NoError(err)
		blacklistedAddress, err := testDB.GetBlacklistedAddress(d.GetTestContext(), blacklistBody.Data.Address)
		d.Require().NoError(err)
		d.Require().NotNil(blacklistedAddress)

		// update the address
		blacklistBody.Type = "update"
		blacklistBody.Data.Remark = "testRemarkUpdated"
		err = testDB.UpdateBlacklistedAddress(d.GetTestContext(), blacklistBody.ID, blacklistBody)
		d.Require().NoError(err)

		// check to make sure it updated
		blacklistedAddress, err = testDB.GetBlacklistedAddress(d.GetTestContext(), blacklistBody.Data.Address)
		d.Require().NoError(err)
		d.Require().NotNil(blacklistedAddress)
		d.Require().Equal("testRemarkUpdated", blacklistedAddress.Data.Remark)

		// check for non blacklisted address
		res, err := testDB.GetBlacklistedAddress(d.GetTestContext(), gofakeit.BitcoinAddress())
		d.Require().EqualError(err, gorm.ErrRecordNotFound.Error())
		d.Require().Nil(res)

		// delete it
		err = testDB.DeleteBlacklistedAddress(d.GetTestContext(), blacklistBody.ID)
		d.Require().NoError(err)

		// delete nonexistent
		err = testDB.DeleteBlacklistedAddress(d.GetTestContext(), "NonexistentId")
		d.Require().NoError(err)
	})
}
