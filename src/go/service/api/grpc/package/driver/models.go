package driver

import "grpc/package/model"

type (
	_enumElement struct {
		value   int
		name    string
		comment string
	}
	_enum struct {
		name     string
		comment  string
		elements []_enumElement
	}
	_option struct {
		name    string
		compact bool
		value   any
	}
	_literalField struct {
		name  string
		value any
	}
	_literal struct {
		singleLine bool
		fields     []_literalField
	}
	_field struct {
		id          int
		name        string
		comment     string
		typ         model.FieldType
		cardinality model.FieldCardinality
		options     []_option
	}
	_extension struct {
		name   string
		fields []_field
	}
	_message struct {
		name       string
		comment    string
		fields     []_field
		oneOfs     []_oneOf
		messages   []_message
		enums      []_enum
		extensions []_extension
		options    []_option
	}
	_oneOf struct {
		name   string
		fields []_field
	}
	_method struct {
		name    string
		input   string
		output  string
		options []_option
	}
	_service struct {
		name    string
		methods []_method
	}
	_import struct {
		name string
		path string
		typ  model.ImportType
	}
	_file struct {
		pkg        string
		imports    []_import
		messages   []_message
		enums      []_enum
		options    []_option
		extensions []_extension
		services   []_service
	}
)
