package controller

import (
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

const defaultIssueBodyTemplate = `{{if .Metadata.PackageName}}packageName: {{.Metadata.PackageName}}{{end}}
{{if .Metadata.GroupName}}groupName: {{.Metadata.GroupName}}{{end}}
{{if .Metadata.DepName}}depName: {{.Metadata.DepName}}{{end}}

This pull request was created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action).`

type BodyParam struct {
	Metadata *Metadata
}

func renderBody(body string, metadata *Metadata) (string, error) {
	return template.Render(body, &BodyParam{
		Metadata: metadata,
	})
}
