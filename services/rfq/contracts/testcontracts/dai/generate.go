//go:build generate
// +build generate

package dai

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol-version 0.6.2 --pkg dai --filename dai --sol /home/ubuntu/repos/sanguine/services/rfq/contracts/testcontracts/dai/DAI.sol
