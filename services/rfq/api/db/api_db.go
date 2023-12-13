package db

// ApiDBReader is the interface for reading from the database.
type ApiDBReader interface {
}

// ApiDBWriter is the interface for writing to the database.
type ApiDBWriter interface {
}

// ApiDB is the interface for the database service.
type ApiDB interface {
	ApiDBReader
	ApiDBWriter
}
