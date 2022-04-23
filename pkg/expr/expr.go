package expr

import (
	"errors"
	"fmt"
	"os"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
)

var errMustBeBool = errors.New("must be a boolean")

type Param struct {
	Metadata *domain.Metadata
	Vars     []*Var
}

type Var struct {
	Name  string
	Type  string
	Value interface{}
	Vars  interface{}
}

func CompileBool(s string) (*vm.Program, error) {
	a, err := expr.Compile(s, expr.AsBool(), expr.Env(&Param{}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] compile an expression\n%+v", err)
		return nil, fmt.Errorf("compile an expression: %w", err)
	}
	return a, nil
}

func RunBool(prog *vm.Program, param *Param) (bool, error) {
	a, err := expr.Run(prog, param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] evaluate an expression\n%+v", err)
		return false, fmt.Errorf("evaluate an expression: %w", err)
	}
	f, ok := a.(bool)
	if !ok {
		return false, errMustBeBool
	}
	return f, nil
}
