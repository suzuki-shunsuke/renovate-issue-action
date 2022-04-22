package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func Parse(s string) (*template.Template, error) {
	tpl, err := template.New("_").Funcs(sprig.TxtFuncMap()).Parse(s)
	if err != nil {
		return nil, fmt.Errorf("parse a string with Go's text/template: %w", err)
	}
	return tpl, nil
}

func Execute(tpl *template.Template, param interface{}) (string, error) {
	if tpl == nil {
		return "", nil
	}
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, param); err != nil {
		return "", fmt.Errorf("render a template: %w", err)
	}
	return buf.String(), nil
}

func Render(s string, param interface{}) (string, error) {
	tpl, err := Parse(s)
	if err != nil {
		return "", fmt.Errorf("parse a template: %w", err)
	}
	return Execute(tpl, param)
}
