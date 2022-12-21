package mapper

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/builder"
	. "github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"go/token"
)

func NewNames(names []string) []*builder.IdentBuilder {
	return lo.Map(names, func(item string, _ int) *builder.IdentBuilder { return builder.Ident(item) })
}
func NewIdents(names []Ident) []builder.ExprBuilder {
	return lo.Map(names, func(item Ident, _ int) builder.ExprBuilder { return NewIdent(item) })
}
func NewMethodTypes(methods []Field) []*builder.FieldBuilder {
	return lo.Map(methods, func(item Field, _ int) *builder.FieldBuilder {
		return builder.Field(NewIdent(item.Type), NewNames(item.Names)...)
	})
}
func NewFuncTypes(methods []FuncType) []*builder.FieldBuilder {
	return lo.Map(methods, func(v FuncType, _ int) *builder.FieldBuilder {
		return builder.Field(NewFuncType(v))
	})
}
func NewMapperFields(fields []Field) []builder.ExprBuilder {
	return lo.Map(fields, func(p Field, index int) builder.ExprBuilder {
		return builder.KeyValue(builder.Ident(p.Names[0]),
			builder.Selector(builder.Ident("in"), builder.Ident(p.Names[0])),
		).Decs(builder.KeyValueDecs().Before(model.NewLine).After(model.NewLine))
	})
}
func NewFields(fields []Field) []*builder.FieldBuilder {
	return lo.Map(fields, func(v Field, _ int) *builder.FieldBuilder {
		return builder.Field(NewIdent(v.Type), NewNames(v.Names)...)
	})
}
func NewFuncType(props FuncType) *builder.FuncTypeBuilder {
	return builder.NewFuncType(NewMethodTypes(props.In), NewMethodTypes(props.Out))
}
func NewFunc(props Func) *builder.FuncBuilder {
	return builder.NewFunc(props.Name, NewFuncType(props.FuncType), nil)
}
func NewIdent(props Ident) builder.ExprBuilder {
	switch len(props.Generic) {
	case 0:
		var ret builder.ExprBuilder = builder.Ident(props.Name).Path(props.Path)
		if props.Ptr {
			ret = builder.Star(ret)
		}
		return ret
	case 1:
		return builder.Index(builder.Selector(builder.Ident(props.Path), builder.Ident(props.Name)), NewIdent(props.Generic[0]))
	default:
		return builder.IndexList(builder.Selector(builder.Ident(props.Path), builder.Ident(props.Name)), NewIdents(props.Generic)...)
	}
}
func NewStruct(props Struct) *builder.GenBuilder {
	return builder.NewStruct(props.Name, NewFields(props.Fields)...)
}
func NewInterface(props Interface) *builder.GenBuilder {
	return builder.NewInterface(props.Name, NewFuncTypes(props.Methods)...)
}
func NewFile(props File) {

}
func NewEnum(props Enum) []builder.DeclBuilder {
	return nil
}
func NewBuilder(props Struct) []builder.DeclBuilder {
	name := props.Name + "Builder"
	ret := []builder.DeclBuilder{
		builder.NewType(name, NewIdent(Ident{Path: "x", Name: "Builder",
			Generic: []Ident{{Path: props.Path, Name: props.Name, Ptr: true, Generic: props.Generic}}})),
		NewConstructor(Ident{Name: name}),
	}
	ret = append(ret, lo.Map(props.Fields, func(item Field, _ int) builder.DeclBuilder {
		return builder.NewFunc(item.Names[0], NewFuncType(FuncType{
			In:  []Field{{Names: []string{"v"}, Type: item.Type}},
			Out: []Field{{Type: Ident{Name: name, Ptr: true}}},
		}), builder.Block(
			builder.Assign(token.ASSIGN,
				[]builder.ExprBuilder{builder.Selector(builder.Selector(
					builder.Ident("b"), builder.Ident("T")), builder.Ident(item.Names[0]))},
				[]builder.ExprBuilder{builder.Ident("v")}),
			builder.Return(builder.Ident("b")).Decs(builder.ReturnDecs().Before(model.NewLine)))).
			Recv(builder.FieldList(builder.Field(NewIdent(Ident{Name: props.Name + "Builder", Ptr: true}), builder.Ident("b"))))
	})...)
	return ret
}
func NewService(props Service) []builder.DeclBuilder {
	ret := []builder.DeclBuilder{
		NewStruct(props.Struct),
		NewConstructor(Ident{Name: props.Name}),
	}
	ret = append(ret, lo.Map(props.Methods, func(item Func, _ int) builder.DeclBuilder {
		return builder.NewFunc(item.Name, NewFuncType(item.FuncType), builder.Block()).
			Recv(builder.FieldList(builder.Field(NewIdent(Ident{Name: props.Name, Ptr: true}), builder.Ident("b"))))
	})...)
	return ret
}
func NewConstructor(ident Ident) *builder.FuncBuilder {
	return builder.NewFunc("New"+ident.Name,
		builder.FuncType().Results(builder.FieldList(builder.Field(builder.Star(NewIdent(ident))))), builder.Block(
			builder.Return(builder.Unary(token.AND, builder.CompositeLit(NewIdent(ident)))).
				Decs(builder.ReturnDecs().Before(model.NewLine))))
}
func NewMapper(props Mapper) []builder.DeclBuilder {
	return []builder.DeclBuilder{
		builder.NewFunc("From"+props.From.Name,
			builder.FuncType().
				Params(builder.FieldList(builder.Field(NewIdent(props.From.Ident), builder.Ident("in")))).
				Results(builder.FieldList(builder.Field(NewIdent(props.To.Ident)))), builder.Block(
				builder.Return(builder.CompositeLit(NewIdent(props.To.Ident), NewMapperFields(props.From.Fields)...)).
					Decs(builder.ReturnDecs().Before(model.NewLine)))),
	}
}
func NewOptions(props Options) []builder.DeclBuilder {
	ret := []builder.DeclBuilder{}
	return ret
}
func NewDecls(props any) (res []builder.DeclBuilder) {
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
