package expr

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type identBuilder Builder[*model.Ident]

// Ident callBuilder
func Ident(name string) IdentBuilder {
	return &identBuilder{T: &model.Ident{Name: name}}
}
func (b *identBuilder) Decs(decs model.IdentDecorations) IdentBuilder {
	b.T.Decs = decs
	return b
}
func (b *identBuilder) Obj(obj *model.Object) IdentBuilder {
	b.T.Obj = obj
	return b
}
func (b *identBuilder) Path(path string) IdentBuilder {
	b.T.Path = path
	return b
}
func (b *identBuilder) Build() *model.Ident { return b.T }
func (b *identBuilder) AsExpr() model.Expr  { return b.T }
