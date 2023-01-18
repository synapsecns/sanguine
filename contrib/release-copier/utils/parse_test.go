package utils_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/releasecopier/utils"
	"testing"
)

func TestParseGithubRepository(t *testing.T) {
	testCase := "octocat/Hello-World"
	repoOwner, repoName := utils.ParseGithubRepository(testCase)
	Equal(t, repoOwner, "octocat")
	Equal(t, repoName, "Hello-World")
}
