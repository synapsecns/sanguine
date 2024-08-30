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

	moduledetector "github.com/synapsecns/sanguine/contrib/git-changes-action/detector/module"
	packagedetector "github.com/synapsecns/sanguine/contrib/git-changes-action/detector/package"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"

	"github.com/sethvargo/go-githubactions"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/git"
)

const defaultTimeout = "1m"

func main() {
	token := githubactions.GetInput("github_token")
	dependencyLevelResolution := githubactions.GetInput("dependencyLevelResolution")

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

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ct, err := git.GetChangeTree(ctx, workingDirectory, ref, token, base)
	if err != nil {
		panic(err)
	}

	noDepChanged, noDepUnchanged, _, err := outputModuleChanges(workingDirectory, ct, false, dependencyLevelResolution)
	if err != nil {
		panic(err)
	}

	depChanged, depUnchanged, allModules, err := outputModuleChanges(workingDirectory, ct, true, dependencyLevelResolution)
	if err != nil {
		panic(err)
	}

	githubactions.SetOutput("changed_modules", noDepChanged)
	githubactions.SetOutput("unchanged_modules", noDepUnchanged)

	githubactions.SetOutput("changed_modules_deps", depChanged)
	githubactions.SetOutput("unchanged_modules_deps", depUnchanged)

	githubactions.SetOutput("all_modules", allModules)

}

// outputModuleChanges outputs the changed modules.
// this wraps detector.DetectChangedModules and handles the output formatting to be parsable by github actions.
// the final output is a json array of strings.
func outputModuleChanges(workingDirectory string, ct tree.Tree, includeDeps bool, dependencyLevelResolution string) (changedJSON string, unchangedJSON string, allModulesJSON string, err error) {
	var modules map[string]bool

	if dependencyLevelResolution == "packages" {
		modules, err = packagedetector.DetectChangedModules(workingDirectory, ct, includeDeps)
	} else {
		modules, err = moduledetector.DetectChangedModules(workingDirectory, ct, includeDeps)
	}

	if err != nil {
		return changedJSON, unchangedJSON, allModulesJSON, fmt.Errorf("failed to detect changed modules w/ include deps set to %v: %w", includeDeps, err)
	}

	var changedModules, unchangedModules, allModules []string
	for module, changed := range modules {
		modName := strings.TrimPrefix(module, "./")
		allModules = append(allModules, modName)

		if changed {
			changedModules = append(changedModules, modName)
		} else {
			unchangedModules = append(unchangedModules, modName)
		}
	}

	sort.Strings(changedModules)
	sort.Strings(unchangedModules)
	sort.Strings(allModules)

	marshalledChanged, err := json.Marshal(changedModules)
	if err != nil {
		return changedJSON, unchangedJSON, allModulesJSON, fmt.Errorf("failed to marshall changed module json w/ include deps set to %v: %w", includeDeps, err)
	}

	marshalledUnchanged, err := json.Marshal(unchangedModules)
	if err != nil {
		return changedJSON, unchangedJSON, allModulesJSON, fmt.Errorf("failed to marshall unchanged module json w/ include deps set to %v: %w", includeDeps, err)
	}

	marshalledAll, err := json.Marshal(allModules)
	if err != nil {
		return changedJSON, unchangedJSON, allModulesJSON, fmt.Errorf("failed to marshall all modules json w/ include deps set to %v: %w", includeDeps, err)
	}

	return string(marshalledChanged), string(marshalledUnchanged), string(marshalledAll), nil
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
