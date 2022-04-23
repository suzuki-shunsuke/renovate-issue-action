package config

type Config struct {
	RenovateLogin string `yaml:"renovate_login"`
	Issue         *Issue
}

type Issue struct {
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

func SetDefault(cfg *Config) {
	if cfg.RenovateLogin == "" {
		cfg.RenovateLogin = "renovate[bot]"
	}
	if cfg.Issue.Title == "" {
		cfg.Issue.Title = defaultIssueTitleTemplate
	}
	if cfg.Issue.DescriptionHeader == "" {
		cfg.Issue.DescriptionHeader = defaultIssueDescriptionHeader
	}
}
