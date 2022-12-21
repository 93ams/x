package builder

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	"go/token"
)

type (
	StmtBuilder       interface{ AsStmt() model2.Stmt }
	AssignBuilder     x.Builder[*model2.Assign]
	BlockBuilder      x.Builder[*model2.Block]
	ExprStmtBuilder   x.Builder[*model2.ExprStmt]
	ReturnBuilder     x.Builder[*model2.Return]
	ReturnDecsBuilder x.Builder[model2.ReturnDecs]
)

func MapStmts(builders []StmtBuilder) []model2.Stmt {
	return lo.Map(builders, func(item StmtBuilder, _ int) model2.Stmt { return item.AsStmt() })
}
func Assign(t token.Token, lhs, rhs []ExprBuilder) *AssignBuilder {
	return &AssignBuilder{T: &model2.Assign{Lhs: MapExprs(lhs), Tok: t, Rhs: MapExprs(rhs)}}
}
func Block(stmts ...StmtBuilder) *BlockBuilder {
	return &BlockBuilder{T: model2.NewBlock(MapStmts(stmts))}
}
func NewExpr(expr model2.Expr) *model2.ExprStmt { return &model2.ExprStmt{X: expr} }
func Expr(expr ExprBuilder) *ExprStmtBuilder {
	return &ExprStmtBuilder{T: NewExpr(expr.AsExpr())}
}
func NewReturn(stmts []model2.Expr) *model2.Return { return &model2.Return{Results: stmts} }
func Return(exprs ...ExprBuilder) *ReturnBuilder {
	return &ReturnBuilder{T: NewReturn(MapExprs(exprs))}
}
func ReturnDecs() *ReturnDecsBuilder { return &ReturnDecsBuilder{T: model2.ReturnDecs{}} }

func (b *BlockBuilder) Decs(decs model2.BlockDecs) *BlockBuilder {
	b.T.Decs = decs
	return b
}
func (b *ExprStmtBuilder) Decs(decs model2.ExprStmtDecorations) *ExprStmtBuilder {
	b.T.Decs = decs
	return b
}
func (b *ReturnBuilder) Decs(decs *ReturnDecsBuilder) *ReturnBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *ReturnDecsBuilder) Before(d model2.SpaceType) *ReturnDecsBuilder {
	b.T.Before = d
	return b
}
func (b *ReturnDecsBuilder) After(d model2.SpaceType) *ReturnDecsBuilder {
	b.T.After = d
	return b
}
func (b *ReturnDecsBuilder) Start(d model2.Decs) *ReturnDecsBuilder { b.T.Start = d; return b }
func (b *ReturnDecsBuilder) End(d model2.Decs) *ReturnDecsBuilder   { b.T.End = d; return b }

func (b *AssignBuilder) Build() *model2.Assign        { return b.T }
func (b *AssignBuilder) AsStmt() model2.Stmt          { return b.T }
func (b *BlockBuilder) Build() *model2.Block          { return b.T }
func (b *BlockBuilder) AsStmt() model2.Stmt           { return b.T }
func (b *ExprStmtBuilder) Build() *model2.ExprStmt    { return b.T }
func (b *ExprStmtBuilder) AsStmt() model2.Stmt        { return b.T }
func (b *ReturnBuilder) Build() *model2.Return        { return b.T }
func (b *ReturnBuilder) AsStmt() model2.Stmt          { return b.T }
func (b *ReturnDecsBuilder) Build() model2.ReturnDecs { return b.T }
