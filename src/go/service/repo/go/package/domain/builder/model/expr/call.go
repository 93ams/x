package expr

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type callBuilder Builder[*model.Call]

// Call callBuilder
func Call(fun ExprBuilder, args ...ExprBuilder) CallBuilder {
	return &callBuilder{T: &model.Call{Fun: fun.AsExpr(), Args: MapExprs(args)}}
}
func (b *callBuilder) Decs(decs model.CallDecorations) CallBuilder {
	b.T.Decs = decs
	return b
}
func (b *callBuilder) Ellipsis(ellipsis bool) CallBuilder {
	b.T.Ellipsis = ellipsis
	return b
}
func (b *callBuilder) Build() *model.Call { return b.T }
func (b *callBuilder) AsExpr() model.Expr { return b.T }
