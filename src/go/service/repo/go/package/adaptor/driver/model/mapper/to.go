package mapper

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	builder2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/builder"
	. "github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"go/token"
)

func NewNames(names []string) []*builder2.IdentBuilder {
	return lo.Map(names, func(item string, _ int) *builder2.IdentBuilder { return builder2.Ident(item) })
}
func NewIdents(names []Ident) []builder2.ExprBuilder {
	return lo.Map(names, func(item Ident, _ int) builder2.ExprBuilder { return NewIdent(item) })
}
func NewMethodTypes(methods []Field) []*builder2.FieldBuilder {
	return lo.Map(methods, func(item Field, _ int) *builder2.FieldBuilder {
		return builder2.Field(NewIdent(item.Type), NewNames(item.Names)...)
	})
}
func NewFuncTypes(methods []FuncType) []*builder2.FieldBuilder {
	return lo.Map(methods, func(v FuncType, _ int) *builder2.FieldBuilder {
		return builder2.Field(NewFuncType(v))
	})
}
func NewMapperFields(fields []StructField) []builder2.ExprBuilder {
	return lo.Map(fields, func(p StructField, index int) builder2.ExprBuilder {
		return builder2.KeyValue(builder2.Ident(p.Names[0]),
			builder2.Selector(builder2.Ident("in"), builder2.Ident(p.Names[0])),
		).Decs(builder2.KeyValueDecs().Before(model.NewLine).After(model.NewLine))
	})
}
func NewStructFields(fields []StructField) []*builder2.FieldBuilder {
	return lo.Map(fields, func(v StructField, _ int) *builder2.FieldBuilder {
		return builder2.Field(NewIdent(v.Type), NewNames(v.Names)...)
	})
}
func NewFuncType(props FuncType) *builder2.FuncTypeBuilder {
	return builder2.NewFuncType(NewMethodTypes(props.In), NewMethodTypes(props.Out))
}
func NewFunc(props Func) *builder2.FuncBuilder {
	return builder2.NewFunc(props.Name, NewFuncType(props.FuncType), nil)
}
func NewIdent(props Ident) builder2.ExprBuilder {
	switch len(props.Generic) {
	case 0:
		var ret builder2.ExprBuilder = builder2.Ident(props.Name).Path(props.Path)
		if props.Ptr {
			ret = builder2.Star(ret)
		}
		return ret
	case 1:
		return builder2.Index(builder2.Selector(builder2.Ident(props.Path), builder2.Ident(props.Name)), NewIdent(props.Generic[0]))
	default:
		return builder2.IndexList(builder2.Selector(builder2.Ident(props.Path), builder2.Ident(props.Name)), NewIdents(props.Generic)...)
	}
}
func NewStruct(props Struct) *builder2.GenBuilder {
	return builder2.NewStruct(props.Name, NewStructFields(props.Fields)...)
}
func NewInterface(props Interface) *builder2.GenBuilder {
	return builder2.NewInterface(props.Name, NewFuncTypes(props.Methods)...)
}
func NewFile(props File) {

}
func NewEnum(props Enum) []builder2.DeclBuilder {
	return nil
}
func NewBuilder(props Struct) []builder2.DeclBuilder {
	name := props.Name + "Builder"
	ret := []builder2.DeclBuilder{
		builder2.NewType(name, NewIdent(Ident{Path: "x", Name: "Builder",
			Generic: []Ident{{Path: props.Path, Name: props.Name, Ptr: true, Generic: props.Generic}}})),
		NewConstructor(Ident{Name: name}),
	}
	ret = append(ret, lo.Map(props.Fields, func(item StructField, _ int) builder2.DeclBuilder {
		return builder2.NewFunc(item.Names[0], NewFuncType(FuncType{
			In:  []Field{{Names: []string{"v"}, Type: item.Type}},
			Out: []Field{{Type: Ident{Name: name, Ptr: true}}},
		}), builder2.Block(
			builder2.Assign(token.ASSIGN,
				[]builder2.ExprBuilder{builder2.Selector(builder2.Selector(
					builder2.Ident("b"), builder2.Ident("T")), builder2.Ident(item.Names[0]))},
				[]builder2.ExprBuilder{builder2.Ident("v")}),
			builder2.Return(builder2.Ident("b")).Decs(builder2.ReturnDecs().Before(model.NewLine)))).
			Recv(builder2.FieldList(builder2.Field(NewIdent(Ident{Name: props.Name + "Builder", Ptr: true}), builder2.Ident("b"))))
	})...)
	return ret
}
func NewConstructor(ident Ident) *builder2.FuncBuilder {
	return builder2.NewFunc("New"+ident.Name,
		builder2.FuncType().Results(builder2.FieldList(builder2.Field(builder2.Star(NewIdent(ident))))), builder2.Block(
			builder2.Return(builder2.Unary(token.AND, builder2.CompositeLit(NewIdent(ident)))).
				Decs(builder2.ReturnDecs().Before(model.NewLine))))
}
func NewMapper(props Mapper) []builder2.DeclBuilder {
	return []builder2.DeclBuilder{
		builder2.NewFunc("From"+props.From.Name,
			builder2.FuncType().
				Params(builder2.FieldList(builder2.Field(NewIdent(props.From.Ident), builder2.Ident("in")))).
				Results(builder2.FieldList(builder2.Field(NewIdent(props.To.Ident)))), builder2.Block(
				builder2.Return(builder2.CompositeLit(NewIdent(props.To.Ident), NewMapperFields(props.From.Fields)...)).
					Decs(builder2.ReturnDecs().Before(model.NewLine)))),
	}
}
func NewOptions(props Options) []builder2.DeclBuilder {
	ret := []builder2.DeclBuilder{}
	return ret
}
func NewService(props Strategy) []builder2.DeclBuilder {
	return []builder2.DeclBuilder{
		NewStruct(props.Struct),
	}
}
func NewDecls(props any) (res []builder2.DeclBuilder) {
	switch p := props.(type) {
	case Func:
		res = append(res, NewFunc(p))
	case Struct:
		res = append(res, NewStruct(p))
	case Interface:
		res = append(res, NewInterface(p))
	}
	return
}
