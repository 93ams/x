package expr

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type (
	compositeLitBuilder     Builder[*model.CompositeLit]
	compositeLitDecsBuilder Builder[model.CompositeLitDecs]
)

func NewCompositeLit(t model.Expr, elts []model.Expr) *model.CompositeLit {
	return &model.CompositeLit{Type: t, Elts: elts}
}
func CompositeLit(t ExprBuilder, elts ...ExprBuilder) CompositeLitBuilder {
	return &compositeLitBuilder{T: NewCompositeLit(t.AsExpr(), MapExprs(elts))}
}
func (b *compositeLitBuilder) Elts(exprs ...ExprBuilder) CompositeLitBuilder {
	b.T.Elts = append(b.T.Elts, MapExprs(exprs)...)
	return b
}
func (b *compositeLitBuilder) Incomplete(incomplete bool) CompositeLitBuilder {
	b.T.Incomplete = incomplete
	return b
}
func (b *compositeLitBuilder) Decs(decs CompositeLitDecsBuilder) CompositeLitBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *compositeLitBuilder) Build() *model.CompositeLit { return b.T }
func (b *compositeLitBuilder) AsExpr() model.Expr         { return b.T }

func CompositeLitDecs() CompositeLitDecsBuilder {
	return &compositeLitDecsBuilder{}
}
func (b *compositeLitDecsBuilder) Before(d model.SpaceType) CompositeLitDecsBuilder {
	b.T.Before = d
	return b
}
func (b *compositeLitDecsBuilder) After(d model.SpaceType) CompositeLitDecsBuilder {
	b.T.After = d
	return b
}
func (b *compositeLitDecsBuilder) Start(d model.Decs) CompositeLitDecsBuilder {
	b.T.Start = d
	return b
}
func (b *compositeLitDecsBuilder) End(d model.Decs) CompositeLitDecsBuilder {
	b.T.End = d
	return b
}
func (b *compositeLitDecsBuilder) Build() model.CompositeLitDecs { return b.T }
