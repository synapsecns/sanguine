package copier_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-github/v41/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/release-copier-action/copier"
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
