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
	Tag   *BasicLit
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
	BasicLit struct {
		Kind  token.Token
		Value string
		Decs  BasicLitDecorations
	}
	FuncLit struct {
		Type *FuncType
		Body *BlockStmt
		Decs FuncLitDecorations
	}
	CompositeLit struct {
		Type       Expr
		Elts       []Expr
		Incomplete bool
		Decs       CompositeLitDecorations
	}
	ParenExpr struct {
		X    Expr
		Decs ParenExprDecorations
	}
	SelectorExpr struct {
		X    Expr
		Sel  *Ident
		Decs SelectorExprDecorations
	}
	IndexExpr struct {
		X     Expr
		Index Expr
		Decs  IndexExprDecorations
	}
	IndexListExpr struct {
		X       Expr
		Indices []Expr
		Decs    IndexListExprDecorations
	}
	SliceExpr struct {
		X      Expr
		Low    Expr
		High   Expr
		Max    Expr
		Slice3 bool
		Decs   SliceExprDecorations
	}
	TypeAssertExpr struct {
		X    Expr
		Type Expr
		Decs TypeAssertExprDecorations
	}
	CallExpr struct {
		Fun      Expr
		Args     []Expr
		Ellipsis bool
		Decs     CallExprDecorations
	}
	StarExpr struct {
		X    Expr
		Decs StarExprDecorations
	}
	UnaryExpr struct {
		Op   token.Token
		X    Expr
		Decs UnaryExprDecorations
	}
	BinaryExpr struct {
		X    Expr
		Op   token.Token
		Y    Expr
		Decs BinaryExprDecorations
	}
	KeyValueExpr struct {
		Key   Expr
		Value Expr
		Decs  KeyValueExprDecorations
	}
)

type ChanDir int

const (
	SEND ChanDir = 1 << iota
	RECV
)

type (
	ArrayType struct {
		Len  Expr
		Elt  Expr
		Decs ArrayTypeDecorations
	}
	StructType struct {
		Fields     *FieldList
		Incomplete bool
		Decs       StructTypeDecorations
	}
	FuncType struct {
		Func       bool
		TypeParams *FieldList
		Params     *FieldList
		Results    *FieldList
		Decs       FuncTypeDecorations
	}
	InterfaceType struct {
		Methods    *FieldList
		Incomplete bool
		Decs       InterfaceTypeDecorations
	}
	MapType struct {
		Key   Expr
		Value Expr
		Decs  MapTypeDecorations
	}
	ChanType struct {
		Dir   ChanDir
		Value Expr
		Decs  ChanTypeDecorations
	}
)

func (*BadExpr) exprNode()        {}
func (*Ident) exprNode()          {}
func (*Ellipsis) exprNode()       {}
func (*BasicLit) exprNode()       {}
func (*FuncLit) exprNode()        {}
func (*CompositeLit) exprNode()   {}
func (*ParenExpr) exprNode()      {}
func (*SelectorExpr) exprNode()   {}
func (*IndexExpr) exprNode()      {}
func (*IndexListExpr) exprNode()  {}
func (*SliceExpr) exprNode()      {}
func (*TypeAssertExpr) exprNode() {}
func (*CallExpr) exprNode()       {}
func (*StarExpr) exprNode()       {}
func (*UnaryExpr) exprNode()      {}
func (*BinaryExpr) exprNode()     {}
func (*KeyValueExpr) exprNode()   {}

