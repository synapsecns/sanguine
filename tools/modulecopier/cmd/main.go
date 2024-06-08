package cmd

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/tools/modulecopier/internal"
	"github.com/urfave/cli/v2"
	"os"
)

// Run runs the module copier.
func Run(args []string, buildInfo config.BuildInfo) {
	app := cli.NewApp()
	app.Name = buildInfo.Name()
	app.Version = buildInfo.Version()
	app.Description = buildInfo.VersionString() + "This is used for copying files out of modules in order to export unused fields. This should only be used for unit testing"
	app.Usage = fmt.Sprintf("%s --help", buildInfo.Name())
	app.Flags = []cli.Flag{
		modulePathFlag,
		filePathFlag,
		packageFlag,
	}
	app.Action = func(c *cli.Context) error {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not determine working directory: %w", err)
		}

		modulePath := c.String(modulePathFlag.Name)
		filePath := core.ExpandOrReturnPath(c.String(filePathFlag.Name))
		packageName := c.String(packageFlag.Name)

		// return an error if neither is specified or both are specified
		if (modulePath == "" && filePath == "") || (modulePath != "" && filePath != "") {
			return fmt.Errorf("exactly one of %s OR %s must be specified", modulePathFlag.Name, filePathFlag.Name)
		}

		// handle module path copy
		if modulePath != "" {
			err = internal.CopyModule(modulePath, wd, packageName)
			if err != nil {
				return fmt.Errorf("could not copy files for %s to %s", c.String("module-path"), wd)
			}
		} else {
			// handle go file copy
			err = internal.CopyFile(filePath, wd, packageName)
			if err != nil {
				return fmt.Errorf("could not copy files for %s to %s", c.String("module-path"), wd)
			}
		}

		return nil
	}
	err := app.Run(args)
	if err != nil {
		// we send an additional alert through beep because go:generate *will* silently fail if ran as
		// go:generate ./...
		logoPath, _ := config.GetLogoPath()
		_ = beeep.Notify("GethExport Failed", "", logoPath)
		panic(err)
	}
}
