package spec

import "github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"

type importBuilder Builder[*model.Import]

func NewImport(name *model.Ident) *model.Import {
	return &model.Import{Name: name, Path: &model.Lit{}}
}
func Import(name IdentBuilder) ImportBuilder {
	return &importBuilder{T: NewImport(name.Build())}
}
func (b *importBuilder) Decs(decs model.ImportDecorations) ImportBuilder {
	b.T.Decs = decs
	return b
}
func (b *importBuilder) Path(path LitBuilder) ImportBuilder {
	b.T.Path = path.Build()
	return b
}
func (b *importBuilder) Build() *model.Import { return b.T }
func (b *importBuilder) AsSpec() model.Spec   { return b.T }
