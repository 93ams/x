package stmt

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type blockBuilder Builder[*model.Block]

func NewBlock(stmts []model.Stmt) *model.Block { return &model.Block{List: stmts} }
func Block(stmts ...StmtBuilder) BlockBuilder {
	return &blockBuilder{T: NewBlock(MapStmts(stmts))}
}
func (b *blockBuilder) Decs(decs model.BlockDecorations) BlockBuilder {
	b.T.Decs = decs
	return b
}
func (b *blockBuilder) Build() *model.Block { return b.T }
func (b *blockBuilder) AsStmt() model.Stmt  { return b.T }
