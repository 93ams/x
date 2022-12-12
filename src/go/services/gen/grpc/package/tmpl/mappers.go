package tmpl

import (
	"github.com/samber/lo"
	"io"
	"text/template"
)

func Mappers(w io.Writer, props MappersProps) error {
	return mappersTemplate.Execute(w, props)
}

var mappersTemplate = lo.Must(template.New("mappers").Parse(`package {{.Pkg}}

import (
{{range $import, $alias := .Imports}}
	{{$alias}} "{{$import}}"
{{end}}
)

{{range .Structs}}
func To{{.Name}}(in {{.Name}}) model.{{.Name}}{
	return model.{{.Name}}{
		{{range .Fields}}
		{{end}}
	}
}
func To{{.Name}}s(in []{{.Name}}) []model.{{.Name}}{
	ret := make([]model.{{.Name}})
	for i := range in {
		ret[i] = To{{.Name}}(in[i])
	}
	return ret
}
func From{{.Name}}(in model.{{.Name}}) {{.Name}}{
	return {{.Name}}{
		{{range .Fields}}
		{{end}}
	}
}
func From{{.Name}}s(in []model.{{.Name}}) []{{.Name}}{
	ret := make([]{{.Name}})
	for i := range in {
		ret[i] = From{{.Name}}(in[i])
	}
	return ret
}
{{end}}
`))
