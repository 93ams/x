package model

type (
	SpaceType   int
	Decorations []string
	NodeDecs    struct {
		Before SpaceType
		Start  Decorations
		End    Decorations
		After  SpaceType
	}
	ArrayTypeDecorations struct {
		NodeDecs
		Lbrack Decorations
		Len    Decorations
	}
	AssignStmtDecorations struct {
		NodeDecs
		Tok Decorations
	}
	BadDeclDecorations    struct{ NodeDecs }
	BadExprDecorations    struct{ NodeDecs }
	BadStmtDecorations    struct{ NodeDecs }
	LitDecs               struct{ NodeDecs }
	BinaryExprDecorations struct {
		NodeDecs
		X  Decorations
		Op Decorations
	}
	BlockDecorations struct {
		NodeDecs
		Lbrace Decorations
	}
	BranchDecorations struct {
		NodeDecs
		Tok Decorations
	}
	CallDecorations struct {
		NodeDecs
		Fun      Decorations
		Lparen   Decorations
		Ellipsis Decorations
	}
	CaseClauseDecorations struct {
		NodeDecs
		Case  Decorations
		Colon Decorations
	}
	ChanTypeDecorations struct {
		NodeDecs
		Begin Decorations
		Arrow Decorations
	}
	CommClauseDecorations struct {
		NodeDecs
		Case  Decorations
		Comm  Decorations
		Colon Decorations
	}
	CompositeLitDecs struct {
		NodeDecs
		Type   Decorations
		Lbrace Decorations
	}
	DeclStmtDecorations  struct{ NodeDecs }
	DeferStmtDecorations struct {
		NodeDecs
		Defer Decorations
	}
	EllipsisDecorations struct {
		NodeDecs
		Ellipsis Decorations
	}
	EmptyStmtDecorations struct{ NodeDecs }
	ExprStmtDecorations  struct{ NodeDecs }
	FieldDecorations     struct {
		NodeDecs
		Type Decorations
	}
	FieldListDecorations struct {
		NodeDecs
		Opening Decorations
	}
	FileDecorations struct {
		NodeDecs
		Package Decorations
		Name    Decorations
	}
	ForStmtDecorations struct {
		NodeDecs
		For  Decorations
		Init Decorations
		Cond Decorations
		Post Decorations
	}
	FuncDecs struct {
		NodeDecs
		Func       Decorations
		Recv       Decorations
		Name       Decorations
		TypeParams Decorations
		Params     Decorations
		Results    Decorations
	}
	FuncLitDecorations struct {
		NodeDecs
		Type Decorations
	}
	FuncTypeDecorations struct {
		NodeDecs
		Func       Decorations
		TypeParams Decorations
		Params     Decorations
	}
	GenDecorations struct {
		NodeDecs
		Tok    Decorations
		Lparen Decorations
	}
	GoStmtDecorations struct {
		NodeDecs
		Go Decorations
	}
	IdentDecorations struct {
		NodeDecs
		X Decorations
	}
	IfStmtDecorations struct {
		NodeDecs
		If   Decorations
		Init Decorations
		Cond Decorations
		Else Decorations
	}
	ImportDecorations struct {
		NodeDecs
		Name Decorations
	}
	IncDecStmtDecorations struct {
		NodeDecs
		X Decorations
	}
	IndexExprDecorations struct {
		NodeDecs
		X      Decorations
		Lbrack Decorations
		Index  Decorations
	}
	IndexListExprDecorations struct {
		NodeDecs
		X       Decorations
		Lbrack  Decorations
		Indices Decorations
	}
	InterfaceDecorations struct {
		NodeDecs
		Interface Decorations
	}
	KeyValueDecs struct {
		NodeDecs
		Key   Decorations
		Colon Decorations
	}
	LabeledStmtDecorations struct {
		NodeDecs
		Label Decorations
		Colon Decorations
	}
	MapTypeDecorations struct {
		NodeDecs
		Map Decorations
		Key Decorations
	}
	PackageDecorations   struct{ NodeDecs }
	ParenExprDecorations struct {
		NodeDecs
		Lparen Decorations
		X      Decorations
	}
	RangeStmtDecorations struct {
		NodeDecs
		For   Decorations
		Key   Decorations
		Value Decorations
		Range Decorations
		X     Decorations
	}
	ReturnDecs struct {
		NodeDecs
		Return Decorations
	}
	SelectStmtDecorations struct {
		NodeDecs
		Select Decorations
	}
	SelectorDecs struct {
		NodeDecs
		X Decorations
	}
	SendStmtDecorations struct {
		NodeDecs
		Chan  Decorations
		Arrow Decorations
	}
	SliceExprDecorations struct {
		NodeDecs
		X      Decorations
		Lbrack Decorations
		Low    Decorations
		High   Decorations
		Max    Decorations
	}
	StarExprDecorations struct {
		NodeDecs
		Star Decorations
	}
	StructDecs struct {
		NodeDecs
		Struct Decorations
	}
	SwitchStmtDecorations struct {
		NodeDecs
		Switch Decorations
		Init   Decorations
		Tag    Decorations
	}
	TypeAssertExprDecorations struct {
		NodeDecs
		X      Decorations
		Lparen Decorations
		Type   Decorations
	}
	TypeDecorations struct {
		NodeDecs
		Name       Decorations
		TypeParams Decorations
	}
	TypeSwitchStmtDecorations struct {
		NodeDecs
		Switch Decorations
		Init   Decorations
		Assign Decorations
	}
	UnaryExprDecorations struct {
		NodeDecs
		Op Decorations
	}
	ValueDecorations struct {
		NodeDecs
		Assign Decorations
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

func (d *Decorations) Append(decs ...string)  { *d = append(*d, decs...) }
func (d *Decorations) Prepend(decs ...string) { *d = append(append([]string{}, decs...), *d...) }
func (d *Decorations) Replace(decs ...string) { *d = append([]string{}, decs...) }
func (d *Decorations) Clear()                 { *d = nil }
func (d *Decorations) All() []string          { return *d }
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
