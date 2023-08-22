package chainmanager_test

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/richardwilkes/toolbox/collection"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	"github.com/synapsecns/sanguine/services/omnirpc/metadata"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"
	"sort"
	"testing"
	"time"
)

func TestRefreshRPCInfoNil(t *testing.T) {
	nullHandler, err := metrics.NewByType(context.Background(), metadata.BuildInfo(), metrics.Null)
	NoError(t, err)

	cm := chainmanager.NewChainManager(nullHandler)

	// make sure we don't panic if the chain is not nil
	NotPanics(t, func() {
		cm.RefreshRPCInfo(context.Background(), gofakeit.Uint32())
	})
}

func TestSortInfoList(t *testing.T) {
	// errors first
	infoList := []rpcinfo.Result{
		{
			URL:      "err",
			HasError: true,
			Error:    errors.New(gofakeit.Sentence(10)),
		},
		{
			URL:      "err",
			HasError: true,
			Error:    errors.New(gofakeit.Sentence(10)),
		},
		{
			URL:      "laggiest",
			BlockAge: time.Second * 20,
			Latency:  time.Second * 20,
		},
		{
			URL:      "laggier",
			BlockAge: time.Second * 20,
			Latency:  time.Second * 10,
		},
		{
			URL:     "slowest",
			Latency: time.Second * 20,
		},
		{
			URL:     "slower",
			Latency: time.Second * 10,
		},
	}

	sortedInfoList := chainmanager.SortInfoList(infoList)
	Equal(t, infoList, sortedInfoList)
}

func TestGetChainIDs(t *testing.T) {
	nullHandler, err := metrics.NewByType(context.Background(), metadata.BuildInfo(), metrics.Null)
	NoError(t, err)

	cm := chainmanager.NewChainManager(nullHandler)

	chainIDs := collection.Set[uint32]{}

	for i := 0; i < 40; i++ {
		testChainID := gofakeit.Uint32()

		chainIDs.Add(testChainID)
		cm.PutChain(testChainID, []string{gofakeit.URL(), gofakeit.URL()}, gofakeit.Uint16())
	}

	// sort both slices to assert equality
	Equal(t, sortSlice(cm.GetChainIDs()), sortSlice(chainIDs.Values()))
}

// sortSlice sorts a uint 32 slice.
func sortSlice(res []uint32) []uint32 {
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	return res
}
