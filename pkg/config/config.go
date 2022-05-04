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
	Vars   []*Var `json:"vars,omitempty"`
}

type Var struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type Issue struct {
	RepoOwner           string   `yaml:"repo_owner,omitempty"`
	RepoName            string   `yaml:"repo_name,omitempty"`
	Title               *string  `json:"title,omitempty"`
	DescriptionHeader   *string  `yaml:"description_header,omitempty"`
	DescriptionBody     *string  `yaml:"description_body,omitempty"`
	AdditionalBody      string   `yaml:"additional_body,omitempty"`
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
	issue.AdditionalBody = domain.JoinBody(issue.AdditionalBody, is.AdditionalBody)
}

func SetDefault(cfg *Config, repo *domain.Repo) {
	if cfg.RenovateLogin == "" {
		cfg.RenovateLogin = "renovate[bot]"
	}
	if cfg.Issue == nil {
		cfg.Issue = &Issue{}
	}
	if domain.GetString(cfg.Issue.Title) == "" {
		if cfg.Issue.Title == nil {
			s := ""
			cfg.Issue.Title = &s
		}
		*cfg.Issue.Title = defaultIssueTitleTemplate
	}
	if domain.GetString(cfg.Issue.DescriptionHeader) == "" {
		if cfg.Issue.DescriptionHeader == nil {
			s := ""
			cfg.Issue.DescriptionHeader = &s
		}
		*cfg.Issue.DescriptionHeader = defaultIssueDescriptionHeader
	}
	if cfg.Issue.RepoOwner == "" {
		cfg.Issue.RepoOwner = repo.Owner
	}
	if cfg.Issue.RepoName == "" {
		cfg.Issue.RepoName = repo.Name
	}
	if cfg.Issue.Labels == nil {
		cfg.Issue.Labels = []string{"renovate-issue-action"}
	}
}

func (issue *Issue) Description() string {
	return domain.JoinBody(domain.GetString(issue.DescriptionHeader), domain.GetString(issue.DescriptionBody), issue.AdditionalBody) + "\n## Closed Pull Requests\n"
}

const defaultIssueTitleTemplate = `renovate-issue-action ({{.RepoOwner}}/{{.RepoName}}): {{.Metadata.Name}}{{if .Metadata.PackageFileDir}} ({{.Metadata.PackageFileDir}}){{end}}{{if eq .Metadata.UpdateType "major"}} major{{end}}`

const defaultIssueDescriptionHeader = `
_This pull request was created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action)._

:warning: Please don't edit the Issue title, because renovate-issue-action searches issue with Issue title.

{{if .Metadata.PackageName}}packageName: {{.Metadata.PackageName}}{{end}}
{{if .Metadata.GroupName}}groupName: {{.Metadata.GroupName}}{{end}}
{{if .Metadata.DepName}}depName: {{.Metadata.DepName}}{{end}}
`
