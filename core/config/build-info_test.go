package config_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"testing"
)

func TestBuildInfo(t *testing.T) {
	buildInfo := config.NewBuildInfo(config.DefaultVersion, config.DefaultCommit, config.AppName, config.DefaultDate)
	Equal(t, buildInfo.Name(), config.AppName)
	Equal(t, buildInfo.Version(), config.DefaultVersion)
	Equal(t, buildInfo.Name(), config.AppName)
	// should use current commit
	NotEqual(t, buildInfo.Commit(), config.DefaultCommit)
	Equal(t, buildInfo.Date(), config.DefaultDate)
}
