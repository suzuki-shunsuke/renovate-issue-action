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
	CreateComment(ctx context.Context, repoOwner, repoName string, issueNumber int, body string) (*github.IssueComment, error)
}

type Issue struct {
	Number int
	Body   string
}

type Repo struct {
	Owner string
	Name  string
}

type Metadata struct {
	PackageFileDir string `json:"packageFileDir"`
	PackageName    string `json:"packageName"`
	GroupName      string `json:"groupName"`
	DepName        string `json:"depName"`
	UpdateType     string `json:"updateType"`
	Manager        string `json:"manager"`
}

func (meta *Metadata) Name() string {
	if meta.GroupName != "" {
		return meta.GroupName
	}
	if meta.PackageName != "" {
		return meta.PackageName
	}
	if meta.DepName != "" {
		return meta.DepName
	}
	return ""
}
