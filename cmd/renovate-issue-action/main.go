package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/cli"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/log"
	"github.com/suzuki-shunsuke/zap-error/logerr"
	"go.uber.org/zap"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	logCfg := log.NewConfig()
	logger, err := logCfg.Build()
	if err != nil {
		zap.L().Fatal("create a logger", zap.Error(err))
	}
	defer logger.Sync()
	core(logCfg, logger)
}

func core(logCfg *zap.Config, logger *zap.Logger) {
	runner := cli.Runner{
		Stdin:     os.Stdin,
		Stdout:    os.Stdout,
		Stderr:    os.Stderr,
		LogConfig: logCfg,
		LDFlags: &cli.LDFlags{
			Version: version,
			Commit:  commit,
			Date:    date,
		},
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := runner.Run(ctx, logger, os.Args...); err != nil {
		logger.Fatal("command failed", logerr.ToFields(err)...)
	}
}
