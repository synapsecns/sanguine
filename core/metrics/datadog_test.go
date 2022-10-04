package metrics_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"golang.org/x/exp/slices"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"strings"
	"testing"
	"time"
)

func TestAllProfileString(t *testing.T) {
	rawProfiles := strings.Split(metrics.AllProfilesString(), ",")
	Equal(t, len(rawProfiles), len(metrics.AllProfileTypes()))

	for profile := range metrics.AllProfileTypes() {
		_, ok := metrics.AllProfileTypes()[strings.ToLower(profile)]
		Truef(t, ok, "could not find %s", ok)
	}
}

func TestGetProfileTypesFromEnv(t *testing.T) {
	testProfiles := []profiler.ProfileType{
		profiler.BlockProfile,
		profiler.HeapProfile,
	}

	// add in a duplicate
	t.Setenv(metrics.DDProfileEnv, strings.Join([]string{testProfiles[0].String(), testProfiles[1].String(), testProfiles[1].String()}, ","))

	returnedTypes := metrics.GetProfileTypesFromEnv()
	Equal(t, len(returnedTypes), len(testProfiles))

	for _, ret := range returnedTypes {
		True(t, slices.Contains(testProfiles, ret))
	}
}

func TestHandler(t *testing.T) {
	NotPanics(t, func() {
		newHandler := metrics.NewDatadogMetricsHandler(config.NewBuildInfo(config.DefaultVersion, config.DefaultCommit, config.AppName, config.DefaultDate))
		err := newHandler.Start(context.Background())
		Nil(t, err)
		time.Sleep(time.Millisecond * 50)
	})
}
