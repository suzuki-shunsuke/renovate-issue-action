package cli

import (
	"context"
	"io"
	"os"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/controller"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/github"
	"go.uber.org/zap"
)

type Runner struct {
	Stdin     io.Reader
	Stdout    io.Writer
	Stderr    io.Writer
	LDFlags   *LDFlags
	LogConfig *zap.Config
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (runner *Runner) Run(ctx context.Context, logger *zap.Logger, args ...string) error {
	httpClient := github.NewHTTPClient(ctx, os.Getenv("GITHUB_TOKEN"))
	ctrl := controller.InitializeController(ctx, httpClient, runner.LogConfig)
	return ctrl.Run(ctx, logger, &controller.RunParam{ //nolint:wrapcheck
		GitHubEventPath: os.Getenv("GITHUB_EVENT_PATH"),
		GitHubActor:     os.Getenv("GITHUB_ACTOR"),
		GitHubRunID:     os.Getenv("GITHUB_RUN_ID"),
	})
}
