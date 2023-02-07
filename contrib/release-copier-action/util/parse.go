// Package util contains utility functions for parsing action inputs
package util

import "strings"

// ParseGithubRepository parses ghte igthub repository from the GITHUB_REPOSITORY environment variable
// this comes in the format owner/repo. This function returns the owner and repo as separate strings.
func ParseGithubRepository(githubRepo string) (repoOwner, repoName string) {
	//nolint: gocritic
	repoOwner = githubRepo[:strings.Index(githubRepo, "/")]
	repoName = githubRepo[strings.Index(githubRepo, "/")+1:]
	return
}
