// Code generated by codegen using template tables.go.tpl; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"

{{range $v := .}}  	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/{{$v.AWSService | ToLower}}"
{{end}}
)

func PluginAutoGeneratedTables() []*schema.Table {
	return []*schema.Table{
{{- range .}}
	{{- if not .Parent}}
		{{.AWSService | ToLower}}.{{.TableFuncName}}(),
	{{- end}}
{{- end}}
	}
}
