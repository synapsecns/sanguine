package anvil

// Hardfork indicates which chain should be hardforked by foundry
// see: https://github.com/foundry-rs/foundry/blob/master/anvil/src/hardfork.rs#L94
// note: because of the way foundry does alias being incompatible w/ stringer, latest will have the wrong id when passed as an int
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Hardfork -linecomment
type Hardfork uint8

const (
	// Frontier is the initial release of the Ethereum protocol.
	Frontier Hardfork = iota + 1 // frontier
	// Homestead is the chain config corresponding to the homestead hardfork.
	Homestead // homestead
	// DAO is the chain config corresponding to the DAO hardfork.
	DAO // dao
	// Tangerine is the chain config corresponding to the tangerine hardfork.
	Tangerine // tangerine
	// Spurious is the chain config corresponding to the spurious hardfork.
	Spurious // spuriousdragon
	// Byzantium is the chain config corresponding to the byzantium hardfork.
	Byzantium // byzantium
	// Constantinople is the chain config corresponding to the constantinople hardfork.
	Constantinople // constantinople
	// Petersburg is the chain config corresponding to the petersburg hardfork.
	Petersburg // petersburg
	// Istanbul is the chain config corresponding to the istanbul hardfork.
	Istanbul // istanbul
	// MuirGlacier is the chain config corresponding to the muir glacier hardfork.
	MuirGlacier // muirglacier
	// Berlin is the chain config corresponding to the berlin hardfork.
	Berlin // berlin
	// London is the chain config corresponding to the london hardfork.
	London // london
	// ArrowGlacier is the chain config corresponding to the arrow glacier hardfork.
	ArrowGlacier // arrowglacier
	// GlayGlacier is the chain config corresponding to the glay glacier hardfork.
	GlayGlacier // glayglacier
	// Latest is the chain config corresponding to the latest hardfork.
	// the int of this is not correct (1 to hig), but it is not used in foundry option builder.
	Latest // latest
)

// allHardforks is a list of all hardforks.
var allHardforks = []Hardfork{
	Frontier,
	Homestead,
	DAO,
	Tangerine,
	Spurious,
	Byzantium,
	Constantinople,
	Petersburg,
	Istanbul,
	MuirGlacier,
	Berlin,
	London,
	ArrowGlacier,
	GlayGlacier,
	Latest,
}

func init() {
	if len(_Hardfork_index)-1 != len(allHardforks) {
		panic("not all hardforks have been added to the stringer")
	}
}
