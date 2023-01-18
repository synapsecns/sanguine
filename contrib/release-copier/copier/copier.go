package copier

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
	"strings"
)

// ReleaseCopier contains the release copier client
type ReleaseCopier struct {
	client *github.Client
}

// NewReleaseCopier creates a new release copier client
func NewReleaseCopier(ctx context.Context, token string) *ReleaseCopier {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})

	tc := oauth2.NewClient(ctx, ts)

	return &ReleaseCopier{
		client: github.NewClient(tc),
	}
}

func (r *ReleaseCopier) CopyRelease(ctx context.Context, sourceOwner, sourceRepo, targetOwner, targetRepo, tagName, stripPrefix string) error {
	// Get the release
	ogRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, sourceOwner, sourceRepo, tagName)
	if err != nil {
		return fmt.Errorf("could not get origin release: %w", err)
	}

	if ogRelease.TagName == nil {
		return errors.New("could not get origin release tag name, tag is required for copying a release")
	}

	strippedTag := strings.TrimPrefix(tagName, stripPrefix)
	// releaes name is optional, so we only set it if it exists
	// we also strip the prfix here
	name := ""
	if ogRelease.Name != nil {
		name = strings.TrimPrefix(*ogRelease.Name, stripPrefix)
	}

	// Create the release
	newRelease := &github.RepositoryRelease{
		TagName: &strippedTag,
		Name:    &name,
	}

	newRelease, _, err = r.client.Repositories.CreateRelease(ctx, targetOwner, targetRepo, newRelease)
}
