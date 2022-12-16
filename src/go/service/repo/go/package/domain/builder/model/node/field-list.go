package node

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type fieldListBuilder Builder[*model.FieldList]

// FieldList FieldListBuilder
func FieldList(fields ...FieldBuilder) FieldListBuilder {
	return &fieldListBuilder{T: &model.FieldList{
		List: MapBuilder[*model.Field](fields),
	}}
}
func (b *fieldListBuilder) Decs(decs model.FieldListDecorations) FieldListBuilder {
	b.T.Decs = decs
	return b
}
func (b *fieldListBuilder) Build() *model.FieldList { return b.T }
