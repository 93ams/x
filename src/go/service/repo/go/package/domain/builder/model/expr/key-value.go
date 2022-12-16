package expr

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type (
	keyValueBuilder     Builder[*model.KeyValue]
	keyValueDecsBuilder Builder[model.KeyValueDecs]
)

func NewKeyValue(key model.Expr, value model.Expr) *model.KeyValue {
	return &model.KeyValue{Key: key, Value: value}
}
func KeyValue(t ExprBuilder, value ExprBuilder) KeyValueBuilder {
	return &keyValueBuilder{T: NewKeyValue(t.AsExpr(), value.AsExpr())}
}

func (b *keyValueBuilder) Decs(decs KeyValueDecsBuilder) KeyValueBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *keyValueBuilder) Build() *model.KeyValue { return b.T }
func (b *keyValueBuilder) AsExpr() model.Expr     { return b.T }

func KeyValueDecs() KeyValueDecsBuilder {
	return &keyValueDecsBuilder{}
}
func (b *keyValueDecsBuilder) Before(d model.SpaceType) KeyValueDecsBuilder {
	b.T.Before = d
	return b
}
func (b *keyValueDecsBuilder) After(d model.SpaceType) KeyValueDecsBuilder {
	b.T.After = d
	return b
}
func (b *keyValueDecsBuilder) Start(d model.Decorations) KeyValueDecsBuilder {
	b.T.Start = d
	return b
}
func (b *keyValueDecsBuilder) End(d model.Decorations) KeyValueDecsBuilder {
	b.T.End = d
	return b
}
func (b *keyValueDecsBuilder) Build() model.KeyValueDecs { return b.T }
