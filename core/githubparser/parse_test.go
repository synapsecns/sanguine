package githubparser_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/githubparser"
	"testing"
)

func TestParseGithubRepository(t *testing.T) {
	testCase := "octocat/Hello-World"
	repoOwner, repoName := githubparser.ParseGithubRepository(testCase)
	Equal(t, repoOwner, "octocat")
	Equal(t, repoName, "Hello-World")
}
