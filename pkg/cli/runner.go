package cli

import (
	"context"
	"io"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/controller"
)

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	LDFlags *LDFlags
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (runner *Runner) Run(ctx context.Context, args ...string) error {
	ctrl := &controller.Controller{}
	return ctrl.Run(ctx)
}
