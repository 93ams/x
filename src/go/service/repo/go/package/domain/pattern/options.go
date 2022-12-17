package pattern

/*
{{$name := .Name}}
func New{{$name}}(...opts x.Opt[*{{$name}}]) *{{$name}} {
	return x.Apply(&{{Name}}{}, opts)
}
{{range .Fields}}
func {{.Name}}(val {{.Type}}) x.Opt[*{{$name}} {
	return func(v *{{$name}}) { v.{{.Name}} = val }
}
{{end}}
*/
