package expr

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type funcTypeBuilder Builder[*model.FuncType]

func NewFuncType() *model.FuncType {
	return &model.FuncType{}
}
func FuncType() FuncTypeBuilder {
	return &funcTypeBuilder{T: NewFuncType()}
}
func (b *funcTypeBuilder) Decs(decs model.FuncTypeDecorations) FuncTypeBuilder {
	b.T.Decs = decs
	return b
}
func (b *funcTypeBuilder) TypeParams(fields FieldListBuilder) FuncTypeBuilder {
	b.T.TypeParams = fields.Build()
	return b
}
func (b *funcTypeBuilder) Params(fields FieldListBuilder) FuncTypeBuilder {
	b.T.Params = fields.Build()
	return b
}
func (b *funcTypeBuilder) Results(fields FieldListBuilder) FuncTypeBuilder {
	b.T.Results = fields.Build()
	return b
}
func (b *funcTypeBuilder) Func(fn bool) FuncTypeBuilder {
	b.T.Func = fn
	return b
}
func (b *funcTypeBuilder) Build() *model.FuncType { return b.T }
func (b *funcTypeBuilder) AsExpr() model.Expr     { return b.T }
