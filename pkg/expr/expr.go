package expr

import (
	"errors"
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
)

var errMustBeBool = errors.New("must be a boolean")

func CompileBool(s string) (*vm.Program, error) {
	a, err := expr.Compile(s, expr.AsBool(), expr.Env(map[string]interface{}{
		"PackageFileDir": "",
		"PackageName":    "",
		"GroupName":      "",
		"DepName":        "",
		"UpdateType":     "",
		"Manager":        "",
	}))
	if err != nil {
		return nil, fmt.Errorf("compile an expression: %w", err)
	}
	return a, nil
}

func RunBool(prog *vm.Program, metadata *domain.Metadata) (bool, error) {
	a, err := expr.Run(prog, metadata)
	if err != nil {
		return false, fmt.Errorf("evaluate an expression: %w", err)
	}
	f, ok := a.(bool)
	if !ok {
		return false, errMustBeBool
	}
	return f, nil
}
