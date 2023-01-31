// Package user defines all end-user defined messages we resolve to in the explorer
package user

import (
	"github.com/synapsecns/sanguine/services/explorer/contracts/user/dfk/dfkhero"
	"github.com/synapsecns/sanguine/services/explorer/contracts/user/dfk/dfkpet"
)

// MessageFormat contains metadata about the use of a message format
type MessageFormat map[string]MessageFormatter

type MessageFormatter struct {
	DataType interface{}
}

var messageFormats MessageFormat

func GetMessageFormats() MessageFormat {
	return messageFormats
}

func init() {
	// add all types here
	messageFormats = make(MessageFormat)
	messageFormats["dfkhero"] = MessageFormatter{DataType: dfkhero.HeroBridgeUpgradeableMessageFormat{}}
	messageFormats["dfkpet"] = MessageFormatter{DataType: dfkpet.PetBridgeUpgradeableMessageFormat{}}
	messageFormats["dfktear"] = MessageFormatter{DataType: dfkpet.PetBridgeUpgradeableMessageFormat{}}
}
