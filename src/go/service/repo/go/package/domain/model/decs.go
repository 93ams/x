package model

type (
	SpaceType int
	Decs      []string
	NodeDecs  struct {
		Before SpaceType
		Start  Decs
		End    Decs
		After  SpaceType
	}
	ArrayTypeDecs struct {
		NodeDecs
		Lbrack Decs
		Len    Decs
	}
	AssignDecs struct {
		NodeDecs
		Tok Decs
	}
	BadDeclDecs struct{ NodeDecs }
	BadExprDecs struct{ NodeDecs }
	BadStmtDecs struct{ NodeDecs }
	LitDecs     struct{ NodeDecs }
	BinaryDecs  struct {
		NodeDecs
		X  Decs
		Op Decs
	}
	BlockDecs struct {
		NodeDecs
		Lbrace Decs
	}
	BranchDecs struct {
		NodeDecs
		Tok Decs
	}
	CallDecs struct {
		NodeDecs
		Fun      Decs
		Lparen   Decs
		Ellipsis Decs
	}
	CaseClauseDecs struct {
		NodeDecs
		Case  Decs
		Colon Decs
	}
	ChanTypeDecs struct {
		NodeDecs
		Begin Decs
		Arrow Decs
	}
	CommClauseDecs struct {
		NodeDecs
		Case  Decs
		Comm  Decs
		Colon Decs
	}
	CompositeLitDecs struct {
		NodeDecs
		Type   Decs
		Lbrace Decs
	}
	DeclStmtDecs  struct{ NodeDecs }
	DeferStmtDecs struct {
		NodeDecs
		Defer Decs
	}
	EllipsisDecs struct {
		NodeDecs
		Ellipsis Decs
	}
	EmptyStmtDecs struct{ NodeDecs }
	ExprStmtDecs  struct{ NodeDecs }
	FieldDecs     struct {
		NodeDecs
		Type Decs
	}
	FieldListDecs struct {
		NodeDecs
		Opening Decs
	}
	FileDecs struct {
		NodeDecs
		Package Decs
		Name    Decs
	}
	ForStmtDecs struct {
		NodeDecs
		For  Decs
		Init Decs
		Cond Decs
		Post Decs
	}
	FuncDecs struct {
		NodeDecs
		Func       Decs
		Recv       Decs
		Name       Decs
		TypeParams Decs
		Params     Decs
		Results    Decs
	}
	FuncLitDecs struct {
		NodeDecs
		Type Decs
	}
	FuncTypeDecs struct {
		NodeDecs
		Func       Decs
		TypeParams Decs
		Params     Decs
	}
	GenDecs struct {
		NodeDecs
		Tok    Decs
		Lparen Decs
	}
	GoStmtDecs struct {
		NodeDecs
		Go Decs
	}
	IdentDecs struct {
		NodeDecs
		X Decs
	}
	IfStmtDecs struct {
		NodeDecs
		If   Decs
		Init Decs
		Cond Decs
		Else Decs
	}
	ImportDecs struct {
		NodeDecs
		Name Decs
	}
	IncDecStmtDecs struct {
		NodeDecs
		X Decs
	}
	IndexDecs struct {
		NodeDecs
		X      Decs
		Lbrack Decs
		Index  Decs
	}
	IndexListDecs struct {
		NodeDecs
		X       Decs
		Lbrack  Decs
		Indices Decs
	}
	InterfaceDecs struct {
		NodeDecs
		Interface Decs
	}
	KeyValueDecs struct {
		NodeDecs
		Key   Decs
		Colon Decs
	}
	LabeledStmtDecs struct {
		NodeDecs
		Label Decs
		Colon Decs
	}
	MapTypeDecs struct {
		NodeDecs
		Map Decs
		Key Decs
	}
	PackageDecs   struct{ NodeDecs }
	ParenExprDecs struct {
		NodeDecs
		Lparen Decs
		X      Decs
	}
	RangeStmtDecs struct {
		NodeDecs
		For   Decs
		Key   Decs
		Value Decs
		Range Decs
		X     Decs
	}
	ReturnDecs struct {
		NodeDecs
		Return Decs
	}
	SelectStmtDecs struct {
		NodeDecs
		Select Decs
	}
	SelectorDecs struct {
		NodeDecs
		X Decs
	}
	SendStmtDecs struct {
		NodeDecs
		Chan  Decs
		Arrow Decs
	}
	SliceExprDecs struct {
		NodeDecs
		X      Decs
		Lbrack Decs
		Low    Decs
		High   Decs
		Max    Decs
	}
	StarDecs struct {
		NodeDecs
		Star Decs
	}
	StructDecs struct {
		NodeDecs
		Struct Decs
	}
	SwitchStmtDecs struct {
		NodeDecs
		Switch Decs
		Init   Decs
		Tag    Decs
	}
	TypeAssertDecs struct {
		NodeDecs
		X      Decs
		Lparen Decs
		Type   Decs
	}
	TypeDecs struct {
		NodeDecs
		Name       Decs
		TypeParams Decs
	}
	TypeSwitchStmtDecs struct {
		NodeDecs
		Switch Decs
		Init   Decs
		Assign Decs
	}
	UnaryDecs struct {
		NodeDecs
		Op Decs
	}
	ValueDecs struct {
		NodeDecs
		Assign Decs
	}
)

