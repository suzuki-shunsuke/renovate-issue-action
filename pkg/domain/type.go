package domain

import (
	"context"

	"github.com/google/go-github/v43/github"
)

type GitHub interface {
	ListIssues(ctx context.Context, repoOwner, repoName, title string) ([]*Issue, error)
	CloseIssue(ctx context.Context, repoOwner, repoName string, issueNumber int) error
	UpdateIssue(ctx context.Context, repoOwner, repoName string, issueNumber int, issue *github.IssueRequest) error
	CreateIssue(ctx context.Context, repoOwner, repoName string, issue *github.IssueRequest) (*github.Issue, error)
}

type Issue struct {
	Number int
	Body   string
}
