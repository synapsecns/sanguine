package copier

import (
	"context"
	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
	"sync"
)

// ReleaseCopier contains the release copier client.
type ReleaseCopier struct {
	client                              *github.Client
	sourceOwner, sourceRepo, targetRepo string
	// mux ensures only one copy can be made at a time
	mux sync.Mutex
}

// NewImageCopier creates a new image copier client.
func NewImageCopier(ctx context.Context, token string) *ReleaseCopier {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})

	tc := oauth2.NewClient(ctx, ts)

	return &ReleaseCopier{
		client: github.NewClient(tc),
	}
}

func (r *ReleaseCopier) CopyImages() {

}
