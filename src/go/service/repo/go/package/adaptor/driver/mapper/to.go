package mapper

import (
	"github.com/samber/lo"
	. "github.com/tilau2328/x/src/go/services/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/builder"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/model"
	"go/token"
)

func NewNames(names []string) []*builder.IdentBuilder {
	return lo.Map(names, func(item string, _ int) *builder.IdentBuilder { return builder.Ident(item) })
}
func NewMethodTypes(methods []MethodType) []*builder.FieldBuilder {
	return lo.Map(methods, func(item MethodType, _ int) *builder.FieldBuilder {
		return builder.Field(builder.Ident(item.Type.Name).Path(item.Type.Path), NewNames(item.Names)...)
	})
}
func NewMethodDefs(methods []MethodDef) []*builder.FieldBuilder {
	return lo.Map(methods, func(v MethodDef, _ int) *builder.FieldBuilder {
		return builder.Field(builder.FuncType().
			Params(builder.FieldList(NewMethodTypes(v.In)...)).
			Results(builder.FieldList(NewMethodTypes(v.Out)...)),
			builder.Ident(v.Name))
	})
}
func NewMapperFields(fields []StructField) []builder.ExprBuilder {
	return lo.Map(fields, func(p StructField, index int) builder.ExprBuilder {
		return builder.KeyValue(builder.Ident(p.Names[0]),
			builder.Selector(builder.Ident("in"), builder.Ident(p.Names[0])),
		).Decs(builder.KeyValueDecs().Before(model.NewLine).After(model.NewLine))
	})
}
func NewStructFields(fields []StructField) []*builder.FieldBuilder {
	return lo.Map(fields, func(v StructField, _ int) *builder.FieldBuilder {
		return builder.Field(builder.Ident(v.Type.Name).Path(v.Type.Path), NewNames(v.Names)...)
	})
}

func NewMethod(props Method) *builder.FuncBuilder {
	return builder.Func(builder.Ident(props.Name)).Type(builder.FuncType().
		Params(builder.FieldList(NewMethodTypes(props.In)...)).
		Results(builder.FieldList(NewMethodTypes(props.Out)...)))
}
func NewStruct(props Struct) *builder.GenBuilder {
	return builder.Gen(token.TYPE, builder.Type(
		builder.Struct(builder.FieldList(NewStructFields(props.Fields)...)),
	).Name(builder.Ident(props.Name)))
}
func NewInterface(props Interface) *builder.GenBuilder {
	return builder.Gen(token.TYPE, builder.Type(
		builder.Interface(builder.FieldList(NewMethodDefs(props.Methods)...)),
	).Name(builder.Ident(props.Name)))
}
func NewFile(props File) {

}
func NewEnum(props Enum) []builder.DeclBuilder {
	return nil
}
func NewBuilder(props Builder) []builder.DeclBuilder {
	return nil
}
func NewMapper(props Mapper) []builder.DeclBuilder {
	return []builder.DeclBuilder{
		builder.Func(builder.Ident("From" + props.From.Name)).
			Type(builder.FuncType().
				Params(builder.FieldList(builder.Field(builder.Ident(props.From.Name).Path(props.From.Path), builder.Ident("in")))).
				Results(builder.FieldList(builder.Field(builder.Ident(props.To.Name).Path(props.To.Path))))).
			Body(builder.Block(
				builder.Return(builder.CompositeLit(builder.Ident(props.To.Name).Path(props.To.Path), NewMapperFields(props.From.Fields)...)).
					Decs(builder.ReturnDecs().Before(model.NewLine)),
			)),
		builder.Func(builder.Ident("To" + props.From.Name)).
			Type(builder.FuncType().
				Params(builder.FieldList(builder.Field(builder.Ident(props.To.Name).Path(props.To.Path), builder.Ident("in")))).
				Results(builder.FieldList(builder.Field(builder.Ident(props.From.Name).Path(props.From.Path)))),
			).Body(builder.Block(
			builder.Return(builder.CompositeLit(builder.Ident(props.From.Name).Path(props.From.Path), NewMapperFields(props.From.Fields)...)).
				Decs(builder.ReturnDecs().Before(model.NewLine)),
		)),
	}
}
func NewOptions(props Options) []builder.DeclBuilder {
	ret := []builder.DeclBuilder{
		NewMethod(Method{
			MethodDef: MethodDef{
				Name: "New" + props.Name,
				In: []MethodType{
					{Names: []string{"opts"}, Type: TypeProps{Path: "x", Name: "Opt", Repeated: true}},
				},
				Out: []MethodType{
					{Type: TypeProps{Name: props.Name, Ptr: true}},
				},
			},
			Body: nil,
		}),
	}
	return ret
}
func NewService(props Service) []builder.DeclBuilder {
	return []builder.DeclBuilder{
		NewStruct(props.Struct),
	}
}
func NewDecls(props any) (res []builder.DeclBuilder) {
	switch p := props.(type) {
	case Method:
		res = append(res, NewMethod(p))
	case Struct:
		res = append(res, NewStruct(p))
	case Interface:
		res = append(res, NewInterface(p))

	case Enum:
		res = NewEnum(p)
	case Mapper:
		res = NewMapper(p)
	case Builder:
		res = NewBuilder(p)
	case Options:
		res = NewOptions(p)
	case Service:
		res = NewService(p)
	}
	return
}
