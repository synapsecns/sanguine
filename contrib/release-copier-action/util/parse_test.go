package util_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/release-copier-action/util"
	"testing"
)

func TestParseGithubRepository(t *testing.T) {
	testCase := "octocat/Hello-World"
	repoOwner, repoName := util.ParseGithubRepository(testCase)
	Equal(t, repoOwner, "octocat")
	Equal(t, repoName, "Hello-World")
}
