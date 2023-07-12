package error

var logger = log.Logger("scribe")

const (
	// ContextCancelled is returned when the context is cancelled
	ContextCancelled ErrorType = iota
	// GetLogsError is returned when the logs cannot be retrieved
	GetLogsError
	// GetTxError is returned when the tx cannot be retrieved
	GetTxError
	// CouldNotGetReceiptError is returned when the receipt cannot be retrieved
	CouldNotGetReceiptError
	// GetBlockError is returned when the block cannot be retrieved
	GetBlockError
	// BlockByNumberError is returned when the block cannot be retrieved
	BlockByNumberError
	// StoreError is returned when data cannot be inserted into the database
	StoreError
	// ReadError is returned when data cannot be read from the database
	ReadError
)

type ErrorType int

func HandleError(err error, errorType ErrorType) error {
	switch errorType {
	case ContextCancelled:
		return contextCancelled(err)
	case GetLogsError:
		return getLogsError(err)
	case GetTxError:
		return getTxError(err)
	case CouldNotGetReceiptError:
		return couldNotGetReceiptError(err)
	case GetBlockError:
		return getBlockError(err)
	case BlockByNumberError:
		return blockByNumberError(err)
	case StoreError:
		return storeError(err)
	case ReadError:
		return readError(err)
	default:
		return err

	}

}
