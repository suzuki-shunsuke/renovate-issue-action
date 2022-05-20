package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v43/github"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/config"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/expr"
	"github.com/suzuki-shunsuke/zap-error/logerr"
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
	ConfigFilePath  string
}

func (ctrl *Controller) Run(ctx context.Context, logger *zap.Logger, param *RunParam) error { //nolint:funlen,cyclop
	ev := &github.PullRequestEvent{}
	if err := readPayload(param.GitHubEventPath, ev); err != nil {
		return err
	}
	pr := ev.GetPullRequest()
	repo := ev.GetRepo()
	prRepo := &domain.Repo{
		Owner: repo.GetOwner().GetLogin(),
		Name:  repo.GetName(),
	}
	metadata := &domain.Metadata{}
	logger = logger.With(zap.String("repo_owner", prRepo.Owner), zap.String("repo_name", prRepo.Name))
	logger.Info("read metadata")
	if err := readMetadata(pr.GetBody(), metadata); err != nil {
		return err
	}
	cfg := &config.Config{}
	if p := config.Find(param.ConfigFilePath); p != "" {
		if err := config.Read(p, cfg); err != nil {
			return fmt.Errorf("read a configuration file: %w", logerr.WithFields(err, zap.String("configuration_file_path", p)))
		}
	}
	config.SetDefault(cfg, prRepo)
	logger.Info("select an entry")
	entry := selectEntry(logger, cfg.Entries, &expr.Param{
		Metadata: metadata,
		Labels:   getLabelsFromPR(pr),
	})
	if entry != nil {
		if entry.Ignore {
			logger.Info("do nothing because entry.ignore is true")
			return nil
		}
		cfg.Issue.Merge(entry.Issue)
	}

	logger.Info("get an issue title")
	title, err := getIssueTitle(cfg.Issue, prRepo, metadata)
	if err != nil {
		return err
	}
	logger = logger.With(zap.String("issue_title", title))
	logger.Info("list issues")
	issues, err := ctrl.github.ListIssues(ctx, cfg.Issue.RepoOwner, cfg.Issue.RepoName, title)
	if err != nil {
		return fmt.Errorf("list issues: %w", err)
	}
	var issue *domain.Issue
	if len(issues) != 0 {
		logger.Info("issue already exists")
		issue = issues[0]
	}
	prURL := pr.GetHTMLURL()
	buildURL := "https://github.com/" + prRepo.Owner + "/" + prRepo.Name + "/actions/runs/" + param.GitHubRunID
	if pr.GetMerged() {
		logger.Info("pr was merged")
		if err := ctrl.runMergedPR(ctx, logger, cfg.Issue.RepoOwner, cfg.Issue.RepoName, issue, prURL, buildURL); err != nil {
			return err
		}
		return nil
	}
	closedByRenovate := param.GitHubActor == cfg.RenovateLogin
	logger = logger.With(zap.Bool("closed_by_renovate", closedByRenovate))
	logger.Info("pr was closed")
	if err := ctrl.runUnmergedPR(ctx, logger, cfg, cfg.Issue.RepoOwner, cfg.Issue.RepoName, title, issue, metadata, prURL, closedByRenovate, buildURL); err != nil {
		return err
	}
	return nil
}

const closeIssueCommentTemplate = `This issue was closed by %s
This comment is created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action) [Build Link](%s)`

func (ctrl *Controller) runMergedPR(ctx context.Context, logger *zap.Logger, repoOwner, repoName string, issue *domain.Issue, prURL, buildURL string) error {
	if issue != nil {
		logger.Info("close an issue")
		cmt, err := ctrl.github.CreateComment(ctx, repoOwner, repoName, issue.Number, fmt.Sprintf(closeIssueCommentTemplate, prURL, buildURL))
		if err != nil {
			logger.Error("create an issue comment", zap.Error(err))
		} else {
			logger.Info("create an issue comment", zap.String("comment_url", cmt.GetHTMLURL()))
		}
		if err := ctrl.github.CloseIssue(ctx, repoOwner, repoName, issue.Number); err != nil {
			return fmt.Errorf("close an issue: %w", err)
		}
		return nil
	}
	return nil
}

