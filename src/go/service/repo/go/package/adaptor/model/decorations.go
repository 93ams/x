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
	BasicLitDecorations   struct{ NodeDecs }
	BinaryExprDecorations struct {
		NodeDecs
		X  Decorations
		Op Decorations
	}
	BlockStmtDecorations struct {
		NodeDecs
		Lbrace Decorations
	}
	BranchStmtDecorations struct {
		NodeDecs
		Tok Decorations
	}
	CallExprDecorations struct {
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
	CompositeLitDecorations struct {
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
	FuncDeclDecorations struct {
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
	GenDeclDecorations struct {
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
	ImportSpecDecorations struct {
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
	InterfaceTypeDecorations struct {
		NodeDecs
		Interface Decorations
	}
	KeyValueExprDecorations struct {
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
	ReturnStmtDecorations struct {
		NodeDecs
		Return Decorations
	}
	SelectStmtDecorations struct {
		NodeDecs
		Select Decorations
	}
	SelectorExprDecorations struct {
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
	StructTypeDecorations struct {
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
	TypeSpecDecorations struct {
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
	ValueSpecDecorations struct {
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
	"Field.Names":       true,
	"LabeledStmt.Label": true,
	"BranchStmt.Label":  true,
	"ImportSpec.Name":   true,
	"ValueSpec.Names":   true,
	"TypeSpec.Name":     true,
	"FuncDecl.Name":     true,
	"File.Name":         true,
	"SelectorExpr.Sel":  true,
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
func (n *ArrayType) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *AssignStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *BadDecl) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *BadExpr) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *BadStmt) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *BasicLit) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *BinaryExpr) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *BlockStmt) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *BranchStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *CallExpr) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *CaseClause) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *ChanType) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *CommClause) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *CompositeLit) Decorations() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *DeclStmt) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *DeferStmt) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *Ellipsis) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *EmptyStmt) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *ExprStmt) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *Field) Decorations() *NodeDecs          { return &n.Decs.NodeDecs }
func (l *FieldList) Decorations() *NodeDecs      { return &l.Decs.NodeDecs }
func (n *File) Decorations() *NodeDecs           { return &n.Decs.NodeDecs }
func (n *ForStmt) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *FuncDecl) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *FuncLit) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *FuncType) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *GenDecl) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *GoStmt) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *Ident) Decorations() *NodeDecs          { return &n.Decs.NodeDecs }
func (n *IfStmt) Decorations() *NodeDecs         { return &n.Decs.NodeDecs }
func (n *ImportSpec) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *IncDecStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *IndexExpr) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *IndexListExpr) Decorations() *NodeDecs  { return &n.Decs.NodeDecs }
func (n *InterfaceType) Decorations() *NodeDecs  { return &n.Decs.NodeDecs }
func (n *KeyValueExpr) Decorations() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *LabeledStmt) Decorations() *NodeDecs    { return &n.Decs.NodeDecs }
func (n *MapType) Decorations() *NodeDecs        { return &n.Decs.NodeDecs }
func (n *Package) Decorations() *NodeDecs        { return nil }
func (n *ParenExpr) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *RangeStmt) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *ReturnStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *SelectStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *SelectorExpr) Decorations() *NodeDecs   { return &n.Decs.NodeDecs }
func (n *SendStmt) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *SliceExpr) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *StarExpr) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *StructType) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *SwitchStmt) Decorations() *NodeDecs     { return &n.Decs.NodeDecs }
func (n *TypeAssertExpr) Decorations() *NodeDecs { return &n.Decs.NodeDecs }
func (n *TypeSpec) Decorations() *NodeDecs       { return &n.Decs.NodeDecs }
func (n *TypeSwitchStmt) Decorations() *NodeDecs { return &n.Decs.NodeDecs }
func (n *UnaryExpr) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
func (n *ValueSpec) Decorations() *NodeDecs      { return &n.Decs.NodeDecs }
