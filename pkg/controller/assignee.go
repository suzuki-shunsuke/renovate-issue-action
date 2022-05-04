package controller

import (
	"strings"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/expr"
	"go.uber.org/zap"
)

func evaluateAssignees(logger *zap.Logger, param *expr.Param, assignees []string) []string {
	arr := make([]string, 0, len(assignees))
	for _, assignee := range assignees {
		if !strings.HasPrefix(assignee, "expr:") {
			arr = append(arr, assignee)
			continue
		}
		list, err := expr.ExecStringSlice(strings.TrimPrefix(assignee, "expr:"), param)
		if err != nil {
			logger.Error("evaluate an assignee", zap.String("invalid_assignee", assignee), zap.Error(err))
			continue
		}
		arr = append(arr, list...)
	}
	return arr
}
