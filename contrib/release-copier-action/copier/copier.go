package copier

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// ReleaseCopier contains the release copier client.
type ReleaseCopier struct {
	client                                           *github.Client
	sourceOwner, sourceRepo, targetOwner, targetRepo string
	// mux ensures only one copy can be made at a time
	mux sync.Mutex
}

// NewReleaseCopier creates a new release copier client.
func NewReleaseCopier(ctx context.Context, token string) *ReleaseCopier {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})

	tc := oauth2.NewClient(ctx, ts)

	return &ReleaseCopier{
		client: github.NewClient(tc),
	}
}

// CopyRelease copies a release from a sourceRepo to a targetRepo and strips a prefix.
// nolint: cyclop
func (r *ReleaseCopier) CopyRelease(ctx context.Context, sourceOwner, sourceRepo, targetOwner, targetRepo, tagName, stripPrefix string) error {
	// make sure only one copy process runs at a time
	r.mux.Lock()
	defer r.mux.Unlock()

	r.sourceRepo = sourceRepo
	r.sourceOwner = sourceOwner
	r.targetOwner = targetOwner
	r.targetRepo = targetRepo

	// Get the release
	ogRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, sourceOwner, sourceRepo, tagName)
	if err != nil {
		return fmt.Errorf("could not get origin release: %w", err)
	}

	// Get the tag for the release
	ogTag, err := r.GetTagForRelease(ctx, ogRelease)
	if err != nil {
		return fmt.Errorf("could not get tag for release: %w", err)
	}

	if ogRelease.TagName == nil {
		return errors.New("could not get origin release tag name, tag is required for copying a release")
	}

	strippedTag := strings.TrimPrefix(tagName, stripPrefix)
	// releaes name is optional, so we only set it if it exists
	// we also strip the prefix here
	name := ""
	if ogRelease.Name != nil {
		name = strings.TrimPrefix(*ogRelease.Name, stripPrefix)
	}

	commits, _, err := r.client.Repositories.ListCommits(ctx, targetOwner, targetRepo, &github.CommitsListOptions{})
	if err != nil {
		return fmt.Errorf("could not get commits for repo %s/%s: %w", targetOwner, targetRepo, err)
	}
	if len(commits) == 0 {
		return errors.New("at least one commit is required for a release")
	}

	// Create the tag
	tag, _, err := r.client.Git.CreateTag(ctx, targetOwner, targetRepo, &github.Tag{
		Tag:     &strippedTag,
		Message: ogTag.Commit.Message,
		Object: &github.GitObject{
			Type: github.String("commit"),
			SHA:  commits[0].SHA,
			URL:  commits[0].URL,
		},
	})
	if err != nil {
		return fmt.Errorf("could not create tag %s: %w", strippedTag, err)
	}

	// Create the release
	newRelease := &github.RepositoryRelease{
		TagName: tag.Tag,
		Name:    &name,
		Body:    ogRelease.Body,
	}

	newRelease, _, err = r.client.Repositories.CreateRelease(ctx, targetOwner, targetRepo, newRelease)
	if err != nil {
		return fmt.Errorf("could not create release: %w", err)
	}

	for _, asset := range ogRelease.Assets {
		err = r.copyReleaseAsset(ctx, asset, newRelease)
		if err != nil {
			return fmt.Errorf("could not copy release asset: %w", err)
		}
	}
	return nil
}

// GetTagForRelease gets the tag for a given release. It does this by iterating through all tags to find
// a matching tag name. We need to do this because the github api does not return the tag by name
// only sha.
func (r *ReleaseCopier) GetTagForRelease(ctx context.Context, release *github.RepositoryRelease) (*github.RepositoryTag, error) {
	page := 1
	for {
		tags, res, err := r.client.Repositories.ListTags(ctx, r.sourceOwner, r.sourceRepo, &github.ListOptions{
			PerPage: 100,
			Page:    page,
		})
		if err != nil {
			return nil, fmt.Errorf("could not get tags: %w", err)
		}

		for _, tag := range tags {
			if tag.GetName() == release.GetTagName() {
				// some fields aren't populated in this response, so we need to get the full tag commit metadata
				commit, _, err := r.client.Git.GetCommit(ctx, r.sourceOwner, r.sourceRepo, tag.GetCommit().GetSHA())
				if err != nil {
					return nil, fmt.Errorf("could not get commit: %w", err)
				}

				tag.Commit = commit

				return tag, nil
			}
		}

		if page == res.LastPage {
			return nil, fmt.Errorf("could not find tag for release %s/%s", r.sourceOwner, r.sourceRepo)
		}
		page = res.NextPage
	}
}

// copyReleaseAsset copies a release asset from the source repo to the target repo.
func (r *ReleaseCopier) copyReleaseAsset(ctx context.Context, asset *github.ReleaseAsset, targetRelease *github.RepositoryRelease) error {
	// Download the original release asset
	reader, _, err := r.client.Repositories.DownloadReleaseAsset(ctx, r.sourceOwner, r.sourceRepo, *asset.ID, http.DefaultClient)
	if err != nil {
		return fmt.Errorf("could not download asset %s: %w", *asset.Name, err)
	}

	folderName, err := os.MkdirTemp("", "release-copier")
	if err != nil {
		return fmt.Errorf("could not create temp folder: %w", err)
	}

	// create a temp file, we'll use this to store the contents of the original
	fileName := fmt.Sprintf("%s/%s", folderName, *asset.Name)

	//nolint: gosec
	tmpFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("could not create temp file %s: %w", fileName, err)
	}

	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()

	toWrite, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("could not read asset %s: %w", *asset.Name, err)
	}

	_, err = tmpFile.Write(toWrite)
	if err != nil {
		return fmt.Errorf("could not write asset %s: %w", *asset.Name, err)
	}

	// release the file handle
	_ = tmpFile.Close()

	// open as readonly
	//nolint:gosec
	tmpFile, err = os.Open(fileName)
	if err != nil {
		return fmt.Errorf("could not open temp file %s: %w", fileName, err)
	}

	// Upload the resulting release asset
	_, res, err := r.client.Repositories.UploadReleaseAsset(ctx, r.targetOwner, r.targetRepo, targetRelease.GetID(), &github.UploadOptions{
		Name:  asset.GetName(),
		Label: asset.GetLabel(),
	}, tmpFile)

	if err != nil {
		return fmt.Errorf("could not upload: %w", err)
	}

	_ = res

	return nil
}
