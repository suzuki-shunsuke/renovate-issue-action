package controller

import (
	"fmt"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

type BodyParam struct {
	Metadata *domain.Metadata
}

func renderBody(body string, metadata *domain.Metadata) (string, error) {
	s, err := template.Render(body, &BodyParam{
		Metadata: metadata,
	})
	if err != nil {
		return "", fmt.Errorf("render the template of issue body: %w", err)
	}
	return s, nil
}
