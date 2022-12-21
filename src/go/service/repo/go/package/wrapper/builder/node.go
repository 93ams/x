package builder

import (
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper/model"
)

type (
	FieldBuilder     x.Builder[*model.Field]
	FieldListBuilder x.Builder[*model.FieldList]
	FileBuilder      x.Builder[*model.File]
	PackageBuilder   x.Builder[*model.Package]
)

func NewField(t model.Expr, names []*model.Ident) *model.Field {
	return &model.Field{Names: names, Type: t}
}
func Field(t ExprBuilder, names ...*IdentBuilder) *FieldBuilder {
	return &FieldBuilder{T: NewField(t.AsExpr(), x.MapBuilder[*model.Ident](names))}
}
func NewFieldList(fields []*model.Field) *model.FieldList {
	return &model.FieldList{List: fields}
}
func FieldList(fields ...*FieldBuilder) *FieldListBuilder {
	return &FieldListBuilder{T: NewFieldList(x.MapBuilder[*model.Field](fields))}
}
func File(name x.IBuilder[*model.Ident], imports ...x.IBuilder[*model.Import]) *FileBuilder {
	return &FileBuilder{T: &model.File{
		Name:    name.Build(),
		Imports: x.MapBuilder[*model.Import](imports),
	}}
}
func Package(name string) *PackageBuilder {
	return &PackageBuilder{T: &model.Package{Name: name}}
}

func (b *FieldBuilder) Decs(decs model.FieldDecorations) *FieldBuilder {
	b.T.Decs = decs
	return b
}
func (b *FieldBuilder) Tag(tag x.IBuilder[*model.Lit]) *FieldBuilder {
	b.T.Tag = tag.Build()
	return b
}
func (b *FieldListBuilder) Decs(decs model.FieldListDecorations) *FieldListBuilder {
	b.T.Decs = decs
	return b
}
func (b *FileBuilder) Decls(decls ...DeclBuilder) *FileBuilder {
	b.T.Decls = append(b.T.Decls, MapDecls(decls)...)
	return b
}
func (b *FileBuilder) Decs(decs model.FileDecorations) *FileBuilder {
	b.T.Decs = decs
	return b
}
func (b *FileBuilder) Scope(name *model.Scope) *FileBuilder {
	b.T.Scope = name
	return b
}
func (b *PackageBuilder) Scope(scope *model.Scope) *PackageBuilder {
	b.T.Scope = scope
	return b
}
func (b *PackageBuilder) Decs(decs model.PackageDecorations) *PackageBuilder {
	b.T.Decs = decs
	return b
}
func (b *PackageBuilder) Files(files map[string]FileBuilder) *PackageBuilder {
	for k, v := range files {
		b.T.Files[k] = v.Build()
	}
	return b
}
func (b *PackageBuilder) Imports(imports map[string]*model.Object) *PackageBuilder {
	for k, v := range imports {
		b.T.Imports[k] = v
	}
	return b
}

func (b *FieldBuilder) Build() *model.Field         { return b.T }
func (b *FieldListBuilder) Build() *model.FieldList { return b.T }
func (b *FileBuilder) Build() *model.File           { return b.T }
func (b *PackageBuilder) Build() *model.Package     { return b.T }
