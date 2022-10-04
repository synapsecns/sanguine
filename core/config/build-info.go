package config

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/integralist/go-findroot/find"
	"time"
)

// AppName is the application name.
const AppName = "Sanguine"

// DefaultVersion is the application version.
const DefaultVersion = "0.1.0"

// DefaultCommit is the default commit. Build info will attempt to replace
// with the current commit if not present.
const DefaultCommit = "none"

// DefaultDate when not passed in by the compiler.
var DefaultDate = time.Now().Format(time.RFC3339)

// VendorName is the vendor named used for versioning schemes that depend on a vendor name
// we use the github name for convince.
const VendorName = "synapsecns"

// BuildInfo will contains build info from https://goreleaser.com/cookbooks/using-main.version
type BuildInfo struct {
	version string
	commit  string
	name    string
	date    string
}

// NewBuildInfo creates a build info struct from buildtime data
// it sets sensible defaults.
func NewBuildInfo(version, commit, name, date string) BuildInfo {
	if commit == DefaultCommit {
		commit = getCurrentCommit()
	}

	return BuildInfo{
		version: version,
		commit:  commit,
		name:    name,
		date:    date,
	}
}

// getCurrentCommit sets the commit from the local repo or uses default if not found.
func getCurrentCommit() string {
	root, err := find.Repo()
	// nothing we can do in this case, we'll use unknown
	if err != nil {
		return DefaultCommit
	}

	repo, err := git.PlainOpen(root.Path)
	if err != nil {
		return DefaultCommit
	}

	// get the current commit hash
	ref, err := repo.Head()
	if err != nil {
		return DefaultCommit
	}

	return ref.Strings()[0]
}

// Version of the build.
func (b BuildInfo) Version() string {
	return b.version
}

// Commit of the build.
func (b BuildInfo) Commit() string {
	return b.commit
}

// Name of the application.
func (b BuildInfo) Name() string {
	return b.name
}

// Date the application was built.
func (b BuildInfo) Date() string {
	return b.date
}

// VersionString pretty prints a version string with the info above.
func (b BuildInfo) VersionString() string {
	return fmt.Sprintf("%s: (commit: %s), commit (date: %s) \n", b.version, b.commit, b.date)
}
