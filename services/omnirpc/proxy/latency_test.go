package proxy_test

import (
	backends2 "github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"sync"
)

func (p *ProxySuite) TestReorder() {
	presets := []preset.Backend{preset.GetAvalancheLocal(), preset.GetRinkeby(), preset.GetMaticMumbai()}

	backends := make([]backends2.SimulatedTestBackend, len(presets))
	var wg sync.WaitGroup

	// setup 3 backends, we'll use these to get do different block heights
	for i, presetBackend := range presets {
		var mux sync.Mutex
		wg.Add(1)
		go func(index int, initializer preset.Backend) {
			be := initializer.Geth(p.GetTestContext(), p.T())

			mux.Lock()
			defer mux.Unlock()
			backends[index] = be
		}(i, presetBackend)
	}

	wg.Done()
}
