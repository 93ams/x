package builder

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper/model"
	"go/token"
)

type (
	ExprBuilder             interface{ AsExpr() model.Expr }
	CallBuilder             x.Builder[*model.Call]
	CompositeLitBuilder     x.Builder[*model.CompositeLit]
	CompositeLitDecsBuilder x.Builder[model.CompositeLitDecs]
	FuncTypeBuilder         x.Builder[*model.FuncType]
	IdentBuilder            x.Builder[*model.Ident]
	IndexBuilder            x.Builder[*model.Index]
	IndexListBuilder        x.Builder[*model.IndexList]
	InterfaceBuilder        x.Builder[*model.Interface]
	KeyValueBuilder         x.Builder[*model.KeyValue]
	KeyValueDecsBuilder     x.Builder[model.KeyValueDecs]
	LitBuilder              x.Builder[*model.Lit]
	SelectorBuilder         x.Builder[*model.Selector]
	SelectorDecsBuilder     x.Builder[model.SelectorDecs]
	StarBuilder             x.Builder[*model.Star]
	StructBuilder           x.Builder[*model.Struct]
	UnaryBuilder            x.Builder[*model.Unary]
)

func MapExprs(builders []ExprBuilder) []model.Expr {
	return lo.Map(builders, func(item ExprBuilder, _ int) model.Expr { return item.AsExpr() })
}
func Call(fun ExprBuilder, args ...ExprBuilder) *CallBuilder {
	return &CallBuilder{T: model.NewCall(fun.AsExpr(), MapExprs(args))}
}
func CompositeLit(t ExprBuilder, elts ...ExprBuilder) *CompositeLitBuilder {
	return &CompositeLitBuilder{T: model.NewCompositeLit(t.AsExpr(), MapExprs(elts))}
}
func CompositeLitDecs() *CompositeLitDecsBuilder { return &CompositeLitDecsBuilder{} }
func FuncType() *FuncTypeBuilder                 { return &FuncTypeBuilder{T: model.NewFuncType()} }
func Ident(name string) *IdentBuilder            { return &IdentBuilder{T: &model.Ident{Name: name}} }
func Interface(methods *FieldListBuilder) *InterfaceBuilder {
	return &InterfaceBuilder{T: &model.Interface{Methods: methods.Build()}}
}
func KeyValue(t ExprBuilder, value ExprBuilder) *KeyValueBuilder {
	return &KeyValueBuilder{T: model.NewKeyValue(t.AsExpr(), value.AsExpr())}
}
func KeyValueDecs() *KeyValueDecsBuilder { return &KeyValueDecsBuilder{} }
func Lit(kind token.Token, value string) *LitBuilder {
	return &LitBuilder{T: &model.Lit{Value: value, Kind: kind}}
}
func IndexList(x ExprBuilder, indicies ...ExprBuilder) *IndexListBuilder {
	return &IndexListBuilder{T: &model.IndexList{X: x.AsExpr(), Indices: MapExprs(indicies)}}
}
func Selector(x ExprBuilder, sel *IdentBuilder) *SelectorBuilder {
	return &SelectorBuilder{T: model.NewSelector(x.AsExpr(), sel.Build())}
}
func Star(x ExprBuilder) *StarBuilder {
	return &StarBuilder{T: &model.Star{X: x.AsExpr()}}
}
func Struct(fields *FieldListBuilder) *StructBuilder {
	return &StructBuilder{T: model.NewStruct(fields.Build())}
}
func Index(x ExprBuilder, i ExprBuilder) *IndexBuilder {
	return &IndexBuilder{T: model.NewIndex(x.AsExpr(), i.AsExpr())}
}
func Unary(op token.Token, x ExprBuilder) *UnaryBuilder {
	return &UnaryBuilder{T: &model.Unary{X: x.AsExpr(), Op: op}}
}
func (b *CallBuilder) Decs(decs model.CallDecs) *CallBuilder {
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
func (b *CompositeLitDecsBuilder) Before(d model.SpaceType) *CompositeLitDecsBuilder {
	b.T.Before = d
	return b
}
func (b *CompositeLitDecsBuilder) After(d model.SpaceType) *CompositeLitDecsBuilder {
	b.T.After = d
	return b
}
func (b *CompositeLitDecsBuilder) Start(d model.Decs) *CompositeLitDecsBuilder {
	b.T.Start = d
	return b
}
func (b *CompositeLitDecsBuilder) End(d model.Decs) *CompositeLitDecsBuilder {
	b.T.End = d
	return b
}
func (b *FuncTypeBuilder) Decs(decs model.FuncTypeDecorations) *FuncTypeBuilder {
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
func (b *IdentBuilder) Decs(decs model.IdentDecorations) *IdentBuilder {
	b.T.Decs = decs
	return b
}
func (b *IdentBuilder) Obj(obj *model.Object) *IdentBuilder {
	b.T.Obj = obj
	return b
}
func (b *IdentBuilder) Path(path string) *IdentBuilder {
	b.T.Path = path
	return b
}
func (b *InterfaceBuilder) Decs(decs model.InterfaceDecorations) *InterfaceBuilder {
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
func (b *KeyValueDecsBuilder) Before(d model.SpaceType) *KeyValueDecsBuilder {
	b.T.Before = d
	return b
}
func (b *KeyValueDecsBuilder) After(d model.SpaceType) *KeyValueDecsBuilder {
	b.T.After = d
	return b
}
func (b *KeyValueDecsBuilder) Start(d model.Decs) *KeyValueDecsBuilder {
	b.T.Start = d
	return b
}
func (b *KeyValueDecsBuilder) End(d model.Decs) *KeyValueDecsBuilder {
	b.T.End = d
	return b
}
func (b *LitBuilder) Decs(decs model.LitDecs) *LitBuilder {
	b.T.Decs = decs
	return b
}
func (b *SelectorBuilder) Decs(decs model.SelectorDecs) *SelectorBuilder {
	b.T.Decs = decs
	return b
}
func (b *StructBuilder) Decs(decs model.StructDecs) *StructBuilder {
	b.T.Decs = decs
	return b
}
func (b *StructBuilder) Incomplete(incomplete bool) *StructBuilder {
	b.T.Incomplete = incomplete
	return b
}

func (b *CallBuilder) Build() *model.Call                        { return b.T }
func (b *CallBuilder) AsExpr() model.Expr                        { return b.T }
func (b *CompositeLitBuilder) Build() *model.CompositeLit        { return b.T }
func (b *CompositeLitBuilder) AsExpr() model.Expr                { return b.T }
func (b *CompositeLitDecsBuilder) Build() model.CompositeLitDecs { return b.T }
func (b *FuncTypeBuilder) Build() *model.FuncType                { return b.T }
func (b *FuncTypeBuilder) AsExpr() model.Expr                    { return b.T }
func (b *IdentBuilder) Build() *model.Ident                      { return b.T }
func (b *IdentBuilder) AsExpr() model.Expr                       { return b.T }
func (b *IndexBuilder) Build() *model.Index                      { return b.T }
func (b *IndexBuilder) AsExpr() model.Expr                       { return b.T }
func (b *IndexListBuilder) Build() *model.IndexList              { return b.T }
func (b *IndexListBuilder) AsExpr() model.Expr                   { return b.T }
func (b *InterfaceBuilder) Build() *model.Interface              { return b.T }
func (b *InterfaceBuilder) AsExpr() model.Expr                   { return b.T }
func (b *KeyValueBuilder) Build() *model.KeyValue                { return b.T }
func (b *KeyValueBuilder) AsExpr() model.Expr                    { return b.T }
func (b *KeyValueDecsBuilder) Build() model.KeyValueDecs         { return b.T }
func (b *LitBuilder) Build() *model.Lit                          { return b.T }
func (b *LitBuilder) AsExpr() model.Expr                         { return b.T }
func (b *SelectorBuilder) Build() *model.Selector                { return b.T }
func (b *SelectorBuilder) AsExpr() model.Expr                    { return b.T }
func (b *StarBuilder) Build() *model.Star                        { return b.T }
func (b *StarBuilder) AsExpr() model.Expr                        { return b.T }
func (b *StructBuilder) Build() *model.Struct                    { return b.T }
func (b *StructBuilder) AsExpr() model.Expr                      { return b.T }
func (b *UnaryBuilder) Build() *model.Unary                      { return b.T }
func (b *UnaryBuilder) AsExpr() model.Expr                       { return b.T }

func NewFuncType(params, results []*FieldBuilder) *FuncTypeBuilder {
	return FuncType().Params(FieldList(params...)).Results(FieldList(results...))
}
