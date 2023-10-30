package logger

import (
	"github.com/ipfs/go-log"
)

var logger = log.Logger("sinner")

const (
	// ContextCancelled is returned when the context is canceled.
	ContextCancelled ErrorType = iota
	// ScribeFetchFailure is returned when a scribe related query fails.
	ScribeFetchFailure
	// SinnerIndexingFailure is for when the sinner indexer fails.
	SinnerIndexingFailure
	// UnknownTopic is for when an unknown topic is encountered while parsing.
	UnknownTopic
)

const (
	// InitiatingIndexing is returned when a contract begins indexing.
	InitiatingIndexing StatusType = iota
	// CreatingSQLStore is returned when a SQL store is being created.
	CreatingSQLStore
)

// ErrorType is a type of error.
type ErrorType int

// StatusType is a type of status for a process in scribe.
type StatusType int

// ReportSinnerError reports an error that occurs anywhere in sinner.
//
// nolint:exhaustive
func ReportSinnerError(err error, chainID uint32, errorType ErrorType) {
	switch errorType {
	case ContextCancelled:
		logger.Errorf("Context canceled for scribe on chain %d. Error: %v", chainID, err)
	case ScribeFetchFailure:
		logger.Errorf("Scribe fetch failure on chain %d. Error: %v", chainID, err)
	case UnknownTopic:
		logger.Errorf("Sinner parse failure on chain %d. Error: %v", chainID, err)
	default:

		logger.Errorf("Error on chain %d: %v", chainID, err)
	}
}
