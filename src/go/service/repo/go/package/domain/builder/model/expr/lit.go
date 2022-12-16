package expr

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
	"go/token"
)

type litBuilder Builder[*model.Lit]

// Lit LitBuilder
func Lit(kind token.Token, value string) LitBuilder {
	return &litBuilder{T: &model.Lit{Value: value, Kind: kind}}
}
func (b *litBuilder) Decs(decs model.LitDecs) LitBuilder {
	b.T.Decs = decs
	return b
}
func (b *litBuilder) Build() *model.Lit  { return b.T }
func (b *litBuilder) AsExpr() model.Expr { return b.T }
