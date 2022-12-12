package proto

import (
	. "github.com/tilau2328/cql/src/go/package/shared/x"
)

type (
	EnumBuilder        Builder[_enum]
	EnumElementBuilder Builder[_enumElement]
	OneOfBuilder       Builder[_oneOf]
	MessageBuilder     Builder[_message]
	LiteralBuilder     Builder[_literal]
	OptionBuilder      Builder[_option]
	ExtensionBuilder   Builder[_extension]
	ServiceBuilder     Builder[_service]
	MethodBuilder      Builder[_method]
	FileBuilder        Builder[_file]
	FieldBuilder       Builder[_field]
	ImportBuilder      Builder[_import]
)

func Enum(name string, elements ..._enumElement) *EnumBuilder {
	return &EnumBuilder{T: _enum{Name: name, Elements: elements}}
}
func (b *EnumBuilder) Elements(el ..._enumElement) *EnumBuilder {
	b.T.Elements = append(b.T.Elements, el...)
	return b
}
func (b *EnumBuilder) Comment(s string) *EnumBuilder {
	b.T.Comment = s
	return b
}
func EnumElement(name string, val int) *EnumElementBuilder {
	return &EnumElementBuilder{T: _enumElement{Value: val, Name: name}}
}
func (b *EnumElementBuilder) Comment(s string) *EnumElementBuilder {
	b.T.Comment = s
	return b
}
func Message(name string, fields ..._field) *MessageBuilder {
	return &MessageBuilder{T: _message{Name: name, Fields: fields}}
}
func (b *MessageBuilder) Options(opts ..._option) *MessageBuilder {
	b.T.Options = append(b.T.Options, opts...)
	return b
}
func (b *MessageBuilder) Extensions(extensions ..._extension) *MessageBuilder {
	b.T.Extensions = append(b.T.Extensions, extensions...)
	return b
}
func (b *MessageBuilder) Messages(messages ..._message) *MessageBuilder {
	b.T.Messages = append(b.T.Messages, messages...)
	return b
}
func (b *MessageBuilder) OneOfs(oneOves ..._oneOf) *MessageBuilder {
	b.T.OneOfs = append(b.T.OneOfs, oneOves...)
	return b
}
func (b *MessageBuilder) Enums(enums ..._enum) *MessageBuilder {
	b.T.Enums = append(b.T.Enums, enums...)
	return b
}
func (b *MessageBuilder) Fields(fields ..._field) *MessageBuilder {
	b.T.Fields = append(b.T.Fields, fields...)
	return b
}
func (b *MessageBuilder) Comment(comment string) *MessageBuilder {
	b.T.Comment = comment
	return b
}
func OneOf(name string, fields ..._field) *OneOfBuilder {
	return &OneOfBuilder{T: _oneOf{Name: name, Fields: fields}}
}
func (b *OneOfBuilder) Fields(fields ..._field) *OneOfBuilder {
	b.T.Fields = append(b.T.Fields, fields...)
	return b
}
func (b *OneOfBuilder) Name(name string) *OneOfBuilder {
	b.T.Name = name
	return b
}
func Option(name string, value any) *OptionBuilder {
	return &OptionBuilder{T: _option{Name: name, Value: value}}
}
func (b *OptionBuilder) Compact(compact bool) *OptionBuilder {
	b.T.Compact = compact
	return b
}
func (b *OptionBuilder) Value(value any) *OptionBuilder {
	b.T.Value = value
	return b
}
func Literal(fields map[string]any) *LiteralBuilder {
	var f []_literalField
	for k, v := range fields {
		f = append(f, _literalField{Name: k, Value: v})
	}
	return &LiteralBuilder{T: _literal{Fields: f}}
}
func (b *LiteralBuilder) Field(name string, value any) *LiteralBuilder {
	b.T.Fields = append(b.T.Fields, _literalField{Name: name, Value: value})
	return b
}
func (b *LiteralBuilder) SingleLine(singleLine bool) *LiteralBuilder {
	b.T.SingleLine = singleLine
	return b
}
func Extension(name string, fields ..._field) *ExtensionBuilder {
	return &ExtensionBuilder{T: _extension{Name: name, Fields: fields}}
}
func (b *ExtensionBuilder) Fields(fields ..._field) *ExtensionBuilder {
	b.T.Fields = append(b.T.Fields, fields...)
	return b
}
func Service(name string, methods ..._method) *ServiceBuilder {
	return &ServiceBuilder{T: _service{Name: name, Methods: methods}}
}
func (b *ServiceBuilder) Methods(methods ..._method) *ServiceBuilder {
	b.T.Methods = append(b.T.Methods, methods...)
	return b
}
func Method(name, input, output string) *MethodBuilder {
	return &MethodBuilder{T: _method{Name: name, Input: input, Output: output}}
}
func (b *MethodBuilder) Options(options ..._option) *MethodBuilder {
	b.T.Options = append(b.T.Options, options...)
	return b
}
func File(pkg string) *FileBuilder { return &FileBuilder{T: _file{Package: pkg}} }
func (b *FileBuilder) Imports(imports ..._import) *FileBuilder {
	b.T.Imports = append(b.T.Imports, imports...)
	return b
}
func (b *FileBuilder) Options(options ..._option) *FileBuilder {
	b.T.Options = append(b.T.Options, options...)
	return b
}
func (b *FileBuilder) Extensions(v ..._extension) *FileBuilder {
	b.T.Extensions = append(b.T.Extensions, v...)
	return b
}
func (b *FileBuilder) Messages(v ..._message) *FileBuilder {
	b.T.Messages = append(b.T.Messages, v...)
	return b
}
func (b *FileBuilder) Services(v ..._service) *FileBuilder {
	b.T.Services = append(b.T.Services, v...)
	return b
}
func (b *FileBuilder) Enums(v ..._enum) *FileBuilder {
	b.T.Enums = append(b.T.Enums, v...)
	return b
}
func Field(name string, t FieldType, id int) *FieldBuilder {
	return &FieldBuilder{T: _field{Name: name, Type: t, ID: id}}
}
func (b *FieldBuilder) Options(options ..._option) *FieldBuilder {
	b.T.Options = append(b.T.Options, options...)
	return b
}
func (b *FieldBuilder) Comment(comment string) *FieldBuilder {
	b.T.Comment = comment
	return b
}
func (b *FieldBuilder) Cardinality(cardinality FieldCardinality) *FieldBuilder {
	b.T.Cardinality = cardinality
	return b
}
func Import(name, path string) *ImportBuilder {
	return &ImportBuilder{T: _import{Name: name, Path: path}}
}
func (b *ImportBuilder) Type(importType ImportType) *ImportBuilder {
	b.T.Type = importType
	return b
}
