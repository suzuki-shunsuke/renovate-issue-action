package controller

import (
	"fmt"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

const defaultIssueBodyTemplate = `
_This pull request was created by [renovate-issue-action](https://github.com/suzuki-shunsuke/renovate-issue-action)._

:warning: Please don't edit the Issue title, because renovate-issue-action searches issue with Issue title.

{{if .Metadata.PackageName}}packageName: {{.Metadata.PackageName}}{{end}}
{{if .Metadata.GroupName}}groupName: {{.Metadata.GroupName}}{{end}}
{{if .Metadata.DepName}}depName: {{.Metadata.DepName}}{{end}}

## Closed Pull Requests

`

type BodyParam struct {
	Metadata *Metadata
}

func renderBody(body string, metadata *Metadata) (string, error) {
	s, err := template.Render(body, &BodyParam{
		Metadata: metadata,
	})
	if err != nil {
		return "", fmt.Errorf("render the template of issue body: %w", err)
	}
	return s, nil
}
