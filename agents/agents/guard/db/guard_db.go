package db

import submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"

// GuardDB is the interface for the guard's database.
type GuardDB interface {
	SubmitterDB() submitterDB.Service
}
