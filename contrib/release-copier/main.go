package main

import (
	"context"
	"github.com/sethvargo/go-githubactions"
	_ "github.com/sethvargo/go-githubactions"
	"github.com/synapsecns/sanguine/contrib/releasecopier/copier"
	"github.com/synapsecns/sanguine/contrib/releasecopier/utils"
	"os"
)

func main() {
	// here we parse a number of variables we use in the action:
	//
	// GITHUB_TOKEN: the github access token
	// GITHUB_REPOSITORY: the repository we are running the action on in the format owner/repo
	repoOwner, repoPath := utils.ParseGithubRepository(os.Getenv("GITHUB_REPOSITORY"))
	token := os.Getenv("GITHUB_TOKEN")

	// we also parse the source and target repositories
	destOwner, destRepo := utils.ParseGithubRepository(githubactions.GetInput("destination_repo"))
	// and the tag to copy
	tagName := githubactions.GetInput("tag_name")
	// the prefix to strip
	stripPrefix := githubactions.GetInput("strip_prefix")

	client := copier.NewReleaseCopier(context.Background(), token)

	client.CopyRelease(context.Background(), repoOwner, repoPath, "synapsecns", "sanguine", "v0.0.1")
}
