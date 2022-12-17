package pattern

/*
{{$name := .Name}}
type {{$name}}Builder x.Builder[{{.SrcPackage}}.{{$name}}]
func {{.Name}}({{range .Mandatory}}
	{{.Name}} {{.Type}},
{{end}}) *{{$name}}Builder {
	return &{{$name}}Builder{T: {{.SrcPackage}}.{{$name}}{

	}}
}
{{range .Optional}}
func (b *{{$name}}Builder) {{.Name}}(v {{.Type}}) *{{$name}}Builder {
	b.T.{{.Name}} = v
	return b
}
{{end}}
*/
