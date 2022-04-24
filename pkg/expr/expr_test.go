package expr_test

import (
	"testing"

	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/domain"
	"github.com/suzuki-shunsuke/renovate-issue-action/pkg/expr"
)

func TestRender(t *testing.T) {
	t.Parallel()
	data := []struct {
		name  string
		isErr bool
		s     string
		param *expr.Param
		exp   bool
	}{
		{
			name: "true",
			s:    `Metadata.PackageName == "actions/checkout"`,
			exp:  true,
			param: &expr.Param{
				Metadata: &domain.Metadata{
					PackageName: "actions/checkout",
				},
			},
		},
		{
			name: "false",
			s:    `Metadata.PackageName != "actions/checkout"`,
			exp:  false,
			param: &expr.Param{
				Metadata: &domain.Metadata{
					PackageName: "actions/checkout",
				},
			},
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
			s, err := expr.ExecBool(d.s, d.param)
			if err != nil {
				if !d.isErr {
					t.Fatal(err)
				}
				return
			}
			if s != d.exp {
				t.Fatalf("wanted %v, got %v", d.exp, s)
			}
		})
	}
}
