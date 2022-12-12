package model

type ArrayTypeDecorations struct {
	NodeDecs
	Lbrack Decorations
	Len    Decorations
}
type AssignStmtDecorations struct {
	NodeDecs
	Tok Decorations
}
type BadDeclDecorations struct {
	NodeDecs
}
type BadExprDecorations struct {
	NodeDecs
}
type BadStmtDecorations struct {
	NodeDecs
}
type BasicLitDecorations struct {
	NodeDecs
}
type BinaryExprDecorations struct {
	NodeDecs
	X  Decorations
	Op Decorations
}
type BlockStmtDecorations struct {
	NodeDecs
	Lbrace Decorations
}
type BranchStmtDecorations struct {
	NodeDecs
	Tok Decorations
}
type CallExprDecorations struct {
	NodeDecs
	Fun      Decorations
	Lparen   Decorations
	Ellipsis Decorations
}
type CaseClauseDecorations struct {
	NodeDecs
	Case  Decorations
	Colon Decorations
}
type ChanTypeDecorations struct {
	NodeDecs
	Begin Decorations
	Arrow Decorations
}
type CommClauseDecorations struct {
	NodeDecs
	Case  Decorations
	Comm  Decorations
	Colon Decorations
}
type CompositeLitDecorations struct {
	NodeDecs
	Type   Decorations
	Lbrace Decorations
}
type DeclStmtDecorations struct {
	NodeDecs
}
type DeferStmtDecorations struct {
	NodeDecs
	Defer Decorations
}
type EllipsisDecorations struct {
	NodeDecs
	Ellipsis Decorations
}
type EmptyStmtDecorations struct {
	NodeDecs
}
type ExprStmtDecorations struct {
	NodeDecs
}
type FieldDecorations struct {
	NodeDecs
	Type Decorations
}
type FieldListDecorations struct {
	NodeDecs
	Opening Decorations
}
type FileDecorations struct {
	NodeDecs
	Package Decorations
	Name    Decorations
}
type ForStmtDecorations struct {
	NodeDecs
	For  Decorations
	Init Decorations
	Cond Decorations
	Post Decorations
}
type FuncDeclDecorations struct {
	NodeDecs
	Func       Decorations
	Recv       Decorations
	Name       Decorations
	TypeParams Decorations
	Params     Decorations
	Results    Decorations
}
type FuncLitDecorations struct {
	NodeDecs
	Type Decorations
}
type FuncTypeDecorations struct {
	NodeDecs
	Func       Decorations
	TypeParams Decorations
	Params     Decorations
}
type GenDeclDecorations struct {
	NodeDecs
	Tok    Decorations
	Lparen Decorations
}
type GoStmtDecorations struct {
	NodeDecs
	Go Decorations
}
type IdentDecorations struct {
	NodeDecs
	X Decorations
}
type IfStmtDecorations struct {
	NodeDecs
	If   Decorations
	Init Decorations
	Cond Decorations
	Else Decorations
}
type ImportSpecDecorations struct {
	NodeDecs
	Name Decorations
}
type IncDecStmtDecorations struct {
	NodeDecs
	X Decorations
}
type IndexExprDecorations struct {
	NodeDecs
	X      Decorations
	Lbrack Decorations
	Index  Decorations
}
type IndexListExprDecorations struct {
	NodeDecs
	X       Decorations
	Lbrack  Decorations
	Indices Decorations
}
type InterfaceTypeDecorations struct {
	NodeDecs
	Interface Decorations
}
type KeyValueExprDecorations struct {
	NodeDecs
	Key   Decorations
	Colon Decorations
}
type LabeledStmtDecorations struct {
	NodeDecs
	Label Decorations
	Colon Decorations
}
type MapTypeDecorations struct {
	NodeDecs
	Map Decorations
	Key Decorations
}
type PackageDecorations struct {
	NodeDecs
}
type ParenExprDecorations struct {
	NodeDecs
	Lparen Decorations
	X      Decorations
}
type RangeStmtDecorations struct {
	NodeDecs
	For   Decorations
	Key   Decorations
	Value Decorations
	Range Decorations
	X     Decorations
}
type ReturnStmtDecorations struct {
	NodeDecs
	Return Decorations
}
type SelectStmtDecorations struct {
	NodeDecs
	Select Decorations
}

type SelectorExprDecorations struct {
	NodeDecs
	X Decorations
}

type SendStmtDecorations struct {
	NodeDecs
	Chan  Decorations
	Arrow Decorations
}

type SliceExprDecorations struct {
	NodeDecs
	X      Decorations
	Lbrack Decorations
	Low    Decorations
	High   Decorations
	Max    Decorations
}
type StarExprDecorations struct {
	NodeDecs
	Star Decorations
}
type StructTypeDecorations struct {
	NodeDecs
	Struct Decorations
}

type SwitchStmtDecorations struct {
	NodeDecs
	Switch Decorations
	Init   Decorations
	Tag    Decorations
}

type TypeAssertExprDecorations struct {
	NodeDecs
	X      Decorations
	Lparen Decorations
	Type   Decorations
}

type TypeSpecDecorations struct {
	NodeDecs
	Name       Decorations
	TypeParams Decorations
}

type TypeSwitchStmtDecorations struct {
	NodeDecs
	Switch Decorations
	Init   Decorations
	Assign Decorations
}

type UnaryExprDecorations struct {
	NodeDecs
	Op Decorations
}

type ValueSpecDecorations struct {
	NodeDecs
	Assign Decorations
}
