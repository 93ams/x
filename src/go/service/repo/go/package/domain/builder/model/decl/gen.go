package decl

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
	"go/token"
)

type genBuilder Builder[*model.Gen]

func NewGen(t token.Token, specs []model.Spec) *model.Gen { return &model.Gen{Tok: t, Specs: specs} }
func Gen(t token.Token, specs ...SpecBuilder) GenBuilder {
	return &genBuilder{T: NewGen(t, MapSpecs(specs))}
}
func (b *genBuilder) Decs(decs model.GenDecorations) GenBuilder {
	b.T.Decs = decs
	return b
}
func (b *genBuilder) Build() *model.Gen  { return b.T }
func (b *genBuilder) AsDecl() model.Decl { return b.T }
