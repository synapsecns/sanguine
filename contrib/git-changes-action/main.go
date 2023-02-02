// Package main provides the entrypoint for the action
package main

import (
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector"
)

func main() {
	workingDirectory := githubactions.GetInput("working_directory")
	ref := githubactions.GetInput("ref")
	rawIncludeDeps := githubactions.GetInput("include_deps")
	// TODO: we might need to use a pr as a base
	includeDeps := false
	if rawIncludeDeps == "true" {
		includeDeps = true
	}

	modules, err := detector.DetectChangedModules(workingDirectory, ref, includeDeps)
	if err != nil {
		panic(err)
	}

	marshalledJSON, err := json.Marshal(modules)
	if err != nil {
		panic(err)
	}

	fmt.Println(marshalledJSON)
	githubactions.SetOutput("changed_modules", string(marshalledJSON))
}
