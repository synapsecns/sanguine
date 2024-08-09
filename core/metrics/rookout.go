package metrics

import (
	rookout "github.com/Rookout/GoSDK"
	"github.com/Rookout/GoSDK/pkg/config"
	"github.com/synapsecns/sanguine/core"
	synconfig "github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"os"
)

// DefaultGitRepo is the default git repo for sanguine.
// exposed here to allow it to be overridden by an ldflag rather than an env var.
const DefaultGitRepo = "https://github.com/synapsecns/sanguine"

// rookout is sideloaded.
// TODO: consider moving this to metrics.
func init() {
	if core.HasEnv(internal.RookoutToken) {
		// some env vars are supported here to allow config docs to be valid: https://docs.rookout.com/docs/setup-guide/#configuration.
		// we do not document these as they are not intended to be used.
		err := rookout.Start(config.RookOptions{
			Token:     os.Getenv(internal.RookoutToken),
			Debug:     core.HasEnv(internal.RookoutDebug),
			GitCommit: core.GetEnv(internal.RookoutCommit, synconfig.DefaultCommit),
			// note: we do not document
			GitOrigin: core.GetEnv(internal.RookoutRemoteOrigin, core.GetEnv(internal.GitRepo, DefaultGitRepo)),
		})
		if err != nil {
			logger.Warn(err)
		}
	}
}
