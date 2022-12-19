package model

import "go/token"

type (
	Node interface{ Decorations() *NodeDecs }
	Expr interface {
		Node
		exprNode()
	}
	Stmt interface {
		Node
		stmtNode()
	}
	Decl interface {
		Node
		declNode()
	}
)
type Field struct {
	Names []*Ident
	Type  Expr
	Tag   *Lit
	Decs  FieldDecorations
}
type FieldList struct {
	Opening bool
	List    []*Field
	Closing bool
	Decs    FieldListDecorations
}

func (l *FieldList) NumFields() int {
	n := 0
	if l != nil {
		for _, g := range l.List {
			m := len(g.Names)
			if m == 0 {
				m = 1
			}
			n += m
		}
	}
	return n
}

type (
	BadExpr struct {
		Length int
		Decs   BadExprDecorations
	}
	Ident struct {
		Name string
		Obj  *Object
		Path string
		Decs IdentDecorations
	}
	Ellipsis struct {
		Elt  Expr
		Decs EllipsisDecorations
	}
	Lit struct {
		Kind  token.Token
		Value string
		Decs  LitDecs
	}
	FuncLit struct {
		Type *FuncType
		Body *Block
		Decs FuncLitDecorations
	}
	CompositeLit struct {
		Type       Expr
		Elts       []Expr
		Incomplete bool
		Decs       CompositeLitDecs
	}
	Paren struct {
		X    Expr
		Decs ParenExprDecorations
	}
	Selector struct {
		X    Expr
		Sel  *Ident
		Decs SelectorDecs
	}
	Index struct {
		X     Expr
		Index Expr
		Decs  IndexExprDecorations
	}
	IndexList struct {
		X       Expr
		Indices []Expr
		Decs    IndexListExprDecorations
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
		Decs TypeAssertExprDecorations
	}
	Call struct {
		Fun      Expr
		Args     []Expr
		Ellipsis bool
		Decs     CallDecs
	}
	Star struct {
		X    Expr
		Decs StarExprDecorations
	}
	Unary struct {
		Op   token.Token
		X    Expr
		Decs UnaryExprDecorations
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
)

type ChanDir int

const (
	SEND ChanDir = 1 << iota
	RECV
)

type (
	Array struct {
		Len  Expr
		Elt  Expr
		Decs ArrayTypeDecorations
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
		Decs       FuncTypeDecorations
	}
	Interface struct {
		Methods    *FieldList
		Incomplete bool
		Decs       InterfaceDecorations
	}
	MapType struct {
		Key   Expr
		Value Expr
		Decs  MapTypeDecorations
	}
	Chan struct {
		Dir   ChanDir
		Value Expr
		Decs  ChanTypeDecorations
	}
)

func (*BadExpr) exprNode()      {}
func (*Ident) exprNode()        {}
func (*Ellipsis) exprNode()     {}
func (*Lit) exprNode()          {}
func (*FuncLit) exprNode()      {}
func (*CompositeLit) exprNode() {}
func (*Paren) exprNode()        {}
func (*Selector) exprNode()     {}
func (*Index) exprNode()        {}
func (*IndexList) exprNode()    {}
func (*Slice) exprNode()        {}
func (*TypeAssert) exprNode()   {}
func (*Call) exprNode()         {}
func (*Star) exprNode()         {}
func (*Unary) exprNode()        {}
func (*Binary) exprNode()       {}
func (*KeyValue) exprNode()     {}

func (*Array) exprNode()     {}
func (*Struct) exprNode()    {}
func (*FuncType) exprNode()  {}
func (*Interface) exprNode() {}
func (*MapType) exprNode()   {}
func (*Chan) exprNode()      {}

func NewIdent(name string) *Ident { return &Ident{name, nil, "", IdentDecorations{}} }

func IsExported(name string) bool { return token.IsExported(name) }

func (id *Ident) IsExported() bool { return token.IsExported(id.Name) }

func (id *Ident) String() string {
	if id != nil {
		if id.Path != "" {
			return id.Path + "." + id.Name
		}
		return id.Name
	}
	return "<nil>"
}

type (
	BadStmt struct {
		Length int
		Decs   BadStmtDecorations
	}
	DeclStmt struct {
		Decl Decl
		Decs DeclStmtDecorations
	}
	Empty struct {
		Implicit bool
		Decs     EmptyStmtDecorations
	}
	Labeled struct {
		Label *Ident
		Stmt  Stmt
		Decs  LabeledStmtDecorations
	}
	ExprStmt struct {
		X    Expr
		Decs ExprStmtDecorations
	}
	Send struct {
		Chan  Expr
		Value Expr
		Decs  SendStmtDecs
	}
	IncDec struct {
		X    Expr
		Tok  token.Token
		Decs IncDecStmtDecorations
	}
	Assign struct {
		Lhs  []Expr
		Tok  token.Token
		Rhs  []Expr
		Decs AssignStmtDecorations
	}
	Go struct {
		Call *Call
		Decs GoStmtDecorations
	}
	Defer struct {
		Call *Call
		Decs DeferStmtDecorations
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
		Decs IfStmtDecorations
	}
	CaseClause struct {
		List []Expr
		Body []Stmt
		Decs CaseClauseDecorations
	}
	Switch struct {
		Init Stmt
		Tag  Expr
		Body *Block
		Decs SwitchStmtDecorations
	}
	TypeSwitch struct {
		Init   Stmt
		Assign Stmt
		Body   *Block
		Decs   TypeSwitchStmtDecorations
	}
	CommClause struct {
		Comm Stmt
		Body []Stmt
		Decs CommClauseDecorations
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
		Decs ForStmtDecorations
	}
	Range struct {
		Key, Value Expr
		Tok        token.Token
		X          Expr
		Body       *Block
		Decs       RangeStmtDecs
	}
)

func (*BadStmt) stmtNode()    {}
func (*DeclStmt) stmtNode()   {}
func (*Empty) stmtNode()      {}
func (*Labeled) stmtNode()    {}
func (*ExprStmt) stmtNode()   {}
func (*Send) stmtNode()       {}
func (*IncDec) stmtNode()     {}
func (*Assign) stmtNode()     {}
func (*Go) stmtNode()         {}
func (*Defer) stmtNode()      {}
func (*Return) stmtNode()     {}
func (*Branch) stmtNode()     {}
func (*Block) stmtNode()      {}
func (*If) stmtNode()         {}
func (*CaseClause) stmtNode() {}
func (*Switch) stmtNode()     {}
func (*TypeSwitch) stmtNode() {}
func (*CommClause) stmtNode() {}
func (*Select) stmtNode()     {}
func (*For) stmtNode()        {}
func (*Range) stmtNode()      {}

type (
	Spec interface {
		Node
		specNode()
	}
	Import struct {
		Name *Ident
		Path *Lit
		Decs ImportDecorations
	}
	Value struct {
		Names  []*Ident
		Type   Expr
		Values []Expr
		Decs   ValueDecorations
	}
	Type struct {
		Name   *Ident
		Params *FieldList
		Assign bool
		Type   Expr
		Decs   TypeDecorations
	}
)

func (*Import) specNode() {}
func (*Value) specNode()  {}
func (*Type) specNode()   {}

type (
	BadDecl struct {
		Length int
		Decs   BadDeclDecorations
	}
	Gen struct {
		Tok    token.Token
		Lparen bool
		Specs  []Spec
		Rparen bool
		Decs   GenDecorations
	}
	Func struct {
		Recv *FieldList
		Name *Ident
		Type *FuncType
		Body *Block
		Decs FuncDecs
	}
)

func (*BadDecl) declNode() {}
func (*Gen) declNode()     {}
func (*Func) declNode()    {}

type File struct {
	Name       *Ident
	Decls      []Decl
	Scope      *Scope
	Imports    []*Import
	Unresolved []*Ident
	Decs       FileDecorations
}

type Package struct {
	Name    string
	Scope   *Scope
	Imports map[string]*Object
	Files   map[string]*File
	Decs    PackageDecorations
}
