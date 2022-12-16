package builder

import (
	"github.com/samber/lo"
	. "github.com/tilau2328/cql/src/go/package/x"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type (
	FieldListBuilder interface {
		Decs(model.FieldListDecorations) FieldListBuilder
		IBuilder[*model.FieldList]
	}
	FieldBuilder interface {
		Decs(model.FieldDecorations) FieldBuilder
		Tag(IBuilder[*model.Lit]) FieldBuilder
		IBuilder[*model.Field]
	}
	FileBuilder interface {
		Decs(model.FileDecorations) FileBuilder
		Decls(...DeclBuilder) FileBuilder
		Scope(*model.Scope) FileBuilder
		IBuilder[*model.File]
	}
	PackageBuilder interface {
		Imports(map[string]*model.Object) PackageBuilder
		Decs(model.PackageDecorations) PackageBuilder
		Files(map[string]FileBuilder) PackageBuilder
		Scope(*model.Scope) PackageBuilder
		IBuilder[*model.Package]
	}
	// SpecBuilder Specification Builders
	SpecBuilder  interface{ AsSpec() model.Spec }
	ValueBuilder interface {
		Decs(model.ValueDecorations) ValueBuilder
		Values(...ExprBuilder) ValueBuilder
		IBuilder[*model.Value]
		SpecBuilder
	}
	ImportBuilder interface {
		Decs(model.ImportDecorations) ImportBuilder
		Path(LitBuilder) ImportBuilder
		IBuilder[*model.Import]
		SpecBuilder
	}
	TypeBuilder interface {
		Decs(model.TypeDecorations) TypeBuilder
		Params(FieldListBuilder) TypeBuilder
		Name(IdentBuilder) TypeBuilder
		Assign(bool) TypeBuilder
		IBuilder[*model.Type]
		SpecBuilder
	}
	DeclBuilder interface{ AsDecl() model.Decl }
	FuncBuilder interface {
		Decs(FuncDecsBuilder) FuncBuilder
		Recv(FieldListBuilder) FuncBuilder
		Type(FuncTypeBuilder) FuncBuilder
		Body(BlockBuilder) FuncBuilder
		IBuilder[*model.Func]
		DeclBuilder
	}
	FuncDecsBuilder interface {
		IBuilder[model.FuncDecs]
	}
	GenBuilder interface {
		Decs(model.GenDecorations) GenBuilder
		IBuilder[*model.Gen]
		DeclBuilder
	}
	// StmtBuilder Statement Builders
	StmtBuilder interface{ AsStmt() model.Stmt }
	IfBuilder   interface {
		IBuilder[*model.If]
		StmtBuilder
	}
	GoBuilder interface {
		IBuilder[*model.Go]
		StmtBuilder
	}
	ForBuilder interface {
		IBuilder[*model.For]
		StmtBuilder
	}
	SendBuilder interface {
		IBuilder[*model.Send]
		StmtBuilder
	}
	DeferBuilder interface {
		IBuilder[*model.Defer]
		StmtBuilder
	}
	EmptyBuilder interface {
		IBuilder[*model.Empty]
		StmtBuilder
	}
	RangeBuilder interface {
		IBuilder[*model.Range]
		StmtBuilder
	}
	BlockBuilder interface {
		IBuilder[*model.Block]
		StmtBuilder
	}
	BranchBuilder interface {
		IBuilder[*model.Branch]
		StmtBuilder
	}
	ReturnBuilder interface {
		Decs(ReturnDecsBuilder) ReturnBuilder
		IBuilder[*model.Return]
		StmtBuilder
	}
	ReturnDecsBuilder interface {
		Start(d model.Decs) ReturnDecsBuilder
		End(d model.Decs) ReturnDecsBuilder
		Before(d model.SpaceType) ReturnDecsBuilder
		After(d model.SpaceType) ReturnDecsBuilder
		IBuilder[model.ReturnDecs]
	}
	AssignBuilder interface {
		IBuilder[*model.Assign]
		StmtBuilder
	}
	SelectBuilder interface {
		IBuilder[*model.Select]
		StmtBuilder
	}
	SwitchBuilder interface {
		IBuilder[*model.Switch]
		StmtBuilder
	}
	IncDecBuilder interface {
		IBuilder[*model.IncDec]
		StmtBuilder
	}
	LabeledBuilder interface {
		IBuilder[*model.Labeled]
		StmtBuilder
	}
	ExprStmtBuilder interface {
		Decs(model.ExprStmtDecorations) ExprStmtBuilder
		IBuilder[*model.ExprStmt]
		StmtBuilder
	}
	DeclStmtBuilder interface {
		IBuilder[*model.DeclStmt]
		StmtBuilder
	}
	TypeSwitchBuilder interface {
		IBuilder[*model.TypeSwitch]
		StmtBuilder
	}
	CaseClauseBuilder interface {
		IBuilder[*model.CaseClause]
		StmtBuilder
	}
	CommClauseBuilder interface {
		IBuilder[*model.CommClause]
		StmtBuilder
	}
	// ExprBuilder Expression Builders
	ExprBuilder         interface{ AsExpr() model.Expr }
	CompositeLitBuilder interface {
		Decs(CompositeLitDecsBuilder) CompositeLitBuilder
		Elts(...ExprBuilder) CompositeLitBuilder
		Incomplete(bool) CompositeLitBuilder
		IBuilder[*model.CompositeLit]
		ExprBuilder
	}
	CompositeLitDecsBuilder interface {
		IBuilder[model.CompositeLitDecs]
	}
	LitBuilder interface {
		IBuilder[*model.Lit]
		ExprBuilder
	}
	LitDecsBuilder interface {
		IBuilder[model.LitDecs]
	}
	SelectorBuilder interface {
		Decs(model.SelectorDecs) SelectorBuilder
		IBuilder[*model.Selector]
		ExprBuilder
	}
	SelectorDecsBuilder interface {
		IBuilder[model.SelectorDecs]
	}
	KeyValueBuilder interface {
		Decs(KeyValueDecsBuilder) KeyValueBuilder
		IBuilder[*model.KeyValue]
		ExprBuilder
	}
	KeyValueDecsBuilder interface {
		Start(model.Decs) KeyValueDecsBuilder
		Before(model.SpaceType) KeyValueDecsBuilder
		After(model.SpaceType) KeyValueDecsBuilder
		End(model.Decs) KeyValueDecsBuilder
		IBuilder[model.KeyValueDecs]
	}
	FuncTypeBuilder interface {
		Decs(model.FuncTypeDecorations) FuncTypeBuilder
		TypeParams(FieldListBuilder) FuncTypeBuilder
		Results(FieldListBuilder) FuncTypeBuilder
		Params(FieldListBuilder) FuncTypeBuilder
		Func(bool) FuncTypeBuilder
		IBuilder[*model.FuncType]
		ExprBuilder
	}
	CallBuilder interface {
		Decs(model.CallDecs) CallBuilder
		Ellipsis(bool) CallBuilder
		IBuilder[*model.Call]
		ExprBuilder
	}
	IdentBuilder interface {
		Decs(model.IdentDecorations) IdentBuilder
		Obj(*model.Object) IdentBuilder
		Path(string) IdentBuilder
		IBuilder[*model.Ident]
		ExprBuilder
	}
	StructBuilder interface {
		IBuilder[*model.Struct]
		ExprBuilder
	}
	EllipsisBuilder interface {
		IBuilder[*model.Ellipsis]
		ExprBuilder
	}
	InterfaceBuilder interface {
		Decs(model.InterfaceDecorations) InterfaceBuilder
		Incomplete(bool) InterfaceBuilder
		IBuilder[*model.Interface]
		ExprBuilder
	}
)

func MapStmts(builders []StmtBuilder) []model.Stmt {
	return lo.Map(builders, func(item StmtBuilder, _ int) model.Stmt { return item.AsStmt() })
}
func MapDecls(builders []DeclBuilder) []model.Decl {
	return lo.Map(builders, func(item DeclBuilder, _ int) model.Decl { return item.AsDecl() })
}
func MapExprs(builders []ExprBuilder) []model.Expr {
	return lo.Map(builders, func(item ExprBuilder, _ int) model.Expr { return item.AsExpr() })
}
func MapSpecs(builders []SpecBuilder) []model.Spec {
	return lo.Map(builders, func(item SpecBuilder, _ int) model.Spec { return item.AsSpec() })
}
