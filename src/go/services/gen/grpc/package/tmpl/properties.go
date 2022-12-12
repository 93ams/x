package tmpl

type (
	MethodProps struct {
		Name    string
		Return  string
		Request string
	}
	RequesterProps struct {
		Pkg     string
		Name    string
		Imports map[string]string
		Methods []MethodProps
	}
	HandlerProps struct {
		Pkg      string
		Name     string
		Provider string
		Imports  map[string]string
		Methods  []MethodProps
	}
	StructProps struct {
		Name   string
		Fields map[string]string
	}
	MappersProps struct {
		Pkg     string
		Imports map[string]string
		Structs []StructProps
	}
)
