package builder

import (
	"grpc/package/domain/model"
)

type (
	EnumBuilder            Builder[model.Enum]
	EnumFieldBuilder       Builder[model.EnumField]
	EnumValueOptionBuilder Builder[model.EnumValueOption]
	ExtendBuilder          Builder[model.Extend]
	ExtensionsBuilder      Builder[model.Extensions]
	FieldBuilder           Builder[model.Field]
	FieldOptionBuilder     Builder[model.FieldOption]
	GroupFieldBuilder      Builder[model.GroupField]
	ImportBuilder          Builder[model.Import]
	MapFieldBuilder        Builder[model.MapField]
	MessageBuilder         Builder[model.Message]
	OneOfBuilder           Builder[model.OneOf]
	OneOfFieldBuilder      Builder[model.OneOfField]
	OptionBuilder          Builder[model.Option]
	PackageBuilder         Builder[model.Package]
	ProtoBuilder           Builder[model.Proto]
	RangeBuilder           Builder[model.Range]
	ReservedBuilder        Builder[model.Reserved]
	RPCBuilder             Builder[model.RPC]
	RPCRequestBuilder      Builder[model.Request]
	RPCResponseBuilder     Builder[model.Response]
	ServiceBuilder         Builder[model.Service]
	SyntaxBuilder          Builder[model.Syntax]
)

func Enum(name string) *EnumBuilder { return &EnumBuilder{T: model.Enum{Name: name}} }
func EnumField(ident, number string) *EnumFieldBuilder {
	return &EnumFieldBuilder{T: model.EnumField{Ident: ident, Number: number}}
}
func EnumValueOption(name, constant string) *EnumValueOptionBuilder {
	return &EnumValueOptionBuilder{T: model.EnumValueOption{Name: name, Constant: constant}}
}
func Extend(messageType string) *ExtendBuilder {
	return &ExtendBuilder{T: model.Extend{Type: messageType}}
}
func Extensions(ranges ...RangeBuilder) *ExtensionsBuilder {
	return &ExtensionsBuilder{T: model.Extensions{}}
}
func Field() *FieldBuilder { return &FieldBuilder{T: model.Field{}} }
func FieldOption(name, constant string) *FieldOptionBuilder {
	return &FieldOptionBuilder{T: model.FieldOption{Name: name, Constant: constant}}
}
func GroupField() *GroupFieldBuilder       { return &GroupFieldBuilder{T: model.GroupField{}} }
func Import() *ImportBuilder               { return &ImportBuilder{T: model.Import{}} }
func MapField() *MapFieldBuilder           { return &MapFieldBuilder{T: model.MapField{}} }
func Message() *MessageBuilder             { return &MessageBuilder{T: model.Message{}} }
func OneOf() *OneOfBuilder                 { return &OneOfBuilder{T: model.OneOf{}} }
func OneOfField() *OneOfFieldBuilder       { return &OneOfFieldBuilder{T: model.OneOfField{}} }
func Option() *OptionBuilder               { return &OptionBuilder{T: model.Option{}} }
func Package() *PackageBuilder             { return &PackageBuilder{T: model.Package{}} }
func Proto() *ProtoBuilder                 { return &ProtoBuilder{T: model.Proto{}} }
func Range() *RangeBuilder                 { return &RangeBuilder{T: model.Range{}} }
func Reserved() *ReservedBuilder           { return &ReservedBuilder{T: model.Reserved{}} }
func RPC() *RPCBuilder                     { return &RPCBuilder{T: model.RPC{}} }
func RPCRequest() *RPCRequestBuilder       { return &RPCRequestBuilder{T: model.Request{}} }
func RPCResponse() *RPCResponseBuilder     { return &RPCResponseBuilder{T: model.Response{}} }
func Service() *ServiceBuilder             { return &ServiceBuilder{T: model.Service{}} }
func Syntax(version string) *SyntaxBuilder { return &SyntaxBuilder{T: model.Syntax{Version: version}} }
