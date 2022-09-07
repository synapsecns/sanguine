// Package dev contains dev dependencies required for running developer tasks (coverage testutils, etc)
// that are not required by the project itself. In order to enforce this constraint, this module panics upon
// being imported. Dependencies here are not included in produced binaries and won't affect the dev build
package dev

// import (
// 	_ "github.com/synapsecns/sanguine/tools/abigen"
// )

func init() {
	panic("could not import dev package: this package is meant to define dependencies, not be imported.")
}