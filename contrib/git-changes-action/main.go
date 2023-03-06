// Package main provides the entrypoint for the action
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/sethvargo/go-githubactions"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector"
)

const defaultTimeout = "1m"

func main() {
	token := githubactions.GetInput("github_token")

	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ref := githubactions.GetInput("ref")
	base := githubactions.GetInput("base")

	timeout, err := getTimeout()
	if err != nil {
		panic(fmt.Errorf("failed to parse timeout: %w", err))
	}

	includeDeps := getIncludeDeps()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ct, err := detector.GetChangeTree(ctx, workingDirectory, ref, token, base)
	if err != nil {
		panic(err)
	}

	modules, err := detector.DetectChangedModules(workingDirectory, ct, includeDeps)
	if err != nil {
		panic(err)
	}

	var changedModules []string
	for module, changed := range modules {
		if !changed {
			continue
		}

		changedModules = append(changedModules, strings.TrimPrefix(module, "./"))
	}

	sort.Strings(changedModules)
	marshalledJSON, err := json.Marshal(changedModules)
	if err != nil {
		panic(err)
	}

	if len(changedModules) == 0 {
		fmt.Println("no modules changed")
	} else {
		fmt.Printf("setting output to %s\n", marshalledJSON)
	}
	githubactions.SetOutput("changed_modules", string(marshalledJSON))
}

// getIncludeDeps gets the include deps setting.
// If it is not set, it defaults to false.
func getIncludeDeps() (includeDeps bool) {
	rawIncludeDeps := githubactions.GetInput("include_deps")

	includeDeps = false
	if rawIncludeDeps == "true" {
		includeDeps = true
	}
	return
}

// getTimeout gets the timeout setting. If it is not set, it defaults to 1 minute.
// Errors if the timeout is not a valid duration.
func getTimeout() (timeout time.Duration, err error) {
	rawTimeout := githubactions.GetInput("timeout")
	if rawTimeout == "" {
		rawTimeout = defaultTimeout
	}

	timeout, err = time.ParseDuration(rawTimeout)
	if err != nil {
		return 0, fmt.Errorf("failed to parse timeout: %w", err)
	}
	return timeout, nil
}
