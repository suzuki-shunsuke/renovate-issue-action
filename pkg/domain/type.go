package domain

import (
	"context"

	"github.com/google/go-github/v43/github"
	"github.com/shurcooL/githubv4"
)

type GitHub interface {
	ListIssues(ctx context.Context, repoOwner, repoName, title string) ([]*Issue, error)
	CloseIssue(ctx context.Context, repoOwner, repoName string, issueNumber int) error
	UpdateIssue(ctx context.Context, repoOwner, repoName string, issueNumber int, issue *github.IssueRequest) error
	CreateIssue(ctx context.Context, repoOwner, repoName string, issue *github.IssueRequest) (*github.Issue, error)
}

type Issue struct {
	Number githubv4.Int
	Body   githubv4.String
}
