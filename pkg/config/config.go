package config

import (
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
)

type Config struct {
	RenovateLogin string   `yaml:"renovate_login,omitempty" jsonschema:"default=renovate[bot]"`
	Issue         *Issue   `json:"issue,omitempty"`
	Entries       []*Entry `json:"entries,omitempty"`
}

type Entry struct {
	Issue  *Issue `json:"issue"`
	If     string `json:"if"`
	Ignore bool   `json:"ignore,omitempty"`
}

type Issue struct {
	RepoOwner           string   `yaml:"repo_owner,omitempty"`
	RepoName            string   `yaml:"repo_name,omitempty"`
	Title               *string  `json:"title,omitempty"`
	DescriptionHeader   *string  `yaml:"description_header,omitempty"`
	DescriptionBody     *string  `yaml:"description_body,omitempty"`
	Labels              []string `json:"labels,omitempty"`
	AdditionalLabels    []string `yaml:"additional_labels,omitempty"`
	Assignees           []string `json:"assignees,omitempty"`
	AdditionalAssignees []string `yaml:"additional_assignees,omitempty"`
	// Projects            []string `json:"projects,omitempty"`
}

func (issue *Issue) Merge(is *Issue) {
	if is.RepoOwner != "" {
		issue.RepoOwner = is.RepoOwner
	}
	if is.RepoName != "" {
		issue.RepoName = is.RepoName
	}
	if is.Title != nil {
		issue.Title = is.Title
	}
	if is.DescriptionHeader != nil {
		issue.DescriptionHeader = is.DescriptionHeader
	}
	if is.DescriptionBody != nil {
		issue.DescriptionBody = is.DescriptionBody
	}
	if is.Labels != nil {
		issue.Labels = is.Labels
	}
	if is.Assignees != nil {
		issue.Assignees = is.Assignees
	}
	issue.AdditionalLabels = append(issue.AdditionalLabels, is.AdditionalLabels...)
	issue.AdditionalAssignees = append(issue.AdditionalAssignees, is.AdditionalAssignees...)
}

func (issue *Issue) Description() string {
	return domain.GetString(issue.DescriptionHeader) + `
` + domain.GetString(issue.DescriptionBody) + `
## Closed Pull Requests
`
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
	if domain.GetString(cfg.Issue.Title) == "" {
		*cfg.Issue.Title = defaultIssueTitleTemplate
	}
	if domain.GetString(cfg.Issue.DescriptionHeader) == "" {
		*cfg.Issue.DescriptionHeader = defaultIssueDescriptionHeader
	}
	if cfg.Issue.RepoOwner == "" {
		cfg.Issue.RepoOwner = repo.Owner
	}
	if cfg.Issue.RepoName == "" {
		cfg.Issue.RepoName = repo.Name
	}
}
