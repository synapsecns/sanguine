//go:build generate
// +build generate

package usdc

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol-version 0.6.2 --pkg usdc --filename usdc --sol /home/ubuntu/repos/sanguine/services/rfq/contracts/testcontracts/usdc/USDC.sol
