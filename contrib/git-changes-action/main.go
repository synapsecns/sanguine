// Package main provides the entrypoint for the action
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector"
	"os"
	"sort"
	"strings"
)

func main() {
	token := githubactions.GetInput("github_token")

	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rawIncludeDeps := githubactions.GetInput("include_deps")
	// TODO: we might need to use a pr as a base
	includeDeps := false
	if rawIncludeDeps == "true" {
		includeDeps = true
	}

	ref := githubactions.GetInput("ref")
	base := githubactions.GetInput("base")

	ct, err := detector.GetChangeTree(context.Background(), workingDirectory, ref, token, base)
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
