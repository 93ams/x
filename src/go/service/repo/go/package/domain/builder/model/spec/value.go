package spec

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type valueBuilder Builder[*model.Value]

// Value ValueBuilder
func Value(t model.Expr, names ...IBuilder[*model.Ident]) ValueBuilder {
	return &valueBuilder{T: &model.Value{
		Names: MapBuilder[*model.Ident](names),
		Type:  t,
	}}
}
func (b *valueBuilder) Decs(decs model.ValueDecorations) ValueBuilder {
	b.T.Decs = decs
	return b
}
func (b *valueBuilder) Values(values ...ExprBuilder) ValueBuilder {
	b.T.Values = MapExprs(values)
	return b
}
func (b *valueBuilder) Build() *model.Value { return b.T }
func (b *valueBuilder) AsSpec() model.Spec  { return b.T }
