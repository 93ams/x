package stmt

import (
	"github.com/tilau2328/x/src/go/services/gen/go/package/domain/model"
)

type exprStmtBuilder Builder[*model.ExprStmt]

func NewExpr(expr model.Expr) *model.ExprStmt { return &model.ExprStmt{X: expr} }
func Expr(expr ExprBuilder) ExprStmtBuilder   { return &exprStmtBuilder{T: NewExpr(expr.AsExpr())} }
func (b *exprStmtBuilder) Decs(decs model.ExprStmtDecorations) ExprStmtBuilder {
	b.T.Decs = decs
	return b
}
func (b *exprStmtBuilder) Build() *model.ExprStmt { return b.T }
func (b *exprStmtBuilder) AsStmt() model.Stmt     { return b.T }
