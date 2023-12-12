// Package db implements the data store
//
// db
// ├── db.go: Houses the interface in which the database can be interacted with. All functions for both read/write/util are here.
// ├── read_test.go: All the tests for the read operations for the database (its testing all the functions in the `Reader` interface in db.go).
// ├── write_test.go: All the tests for the write operations for the database (its testing all the functions in the `Writer` interface in db.go).
// ├── suite_test.go: Sets up the test environment (inits databases).
// ├── sql: Contains all implementations from `db.go` and the init functions for both mysql and sqlite.
// │ ├── base: Holds all the implementations of `Reader` and `Writer`. Also has the database init functionality.
// │ ├── mysql: Mysql init functionality.
// │ └── sqlite: Sqlite init functionality.
// └── model: Has all the models for each table, as well as all the namer initializing.
package db
