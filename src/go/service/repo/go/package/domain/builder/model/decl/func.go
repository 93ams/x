package decl

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type (
	funcBuilder     Builder[*model.Func]
	funcDecsBuilder Builder[model.FuncDecs]
)

func NewFunc(name *model.Ident) *model.Func {
	return &model.Func{Name: name, Type: &model.FuncType{}, Body: &model.Block{}}
}
func Func(name IBuilder[*model.Ident]) FuncBuilder {
	return &funcBuilder{T: NewFunc(name.Build())}
}
func (b *funcBuilder) Decs(decs FuncDecsBuilder) FuncBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *funcBuilder) Recv(fields FieldListBuilder) FuncBuilder {
	b.T.Recv = fields.Build()
	return b
}
func (b *funcBuilder) Type(t FuncTypeBuilder) FuncBuilder {
	b.T.Type = t.Build()
	return b
}
func (b *funcBuilder) Body(body BlockBuilder) FuncBuilder {
	b.T.Body = body.Build()
	return b
}
func (b *funcBuilder) Build() *model.Func { return b.T }
func (b *funcBuilder) AsDecl() model.Decl { return b.T }

func FuncDecs() FuncDecsBuilder { return &funcDecsBuilder{} }
func (f *funcDecsBuilder) TypeParams(d model.Decorations) FuncDecsBuilder {
	f.T.TypeParams = d
	return f
}
func (f *funcDecsBuilder) Results(d model.Decorations) FuncDecsBuilder {
	f.T.Results = d
	return f
}
func (f *funcDecsBuilder) Params(d model.Decorations) FuncDecsBuilder {
	f.T.Params = d
	return f
}
func (f *funcDecsBuilder) Start(d model.Decorations) FuncDecsBuilder {
	f.T.Start = d
	return f
}
func (f *funcDecsBuilder) Name(d model.Decorations) FuncDecsBuilder {
	f.T.Name = d
	return f
}
func (f *funcDecsBuilder) Func(d model.Decorations) FuncDecsBuilder {
	f.T.Func = d
	return f
}
func (f *funcDecsBuilder) Recv(d model.Decorations) FuncDecsBuilder {
	f.T.Recv = d
	return f
}
func (f *funcDecsBuilder) End(d model.Decorations) FuncDecsBuilder {
	f.T.End = d
	return f
}
func (f *funcDecsBuilder) Build() model.FuncDecs { return f.T }
