package cover

import (
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/tools/cover/cmd"
	"os"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "cover", date)
	cmd.Run(os.Args, buildInfo)

}
