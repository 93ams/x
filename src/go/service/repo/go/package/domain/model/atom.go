package model

type (
	TypeProps struct {
		Name     string
		Path     string
		Ptr      bool
		Repeated bool
	}
	File struct {
		Pkg, Name string
	}
	MethodType struct {
		Names []string
		Type  TypeProps
	}
	MethodDef struct {
		Name    string
		In, Out []MethodType
	}
	Method struct {
		MethodDef
		Receiver *TypeProps
		Body     any
	}
	StructField struct {
		Names []string
		Type  TypeProps
	}
	Struct struct {
		Fields     []StructField
		Path, Name string
	}
)
