package node

import (
	. "github.com/tilau2328/x/src/go/package/x"
	. "github.com/tilau2328/x/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type packageBuilder Builder[*model.Package]

// Package PackageBuilder
func Package(name string) PackageBuilder {
	return &packageBuilder{T: &model.Package{Name: name}}
}
func (b *packageBuilder) Scope(scope *model.Scope) PackageBuilder {
	b.T.Scope = scope
	return b
}
func (b *packageBuilder) Decs(decs model.PackageDecorations) PackageBuilder {
	b.T.Decs = decs
	return b
}
func (b *packageBuilder) Files(files map[string]FileBuilder) PackageBuilder {
	for k, v := range files {
		b.T.Files[k] = v.Build()
	}
	return b
}
func (b *packageBuilder) Imports(imports map[string]*model.Object) PackageBuilder {
	for k, v := range imports {
		b.T.Imports[k] = v
	}
	return b
}
func (b *packageBuilder) Build() *model.Package { return b.T }
