package expr

import (
	"errors"
	"fmt"
	"os"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
)

var (
	errMustBeBool        = errors.New("must be a boolean")
	errMustBeStringSlice = errors.New("must be a string slice")
)

type Param struct {
	Metadata *domain.Metadata
	Labels   []string
	Vars     map[string]interface{}
}

func compileBool(s string) (*vm.Program, error) {
	a, err := expr.Compile(s, expr.AsBool(), expr.Env(&Param{}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] compile an expression\n%+v", err)
		return nil, fmt.Errorf("compile an expression: %w", err)
	}
	return a, nil
}

func runBool(prog *vm.Program, param *Param) (bool, error) {
	a, err := run(prog, param)
	if err != nil {
		return false, err
	}
	f, ok := a.(bool)
	if !ok {
		return false, errMustBeBool
	}
	return f, nil
}

func ExecBool(s string, param *Param) (bool, error) {
	prog, err := compileBool(s)
	if err != nil {
		return false, err
	}
	return runBool(prog, param)
}

func ExecStringSlice(s string, param *Param) ([]string, error) {
	prog, err := compile(s)
	if err != nil {
		return nil, err
	}
	return runStringSlice(prog, param)
}

func compile(s string) (*vm.Program, error) {
	a, err := expr.Compile(s, expr.Env(&Param{}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] compile an expression\n%+v", err)
		return nil, fmt.Errorf("compile an expression: %w", err)
	}
	return a, nil
}

func run(prog *vm.Program, param *Param) (interface{}, error) {
	a, err := expr.Run(prog, param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] evaluate an expression\n%+v", err)
		return false, fmt.Errorf("evaluate an expression: %w", err)
	}
	return a, nil
}

func runStringSlice(prog *vm.Program, param *Param) ([]string, error) {
	a, err := run(prog, param)
	if err != nil {
		return nil, err
	}
	s, ok := a.([]string)
	if !ok {
		return nil, errMustBeStringSlice
	}
	return s, nil
}
