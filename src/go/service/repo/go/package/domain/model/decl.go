package model

type (
	Field struct {
		Names []string
		Type  Ident
	}
	FuncType struct {
		Name    string
		In, Out []Field
	}
	Func struct {
		Receiver *Ident
		Body     any
		FuncType
	}
	Ident struct {
		Name, Path    string
		Repeated, Ptr bool
		Generic       []Ident
	}
	Selector struct {
		Ident
	}
	TypeDef struct {
		Generic map[string]Ident
		Name    string
		Ident
	}
	Struct struct {
		Fields []Field
		Ident
	}
	Interface struct {
		Methods    []FuncType
		Path, Name string
	}
)
