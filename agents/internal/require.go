// Package internal contains dev dependencies required for running developer tasks (coverage testing, etc)
// that are not required by the project itself. In order to enforce this constraint, this module panics upon
// being imported. Dependencies here are not included in produced binaries and won't affect the dev build
package internal

import (
	"github.com/BurntSushi/toml"
	"github.com/dgraph-io/ristretto"
	"github.com/go-playground/validator/v10"
	"github.com/ugorji/go/codec"
	"github.com/urfave/cli/v2"
	"github.com/vburenin/ifacemaker/maker"
	"github.com/vektra/mockery/v2/pkg"
	"golang.org/x/exp/rand"
)

func init() {
	panic("could not import dev package: this package is meant to define dependencies, not be imported.")
}

var _ = codec.Decoder{}
var _ = maker.ParseStruct
var _ = toml.Unmarshal

// required by abigen.
var _ = rand.Int
var _ = ristretto.Config{}
var _ = validator.Validate{}
var _ = cli.StringFlag{}

// required by mockery.
var _ = pkg.Method{}
