// Package main contains a generator for copying files exported files from geth
// in order to use private fields. The resulting files should not be modified directly
// but if there are new methods you need exported, generators, etc that can be done in other files
// that will now have access to the private fields. These generated files should only be used for testing
//
// TODO: look into implementing a tag for tests in order to make sure nothing in testutils/ is used in a production build
// we haven't done this yet because of the poor ux in an ide as far as having to add a `-tag`.
package cmd

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/synapsecns/sanguine/tools/modulecopier/internal"
	"github.com/synapsecns/synapse-node/config"
	"github.com/urfave/cli/v2"
	"os"
)

const appName = "modulecopier"

func Run(args []string) {
	app := cli.NewApp()
	app.Name = appName
	app.Version = config.AppVersion
	app.Description = "This is used for copying files out of modules in order to export unused fields. This should only be used for unit testing"
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
		filePath := c.String(filePathFlag.Name)
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
