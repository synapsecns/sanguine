package cmd

import (
	// used to embed markdown.
	_ "embed"
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/jftuga/termsize"
	"github.com/urfave/cli/v2"
)

//go:embed cmd.md
var help string

// infoCommand references the help info from the cmd.md file and presents it.
var infoCommand = &cli.Command{
	Name:        "info",
	Description: "learn how to use explorer cli",
	Action: func(c *cli.Context) error {
		fmt.Println(string(markdown.Render(help, termsize.Width(), 6)))
		return nil
	},
}

//var portFlag = &cli.UintFlag{
//	Name:  "port",
//	Usage: "--port 5121",
//	Value: 0,
//}
//
//var serverCommand = &cli.Command{
//	Name:        "server",
//	Description: "starts a graphql server",
//	Flags:       []cli.Flag{portFlag, dbFlag, pathFlag},
//	Action: func(c *cli.Context) error {
//		err := api.Start(c.Context, api.Config{
//			HTTPPort: uint16(c.Uint(portFlag.Name)),
//			Database: c.String(dbFlag.Name),
//			Path:     c.String(pathFlag.Name),
//			GRPCPort: uint16(c.Uint(grpcPortFlag.Name)),
//		})
//		if err != nil {
//			return fmt.Errorf("could not start server: %w", err)
//		}
//
//		return nil
//	},
//}
//
//func init() {
//	ports, err := freeport.Take(1)
//	if len(ports) > 0 && err != nil {
//		portFlag.Value = uint(ports[0])
//	}
//}
