package commandline

import (
	"context"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"syscall"
)

const shellCommandName = "shell"

// GenerateShellCommand generates the shell command with a list of commands that the shell should take.
// TODO: this needs a more comprehensive test suite.
// TODO: support ctrl+c.
func GenerateShellCommand(shellCommands []*cli.Command) *cli.Command {
	// explicitly exclude shell if included
	capturedCommands := pruneShellCommands(shellCommands)

	// make sure tty is open, this will not be the case in distroless containers
	_, err := syscall.Open("/dev/tty", syscall.O_RDONLY, 0)
	shellAvailable := err == nil

	return &cli.Command{
		Name:  shellCommandName,
		Usage: "start an interactive shell.",
		Flags: []cli.Flag{
			&LogLevel,
		},
		Action: func(c *cli.Context) (err error) {
			SetLogLevel(c)

			console := cli.NewApp()
			console.Commands = capturedCommands
			console.Action = func(c *cli.Context) error {
				fmt.Printf("Command not found. Type 'help' for a list of commands or \"%s\", \"%s\" or \"%s\" to exit.\n", quitCommand, exitCommand, quitCommandShort)
				return nil
			}

			if c.Args().Len() == 0 {
				err := console.RunContext(c.Context, strings.Fields("cmd help"))
				if err != nil {
					return errors.Wrap(err, "could not show help")
				}
			}

			// warn user about sigterms
			sigs := make(chan os.Signal)
			go func() {
				for range sigs {
					fmt.Printf("\n(type \"%s\", \"%s\" or \"%s\" to exit)\n\n >", quitCommand, exitCommand, quitCommandShort)
				}
			}()
			//nolint: govet
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			defer func() {
				signal.Stop(sigs)
				close(sigs)
			}()

			if !shellAvailable {
				fmt.Println("Shell is not available in this environment because /dev/tty is not available. This is expected in containerized images")
				return nil
			}

			interactive := newInteractiveClient(c.Context, capturedCommands, console)
			for {
				p := prompt.New(
					interactive.executor,
					interactive.completor,
					prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
					prompt.OptionMaxSuggestion(3),
					prompt.OptionLivePrefix(livePrefix),
				)
				p.Run()
			}
		},
	}
}

// pruneShellCommands gets a list of commands including the shell command.
func pruneShellCommands(commands []*cli.Command) (prunedCommands []*cli.Command) {
	// initialize shell commands
	for _, command := range commands {
		if command.Name != shellCommandName {
			prunedCommands = append(prunedCommands, command)
		}
	}
	return prunedCommands
}

// ShellCmd is used to launch an interactive shell. This is useful for repeatedly interacting with the cli or testing.

// livePrefix generates a prefix with the current user directory.
// this is useful for file relative commands like creating a new geth account.
func livePrefix() (prefix string, useLivePrefix bool) {
	pwd, err := os.Getwd()
	useLivePrefix = true
	if err != nil {
		prefix += " $ "
		return
	}

	prefix += " " + pwd + " $ "

	u, err := user.Current()
	if err != nil {
		return
	}

	prefix = strings.ReplaceAll(prefix, u.HomeDir, "~")
	return
}

// interactiveClient object.
type interactiveClient struct {
	// cli app
	app *cli.App
	// ctx is the parent context
	//nolint: containedctx
	ctx context.Context
	// shellCommands are all shell commands supported by the interactive client
	shellCommands []*cli.Command
}

// newInteractiveClient creates a new interactive client.
func newInteractiveClient(ctx context.Context, shellCommands []*cli.Command, app *cli.App) *interactiveClient {
	return &interactiveClient{
		app:           app,
		shellCommands: shellCommands,
		ctx:           ctx,
	}
}

// completor handles autocompletion for the interactive client.
func (i *interactiveClient) completor(in prompt.Document) []prompt.Suggest {
	// commandPrompts are prompts for commands (no flags)
	commandPrompts := []prompt.Suggest{
		{
			Text:        "help",
			Description: "Shows a list of commands or help for one command",
		},
	}

	// flagPrompts promp flags for each command
	var flagPrompts []prompt.Suggest

	for _, command := range i.shellCommands {
		commandPrompts = append(commandPrompts, prompt.Suggest{
			Text:        command.Name,
			Description: command.Usage,
		})

		for _, flag := range command.Flags {
			docFlag, ok := flag.(cli.DocGenerationFlag)
			if ok {
				flagPrompts = append(flagPrompts, prompt.Suggest{
					Text:        fmt.Sprintf("%s --%s ", command.Name, longestFlag(flag)),
					Description: docFlag.GetUsage(),
				})
			}
		}
	}

	// get the prompts for the current text. if the user has entered part of a command this will be
	// command prompts. if the user has entered a whole command flags will be suggested
	prompts := prompt.FilterHasPrefix(commandPrompts, in.CurrentLineBeforeCursor(), true)
	if len(prompts) == 0 {
		prompts = prompt.FilterHasPrefix(flagPrompts, in.CurrentLineBeforeCursor(), true)

		// right now, autocomplete for the flags will not take into account the first word so
		// autocompleting `start --config` will return `start start --config`. Here, since we've
		// already done the matching we cut off hte command prefix
		for i, currentPrompt := range prompts {
			splitText := strings.Split(currentPrompt.Text, " ")
			currentPrompt.Text = strings.Join(splitText[1:], " ")
			prompts[i] = currentPrompt
		}
	}

	return prompts
}

// longestFlag gets the longest flag from all flag names
// this is used to avoid short (non-descriptive) flags from coming up in the autocomplete.
func longestFlag(flag cli.Flag) (flagName string) {
	maxLength := 0
	for _, name := range flag.Names() {
		flagLength := len(name)
		if flagLength > maxLength {
			maxLength = flagLength
			flagName = name
		}
	}
	return flagName
}

// executor handles executing interactive commands.
func (i *interactiveClient) executor(line string) {
	if line == "" {
		return
	}

	if line == quitCommand || line == quitCommandShort || line == exitCommand {
		os.Exit(0)
	}

	if line == "shell" {
		fmt.Println("cannot start a shell from within a shell!")
		return
	}

	err := i.app.RunContext(i.ctx, strings.Fields("cmd "+line))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

const (
	quitCommand      = "quit"
	quitCommandShort = "q"
	exitCommand      = "exit"
)
