package expr

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type (
	selectorBuilder     Builder[*model.Selector]
	selectorDecsBuilder Builder[model.SelectorDecs]
)

func NewSelector(x model.Expr, sel *model.Ident) *model.Selector {
	return &model.Selector{X: x, Sel: sel}
}
func Selector(x ExprBuilder, sel IdentBuilder) SelectorBuilder {
	return &selectorBuilder{T: NewSelector(x.AsExpr(), sel.Build())}
}
func (b *selectorBuilder) Decs(decs model.SelectorDecs) SelectorBuilder {
	b.T.Decs = decs
	return b
}
func (b *selectorBuilder) Build() *model.Selector { return b.T }
func (b *selectorBuilder) AsExpr() model.Expr     { return b.T }
