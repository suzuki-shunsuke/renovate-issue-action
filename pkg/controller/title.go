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
	Metadata  *domain.Metadata
}

func getIssueTitle(issue *config.Issue, repo *domain.Repo, metadata *domain.Metadata) (string, error) {
	s, err := template.Render(issue.Title, &TitleParam{
		RepoOwner: repo.Owner,
		RepoName:  repo.Name,
		Metadata:  metadata,
	})
	if err != nil {
		return "", fmt.Errorf("render the template of issue title: %w", err)
	}
	return s, nil
}
