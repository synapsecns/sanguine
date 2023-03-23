package detector

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/google/go-github/v41/github"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/actionscore"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"golang.org/x/exp/slices"
	"golang.org/x/oauth2"
	"os"
	"strings"
	"time"
)

// GetChangeTree returns the ref for the given event name.
// it is based on  https://github.com/dorny/paths-filter/blob/4067d885736b84de7c414f582ac45897079b0a78/src/main.ts#L36
func GetChangeTree(ctx context.Context, repoPath, ref, token, base string) (tree.Tree, error) {
	ghContext := actionscore.NewContext()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	isPrEvent := slices.ContainsFunc([]EventType{EventPullRequest, EventPullRequestReview, EventPullRequestReviewComment, EventPullRequestTarget}, func(eventType EventType) bool {
		return strings.EqualFold(ghContext.EventName, eventType.String())
	})

	if isPrEvent {
		return getChangedFilesFromAPI(ctx, ghContext, client)
	}

	ct, err := getChangeTreeFromGit(repoPath, ghContext, ref, base)
	if err != nil {
		return nil, fmt.Errorf("could not get change tree: %w", err)
	}
	return ct, nil
}

// nolint: cyclop
func getChangedFilesFromAPI(ctx context.Context, ghContext *actionscore.Context, client *github.Client) (ct tree.Tree, err error) {
	var gpe github.PullRequestEvent
	// TODO: should rap into context
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

	repoOwner, repoName := ghContext.Repo()

	prNumber := gpe.GetPullRequest().GetNumber()

	ct = tree.NewTree()

	page := 1
	const retryCount = 10
	const perPage = 100
	for {
		var files []*github.CommitFile
		var res *github.Response
		err = retry.Do(func() error {
			reqCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
			defer cancel()

			files, res, err = client.PullRequests.ListFiles(reqCtx, repoOwner, repoName, prNumber, &github.ListOptions{
				Page:    page,
				PerPage: perPage,
			})
			if err != nil {
				return fmt.Errorf("could not get files for repoOwner %s, repoName %s, prNumber %d, page number %d with page size %d: %w",
					repoOwner, repoName, prNumber, page, perPage, err)
			}
			return nil
		}, retry.Context(ctx), retry.Attempts(retryCount))
		if err != nil {
			return nil, fmt.Errorf("could not get files after %d retries: %w", retryCount, err)
		}

		for _, file := range files {
			if file.Filename != nil {
				ct.Add(file.GetFilename())
			}
			if file.PreviousFilename != nil {
				ct.Add(file.GetPreviousFilename())
			}
		}

		if res.NextPage == 0 {
			break
		}

		page = res.NextPage
	}

	return ct, nil
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
