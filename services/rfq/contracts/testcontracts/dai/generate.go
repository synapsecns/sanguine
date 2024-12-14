package dai

import "os"

func init() {
	os.Setenv("1_KEY", os.Getenv("ETHERSCAN_KEY"))
}

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate-from-etherscan --address=0x6b175474e89094c44da98b954eedeac495271d0f --chainID 1 --pkg dai --sol-version 0.5.12 --filename=dai --disable-ci --url https://api.etherscan.io/api?
