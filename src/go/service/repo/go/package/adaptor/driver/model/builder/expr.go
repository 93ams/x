package builder

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	"go/token"
)

type (
	ExprBuilder             interface{ AsExpr() model2.Expr }
	CallBuilder             x.Builder[*model2.Call]
	CompositeLitBuilder     x.Builder[*model2.CompositeLit]
	CompositeLitDecsBuilder x.Builder[model2.CompositeLitDecs]
	FuncTypeBuilder         x.Builder[*model2.FuncType]
	IdentBuilder            x.Builder[*model2.Ident]
	IndexBuilder            x.Builder[*model2.Index]
	IndexListBuilder        x.Builder[*model2.IndexList]
	InterfaceBuilder        x.Builder[*model2.Interface]
	KeyValueBuilder         x.Builder[*model2.KeyValue]
	KeyValueDecsBuilder     x.Builder[model2.KeyValueDecs]
	LitBuilder              x.Builder[*model2.Lit]
	SelectorBuilder         x.Builder[*model2.Selector]
	SelectorDecsBuilder     x.Builder[model2.SelectorDecs]
	StarBuilder             x.Builder[*model2.Star]
	StructBuilder           x.Builder[*model2.Struct]
	UnaryBuilder            x.Builder[*model2.Unary]
)

func MapExprs(builders []ExprBuilder) []model2.Expr {
	return lo.Map(builders, func(item ExprBuilder, _ int) model2.Expr { return item.AsExpr() })
}
func Call(fun ExprBuilder, args ...ExprBuilder) *CallBuilder {
	return &CallBuilder{T: model2.NewCall(fun.AsExpr(), MapExprs(args))}
}
func CompositeLit(t ExprBuilder, elts ...ExprBuilder) *CompositeLitBuilder {
	return &CompositeLitBuilder{T: model2.NewCompositeLit(t.AsExpr(), MapExprs(elts))}
}
func CompositeLitDecs() *CompositeLitDecsBuilder { return &CompositeLitDecsBuilder{} }
func FuncType() *FuncTypeBuilder                 { return &FuncTypeBuilder{T: model2.NewFuncType()} }
func Ident(name string) *IdentBuilder            { return &IdentBuilder{T: &model2.Ident{Name: name}} }
func Interface(methods *FieldListBuilder) *InterfaceBuilder {
	return &InterfaceBuilder{T: &model2.Interface{Methods: methods.Build()}}
}
func KeyValue(t ExprBuilder, value ExprBuilder) *KeyValueBuilder {
	return &KeyValueBuilder{T: model2.NewKeyValue(t.AsExpr(), value.AsExpr())}
}
func KeyValueDecs() *KeyValueDecsBuilder { return &KeyValueDecsBuilder{} }
func Lit(kind token.Token, value string) *LitBuilder {
	return &LitBuilder{T: &model2.Lit{Value: value, Kind: kind}}
}
func IndexList(x ExprBuilder, indicies ...ExprBuilder) *IndexListBuilder {
	return &IndexListBuilder{T: &model2.IndexList{X: x.AsExpr(), Indices: MapExprs(indicies)}}
}
func Selector(x ExprBuilder, sel *IdentBuilder) *SelectorBuilder {
	return &SelectorBuilder{T: model2.NewSelector(x.AsExpr(), sel.Build())}
}
func Star(x ExprBuilder) *StarBuilder {
	return &StarBuilder{T: &model2.Star{X: x.AsExpr()}}
}
func Struct(fields *FieldListBuilder) *StructBuilder {
	return &StructBuilder{T: model2.NewStruct(fields.Build())}
}
func Index(x ExprBuilder, i ExprBuilder) *IndexBuilder {
	return &IndexBuilder{T: model2.NewIndex(x.AsExpr(), i.AsExpr())}
}
func Unary(op token.Token, x ExprBuilder) *UnaryBuilder {
	return &UnaryBuilder{T: &model2.Unary{X: x.AsExpr(), Op: op}}
}
func (b *CallBuilder) Decs(decs model2.CallDecs) *CallBuilder {
	b.T.Decs = decs
	return b
}
func (b *CallBuilder) Ellipsis(ellipsis bool) *CallBuilder {
	b.T.Ellipsis = ellipsis
	return b
}
func (b *CompositeLitBuilder) Elts(exprs ...ExprBuilder) *CompositeLitBuilder {
	b.T.Elts = append(b.T.Elts, MapExprs(exprs)...)
	return b
}
func (b *CompositeLitBuilder) Incomplete(incomplete bool) *CompositeLitBuilder {
	b.T.Incomplete = incomplete
	return b
}
func (b *CompositeLitBuilder) Decs(decs *CompositeLitDecsBuilder) *CompositeLitBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *CompositeLitDecsBuilder) Before(d model2.SpaceType) *CompositeLitDecsBuilder {
	b.T.Before = d
	return b
}
func (b *CompositeLitDecsBuilder) After(d model2.SpaceType) *CompositeLitDecsBuilder {
	b.T.After = d
	return b
}
func (b *CompositeLitDecsBuilder) Start(d model2.Decs) *CompositeLitDecsBuilder {
	b.T.Start = d
	return b
}
func (b *CompositeLitDecsBuilder) End(d model2.Decs) *CompositeLitDecsBuilder {
	b.T.End = d
	return b
}
func (b *FuncTypeBuilder) Decs(decs model2.FuncTypeDecorations) *FuncTypeBuilder {
	b.T.Decs = decs
	return b
}
func (b *FuncTypeBuilder) TypeParams(fields *FieldListBuilder) *FuncTypeBuilder {
	b.T.TypeParams = fields.Build()
	return b
}
func (b *FuncTypeBuilder) Params(fields *FieldListBuilder) *FuncTypeBuilder {
	b.T.Params = fields.Build()
	return b
}
func (b *FuncTypeBuilder) Results(fields *FieldListBuilder) *FuncTypeBuilder {
	b.T.Results = fields.Build()
	return b
}
func (b *FuncTypeBuilder) Func(fn bool) *FuncTypeBuilder {
	b.T.Func = fn
	return b
}
func (b *IdentBuilder) Decs(decs model2.IdentDecorations) *IdentBuilder {
	b.T.Decs = decs
	return b
}
func (b *IdentBuilder) Obj(obj *model2.Object) *IdentBuilder {
	b.T.Obj = obj
	return b
}
func (b *IdentBuilder) Path(path string) *IdentBuilder {
	b.T.Path = path
	return b
}
func (b *InterfaceBuilder) Decs(decs model2.InterfaceDecorations) *InterfaceBuilder {
	b.T.Decs = decs
	return b
}
func (b *InterfaceBuilder) Incomplete(incomplete bool) *InterfaceBuilder {
	b.T.Incomplete = incomplete
	return b
}
func (b *KeyValueBuilder) Decs(decs *KeyValueDecsBuilder) *KeyValueBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *KeyValueDecsBuilder) Before(d model2.SpaceType) *KeyValueDecsBuilder {
	b.T.Before = d
	return b
}
func (b *KeyValueDecsBuilder) After(d model2.SpaceType) *KeyValueDecsBuilder {
	b.T.After = d
	return b
}
func (b *KeyValueDecsBuilder) Start(d model2.Decs) *KeyValueDecsBuilder {
	b.T.Start = d
	return b
}
func (b *KeyValueDecsBuilder) End(d model2.Decs) *KeyValueDecsBuilder {
	b.T.End = d
	return b
}
func (b *LitBuilder) Decs(decs model2.LitDecs) *LitBuilder {
	b.T.Decs = decs
	return b
}
func (b *SelectorBuilder) Decs(decs model2.SelectorDecs) *SelectorBuilder {
	b.T.Decs = decs
	return b
}
func (b *StructBuilder) Decs(decs model2.StructDecs) *StructBuilder {
	b.T.Decs = decs
	return b
}
func (b *StructBuilder) Incomplete(incomplete bool) *StructBuilder {
	b.T.Incomplete = incomplete
	return b
}