const (
	None      SpaceType = 0
	NewLine   SpaceType = 1
	EmptyLine SpaceType = 2
)

var Avoid = map[string]bool{
	"Field.Names":   true,
	"Labeled.Label": true,
	"Branch.Label":  true,
	"Import.Name":   true,
	"Value.Names":   true,
	"Type.Name":     true,
	"Func.Name":     true,
	"File.Name":     true,
	"Selector.Sel":  true,
}

func (d *Decs) Append(decs ...string)  { *d = append(*d, decs...) }
func (d *Decs) Prepend(decs ...string) { *d = append(append([]string{}, decs...), *d...) }
func (d *Decs) Replace(decs ...string) { *d = append([]string{}, decs...) }
func (d *Decs) Clear()                 { *d = nil }
func (d *Decs) All() []string          { return *d }
func (s SpaceType) String() string {
	switch s {
	case None:
		return "None"
	case NewLine:
		return "NewLine"
	case EmptyLine:
		return "EmptyLine"
	}
	return ""
}
func (n *Array) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Assign) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *BadDecl) Decors() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *BadExpr) Decors() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *BadStmt) Decors() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *Lit) Decors() *NodeDecs          { return &n.Decs.NodeDecs }
func (n *Binary) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Block) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Branch) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Call) Decors() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *CaseClause) Decors() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *Chan) Decors() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *CommClause) Decors() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *CompositeLit) Decors() *NodeDecs { return &n.Decs.NodeDecs }
func (n *DeclStmt) Decors() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Defer) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Ellipsis) Decors() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Empty) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *ExprStmt) Decors() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Field) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *FieldList) Decors() *NodeDecs    { return &n.Decs.NodeDecs }
func (n *File) Decors() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *For) Decors() *NodeDecs          { return &n.Decs.NodeDecs }
func (n *Func) Decors() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *FuncLit) Decors() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *FuncType) Decors() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Gen) Decors() *NodeDecs          { return &n.Decs.NodeDecs }
func (n *Go) Decors() *NodeDecs           { return &n.Decs.NodeDecs }
func (n *Ident) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *If) Decors() *NodeDecs           { return &n.Decs.NodeDecs }
func (n *Import) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *IncDec) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Index) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *IndexList) Decors() *NodeDecs    { return &n.Decs.NodeDecs }
func (n *Interface) Decors() *NodeDecs    { return &n.Decs.NodeDecs }
func (n *KeyValue) Decors() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Labeled) Decors() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *MapType) Decors() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *Package) Decors() *NodeDecs      { return nil }
func (n *Paren) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Range) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Return) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Select) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Selector) Decors() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Send) Decors() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *Slice) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Star) Decors() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *Struct) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Switch) Decors() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *TypeAssert) Decors() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *Type) Decors() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *TypeSwitch) Decors() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *Unary) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Value) Decors() *NodeDecs        { return &n.Decs.NodeDecs }
