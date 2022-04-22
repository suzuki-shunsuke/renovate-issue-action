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

type RunParam struct {
	GitHubEventPath string
	GitHubActor     string
	GitHubRunID     string
}

func readPayload(p string, ev *github.PullRequestEvent) error {
	f, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("open GITHUB_EVENT_PATH: %w", err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(ev); err != nil {
		return fmt.Errorf("parse payload as JSON: %w", err)
	}
	return nil
}

func (ctrl *Controller) Run(ctx context.Context, logger *zap.Logger, param *RunParam) error {
	ev := &github.PullRequestEvent{}
	if err := readPayload(param.GitHubEventPath, ev); err != nil {
		return err
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
	title, err := getIssueTitle(cfg, repoOwner, repoName, metadata)
	if err != nil {
		return err
	}
	logger = logger.With(zap.String("issue_title", title))
	logger.Info("list issues")
	issues, err := ctrl.github.ListIssues(ctx, repoOwner, repoName, title)
	if err != nil {
		return fmt.Errorf("list issues: %w", err)
	}
	var issue *domain.Issue
	if len(issues) != 0 {
		logger.Info("issue already exists")
		issue = issues[0]
	}
	prURL := pr.GetHTMLURL()
	buildURL := "https://github.com/" + repoOwner + "/" + repoName + "/actions/runs/" + param.GitHubRunID
	if pr.GetMerged() {
		logger.Info("pr was merged")
		if err := ctrl.runMergedPR(ctx, logger, repoOwner, repoName, issue, prURL, buildURL); err != nil {
			return err
		}
		return nil
	}
	closedByRenovate := param.GitHubActor == "renovate[bot]"
	logger = logger.With(zap.Bool("closed_by_renovate", closedByRenovate))
	logger.Info("pr was closed")
	if err := ctrl.runUnmergedPR(ctx, logger, repoOwner, repoName, title, issue, metadata, prURL, closedByRenovate, buildURL); err != nil {
		return err
	}
	return nil
}

const closeIssueCommentTemplate = `This issue was closed by %s
This comment is created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action) [Build Link](%s)`

func (ctrl *Controller) runMergedPR(ctx context.Context, logger *zap.Logger, repoOwner, repoName string, issue *domain.Issue, prURL, buildURL string) error {
	if issue != nil {
		issueNumber := int(issue.Number)
		logger.Info("close an issue")
		cmt, err := ctrl.github.CreateComment(ctx, repoOwner, repoName, issueNumber, fmt.Sprintf(closeIssueCommentTemplate, prURL, buildURL))
		if err != nil {
			logger.Error("create an issue comment", zap.Error(err))
		} else {
			logger.Info("create an issue comment", zap.String("comment_url", cmt.GetHTMLURL()))
		}
		if err := ctrl.github.CloseIssue(ctx, repoOwner, repoName, int(issue.Number)); err != nil {
			return fmt.Errorf("close an issue: %w", err)
		}
		return nil
	}
	return nil
}

func (ctrl *Controller) runUnmergedPR(ctx context.Context, logger *zap.Logger, repoOwner, repoName, title string, issue *domain.Issue, metadata *Metadata, prURL string, closedByRenovate bool, buildURL string) error {
	issueNumber := int(issue.Number)
	body := string(issue.Body)
	if closedByRenovate {
		if issue != nil {
			logger.Info("close an issue")
			cmt, err := ctrl.github.CreateComment(ctx, repoOwner, repoName, issueNumber, fmt.Sprintf(closeIssueCommentTemplate, prURL, buildURL))
			if err != nil {
				logger.Error("create an issue comment", zap.Error(err))
			} else {
				logger.Info("create an issue comment", zap.String("comment_url", cmt.GetHTMLURL()))
			}
			if err := ctrl.github.CloseIssue(ctx, repoOwner, repoName, issueNumber); err != nil {
				return fmt.Errorf("close an issue: %w", err)
			}
			return nil
		}
		return nil
	}
	if issue != nil {
		// If issue is found, update issue description.
		if strings.Contains(body, prURL) {
			logger.Info("the pull request URL is already included in the issue, so skip updating the issue")
			return nil
		}
		logger.Info("update an issue")
		if err := ctrl.github.UpdateIssue(ctx, repoOwner, repoName, issueNumber, &github.IssueRequest{
			Body: github.String(body + `
* ` + prURL),
		}); err != nil {
			return fmt.Errorf("update an issue: %w", err)
		}
		return nil
	}

	logger.Info("render an issue body")
	body, err := renderBody(defaultIssueBodyTemplate, metadata)
	if err != nil {
		return err
	}
	body += `
* ` + prURL

	logger.Info("create an issue")
	is, err := ctrl.github.CreateIssue(ctx, repoOwner, repoName, &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
		// Labels    *[]string
		// Assignees *[]string
	})
	if err != nil {
		return fmt.Errorf("create an issue: %w", err)
	}
	logger.Info("created an issue", zap.Int("pr_number", is.GetNumber()))
	return nil
}

type Config struct{}
