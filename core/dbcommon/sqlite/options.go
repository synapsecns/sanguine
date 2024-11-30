package sqlite3dump

// Option is SQL dump option.
type Option func(dumper *sqlite3dumper)

// WithMigration option won't include creation tables and will include table column names.
func WithMigration() Option {
	return func(dumper *sqlite3dumper) {
		dumper.migration = true
	}
}

// WithDropIfExists option drops existing table or index if it already exists.
func WithDropIfExists(dropIfExists bool) Option {
	return func(dumper *sqlite3dumper) {
		dumper.dropIfExists = dropIfExists
	}
}

// WithTransaction wraps query with transaction.
//
// Adds 'BEGIN TRANSACTION' at start and 'COMMIT' at the end.
func WithTransaction(addTransaction bool) Option {
	return func(dumper *sqlite3dumper) {
		dumper.wrapWithTransaction = addTransaction
	}
}
