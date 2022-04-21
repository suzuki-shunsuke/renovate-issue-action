package controller

import (
	"context"
	"encoding/json"
	"os"

	"github.com/google/go-github/v43/github"
)

type Controller struct {
	github GitHub
}

type GitHub interface {
	GetIssue(ctx context.Context, repoOwner, repoName, title string) (*github.Issue, error)
}

func (ctrl *Controller) Run(ctx context.Context) error {
	ev := &github.PullRequestEvent{}
	if err := json.Unmarshal([]byte(os.Getenv("EVENT_PAYLOAD")), ev); err != nil {
		return err
	}
	pr := ev.GetPullRequest()
	repo := ev.GetRepo()
	repoOwner := repo.GetOwner().GetLogin()
	repoName := repo.GetName()
	metadata := &Metadata{}
	if err := readMetadata(pr.GetBody(), metadata); err != nil {
		return err
	}
	cfg := &Config{}
	title, err := getIssueTitle(cfg, metadata)
	if err != nil {
		return err
	}
	issue, err := ctrl.github.GetIssue(ctx, repoOwner, repoName, title)
	if err != nil {
		return err
	}
	if pr.GetMerged() {
		if err := ctrl.runMergedPR(ctx, issue); err != nil {
			return err
		}
		return nil
	}
	if err := ctrl.runUnmergedPR(ctx, issue); err != nil {
		return err
	}
	return nil
}

func (ctrl *Controller) runMergedPR(ctx context.Context, issue *github.Issue) error {
	if issue != nil {
		if err := ctrl.closeIssue(ctx, issue); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (ctrl *Controller) runUnmergedPR(ctx context.Context, issue *github.Issue) error {
	if issue != nil {
		// If issue is found, update issue description.
		if err := ctrl.updateIssue(ctx, issue); err != nil {
			return err
		}
		return nil
	}
	if err := ctrl.createIssue(ctx); err != nil {
		return err
	}
	return nil
}

func (ctrl *Controller) closeIssue(ctx context.Context, issue *github.Issue) error {
	return nil
}

func (ctrl *Controller) updateIssue(ctx context.Context, issue *github.Issue) error {
	return nil
}

func (ctrl *Controller) createIssue(ctx context.Context) error {
	return nil
}

func readMetadata(body string, metadata *Metadata) error {
	return nil
}

func getIssueTitle(cfg *Config, metadata *Metadata) (string, error) {
	return "", nil
}

type Config struct{}

type Metadata struct {
	PackageFileDir string `json:"packageFileDir"`
	PackageName    string `json:"packageName"`
	GroupName      string `json:"groupName"`
}
