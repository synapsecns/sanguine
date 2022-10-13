// Package internal contains dev dependencies required for running developer tasks (coverage testing, etc)
// that are not required by the project itself. In order to enforce this constraint, this module panics upon
// being imported. Dependencies here are not included in produced binaries and won't affect the dev build
package internal

import (
	"github.com/BurntSushi/toml"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/tools"
	"github.com/ugorji/go/codec"
	"github.com/vburenin/ifacemaker/maker"
	"github.com/vektra/mockery/v2/pkg"
)

func init() {
	// This package should never be imported. If it is, panic.
	// we ignore this in tests because -covermode will recursively try to run all packages
	if !core.IsTest() {
		panic("could not import internal package: this package is meant to define dependencies, not be imported.")
	}
}

var _ = codec.Decoder{}
var _ = maker.ParseStruct
var _ = toml.Unmarshal

// required by abigen.
var _ = tools.Importable{}

// required by mockery.
var _ = pkg.Method{}
