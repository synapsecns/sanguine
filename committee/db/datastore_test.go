package db_test

import (
	"errors"
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

func (d *DBSuite) TestHas() {
	d.RunOnAllDatastores(func(ds datastore.Batching) {
		d.addTestCases(ds)

		has, err := ds.Has(d.GetTestContext(), datastore.NewKey("/a/b/c"))
		if err != nil {
			d.T().Error(err)
		}

		if !has {
			d.T().Error("Key should be found")
		}

		has, err = ds.Has(d.GetTestContext(), datastore.NewKey("/a/b/c/d"))
		if err != nil {
			d.T().Error(err)
		}

		if has {
			d.T().Error("Key should not be found")
		}
	})
}

func (d *DBSuite) TestGetSize() {
	d.RunOnAllDatastores(func(testStore datastore.Batching) {
		d.addTestCases(testStore)

		size, err := testStore.GetSize(d.GetTestContext(), datastore.NewKey("/a/b/c"))
		if err != nil {
			d.T().Error(err)
		}

		if size != len(testcases["/a/b/c"]) {
			d.T().Error("")
		}

		_, err = testStore.GetSize(d.GetTestContext(), datastore.NewKey("/a/b/c/d"))
		if !errors.Is(err, datastore.ErrNotFound) {
			d.T().Error(err)
		}
	})
}

func (d *DBSuite) TestNotExistsGet() {
	d.RunOnAllDatastores(func(testStore datastore.Batching) {
		d.addTestCases(testStore)

		has, err := testStore.Has(d.GetTestContext(), datastore.NewKey("/a/b/c/d"))
		if err != nil {
			d.T().Error(err)
		}

		if has {
			d.T().Error("Key should not be found")
		}

		val, err := testStore.Get(d.GetTestContext(), datastore.NewKey("/a/b/c/d"))
		if val != nil {
			d.T().Error("Key should not be found")
		}

		if !errors.Is(err, datastore.ErrNotFound) {
			d.T().Error("Error was not set to ds.ErrNotFound")
			if err != nil {
				d.T().Error(err)
			}
		}
	})
}

func (d *DBSuite) TestDelete() {
	d.RunOnAllDatastores(func(testStore datastore.Batching) {
		d.addTestCases(testStore)

		has, err := testStore.Has(d.GetTestContext(), datastore.NewKey("/a/b/c"))
		if err != nil {
			d.T().Error(err)
		}
		if !has {
			d.T().Error("Key should be found")
		}

		err = testStore.Delete(d.GetTestContext(), datastore.NewKey("/a/b/c"))
		if err != nil {
			d.T().Error(err)
		}

		has, err = testStore.Has(d.GetTestContext(), datastore.NewKey("/a/b/c"))
		if err != nil {
			d.T().Error(err)
		}
		if has {
			d.T().Error("Key should not be found")
		}

	})
}

func (d *DBSuite) TestGetEmpty() {
	d.RunOnAllDatastores(func(testStore datastore.Batching) {
		err := testStore.Put(d.GetTestContext(), datastore.NewKey("/a"), []byte{})
		if err != nil {
			d.T().Error(err)
		}

		v, err := testStore.Get(d.GetTestContext(), datastore.NewKey("/a"))
		if err != nil {
			d.T().Error(err)
		}

		if len(v) != 0 {
			d.T().Error("expected 0 len []byte form get")
		}
	})
}

func (d *DBSuite) TestBatching() {
	d.RunOnAllDatastores(func(testStore datastore.Batching) {
		b, err := testStore.Batch(d.GetTestContext())
		if err != nil {
			d.T().Fatal(err)
		}

		for k, v := range testcases {
			err := b.Put(d.GetTestContext(), datastore.NewKey(k), []byte(v))
			if err != nil {
				d.T().Fatal(err)
			}
		}

		err = b.Commit(d.GetTestContext())
		if err != nil {
			d.T().Fatal(err)
		}

		for k, v := range testcases {
			val, err := testStore.Get(d.GetTestContext(), datastore.NewKey(k))
			if err != nil {
				d.T().Fatal(err)
			}

			if v != string(val) {
				d.T().Fatal("got wrong data!")
			}
		}

		//Test delete

		b, err = testStore.Batch(d.GetTestContext())
		if err != nil {
			d.T().Fatal(err)
		}

		err = b.Delete(d.GetTestContext(), datastore.NewKey("/a/b"))
		if err != nil {
			d.T().Fatal(err)
		}

		err = b.Delete(d.GetTestContext(), datastore.NewKey("/a/b/c"))
		if err != nil {
			d.T().Fatal(err)
		}

		err = b.Commit(d.GetTestContext())
		if err != nil {
			d.T().Fatal(err)
		}

		rs, err := testStore.Query(d.GetTestContext(), query.Query{Prefix: "/"})
		if err != nil {
			d.T().Fatal(err)
		}

		d.expectMatches([]string{
			"/a",
			"/a/b/d",
			"/a/c",
			"/a/d",
			"/e",
			"/f",
			"/g",
		}, rs)

		// TODO: test cancel
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

// TODO: more tests
