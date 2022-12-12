package tmpl

import (
	"github.com/samber/lo"
	"io"
	"text/template"
)

func Requester(w io.Writer, props RequesterProps) error {
	return requesterTemplate.Execute(w, props)
}

var requesterTemplate = lo.Must(template.New("requester").Parse(`package {{.Pkg}}
{{range .Imports}}
{{end}}

type {{.Name}}Requester struct {
}

func New{{.Name}}Requester() *{{.Name}}Requester{
	return &{{.Name}}Requester{}
}

{{range .Methods}}
func (r *{{.Name}}Requester) {{.Name}}() {
}
{{end}}
`))
