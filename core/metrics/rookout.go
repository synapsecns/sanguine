package metrics

import (
	rookout "github.com/Rookout/GoSDK"
	"github.com/Rookout/GoSDK/pkg/config"
	"github.com/synapsecns/sanguine/core"
	synconfig "github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"os"
)

// rookout is sideloaded.
// TODO: consider moving this to metrics
func init() {
	if core.HasEnv(internal.RookoutToken) {
		// TODO: Consider doing git stuff here.
		err := rookout.Start(config.RookOptions{
			Token:     os.Getenv(internal.RookoutToken),
			Debug:     core.HasEnv(internal.RookoutDebug),
			GitCommit: synconfig.DefaultCommit,
			// note: will chagne on fork
			GitOrigin: "https://github.com/synapsecns/sanguine",
		})
		if err != nil {
			logger.Warn(err)
		}
	}
}
