package anvil

// Hardfork contains the hardfork id used for the --hardfork paramater in anvil.
// these are harcoded to match https://github.com/Rjected/foundry/blob/f3546b91e88d8990ec95bb09b3f7bb88e5e2f910/anvil/src/hardfork.rs#L95
// and don't use iota for tha treason
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Hardfork -linecomment
type Hardfork int

const (
	Frontier       Hardfork = 1  // frontier
	Homestead      Hardfork = 2  // homestead
	Dao            Hardfork = 3  // dao
	Tangerine      Hardfork = 4  // tangerine
	SpuriousDragon Hardfork = 5  // spuriousdragon
	Byzantium      Hardfork = 6  // byzantium
	Constantinople Hardfork = 7  // constantinople
	Petersburg     Hardfork = 8  // petersburg
	Istanbul       Hardfork = 9  // istanbul
	Muirglacier    Hardfork = 10 // muirglacier
	Berlin         Hardfork = 11 // berlin
	London         Hardfork = 12 // london
	ArrowGlacier   Hardfork = 13 // arrowglacier
	GrayGlacier    Hardfork = 14 // grayglacier
	Latest                  = GrayGlacier
)
