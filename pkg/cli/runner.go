package cli

import (
	"context"
	"flag"
	"fmt"
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

func (flags *LDFlags) ShowVersion() string {
	if flags.Version == "" {
		if flags.Commit == "" {
			return ""
		}
		return flags.Commit
	}
	return flags.Version + " (" + flags.Commit + ")"
}

func parseFlag(param *controller.RunParam, verFlag, helpFlag *bool) {
	flag.StringVar(&param.ConfigFilePath, "config", "", "configuration file path")
	flag.BoolVar(verFlag, "version", false, "show renovate-issue-action's version")
	flag.BoolVar(helpFlag, "help", false, "show the help message")
	flag.Parse()
}

const helpMsg = `NAME:
   renovate-issue-action - Create, update, and close GitHub Issues with GitHub Actions according to Renovate Pull Requests

USAGE:
   renovate-issue-action [--help] [--version] [--config <configuration file path>]

VERSION:
   %s

OPTIONS:
   --config value   configuration file path [$AQUA_CONFIG]
   --help           show help (default: false)
   --version        print the version (default: false)`

func (runner *Runner) Run(ctx context.Context, logger *zap.Logger, args ...string) error {
	httpClient := github.NewHTTPClient(ctx, os.Getenv("GITHUB_TOKEN"))
	ctrl := controller.InitializeController(ctx, httpClient, runner.LogConfig)
	param := &controller.RunParam{
		GitHubEventPath: os.Getenv("GITHUB_EVENT_PATH"),
		GitHubActor:     os.Getenv("GITHUB_ACTOR"),
		GitHubRunID:     os.Getenv("GITHUB_RUN_ID"),
	}
	verFlag, helpFlag := false, false
	parseFlag(param, &verFlag, &helpFlag)
	if helpFlag {
		fmt.Fprintf(runner.Stdout, helpMsg, runner.LDFlags.ShowVersion())
		return nil
	}
	if verFlag {
		fmt.Fprintln(runner.Stdout, runner.LDFlags.ShowVersion())
		return nil
	}
	return ctrl.Run(ctx, logger, param) //nolint:wrapcheck
}
