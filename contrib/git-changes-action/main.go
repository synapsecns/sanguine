// Package main provides the entrypoint for the action
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector"
	"os"
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

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	ref := os.Getenv("GITHUB_REF")
	explicitRef := githubactions.GetInput("ref")
	if explicitRef != "" {
		ref = explicitRef
	}

	base := githubactions.GetInput("base")

	ct, err := detector.GetChangeTree(context.Background(), workingDirectory, eventName, ref, token, base)
	if err != nil {
		panic(err)
	}

	modules, err := detector.DetectChangedModules(workingDirectory, ct, includeDeps)
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
