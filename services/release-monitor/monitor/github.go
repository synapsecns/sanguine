package monitor

import (
	"context"
	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
	"net/http"
)

// newGithubClient creates a new github client
func newGithubClient(ctx context.Context, apiKey string) *github.Client {
	httpClient := http.DefaultClient
	if apiKey != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})
		httpClient = oauth2.NewClient(ctx, ts)
	}

	return github.NewClient(httpClient)
}