func (ctrl *Controller) runUnmergedPR(ctx context.Context, logger *zap.Logger, cfg *config.Config, repoOwner, repoName, title string, issue *domain.Issue, metadata *domain.Metadata, prURL string, closedByRenovate bool, buildURL string) error { //nolint:cyclop
	if closedByRenovate {
		if issue == nil {
			return nil
		}
		logger.Info("close an issue")
		cmt, err := ctrl.github.CreateComment(ctx, repoOwner, repoName, issue.Number, fmt.Sprintf(closeIssueCommentTemplate, prURL, buildURL))
		if err != nil {
			logger.Error("create an issue comment", zap.Error(err))
		} else {
			logger.Info("create an issue comment", zap.String("comment_url", cmt.GetHTMLURL()))
		}
		if err := ctrl.github.CloseIssue(ctx, repoOwner, repoName, issue.Number); err != nil {
			return fmt.Errorf("close an issue: %w", err)
		}
		return nil
	}
	if issue != nil {
		// If issue is found, update issue description.
		if strings.Contains(issue.Body, prURL) {
			logger.Info("the pull request URL is already included in the issue, so skip updating the issue")
			return nil
		}
		logger.Info("update an issue")
		if err := ctrl.github.UpdateIssue(ctx, repoOwner, repoName, issue.Number, &github.IssueRequest{
			Body: github.String(issue.Body + "\n* " + prURL),
		}); err != nil {
			return fmt.Errorf("update an issue: %w", err)
		}
		return nil
	}

	logger.Info("render an issue body")
	body, err := renderBody(cfg.Issue.Description(), metadata)
	if err != nil {
		return err
	}
	body += "\n* " + prURL

	logger.Info("creating an issue")
	is, err := ctrl.github.CreateIssue(ctx, repoOwner, repoName, &github.IssueRequest{
		Title:     github.String(title),
		Body:      github.String(body),
		Labels:    domain.GetStringSlicePointer(append(cfg.Issue.Labels, cfg.Issue.AdditionalLabels...)),
		Assignees: domain.GetStringSlicePointer(append(cfg.Issue.Assignees, cfg.Issue.AdditionalAssignees...)),
	})
	if err != nil {
		return fmt.Errorf("create an issue: %w", err)
	}
	logger.Info("created an issue", zap.Int("pr_number", is.GetNumber()))
	if err := ctrl.addIssueToProjects(ctx, logger, cfg, is.GetNodeID()); err != nil {
		return err
	}
	return nil
}

func (ctrl *Controller) addIssueToProjects(ctx context.Context, logger *zap.Logger, cfg *config.Config, issueID string) error {
	projectMap := make(map[string]*config.Project, len(cfg.Projects))
	for _, project := range cfg.Projects {
		projectMap[project.Alias] = project
	}
	logger = logger.With(zap.String("issue_id", issueID))
	for _, projectName := range cfg.Issue.Projects {
		logger := logger.With(zap.String("project_name", projectName))
		project, ok := projectMap[projectName]
		if !ok {
			logger.Error("project name isn't found")
			continue
		}
		if project.ColumnID != "" {
			logger = logger.With(zap.String("project_column_id", project.ColumnID))
		}
		if project.NextID != "" {
			logger = logger.With(zap.String("project_next_id", project.NextID))
		}
		logger.Info("adding an issue to project")
		if project.ColumnID != "" {
			if err := ctrl.github.AddProjectCard(ctx, issueID, project.ColumnID); err != nil {
				logger.Error("add an issue to a project", zap.Error(err))
			} else {
				logger.Info("added an issue to project")
			}
			continue
		}
		if project.NextID != "" {
			if err := ctrl.github.AddProjectNextItem(ctx, issueID, project.NextID); err != nil {
				logger.Error("add an issue to a project", zap.Error(err))
			} else {
				logger.Info("added an issue to project")
			}
			continue
		}
	}
	return nil
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

func selectEntry(logger *zap.Logger, entries []*config.Entry, param *expr.Param) *config.Entry {
	for i, entry := range entries {
		vars := make(map[string]interface{}, len(entry.Vars))
		for _, v := range entry.Vars {
			vars[v.Name] = v.Value
		}
		f, err := expr.ExecBool(entry.If, &expr.Param{
			Metadata: param.Metadata,
			Labels:   param.Labels,
			Vars:     vars,
		})
		if err != nil {
			logger.Error("evaluate entry.if", zap.Int("entry_index", i), zap.Error(err))
			continue
		}
		if f {
			return entry
		}
	}
	return nil
}

func getLabelsFromPR(pr *github.PullRequest) []string {
	labels := make([]string, len(pr.Labels))
	for i, label := range pr.Labels {
		labels[i] = label.GetName()
	}
	return labels
}
