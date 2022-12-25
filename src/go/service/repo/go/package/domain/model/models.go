package model

import (
	"go/token"
)

type (
	ChanDir int
	Node    interface{ Decors() *NodeDecs }
	Expr    interface {
		Node
		AsExpr() Expr
	}
	Stmt interface {
		Node
		AsStmt() Stmt
	}
	Decl interface {
		Node
		AsDecl() Decl
	}
	Spec interface {
		Node
		AsSpec() Spec
	}
	Field struct {
		Names []*Ident
		Type  Expr
		Tag   *Lit
		Decs  FieldDecs
	}
	FieldList struct {
		Opening bool
		List    []*Field
		Closing bool
		Decs    FieldListDecs
	}
	BadExpr struct {
		Length int
		Decs   BadExprDecs
	}
	Ident struct {
		Name string
		Obj  *Object
		Path string
		Decs IdentDecs
	}
	Ellipsis struct {
		Elt  Expr
		Decs EllipsisDecs
	}
	Lit struct {
		Kind  token.Token
		Value string
		Decs  LitDecs
	}
	FuncLit struct {
		Type *FuncType
		Body *Block
		Decs FuncLitDecs
	}
	CompositeLit struct {
		Type       Expr
		Elts       []Expr
		Incomplete bool
		Decs       CompositeLitDecs
	}
	Paren struct {
		X    Expr
		Decs ParenExprDecs
	}
	Selector struct {
		X    Expr
		Sel  *Ident
		Decs SelectorDecs
	}
	Index struct {
		X     Expr
		Index Expr
		Decs  IndexDecs
	}
	IndexList struct {
		X       Expr
		Indices []Expr
		Decs    IndexListDecs
	}
	Slice struct {
		X      Expr
		Low    Expr
		High   Expr
		Max    Expr
		Slice3 bool
		Decs   SliceExprDecs
	}
	TypeAssert struct {
		X    Expr
		Type Expr
		Decs TypeAssertDecs
	}
	Call struct {
		Fun      Expr
		Args     []Expr
		Ellipsis bool
		Decs     CallDecs
	}
	Star struct {
		X    Expr
		Decs StarDecs
	}
	Unary struct {
		Op   token.Token
		X    Expr
		Decs UnaryDecs
	}
	Binary struct {
		X    Expr
		Op   token.Token
		Y    Expr
		Decs BinaryDecs
	}
	KeyValue struct {
		Key   Expr
		Value Expr
		Decs  KeyValueDecs
	}

	Array struct {
		Len  Expr
		Elt  Expr
		Decs ArrayTypeDecs
	}
	Struct struct {
		Fields     *FieldList
		Incomplete bool
		Decs       StructDecs
	}
	FuncType struct {
		Func       bool
		TypeParams *FieldList
		Params     *FieldList
		Results    *FieldList
		Decs       FuncTypeDecs
	}
	Interface struct {
		Methods    *FieldList
		Incomplete bool
		Decs       InterfaceDecs
	}
	MapType struct {
		Key   Expr
		Value Expr
		Decs  MapTypeDecs
	}
	Chan struct {
		Dir   ChanDir
		Value Expr
		Decs  ChanTypeDecs
	}
	// Spec spec related models
	BadStmt struct {
		Length int
		Decs   BadStmtDecs
	}
	DeclStmt struct {
		Decl Decl
		Decs DeclStmtDecs
	}
	Empty struct {
		Implicit bool
		Decs     EmptyStmtDecs
	}
	Labeled struct {
		Label *Ident
		Stmt  Stmt
		Decs  LabeledStmtDecs
	}
	ExprStmt struct {
		X    Expr
		Decs ExprStmtDecs
	}
	Send struct {
		Chan  Expr
		Value Expr
		Decs  SendStmtDecs
	}
	IncDec struct {
		X    Expr
		Tok  token.Token
		Decs IncDecStmtDecs
	}
	Assign struct {
		Lhs  []Expr
		Tok  token.Token
		Rhs  []Expr
		Decs AssignDecs
	}
	Go struct {
		Call *Call
		Decs GoStmtDecs
	}
	Defer struct {
		Call *Call
		Decs DeferStmtDecs
	}
	Return struct {
		Results []Expr
		Decs    ReturnDecs
	}
	Branch struct {
		Tok   token.Token
		Label *Ident
		Decs  BranchDecs
	}
	Block struct {
		List           []Stmt
		RbraceHasNoPos bool
		Decs           BlockDecs
	}
	If struct {
		Init Stmt
		Cond Expr
		Body *Block
		Else Stmt
		Decs IfStmtDecs
	}
	CaseClause struct {
		List []Expr
		Body []Stmt
		Decs CaseClauseDecs
	}
	Switch struct {
		Init Stmt
		Tag  Expr
		Body *Block
		Decs SwitchStmtDecs
	}
	TypeSwitch struct {
		Init   Stmt
		Assign Stmt
		Body   *Block
		Decs   TypeSwitchStmtDecs
	}
	CommClause struct {
		Comm Stmt
		Body []Stmt
		Decs CommClauseDecs
	}
	Select struct {
		Body *Block
		Decs SelectStmtDecs
	}
	For struct {
		Init Stmt
		Cond Expr
		Post Stmt
		Body *Block
		Decs ForStmtDecs
	}
	Range struct {
		Key, Value Expr
		Tok        token.Token
		X          Expr
		Body       *Block
		Decs       RangeStmtDecs
	}
	Import struct {
		Name *Ident
		Path *Lit
		Decs ImportDecs
	}
	Value struct {
		Names  []*Ident
		Type   Expr
		Values []Expr
		Decs   ValueDecs
	}
	Type struct {
		Name   *Ident
		Params *FieldList
		Assign bool
		Type   Expr
		Decs   TypeDecs
	}
	BadDecl struct {
		Length int
		Decs   BadDeclDecs
	}
	Gen struct {
		Tok    token.Token
		Lparen bool
		Specs  []Spec
		Rparen bool
		Decs   GenDecs
	}
	Func struct {
		Recv *FieldList
		Name *Ident
		Type *FuncType
		Body *Block
		Decs FuncDecs
	}
	// File file related models
	File struct {
		Name       *Ident
		Decls      []Decl
		Scope      *Scope
		Imports    []*Import
		Unresolved []*Ident
		Decs       FileDecs
	}
	Package struct {
		Name    string
		Scope   *Scope
		Imports map[string]*Object
		Files   map[string]*File
		Decs    PackageDecs
	}
	Module struct {
		Name     string
		Packages map[string]*Package
	}
)

const (
	SendChan ChanDir = 1 << iota
	RecvChan
)

func IsExported(name string) bool { return token.IsExported(name) }
func (n *Ident) IsExported() bool { return token.IsExported(n.Name) }
func (n *Ident) String() string {
	if n != nil {
		if n.Path != "" {
			return n.Path + "." + n.Name
		}
		return n.Name
	}
	return "<nil>"
}
func (n *FieldList) NumFields() (ret int) {
	if n != nil {
		for _, g := range n.List {
			m := len(g.Names)
			if m == 0 {
				m = 1
			}
			ret += m
		}
	}
	return
}
