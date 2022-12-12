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
		X    Expr // parenthesized expression
		Decs ParenExprDecorations
	}
	SelectorExpr struct {
		X    Expr   // expression
		Sel  *Ident // field selector
		Decs SelectorExprDecorations
	}
	IndexExpr struct {
		X     Expr // expression
		Index Expr // index expression
		Decs  IndexExprDecorations
	}
	IndexListExpr struct {
		X       Expr // expression
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
		Len  Expr // Ellipsis node for [...]T array types, nil for slice types
		Elt  Expr // element type
		Decs ArrayTypeDecorations
	}
	StructType struct {
		Fields     *FieldList // list of field declarations
		Incomplete bool       // true if (source) fields are missing in the Fields list
		Decs       StructTypeDecorations
	}
	FuncType struct {
		Func       bool
		TypeParams *FieldList // type parameters; or nil
		Params     *FieldList // (incoming) parameters; non-nil
		Results    *FieldList // (outgoing) results; or nil
		Decs       FuncTypeDecorations
	}
	InterfaceType struct {
		Methods    *FieldList // list of embedded interfaces, methods, or types
		Incomplete bool       // true if (source) methods or types are missing in the Methods list
		Decs       InterfaceTypeDecorations
	}
	MapType struct {
		Key   Expr
		Value Expr
		Decs  MapTypeDecorations
	}
	ChanType struct {
		Dir   ChanDir // channel direction
		Value Expr    // value type
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
		Length int // position range of bad statement
		Decs   BadStmtDecorations
	}
	DeclStmt struct {
		Decl Decl // *GenDecl with CONST, TYPE, or VAR token
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
		X    Expr // expression
		Decs ExprStmtDecorations
	}
	SendStmt struct {
		Chan  Expr
		Value Expr
		Decs  SendStmtDecorations
	}
	IncDecStmt struct {
		X    Expr
		Tok  token.Token // INC or DEC
		Decs IncDecStmtDecorations
	}
	AssignStmt struct {
		Lhs  []Expr
		Tok  token.Token // assignment token, DEFINE
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
		Results []Expr // result expressions; or nil
		Decs    ReturnStmtDecorations
	}
	BranchStmt struct {
		Tok   token.Token // keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)
		Label *Ident      // label name; or nil
		Decs  BranchStmtDecorations
	}
	BlockStmt struct {
		List           []Stmt
		RbraceHasNoPos bool // Rbrace may be absent due to syntax error, so we duplicate this in the output for compatibility.
		Decs           BlockStmtDecorations
	}
	IfStmt struct {
		Init Stmt // initialization statement; or nil
		Cond Expr // condition
		Body *BlockStmt
		Else Stmt // else branch; or nil
		Decs IfStmtDecorations
	}
	CaseClause struct {
		List []Expr // list of expressions or types; nil means default case
		Body []Stmt // statement list; or nil
		Decs CaseClauseDecorations
	}
	SwitchStmt struct {
		Init Stmt       // initialization statement; or nil
		Tag  Expr       // tag expression; or nil
		Body *BlockStmt // CaseClauses only
		Decs SwitchStmtDecorations
	}
	TypeSwitchStmt struct {
		Init   Stmt       // initialization statement; or nil
		Assign Stmt       // x := y.(type) or y.(type)
		Body   *BlockStmt // CaseClauses only
		Decs   TypeSwitchStmtDecorations
	}
	CommClause struct {
		Comm Stmt   // send or receive statement; nil means default case
		Body []Stmt // statement list; or nil
		Decs CommClauseDecorations
	}
	SelectStmt struct {
		Body *BlockStmt // CommClauses only
		Decs SelectStmtDecorations
	}
	ForStmt struct {
		Init Stmt // initialization statement; or nil
		Cond Expr // condition; or nil
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
		Name *Ident    // local package name (including "."); or nil
		Path *BasicLit // import path
		Decs ImportSpecDecorations
	}

	ValueSpec struct {
		Names  []*Ident
		Type   Expr
		Values []Expr
		Decs   ValueSpecDecorations
	}

	// A TypeSpec node represents a type declaration (TypeSpec production).
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
		Length int // position range of bad declaration
		Decs   BadDeclDecorations
	}
	GenDecl struct {
		Tok    token.Token // IMPORT, CONST, TYPE, or VAR
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
