package util

import (
	"fmt"
	// we don't import a specific driver to let the user choose
)

// Options are the sqlite datastore options, reexported here for convenience.
type Options struct {
	Driver string
	DSN    string
	Table  string
	// Don't try to create table
	NoCreate bool

	// sqlcipher extension specific
	Key            []byte
	CipherPageSize uint
}

// Queries are the sqlite queries for a given table.
type Queries struct {
	deleteQuery  string
	existsQuery  string
	getQuery     string
	putQuery     string
	queryQuery   string
	prefixQuery  string
	limitQuery   string
	offsetQuery  string
	getSizeQuery string
}

// NewQueries creates a new sqlite set of queries for the passed table
func NewQueries(tbl string) Queries {
	return Queries{
		deleteQuery:  fmt.Sprintf("DELETE FROM %s WHERE `key` = ?", tbl),
		existsQuery:  fmt.Sprintf("SELECT exists(SELECT 1 FROM %s WHERE `key`=?)", tbl),
		getQuery:     fmt.Sprintf("SELECT data FROM %s WHERE `key` = ?", tbl),
		putQuery:     fmt.Sprintf("REPLACE INTO %s(`key`, `data`) VALUES(?, ?)", tbl),
		queryQuery:   fmt.Sprintf("SELECT `key`, `data` FROM %s", tbl),
		prefixQuery:  " WHERE BINARY `key` LIKE '%s%%' ORDER BY `key`",
		limitQuery:   ` LIMIT %d`,
		offsetQuery:  ` OFFSET %d`,
		getSizeQuery: fmt.Sprintf("SELECT length(`data`) FROM %s WHERE `key` = ?", tbl),
	}
}

// Delete returns the sqlite query for deleting a row.
func (q Queries) Delete() string {
	return q.deleteQuery
}

// Exists returns the sqlite query for determining if a row exists.
func (q Queries) Exists() string {
	return q.existsQuery
}

// Get returns the sqlite query for getting a row.
func (q Queries) Get() string {
	return q.getQuery
}

// Put returns the sqlite query for putting a row.
func (q Queries) Put() string {
	return q.putQuery
}

// Query returns the sqlite query for getting multiple rows.
func (q Queries) Query() string {
	return q.queryQuery
}

// Prefix returns the sqlite query fragment for getting a rows with a key prefix.
func (q Queries) Prefix() string {
	return q.prefixQuery
}

// Limit returns the sqlite query fragment for limiting results.
func (q Queries) Limit() string {
	return q.limitQuery
}

// Offset returns the sqlite query fragment for returning rows from a given offset.
func (q Queries) Offset() string {
	return q.offsetQuery
}

// GetSize returns the sqlite query for determining the size of a value.
func (q Queries) GetSize() string {
	return q.getSizeQuery
}
