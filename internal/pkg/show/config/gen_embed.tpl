// Code generated by gen_embed_var.go; DO NOT EDIT.
package [[.Package]]

import _ "embed"

//go:embed default.yaml
var DefaultConfig string

// plugin default config
var (
[[range .Plugins]]
//go:embed [[$.Dir]]/[[.]].yaml
[[UpperCamelCase .]]DefaultConfig string
[[end]]
)

var pluginDefaultConfigs = map[string]string{
	[[- range .Plugins]]
	"[[.]]":[[UpperCamelCase . -]]DefaultConfig,
[[- end]]
}

//go:embed templates/quickstart.yaml
var QuickStart string

//go:embed templates/gitops.yaml
var GitOps string
