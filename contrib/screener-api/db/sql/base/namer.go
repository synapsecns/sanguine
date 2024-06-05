package base

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
)

func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	addressName = namer.GetConsistentName("Address")
	indicatorName = namer.GetConsistentName("Indicators")
}

var (
	addressName   string
	indicatorName string
)
