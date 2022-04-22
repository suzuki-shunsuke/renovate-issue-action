package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v43/github"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
	"go.uber.org/zap"
)

type Controller struct {
	github domain.GitHub
	logCfg *zap.Config
}

func New(gh domain.GitHub, logCfg *zap.Config) *Controller {
	return &Controller{
		github: gh,
		logCfg: logCfg,
	}
}

func (ctrl *Controller) Run(ctx context.Context, logger *zap.Logger) error {
	ev := &github.PullRequestEvent{}
	if err := json.Unmarshal([]byte(os.Getenv("EVENT_PAYLOAD")), ev); err != nil {
		return fmt.Errorf("unmarshal the event payload as JSON: %w", err)
	}
	pr := ev.GetPullRequest()
	repo := ev.GetRepo()
	repoOwner := repo.GetOwner().GetLogin()
	repoName := repo.GetName()
	metadata := &Metadata{}
	logger = logger.With(zap.String("repo_owner", repoOwner), zap.String("repo_name", repoName))
	logger.Info("read metadata")
	if err := readMetadata(pr.GetBody(), metadata); err != nil {
		return err
	}
	cfg := &Config{}
	logger.Info("get an issue title")
	title, err := getIssueTitle(cfg, metadata)
	if err != nil {
		return err
	}
	logger = logger.With(zap.String("issue_title", title))
	logger.Info("list issues")
	issues, err := ctrl.github.ListIssues(ctx, repoOwner, repoName, title)
	if err != nil {
		return err
	}
	var issue *domain.Issue
	if len(issues) != 0 {
		logger.Info("issue already exists")
		issue = issues[0]
	}
	if pr.GetMerged() {
		logger.Info("pr was merged")
		if err := ctrl.runMergedPR(ctx, logger, repoOwner, repoName, issue); err != nil {
			return err
		}
		return nil
	}
	logger.Info("pr was closed")
	if err := ctrl.runUnmergedPR(ctx, logger, repoOwner, repoName, title, issue, metadata, pr.GetURL()); err != nil {
		return err
	}
	return nil
}

func (ctrl *Controller) runMergedPR(ctx context.Context, logger *zap.Logger, repoOwner, repoName string, issue *domain.Issue) error {
	if issue != nil {
		logger.Info("close an issue")
		if err := ctrl.github.CloseIssue(ctx, repoOwner, repoName, issue.Number); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (ctrl *Controller) runUnmergedPR(ctx context.Context, logger *zap.Logger, repoOwner, repoName, title string, issue *domain.Issue, metadata *Metadata, prURL string) error {
	if issue != nil {
		// If issue is found, update issue description.
		if strings.Contains(issue.Body, prURL) {
			return nil
		}
		logger.Info("update an issue")
		if err := ctrl.github.UpdateIssue(ctx, repoOwner, repoName, issue.Number, &github.IssueRequest{
			Body: github.String(issue.Body + `
* ` + prURL),
		}); err != nil {
			return err
		}
		return nil
	}

	logger.Info("render an issue body")
	body, err := renderBody(defaultIssueBodyTemplate, metadata)
	if err != nil {
		return err
	}

	logger.Info("create an issue")
	is, err := ctrl.github.CreateIssue(ctx, repoOwner, repoName, &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
		// Labels    *[]string
		// Assignees *[]string
	})
	if err != nil {
		return err
	}
	logger.Info("created an issue", zap.Int("pr_number", is.GetNumber()))
	return nil
}

type Config struct{}
