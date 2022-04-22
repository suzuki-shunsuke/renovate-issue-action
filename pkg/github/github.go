package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v43/github"
	"github.com/shurcooL/githubv4"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
	"golang.org/x/oauth2"
)

func NewHTTPClient(ctx context.Context, token string) *http.Client {
	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
}

type client struct {
	github   *github.Client
	v4Client *githubv4.Client
}

func New(gh *github.Client, v4Client *githubv4.Client) domain.GitHub {
	return &client{
		github:   gh,
		v4Client: v4Client,
	}
}

func (cl *client) CloseIssue(ctx context.Context, repoOwner, repoName string, issueNumber int) error {
	_, _, err := cl.github.Issues.Edit(ctx, repoOwner, repoName, issueNumber, &github.IssueRequest{
		State: github.String("closed"),
	})
	if err != nil {
		return fmt.Errorf("close an issue by GitHub API v3: %w", err)
	}
	return nil
}

func (cl *client) UpdateIssue(ctx context.Context, repoOwner, repoName string, issueNumber int, issue *github.IssueRequest) error {
	_, _, err := cl.github.Issues.Edit(ctx, repoOwner, repoName, issueNumber, issue)
	if err != nil {
		return fmt.Errorf("update an issue by GitHub API v3: %w", err)
	}
	return nil
}

func (cl *client) CreateIssue(ctx context.Context, repoOwner, repoName string, issue *github.IssueRequest) (*github.Issue, error) {
	ret, _, err := cl.github.Issues.Create(ctx, repoOwner, repoName, issue)
	if err != nil {
		return nil, fmt.Errorf("create an issue by GitHub API v3: %w", err)
	}
	return ret, nil
}

func (cl *client) ListIssues(ctx context.Context, repoOwner, repoName, title string) ([]*domain.Issue, error) {
	var q struct {
		Search struct {
			Edges []struct {
				Node struct {
					Issue *domain.Issue `graphql:"... on Issue"`
				}
			}
		} `graphql:"search(first: 100, query: $searchQuery, type: $searchType)"`
	}
	variables := map[string]interface{}{
		"searchQuery": githubv4.String(fmt.Sprintf(`repo:%s/%s state:open "%s" in:title`, repoOwner, repoName, title)),
		"searchType":  githubv4.SearchTypeIssue,
	}

	if err := cl.v4Client.Query(ctx, &q, variables); err != nil {
		return nil, fmt.Errorf("get an issue by GitHub GraphQL API: %w", err)
	}
	issues := make([]*domain.Issue, len(q.Search.Edges))
	for i, edge := range q.Search.Edges {
		issues[i] = edge.Node.Issue
	}
	return issues, nil
}
