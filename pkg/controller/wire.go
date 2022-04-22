//go:build wireinject
// +build wireinject

package controller

import (
	"context"
	"net/http"

	"github.com/google/go-github/v43/github"
	"github.com/google/wire"
	"github.com/shurcooL/githubv4"
	gh "github.com/suzuki-shunsuke/renovate-issue-action/pkg/github"
	"go.uber.org/zap"
)

func InitializeController(ctx context.Context, httpClient *http.Client, logCfg *zap.Config) *Controller {
	wire.Build(New, gh.New, github.NewClient, githubv4.NewClient)
	return &Controller{}
}
