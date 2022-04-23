package cli

import (
	"context"
	"flag"
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

func parseFlag(param *controller.RunParam) {
	flag.StringVar(&param.ConfigFilePath, "config", "", "configuration file path")
	flag.Parse()
}

func (runner *Runner) Run(ctx context.Context, logger *zap.Logger, args ...string) error {
	httpClient := github.NewHTTPClient(ctx, os.Getenv("GITHUB_TOKEN"))
	ctrl := controller.InitializeController(ctx, httpClient, runner.LogConfig)
	param := &controller.RunParam{
		GitHubEventPath: os.Getenv("GITHUB_EVENT_PATH"),
		GitHubActor:     os.Getenv("GITHUB_ACTOR"),
		GitHubRunID:     os.Getenv("GITHUB_RUN_ID"),
	}
	parseFlag(param)
	return ctrl.Run(ctx, logger, param) //nolint:wrapcheck
}
