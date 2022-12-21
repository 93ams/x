package builder

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper/model"
	"go/token"
)

type (
	StmtBuilder       interface{ AsStmt() model.Stmt }
	AssignBuilder     x.Builder[*model.Assign]
	BlockBuilder      x.Builder[*model.Block]
	ExprStmtBuilder   x.Builder[*model.ExprStmt]
	ReturnBuilder     x.Builder[*model.Return]
	ReturnDecsBuilder x.Builder[model.ReturnDecs]
)

func MapStmts(builders []StmtBuilder) []model.Stmt {
	return lo.Map(builders, func(item StmtBuilder, _ int) model.Stmt { return item.AsStmt() })
}
func Assign(t token.Token, lhs, rhs []ExprBuilder) *AssignBuilder {
	return &AssignBuilder{T: &model.Assign{Lhs: MapExprs(lhs), Tok: t, Rhs: MapExprs(rhs)}}
}
func Block(stmts ...StmtBuilder) *BlockBuilder {
	return &BlockBuilder{T: model.NewBlock(MapStmts(stmts))}
}
func NewExpr(expr model.Expr) *model.ExprStmt { return &model.ExprStmt{X: expr} }
func Expr(expr ExprBuilder) *ExprStmtBuilder {
	return &ExprStmtBuilder{T: NewExpr(expr.AsExpr())}
}
func NewReturn(stmts []model.Expr) *model.Return { return &model.Return{Results: stmts} }
func Return(exprs ...ExprBuilder) *ReturnBuilder {
	return &ReturnBuilder{T: NewReturn(MapExprs(exprs))}
}
func ReturnDecs() *ReturnDecsBuilder { return &ReturnDecsBuilder{T: model.ReturnDecs{}} }

func (b *BlockBuilder) Decs(decs model.BlockDecs) *BlockBuilder {
	b.T.Decs = decs
	return b
}
func (b *ExprStmtBuilder) Decs(decs model.ExprStmtDecorations) *ExprStmtBuilder {
	b.T.Decs = decs
	return b
}
func (b *ReturnBuilder) Decs(decs *ReturnDecsBuilder) *ReturnBuilder {
	b.T.Decs = decs.Build()
	return b
}
func (b *ReturnDecsBuilder) Before(d model.SpaceType) *ReturnDecsBuilder {
	b.T.Before = d
	return b
}
func (b *ReturnDecsBuilder) After(d model.SpaceType) *ReturnDecsBuilder {
	b.T.After = d
	return b
}
func (b *ReturnDecsBuilder) Start(d model.Decs) *ReturnDecsBuilder { b.T.Start = d; return b }
func (b *ReturnDecsBuilder) End(d model.Decs) *ReturnDecsBuilder   { b.T.End = d; return b }

func (b *AssignBuilder) Build() *model.Assign        { return b.T }
func (b *AssignBuilder) AsStmt() model.Stmt          { return b.T }
func (b *BlockBuilder) Build() *model.Block          { return b.T }
func (b *BlockBuilder) AsStmt() model.Stmt           { return b.T }
func (b *ExprStmtBuilder) Build() *model.ExprStmt    { return b.T }
func (b *ExprStmtBuilder) AsStmt() model.Stmt        { return b.T }
func (b *ReturnBuilder) Build() *model.Return        { return b.T }
func (b *ReturnBuilder) AsStmt() model.Stmt          { return b.T }
func (b *ReturnDecsBuilder) Build() model.ReturnDecs { return b.T }
