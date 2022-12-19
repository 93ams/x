package model

type (
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
	Builder struct {
		Name string
	}
	Options struct {
		Struct
	}
	Service struct {
		Struct
		Methods []Method
	}
	StructField struct {
		Names []string
		Type  TypeProps
	}
	Struct struct {
		Fields     []StructField
		Path, Name string
	}
	CommandProps struct {
	}
	Interface struct {
		Methods []MethodDef
		Name    string
	}
	TypeProps struct {
		Name     string
		Path     string
		Ptr      bool
		Repeated bool
	}
	Enum struct {
	}
	Mapper struct {
		From, To Struct
	}
)
