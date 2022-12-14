package driver

import (
	"github.com/samber/lo"
	"grpc/package/model"
	"io"
)

type (
	EnumBuilder        _enum
	EnumElementBuilder _enumElement
	OneOfBuilder       _oneOf
	MessageBuilder     _message
	LiteralBuilder     _literal
	OptionBuilder      _option
	ExtensionBuilder   _extension
	ServiceBuilder     _service
	MethodBuilder      _method
	FileBuilder        _file
	FieldBuilder       _field
	ImportBuilder      _import
)

func Enum(name string, elements ...*EnumElementBuilder) *EnumBuilder {
	b := EnumBuilder(_enum{name: name, elements: lo.Map(elements, func(i *EnumElementBuilder, _ int) _enumElement { return _enumElement(*i) })})
	return &b
}
func (b *EnumBuilder) Elements(el ...*EnumElementBuilder) *EnumBuilder {
	b.elements = append(b.elements, lo.Map(el, func(i *EnumElementBuilder, _ int) _enumElement { return _enumElement(*i) })...)
	return b
}
func (b *EnumBuilder) Comment(s string) *EnumBuilder {
	b.comment = s
	return b
}
func EnumElement(name string, val int) *EnumElementBuilder {
	b := EnumElementBuilder(_enumElement{value: val, name: name})
	return &b
}
func (b *EnumElementBuilder) Comment(s string) *EnumElementBuilder {
	b.comment = s
	return b
}
func Message(name string, fields ...*FieldBuilder) *MessageBuilder {
	b := MessageBuilder(_message{name: name, fields: lo.Map(fields, func(i *FieldBuilder, _ int) _field { return _field(*i) })})
	return &b
}
func (b *MessageBuilder) Options(options ...*OptionBuilder) *MessageBuilder {
	b.options = append(b.options, lo.Map(options, func(i *OptionBuilder, _ int) _option { return _option(*i) })...)
	return b
}
func (b *MessageBuilder) Extensions(extensions ...*ExtensionBuilder) *MessageBuilder {
	b.extensions = append(b.extensions, lo.Map(extensions, func(i *ExtensionBuilder, _ int) _extension { return _extension(*i) })...)
	return b
}
func (b *MessageBuilder) Messages(messages ...*MessageBuilder) *MessageBuilder {
	b.messages = append(b.messages, lo.Map(messages, func(i *MessageBuilder, _ int) _message { return _message(*i) })...)
	return b
}
func (b *MessageBuilder) OneOfs(oneOves ...*OneOfBuilder) *MessageBuilder {
	b.oneOfs = append(b.oneOfs, lo.Map(oneOves, func(i *OneOfBuilder, _ int) _oneOf { return _oneOf(*i) })...)
	return b
}
func (b *MessageBuilder) Enums(enums ...*EnumBuilder) *MessageBuilder {
	b.enums = append(b.enums, lo.Map(enums, func(i *EnumBuilder, _ int) _enum { return _enum(*i) })...)
	return b
}
func (b *MessageBuilder) Fields(fields ...*FieldBuilder) *MessageBuilder {
	b.fields = append(b.fields, lo.Map(fields, func(i *FieldBuilder, _ int) _field { return _field(*i) })...)
	return b
}
func (b *MessageBuilder) Comment(comment string) *MessageBuilder {
	b.comment = comment
	return b
}
func OneOf(name string, fields ...*FieldBuilder) *OneOfBuilder {
	b := OneOfBuilder(_oneOf{name: name, fields: lo.Map(fields, func(i *FieldBuilder, _ int) _field { return _field(*i) })})
	return &b
}
func (b *OneOfBuilder) Fields(fields ...*FieldBuilder) *OneOfBuilder {
	b.fields = append(b.fields, lo.Map(fields, func(i *FieldBuilder, _ int) _field { return _field(*i) })...)
	return b
}
func Option(name string, value any) *OptionBuilder {
	b := OptionBuilder(_option{name: name, value: value})
	return &b
}
func (b *OptionBuilder) Compact(compact bool) *OptionBuilder {
	b.compact = compact
	return b
}
func (b *OptionBuilder) Value(value any) *OptionBuilder {
	b.value = value
	return b
}
func Literal(fields map[string]any) *LiteralBuilder {
	var f []_literalField
	for k, v := range fields {
		f = append(f, _literalField{name: k, value: v})
	}
	b := LiteralBuilder(_literal{fields: f})
	return &b
}
func (b *LiteralBuilder) Field(name string, value any) *LiteralBuilder {
	b.fields = append(b.fields, _literalField{name: name, value: value})
	return b
}
func (b *LiteralBuilder) SingleLine(singleLine bool) *LiteralBuilder {
	b.singleLine = singleLine
	return b
}
func Extension(name string, fields ...*FieldBuilder) *ExtensionBuilder {
	b := ExtensionBuilder(_extension{name: name, fields: lo.Map(fields, func(i *FieldBuilder, _ int) _field { return _field(*i) })})
	return &b
}
func (b *ExtensionBuilder) Fields(fields ...*FieldBuilder) *ExtensionBuilder {
	b.fields = append(b.fields, lo.Map(fields, func(i *FieldBuilder, _ int) _field { return _field(*i) })...)
	return b
}
func Service(name string, methods ...*MethodBuilder) *ServiceBuilder {
	b := ServiceBuilder(_service{name: name, methods: lo.Map(methods, func(i *MethodBuilder, _ int) _method { return _method(*i) })})
	return &b
}
func (b *ServiceBuilder) Methods(methods ...*MethodBuilder) *ServiceBuilder {
	b.methods = append(b.methods, lo.Map(methods, func(i *MethodBuilder, _ int) _method { return _method(*i) })...)
	return b
}
func Method(name, input, output string) *MethodBuilder {
	b := MethodBuilder(_method{name: name, input: input, output: output})
	return &b
}
func (b *MethodBuilder) Options(options ...*OptionBuilder) *MethodBuilder {
	b.options = append(b.options, lo.Map(options, func(i *OptionBuilder, _ int) _option { return _option(*i) })...)
	return b
}
func File(pkg string) *FileBuilder {
	b := FileBuilder(_file{pkg: pkg})
	return &b
}
func (b *FileBuilder) Imports(imports ...*ImportBuilder) *FileBuilder {
	b.imports = append(b.imports, lo.Map(imports, func(i *ImportBuilder, _ int) _import { return _import(*i) })...)
	return b
}
func (b *FileBuilder) Options(options ...*OptionBuilder) *FileBuilder {
	b.options = append(b.options, lo.Map(options, func(i *OptionBuilder, _ int) _option { return _option(*i) })...)
	return b
}
func (b *FileBuilder) Extensions(extensions ...*ExtensionBuilder) *FileBuilder {
	b.extensions = append(b.extensions, lo.Map(extensions, func(i *ExtensionBuilder, _ int) _extension { return _extension(*i) })...)
	return b
}
func (b *FileBuilder) Messages(messages ...*MessageBuilder) *FileBuilder {
	b.messages = append(b.messages, lo.Map(messages, func(i *MessageBuilder, _ int) _message { return _message(*i) })...)
	return b
}
func (b *FileBuilder) Services(services ...*ServiceBuilder) *FileBuilder {
	b.services = append(b.services, lo.Map(services, func(i *ServiceBuilder, _ int) _service { return _service(*i) })...)
	return b
}
func (b *FileBuilder) Enums(enums ...*EnumBuilder) *FileBuilder {
	b.enums = append(b.enums, lo.Map(enums, func(i *EnumBuilder, _ int) _enum { return _enum(*i) })...)
	return b
}
func Field(name string, t model.FieldType, id int) *FieldBuilder {
	b := FieldBuilder(_field{name: name, typ: t, id: id})
	return &b
}
func (b *FieldBuilder) Options(options ...*OptionBuilder) *FieldBuilder {
	b.options = append(b.options, lo.Map(options, func(i *OptionBuilder, _ int) _option { return _option(*i) })...)
	return b
}
func (b *FieldBuilder) Comment(comment string) *FieldBuilder {
	b.comment = comment
	return b
}
func (b *FieldBuilder) Cardinality(cardinality model.FieldCardinality) *FieldBuilder {
	b.cardinality = cardinality
	return b
}
func Import(name, path string) *ImportBuilder {
	b := ImportBuilder(_import{name: name, path: path})
	return &b
}
func (b *ImportBuilder) Type(importType model.ImportType) *ImportBuilder {
	b.typ = importType
	return b
}
func (b *FileBuilder) Encode(indent int, dst io.Writer) error {
	return _file(*b).Encode(indent, dst)
}
