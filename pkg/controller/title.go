package controller

import (
	"fmt"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/config"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

const defaultIssueTitleTemplate = `Renovate Automerge Failure({{.RepoOwner}}/{{.RepoName}}): {{if .Metadata.GroupName}}{{.Metadata.GroupName}}{{else}}{{.Metadata.PackageName}}{{.Metadata.DepName}}{{end}} {{if .Metadata.PackageFileDir}}({{.Metadata.PackageFileDir}}){{end}}`

type TitleParam struct {
	RepoOwner string
	RepoName  string
	Metadata  *Metadata
}

func getIssueTitle(cfg *config.Config, repoOwner, repoName string, metadata *Metadata) (string, error) {
	s, err := template.Render(defaultIssueTitleTemplate, &TitleParam{
		RepoOwner: repoOwner,
		RepoName:  repoName,
		Metadata:  metadata,
	})
	if err != nil {
		return "", fmt.Errorf("render the template of issue title: %w", err)
	}
	return s, nil
}
