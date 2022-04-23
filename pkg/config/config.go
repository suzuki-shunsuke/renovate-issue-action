package config

import (
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
)

type Config struct {
	RenovateLogin string `yaml:"renovate_login"`
	Issue         *Issue
}

type Issue struct {
	RepoOwner         string `yaml:"repo_owner"`
	RepoName          string `yaml:"repo_name"`
	Title             string
	DescriptionHeader string `yaml:"description_header"`
	DescriptionBody   string `yaml:"description_body"`
	Labels            []string
	Assignees         []string
	Projects          []string
}

const defaultIssueTitleTemplate = `Renovate Automerge Failure({{.RepoOwner}}/{{.RepoName}}): {{if .Metadata.GroupName}}{{.Metadata.GroupName}}{{else}}{{.Metadata.PackageName}}{{.Metadata.DepName}}{{end}} {{if .Metadata.PackageFileDir}}({{.Metadata.PackageFileDir}}){{end}}`

const defaultIssueDescriptionHeader = `
_This pull request was created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action)._

:warning: Please don't edit the Issue title, because renovate-issue-action searches issue with Issue title.

{{if .Metadata.PackageName}}packageName: {{.Metadata.PackageName}}{{end}}
{{if .Metadata.GroupName}}groupName: {{.Metadata.GroupName}}{{end}}
{{if .Metadata.DepName}}depName: {{.Metadata.DepName}}{{end}}
`

func SetDefault(cfg *Config, repo *domain.Repo) {
	if cfg.RenovateLogin == "" {
		cfg.RenovateLogin = "renovate[bot]"
	}
	if cfg.Issue.Title == "" {
		cfg.Issue.Title = defaultIssueTitleTemplate
	}
	if cfg.Issue.DescriptionHeader == "" {
		cfg.Issue.DescriptionHeader = defaultIssueDescriptionHeader
	}
	if cfg.Issue.RepoOwner == "" {
		cfg.Issue.RepoOwner = repo.Owner
	}
	if cfg.Issue.RepoName == "" {
		cfg.Issue.RepoName = repo.Name
	}
}
