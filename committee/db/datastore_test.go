package db_test

import (
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/synapsecns/sanguine/committee/db"
)

// This test along with several others are adaptation copypasta from ds_test.go. This was needed to test our mysql driver without pushing it upstream.
// why not push upstream? because we rely on an orm and we wouldn't want to force others to have to download gorm
// to utilize ds-test especially when it may conflict w/ their own versioning considering go's lack of a peer dependency concept.
var testcases = map[string]string{
	"/a":     "a",
	"/a/b":   "ab",
	"/a/b/c": "abc",
	"/a/b/d": "a/b/d",
	"/a/c":   "ac",
	"/a/d":   "ad",
	"/e":     "e",
	"/f":     "f",
	"/g":     "",
}

func (d *DBSuite) addTestCases(ds datastore.Batching) {
	d.T().Helper()

	for k, v := range testcases {
		dsk := datastore.NewKey(k)
		d.NoError(ds.Put(d.GetTestContext(), dsk, []byte(v)))
	}

	for k, v := range testcases {
		dsk := datastore.NewKey(k)
		value, err := ds.Get(d.GetTestContext(), dsk)
		d.NoError(err)
		d.Equal([]byte(v), value)
	}
}

func (d *DBSuite) expectMatches(expect []string, actualR query.Results) {
	actual, err := actualR.Rest()
	if err != nil {
		d.T().Error(err)
	}

	if len(actual) != len(expect) {
		d.T().Error("not enough", expect, actual)
	}
	for _, k := range expect {
		found := false
		for _, e := range actual {
			if e.Key == k {
				found = true
			}
		}
		if !found {
			d.T().Error(k, "not found")
		}
	}
}

func (d *DBSuite) TestQuery() {
	d.RunOnAllDatastores(func(ds datastore.Batching) {
		d.addTestCases(ds)

		rs, err := ds.Query(d.GetTestContext(), query.Query{Prefix: "/a/"})
		d.NoError(err)

		d.expectMatches([]string{
			"/a/b",
			"/a/b/c",
			"/a/b/d",
			"/a/c",
			"/a/d",
		}, rs)

		// test offset and limit

		rs, err = ds.Query(d.GetTestContext(), query.Query{Prefix: "/a/", Offset: 2, Limit: 2})
		if err != nil {
			d.T().Fatal(err)
		}

		d.expectMatches([]string{
			"/a/b/d",
			"/a/c",
		}, rs)
	})
}

func (d *DBSuite) TestDBAdapters() {
	d.RunOnAllDBs(func(testDB db.Service) {
		ds, err := testDB.GlobalDatastore()
		d.NoError(err)

		err = ds.Put(d.GetTestContext(), datastore.NewKey("key"), []byte("value"))
		d.NoError(err)

		value, err := ds.Get(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)
		d.Equal([]byte("value"), value)

		var has bool
		has, err = ds.Has(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)
		d.True(has)

		_, err = ds.GetSize(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)

		err = ds.Delete(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)

		has, err = ds.Has(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)
		d.False(has)
	})
}
