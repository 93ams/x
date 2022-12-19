package builder

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/model"
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
func NewImport(name *model.Ident) *model.Import {
	return &model.Import{Name: name, Path: &model.Lit{}}
}
func Import(name IdentBuilder) *ImportBuilder {
	return &ImportBuilder{T: NewImport(name.Build())}
}
func NewType(t model.Expr) *model.Type { return &model.Type{Type: t} }
func Type(t ExprBuilder) *TypeBuilder {
	return &TypeBuilder{T: NewType(t.AsExpr())}
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
