package pbscribe

import (
	"github.com/synapsecns/sanguine/services/scribe/db"
)

// ToNative converts the log filter to native.
func (x *LogFilter) ToNative() db.LogFilter {
	logFilter := db.LogFilter{}

	var txIndex, blockNumber, index *int

	if x.ContractAddress != nil {
		logFilter.ContractAddress = x.ContractAddress.GetData()
	}

	if x.BlockNumber != nil {
		logFilter.BlockNumber = x.BlockNumber.GetData()
		castedInt := int(logFilter.BlockNumber)
		blockNumber = &castedInt
	}

	if x.TxHash != nil {
		logFilter.TxHash = x.TxHash.GetData()
	}

	if x.TxIndex != nil {
		logFilter.TxIndex = x.TxIndex.GetData()
		castedInt := int(logFilter.TxIndex)
		txIndex = &castedInt
	}

	if x.BlockHash != nil {
		logFilter.BlockHash = x.BlockHash.GetData()
	}

	if x.Index != nil {
		logFilter.Index = x.Index.GetData()
		castedInt := int(logFilter.Index)
		index = &castedInt
	}

	if x.Confirmed != nil {
		logFilter.Confirmed = x.Confirmed.GetData()
	}

	// we use this function to assure new functionality gets included here
	return db.BuildLogFilter(&logFilter.ContractAddress, blockNumber, &logFilter.TxHash, txIndex, &logFilter.BlockHash, index, &logFilter.Confirmed)
}
