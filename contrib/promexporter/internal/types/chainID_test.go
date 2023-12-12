package types_test

import (
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/types"
	"reflect"
	"testing"
)

func TestToInts(t *testing.T) {
	tests := []struct {
		name string
		ids  []types.ChainID
		want []int
	}{
		{
			name: "Test with multiple ChainIDs",
			ids: []types.ChainID{
				types.ETH, types.ROPSTEN, types.RINKEBY, types.GOERLI, types.OPTIMISM,
			},
			want: []int{1, 3, 4, 5, 10},
		},
		{
			name: "Test with a single ChainID",
			ids:  []types.ChainID{types.KOVAN},
			want: []int{42},
		},
		{
			name: "Test with an empty input",
			ids:  []types.ChainID{},
			want: []int{},
		},
	}

	for i := range tests {
		tt := tests[i]

		t.Run(tt.name, func(t *testing.T) {
			got := types.ToInts(tt.ids...)
			if len(got) == len(tt.want) && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chainID.ToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
