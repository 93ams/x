package model

type (
	EnumElement struct {
		Value   int
		Name    string
		comment string
	}
	Enum struct {
		Name     string
		comment  string
		elements []EnumElement
	}
	Option struct {
		Name    string
		compact bool
		Value   any
	}
	LiteralField struct {
		Name  string
		Value any
	}
	Literal struct {
		SingleLine bool
		Fields     []LiteralField
	}
	Field struct {
		ID          int
		Name        string
		comment     string
		Type        FieldType
		Cardinality FieldCardinality
		Options     []Option
	}
	Extension struct {
		Name   string
		Fields []Field
	}
	Message struct {
		Name       string
		comment    string
		fields     []Field
		oneOfs     []OneOf
		messages   []Message
		enums      []Enum
		extensions []Extension
		options    []Option
	}
	OneOf struct {
		Name   string
		fields []Field
	}
	Method struct {
		Name    string
		input   string
		output  string
		options []Option
	}
	Service struct {
		Name    string
		Methods []Method
	}
	Import struct {
		Name string
		Path string
		Type ImportType
	}
	File struct {
		Pkg        string
		Imports    []Import
		Messages   []Message
		Enums      []Enum
		Options    []Option
		Extensions []Extension
		Services   []Service
	}
)
