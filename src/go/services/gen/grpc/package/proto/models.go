package proto

type (
	_enumElement struct {
		Value   int
		Name    string
		Comment string
	}
	_enum struct {
		Name     string
		Comment  string
		Elements []_enumElement
	}
	_option struct {
		Name    string
		Compact bool
		Value   any
	}
	_literalField struct {
		Name  string
		Value any
	}
	_literal struct {
		SingleLine bool
		Fields     []_literalField
	}
	FieldCardinality int
	FieldType        string
	_field           struct {
		ID          int
		Name        string
		Comment     string
		Type        FieldType
		Cardinality FieldCardinality
		Options     []_option
	}
	_extension struct {
		Name   string
		Fields []_field
	}
	_message struct {
		Name       string
		Comment    string
		Fields     []_field
		OneOfs     []_oneOf
		Messages   []_message
		Enums      []_enum
		Extensions []_extension
		Options    []_option
	}
	_oneOf struct {
		Name   string
		Fields []_field
	}
	_method struct {
		Name    string
		Input   string
		Output  string
		Options []_option
	}
	_service struct {
		Name    string
		Methods []_method
	}
	ImportType int
	_import    struct {
		Name string
		Path string
		Type ImportType
	}
	_file struct {
		Package    string
		Imports    []_import
		Messages   []_message
		Enums      []_enum
		Options    []_option
		Extensions []_extension
		Services   []_service
	}
)

const (
	ImportDefault ImportType = iota
	ImportPublic
	ImportWeak
)
const (
	CardinalityDefault FieldCardinality = iota
	CardinalityRequired
	CardinalityOptional
	CardinalityRepeated
)
const (
	TypeString = "string"
	TypeUint64 = "uint64"
)
