package db_test

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"gorm.io/gorm"
)

func (d *DBSuite) TestEmpty() {
	d.RunOnAllDBs(func(testDB db.DB) {
		testAddress := gofakeit.BitcoinAddress()

		// 5 mins ago
		_, err := testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now().Add(time.Minute*-5))
		d.Require().Error(err, db.ErrNoAddressNotCached)

		err = testDB.PutAddressIndicators(d.GetTestContext(), testAddress, []trmlabs.AddressRiskIndicator{})
		d.Require().NoError(err)

		// 5 mins ago
		_, err = testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now().Add(time.Minute*-5))
		d.Require().NoError(err)

		// also make sure expiry works correctly, this should error
		_, err = testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now())
		d.Require().Error(err, db.ErrNoAddressNotCached)
	})
}

func (d *DBSuite) TestAddressUpdate() {
	d.RunOnAllDBs(func(testDB db.DB) {
		testAddress := gofakeit.BitcoinAddress()

		// 5 mins ago
		_, err := testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now().Add(time.Minute*-5))
		d.Require().Error(err, db.ErrNoAddressNotCached)

		err = testDB.PutAddressIndicators(d.GetTestContext(), testAddress, []trmlabs.AddressRiskIndicator{})
		d.Require().NoError(err)

		// 5 mins ago
		_, err = testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now().Add(time.Minute*-5))
		d.Require().NoError(err)

		// also make sure expiry works correctly, this should error
		_, err = testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now())
		d.Require().Error(err, db.ErrNoAddressNotCached)

		// update the address
		err = testDB.PutAddressIndicators(d.GetTestContext(), testAddress, []trmlabs.AddressRiskIndicator{
			{
				Category: "test",
			},
		})
		d.Require().NoError(err)

		// 5 mins ago
		_, err = testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now().Add(time.Minute*-5))
		d.Require().NoError(err)

		// also make sure expiry works correctly, this should error
		_, err = testDB.GetAddressIndicators(d.GetTestContext(), testAddress, time.Now())
		d.Require().Error(err, db.ErrNoAddressNotCached)
	})
}

func (d *DBSuite) TestBlacklist() {
	d.RunOnAllDBs(func(testDB db.DB) {

		testAddress := gofakeit.BitcoinAddress()

		blacklistBody := db.BlacklistedAddress{
			TypeReq: "create",
			Id:      "testId",
			Address: testAddress,
			Network: "bitcoin",
			Tag:     "testTag",
			Remark:  "testRemark",
		}

		// blacklist the address
		err := testDB.PutBlacklistedAddress(d.GetTestContext(), blacklistBody)
		d.Require().NoError(err)
		blacklistedAddress, err := testDB.GetBlacklistedAddress(d.GetTestContext(), blacklistBody.Address)
		d.Require().NoError(err)
		d.Require().NotNil(blacklistedAddress)

		// update the address
		blacklistBody.TypeReq = "update"
		blacklistBody.Remark = "testRemarkUpdated"
		err = testDB.UpdateBlacklistedAddress(d.GetTestContext(), blacklistBody.Id, blacklistBody)
		d.Require().NoError(err)

		// check to make sure it updated
		blacklistedAddress, err = testDB.GetBlacklistedAddress(d.GetTestContext(), blacklistBody.Address)
		d.Require().NoError(err)
		d.Require().NotNil(blacklistedAddress)
		d.Require().Equal("testRemarkUpdated", blacklistedAddress.Remark)

		// check for non blacklisted address
		res, err := testDB.GetBlacklistedAddress(d.GetTestContext(), gofakeit.BitcoinAddress())
		d.Require().EqualError(err, gorm.ErrRecordNotFound.Error())
		d.Require().Nil(res)

		// delete it
		err = testDB.DeleteBlacklistedAddress(d.GetTestContext(), blacklistBody.Id)
		d.Require().NoError(err)

		// delete nonexistent
		err = testDB.DeleteBlacklistedAddress(d.GetTestContext(), "NonexistentId")
		d.Require().Error(err)

	})
}
