package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"time"
)

func (d *DBSuite) TestEmpty() {
	d.RunOnAllDBs(func(testDB db.RuleDB) {
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

func (d *DBSuite) TestAdressUpdate() {
	d.RunOnAllDBs(func(testDB db.RuleDB) {
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
