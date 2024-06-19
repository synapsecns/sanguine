package base

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
)

func init() {
	namer := dbcommon.NewNamer(GetAllModels())

	addressName = namer.GetConsistentName("Address")
	typeName = namer.GetConsistentName("Type")
	idName = namer.GetConsistentName("ID")
	dataName = namer.GetConsistentName("Data")
	networkName = namer.GetConsistentName("Network")
	tagName = namer.GetConsistentName("Tag")
	remarkName = namer.GetConsistentName("Remark")
}

var (
	addressName string
	typeName    string
	idName      string
	dataName    string
	networkName string
	tagName     string
	remarkName  string
)
