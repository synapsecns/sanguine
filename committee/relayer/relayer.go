package relayer

import "github.com/synapsecns/sanguine/ethergo/submitter"

// get the events
// get the signatures
// call receiveModuleMessage

type Relayer interface {
}

type SignatureStore interface {
}

type relayerImpl struct {
	submitter *submitter.TransactionSubmitter
}
