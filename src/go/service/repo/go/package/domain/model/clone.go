package model

func (n *BadExpr) Clone() *BadExpr {
	v := *n
	return &v
}
func (n *Ident) Clone() *Ident {
	v := *n
	v.Obj = nil
	return &v
}
func (n *Ellipsis) Clone() *Ellipsis {
	v := *n
	return &v
}
func (n *Lit) Clone() *Lit {
	v := *n
	return &v
}
func (n *FuncLit) Clone() *FuncLit {
	v := *n
	return &v
}
func (n *CompositeLit) Clone() *CompositeLit {
	v := *n
	return &v
}
func (n *Paren) Clone() *Paren {
	v := *n
	return &v
}
func (n *Selector) Clone() *Selector {
	v := *n
	return &v
}
func (n *Index) Clone() *Index {
	v := *n
	return &v
}
func (n *IndexList) Clone() *IndexList {
	v := *n
	return &v
}
func (n *Slice) Clone() *Slice {
	v := *n
	return &v
}
func (n *TypeAssert) Clone() *TypeAssert {
	v := *n
	return &v
}
func (n *Call) Clone() *Call {
	v := *n
	return &v
}
func (n *Star) Clone() *Star {
	v := *n
	return &v
}
func (n *Unary) Clone() *Unary {
	v := *n
	return &v
}
func (n *Binary) Clone() *Binary {
	v := *n
	return &v
}
func (n *KeyValue) Clone() *KeyValue {
	v := *n
	return &v
}
func (n *Array) Clone() *Array {
	v := *n
	return &v
}
func (n *Struct) Clone() *Struct {
	v := *n
	return &v
}
func (n *FuncType) Clone() *FuncType {
	v := *n
	return &v
}
func (n *Interface) Clone() *Interface {
	v := *n
	return &v
}
func (n *MapType) Clone() *MapType {
	v := *n
	return &v
}
func (n *Chan) Clone() *Chan {
	v := *n
	return &v
}
func (n *BadStmt) Clone() *BadStmt {
	v := *n
	return &v
}
func (n *DeclStmt) Clone() *DeclStmt {
	v := *n
	return &v
}
func (n *Empty) Clone() *Empty {
	v := *n
	return &v
}
func (n *Labeled) Clone() *Labeled {
	v := *n
	return &v
}
func (n *ExprStmt) Clone() *ExprStmt {
	v := *n
	return &v
}
func (n *Send) Clone() *Send {
	v := *n
	return &v
}
func (n *IncDec) Clone() *IncDec {
	v := *n
	return &v
}
func (n *Assign) Clone() *Assign {
	v := *n
	return &v
}
func (n *Go) Clone() *Go {
	v := *n
	return &v
}
func (n *Defer) Clone() *Defer {
	v := *n
	return &v
}
func (n *Return) Clone() *Return {
	v := *n
	return &v
}
func (n *Branch) Clone() *Branch {
	v := *n
	return &v
}
func (n *Block) Clone() *Block {
	v := *n
	return &v
}
func (n *If) Clone() *If {
	v := *n
	return &v
}
func (n *CaseClause) Clone() *CaseClause {
	v := *n
	return &v
}
func (n *Switch) Clone() *Switch {
	v := *n
	return &v
}
func (n *TypeSwitch) Clone() *TypeSwitch {
	v := *n
	return &v
}
func (n *CommClause) Clone() *CommClause {
	v := *n
	return &v
}
func (n *Select) Clone() *Select {
	v := *n
	return &v
}
func (n *For) Clone() *For {
	v := *n
	return &v
}
func (n *Range) Clone() *Range {
	v := *n
	return &v
}
func (n *Import) Clone() *Import {
	v := *n
	return &v
}
func (n *Value) Clone() *Value {
	v := *n
	return &v
}
func (n *Type) Clone() *Type {
	v := *n
	return &v
}
func (n *BadDecl) Clone() *BadDecl {
	v := *n
	return &v
}
func (n *Gen) Clone() *Gen {
	v := *n
	return &v
}
func (n *Func) Clone() *Func {
	v := *n
	return &v
}

func (n *BadExpr) AsExpr() Expr      { return n.Clone() }
func (n *Ident) AsExpr() Expr        { return n.Clone() }
func (n *Ellipsis) AsExpr() Expr     { return n.Clone() }
func (n *Lit) AsExpr() Expr          { return n.Clone() }
func (n *FuncLit) AsExpr() Expr      { return n.Clone() }
func (n *CompositeLit) AsExpr() Expr { return n.Clone() }
func (n *Paren) AsExpr() Expr        { return n.Clone() }
func (n *Selector) AsExpr() Expr     { return n.Clone() }
func (n *Index) AsExpr() Expr        { return n.Clone() }
func (n *IndexList) AsExpr() Expr    { return n.Clone() }
func (n *Slice) AsExpr() Expr        { return n.Clone() }
func (n *TypeAssert) AsExpr() Expr   { return n.Clone() }
func (n *Call) AsExpr() Expr         { return n.Clone() }
func (n *Star) AsExpr() Expr         { return n.Clone() }
func (n *Unary) AsExpr() Expr        { return n.Clone() }
func (n *Binary) AsExpr() Expr       { return n.Clone() }
func (n *KeyValue) AsExpr() Expr     { return n.Clone() }
func (n *Array) AsExpr() Expr        { return n.Clone() }
func (n *Struct) AsExpr() Expr       { return n.Clone() }
func (n *FuncType) AsExpr() Expr     { return n.Clone() }
func (n *Interface) AsExpr() Expr    { return n.Clone() }
func (n *MapType) AsExpr() Expr      { return n.Clone() }
func (n *Chan) AsExpr() Expr         { return n.Clone() }
func (n *BadStmt) AsStmt() Stmt      { return n.Clone() }
func (n *DeclStmt) AsStmt() Stmt     { return n.Clone() }
func (n *Empty) AsStmt() Stmt        { return n.Clone() }
func (n *Labeled) AsStmt() Stmt      { return n.Clone() }
func (n *ExprStmt) AsStmt() Stmt     { return n.Clone() }
func (n *Send) AsStmt() Stmt         { return n.Clone() }
func (n *IncDec) AsStmt() Stmt       { return n.Clone() }
func (n *Assign) AsStmt() Stmt       { return n.Clone() }
func (n *Go) AsStmt() Stmt           { return n.Clone() }
func (n *Defer) AsStmt() Stmt        { return n.Clone() }
func (n *Return) AsStmt() Stmt       { return n.Clone() }
func (n *Branch) AsStmt() Stmt       { return n.Clone() }
func (n *Block) AsStmt() Stmt        { return n.Clone() }
func (n *If) AsStmt() Stmt           { return n.Clone() }
func (n *CaseClause) AsStmt() Stmt   { return n.Clone() }
func (n *Switch) AsStmt() Stmt       { return n.Clone() }
func (n *TypeSwitch) AsStmt() Stmt   { return n.Clone() }
func (n *CommClause) AsStmt() Stmt   { return n.Clone() }
func (n *Select) AsStmt() Stmt       { return n.Clone() }
func (n *For) AsStmt() Stmt          { return n.Clone() }
func (n *Range) AsStmt() Stmt        { return n.Clone() }
func (n *Import) AsSpec() Spec       { return n.Clone() }
func (n *Value) AsSpec() Spec        { return n.Clone() }
func (n *Type) AsSpec() Spec         { return n.Clone() }
func (n *BadDecl) AsDecl() Decl      { return n.Clone() }
func (n *Gen) AsDecl() Decl          { return n.Clone() }
func (n *Func) AsDecl() Decl         { return n.Clone() }
