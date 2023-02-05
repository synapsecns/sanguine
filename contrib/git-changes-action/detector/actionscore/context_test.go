// nolint: goconst
package actionscore_test

import (
	"github.com/Flaque/filet"
	"github.com/google/go-github/v41/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/actionscore"
	"os"
	"strings"
	"testing"
)

func unsetGithubRepo(tb testing.TB) {
	tb.Helper()
	// unset all github env vars
	for _, osVar := range os.Environ() {
		splitVar := strings.Split(osVar, "=")
		key := splitVar[0]

		if strings.HasPrefix(key, "GITHUB_") {
			tb.Setenv(key, "")
		}
	}
}

// nolint: gocognit, cyclop
func TestNewContext(t *testing.T) {
	unsetGithubRepo(t)
	t.Run("with GITHUB_EVENT_PATH set", func(t *testing.T) {
		res := mock.MustMarshal(github.PushEvent{
			Repo: &github.PushEventRepository{
				Name: github.String("test-repo"),
			},
		})
		eventPath := filet.TmpFile(t, "", string(res))
		t.Setenv("GITHUB_EVENT_PATH", eventPath.Name())

		c := actionscore.NewContext()
		if c.Payload.Repository == nil {
			t.Error("Expected Payload.Repository to be set, got nil")
		}
	})

	t.Run("with GITHUB_EVENT_PATH not set", func(t *testing.T) {
		c := actionscore.NewContext()
		if c.Payload.Repository != nil {
			t.Error("Expected Payload.Repository to be nil, got non-nil")
		}
		if c.Payload.Issue != nil {
			t.Error("Expected Payload.Issue to be nil, got non-nil")
		}
	})

	t.Run("with GITHUB_EVENT_NAME set", func(t *testing.T) {
		t.Setenv("GITHUB_EVENT_NAME", "push")

		c := actionscore.NewContext()
		if c.EventName != "push" {
			t.Errorf("Expected EventName to be 'push', got %s", c.EventName)
		}
	})

	t.Run("with GITHUB_SHA set", func(t *testing.T) {
		t.Setenv("GITHUB_SHA", "abc123")

		c := actionscore.NewContext()
		if c.SHA != "abc123" {
			t.Errorf("Expected SHA to be 'abc123', got %s", c.SHA)
		}
	})

	t.Run("with GITHUB_REF set", func(t *testing.T) {
		t.Setenv("GITHUB_REF", "refs/heads/main")

		c := actionscore.NewContext()
		if c.Ref != "refs/heads/main" {
			t.Errorf("Expected Ref to be 'refs/heads/main', got %s", c.Ref)
		}
	})

	t.Run("with GITHUB_WORKFLOW set", func(t *testing.T) {
		t.Setenv("GITHUB_WORKFLOW", "main")

		c := actionscore.NewContext()
		if c.Workflow != "main" {
			t.Errorf("Expected Workflow to be 'main', got %s", c.Workflow)
		}
	})

	t.Run("GITHUB_ACTION is set", func(t *testing.T) {
		t.Setenv("GITHUB_ACTION", "test")
		c := actionscore.NewContext()
		if c.Action != "test" {
			t.Errorf("Expected action to be 'test', but got '%s'", c.Action)
		}
	})

	t.Run("GITHUB_ACTOR is set", func(t *testing.T) {
		t.Setenv("GITHUB_ACTOR", "test")
		c := actionscore.NewContext()
		if c.Actor != "test" {
			t.Errorf("Expected actor to be 'test', but got '%s'", c.Actor)
		}
	})

	t.Run("GITHUB_ACTOR is not set", func(t *testing.T) {
		t.Setenv("GITHUB_ACTOR", "")
		c := actionscore.NewContext()
		if c.Actor != "" {
			t.Errorf("Expected actor to be empty, but got '%s'", c.Actor)
		}
	})

	t.Run("GITHUB_JOB is set", func(t *testing.T) {
		t.Setenv("GITHUB_JOB", "test")
		c := actionscore.NewContext()
		if c.Job != "test" {
			t.Errorf("Expected job to be 'test', but got '%s'", c.Job)
		}
	})

	t.Run("GITHUB_JOB is not set", func(t *testing.T) {
		t.Setenv("GITHUB_JOB", "")
		c := actionscore.NewContext()
		if c.Job != "" {
			t.Errorf("Expected job to be empty, but got '%s'", c.Job)
		}
	})

	t.Run("GITHUB_RUN_NUMBER is not set", func(t *testing.T) {
		c := actionscore.NewContext()
		if c.RunNumber != 0 {
			t.Errorf("expected RunNumber to be 0, but got %d", c.RunNumber)
		}
	})

	t.Run("GITHUB_RUN_NUMBER is set", func(t *testing.T) {
		t.Setenv("GITHUB_RUN_NUMBER", "123")
		c := actionscore.NewContext()
		if c.RunNumber != 123 {
			t.Errorf("expected RunNumber to be 123, but got %d", c.RunNumber)
		}
	})

	t.Run("GITHUB_RUN_ID is not set", func(t *testing.T) {
		c := actionscore.NewContext()
		if c.RunID != 0 {
			t.Errorf("expected RunID to be 0, but got %d", c.RunID)
		}
	})

	t.Run("GITHUB_RUN_ID is set", func(t *testing.T) {
		t.Setenv("GITHUB_RUN_ID", "456")
		c := actionscore.NewContext()
		if c.RunID != 456 {
			t.Errorf("expected RunID to be 456, but got %d", c.RunID)
		}
	})

	t.Run("GITHUB_API_URL is not set", func(t *testing.T) {
		c := actionscore.NewContext()
		if c.APIURL != "https://api.github.com" {
			t.Errorf("expected APIURL to be https://api.github.com, but got %s", c.APIURL)
		}
	})

	t.Run("GITHUB_API_URL is set", func(t *testing.T) {
		t.Setenv("GITHUB_API_URL", "https://my-api.github.com")
		c := actionscore.NewContext()
		if c.APIURL != "https://my-api.github.com" {
			t.Errorf("expected APIURL to be https://my-api.github.com, but got %s", c.APIURL)
		}
	})

	t.Run("GITHUB_SERVER_URL is not set", func(t *testing.T) {
		c := actionscore.NewContext()
		if c.ServerURL != "https://github.com" {
			t.Errorf("expected APIURL to be https://github.com, but got %s", c.APIURL)
		}
	})

	t.Run("GITHUB_SERVER_URL is set", func(t *testing.T) {
		t.Setenv("GITHUB_SERVER_URL", "https://my-github.com")
		c := actionscore.NewContext()
		if c.ServerURL != "https://my-github.com" {
			t.Errorf("expected APIURL to be https://my-github.com, but got %s", c.APIURL)
		}
	})

	t.Run("GITHUB_GRAPHQL_URL is not set", func(t *testing.T) {
		c := actionscore.NewContext()
		if c.GraphQLURL != "https://api.github.com/graphql" {
			t.Errorf("expected APIURL to be https://api.github.com/graphql, but got %s", c.APIURL)
		}
	})

	t.Run("GITHUB_GRAPHQL_URL is set", func(t *testing.T) {
		t.Setenv("GITHUB_GRAPHQL_URL", "https://my-github.com/graphql")
		c := actionscore.NewContext()
		if c.GraphQLURL != "https://my-github.com/graphql" {
			t.Errorf("expected APIURL to be https://my-github.com/graphql, but got %s", c.APIURL)
		}
	})
}

func TestIssue(t *testing.T) {
	unsetGithubRepo(t)
	// Test with issue
	t.Run("Test issue", func(t *testing.T) {
		issue := &github.Issue{Number: github.Int(1)}
		c := &actionscore.Context{Payload: actionscore.WebhookPayload{Issue: issue}}
		owner, repo, number := c.Issue()
		assert.Equal(t, "", owner)
		assert.Equal(t, "", repo)
		assert.Equal(t, 1, number)
	})

	// Test with pull request
	t.Run("Test pull request", func(t *testing.T) {
		pr := &github.PullRequest{Number: github.Int(2)}
		c := &actionscore.Context{Payload: actionscore.WebhookPayload{PR: pr}}
		owner, repo, number := c.Issue()
		assert.Equal(t, "", owner)
		assert.Equal(t, "", repo)
		assert.Equal(t, 2, number)
	})

	// Test with GITHUB_REPOSITORY
	t.Run("Test GITHUB_REPOSITORY", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "owner/repo")
		c := &actionscore.Context{}
		owner, repo, number := c.Issue()
		assert.Equal(t, "owner", owner)
		assert.Equal(t, "repo", repo)
		assert.Equal(t, 0, number)
	})
}
