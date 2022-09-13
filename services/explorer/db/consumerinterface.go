package db

// ConsumerDBWriter is the interface for writing events to the ConsumerDB.
type ConsumerDBWriter interface {
}

// ConsumerDBReader is the interface for reading events from the ConsumerDB.
type ConsumerDBReader interface {
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
