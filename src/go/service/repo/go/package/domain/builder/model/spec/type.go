package spec

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type typeBuilder Builder[*model.Type]

func NewType(t model.Expr) *model.Type { return &model.Type{Type: t} }
func Type(t ExprBuilder) TypeBuilder {
	return &typeBuilder{T: NewType(t.AsExpr())}
}
func (b *typeBuilder) Decs(decs model.TypeDecorations) TypeBuilder {
	b.T.Decs = decs
	return b
}
func (b *typeBuilder) Params(params FieldListBuilder) TypeBuilder {
	b.T.Params = params.Build()
	return b
}
func (b *typeBuilder) Name(params IdentBuilder) TypeBuilder {
	b.T.Name = params.Build()
	return b
}
func (b *typeBuilder) Assign(assign bool) TypeBuilder {
	b.T.Assign = assign
	return b
}
func (b *typeBuilder) Build() *model.Type { return b.T }
func (b *typeBuilder) AsSpec() model.Spec { return b.T }
