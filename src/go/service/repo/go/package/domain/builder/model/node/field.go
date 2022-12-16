package node

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type fieldBuilder Builder[*model.Field]

// Field FieldBuilder
func Field(t ExprBuilder, names ...IdentBuilder) FieldBuilder {
	return &fieldBuilder{T: &model.Field{
		Names: MapBuilder[*model.Ident, IdentBuilder](names),
		Type:  t.AsExpr(),
	}}
}
func (b *fieldBuilder) Decs(decs model.FieldDecorations) FieldBuilder {
	b.T.Decs = decs
	return b
}
func (b *fieldBuilder) Tag(tag IBuilder[*model.Lit]) FieldBuilder {
	b.T.Tag = tag.Build()
	return b
}
func (b *fieldBuilder) Build() *model.Field { return b.T }
