package fastyaml

import "text/template"

var fileTpl = template.Must(template.New("file").Parse(`package {{ .Package }}

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

	d := p.Depth
	for d == p.Depth {
		key := p.ReadKey()

		switch key {
		{{ range .Fields -}}
		case "{{ .Name }}":
			{{ if .AdvanceBefore }}p.AdvanceLine(){{ end }}
			o, err := {{ .Method }}()
			if err != nil {
				return {{ $type.ErrType }}, err
			}
			result.{{ .Field }} = o
			{{ if not .AdvanceBefore }}p.AdvanceLine(){{ end }}
		{{ end }}
		default:
			p.SkipLine()
		}
	}
	return result, nil
}
{{ end }}
`))
