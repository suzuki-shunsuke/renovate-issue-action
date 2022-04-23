package controller

import (
	"fmt"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/config"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

type TitleParam struct {
	RepoOwner string
	RepoName  string
	Metadata  *Metadata
}

func getIssueTitle(cfg *config.Config, repo *domain.Repo, metadata *Metadata) (string, error) {
	s, err := template.Render(cfg.Issue.Title, &TitleParam{
		RepoOwner: repo.Owner,
		RepoName:  repo.Name,
		Metadata:  metadata,
	})
	if err != nil {
		return "", fmt.Errorf("render the template of issue title: %w", err)
	}
	return s, nil
}
