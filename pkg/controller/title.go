package controller

import (
	"fmt"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/config"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

type TitleParam struct {
	RepoOwner string
	RepoName  string
	Metadata  *Metadata
}

func getIssueTitle(cfg *config.Config, repoOwner, repoName string, metadata *Metadata) (string, error) {
	s, err := template.Render(cfg.Issue.Title, &TitleParam{
		RepoOwner: repoOwner,
		RepoName:  repoName,
		Metadata:  metadata,
	})
	if err != nil {
		return "", fmt.Errorf("render the template of issue title: %w", err)
	}
	return s, nil
}
