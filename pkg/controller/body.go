package controller

import (
	"fmt"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

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
