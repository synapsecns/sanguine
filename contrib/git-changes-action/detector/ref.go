package detector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-github/v37/github"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"golang.org/x/exp/slices"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

// GetChangeTree returns the ref for the given event name.
// it is based on  https://github.com/dorny/paths-filter/blob/4067d885736b84de7c414f582ac45897079b0a78/src/main.ts#L36
func GetChangeTree(ctx context.Context, repoPath, eventName, ref, token string) (tree.Tree, error) {
	isPrEvent := slices.ContainsFunc([]EventType{EventPullRequest, EventPullRequestReview, EventPullRequestReviewComment, EventPullRequestTarget}, func(eventType EventType) bool {
		return strings.EqualFold(eventName, eventType.String())
	})

	if isPrEvent {
		return getChangedFilesFromAPI(ctx, token)
	}

	ct, err := getChangeTreeFromGit(repoPath, ref)
	if err != nil {
		return nil, fmt.Errorf("could not get change tree: %w", err)
	}
	return ct, nil
}

// nolint: cyclop
func getChangedFilesFromAPI(ctx context.Context, token string) (ct tree.Tree, err error) {
	var gpe github.PullRequestEvent
	f, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return nil, fmt.Errorf("could not open event path: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	if err := json.NewDecoder(f).Decode(&gpe); err != nil {
		return nil, fmt.Errorf("could not decode event: %w", err)
	}

	owner, err := getOwner(gpe.Repo)
	if err != nil {
		return nil, fmt.Errorf("could not get owner: %w", err)
	}

	name, err := getName(gpe.Repo)
	if err != nil {
		return nil, fmt.Errorf("could not get name: %w", err)
	}

	prNumber := gpe.GetPullRequest().GetNumber()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	ct = tree.NewTree()

	page := 1
	for {
		files, res, err := client.PullRequests.ListFiles(ctx, owner, name, prNumber, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return nil, fmt.Errorf("could not get files: %w", err)
		}

		for _, file := range files {
			if file.Filename != nil {
				ct.Add(file.GetFilename())
			}
			if file.PreviousFilename != nil {
				ct.Add(file.GetPreviousFilename())
			}
		}

		if page == res.LastPage {
			break
		}

		page = res.NextPage
	}

	return ct, nil
}

func getName(repo *github.Repository) (string, error) {
	if repo == nil {
		return "", errors.New("repository is nil")
	}
	if repo.GetName() == "" {
		return "", errors.New("repository name is empty")
	}

	return repo.GetName(), nil
}

// getOwner returns the owner of the repository.
func getOwner(repo *github.Repository) (string, error) {
	if repo == nil {
		return "", errors.New("repository is nil")
	}

	if repo.GetOwner() == nil {
		return "", errors.New("repository owner is nil")
	}

	if repo.GetOwner().GetName() != "" {
		return "", errors.New("repository owner name is empty")
	}

	return repo.GetOwner().GetName(), nil
}

// EventType is the type of github api event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType -linecomment
type EventType uint8

const (
	// EventPullRequest is a pull request event.
	EventPullRequest EventType = iota // pull_request
	// EventPush is a push event.
	EventPush // push
	// EventRelease is a release event.
	EventRelease // release
	// EventTag is a tag event.
	EventTag // tag
	// EventPullRequestReview is a pull request review event.
	EventPullRequestReview // pull_request_review
	// EventPullRequestReviewComment is a pull request review comment event.
	EventPullRequestReviewComment // pull_request_review_comment
	// EventPullRequestTarget is a pull request target event.
	EventPullRequestTarget // pull_request_target
	// EventIssueComment is an issue comment event.
	EventIssueComment // issue_comment
	// EventIssues is an issues event.
	EventIssues // issues
	// EventCreate is a create event.
	EventCreate // create
	// EventDelete is a delete event.
	EventDelete // delete
	// EventDeployment is a deployment event.
	EventDeployment // deployment
)
