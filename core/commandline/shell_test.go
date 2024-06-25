package commandline

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/urfave/cli/v2"
)

func TestSignalHandling(t *testing.T) {
	shellCommands := []*cli.Command{
		{
			Name:  "test",
			Usage: "test command",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app := &cli.App{
		Commands: []*cli.Command{
			GenerateShellCommand(shellCommands),
		},
	}

	// Set up a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(1 * time.Second)
		p, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Errorf("Failed to find process: %v", err)
			return
		}
		if err := p.Signal(syscall.SIGINT); err != nil {
			t.Errorf("Failed to send SIGINT: %v", err)
		}
	}()

	err := app.RunContext(ctx, []string{"app", "shell"})
	if err != nil {
		t.Fatalf("app.RunContext failed: %v", err)
	}
}

func TestMain(m *testing.M) {
	// Call signal.Notify so that the test process does not exit on SIGINT
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %v", sig)
	}()
}