func (b *CallBuilder) Build() *model2.Call                        { return b.T }
func (b *CallBuilder) AsExpr() model2.Expr                        { return b.T }
func (b *CompositeLitBuilder) Build() *model2.CompositeLit        { return b.T }
func (b *CompositeLitBuilder) AsExpr() model2.Expr                { return b.T }
func (b *CompositeLitDecsBuilder) Build() model2.CompositeLitDecs { return b.T }
func (b *FuncTypeBuilder) Build() *model2.FuncType                { return b.T }
func (b *FuncTypeBuilder) AsExpr() model2.Expr                    { return b.T }
func (b *IdentBuilder) Build() *model2.Ident                      { return b.T }
func (b *IdentBuilder) AsExpr() model2.Expr                       { return b.T }
func (b *IndexBuilder) Build() *model2.Index                      { return b.T }
func (b *IndexBuilder) AsExpr() model2.Expr                       { return b.T }
func (b *IndexListBuilder) Build() *model2.IndexList              { return b.T }
func (b *IndexListBuilder) AsExpr() model2.Expr                   { return b.T }
func (b *InterfaceBuilder) Build() *model2.Interface              { return b.T }
func (b *InterfaceBuilder) AsExpr() model2.Expr                   { return b.T }
func (b *KeyValueBuilder) Build() *model2.KeyValue                { return b.T }
func (b *KeyValueBuilder) AsExpr() model2.Expr                    { return b.T }
func (b *KeyValueDecsBuilder) Build() model2.KeyValueDecs         { return b.T }
func (b *LitBuilder) Build() *model2.Lit                          { return b.T }
func (b *LitBuilder) AsExpr() model2.Expr                         { return b.T }
func (b *SelectorBuilder) Build() *model2.Selector                { return b.T }
func (b *SelectorBuilder) AsExpr() model2.Expr                    { return b.T }
func (b *StarBuilder) Build() *model2.Star                        { return b.T }
func (b *StarBuilder) AsExpr() model2.Expr                        { return b.T }
func (b *StructBuilder) Build() *model2.Struct                    { return b.T }
func (b *StructBuilder) AsExpr() model2.Expr                      { return b.T }
func (b *UnaryBuilder) Build() *model2.Unary                      { return b.T }
func (b *UnaryBuilder) AsExpr() model2.Expr                       { return b.T }

func NewFuncType(params, results []*FieldBuilder) *FuncTypeBuilder {
	return FuncType().Params(FieldList(params...)).Results(FieldList(results...))
}
