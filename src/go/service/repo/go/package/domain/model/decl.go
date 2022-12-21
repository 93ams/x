package model

type (
	MethodType struct {
		Names []string
		Type  Ident
	}
	FuncType struct {
		In, Out []MethodType
	}
	Func struct {
		Receiver *Ident
		Name     string
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
	StructField struct {
		Names []string
		Type  Ident
	}
	Struct struct {
		Fields []StructField
		Ident
	}
	Interface struct {
		Methods    []FuncType
		Path, Name string
	}
)
