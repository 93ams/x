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
	ArrayTypeDecorations struct {
		NodeDecs
		Lbrack Decs
		Len    Decs
	}
	AssignStmtDecorations struct {
		NodeDecs
		Tok Decs
	}
	BadDeclDecorations struct{ NodeDecs }
	BadExprDecorations struct{ NodeDecs }
	BadStmtDecorations struct{ NodeDecs }
	LitDecs            struct{ NodeDecs }
	BinaryDecs         struct {
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
	CaseClauseDecorations struct {
		NodeDecs
		Case  Decs
		Colon Decs
	}
	ChanTypeDecorations struct {
		NodeDecs
		Begin Decs
		Arrow Decs
	}
	CommClauseDecorations struct {
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
	DeclStmtDecorations  struct{ NodeDecs }
	DeferStmtDecorations struct {
		NodeDecs
		Defer Decs
	}
	EllipsisDecorations struct {
		NodeDecs
		Ellipsis Decs
	}
	EmptyStmtDecorations struct{ NodeDecs }
	ExprStmtDecorations  struct{ NodeDecs }
	FieldDecorations     struct {
		NodeDecs
		Type Decs
	}
	FieldListDecorations struct {
		NodeDecs
		Opening Decs
	}
	FileDecorations struct {
		NodeDecs
		Package Decs
		Name    Decs
	}
	ForStmtDecorations struct {
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
	FuncLitDecorations struct {
		NodeDecs
		Type Decs
	}
	FuncTypeDecorations struct {
		NodeDecs
		Func       Decs
		TypeParams Decs
		Params     Decs
	}
	GenDecorations struct {
		NodeDecs
		Tok    Decs
		Lparen Decs
	}
	GoStmtDecorations struct {
		NodeDecs
		Go Decs
	}
	IdentDecorations struct {
		NodeDecs
		X Decs
	}
	IfStmtDecorations struct {
		NodeDecs
		If   Decs
		Init Decs
		Cond Decs
		Else Decs
	}
	ImportDecorations struct {
		NodeDecs
		Name Decs
	}
	IncDecStmtDecorations struct {
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
	InterfaceDecorations struct {
		NodeDecs
		Interface Decs
	}
	KeyValueDecs struct {
		NodeDecs
		Key   Decs
		Colon Decs
	}
	LabeledStmtDecorations struct {
		NodeDecs
		Label Decs
		Colon Decs
	}
	MapTypeDecorations struct {
		NodeDecs
		Map Decs
		Key Decs
	}
	PackageDecorations   struct{ NodeDecs }
	ParenExprDecorations struct {
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
	StarExprDecorations struct {
		NodeDecs
		Star Decs
	}
	StructDecs struct {
		NodeDecs
		Struct Decs
	}
	SwitchStmtDecorations struct {
		NodeDecs
		Switch Decs
		Init   Decs
		Tag    Decs
	}
	TypeAssertExprDecorations struct {
		NodeDecs
		X      Decs
		Lparen Decs
		Type   Decs
	}
	TypeDecorations struct {
		NodeDecs
		Name       Decs
		TypeParams Decs
	}
	TypeSwitchStmtDecorations struct {
		NodeDecs
		Switch Decs
		Init   Decs
		Assign Decs
	}
	UnaryExprDecorations struct {
		NodeDecs
		Op Decs
	}
	ValueDecorations struct {
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
func (n *Array) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Assign) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *BadDecl) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *BadExpr) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *BadStmt) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *Lit) Decorations() *NodeDecs          { return &n.Decs.NodeDecs }
func (n *Binary) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Block) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Branch) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Call) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *CaseClause) Decorations() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *Chan) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *CommClause) Decorations() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *CompositeLit) Decorations() *NodeDecs { return &n.Decs.NodeDecs }
func (n *DeclStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Defer) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Ellipsis) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Empty) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *ExprStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Field) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (l *FieldList) Decorations() *NodeDecs    { return &l.Decs.NodeDecs }
func (n *File) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *For) Decorations() *NodeDecs          { return &n.Decs.NodeDecs }
func (n *Func) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *FuncLit) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *FuncType) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Gen) Decorations() *NodeDecs          { return &n.Decs.NodeDecs }
func (n *Go) Decorations() *NodeDecs           { return &n.Decs.NodeDecs }
func (n *Ident) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *If) Decorations() *NodeDecs           { return &n.Decs.NodeDecs }
func (n *Import) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *IncDec) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Index) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *IndexList) Decorations() *NodeDecs    { return &n.Decs.NodeDecs }
func (n *Interface) Decorations() *NodeDecs    { return &n.Decs.NodeDecs }
func (n *KeyValue) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Labeled) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *MapType) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *Package) Decorations() *NodeDecs      { return nil }
func (n *Paren) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Range) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Return) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Select) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Selector) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *Send) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *Slice) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Star) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *Struct) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Switch) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *TypeAssert) Decorations() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *Type) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *TypeSwitch) Decorations() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *Unary) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Value) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
