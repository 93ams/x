package builder

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper/model"
	"go/token"
)

type (
	SpecBuilder   interface{ AsSpec() model.Spec }
	ImportBuilder x.Builder[*model.Import]
	TypeBuilder   x.Builder[*model.Type]
	ValueBuilder  x.Builder[*model.Value]
)

func MapSpecs(builders []SpecBuilder) []model.Spec {
	return lo.Map(builders, func(item SpecBuilder, _ int) model.Spec { return item.AsSpec() })
}
func Import(name *IdentBuilder, path *LitBuilder) *ImportBuilder {
	return &ImportBuilder{T: model.NewImport(name.Build(), path.Build())}
}
func Type(t ExprBuilder) *TypeBuilder {
	return &TypeBuilder{T: model.NewType(t.AsExpr())}
}
func NewValue(t model.Expr, names []*model.Ident) *model.Value {
	return &model.Value{Names: names, Type: t}
}
func Value(t ExprBuilder, names ...x.IBuilder[*model.Ident]) *ValueBuilder {
	return &ValueBuilder{T: NewValue(t.AsExpr(), x.MapBuilder[*model.Ident](names))}
}

func (b *ImportBuilder) Decs(decs model.ImportDecorations) *ImportBuilder {
	b.T.Decs = decs
	return b
}
func (b *ImportBuilder) Name(path *IdentBuilder) *ImportBuilder {
	b.T.Name = path.Build()
	return b
}
func (b *ImportBuilder) Path(path *LitBuilder) *ImportBuilder {
	b.T.Path = path.Build()
	return b
}
func (b *TypeBuilder) Decs(decs model.TypeDecorations) *TypeBuilder {
	b.T.Decs = decs
	return b
}
func (b *TypeBuilder) Params(params *FieldListBuilder) *TypeBuilder {
	b.T.Params = params.Build()
	return b
}
func (b *TypeBuilder) Name(params *IdentBuilder) *TypeBuilder {
	b.T.Name = params.Build()
	return b
}
func (b *TypeBuilder) Assign(assign bool) *TypeBuilder {
	b.T.Assign = assign
	return b
}
func (b *ValueBuilder) Decs(decs model.ValueDecorations) *ValueBuilder {
	b.T.Decs = decs
	return b
}
func (b *ValueBuilder) Values(values ...ExprBuilder) *ValueBuilder {
	b.T.Values = MapExprs(values)
	return b
}

func (b *ImportBuilder) Build() *model.Import { return b.T }
func (b *ImportBuilder) AsSpec() model.Spec   { return b.T }
func (b *TypeBuilder) Build() *model.Type     { return b.T }
func (b *TypeBuilder) AsSpec() model.Spec     { return b.T }
func (b *ValueBuilder) Build() *model.Value   { return b.T }
func (b *ValueBuilder) AsSpec() model.Spec    { return b.T }

func NewType(name string, t ExprBuilder) *GenBuilder {
	return Gen(token.TYPE, Type(t).Name(Ident(name)))
}
func NewStruct(name string, fields ...*FieldBuilder) *GenBuilder {
	return NewType(name, Struct(FieldList(fields...)))
}
func NewInterface(name string, fields ...*FieldBuilder) *GenBuilder {
	return NewType(name, Interface(FieldList(fields...)))
}
