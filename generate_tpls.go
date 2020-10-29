package fastyaml

import (
	"strings"
	"text/template"
)

var fileTpl = template.Must(template.New("file").
	Funcs(template.FuncMap{
		"count": func(n int) []int {
			result := make([]int, 0)
			for i := 0; i < n; i++ {
				result = append(result, i)
			}
			return result
		},
		"add": func(a, b int) int {
			return a + b
		},
		"mapType": func(n int) string {
			result := make([]string, 0)
			for i := 0; i < n; i++ {
				result = append(result, "map[string]")
			}
			result = append(result, "string")
			return strings.Join(result, "")
		},
	}).
	Parse(`package {{ .Package }}

// ---------------------------------------
//    Auto-generated file, do not edit!
// ---------------------------------------

{{ if .Imports }}
import (
{{ range .Imports }}"{{ . }}"{{ end }}
)
{{ end }}

type parse{{ .Name }} struct {
	*parser.Parser
}

func Parse{{ .Name }}(in string) (*{{ .Name }}, error) {
	p := &parse{{ .Name }}{
		Parser: parser.New(in),
	}
	return p.parse{{ .Name }}()
}

{{ range $index, $type := .Types }}
func (p *parse{{ $.Name }}) parse{{ .Name }}() ({{ .Type }}, error) {
	result := {{ .TypeInit }}
	{{- if ne .Rest "" }}
	result.{{ .Rest }} = {{ .RestType }}{}
	{{- end }}

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		{{- range .Fields -}}
		case "{{ .Name }}":
			{{- if .AdvanceBefore }}p.AdvanceLine(){{ end }}
			{{- if .Slice }}
			ad := p.Depth
			for ad == p.Depth {
			{{- end }}
			o, err := {{ .Method }}()
			if err != nil {
				return {{ $type.ErrType }}, err
			}
			{{- if .Slice }}
				result.{{ .Field }} = append(result.{{ .Field }}, o)
			{{- else }}
				result.{{ .Field }} = o
			{{- end }}
			{{ if not .AdvanceBefore }}p.AdvanceLine(){{ end }}
			{{- if .Slice }} } {{ end }}
		{{ end -}}
		default:
		{{- if eq .Rest "" }}
			p.SkipLine()
		{{- else }}
			p.AdvanceLine()
			o, err := p.consumeMap{{ add .RestDepth -1 }}()
			if err != nil {
				return {{ $type.ErrType }}, err
			}
			result.{{ .Rest }}[key] = o
		{{- end }}
		}

		if p.IsNew {
		    break
		}
	}
	return result, nil
}
{{ end }}

{{ range $index, $n := count .MapDepth }}
{{ if gt $n 0 }}
func (p *parse{{ $.Name }}) consumeMap{{ $n }}() ({{ mapType $n }}, error) {
	result := make({{ mapType $n }})

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()
		{{- if eq $n 1 }}
		o, err := p.ReadString()
		if err != nil {
			return nil, err
		}
		result[key] = o
		p.AdvanceLine()
		{{- else }}
		p.AdvanceLine()
		o, err := p.consumeMap{{ add $n -1 }}()
		if err != nil {
			return nil, err
		}
		result[key] = o
		{{- end }}
	}

	return result, nil
}
{{ end }}
{{ end }}
`))
