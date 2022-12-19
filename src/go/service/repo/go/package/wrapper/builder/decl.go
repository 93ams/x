package builder

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/model"
	"go/token"
)

type (
	DeclBuilder     interface{ AsDecl() model.Decl }
	GenBuilder      x.Builder[*model.Gen]
	FuncBuilder     x.Builder[*model.Func]
	FuncDecsBuilder x.Builder[model.FuncDecs]
)

func MapDecls(builders []DeclBuilder) []model.Decl {
	return lo.Map(builders, func(item DeclBuilder, _ int) model.Decl { return item.AsDecl() })
}
func NewGen(t token.Token, specs []model.Spec) *model.Gen { return &model.Gen{Tok: t, Specs: specs} }
func Gen(t token.Token, specs ...SpecBuilder) *GenBuilder {
	return &GenBuilder{T: NewGen(t, MapSpecs(specs))}
}
func NewFunc(name *model.Ident) *model.Func {
	return &model.Func{Name: name, Type: &model.FuncType{}, Body: &model.Block{}}
}
func Func(name *IdentBuilder) *FuncBuilder { return &FuncBuilder{T: NewFunc(name.Build())} }
func FuncDecs() *FuncDecsBuilder           { return &FuncDecsBuilder{} }
func (b *GenBuilder) Decs(decs model.GenDecorations) *GenBuilder {
	b.T.Decs = decs
	return b
}
func (b *FuncBuilder) Decs(decs *FuncDecsBuilder) *FuncBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *FuncBuilder) Recv(fields *FieldListBuilder) *FuncBuilder {
	b.T.Recv = fields.Build()
	return b
}
func (b *FuncBuilder) Type(t *FuncTypeBuilder) *FuncBuilder {
	b.T.Type = t.Build()
	return b
}
func (b *FuncBuilder) Body(body *BlockBuilder) *FuncBuilder {
	b.T.Body = body.Build()
	return b
}
func (f *FuncDecsBuilder) TypeParams(d model.Decs) *FuncDecsBuilder {
	f.T.TypeParams = d
	return f
}
func (f *FuncDecsBuilder) Results(d model.Decs) *FuncDecsBuilder {
	f.T.Results = d
	return f
}
func (f *FuncDecsBuilder) Params(d model.Decs) *FuncDecsBuilder {
	f.T.Params = d
	return f
}
func (f *FuncDecsBuilder) Start(d model.Decs) *FuncDecsBuilder {
	f.T.Start = d
	return f
}
func (f *FuncDecsBuilder) Name(d model.Decs) *FuncDecsBuilder {
	f.T.Name = d
	return f
}
func (f *FuncDecsBuilder) Func(d model.Decs) *FuncDecsBuilder {
	f.T.Func = d
	return f
}
func (f *FuncDecsBuilder) Recv(d model.Decs) *FuncDecsBuilder {
	f.T.Recv = d
	return f
}
func (f *FuncDecsBuilder) End(d model.Decs) *FuncDecsBuilder {
	f.T.End = d
	return f
}
func (b *GenBuilder) Build() *model.Gen          { return b.T }
func (b *GenBuilder) AsDecl() model.Decl         { return b.T }
func (b *FuncBuilder) Build() *model.Func        { return b.T }
func (b *FuncBuilder) AsDecl() model.Decl        { return b.T }
func (f *FuncDecsBuilder) Build() model.FuncDecs { return f.T }
