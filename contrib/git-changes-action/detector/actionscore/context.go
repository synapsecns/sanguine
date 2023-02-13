package actionscore

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v41/github"
)

// WebhookPayload represents the Github webhook payload.
type WebhookPayload struct {
	Repository   *github.Repository   `json:"repository,omitempty"`
	Issue        *github.Issue        `json:"issue,omitempty"`
	PR           *github.PullRequest  `json:"pull_request,omitempty"`
	Sender       *github.User         `json:"sender,omitempty"`
	Action       string               `json:"action,omitempty"`
	Installation *github.Installation `json:"installation,omitempty"`
	Comment      *github.IssueComment `json:"comment,omitempty"`
}

// Context represents the context of the Github action.
type Context struct {
	Payload    WebhookPayload
	EventName  string
	SHA        string
	Ref        string
	Workflow   string
	Action     string
	Actor      string
	Job        string
	RunNumber  int
	RunID      int
	APIURL     string
	ServerURL  string
	GraphQLURL string
}

// NewContext creates a new context.
func NewContext() *Context {
	c := &Context{
		APIURL:     "https://api.github.com",
		ServerURL:  "https://github.com",
		GraphQLURL: "https://api.github.com/graphql",
	}

	if os.Getenv("GITHUB_EVENT_PATH") != "" {
		file, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
		if err != nil {
			return c
		}
		defer func() {
			_ = file.Close()
		}()

		jsonParser := json.NewDecoder(file)
		err = jsonParser.Decode(&c.Payload)
		if err != nil {
			return c
		}
	}

	c.EventName = os.Getenv("GITHUB_EVENT_NAME")
	c.SHA = os.Getenv("GITHUB_SHA")
	c.Ref = os.Getenv("GITHUB_REF")
	c.Workflow = os.Getenv("GITHUB_WORKFLOW")
	c.Action = os.Getenv("GITHUB_ACTION")
	c.Actor = os.Getenv("GITHUB_ACTOR")
	c.Job = os.Getenv("GITHUB_JOB")
	c.RunNumber = 0
	c.RunID = 0
	if os.Getenv("GITHUB_RUN_NUMBER") != "" {
		c.RunNumber, _ = strconv.Atoi(os.Getenv("GITHUB_RUN_NUMBER"))
	}
	if os.Getenv("GITHUB_RUN_ID") != "" {
		c.RunID, _ = strconv.Atoi(os.Getenv("GITHUB_RUN_ID"))
	}
	if os.Getenv("GITHUB_API_URL") != "" {
		c.APIURL = os.Getenv("GITHUB_API_URL")
	}
	if os.Getenv("GITHUB_SERVER_URL") != "" {
		c.ServerURL = os.Getenv("GITHUB_SERVER_URL")
	}
	if os.Getenv("GITHUB_GRAPHQL_URL") != "" {
		c.GraphQLURL = os.Getenv("GITHUB_GRAPHQL_URL")
	}

	return c
}

// Issue returns the issue information.
func (c *Context) Issue() (owner, repo string, number int) {
	if os.Getenv("GITHUB_REPOSITORY") != "" {
		parts := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
		owner = parts[0]
		repo = parts[1]
	} else if c.Payload.Repository != nil {
		user := c.Payload.Repository.GetOwner()
		if user != nil {
			owner = user.GetName()
		}
		repo = c.Payload.Repository.GetName()
	}
	if c.Payload.Issue != nil {
		number = c.Payload.Issue.GetNumber()
	} else if c.Payload.PR != nil {
		number = c.Payload.PR.GetNumber()
	}
	return
}

// Repo returns the repository information.
func (c *Context) Repo() (owner, repo string) {
	if os.Getenv("GITHUB_REPOSITORY") != "" {
		parts := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
		owner = parts[0]
		repo = parts[1]
		return
	}
	if c.Payload.Repository != nil {
		user := c.Payload.Repository.GetOwner()
		if user != nil {
			owner = user.GetName()
		}
		repo = c.Payload.Repository.GetName()
		return
	}
	panic("Context.Repo requires a GITHUB_REPOSITORY environment variable like 'owner/repo'")
}
