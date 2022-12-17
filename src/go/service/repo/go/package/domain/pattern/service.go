package pattern

/*
{{$name := .Name}}
type (
	{{$name}}Service struct {
		{{$name}}ServiceOpts
	}
	{{$name}}ServiceOpts struct {
{{range .Deps}}
		{{.Name}} {{.Type}}
{{end}}
	}
)
func New{{$name}}Service(opts {{$name}}ServiceOpts) *{{$name}}Service {
	return &{{$name}}Service{{{$name}}ServiceOpts: opts}
}
{{range .Methods}}
func (s *{{$name}}Service) {{.Name}}({{.In}}) {{.Out}} {
	{{.Body}}
	return {{.Return}}
}
{{end}}
*/
