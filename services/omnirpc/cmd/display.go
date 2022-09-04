package cmd

import (
	"fmt"
	"github.com/jftuga/ellipsis"
	"github.com/olekukonko/tablewriter"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"
	"os"
	"sort"
)

// DisplayLatency displays latency results in a cli.
func DisplayLatency(lat []rpcinfo.Result) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"URL", "Latency", "HasError (if any)"})

	sort.Slice(lat, func(i, j int) bool {
		return lat[i].Latency < lat[j].Latency
	})

	for _, latencyResult := range lat {
		var reason string
		if latencyResult.HasError {
			reason = ellipsis.Shorten(latencyResult.Error.Error(), 20)
		}

		table.Append([]string{
			latencyResult.URL,
			fmt.Sprintf("%s milliseconds", latencyResult.Latency),
			reason,
		})
	}
	table.Render()
}
