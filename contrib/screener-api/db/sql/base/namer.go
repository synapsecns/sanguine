package base

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
)

func init() {
	namer := dbcommon.NewNamer(GetAllModels())

	typeName = namer.GetConsistentName("Type")
	idName = namer.GetConsistentName("ID")
	dataAddressName = namer.GetConsistentName("Address")
	dataNetworkName = namer.GetConsistentName("Network")
	dataTagName = namer.GetConsistentName("Tag")
	dataRemarkName = namer.GetConsistentName("Remark")
}

var (
	typeName        string
	idName          string
	dataAddressName string
	dataNetworkName string
	dataTagName     string
	dataRemarkName  string
)
