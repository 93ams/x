package expr

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type structBuilder Builder[*model.Struct]

func NewStruct(fields *model.FieldList) *model.Struct {
	return &model.Struct{Fields: fields}
}
func Struct(fields FieldListBuilder) StructBuilder {
	return &structBuilder{T: NewStruct(fields.Build())}
}
func (b *structBuilder) Decs(decs model.StructDecs) StructBuilder {
	b.T.Decs = decs
	return b
}
func (b *structBuilder) Incomplete(incomplete bool) StructBuilder {
	b.T.Incomplete = incomplete
	return b
}
func (b *structBuilder) Build() *model.Struct { return b.T }
func (b *structBuilder) AsExpr() model.Expr   { return b.T }
