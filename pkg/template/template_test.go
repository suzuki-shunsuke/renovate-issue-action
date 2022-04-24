package template_test

import (
	"testing"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/template"
)

func TestRender(t *testing.T) {
	t.Parallel()
	data := []struct {
		name  string
		isErr bool
		s     string
		param interface{}
		exp   string
	}{
		{
			name: "normal",
			s:    `{{trimPrefix "foo-" .Name}}`,
			param: map[string]interface{}{
				"Name": "foo-yoo",
			},
			exp: "yoo",
		},
		{
			name:  "err",
			s:     `{{trmPref "foo-" .Name}}`,
			isErr: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			s, err := template.Render(d.s, d.param)
			if err != nil {
				if !d.isErr {
					t.Fatal(err)
				}
				return
			}
			if s != d.exp {
				t.Fatalf("wanted %s, got %s", d.exp, s)
			}
		})
	}
}
