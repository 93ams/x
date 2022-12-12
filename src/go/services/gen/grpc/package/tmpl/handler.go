package tmpl

import (
	"github.com/samber/lo"
	"io"
	"text/template"
)

func Handler(w io.Writer, props HandlerProps) error {
	return handlerTemplate.Execute(w, props)
}

var handlerTemplate = lo.Must(template.New("handler").Parse(`package {{.Pkg}}
import (
{{- range $import, $alias := .Imports}}
	{{$alias}} "{{$import}}"
{{- end}}
)
type {{.Name}}Handler struct {
	provider {{.Provider}}
}

func New{{.Name}}Handler() *{{.Name}}Handler{
	return &{{.Name}}Handler{}
}
{{- range .Methods}}
func (h *{{.Name}}Handler) {{.Name}}(ctx context.Context, req {{.Request}}) ({{.Return}}, error) {
}
{{end}}
`))
