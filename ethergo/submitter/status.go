package submitter

import "github.com/ethereum/go-ethereum/common"

// SubmissionState is the status of a submission. This is not used internally and only serves as
// a way to communicate the status of a submission to the caller.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=SubmissionState -linecomment
type SubmissionState uint8

const (
	// NotFound indicates that the submission was not found.
	NotFound SubmissionState = iota
	// Pending indicates that the submission is pending.
	Pending // pending
	// Confirming indicates that the submission is confirming. The tx has completed on chain, but
	// no txhash has been associated with it yet.
	Confirming // confirming
	// Confirmed indicates that the submission is confirmed and txhash data is available.
	Confirmed // confirmed
)

// SubmissionStatus is the status of a submission.
type SubmissionStatus interface {
	// State is the state of the submission.
	State() SubmissionState
	// HasTx indicates whether the submission has a transaction.
	HasTx() bool
	// TxHash is the hash of the transaction. This will be the zero hash if HasTx is false.
	TxHash() common.Hash
}

type submissionStatusImpl struct {
	state  SubmissionState
	txHash common.Hash
}

func (s submissionStatusImpl) State() SubmissionState {
	return s.state
}

func (s submissionStatusImpl) HasTx() bool {
	return s.state == Confirmed
}

func (s submissionStatusImpl) TxHash() common.Hash {
	return s.txHash
}

var _ SubmissionStatus = &submissionStatusImpl{}
