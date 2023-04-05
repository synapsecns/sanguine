## Command Line

The `commandline` package provides utilities for creating an interactive shell command for a [`cli`](github.com/urfave/cli) using the [`go-prompt`](github.com/c-bata/go-prompt) package.

## Usage

To use the `commandline` package, first import the package:

```go
import "github.com/synapsecns/sanguine/core/commandline"
```

Then, create a new `commandline.CommandLine` with the desired value type:


```go
shellCommands := []*cli.Command{
	{
		Name:  "example",
		Usage: "example command",
		Action: func(c *cli.Context) error {
      return nil
    },
  }
}
app := &cli.App{
  Name:    "example",
  Usage:   "example interactive command line",
  Commands: shellCommands,
}
shellCommand := commandline.GenerateShellCommand(shellCommands)
app.Commands = append(app.Commands, shellCommand)
```

## Interactive Shell

The `commandline` package provides an interactive shell for the `cli` application. The shell can be started by running the `run` command:

```bash
$ example shell
```

## Autocompletion

Autocompletion is supported for both commands and flags. To see the available commands, enter help. To see `help` for a specific command, enter help <command>.

Autocompletion for flags will only suggest flags that are valid for the current command. For example, if the current command is `example`, then the flag `--example` will be suggested, but not `--example2`.

## Exiting

To exit the shell, you can type `quit`, `q` or `exit`.


