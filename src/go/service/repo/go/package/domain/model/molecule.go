package model

type (
	Builder struct {
		Name string
	}
	Options struct {
		Struct
	}
	Mapper struct {
		From, To Struct
	}
	Enum struct {
	}
	Service struct {
		Struct
		Methods []Method
	}
)
