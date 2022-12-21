package builder

import (
	"github.com/tilau2328/x/src/go/package/x"
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
)

type (
	FieldBuilder     x.Builder[*model2.Field]
	FieldListBuilder x.Builder[*model2.FieldList]
	FileBuilder      x.Builder[*model2.File]
	PackageBuilder   x.Builder[*model2.Package]
)

func NewField(t model2.Expr, names []*model2.Ident) *model2.Field {
	return &model2.Field{Names: names, Type: t}
}
func Field(t ExprBuilder, names ...*IdentBuilder) *FieldBuilder {
	return &FieldBuilder{T: NewField(t.AsExpr(), x.MapBuilder[*model2.Ident](names))}
}
func NewFieldList(fields []*model2.Field) *model2.FieldList {
	return &model2.FieldList{List: fields}
}
func FieldList(fields ...*FieldBuilder) *FieldListBuilder {
	return &FieldListBuilder{T: NewFieldList(x.MapBuilder[*model2.Field](fields))}
}
func File(name x.IBuilder[*model2.Ident], imports ...x.IBuilder[*model2.Import]) *FileBuilder {
	return &FileBuilder{T: &model2.File{
		Name:    name.Build(),
		Imports: x.MapBuilder[*model2.Import](imports),
	}}
}
func Package(name string) *PackageBuilder {
	return &PackageBuilder{T: &model2.Package{Name: name}}
}

func (b *FieldBuilder) Decs(decs model2.FieldDecorations) *FieldBuilder {
	b.T.Decs = decs
	return b
}
func (b *FieldBuilder) Tag(tag x.IBuilder[*model2.Lit]) *FieldBuilder {
	b.T.Tag = tag.Build()
	return b
}
func (b *FieldListBuilder) Decs(decs model2.FieldListDecorations) *FieldListBuilder {
	b.T.Decs = decs
	return b
}
func (b *FileBuilder) Decls(decls ...DeclBuilder) *FileBuilder {
	b.T.Decls = append(b.T.Decls, MapDecls(decls)...)
	return b
}
func (b *FileBuilder) Decs(decs model2.FileDecorations) *FileBuilder {
	b.T.Decs = decs
	return b
}
func (b *FileBuilder) Scope(name *model2.Scope) *FileBuilder {
	b.T.Scope = name
	return b
}
func (b *PackageBuilder) Scope(scope *model2.Scope) *PackageBuilder {
	b.T.Scope = scope
	return b
}
func (b *PackageBuilder) Decs(decs model2.PackageDecorations) *PackageBuilder {
	b.T.Decs = decs
	return b
}
func (b *PackageBuilder) Files(files map[string]FileBuilder) *PackageBuilder {
	for k, v := range files {
		b.T.Files[k] = v.Build()
	}
	return b
}
func (b *PackageBuilder) Imports(imports map[string]*model2.Object) *PackageBuilder {
	for k, v := range imports {
		b.T.Imports[k] = v
	}
	return b
}

func (b *FieldBuilder) Build() *model2.Field         { return b.T }
func (b *FieldListBuilder) Build() *model2.FieldList { return b.T }
func (b *FileBuilder) Build() *model2.File           { return b.T }
func (b *PackageBuilder) Build() *model2.Package     { return b.T }
