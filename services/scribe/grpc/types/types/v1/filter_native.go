package pbscribe

import (
	"github.com/synapsecns/sanguine/services/scribe/db"
)

func (x *LogFilter) ToNative() db.LogFilter {
	logFilter := db.LogFilter{}

	if x.ContractAddress != nil {
		logFilter.ContractAddress = x.ContractAddress.GetData()
	}

	if x.BlockNumber != nil {
		logFilter.BlockNumber = x.BlockNumber.GetData()
	}

	if x.TxHash != nil {
		logFilter.TxHash = x.TxHash.GetData()
	}

	if x.TxIndex != nil {
		logFilter.TxIndex = x.TxIndex.GetData()
	}

	if x.BlockHash != nil {
		logFilter.BlockHash = x.BlockHash.GetData()
	}

	if x.Index != nil {
		logFilter.Index = x.Index.GetData()
	}

	if x.Confirmed != nil {
		logFilter.Confirmed = x.Confirmed.GetData()
	}

	return logFilter
}
