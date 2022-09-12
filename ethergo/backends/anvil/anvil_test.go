package anvil_test

import "github.com/synapsecns/sanguine/ethergo/backends/anvil"

func (a *AnvilSuite) TestAnvilBackend() {
	anvil.NewAnvilBackend(a.GetTestContext(), a.T(), &anvil.Config{})
}
