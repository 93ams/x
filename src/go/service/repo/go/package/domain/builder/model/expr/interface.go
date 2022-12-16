package expr

import (
	. "github.com/tilau2328/cql/src/go/package/x"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type interfaceBuilder Builder[*model.Interface]

// Interface InterfaceBuilder
func Interface(methods FieldListBuilder) InterfaceBuilder {
	return &interfaceBuilder{T: &model.Interface{Methods: methods.Build()}}
}
func (b *interfaceBuilder) Decs(decs model.InterfaceDecorations) InterfaceBuilder {
	b.T.Decs = decs
	return b
}
func (b *interfaceBuilder) Incomplete(incomplete bool) InterfaceBuilder {
	b.T.Incomplete = incomplete
	return b
}
func (b *interfaceBuilder) Build() *model.Interface { return b.T }

func (b *interfaceBuilder) AsExpr() model.Expr { return b.T }
