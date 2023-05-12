package types

const (
	gasDataOffsetGasPrice     = 0
	gasDataOffsetDataPrice    = 2
	gasDataOffsetExecBuffer   = 4
	gasDataOffsetAmortAttCost = 6
	gasDataOffsetEtherPrice   = 8
	gasDataOffsetMarkup       = 10
	gasDataSize               = 12
)

// GasData is the GasData interface.
type GasData interface {
	// GasPrice is the gas price for the chain (in Wei per gas unit).
	GasPrice() uint16
	// DataPrice is the calldata price (in Wei per byte of content).
	DataPrice() uint16
	// ExecBuffer is the tx fee safety buffer for message execution (in Wei).
	ExecBuffer() uint16
	// AmortAttCost is the amortized cost for attestation submission (in Wei).
	AmortAttCost() uint16
	// EtherPrice is the chain's Ether Price / Mainnet Ether Price (in BWAD).
	EtherPrice() uint16
	// Markup is the markup for the message execution (in BWAD).
	Markup() uint16
}

type gasData struct {
	gasPrice     uint16
	dataPrice    uint16
	execBuffer   uint16
	amortAttCost uint16
	etherPrice   uint16
	markup       uint16
}

// NewGasData creates a new gasdata.
func NewGasData(gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup uint16) GasData {
	return &gasData{
		gasPrice:     gasPrice,
		dataPrice:    dataPrice,
		execBuffer:   execBuffer,
		amortAttCost: amortAttCost,
		etherPrice:   etherPrice,
		markup:       markup,
	}
}

func (g gasData) GasPrice() uint16 {
	return g.gasPrice
}

func (g gasData) DataPrice() uint16 {
	return g.dataPrice
}

func (g gasData) ExecBuffer() uint16 {
	return g.execBuffer
}

func (g gasData) AmortAttCost() uint16 {
	return g.amortAttCost
}

func (g gasData) EtherPrice() uint16 {
	return g.etherPrice
}

func (g gasData) Markup() uint16 {
	return g.markup
}

var _ GasData = gasData{}

// GasDataFromSnapGas takes the raw snapGas and returns a map keyed by domain id to the gas data of that chain
// func GasDataFromSnapGas(snapGas []*big.Int) (map[uint32]GasData, error) {
//	for _, rawGasAsBigInt := range snapGas {
//		rawGasAsBytes := rawGasAsBigInt.Bytes()
//
//	}
//}
