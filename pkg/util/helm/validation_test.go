package helm

import (
	"testing"
)

func Test_validate(t *testing.T) {
	type args struct {
		param *HelmParam
	}
	tests := []struct {
		name string
		args args
		want int // error count
	}{
		// TODO: Add test cases.
		{"base", args{&HelmParam{
			Repo{Name: "argo", URL: "https://argoproj.github.io/argo-helm"},
			Chart{ChartName: "argo/argo-cd"},
		}}, 0},
		{"one required field validation error", args{&HelmParam{
			Repo{Name: "argo", URL: "https://argoproj.github.io/argo-helm"},
			Chart{ChartName: ""},
		}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.param); len(got) != tt.want {
				t.Logf("got errors' length: %d\n", len(got))
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
