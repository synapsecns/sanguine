package db

import submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"

// NotaryDB is the interface for the notary's database.
type NotaryDB interface {
	SubmitterDB() submitterDB.Service
}
