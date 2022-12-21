package model

type (
	Builder struct {
		New    bool
		Name   string
		Pkg    string
		From   string
		Filter []string
	}
	Options struct {
		Struct
		From   string
		Filter []string
	}
	Mapper struct {
		From, To Struct
	}
	Enum struct {
	}
	Service struct {
		Struct
		Methods []Func
	}
	Command struct {
	}
	Visitor struct {
	}
)
