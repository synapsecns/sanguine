package copier_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-github/v41/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/release-copier-action/copier"
	"io"
	"net/http"
)

func makeRepoTags(count int) []github.RepositoryTag {
	var tags []github.RepositoryTag
	for range count {
		tags = append(tags, github.RepositoryTag{
			Name: new(gofakeit.Name()),
		})
	}
	return tags
}

func (c *CopierSuite) TestGetTagsForRelease() {
	const targetTag = "v1.0.0"
	const targetCommit = "i-am-a-commit"
	const targetMessage = "i-am-a-message"

	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatchPages(
			mock.GetReposTagsByOwnerByRepo,
			makeRepoTags(50),
			makeRepoTags(50),
			[]github.RepositoryTag{
				{
					Name: new(targetTag),
					Commit: &github.Commit{
						SHA: new(targetCommit),
					},
				},
			},
		),
		mock.WithRequestMatch(
			mock.GetReposGitCommitsByOwnerByRepoByCommitSha,
			github.Commit{
				SHA:     new(targetCommit),
				Message: new(targetMessage),
			},
		))

	cp := copier.NewReleaseCopier(c.GetTestContext(), "")
	cp.SetSourceOwner("testowner")
	cp.SetSourceRepo("testrepo")

	cp.SetClient(github.NewClient(mockedHTTPClient))

	tag, err := cp.GetTagForRelease(c.GetTestContext(), &github.RepositoryRelease{
		TagName: new("v1.0.0"),
	})

	Nil(c.T(), err)
	Equal(c.T(), targetMessage, tag.GetCommit().GetMessage())
}

func (c *CopierSuite) TestCopyRelease() {
	const (
		sourceOwner  = "source-owner"
		sourceRepo   = "source-repo"
		targetOwner  = "target-owner"
		targetRepo   = "target-repo"
		sourceTag    = "prefix-v1.0.0"
		strippedTag  = "v1.0.0"
		sourceName   = "prefix-release v1.0.0"
		strippedName = "release v1.0.0"
		stripPrefix  = "prefix-"
		commitSHA    = "target-commit"
		commitURL    = "https://example.com/target-commit"
		tagMessage   = "release tag"
		assetName    = "release.zip"
		assetLabel   = "release asset"
		assetBody    = "asset contents"
	)

	assetID := int64(42)
	releaseID := int64(101)
	uploadedAsset := false

	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetReposReleasesTagsByOwnerByRepoByTag,
			github.RepositoryRelease{
				TagName: new(sourceTag),
				Name:    new(sourceName),
				Body:    new("release body"),
				Assets: []*github.ReleaseAsset{
					{
						ID:    new(assetID),
						Name:  new(assetName),
						Label: new(assetLabel),
					},
				},
			},
		),
		mock.WithRequestMatch(
			mock.GetReposTagsByOwnerByRepo,
			[]github.RepositoryTag{
				{
					Name: new(sourceTag),
					Commit: &github.Commit{
						SHA: new("source-commit"),
					},
				},
			},
		),
		mock.WithRequestMatch(
			mock.GetReposGitCommitsByOwnerByRepoByCommitSha,
			github.Commit{
				Message: new(tagMessage),
			},
		),
		mock.WithRequestMatch(
			mock.GetReposCommitsByOwnerByRepo,
			[]github.RepositoryCommit{
				{
					SHA: new(commitSHA),
					URL: new(commitURL),
				},
			},
		),
		mock.WithRequestMatchHandler(
			mock.PostReposGitTagsByOwnerByRepo,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requestBody, err := io.ReadAll(r.Body)
				Nil(c.T(), err)
				Contains(c.T(), string(requestBody), strippedTag)
				Contains(c.T(), string(requestBody), tagMessage)
				_, _ = w.Write(mock.MustMarshal(github.Tag{
					Tag: new(strippedTag),
				}))
			}),
		),
		mock.WithRequestMatchHandler(
			mock.PostReposReleasesByOwnerByRepo,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requestBody, err := io.ReadAll(r.Body)
				Nil(c.T(), err)
				Contains(c.T(), string(requestBody), strippedTag)
				Contains(c.T(), string(requestBody), strippedName)
				_, _ = w.Write(mock.MustMarshal(github.RepositoryRelease{
					ID:      new(releaseID),
					TagName: new(strippedTag),
					Name:    new(strippedName),
				}))
			}),
		),
		mock.WithRequestMatch(
			mock.GetReposReleasesAssetsByOwnerByRepoByAssetId,
			[]byte(assetBody),
		),
		mock.WithRequestMatchHandler(
			mock.PostReposReleasesAssetsByOwnerByRepoByReleaseId,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				uploadedAsset = true
				query := r.URL.Query()
				Equal(c.T(), assetName, query.Get("name"))
				Equal(c.T(), assetLabel, query.Get("label"))

				requestBody, err := io.ReadAll(r.Body)
				Nil(c.T(), err)
				Equal(c.T(), assetBody, string(requestBody))
				_, _ = w.Write(mock.MustMarshal(github.ReleaseAsset{
					ID: new(assetID),
				}))
			}),
		),
	)

	cp := copier.NewReleaseCopier(c.GetTestContext(), "")
	cp.SetClient(github.NewClient(mockedHTTPClient))

	err := cp.CopyRelease(c.GetTestContext(), sourceOwner, sourceRepo, targetOwner, targetRepo, sourceTag, stripPrefix)

	Nil(c.T(), err)
	True(c.T(), uploadedAsset)
}

func (c *CopierSuite) TestCopyReleaseRequiresTargetCommit() {
	const sourceTag = "v1.0.0"

	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetReposReleasesTagsByOwnerByRepoByTag,
			github.RepositoryRelease{
				TagName: new(sourceTag),
			},
		),
		mock.WithRequestMatch(
			mock.GetReposTagsByOwnerByRepo,
			[]github.RepositoryTag{
				{
					Name: new(sourceTag),
					Commit: &github.Commit{
						SHA: new("source-commit"),
					},
				},
			},
		),
		mock.WithRequestMatch(
			mock.GetReposGitCommitsByOwnerByRepoByCommitSha,
			github.Commit{
				Message: new("tag message"),
			},
		),
		mock.WithRequestMatch(
			mock.GetReposCommitsByOwnerByRepo,
			[]github.RepositoryCommit{},
		),
	)

	cp := copier.NewReleaseCopier(c.GetTestContext(), "")
	cp.SetClient(github.NewClient(mockedHTTPClient))

	err := cp.CopyRelease(c.GetTestContext(), "source-owner", "source-repo", "target-owner", "target-repo", sourceTag, "")

	Error(c.T(), err)
	Contains(c.T(), err.Error(), "at least one commit is required")
}
