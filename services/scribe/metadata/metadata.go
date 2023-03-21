// Package metadata provides a metadata service for Scribe.
package metadata

import "github.com/synapsecns/sanguine/core/config"

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

// BuildInfo returns the build info for the service.
func BuildInfo() config.BuildInfo {
	return config.NewBuildInfo(version, commit, "scribe", date)
}
