package node

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type fileBuilder Builder[*model.File]

// File FileBuilder
func File(name IBuilder[*model.Ident], imports ...IBuilder[*model.Import]) FileBuilder {
	return &fileBuilder{T: &model.File{
		Name:    name.Build(),
		Imports: MapBuilder[*model.Import](imports),
	}}
}
func (b *fileBuilder) Decls(decls ...DeclBuilder) FileBuilder {
	b.T.Decls = append(b.T.Decls, MapDecls(decls)...)
	return b
}
func (b *fileBuilder) Decs(decs model.FileDecorations) FileBuilder {
	b.T.Decs = decs
	return b
}
func (b *fileBuilder) Scope(name *model.Scope) FileBuilder {
	b.T.Scope = name
	return b
}
func (b *fileBuilder) Build() *model.File { return b.T }
