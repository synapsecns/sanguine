// Package main provides the entrypoint for the action
package main

import (
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector"
	"os"
)

func main() {
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

	modules, err := detector.DetectChangedModules(workingDirectory, os.Getenv("GITHUB_REF"), includeDeps)
	if err != nil {
		panic(err)
	}

	marshalledJSON, err := json.Marshal(modules)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshalledJSON))
	githubactions.SetOutput("changed_modules", string(marshalledJSON))
}