func (*ArrayType) exprNode()     {}
func (*StructType) exprNode()    {}
func (*FuncType) exprNode()      {}
func (*InterfaceType) exprNode() {}
func (*MapType) exprNode()       {}
func (*ChanType) exprNode()      {}

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
	EmptyStmt struct {
		Implicit bool
		Decs     EmptyStmtDecorations
	}
	LabeledStmt struct {
		Label *Ident
		Stmt  Stmt
		Decs  LabeledStmtDecorations
	}
	ExprStmt struct {
		X    Expr
		Decs ExprStmtDecorations
	}
	SendStmt struct {
		Chan  Expr
		Value Expr
		Decs  SendStmtDecorations
	}
	IncDecStmt struct {
		X    Expr
		Tok  token.Token
		Decs IncDecStmtDecorations
	}
	AssignStmt struct {
		Lhs  []Expr
		Tok  token.Token
		Rhs  []Expr
		Decs AssignStmtDecorations
	}
	GoStmt struct {
		Call *CallExpr
		Decs GoStmtDecorations
	}
	DeferStmt struct {
		Call *CallExpr
		Decs DeferStmtDecorations
	}
	ReturnStmt struct {
		Results []Expr
		Decs    ReturnStmtDecorations
	}
	BranchStmt struct {
		Tok   token.Token
		Label *Ident
		Decs  BranchStmtDecorations
	}
	BlockStmt struct {
		List           []Stmt
		RbraceHasNoPos bool
		Decs           BlockStmtDecorations
	}
	IfStmt struct {
		Init Stmt
		Cond Expr
		Body *BlockStmt
		Else Stmt
		Decs IfStmtDecorations
	}
	CaseClause struct {
		List []Expr
		Body []Stmt
		Decs CaseClauseDecorations
	}
	SwitchStmt struct {
		Init Stmt
		Tag  Expr
		Body *BlockStmt
		Decs SwitchStmtDecorations
	}
	TypeSwitchStmt struct {
		Init   Stmt
		Assign Stmt
		Body   *BlockStmt
		Decs   TypeSwitchStmtDecorations
	}
	CommClause struct {
		Comm Stmt
		Body []Stmt
		Decs CommClauseDecorations
	}
	SelectStmt struct {
		Body *BlockStmt
		Decs SelectStmtDecorations
	}
	ForStmt struct {
		Init Stmt
		Cond Expr
		Post Stmt
		Body *BlockStmt
		Decs ForStmtDecorations
	}
	RangeStmt struct {
		Key, Value Expr
		Tok        token.Token
		X          Expr
		Body       *BlockStmt
		Decs       RangeStmtDecorations
	}
)

func (*BadStmt) stmtNode()        {}
func (*DeclStmt) stmtNode()       {}
func (*EmptyStmt) stmtNode()      {}
func (*LabeledStmt) stmtNode()    {}
func (*ExprStmt) stmtNode()       {}
func (*SendStmt) stmtNode()       {}
func (*IncDecStmt) stmtNode()     {}
func (*AssignStmt) stmtNode()     {}
func (*GoStmt) stmtNode()         {}
func (*DeferStmt) stmtNode()      {}
func (*ReturnStmt) stmtNode()     {}
func (*BranchStmt) stmtNode()     {}
func (*BlockStmt) stmtNode()      {}
func (*IfStmt) stmtNode()         {}
func (*CaseClause) stmtNode()     {}
func (*SwitchStmt) stmtNode()     {}
func (*TypeSwitchStmt) stmtNode() {}
func (*CommClause) stmtNode()     {}
func (*SelectStmt) stmtNode()     {}
func (*ForStmt) stmtNode()        {}
func (*RangeStmt) stmtNode()      {}

type (
	Spec interface {
		Node
		specNode()
	}

	ImportSpec struct {
		Name *Ident
		Path *BasicLit
		Decs ImportSpecDecorations
	}

	ValueSpec struct {
		Names  []*Ident
		Type   Expr
		Values []Expr
		Decs   ValueSpecDecorations
	}

	TypeSpec struct {
		Name       *Ident
		TypeParams *FieldList
		Assign     bool
		Type       Expr
		Decs       TypeSpecDecorations
	}
)

func (*ImportSpec) specNode() {}
func (*ValueSpec) specNode()  {}
func (*TypeSpec) specNode()   {}

type (
	BadDecl struct {
		Length int
		Decs   BadDeclDecorations
	}
	GenDecl struct {
		Tok    token.Token
		Lparen bool
		Specs  []Spec
		Rparen bool
		Decs   GenDeclDecorations
	}
	FuncDecl struct {
		Recv *FieldList
		Name *Ident
		Type *FuncType
		Body *BlockStmt
		Decs FuncDeclDecorations
	}
)

func (*BadDecl) declNode()  {}
func (*GenDecl) declNode()  {}
func (*FuncDecl) declNode() {}

type File struct {
	Name       *Ident
	Decls      []Decl
	Scope      *Scope
	Imports    []*ImportSpec
	Unresolved []*Ident
	Decs       FileDecorations
}

type Package struct {
	Name    string
	Scope   *Scope
	Imports map[string]*Object
	Files   map[string]*File
}
