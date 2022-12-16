package stmt

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type (
	returnBuilder     Builder[*model.Return]
	returnDecsBuilder Builder[model.ReturnDecs]
)

func NewReturn(stmts []model.Expr) *model.Return { return &model.Return{Results: stmts} }
func Return(exprs ...ExprBuilder) ReturnBuilder {
	return &returnBuilder{T: NewReturn(MapExprs(exprs))}
}
func (b *returnBuilder) Decs(decs ReturnDecsBuilder) ReturnBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *returnBuilder) Build() *model.Return                           { return b.T }
func (b *returnBuilder) AsStmt() model.Stmt                             { return b.T }
func ReturnDecs() ReturnDecsBuilder                                     { return &returnDecsBuilder{T: model.ReturnDecs{}} }
func (b *returnDecsBuilder) Before(d model.SpaceType) ReturnDecsBuilder { b.T.Before = d; return b }
func (b *returnDecsBuilder) After(d model.SpaceType) ReturnDecsBuilder  { b.T.After = d; return b }
func (b *returnDecsBuilder) Start(d model.Decs) ReturnDecsBuilder       { b.T.Start = d; return b }
func (b *returnDecsBuilder) End(d model.Decs) ReturnDecsBuilder         { b.T.End = d; return b }
func (b *returnDecsBuilder) Build() model.ReturnDecs                    { return b.T }
